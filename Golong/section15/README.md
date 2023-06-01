The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS


コマンドのドキュメントを参照


## バージョン確認
```
go version
```


## 環境変数の確認
```
go env
```


## ソースコードの整形

ソースコードを推奨される形式へ自動的に変換する

インデントなどの乱れがあるソースコードがあるディレクトリで
```
go fmt
```
を打つと、きれいに整形しなおしてくれる。できるだけ使った方がきれいなコードになる。

```
-n 実行されるコマンドの表示(ファイルは書き換えない)　対象のファイルを確認する時に使う
-x 実行されるコマンドの表示
```

## パッケージのドキュメントを参照する
```
go doc パッケージ名
```
例）
```
go doc math/rand
```


識別子（関数や定数、変数など）を指定
```
go doc パッケージ名.識別子
```
例）
```
go doc math/rand.Intn
```


## メソッド名の指定

例）timeパッケージのTime型のUnixのドキュメントを参照
```
go doc time.Time.Unix
```
```
-c 識別子のマッチングで大文字小文字を厳密に判定する
-u 非公開な識別子やメソッドについてもドキュメントを出力する
```



## プログラムのビルド
```
go build
```
ファイルやパッケージを指定しないビルド


```
例）app - main.go
          config.go
```


   appディレクトリ内で

   go build を実行した場合、app（Windowsではapp.exe)という実行ファイルが生成されれば成功

   この様に、ファイルやパッケージ名を指定しないビルドはカレントディレクトリ内の*.goファイルを対象にコンパイルする。

   ビルド結果として、カレントディレクトリの名前を持つ実行ファイルを生成する。

   １つのディレクトリ内には１つのパッケージのみ定義可能

## パッケージのビルド
```
例）src - foo - bar1.go   

               　　　bar2.go
```
```
go build app
```
   fooパッケージを構成するファイルがコンパイルされる。



## ファイルを指定するビルド

通常のgo build　では１つのディレクトリ内に複数パッケージを含む状態ではエラーを発生させる。個別にファイルを指定すればビルド可能
```
例）app - main.go  (package main)

     　　　 config.go(package main)

       　 　  foo.go   (package foo)
```
```
go build main.go config.go
```
   この場合生成される実行ファイル名は最初に指定したファイル名になる。(この場合main)



   実行ファイル名の指定の場合
```
go build -o appname main.go config.go
```


```
-x 内部処理の出力 go build -x
-o 生成する実行ファイル名の指定 go build -o ファイル名
```



## パッケージや実行ファイルをビルドした結果をGOPATH内にインストールする
```
go install パッケージ名
```
$GOPATHな/srcに置かれたパッケージのビルドした結果が実行ファイルであれば$GOPATH/binへ、それ以外であれば$GOPATH/pkgへインストールされる。
```
$GOPATH - bin (実行ファイルのインストール先)

                     pkg (ビルドしたパッケージのインストール先)

                     src (パッケージのソースコードのインストール先)
```




## 外部パッケージのダウンロードとインストールをまとめて実行する
```
go get パッケージ
```
GitHubなどのサービス上で開発されているGoのパッケージなどに使用することができる。

ダウンロード対象のパッケージがさらに別のパッケージに依存している場合でも、自動的に依存関係を抽出しつつ必要なパッケージのダウンロードとインストールを実行してくれる。

GitHubでは数多くのGoのパッケージが開発されていて、これらの資産を生かすためにもGitの導入は必須である。

Goを使用してAndroidやiOSで動作するアプリの開発をサポートするgolang/mobileがある。これらの拡張パッケージは実験的なものであったり、多くのバグを含むものなどが混在しているので、

使用には十分に注意が必要である。



例）

HTTPプロトコルの最新版であるHTTP2に対応した拡張パッケージをインストールする
```
go get golang.org/x/net/http2
```
このパッケージはgolang.orgによってホストされているため、パッケージにはgolang.org/x/net/http2を指定する。

何も出力されなければ成功。内部の処理を確認したい場合は-xオプションを付与すれば確認できる。

go install同様に$GOPATH内に保存される。

このパッケージを使用する場合は



import "golang.org/x/net/http2"



と指定する。

GitHubでホストされているパッケージも同様の手順で導入できる。

```
-d 対象パッケージのダウンロードのみ実行して停止する
-f 対象パッケージのパスから推測されるリポジトリへの検証をスキップする (-u指定時のみ）
-fix 対象パッケージの依存関係を解決する前にgo fixツールを適用する。
-insecure カスタムリポジトリを使用する場合に非セキュアなプロトコルの使用を許可する。（例：HTTP)
-t 対象パッケージに付属するテストが依存するパッケージも合わせてダウンロードする。
-u 対象パッケージの更新と依存パッケージの更新を検出して再ダウンロードとインストールを実行する。
```

## テストの実行
```
go test パッケージ名
```