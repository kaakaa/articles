---
title: "NotionのようなUIのTrelloっぽいKanbanツールのOSSの Focalboard を触ってみた"
emoji: "🗒"
type: "tech"
topics: [focalboard, oss, mattermost, trello, notion]
published: true
---

# はじめに

Slack Alternative な OSS のチャットツールを開発している [Mattermost](https://mattermost.com)社の [GitHub Organization](https://github.com/mattermost/focalboard) に、[Focalboard](https://github.com/mattermost/focalboard)という[Trello](https://trello.com/ja)・[Notion](https://notion.so)・[Asana](https://asana.com/ja) Alternative を名乗るリポジトリが作成されていたため、どのようなことができるのか動かして確かめてみました。

https://github.com/mattermost/focalboard

![screenshot](https://storage.googleapis.com/zenn-user-upload/6z7g98ykul3bc2nebx19m14dpkf8)

Focalboard は、記事執筆時点での最新バージョンが`v0.6.1`のため、現在 **early-access beta** の段階であり、動作は保証されていないものと思われます。

# Focalboard とは

https://www.focalboard.com/

Focalboard は、OSS で開発されているセルフホスト可能な Trello/Notion/Asana Alternative と紹介されています。

> Focalboard is an open source, self-hosted alternative to Trello, Notion, and Asana.

`Folcalboard Personal Desktop`という個人で使うデスクトップ版と、チームコラボレーション用の`Focalboard Personal Server`という 2 つの動作形態があります。それぞれ動かして試してみました。

## Focalboard Desktop

### インストール

デスクトップ版は公式サイトからダウンロードすることができます。
https://www.focalboard.com/download/personal-edition/desktop/

macOS / Windows / Linux Desktop 版があり、今回は macOS 版をダウンロードします。
（リポジトリを見ると、Electron 等を使用している訳ではなく、各 OS 向けのビルド用資産があるようです。さすがにコア機能は共通化されてそうですが。）

ダウンロードされた ZIP ファイルを解凍し、中にある`Focalboard.app`をダブルクリックするだけで起動しました。

![screen shot](https://storage.googleapis.com/zenn-user-upload/1ofl6ere2xvtgncsgtv5vrm6rcht)

### ボードを作成してみる

画面左下の`+ Add Board`というボタンをクリックすると、作成できるボードのテンプレートが表示されます。

![templates](https://storage.googleapis.com/zenn-user-upload/5mrqyaeq792r14gvqhqjm2sywh69 =250x)

とりあえず一番上の **Meeting Notes** を作成してみます。

![meeting notes](https://storage.googleapis.com/zenn-user-upload/jseqiardm3b6lmkfk1w1say2an43 =800x)

Notion っぽい...。

タイトル部分は絵文字含めて変更できます。
(`Meeting Notes` = 議事録かと思いましたが、このボードは議事録ではないですね...)

![edit header](https://storage.googleapis.com/zenn-user-upload/wwllki3bigkxegyg72dmz4fy0c2x =800x)

### カードを作成してみる

新しいカード作成画面。(Notion だ...)
カードの内容は Markdown で書くことができ、ブロックエディタのように記述したブロックの移動を行うこともできます。（Notion ほどの自由度はなさそうですが）

![card body](https://storage.googleapis.com/zenn-user-upload/mb00hesd5tv353w3tffvqm44pw2r)

作成したカードは、ボード上で自由に動かすことができます。

![move card](https://storage.googleapis.com/zenn-user-upload/3cungal7cbgjm15dbgaon6uxz3nm)

作成したカードをテーブル表示し、各カラムでソートすることも可能です。

![table view](https://storage.googleapis.com/zenn-user-upload/q6og7hrokvpwkdlt39fij2w8q1fg)

---

一通りのテンプレートを見てみた感じ、テンプレートごとに扱っているデータの構造はすべて同じで、最初に設定されている Group や Type、Property などの属性値が違うだけのようです。`Trello`的なボード管理ツールで、UI が Notion チックなもの、という感じを受けました。(最近の Trello や Asana を使ってないので的外れなことを言ってるかも)
Notion だとページの下にページを作成するなど入れ子構造なデータを作成できますが、そのようなことは今のところできないようです。

## Focalboard Server

Focalboard Desktop は、デスクトップ上でのみ動作し、データも起動したマシン内にのみ保存されるため、個人用に使われることが想定されていますが、Focalboard Server を使うことで チーム共通の Focalboard を立てる事ができるようです。Focalboard Server の画面はブラウザでしかみる事ができず、Focalboard Server の内容を Focalboard Desktop で閲覧・編集するというようなことは、まだできないようです。

### セットアップ

サーバー版の Focalboard のセットアップ手順は公式サイトに書かれています。

https://www.focalboard.com/download/personal-edition/ubuntu/

Ubuntu でのセットアップ方法しかなかったため、Focalboard Server 用の Dockerfile を書いて起動しました。

:::message
動作させるには、下記の Docker 関連ファイル以外に、上記公式ドキュメントに記載されている nginx 設定ファイルを`focalboard/focalboard`というファイル名で置いておく必要があります。
:::

```yaml:docker-compose.yml
version: "3"

services:
  focalboard:
    build: ./focalboard
    ports:
      - "8000:8000"
```

```docker:focalboard/Dockerfile
FROM ubuntu:18.04

RUN apt update && apt install -y nginx wget gzip

RUN wget https://releases.mattermost.com/focalboard/0.5.0/focalboard-server-linux-amd64.tar.gz \
  && gzip -dc focalboard-server-linux-amd64.tar.gz | tar -xvzf - \
  && mv focalboard /opt \
  && rm focalboard-server-linux-amd64.tar.gz

ADD focalboard /etc/nginx/sites-available/

RUN ln -s /etc/nginx/sites-available/focalboard /etc/nginx/sites-enabled/focalboard \
  && rm /etc/nginx/sites-enabled/default \
  && nginx -t \
  && /etc/init.d/nginx reload \
  && /etc/init.d/nginx restart

WORKDIR /opt/focalboard
ENTRYPOINT ["bin/focalboard-server"]
```

上記ファイルを用意した後、`docker-compose up -d`で起動するはずです。
公式ドキュメントにある HTTPS 設定や PostgreSQL の設定は行っていません(PostgreSQL の設定を行わないと SQLite が使われる)。

### アカウント作成

Focalboard Server 起動後、`http://localhost:8000` にアクセスするとログイン画面に移動します。

![login](https://storage.googleapis.com/zenn-user-upload/kqwq9hu2ohmnj1qhal3vl22al4mw =400x)

`or create an account if you don't have one`のリンクをクリックすると、アカウント作成画面に移動します。

![create account](https://storage.googleapis.com/zenn-user-upload/up8oyf48jkwnwrb77u77cte4c6it =400x)

アカウントを作成すると、デスクトップ版と同様のトップ画面が表示されます。

#### ユーザーの招待

二人目以降のユーザーは招待リンクから作成します。招待リンクは、左上のメインメニューの`Invite Users`から発行できます。

![Main menu](https://storage.googleapis.com/zenn-user-upload/yxuausoybb22put1xnllqf74tbbh =300x)

### ボードの作成

別々のアカウントでログインした 2 つのブラウザを並べて操作してみました。

![realtime sync](https://storage.googleapis.com/zenn-user-upload/hsufvgcehjl3l56mqs0g69q7jo30 =800x)

nginx の設定ファイルにあるように、Focalboard 上のイベントは WebSocket でやりとりされているため、一方のユーザーによる操作が他方のブラウザにリアルタイムで反映されているのがわかります。

その他の使い勝手については Focalboard Desktop と同様です。

## 設定

### ユーザー設定

ユーザー設定は Focalboard の左下の`Settings`メニューから操作できます。

![settings](https://storage.googleapis.com/zenn-user-upload/3nrdogvfhp61lfnna8f84hc5tmar =400x)

#### Import/Export archive

`Import archive` / `Export archive` メニューから、作成したボードやカードの情報を Import/Export できるようです。Export されるデータの形式は JSON です。
また、Trello、Notion、Asana の情報を Focalboard に Import するためのスクリプトも用意されているようです。今回は試していないため、どの程度の情報が Import できるかはわかりません。

https://github.com/mattermost/focalboard/tree/main/import

#### Set language

Focalboard 画面の表示言語を指定することができます。現在は`English`と`Spanish`のみ指定できます。時間があれば日本語化のコントリビューションをしたいと思います。

#### Set theme

画面のテーマを切り替えることができます。`Default`、`Dark`、`Light`が利用可能です。

![default theme](https://storage.googleapis.com/zenn-user-upload/ib222g4sa1ibhjfho1mpwef3mh9b =600x)
_Default Theme_

![dark theme](https://storage.googleapis.com/zenn-user-upload/3vcd3vrr0rgixyiqw6i00b7o06bg =600x)
_Dark Theme_

![light theme](https://storage.googleapis.com/zenn-user-upload/qdrm4a0daen8yntwdydm6serigll =600x)
_Light Theme_

### サーバー設定

Focalboard Server のサーバー設定は、公式ドキュメントに記述があります。

https://www.focalboard.com/guide/admin/#personal-server-configuration

## 感想

オンプレミス版のチャット基盤として根強い人気のある Mattermost 社が、[Trello](https://trello.com/ja)や GitHub Projects のような Kanban?ツールのオンプレミス版を開発し始めたというのは面白いなと感じました。[Trello](https://trello.com/ja)の OSS クローンとしては[Wekan](https://wekan.github.io/)や[Restyboard](https://restya.com/board)などがありますが、[Notion](https://notion.so)的な UI を加味してよりモダンになった Kanban ツールを目指しているような気がします。

Mattermost も近年では従業員が 100 人を超えるようになり、開発だけでなくマーケティングやコミュニティマネージメントなどの多くのチームができているようで、チームごとのタスクを俯瞰して管理しやすいようにこのようなツールの開発を初めたのではないかと個人的には思っています。Mattermost 自体の開発動機が、「SaaS 上のチャットツールでは自分たちのコミュニケーションに関するデータを自由に扱えない」というフラストレーションを解消するためというのを考えると、タスク管理に関するデータも SaaS だけに頼るのは辞めたいという意思が見えるように思います。

Focalboard は現時点ではまだベータ版の段階のため、モバイル版がまだなかったり、細かなところの UI などは改善が必要な点も多そうですが、現時点でも十分に使えるツールにはなっているため、今後どのように進化していくのか注目していきたいと思います。
