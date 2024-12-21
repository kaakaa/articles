---
title: "Mattermost 10.3の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2024-12-21T20:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/12/16 に Mattermost のアップデートとなる `v10.3.0` がリリースされました。  
同日に幾つかの修正が加わった`v10.3.1`がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-3-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

https://www.youtube.com/watch?v=iSLZNrz597M

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## (Professional/Enterprise) メッセージの予約投稿機能

> Added a feature to schedule a message at a future date (Professional and Enterprise plans).

Professional/Enterprise版限定ですが、メッセージの予約投稿機能が追加されました。

メッセージを入力し、投稿ボタンの横にあるメニューをクリックすると、メッセージスケジューリング画面が表示されます。 **任意の時刻を選択する**メニューを選択すると、投稿時刻を30分単位で選択できます。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.3/channels-schedule-message.png)

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.3/channels-schedule-message-custom.png)

スケジュールされたメッセージは**下書き > スケジュール済み**から確認でき、メッセージの編集や投稿時刻の再スケジューリングなどを行うこともできます。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.3/channels-schedule-message-list.png)

また、公式ブログのこちらの記事でも予約投稿機能が紹介されています。  
[Send messages at the perfect time with Scheduled Messages in Mattermost](https://mattermost.com/blog/introducing-scheduled-messages/)


## 通知テスト機能

> Added an option to test notifications.

通知の設定を確認するために`system-bot`からのDMをトリガーできる機能が追加されました。

**設定 > 通知 > Send a test notification**ボタンをクリックすると、`system-bot`からテスト用のDMが送信され、各種デバイスに通知を飛ばすことができます。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.3/channels-test-notification.png)

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.3/channels-test-notification-message.png)

## 新しい検索インターフェース

> Added a new search interface.

画面上部にある検索バーをクリックしたときのインターフェースが新しくなりました。

メッセージ検索とファイル検索を切り替えることができ、各種検索クエリをボタンから指定しやすくなりました。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.3/channels-search-interface.png)

* `From:` 指定したユーザーからの投稿のみを検索
* `In:` 指定したチャンネルへの投稿のみを検索
* `On:` 投稿された日付を指定
* `Before:` 指定した日付以前の投稿を検索
* `After:`　指定した日付以降の投稿を検索


## Classic Mobile Appが廃止に

> The Classic Mobile App has been phased out. Please download the new v2 Mobile App from the Apple App Store or Google Play Store. See more details in the classic mobile app deprecation Mattermost forum post.

Mattermost Mobileアプリv2が2023/01/16にリリースされた際、旧バージョンであるMobileアプリv1も後方互換のためにMattemost Classic Mobile Appとして引き続き公開されていましたが、2024/12/16をもって公開が停止されたようです。

[Classic Mobile App Deprecation \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/classic-mobile-app-deprecation/18703)


## その他のトピック

特になし

## おわりに

次の`v10.4`のリリースは 2025/01/16(Mon)を予定しています。  


