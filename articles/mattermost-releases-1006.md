---
title: "Mattermost 10.6の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-03-14T23:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2025/03/14 に Mattermost のアップデートとなる `v10.6.0` がリリースされました。  
本リリースでは、目立った機能の追加等はないようです。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-6-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## Free版環境のユーザー数制限に関するアップデート

> Unlicensed server limits have been updated: a soft limit of 2500 users now results in a banner notification visible by admins, and a 5K hard limit that prevents the creation or activation of users until the user count is reduced below the hard limit.

Mattermostの[Free版](https://mattermost.com/pricing/)を利用している環境のユーザー数制限の値が更新され、2,500人以上のアクティブユーザーが存在する場合は管理者向けのバナー警告が表示されるようになり、5,000人以上のアクティブユーザーがいる場合はユーザー作成やユーザーのアクティベーションを行うことができなくなります。

この制限は、今年中に最大1,000人までに制限されることが明かされています。詳しくは、以下のフォーラムの投稿を参照ください。  
[Feedback on Collaboration for Mission\-Critical Work \- Feedback \- Mattermost Discussion Forums](https://forum.mattermost.com/t/feedback-on-collaboration-for-mission-critical-work/17563)

## アップグレード時の注意事項

### PostgreSQLの最小サポートバージョンがv13に

前回の記事でも触れましたが、v10.6からPostgreSQLのサポートバージョンが v13 以降になりました。
[PostgreSQLのサポートポリシーについて - Mattermost 10\.5の新機能](https://blog.kaakaa.dev/post/mattermost-releases-1005/#postgresql%E3%81%AE%E3%82%B5%E3%83%9D%E3%83%BC%E3%83%88%E3%83%9D%E3%83%AA%E3%82%B7%E3%83%BC%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)

PostgreSQL v11.x, v12.x を利用している環境をMattermost v10.6にアップデートする場合、先にPostgreSQLのバージョンをv13以降にする必要があるかと思います。

## おわりに

次の`v10.7`のリリースは 2025/04/16(Wed)を予定しています。  
