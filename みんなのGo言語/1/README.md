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
- error をちゃんと返す。panicは使わない
- 正規表現はできるだけ使わずStringsパッケージを使う。使っても regexp.MustCompile 。
- map を避ける。構造体を使って型を定義する
```go
a := map[string]string(
  "foo": "bar",
  "baz": "hoge"
)

type data struct {
  foo string
  baz string
}
d := data{
  "foo": "bar",
  "baz": "hoge"
}
```
- reflect を避ける。できるだけ型をつける 5章にて
- 巨大な struct を作らず継承させようとしない。
- 並行処理を使いすぎない
- Go のコードを読もう
- テストとCI
- ビルドとデプロイ
- モニタリング