---
title: "Mattermost 10.10の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-07-17T10:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2025/07/16 に Mattermost のアップデートとなる `v10.10.0` がリリースされました。  
同日にセキュリティ修正が加わった`v10.10.1`がリリースされています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-10-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)
- [Enterprise Advanced](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## 引用文とスレッドの表示改善

> Improved the visual styling of blockquotes and comment threads for better readability and a modern appearance.

引用文とコメントスレッドの表示スタイルが改善され、より読みやすく現代的な見た目になりました。（コメントスレッドの表示改善の方はよくわかりませんでした）

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.10/channels-improve-blockquotes.png)

## エモーティコンの絵文字変換設定の追加

> Added a display setting for users to toggle rendering of emoticons (e.g., :D) as emojis (e.g., 😄).

ユーザー設定に、テキストのエモーティコン（例: `:D`）を絵文字（例: 😄）に自動変換するかどうかを切り替える設定が追加されました。  
**設定 > 表示 > Render emoticons as emojis** から設定することができます。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.10/channels-disable-emoticon.png)





## チャンネル自動分類の実験的機能追加

> Added ExperimentalChannelCategorySorting configuration setting to add the ability to automatically sort channels into categories upon creation/renaming.

`ExperimentalChannelCategorySorting` という設定が追加され、チャンネルの作成や名前変更時に自動的にカテゴリへの分類が行えるようになりました。  
**システムコンソール > 実験的機能 > 機能　> Channel Category Sorting**から設定できます。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.10/channels-category-sort-setting.png)

上記設定を有効にした後、チャンネル作成時のチャンネル名を`category-sort/test-3`のように、`/`の前にカテゴリ名を指定してからチャンネルを作成すると、自動でそのカテゴリに配置されます。以下の操作では、チャンネルを作成すると同時に既存のカテゴリ`category-sort`に配置される様子が確認できます。（存在しないカテゴリ名を指定すると、自動でそのカテゴリが作成されるようです。）

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.10/channels-category-sort.gif)

ただ、現在のバージョンで試したところ、チャンネル作成時もチャンネル名変更時も、指定したカテゴリだけでなく元のカテゴリにもチャンネルが重複して表示されてしまうバグがあるようで、まだ実験的な機能の位置づけのようです。

別のユーザーがそのチャンネルに参加した場合も、作成時に指定したカテゴリに属したチャンネルとして表示されるようなので、チャンネル作成時にどのようにグルーピングされるべきチャンネルかを指定できる機能となるようです。

## その他の機能

### [Enterprise] データ保持設定におけるピン留めされた投稿の保護機能

> Added `PreservePinnedPosts` configuration setting. If it's set to true, pinned posts will not be deleted by data retention.

データ保持設定に `PreservePinnedPosts` という新しい設定が追加されました。この設定を `true` にすると、データ保持ポリシーによる投稿の自動削除からピン留めされた投稿を保護できるようになります。重要な情報をピン留めしている場合に、誤って削除されることを防ぐことができるようになります。

## アップグレード時の注意事項

v10.10では、いくつかのデータベーススキーマ変更が行われますが、アップデート時のダウンタイムは発生しない見込みです。  
詳しくは以下ページの `v10.10` の項を参照ください。

[Important Upgrade Notes \- Mattermost documentation](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

## その他のトピック

### Extended Support Release (ESR) ライフサイクルの更新

[Extended Support Release lifecycle update: Enhanced upgrade experience for admins](https://mattermost.com/blog/extended-support-release-lifecycle-update-enhanced-upgrade-experience-for-admins/)

2025年8月のリリース（サーバー v10.11、デスクトップアプリ v5.13）から、MattermostのESR（Extended Support Release）のライフサイクルが以下のように変更されます：

- **リリース間隔の延長**: ESRのリリースサイクルが6ヶ月から9ヶ月に変更
- **サポート期間の延長**: サポート期間が9ヶ月から12ヶ月に延長

この変更により、セキュリティ修正と重要度の高いバグ修正が、サポート対象のESRに対して1年間バックポートされることになり、システムの安定性に関してより安心できる環境が提供されるようになります。

現在のESRリリースであるMattermost Server v10.5 とデスクトップアプリ v5.11 は、新しいサイクルに移行する前に 2025年11月15日まで引き続きサポートされる予定です。

## おわりに

次の `v10.11` のリリースは 2025/08/15(Fri) を予定しています。
