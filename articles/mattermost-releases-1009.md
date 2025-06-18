---
title: "Mattermost 10.9の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-06-17T09:50:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2025/06/16 に Mattermost のアップデートとなる `v10.9.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-9-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## Enterprise Advancedライセンスの追加

v10.9から**Enterprise Advanced**ライセンスが追加されたようです。

* [Mattermost Plans \- Mattermost documentation](https://docs.mattermost.com/about/plans.html)

上記のページを見るに、Enterpriseライセンスとほぼ同等の機能のように見えますが、v10.9の新機能であるチャンネルバナーや、ユーザー属性(user attributes)機能などは**Enterprise Advanced**ライセンス限定の機能だと書かれています。

> * Introduced a configurable channel banner feature for channel admins, visible across desktop, web, and mobile platforms. This feature requires an Enterprise Advanced license.

> * Added support for user attributes for Enterprise Advanced licensed servers. Defined policies that automatically grant channel memberships based on user attributes. Membership updates happen automatically when user attributes change — no need for manual role adjustments.

(2025/06/18 追記)

Enterprise Advancedプランに関する公式の紹介動画が公開されていました。  
Enterpriseプランに加えポスト量子暗号やEnd-to-End暗号化など、National Securityレベルのセキュリティやミッションクリティカル性を求められる環境向けのプランとして導入されたものという印象を受けました。

https://youtu.be/-pTC_3oyc9Q?si=Mp4ohlVDfAn-sTUK

## チャンネル設定メニューの統合

> Consolidated all channel editing functionality into a single, accessible modal located in the channel header menu. Users can now update channel names, URL slugs, convert to private, modify/add a purpose and header (with a live markdown preview), manage channel banners, and archive the channel—all in one place. Updates include safeguards for unsaved edits, improved URL-slug editing, and enhanced keyboard and navigation accessibility.

チャンネル設定に関するメニューが一つのダイアログに統一されました。

今までのバージョンでは、チャンネル名の変更やチャンネルヘッダーの編集などは別々のメニューとして存在していましたが、それらが **Channel Settings (チャンネル設定)**に統一され、一つの画面でチャンネルに関する設定の編集を行うことができるようになりました。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.9/channels-channel-settings.png)


## "<..>" で囲まれたメールアドレスが使用禁止に

> Updated the email validation logic to ensure Mattermost no longer accepts email addresses enclosed in angle brackets (e.g., <billy@example.com>). Although these comply with RFC 5322 and RFC 6532 standards, they introduce unnecessary complexity. If a user already has such an email in your installation, they may face issues like being unable to update their profile. To resolve this, the email must be modified manually using the command: mmctl user email "<affecteduser@domain.com>" affecteduser@domain.com.

電子メールアドレスの検証ロジックが更新され、"<..>"で囲まれたメールアドレス (例: `<billy@example.com>`) が使用できなくなりました。

すでに "<..>" で囲まれた電子メールアドレスを使用している場合、プロフィール更新等が実施できなくなるようです。解消するには `mmctl` コマンドを使って手動で登録電子メールアドレスの変更が必要です。

## サーバーログをプレーンテキスト形式で表示

**システムコンソール > サーバーログ**画面でJSONL形式のログを確認する際、プレーンテキスト形式で表示できるオプションが追加されました。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.9/admin-serverlog-plain.png)


## アップグレード時の注意事項

いくつかデータベーススキーマの変更があるようですが、アップデート時のダウンタイム等は発生しない見込みだそうです。  
詳しくは以下ページの`v10.9`の項を参照ください。  

[Important Upgrade Notes \- Mattermost documentation](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

## その他のトピック

特になし

## おわりに

次の`v10.10`のリリースは 2025/07/16(Wed)を予定しています。  
