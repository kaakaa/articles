---
title: "Mattermost 10.8の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-05-17T13:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2025/05/16 に Mattermost のアップデートとなる `v10.8.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-7-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。


## チャンネルメニューの改善

> Added an improved channel menu.

チャンネルメニューの表示内容が変更されました。  
いくつかのメニューがサブメニューに移動するなど、スッキリした見た目になっています。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.8/channels-channel-menu-improvement.png)

## 電子メール通知設定の改善

> Updated email notification settings to provide clearer wording and descriptions for both batched and non-batched scenarios. The settings dialog now reflects the selected status more accurately in both collapsed and expanded views, enhancing consistency and usability.

**設定 > 通知 > 電子メール通知**で表示される、電子メール通知設定の内容が変更されました。  
今までは5分以上離席状態が続いている間にイベントがあった場合に都度メールが送信されていましたが、15分/1時間ごとにまとめて(batch形式で)通知メールを送信するオプションが選択できるようになりました。


![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.8/channels-email-notification-setting.png)


## アップグレード時の注意事項

データベースに新たなテーブルが追加されるのと、レガシーなライセンスに対するサポートが終了となるようです。

詳しくは以下ページの`v10.8`の項を参照ください。  

[Important Upgrade Notes \- Mattermost documentation](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

## その他のトピック

特になし

## おわりに

次の`v10.9`のリリースは 2025/06/16(Mon)を予定しています。  
