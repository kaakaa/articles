---
title: "Mattermost 10.1の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2024-10-18T23:50:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/10/16 に Mattermost のアップデートとなる `v10.1.0` がリリースされました。  
また、同日にshared server関連のIssueを修正した`v10.1.1`がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-1-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## (Professional/Enterprise) チャンネルブックマーク機能の有効化

> Enabled Channel Bookmarks, added re-ordering, and fixed URL validity checking.

Mattermost v9.8でサーバー側のみ機能が追加されていた[チャンネルブックマーク機能](https://blog.kaakaa.dev/post/mattermost-releases-908/#%E3%83%81%E3%83%A3%E3%83%B3%E3%83%8D%E3%83%AB%E3%83%96%E3%83%83%E3%82%AF%E3%83%9E%E3%83%BC%E3%82%AF)がリリースされました。  
Professional版とEnterprise版のユーザーのみ利用可能です。

[Manage channel bookmarks \- Mattermost documentation](https://docs.mattermost.com/collaborate/manage-channel-bookmarks.html)

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.1/channels-bookmark.gif)

## チャンネルごとの通知音設定

> Added channel specific message notification sounds configuration.

チャンネルに新規投稿があった場合の通知音について、アカウントに対する設定とは別にチャンネル毎に異なる通知音を設定することができるようになりました。  
重要なチャンネルのみ通知音を変えて注意を向けやすくするなどの対応が取れるようになります。

[Manage your channel\-specific notifications \- Mattermost documentation](https://docs.mattermost.com/preferences/manage-your-channel-specific-notifications.html)

**チャンネルメニュー > 通知の設定 > サウンド**から設定できます。


![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.1/channels-specified-notification-sound.gif)

## その他の変更

### Pluginファイルアップロード時のメッセージ改善

> Added a more descriptive error message, “Uploaded plugin size exceeds limit.” for plugin uploads that are too large.

Mattermost Pluginをアップロードする際、Mattermostのアップロードファイル制限を超えるサイズのファイルをアップロードしようとした場合に表示されるエラーメッセージから原因を特定しづらいという問題がありました。

> Received invalid response from the server

本バージョンで以下のようにメッセージが改善され、メッセージから原因が特定しやすくなりました。[PR](https://github.com/mattermost/mattermost/pull/26271)

> Uploaded plugin size exceeds limit. This limit can be changed in the System Console via File Storage > Maximum File Size

### Mattermost Pluginsの設定でSecret値が扱えるように

> Plugins are now allowed to mark setting fields as secret, obfuscating them in the System Console and the Support Packet.

Mattermost Pluginで設定を扱う際、WebUI上で設定値を見えなくするには Webapp Pluginの[registerAdminConsoleCustomSetting](https://developers.mattermost.com/integrate/reference/webapp/webapp-reference/#registerAdminConsoleCustomSetting)を利用した独自コンポーネントの開発などを実施する必要がありました。  
本バージョンから、Plugin Manifestの`settings_schema`に`secret: bool`というフィールドが追加され、`true`を指定するとその設定値をWebUI上から見えなくすることができるようになりました。

[Manifest reference at Mattermost](https://developers.mattermost.com/integrate/plugins/manifest-reference/)

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.1/plugins-secret-setting.png)

設定入力中は入力値が見えますが、設定を保存して、画面を切り替えるとダミー文字列が表示されるようになるようです。

## アップグレード時の注意事項

特になし

## その他のトピック

### Hacktoberfest 2024

今年もMattermostではHacktober向けのキャンペーンを行っており、以下のような特設ページが用意されています。  
[Mattermost \- Hacktoberfest 2024](https://mattermost.com/wp-content/uploads/2024/09/Hacktoberfest_Hero_Computer.png). 

また、Hacktoberfestの開幕を記念して、[Holopin](https://www.holopin.io/)のバッジを配布しているようです。  
[Hacktoberfest 2024: Contribute, connect & collect digital rewards\! \- Mattermost](https://mattermost.com/blog/hacktoberfest-2024/)

### Reference Architecture検討

[v10 to the power 5: Scaling Mattermost to support 100,000 users](https://mattermost.com/blog/scaling-mattermost-for-100000-users/)

上記エントリでは、Mattermostを10万ユーザー以上で利用する場合のReference Architectureを検討する際に実施したことについて説明されています。  
Mattermostの使われ方について調査し、[lord test](https://github.com/mattermost/mattermost-load-test-ng)ツールを改良し、Reference Architectureを導出していく流れが描かれており、個人的に興味深い記事でした。

ユーザー数ごとのReference Architectureについては以下で公開されています。  
[Scaling for Enterprise \- Mattermost documentation](https://docs.mattermost.com/scale/scaling-for-enterprise.html)


## おわりに

次の`v10.2`のリリースは 2024/11/15(Fri)を予定しています。  


