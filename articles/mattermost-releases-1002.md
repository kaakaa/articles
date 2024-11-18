---
title: "Mattermost 10.2の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2024-11-18T23:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/11/15 に Mattermost のアップデートとなる `v10.2.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-2-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## チャンネルヘッダー部のピン留め投稿アイコンの非表示

> Updated the channel header to hide pinned posts when there aren’t any in the channel.

表示しているチャンネルにピン留めされたメッセージが存在しない場合、チャンネルヘッダー部に表示されるピン留めアイコンが表示されなくなりました。  
メッセージをピン留めするとアイコンが表示されるようになります。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.2/channels-header-pinned-icon.png)

(左: 新版 | 右: 旧版)

## Message Attachmentsでのフィールドごとのメンションサポート状況の変更

> Added full support for @mentions in the values of fields in message attachments.

[Message Attahcments](https://developers.mattermost.com/integrate/reference/message-attachments/)の`Title`と`Field Title`で@メンションがメンションとして表示されなくなりました。

参考: [MM\-59854 Fully allow at mentions in message attachment field values and add E2E tests by hmhealey · Pull Request \#28018 · mattermost/mattermost](https://github.com/mattermost/mattermost/pull/28018)

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.2/channels-mention-in-attachments.png)

(左: 新版 | 右: 旧版)


## REST APIによる投稿の完全な削除

> * Added a new URL parameter called permanent to DELETE /api/v4/posts/<post-id>, and set permanent to true in order to permanently delete a post and its attachments.
> * Added a new mmctl command, mmctl post delete <post-id>, in order to permanently delete a post and its attachments.

REST API `DELETE /api/v4/posts/<post-id>` のエンドポイントに `?permanent=true` というクエリを指定できるようになり、このクエリを指定して投稿の削除を実行すると投稿の情報がデータベースからも完全に削除され、投稿に添付されたファイルも削除されるようになりました。また、`mmctl`にも同等の機能として`mmctl post delete <post-id>`コマンドが追加されました。

この機能を利用するには、`config.json`の `EnableAPIPostDeletion`設定を`true`に設定する必要があります。この設定が`true`になっていないと、501エラーでAPIリクエストが失敗します。

```
$ curl -X DELETE -H "Authorization: Bearer abcdefghijklmnopqrstuvwxyz" "http://mattermost.example.com/api/v4/posts/12345abcdefghijklmnopqrstu?permanent=true"

{"id":"api.post.delete_post.not_enabled.app_error","message":"Cannot delete post, ServiceSettings.EnableAPIPostDeletion is not enabled.","detailed_error":"","request_id":"5cuuhbk64inx8rounjhy91f6ne","status_code":501}
```


## アップグレード時の注意事項

Dockerイメージの署名方法がDocker Content Trust (DCT) から[Cosign](https://github.com/sigstore/cosign)に変更されたため、イメージの検証方法も変更されています。詳しくは、以下のMattermost Forumの投稿を参照してください。

[Upcoming DCT deprecation \- Community \- Mattermost Discussion Forums](https://forum.mattermost.com/t/upcoming-dct-deprecation/19275)


## その他のトピック

### Hacktoberfest 2024 Recap

今年のHacktoberfestでは、38名のコントリビューターから174のPull Requestが提出されたそうです。  
[Hacktoberfest 2024: Community contributions from around the world](https://mattermost.com/blog/hacktoberfest-2024-recap/)

### Mattermost on Azure Marketplace

Azure MarketplaceからMattermostが利用可能になりました。Azure AKS上に展開されるサービスとしてMattermostをデプロイできるようです。

[Microsoft Azure Marketplace](https://azuremarketplace.microsoft.com/en-us/marketplace/apps/mattermost.mattermost-operator?tab=Overview)  
[Mattermost is now available on the Azure Marketplace \- Mattermost](https://mattermost.com/blog/mattermost-on-azure/)

### 新しいBug Bountyプログラムについて

11月からBug Bounty Programが[HackerOne](https://hackerone.com/mattermost?type=team)から[BugCrowd](https://bugcrowd.com/engagements/mattermost-mbb-public)に移行されるそうです。  
[Unveiling the future of our bug bounty program \- Mattermost](https://mattermost.com/blog/unveiling-the-future-of-our-bug-bounty-program/)

上記のエントリを見ると、新しいBug Bounty Programの対象として追加されるプロダクトの中に [`Mattermost Boards plugin`](https://github.com/mattermost/mattermost-plugin-boards)が存在しています。
Boards Pluginについては、[2023年9月にリリースされたMattermost v9.0で公式チームからのサポートがなくなり](https://forum.mattermost.com/t/upcoming-product-changes-to-boards-and-various-plugins/16669)コミュニティ開発に移行したはずですが、GitHubリポジトリを見ると細々とですがまた開発が再開している雰囲気もあります。また開発が再開されるのかもしれません。(スタンドアロン版のFocalboardについては、メンテナの募集をしているようです。[Issue](https://github.com/mattermost/mattermost-plugin-boards))

## おわりに

次の`v10.3`のリリースは 2024/12/16(Mon)を予定しています。  


