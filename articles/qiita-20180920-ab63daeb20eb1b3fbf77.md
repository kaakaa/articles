---
title: "Mattermost5.3の新機能"
date: 2018-09-20T00:00:00+09:00
emoji: "📣"
type: "tech"
topics: ["Mattermost"]
published: true
---

# はじめに

2018/09/14 に Mattermost 5.3.0 がリリースされたので、アップデートの内容について簡単に紹介します。

詳細については公式のリリースを確認ください。

- [Mattermost Changelog — Mattermost 5\.3 documentation](https://docs.mattermost.com/administration/changelog.html#release-v5-3)
- [Mattermost 5\.3: Enhanced search on desktop and mobile, plugin hackathon highlights and more \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-3-enhanced-search-on-desktop-and-mobile-plugin-hackathon-highlights-and-more/)

---

v5.3.0 のリリース作業でミスがあり翻訳が一部壊れています。英語以外の翻訳版を利用されている場合は v5.3.1 を利用してください。
[Mattermost 5\.3\.1 released \- Mattermost Private Cloud Messaging](https://mattermost.com/blog/mattermost-5-3-1-released/)

# アップデート内容

## 日付指定検索

Mattermost メッセージ検索機能で投稿日付を指定できるようになりました。

- `before:YYYY-MM-DD`: 指定日付より前の投稿を検索
- `on:YYYY-MM-DD`: 指定日付に行われた投稿を検索
- `after:YYYY-MM-DD`: 指定日付より後の投稿を検索

また、右上の検索欄に`after:`、`on:`、`before:`を入力することで、日付選択モーダルが現れます。

![](https://qiita-image-store.s3.amazonaws.com/0/9891/1eb95c86-7eea-7103-8a61-99e7f447a3e1.png)

## SAML ユーザーへの ID 設定

ID Provider として SAML を設定している場合、過去に無効化されたユーザーと同じメールアドレスを持つユーザーを新たに追加すると、無効化されたユーザーの履歴にアクセスできてしまうという問題がありました。
この問題を解決するために、メールアドレスの代わりとなる ID を指定することができるようになりました。
ID を指定することができるようになったことにより、結婚による姓の変更などに起因するメールアドレスの変更にも対応できるようになりました。

## プラグインハッカソンの成果

2018/8/16,17 に行われた Mattermost プラグインハッカソンで開発された~~プラグイン~~機能の紹介です。

- [**RSS feed**](https://github.com/jespino/matterfeed): ~~RSS のフィードを行えるプラグインのようです~~ RSS のフィードを購読する連携機能のようです。この機能はプラグインではないですが、プラグインハッカソン中に開発された成果として紹介されていたようです
- [**Webcome bot**](https://github.com/mattermost/mattermost-plugin-welcomebot/tree/master): 新しくチームに加入したユーザーにウェルカムメッセージを投稿できます
- [**Matterpoll**](https://github.com/matterpoll/matterpoll): Mattermost 上で投票を行える機能です
- [**/remind**](https://github.com/scottleedavis/mattermost-plugin-remind): 指定した時間が経過した後にリマインドをしてくれるプラグインです。
- **Crosspost**: つのチャンネルを複数の Mattermost サーバーとリンクすることができ流ようです
- [**Auto-translator**](https://github.com/mattermost/mattermost-plugin-autotranslate): Amazon Translate を使って様々な言語を英語に自動変換するプラグインです

その他にもリッチなエディタ機能を追加するプラグインや、[表示時間をユーザーのローカルタイムゾーンに合わせてくれるプラグイン](https://github.com/mattermost/mattermost-plugin-walltime)などの開発が行われています。
また、Mattermost CTO によるハッカソンの[振り返り](https://mattermost.com/blog/plugin-hackathon/)やプラグインの[ビデオチュートリアル](https://www.youtube.com/watch?v=Cx2EBkGkz00)などが公開されています。

## その他

### リックソフトが Mattermost とパートナー契約を締結

[ニュースリリース：20180920 リックソフト Slack ライクなオンプレミス型チャットツール Mattermost のパートナー契約を締結](https://www.ricksoft.jp/news/n20180920.html)

アトラシアン社のパートナー企業として有名なリックソフト社が HipChat Server 販売終了を受け、Mattermost とのパートナー契約を締結したそうです。
私の知る限りでは、国内初のパートナー企業となるかと思います。（他にあったらゴメンナサイ）
[Mattermost Partner Programs — Mattermost 5\.3 documentation](https://docs.mattermost.com/process/partner-programs.html)

### Matterpoll

現在、Mattermost プラグイン機能を活用した投票機能である`Matterpoll`という OSS を開発しています。
まだ安定バージョンではないため今後も破壊的変更が入る可能性は高いですが、もしご興味がある方がいらっしゃいましたら導入・Contribute をお願いします。
[matterpoll/matterpoll: Poll command for Mattermost](https://github.com/matterpoll/matterpoll)

![](https://qiita-image-store.s3.amazonaws.com/0/9891/dd7dd94a-78cf-c6cb-6b2d-4a1bce1e13aa.png)

# おわりに

次回の`v5.4`のリリースは 2018/10/16(Tue)を予定しています。

---

[Mattermost 日本語\(@mattermost_jp\)さん \| Twitter](https://twitter.com/mattermost_jp?lang=ja) で Mattermost に関する日本語の情報を提供しています。
