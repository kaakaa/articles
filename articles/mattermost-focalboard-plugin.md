---
title: "FocalboardのMattermostプラグインについて"
emoji: "🔌"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: [focalboard, oss, mattermost, plugin]
date: 2021-06-03T00:00:00+09:00
published: true
---

# Mattermost Focalboard Plugin

[Mattermost](https://mattermost.com/)社が開発しているNotion/TrelloライクなKanbanツールである[Focalboard](https://www.focalboard.com/)プロジェクトに、Mattermostと連携するためのプラグインが追加されたようなので動かしてみました。  

Focalboardについては以下にまとめています。
[NotionのようなUIのTrelloっぽいKanbanツールのOSSの Focalboard を触ってみた](https://zenn.dev/kaakaa/articles/mattermost-focalboard-first)

## Set up

### Mattermost

:::message alert
現在公開されているMattermost Focalboard Plugin v0.6.7 は、2021/06/16リリース予定のMattermost v5.36以降のバージョンでないと動作しません。Mattermost v5.36リリース前に試すには、開発者用のMattermost環境を自分で構築する必要があります。[Developer Setup](https://developers.mattermost.com/contribute/server/developer-setup/)  
:::

:::message alert
以下の記述は、Mattermost v5.36が公開された後の手順を書いているものであり、記事公開時点では動作しないことに注意してください。
:::

公式のDocker Composeを使いMattermostも起動しておきます。  
[mattermost/mattermost\-docker: Dockerfile for mattermost in production](https://github.com/mattermost/mattermost-docker)

上記リポジトリをCloneした後、Mattermostのバージョンを `app/Dockerfile` から指定します。  
(**記事執筆時点では`v5.36.0`は公開されていないため、以下の設定ではDocker Compose起動時にエラーとなるはずです。`v5.36.0`が公開される2021/06/16以降に実行するか、[Developer Setup](https://developers.mattermost.com/contribute/server/developer-setup/)を参考に開発者環境でMattermostを起動する必要があります。**)

```diff:app/Dockerfile
...
-ARG MM_VERSION=5.31.0
+ARG MM_VERSION=5.36.0
...
```


また、FocalboardプラグインをMattermostのシステムコンソールからアップロードするための設定として、Docker Compose内の`app`サービスの環境変数に`MM_PLUGINSETTINGS_ENABLEUPLOADS=true`を追加しておきます。これを設定しておかないと、画面上からプラグインをアップロードできません。また、Mattermostの起動ポートを`8065`に書き換えています。

```diff:docker-compose.yml
version: "3"

services:
  ...
  app:
    ...
    environment:
      ...
      # in case your config is not in default location
      #- MM_CONFIG=/mattermost/config/config.json
+      - MM_PLUGINSETTINGS_ENABLEUPLOADS=true
  web:
    build: web
    ports:
-      - "80:8080"
-      - "443:8443"
+      - "8065:8080"
+      - "8443:8443"
    ...
```

上記を書き換えたあと、Docker Composeを起動します。

```bash
$ docker-compose up -d --build
```

起動後、ブラウザから`http://localhost:8065`にアクセスし、Adminユーザー・チームをそれぞれ作成しておきます。

## Mattermost Focalboard Pluginの起動

### Pluginの入手

以下のページからMattermost Focalboard Plugin (`mattermost-plugin-focalboard.tar.gz`) をダウンロードします。  
[Release Plugin v0\.6\.7 · mattermost/focalboard](https://github.com/mattermost/focalboard/releases/tag/v0.6.7-plugin)

### Pluginのアップロード

Mattermostの **System Console > PLUGINS (BETA) > Plugin Management > Upload Plugin** から、先ほどダウンロードしたPluginファイルをアップロードします。

![Mattermost Main Menu](https://github.com/kaakaa/zenn-articles/raw/master/articles/img/mattermost-focalboard-plugin/mattermost-main-menu.png)

![Mattermost Upload Plugin](https://github.com/kaakaa/zenn-articles/raw/master/articles/img/mattermost-focalboard-plugin/mattermost-upload-plugin.png)

Pluginのアップロードが成功すると、**System Console > PLUGINS (BETA) > Plugin Management > Installed Plugin** に `Focalboard` プラグインが表示されます。  
**Enable** リンクをクリックし、Pluginを起動します。

![Mattermost Enable Plugin](https://github.com/kaakaa/zenn-articles/raw/master/articles/img/mattermost-focalboard-plugin/mattermost-enable-plugin.png)

### Pluginの利用
Focalboard Plugin が有効化されると、Mattermostの各チャンネルのヘッダー部分にFocalboardのアイコンが表示されるようになります。

![Mattermost Focalboard Icon](https://github.com/kaakaa/zenn-articles/raw/master/articles/img/mattermost-focalboard-plugin/mattermost-focalboard-icon.png)

このアイコンをクリックすると、Focalboardが別タブとして開きFocalboardが利用できるようになります。  

![Focalboard](https://github.com/kaakaa/zenn-articles/raw/master/articles/img/mattermost-focalboard-plugin/focalboard.png)

その後の操作はFocalboardと一緒です。

## Mattermost Focalboard Pluginの特徴

### チャンネルごとにBoardが存在する

Mattermost Focalboard Pluginでは、**チャンネルごとに別々のBoardが作成される**ようです。  
それぞれのチャンネルごとにBoardを管理できるのは便利な一方、Focalboardが運用されているチャンネルと運用されていないチャンネルがBoardにアクセスしてみないと分からないと言うのが課題になりそうです。

### ユーザー情報をMattermostと共有している

Mattermost Pluginとして起動したFocalboardでは、**ユーザー情報をMattermostと共有しているため、Focalboardでのアカウント作成やログインを行う必要がありません**。
逆に言うと、この方で立てたFocalboardにアクセスするには、Mattermostのアカウントを作成する必要があります。

## まとめ

Focalboard用のサーバーを立てる必要なく、Mattermostインスタンス内でFocalboardを動作させることができるMattermost Focalboard Pluginを試してみました。
Focalboard専用のサーバーを立てるのと比べ、ユーザーアカウントの作成が不要なため利用しやすくなりそうです。また、Mattermostのチャンネルからアクセスしやすくなるのも便利そうです。

Focalboard上のイベントをMattermostに通知する機能などはまだないようですが、今後はそう言う機能も追加されていくのかなと思います。

