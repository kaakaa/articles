---
title: "Mattermost Apps Frameworkを触ってみた"
emoji: "🔌"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: ["mattermost", "integrations"]
date: 2021-05-08T00:00:00+09:00
published: true
aliases: "/post/mattermost/apps-sample"
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

# はじめに

Mattermost の新たな統合機能である Apps Framework が Developer Preview になりました。

https://twitter.com/Mattermost/status/1390333864820813831

既存の統合機能やプラグイン機能との違いを確かめるべく、公式開発者向けドキュメントにある [Quick start guide \(Go\)](https://developers.mattermost.com/integrate/apps/quick-start-go/) を動かしてみました。

注: Mattermost Apps Framework は、記事執筆時点では Developer Preview 段階のためプロダクションリリースには含まれていません。Apps Framework を利用するには、Mattermost 開発環境を用意する必要があります。

# Mattermost Apps の開発・インストール

Mattermost Apps は、Mattermost インスタンスとは別の Mattermot Apps 用のアプリケーションサーバーとして起動します。

ここでは、そのアプリケーションサーバーの開発方法と、Mattermost Apps のインストール・利用方法について書いています。以下にある内容は、公式ドキュメントの[Quick start guide \(Go\)](https://developers.mattermost.com/integrate/apps/quick-start-go/)の内容に沿ったものです。

## 0. 環境設定

Mattermost Apps Framework は、まだ Developer Preview 段階ということで、Mattermost のプロダクションリリースには含まれていません。  
Mattermost Apps Framework を動かすには、[Mattermost 開発環境](https://developers.mattermost.com/contribute/webapp/developer-setup/)を構築した上で、以下の環境変数を設定してから Mattermost 開発環境を起動する必要があります。

```
$ export MM_FEATUREFLAGS_AppsEnabled=true
```

また、Mattermost Apps Framework は Bot アカウント機能と OAuth 2.0 Service Provider 機能を利用しています。そのため、起動した Mattermost インスタンスの**システムコンソール > 統合機能 > 統合機能管理**から、これらの機能を有効化しておく必要があります。

## 1. Mattermpst Apps Plugin のインストール

Mattermost Apps は [mattermost-plugin-apps](https://github.com/mattermost/mattermost-plugin-apps) で管理されるため、まず [mattermost-plugin-apps](https://github.com/mattermost/mattermost-plugin-apps) をインストールする必要があります。

以下のコマンドで `mattermost-plugin-apps` を Mattermost インスタンスにインストールします。

```bash:install mattermost-plugin-apps
$ git clone https://github.com/mattermost/mattermost-plugin-apps.git
$ cd mattermost-plugin-apps
$ export MM_SERVICESETTINGS_SITEURL=http://localhost:8065
$ export MM_ADMIN_USERNAME=kaaka
$ export MM_ADMIN_PASSWORD=PASSWORD
$ make deploy
```

環境変数に指定しているのは、Mattermost インスタンス の SiteURL と、管理者権限を持つユーザーのユーザー名/パスワードです。

## 2. サンプル Mattermost Apps の作成

ここまでで Mattermost Apps を動作させる準備ができました。ここからは Mattermost Apps 本体を開発していきます。

Mattermost Apps サーバーと Mattermost インスタンスは、(恐らく)以下のように連携して動作しています。(以下のシーケンス図は、[公式ドキュメントの図](https://developers.mattermost.com/integrate/apps/lifecycle/)に少し手を加えたものです)

![](https://storage.googleapis.com/zenn-user-upload/e5pgz2kistlm5rwjowf3rodfciv0)

一番右にある `Apps Server` が、これから開発する Mattermost Apps 本体です。

シーケンス図の一番下の `post bot message` の部分に関しては、 Mattermost Apps から Mattermost にメッセージを投稿するロジックを書く必要がありますが、それ以外の部分は Mattermost からのリクエストに対して JSON ファイルを返しているだけです。

---

今回は Go 言語で Mattermost Apps を開発するため、`go mod init` で Go プロジェクトを作成し、`github.com/mattermost/mattermost-plugin-apps/apps`を依存関係に加えます。（`github.com/mattermost/mattermost-plugin-apps/apps` には、Mattermost Apps を開発する際に利用可能なユーティリティライブラリが含まれています）
）

```
$ go mod init
$ go get github.com/mattermost/mattermost-plugin-apps/apps@master
```

以下のドキュメントにもあるように、`github.com/mattermost/mattermost-plugin-apps/apps@master`を通じて、Mattermost REST API を実行したり、Apps 用の Key-Value Store を利用できるようになります。

[Apps APIs](https://developers.mattermost.com/integrate/apps/using-mattermost-api/)

### `manifest.json`

Go プロジェクトのルートディレクトリに以下の内容で`manifest.json`というファイルを作成します。

```json
{
  "app_id": "hello-world",
  "display_name": "Hello, world!",
  "app_type": "http",
  "root_url": "http://localhost:8080",
  "requested_permissions": ["act_as_bot"],
  "requested_locations": ["/channel_header", "/command"]
}
```

この`manifest.json`は、Mattermost Apps の概要や種別、Apps が要求する Permission、Apps が表示される箇所などを表すファイルです。

- `app_id`、`display_name`はそれぞれ Apps の ID と表示名を表します。`app_id`は Apps をインストールする際に使用されます。
- `app_type`は Apps の種別を表します。`http`、`aws_lambda`、`builtin`の中から 1 つを選ぶようですが、`http`以外を指定した場合の動作についてはまだわかっていません。
- `app_type`に`http`を指定した場合には、`root_url`に Apps が起動しているサーバーの URL を記述します。
- `requested_permissions`には、Apps の動作に必要な権限を指定します。 記述できる内容については[Permissions](https://developers.mattermost.com/integrate/apps/manifest/#permissions)を参照。
- `requested_locations`には、Apps が干渉する Mattermost UI 上の箇所を指定します。次の`bindings.json`の記述内容と関連します。 記述できる内容については[Locations](https://developers.mattermost.com/integrate/apps/manifest/#locations)を参照。

Manifest ファイルに記述できる内容については以下の公式ドキュメントで紹介されています。

[Manifest](https://developers.mattermost.com/integrate/apps/manifest/)

### `bindings.json`

次に、Mattermost 上のどこに・どんな機能を配置するかについて記述した`bindings.json`というファイルを以下の内容で作成します。

```json
{
  "type": "ok",
  "data": [
    {
      "location": "/channel_header",
      "bindings": [
        {
          "location": "send-button",
          "icon": "http://localhost:8080/static/icon.png",
          "label": "send hello message",
          "call": {
            "path": "/send-modal"
          }
        }
      ]
    },
    {
      "location": "/command",
      "bindings": [
        {
          "icon": "http://localhost:8080/static/icon.png",
          "label": "helloworld",
          "description": "Hello World app",
          "hint": "[send]",
          "bindings": [
            {
              "location": "send",
              "label": "send",
              "call": {
                "path": "/send"
              }
            }
          ]
        }
      ]
    }
  ]
}
```

`bindings.json`には、`manifest.json`の`requested_locations`に記述した`locations`に対する詳細な定義を記述していきます。
例えば、`"location": "/channel_header"`は、Mattermost のチャンネルヘッダー部分に表示されるボタンを表しています。`bindings.json`には、このチャンネルヘッダーに表示されるボタンの概要(アイコン、ラベル)や、ボタンが押された時の動作を定義します。

```json
...
		{
			"location": "/channel_header",
			"bindings": [
				{
					"location": "send-button",
					"icon": "http://localhost:8080/static/icon.png",
					"label":"send hello message",
					"call": {
						"path": "/send-modal"
					}
				}
			]
		},
...
```

上記の定義により、`http://localhost:8080/static/icon.png`をアイコンとし、`send hello message`をラベルとするボタンがチャンネルヘッダーに表示されます。

![](https://storage.googleapis.com/zenn-user-upload/l5i8afh19drgtphwwd7js1qkk9bg)

このボタンを押した時の動作は`bindings.call`に書かれます。

```json
"call": {
	"path": "/send-modal"
}
```

この定義により、`manifest.json`の`root_url`に書かれた URL をベースとした`/send-modal`のパス、すなわち`http://localhost:8080/send-modal/submit`に Mattermost からのリクエストが送信されます。(末尾の`/submit`が何故つくのか分かっていませんが、サンプルコード上では`/submit`が末尾に付いた URL へリクエストが送信されているようです)

この`bindings.call`には他にも様々なパラメータを指定でき、指定したパラメータの内容によって Mattermost から様々な情報を Apps に対して送信するよう要求できます。

[Call](https://developers.mattermost.com/integrate/apps/call/)

### Modal Forms

前述の `http://localhost:8080/send-modal/submit` へ送信されたリクエストに対しては、今回のサンプル Apps ではモーダルを表示するための JSON ファイルを返却しています。(この Go プログラムの全容は以降のセクションで紹介しています)

```go:main.go
...
//go:embed send_form.json
var formData []byte

func main() {
...
	http.HandleFunc("/send-modal/submit", writeJSON(formData))
...
```

この時、返却する JSON の内容は下記の通りです。

```JSON:send_form.json
{
	"type": "form",
	"form": {
		"title": "Hello, world!",
		"icon": "http://localhost:8080/static/icon.png",
		"fields": [
			{
				"type": "text",
				"name": "message",
				"label": "message"
			}
		],
		"call": {
			"path": "/send"
		}
	}
}
```

Mattermost が、上記の JSON を Apps サーバーから受け取ると、以下のようなモーダルウィンドウを表示します。

![](https://storage.googleapis.com/zenn-user-upload/jdztut38p9pegb9kwns2cbbo8hqk)

返却した JSON の`form`フィールドの内容が、表示されるモーダルの内容となります。  
今回は`fields`配列に`"type": "text"`のオブジェクト 1 つしかないため、1 つのテキストボックスのみを持つモーダルが表示されました。`fields`内には他にも`text`、`static_select`、`dynamic_select`、`bool`、`user`、`channel`など、様々なタイプの要素を指定できます。

詳細については以下のドキュメントに記載があります。

[Interactivity](https://developers.mattermost.com/integrate/apps/interactivity/#modal-forms)

モーダルの送信ボタンを押した場合、`send_form.json`の`form.call`に書かれたパス (`http://localhost:8080/send/submit`) へ Mattermost からリクエストが送信されます。（この場合も末尾に`/submit`が自動で付与されるようです）

送信されたリクエストは、再び Apps サーバーで扱われます。Apps サーバーの内容は、後述のセクションで紹介します。

### アイコンファイル

上記で作成してきた JSON ファイル内で参照されているアイコンファイルを取得しておきます。
公式ドキュメントでは、以下のコマンドでダウンロードするよう紹介されていますが、大きすぎるファイルでなければ恐らくどんな画像ファイルでも大丈夫です。

```
$ wget https://github.com/mattermost/mattermost-plugin-apps/raw/master/examples/go/helloworld/icon.png
```

### Go サーバーアプリケーション

今回開発している Apps サーバーの本体であるサーバーアプリケーションは、下記の Go コードになります。Go 1.16 から導入された[Go embed](https://golang.org/pkg/embed/)を使って、今まで作成してきた JSON ファイルを直接読み込んでいます。

```go:main.go
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
)

//go:embed icon.png
var iconData []byte

//go:embed manifest.json
var manifestData []byte

//go:embed bindings.json
var bindingsData []byte

//go:embed send_form.json
var formData []byte

func main() {
	http.HandleFunc("/manifest.json", writeJSON(manifestData))
	http.HandleFunc("/bindings", writeJSON(bindingsData))
	http.HandleFunc("/send/form", writeJSON(formData))
	http.HandleFunc("/send/submit", send)
	http.HandleFunc("/send-modal/submit", writeJSON(formData))
	http.HandleFunc("/static/icon.png", writeData("image/png", iconData))

	http.ListenAndServe(":8080", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	c := apps.CallRequest{}
	json.NewDecoder(r.Body).Decode(&c)

	message := "Hello, world!"
	v, ok := c.Values["message"]
	if ok && v != nil {
		message += fmt.Sprintf(" ... and %s!", v)
	}
	mmclient.AsBot(c.Context).DM(c.Context.ActingUserID, message)

	json.NewEncoder(w).Encode(apps.CallResponse{})
}

func writeData(ct string, data []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", ct)
		w.Write(data)
	}
}

func writeJSON(data []byte) func(w http.ResponseWriter, r *http.Request) {
	return writeData("application/json", data)
}
```

前述のモーダルウィンドウの送信ボタンをした際に Mattermost から送信される `http://localhost:8080/send/submit` へのリクエストは、上記の`send`関数で処理されます。`send`関数では、モーダルの入力内容を元にメッセージを作成し、`github.com/mattermost/mattermost-plugin-apps/apps/mmclient`の関数を使って、Bot として Mattermost にダイレクトメッセージを投稿しています。

```go
...
func send(w http.ResponseWriter, r *http.Request) {
	c := apps.CallRequest{}
	json.NewDecoder(r.Body).Decode(&c)

	message := "Hello, world!"
	v, ok := c.Values["message"]
	if ok && v != nil {
		message += fmt.Sprintf(" ... and %s!", v)
	}
	mmclient.AsBot(c.Context).DM(c.Context.ActingUserID, message)

	json.NewEncoder(w).Encode(apps.CallResponse{})
}
...
```

以上で Mattermost Apps サーバーの開発は完了です。早速 Mattermost Apps をインストールしてみましょう。

## 3. サンプル Mattermost Apps のインストール

`go run .`コマンドで Mattermost Apps サーバーを立ち上げます。

次に、`mattermost-plugin-apps`をインストールした Mattermost をブラウザで開き、`/apps`コマンドを使って Mattermost Apps をインストールします。

```
/apps debug-add-manifest --url http://localhost:8080/manifest.json
/apps install hello-world
```

`/apps debug-add-manifest`で、Mattermost Apps サーバーから`manifest.json`を取得し、その内容を検証します。問題がなければ`/apps install`コマンドで Apps ID を指定してインストールを実行します。

インストール実行時に、インストールしようとしている Apps が要求する Permission と Locations が表示され、App Secret の入力を促されます。App Secret の内容はなんでも良いらしいです。

![](https://storage.googleapis.com/zenn-user-upload/bwq66dapqo4tffu2m40jvdv3pltj)

App Secret を入力して、`Approve and Install`ボタンをクリックすると、Apps のインストールが完了します。Apps をインストールした後、ページをリロードすると、チャンネルヘッダー部分にボタンが表示されているはずです。

![](https://storage.googleapis.com/zenn-user-upload/l5i8afh19drgtphwwd7js1qkk9bg)

また、Apps インストール後に`/apps list`コマンドを実行すると、インストールされている Mattermost Apps の一覧を確認できます。

![](https://storage.googleapis.com/zenn-user-upload/wgrkbyx67xra1t5jlj565rzsjzyt)

## 4. サンプル Mattermost Apps の実行

チャンネルヘッダー部分のボタンをクリックすると、モーダルウィンドウが開きます。

![](https://storage.googleapis.com/zenn-user-upload/jdztut38p9pegb9kwns2cbbo8hqk)

メッセージを入力し、送信ボタンを押すと、Bot からダイレクトメッセージが届きます。

![](https://storage.googleapis.com/zenn-user-upload/uqxdln0cf4rhuo42w8zy5lxlu3pj)

# さいごに

Developer Preview 版として公開された Mattermost Apps Framework を動かしてみました。

Mattermost Apps Framework は、既存の統合機能(WebHook、Slash Command)とプラグインの中間のような位置付けのものという印象を受けましたが、使用できるパラメータの数も多く、どのような場面で使うべきかはもう少し触ってみないとなんとも言えません。
ただ、AWS Lambda と連携するタイプの Apps が現段階で利用できることから、FaaS 基盤と連携可能という点がポイントになるような気がしています。実際、Apps のインストールや Mattermost UI に関わる部分は Apps サーバーから JSON を返すだけで実現できるため、フルスタックなサーバーアプリケーションを立てておくことなく、必要な部分の処理だけをコードとして書けば良いという形式になっています。この点より、プラグインより軽量であり、かつプラグインのように Mattermost と密接に連携した機能を開発できる基盤なのかなという印象があります。また、FaaS 的な基盤として[n8n](https://n8n.io/)なんかと上手く連携する方法が見つかルト、いろいろ幅が広がりそうです。

以下は妄想です。

今回実装したようなサンプル Apps ぐらいなら、[プラグイン](https://developers.mattermost.com/extend/plugins/)として実装できますが、プラグインは Mattermost インスタンスと密接に連携して動作するため、Mattermost インスタンスの動作に意図しない影響を与えてしまう可能性もあります。また、マルチクラスタ構成を組んでいる場合、プラグインによる機能追加を全てのインスタンスで同一の状態に揃えることが難しいということも、背景としてあったのでは無いかと考えられます（妄想ですが）。Mattermost プラグイン自体は Mattermost を自組織向けにカスタマイズするための強力な仕組みですが、やはり統合機能は別のサーバーで管理した方が利用しやすいという考えもあったのではないかと思われます。ただ、既存の統合機能では HTTP リクエストのやり取りぐらいしかできないため、もう少し踏み込んだ統合機能の実装として Mattermost Apps Framework が生まれたのではないかと想像しています。
