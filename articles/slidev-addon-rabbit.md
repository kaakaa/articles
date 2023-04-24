---
title: "Slidevでプレゼン時間管理のためにウサギとカメを表示する"
emoji: "🐇"
type: "tech"
published: true
date: 2023-04-23T23:30:00+09:00
toc: true
tags: ["slidev", "rabbit"]
---


# はじめに

最近、Slidevでプレゼン時間管理のためにウサギとカメを表示するというのを作っていた。  
https://zenn.dev/kaakaa/articles/slidev-rabbit-turtle

![screen shot](https://blog.kaakaa.dev/images/posts/tech/slidev-rabbit-turtle/screen.gif)

Slidevには[Slidev Addon](https://sli.dev/addons/write-an-addon.html)という機能があり、自作のコンポーネントなどを使い回ししやすくする仕組みがあるらしい。  
ウサギとカメも簡単に適用できるようにSlidev Addon化してみた。

https://www.npmjs.com/package/slidev-addon-rabbit

# 使い方

1. Slidevプロジェクトを作成する (参考: [Getting Started \| Slidev](https://sli.dev/guide/))
2. 作成したSlidevプロジェクトにAddonを適用する (参考: [Use Addon \| Slidev](https://sli.dev/addons/use.html))
   - `npm install slidev-addon-rabbit`
   - `package.json` か SlidevのFrontmatterにAddonの設定を追記する

以上でAddonの適用は完了するらしい。簡単。

あとは`npm run dev`や`npm run build`でプレゼンを起動し、URL末尾に `?time=10` などを追加してプレゼン時間を設定するだけ。

# おわりに

Slidev Addonは、本当にSlidevプロジェクトの部分切り出しという感じだった。  
あと、初npm公開だった。
