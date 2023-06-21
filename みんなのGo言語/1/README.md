Goプロジェクトコード管理
[motemen/ghq](https://github.com/x-motemen/ghq) を使う。

https://zenn.dev/oreo2990/articles/13c80cf34a95af
https://zenn.dev/obregonia1/articles/e82868e8f66793

Goプロジェクトコード検索
[peco/peco](https://github.com/peco/peco) を使う。

https://qiita.com/vintersnow/items/08852df841e8d5faa7c2
https://qiita.com/ryoppy/items/4819d0911bf937dba821
https://zenn.dev/lilpacy/articles/68adb372f67e67

## Goらしいコードを書く
error をちゃんと使う。1.14で try の導入見送りになったっぽいので多値返却でいく
正規表現はできるだけ使わず、使っても regexp.MustCompile で確定させておく
map を避ける。できるだけ構造体を使って型を定義する
reflect を避ける。できるだけ型をつける
巨大な struct を作らず継承させようとしない。これは API の JSON Response を分割して作る時に思う
並行処理を使いすぎない
Go のコードを読もう
Go のバージョンは古いけど「GoのためのGo」も参考になるかも
テストとCI。 go vet や golint などでのチェックを入れる
ビルドとデプロイ。ビルド時の埋め込みやフラグ分岐など
モニタリング