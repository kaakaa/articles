---
title: "Mattermost 9.11の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2024-07-19T22:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/08/16 に Mattermost のアップデートとなる `v9.11.0` がリリースされました。  

本バージョンは[ESR(Extended Support Release)](https://docs.mattermost.com/upgrade/extended-support-release.html)に設定されており、2025/5/15までセキュリティ対応/バグ修正等のサポートが行われる予定です。(ESRでないバージョンのサポート期間は3ヶ月間です)

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v9 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v9-changelog.html#release-v9-11-extended-support-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

https://www.youtube.com/watch?v=W9X5Bflf75c

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## キーボードショートカット一覧画面の更新

`Ctrl + /`もしくは`Cmd + /`で開くことのできるキーボードショートカット一覧画面の表示が更新されました。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.11/channels-keyboard-shortcut.png)

## @メンションを含むメッセージを編集する際の注意

@メンションを含むメッセージを編集する際、"このメッセージを@メンション付きで編集しても、相手には通知されません。"という旨のメッセージが表示されるようになりました。メッセージを編集する際に、メンションが飛んでしまうかどうかを心配する必要がなくなります。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.11/channels-edit-message.png)

## (Enterprise) システム管理者によるユーザー設定の変更

Mattermost Enterprise版を利用している場合、システム管理者が各ユーザーの設定を変更することができるようになったそうです。  
[User management configuration settings \- Mattermost documentation](https://docs.mattermost.com/configure/user-management-configuration-settings.html)

**システムコンソール > ユーザー管理 > ユーザー**から、設定を変更したいユーザーを選択し、**Manage User Settings**から変更できます。(Free版利用のため、その先の画面については分からず)

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.11/channels-manage-user-settings.png)


## アップグレード時の注意事項


### Elasticsearch v8のサポート

Elasticsearch v8のサポートが追加されました。また、Opensearchのサポートがベータ版として追加され、`ElasticsearchSettings.Backend`で Elasticsearch と Opensearch を選択できるようになりました。この変更により、AWS Elasticsearch v7.10はサポートされなくなるため、AWS Elasticsearchを利用しているユーザーはAWS Opensearchへのアップグレードが必要となります。

詳しくは、以下のページを参照ください。  
[Important Upgrade Notes \- Mattermost documentation](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

## その他のトピック

特になし

## おわりに
次のリリースは、メジャーバージョンアップとなる`v10.0`となります。リリースは 2024/9/16(Mon)を予定しています。  
`v10.0`で削除予定の機能については、以下のページを参照ください。  
[v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html)

