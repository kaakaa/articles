---
title: "Mattermost 11.0の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-10-19T22:30:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

## はじめに

2025/10/16 に Mattermost のメジャーアップデートとなる `v11.0` がリリースされました。  
その後、パッチリリースが行われ、現時点（2025/10/19）では `v11.0.2` が最新となっています。

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。  
変更内容の詳細については、公式のリリースを確認してください。

- [Mattermost v11 Changelog — Mattermost documentation](https://docs.mattermost.com/product-overview/mattermost-v11-changelog.html)
- [Mattermost V11: Powering More Mission-Critical Collaboration](https://mattermost.com/blog/mattermost-v11-powering-more-mission-critical-collaboration/)

---

## Mattermost Entry Edition

> Introduced support for new default free edition Mattermost Entry with usage limits. See more details in this forum post.

Mattermost を無償で利用するための新しい形態として、Mattermost Entry Edition が導入されました。  
Entry Edition は Mattermost Enterprise Edition の一部で、サブスクリプション無しで本来は有償版で提供される機能（Professional 版相当）を使用できる代わりに、以下のような制限を持つプランです。

* 50 ユーザーまで
* メッセージ履歴は 10,000 件まで
* 1対1の音声通話は最大 40 分まで
* AI アシスタントへのクエリは 250 件/月まで

Mattermost Entry Edition について、詳しくは以下の公式エントリを参照してください。  
[Introducing Mattermost Entry](https://mattermost.com/blog/introducing-mattermost-entry/)

Mattermost Entry Edition で使用できる機能については、以下に記載されています。  
[Mattermost Plans \- Mattermost documentation](https://docs.mattermost.com/product-overview/plans.html)

## v11 における Breaking Changes

Mattermost v11 では、変更・廃止となる機能が多くあるため、v11 にアップデートする前に以下の Important Upgrade Notes を確認することをお勧めします。  
[Important Upgrade Notes \- Mattermost documentation](https://docs.mattermost.com/administration-guide/upgrade/important-upgrade-notes.html)

以下では、いくつかの廃止や制限の追加について紹介します。

### Mattermost Team Edition の制限追加

> GitLab SSO has been deprecated from Team Edition. Deployments using GitLab SSO can remain on v10.11 ESR (with 12 months of security updates) while transitioning to our new free offering Mattermost Entry, or can explore commercial/nonprofit options. See more details in this forum post.

> User limits were lowered to final threshold of 250 for Mattermost Team Edition (MIT-Compiled License).

Mattermost の無償版としては、MIT ライセンスの元で配布されている Team Edition がありますが、こちらについても以下の変更が行われます。

* ユーザー数の上限が 1,000 → 250 に変更
* 今まで無償版で使用できていた GitLab SSO の機能が廃止
* GitLab Omnibus として Mattermost を GitLab に同梱する形で配布する形態も廃止

Mattermost Entry Edition や Team Edition 等、無償版に対する影響については、以下の公式フォーラムの記事を参照してください。  
[Mattermost v11 Changes in Free Offerings \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/mattermost-v11-changes-in-free-offerings/25126)

特に、GitLab SSO を使って Mattermost を使用していた組織や、250 名以上の組織で無償版を運用していた場合に影響がありそうです。

### Playbook 機能の制限

> Playbooks has stopped working for Team Edition. Entry, Professional, Enterprise, and Enterprise Advanced plans are automatically upgraded to Playbooks v2 with no expected downtime. See more details in this forum post.

[以前からアナウンスがありましたが](https://blog.kaakaa.dev/post/mattermost-releases-910/#mattermost-v10)、[Playbook](https://mattermost.com/playbooks/) v2 以降の機能が Enterprise Edition 限定の機能となり、Team Edition では使用できなくなりました。Entry Edition からは、機能が制限された状態で利用できるようです。

詳しくは、以下の公式フォーラムの投稿を参照してください（Team Edition でも Playbook v1 であれば使用し続けられるとの記載があります）。  
[Clarification and Update on the Playbooks Plugin \(v11\) \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/clarification-and-update-on-the-playbooks-plugin-v11/25192)

### Bleve 検索機能の廃止

> Experimental Bleve Search functionality has been retired. If Bleve is enabled, search will not work until DisableDatabaseSearch is set to false. See more details in this forum post.

Bleve を使用した検索機能が廃止となります。  
今まで Bleve を使用していた場合、Mattermost v11 にアップグレードした後は [DisableDatabaseSearch](https://docs.mattermost.com/administration-guide/configure/environment-configuration-settings.html#disable-database-search) の設定を `false` にするまで検索が機能しなくなります。

詳しくは、以下の公式フォーラムの投稿を参照してください。  
[Transitioning from Bleve Search in Mattermost v11 \- Community \- Mattermost Discussion Forums](https://forum.mattermost.com/t/transitioning-from-bleve-search-in-mattermost-v11/22982)

### MySQL サポートの廃止

> Support for MySQL has ended. See more details in this forum post.

こちらも[以前からアナウンス](https://blog.kaakaa.dev/post/mattermost-releases-900/#mysql%E9%9D%9E%E6%8E%A8%E5%A5%A8%E5%8C%96%E3%81%AB%E5%90%91%E3%81%91%E3%81%9F%E6%B4%BB%E5%8B%95)されていましたが、Mattermost の DB として MySQL を使用できなくなりました。今後は PostgreSQL を使用する必要があります。

詳しくは以下の公式フォーラムの投稿を参照してください。  
[Transition to PostgreSQL \- Announcements \- Mattermost Discussion Forums](https://forum.mattermost.com/t/transition-to-postgresql/19551)


## (Enterprise Advanced) カスタムプロフィール属性による Attribute Based Access Control (ABAC)

> Admin managed Custom Profile Attribute fields can now be used as part of Attribute Based Access Control policies.

Mattermost 上で定義したカスタムプロフィール属性の設定値を元に、チャンネルに対するアクセスポリシーを作成できるようになりました。

今までのバージョンでも、AD/LDAP や SAML で管理されたユーザー属性を Mattermost のユーザー情報に同期し、そのユーザー属性値によりチャンネルへのアクセスポリシーを設定できたようですが（未確認）、本バージョンから Mattermost 上で作成できる Custom Profile Attribute の設定値を元に、チャンネルに対するアクセスポリシーを作成できるようになりました。

完全な機能を使用するには、[Enterprise Advanced](https://mattermost.com/pricing/) ライセンスが必要です。  
（Mattermost v11 から導入された Entry Edition でも一部機能を試せるので、それを元に紹介します。）


まず、カスタムプロフィール属性 の作成方法について紹介します。(Team Edition では使用できません。)

カスタムプロフィール属性は、各ユーザーのユーザープロフィールとして表示できる属性値を Mattermost 上で定義できる機能です。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-11.0/channels-cpa-profile.png)

カスタムプロフィール属性を定義するには、**システムコンソール > システム属性 > ユーザー属性** へ移動し、**+ 属性を追加** ボタンをクリックします。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-11.0/channels-cpa-definitions.png)

設定種別は `テキスト`、`電話番号`、`URL`、`選択(Select)`、`複数選択` から選ぶことができ、それぞれの属性に対して各ユーザーによる変更を許可するか（**Editable by users**）や、AD/LDAP または SAML のユーザー情報と同期するかどうかなどを設定することができます。

各ユーザーによる変更を許可（**Editable by users** を ON）にした場合、各ユーザーが自身のプロフィール設定画面で、カスタムプロフィール属性の値を変更できるようになります。各ユーザーによる変更を禁止した場合は、[`mmctl cpa`](https://docs.mattermost.com/administration-guide/manage/mmctl-command-line-tool.html#mmctl-cpa) コマンドから設定します（`mmctl cpa` コマンドは Enterprise Advanced プランでしか利用できません）。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-11.0/channels-cpa-settings.png)

チャンネルに対するアクセスポリシーに設定できる属性は、各ユーザーによる変更を禁止（**Editable by users** を OFF）にした属性のみとなります。

---

ここからは、カスタムプロフィール属性を使用したチャンネルに対するアクセスポリシーの設定方法について紹介します。

**システムコンソール > システム属性 > 属性ベースのアクセス** を開き、**"このサーバーで属性ベースのアクセス制御を許可する"** を有効にした後、**アクセス制御ポリシー** の **+ ポリシーを追加** ボタンをクリックします。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-11.0/channels-abac-definitions.png)

アクセス制御ポリシー名を入力し、**"アクセスルールに基づく自動メンバー追加"** の設定を行います。  
**"アクセスルールに基づく自動メンバー追加"** の設定を有効にすると、このアクセス制御ポリシーで設定した属性値と一致する属性値を持つユーザーアカウントが自動で設定されたチャンネルに追加され、当該属性値に一致しなくなると自動で削除されるようになります。ロールやチームなどによって必ず所属すべきチャンネル等がある場合、新しいメンバーが参画した際などに、毎回当該チャンネルへユーザーを招待するという操作が不要になります。

次に **"属性ベースのアクセスルール"** の **+ 属性を追加** ボタンをクリックすると、属性値に関する条件式を入力する項目が表示されるため、このアクセスポリシーで対象とする属性値に関するルールを設定します。カスタムプロフィール属性については、**"Editable by users"** が ON になっている属性は、ここで使用することはできません。

次に、**"割り当て済みチャンネル"** から、このアクセスポリシーによって管理される非公開チャンネルを **チャンネル追加** ボタンから追加します。ここでは、公開チャンネルは指定できず、非公開チャンネルのみが選択可能なので注意してください。

以上の設定を行い保存した後、**+ Run Sync Job** ボタンを押すなどして同期ジョブが実行されると、作成したアクセスポリシーによる各チャンネルに対するメンバーシップの操作が行われます。

![Alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-11.0/channels-abac-done.png)

---

## おわりに

Mattermost v11 のメジャーバージョンアップとなりましたが、どちらかというと機能廃止や制約の強化が目立ったように感じました。よりミッションクリティカルかつエンタープライズ領域に注力する姿勢が強く打ち出され、プロダクトとしては成熟してきていることを感じます。しかし、無償版で数百ユーザー程度の中規模デプロイメント向けには、制限の方が強く感じられるようになりました。

元々、MIT ライセンスで配布されていた Team Edition は、GitLab Omnibus に組み込むために求められたものと考えると、GitLab Omnibus に組み込まれなくなった今、存在意義が薄まっているような気もします。今後、無償版については Mattermost Entry Edition に集約されていくのではないかと感じました。

Slack も無償版に対しては過去メッセージの参照に関する制約などを強めており、メッセージングプラットフォームの分野においても **The Free Lunch is Over** の雰囲気を感じています。

次の `v11.1` のリリースは 2025/11/14(Fri) を予定しています。  
