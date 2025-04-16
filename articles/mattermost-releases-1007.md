---
title: "Mattermost 10.7の新機能"
emoji: "🎉"
type: "tech"
published: true
date: 2025-04-16T23:00:00+09:00
toc: true
tags: ["mattermost", "releases"]
---

Mattermost 記事まとめ: https://blog.kaakaa.dev/tags/mattermost/

[Twitter: @mattermost_jp](https://twitter.com/mattermost_jp) で Mattermost に関する日本語の情報を提供しています。

# はじめに

2025/04/16 に Mattermost のアップデートとなる `v10.7.0` がリリースされました。  

本記事は、個人的に気になった新しい機能などを動かしてみることを目的としています。
変更内容の詳細については公式のリリースを確認してください。

- [v10 changelog \- Mattermost documentation](https://docs.mattermost.com/about/mattermost-v10-changelog.html#release-v10-7-feature-release)

---

各機能の見出し前の記号は、その機能が利用可能なエディションを表しています。

- [Professional](https://mattermost.com/pricing/)
- [Enterprise](https://mattermost.com/pricing/)

見出しの前に何もない場合、Free版も利用可能な機能です。

## URL入力時に全角句読点等が区切り文字として認識されない問題の修正

> Updated the marked package which includes full-width punctuation intervals for Unicode characters fix.

Mattermost内で利用されている`marked`パッケージがアップデートされ、URL文字列に続いて全角句読点を入力した場合に、句読点とそれに続く文字列もURLの一部として認識されてしまう問題が解消しました。

今までのバージョンでは、`https://mattermost.com、..`のようにURL文字列の直後に全角句読点が存在するような投稿を作成すると、句読点以降の文字も含めてURLとして認識されてしまい、意図とは異なるリンクが生成されてしまっていました

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.7/channels-marked-punctuation-fail.png)

この問題が本バージョンで解消され、全角句読点以降はURLとして認識されなくなりました。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.7/channels-marked-punctuation-success.png)


## iframe埋め込み時のFrame Ancestor設定

(あまり使用すべきでない機能な気もしますが、内容把握のために動かしてみたので記録として書き残します)

> Added a new System Console page called Embedding which allows frame ancestor domains to be specified when embedding Mattermost in other web sites. Note, teams.microsoft.com is no longer added automatically to the frame ancestors list.

`iframe`を使ってMattermostを別サイトに埋め込んで表示する際、埋め込みを可能とするサイトを[frame-ancestors](https://developer.mozilla.org/ja/docs/Web/HTTP/Reference/Headers/Content-Security-Policy/frame-ancestors)として指定できるようになりました。  
**※ iframeによる埋め込みはセキュリティリスクを伴うため、使用するとしてもイントラサイト内等の安全な環境でのみ利用するようにしましょう**

---

Mattermostは他のサイトに埋め込んで表示することができます。  
[Embed Mattermost](https://developers.mattermost.com/integrate/customization/embedding/)

今までのバージョンでも、iframeによる埋め込みは可能でしたが、利用するには以下のフォーラム投稿で書かれているようなnginx等の設定が必要でした。  
[Recipe: Embedding Mattermost in web applications using an iframe \[unsupported recipe\] \- Community \- Mattermost Discussion Forums](https://forum.mattermost.com/t/recipe-embedding-mattermost-in-web-applications-using-an-iframe-unsupported-recipe/10233)

Mattermostへのアクセス時に上記のような設定を行えない場合、以下のようなHTMLでMattermostをiframe埋め込みしても、アクセスがブロックされ、表示できませんでした。

```html:embed.html
<html>
  <body>
    <h1>Embedded Mattermost</h1>
    <div style="margin: 1px; padding: 5px; border: 4px solid #000; width: 600;">
      <iframe
        src="http://192.168.11.99:8065/test/channels/off-topic"
        width="600"
        height="00"
        frameborder="0"
        scrolling="no"
      ></iframe>
    </div>
  </body>
</html>
```

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.7/channels-frame-ancestors-fail.png)
> Refused to frame 'http://192.168.11.99:8065/' because an ancestor violates the following Content Security Policy directive: "frame-ancestors 'self'".

**システムコンソール > 統合機能 > Embedding > Frame Ancestors**にiframe埋め込みを行なっているサイトを追加することでアクセスできるようになります。  
今回は、前述のHTMLファイルに`http://192.168.11.99:8000/embed.html`でアクセスできるようにしているため、`http://192.168.11.99:8000`を追加します。(ここで設定された値はContents-Security-Policyの[`frame-ancestors`ディレクティブ](https://developer.mozilla.org/ja/docs/Web/HTTP/Reference/Headers/Content-Security-Policy/frame-ancestors)として扱われます。)


![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.7/channels-frame-ancestors-settings.png)

Frame Ancestorsの設定を行った後、HTMLファイルを再度開くと、ページ内にMattermostを埋め込んで表示できるようになります。

![alt text](https://blog.kaakaa.dev/images/posts/mattermost/releases-10.7/channels-frame-ancestors-success.png)


## アップグレード時の注意事項

あまり影響はなさそうですが、アップグレード時にDBのマイグレーションが走るようです。詳しくは以下ページの`v10.7`の項を参照ください。  

[Important Upgrade Notes \- Mattermost documentation](https://docs.mattermost.com/upgrade/important-upgrade-notes.html)

## その他のトピック

特になし

## おわりに

次の`v10.8`のリリースは 2025/05/16(Fri)を予定しています。  
