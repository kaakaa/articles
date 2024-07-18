---
title: "Mattermost 9.10の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2024-07-18T22:50:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2024/07/16 に Mattermost のアップデートとなる `v9.10.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v9 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v9-changelog.html#release-v9-10-feature-release)

本バージョンでの主な変更点を紹介する動画がMattermostの公式YouTubeチャンネルで公開されています。  
実際の動作を確認したい方は、こちらを参照ください。

https://www.youtube.com/watch?v=lDZQ_fZzM1I

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## ユーザープロフィール画面の更新

ユーザーアイコンをクリックすることで表示されるプロフィールポップオーバーの表示が少し更新されました。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.10/channels-profile-popover.png)

メールアドレスの先頭にアイコンが付いたり、各要素間のスペースが少し増えてスッキリした見た目となったようです。

## アクセシビリティの改善

絵文字ピッカーやプラグインによるチャンネルヘッダー部分にアクセシビリティの改善が行われたようです。

絵文字ピッカーでは、絵文字カテゴリをVoice Overした際の説明が、以前のバージョンでは`リンク、emoji_picker.smileys-emotion`のようなシステム的な文字列が含まれていたのが、`リンク、顔文字&表情`のような自然な言葉に修正されています。

![](https://blog.kaakaa.dev/images/posts/mattermost/releases-9.10/channels-a11y-emoji-picker.png)

## その他のトピック

### Mattermost v10

Mattermostもメジャーバージョンアップである **v10** のリリース予定日が2024/09/16に設定されたようです。  
[v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html)

Mattermost v10で廃止予定となる機能について、上記ページにいくつか記載されています。

**MySQL Databases**

Mattermost v10以降、MySQLを使用した新規環境はサポート対象外となり、基本的にはサポート対象バージョンのPostgreSQLを利用した環境のみがサポート対象となるようです。  
既存のMySQLを使用したMattermost環境は、(恐らく)Mattermost v10でもサポートされますが、Mattermost v11リリース時にサポートされなくなるようです。

**Apps Framework**

[Mattermost Apss Framework](https://zenn.dev/kaakaa/articles/mattermost-apps-sample)は、Mattermost v10でdeprecatedとなるようです。Mattermostの拡張機能を開発する際は、Webhook、スラッシュコマンド、OAuth2アプリ、プラグインのいずれかを使用するよう案内されています。

**PlaybookがEnterprise版限定機能に**

Mattermost v10と同時(?)にリリースされる[Mattermost Playbook Plugin](https://github.com/mattermost/mattermost-plugin-playbooks) v2.0以降を利用するには、[Enterpirseライセンス](https://mattermost.com/pricing/)が必要になるようです。Playbook v1.x系は、引き続き利用することができますがセキュリティ対応含め、更新は行われなくなるとのことです。

### User Survey Plugin

Mattermost公式ブログにUser Survey Pluginを紹介する記事が投稿されています。　　
[The User Survey plugin: A new way for customers to gain valuable insights from end users](https://mattermost.com/blog/mattermost-user-survey-plugin/)

User Survey Pluginは、Mattermost上のユーザーに対してNPS形式の質問をDMで送信し、ユーザーの満足度を確認する機能です。  
回答内容はMattermost上に保存され、過去のサーベイ結果含めCSV形式でダウンロードすることが可能なため、データを外部に置けない組織でも利用することができます。

User Survey Pluginは以下のリポジトリで公開されています。  
[mattermost/mattermost\-plugin\-user\-survey](https://github.com/mattermost/mattermost-plugin-user-survey)

## おわりに
次の`v9.11`のリリースは 2024/8/16(Fri)を予定しています。  
