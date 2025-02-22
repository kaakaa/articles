---
title: "Mattermost 10.5の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-02-22T10:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2025/02/19 に Mattermost のアップデートとなる `v10.5.1` がリリースされました。  
毎月16日に新しいバージョンがリリースされる予定ですが、今回は`v10.5.0`にパフォーマンス等の問題が見つかったため、`v10.5.0`の正式なリリースは無く `v10.5.1`がリリースされた形です。  
また、先月リリースされた `v10.4.0` は、開発期間が年末年始を挟んだため、目立った変更がなく、本記事の執筆はしていませんでした。

本バージョンは[ESR(Extended Support Release)](https://docs.mattermost.com/upgrade/extended-support-release.html)に設定されており、2025/11/16までセキュリティ対応/バグ修正等のサポートが行われる予定です。(ESRでないバージョンのサポート期間は3ヶ月間です)


本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-5-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## Desktopアプリで画面がホワイトアウトする事象

Mattermost Desktopアプリを2025/02/14にリリースされた`v5.11.0`にアップデートすると、Mattermost画面がホワイトアウトするという事象が報告されていました。  
[Updating to the Version 5\.11\.0 created an issue · Issue \#3333 · mattermost/desktop](https://github.com/mattermost/desktop/issues/3333)

影響範囲や解消方法については以下のForumの投稿にまとめられています。  
[Mattermost Desktop App 5\.11 – Important Compatibility Notice \- Troubleshooting \- Mattermost Discussion Forums](https://forum.mattermost.com/t/mattermost-desktop-app-5-11-important-compatibility-notice/22599)

本事象については、接続先のMattermostサーバーのバージョンが現在のサポート対象外である`v9.4`以前である場合に発生する可能性があるそうで、
現時点での解消方法としてDesktopアプリを`v5.10.2`にダウングレードする方法が紹介されています。  

また、Desktop `v5.11.1`にて、互換性のないMattermostサーバーに接続した際にエラー画面を表示する機能が追加されるようです。  
[\[MM\-63224\] Add incompatible server screen by devinbinnie · Pull Request \#3348 · mattermost/desktop](https://github.com/mattermost/desktop/pull/3348)


## 投稿編集時に添付ファイルも編集可能に

本バージョンから、投稿の編集画面で添付ファイルの削除等が行えるようになりました。  
誤ってファイルを添付してしまった場合などでも、手軽にファイルだけ入れ替えることができるので地味ですが非常に助かる機能追加です。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.5/channels-edit-attachments.png)


## アップグレード時の注意事項

### Windows環境向けビルドの廃止

[Deprecation Notice: Server Builds for Microsoft Windows \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/deprecation-notice-server-builds-for-microsoft-windows/21498)

今回のリリースから、Windows環境向けのMattermostデプロイがサポート外となるそうです。  
Windows OS上に直接展開するのではなく、k8x/docker等を通じたデプロイへ移行するよう呼びかけられています。

### 手動でのプラグインデプロイの禁止

[Deprecation Notice: Manual Plugin Deployment \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/deprecation-notice-manual-plugin-deployment/21192)

今までのバージョンでは、Mattermostの資産が展開されたディレクトリ内に直接Mattermostプラグインファイルを配備することでMattermostプラグインを利用することができましたが、今回のバージョンから、その方法によるプラグインのデプロイができなくなりました。

引き続きWebUIのシステムコンソールからプラグインをアップロードする方法や、REST API等を使ったデプロイは実施可能なため、今後はそちらの方法を取るようにとのことです。

## その他のトピック

### PostgreSQLのサポートポリシーについて

Mattermostを構築する際に利用できるPostgreSQLのバージョンポリシーに関する記事が公開されていました。  
[Mattermost PostgreSQL Support Policy \- Mattermost](https://mattermost.com/blog/mattermost-postgresql-support-policy/)

Mattermostでは、基本的に[PostgreSQLコミュニティでサポートされているバージョン](https://www.postgresql.org/support/versioning/)をサポートする方針のようです。  
また、Mattermostでは[ESR(Extended Support Release)](https://docs.mattermost.com/about/release-policy.html#extended-support-releases)のバージョンを対象にサポートするPostgreSQLのバージョンを決定しているようで、2025/08/15にリリース予定のMattermostの次のESRである `v10.11`の時点では、PostgreSQL v13がminimum supported versionとなるため、来月リリースされるMattermost v10.6から、PostgreSQL v13以上のサポートに切り替わるそうです。  
[Prepare your Mattermost database \- Mattermost documentation](https://docs.mattermost.com/install/prepare-mattermost-database.html#minimum-supported-version-policy)

### Skype for Businessのサポート終了に関する記事

MicrosoftのSkype for Businessが2025/10/15にサポートを停止する件を受けて、その影響や移行先の紹介を行う記事がいくつか公開されています。  
[Mattermost Callsプラグイン](https://github.com/mattermost/mattermost-plugin-calls)も移行先の一つとして紹介されています。

* [Skype for Business Replacement: What You Need to Know](https://mattermost.com/blog/skype-for-business-replacement/)
* [The Risks of Not Replacing Skype for Business Before End\-of\-Life](https://mattermost.com/blog/risks-replacing-skype-for-business/)
* [4 Skype for Business Alternatives \| Mattermost](https://mattermost.com/blog/alternatives-to-skype-for-business/)
* [What to Look for in a Skype for Business Alternative \- Mattermost](https://mattermost.com/blog/what-to-look-for-in-skype-for-business-alternative/)

## おわりに

次の`v10.6`のリリースは 2025/03/14(Fri)を予定しています。  
