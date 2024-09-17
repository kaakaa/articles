---
title: "Mattermost 10.0の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2024-09-17T22:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/09/16 に Mattermost のアップデートとなる `v10.0.0` がリリースされました。  

本バージョンはメジャーバージョンアップとなりますが、大きな機能追加や変更などがあるわけでは無いように見えています。  
どちらかというと非推奨となる機能の追加や、後方互換を満たせなくなるようなプラグインのメジャーアップデート等の導入が目的のようです。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-0-major-release)

本バージョンでの主な変更点を紹介する動画は、まだMattermostの公式YouTubeチャンネルで公開されていないようです。  
メジャーバージョンアップにより動画公開方針が変更になったりしていなければ数日中にアップロードされるかと思うので、以下のMattermost公式YouTubeチャンネルをチェックしてみてください。

https://www.youtube.com/@MattermostHQ/videos

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## 非推奨/利用条件変更となった機能

以前からアナウンスがありましたが、Mattermost v10のリリースに合わせて非推奨、または利用条件が変更となった機能があります。

* MySQLの非推奨化
* Apps Frameworkの非推奨化
* Playbookの最新版を利用するにはEnterpriseライセンスが必要に

詳しくは、過去の記事を参照してください。  
[Mattermost 9\.10の新機能](https://blog.kaakaa.dev/post/mattermost-releases-910/#mattermost-v10)

## 用語の変更

Mattermost v10のリリースに合わせていくつかの機能の呼称が変更されました。

* [**Threaded Discussions**](https://docs.mattermost.com/collaborate/organize-conversations.html) ← `Collapsed Reply Threads`
* [**Advanced Access Control**](https://docs.mattermost.com/manage/team-channel-members.html#advanced-access-controls) ← `Channel Moderation`
* [**System-wide notifications**](https://docs.mattermost.com/manage/system-wide-notifications.html) ← `Announcement banner`
* [**Delegated Granular Administration**](https://docs.mattermost.com/onboard/delegated-granular-administration.html) ← `System Roles`
* [**EntraID**]() ← `Office 365`
  * ref: [Azure AD が Microsoft Entra ID に名称変更 \- News Center Japan](https://news.microsoft.com/ja-jp/2023/07/12/230712-azure-ad-is-becoming-microsoft-entra-id/)

## パッケージ済みプラグインのメジャーアップデート

各種パッケージ済みプラグインのメジャーアップデート版が同梱されています。

* [Mattermost Copilot plugin v1.0.0](https://github.com/mattermost/mattermost-plugin-ai/releases)
* [Mattermost Calls plugin v1.0.1](https://github.com/mattermost/mattermost-plugin-calls/releases)
* [Mattermost Playbooks plugin v2.0.1](https://github.com/mattermost/mattermost-plugin-playbooks/releases)

Playbook pluginはv2以降、Enterpriseライセンスユーザーのみが利用できます。その他のユーザーも、Playbook plugin v1を引き続き利用できますが、機能アップデートやセキュリティ修正等は今後行われないようです。

## GIF Picker、カスタムグループ、メッセージ優先度機能のBetaラベル削除

GIF Picker、カスタムグループ、メッセージ優先度のBetaラベルが外れ、正式な機能となりました。

**GIF Picker**

絵文字選択画面の隣のタブで、GIPHYのGif画像を検索することができる機能です。  
**システムコンソール > 統合機能 > GIF > GIF選択機能を有効にする** から、機能のON/OFFを設定することができます。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.0/channels-gif-picker.png)

**カスタムグループ**

複数のユーザーに一度にメンションできる、任意のグループを作成できる機能です。Professional/Enterprise版の利用者のみ利用できる機能です。  
[Manage custom groups \- Mattermost documentation](https://docs.mattermost.com/collaborate/organize-using-custom-user-groups.html)

**メッセージ優先度**

投稿するメッセージを目立たせるためのメッセージ優先度ラベルを付与できる機能です。  
![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.0/channels-message-priority.png)

## アップグレード時の注意事項

いくつかは上記で触れていますが、詳しくは以下を参照ください。  
https://docs.mattermost.com/about/mattermost-v10-changelog.html#important-upgrade-notes

## その他のトピック

特になし

## おわりに

次の`v10.1`のリリースは 2024/10/16(Wed)を予定しています。  


