---
title: "QiitaからZennに移行する"
emoji: "👋"
type: "idea" # tech: 技術記事 / idea: アイデア
topics: ["qiita", "zenn"]
published: true
date: 2021-06-19T00:00:00+09:00
---

2023/02/21 掲載しているコードの内、Qiitaの制限に抵触する部分を修正しました。(詳細は[コメント](https://zenn.dev/link/comments/991082a51e4149)参照)

---

ここ数年、Qiita は「アップデートしました」と言いつつ使い勝手は何も変わってなかったり、ログインするたび興味もないキャンペーンの通知を送りつけてくるなど体験があまり良くないので Zenn に移ろうと思いました。

## 1. Qiita の記事を Zenn に移行する

Qiita の記事を Zenn に移すのは、以下の記事で紹介されている [ikawaha/zenn-importer](https://github.com/ikawaha/zenn-importer) を実行するだけでした。

[Qiita/はてブの記事を Zenn でも管理する](https://zenn.dev/ikawaha/articles/20201012-e56b19cd33c396ae0465)

Qiita API はユーザーを指定するだけで公開記事の一覧とその内容をダウンロードできるようなので、Qiita のアクセストークンの発行なども必要なく、記事を一括でダウンロードできました。

## 2. Qiita の記事の内容を Zenn へのリダイレクトに書き換える

必須ではないけど、Qiita と Zenn に同じ記事が残っているのは少しお行儀が悪く感じるので、Qiita の記事内容を Zenn の記事へのリダイレクトに書き換えてます。
移行先の Zenn の記事の URL は、`1.`の手順で使った [ikawaha/zenn-importer](https://github.com/ikawaha/zenn-importer) のREADMEに書いてある規則に従っているため、リダイレクト先のURLも機械的に生成可能。

一括で書き換えるためのスクリプトとして下記を書いた。(記事内容を書き換えるため、こちらの実行には流石に Qiita のアクセストークンが必要)
古すぎる記事は移行するまでもないなと思い、最初に書いたMattermost関連の記事以降の記事を移行対象としたので、`QiitaPerPage = 74`という中途半端な数字になっている。

Qiitaへのアクセス用に`QiitaAccessToken`と`QiitaUser`と、どの記事までを移行対象にするかを決定するための定数`QiitaPerPage`、`QiitaMaxPage`あたりを書き換えて実行するだけで、全ての記事の内容が Zenn 記事へのリダイレクトへ書き換わる。(自分用のスクリプトなので動作は保証しません)

```go:main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	QiitaAccessToken = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	QiitaUser        = "kaakaa_hoe"
	QiitaPerPage     = 74
	QiitaMaxPage = 1
)

type QiitaItem struct {
	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Private   bool   `json:"private"`
}

type QiitaPatch struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func main() {
	client := &http.Client{}

	var all []QiitaItem
	page := 1
	for {
		log.Printf("Get items. page %d", page)
		items, err := getItems(client, page)
		if err != nil {
			log.Fatal(err)
		}
		all = append(all, items...)
		if len(items) < QiitaPerPage {
			break
		}
		log.Println(len(items))
		page += 1
		if QiitaMaxPage < page {
			break
		}
	}

	nums := len(all)
	for i, v := range all {
		if v.Private {
			// skip private article
			continue
		}
		log.Printf("Patching article(%d/%d): %s", i+1, nums, v.Title)
		if err := patchItem(client, v); err != nil {
			log.Printf(" [ERROR] %s", err.Error())
		}
	}
}

func getItems(client *http.Client, page int) ([]QiitaItem, error) {
	req, err := http.NewRequest(http.MethodGet, "https://qiita.com/api/v2/authenticated_user/items", nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	params.Add("page", strconv.Itoa(page))
	params.Add("per_page", fmt.Sprintf("%d", QiitaPerPage))
	params.Add("query", fmt.Sprintf("user:%s", QiitaUser))
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", QiitaAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get items %s", resp.Status)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var items []QiitaItem
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func patchItem(client *http.Client, item QiitaItem) error {
	t, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://zenn.dev/kaakaa/articles/qiita-%s-%s", t.Format("20060102"), item.Id)

	b, err := json.Marshal(QiitaPatch{
		// 短い内容で「移行」という言葉が含まれるとQiitaに弾かれるため文言を修正
		// Body:  fmt.Sprintf("この記事は Zenn に移行しました。\n%s", url),
		Body:  fmt.Sprintf("記事の内容は以下を参照してください。\n%s", url),
		Title: item.Title,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("https://qiita.com/api/v2/items/%s", item.Id), bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", QiitaAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to request patch: %s", resp.Status)
	}
	return nil
}
```

上記のスクリプトを実行することでQiitaの記事を一括で書き換えることができました。

![](https://storage.googleapis.com/zenn-user-upload/82d211bdec1c67cbcee986c9.png)

### 移行記事の投稿日時について

移行するときにネックとなった点。
ZennではFrontmatterで公開日時などを指定できないようなので、Qiita上では古い記事でもZenn上では最新の記事として扱われてしまう。
それによる影響は、自分のアカウントの記事一覧が古いもので埋められてしまうのと、タグ(ZennではTopics)が古い記事で汚染されてしまうこと。

前者については、一旦記事を全部投稿してから最新にしたい記事へ更新を加えることで記事の順序を操作できるかと思ったけど、そんなことはなく投稿日時で決まってしまうようだった。ここは諦めるしかないかな。

後者については自分以外のユーザーまで影響が及んでしまう部分だけど、自分が書いているのは Mattermost の記事が多く、Mattermotは幸い (?) まだZenn上での記事の数が多くないので目を瞑ることにした。閲覧者が多そうなTopicsをつけてしまってる場合は、そのTopicsを外しておいた。

Frontmatterに`date`を設定できるよう要望を出そうかとも思ったけど、今は[`date`が設定されていても無視するという動作をしているよう](https://github.com/zenn-dev/zenn-community/issues/254)なので、`date`相当の情報を有効にするという変更は既存の記事に対する影響がとても大きそうなのであまり良い手じゃないなと思って要望を出すのはやめました。
