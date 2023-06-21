package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	res, _ := http.Get("http://www.pixiv.net/")
	fmt.Println(res.StatusCode)             // 200
	fmt.Println(res.Proto)                  // HTTP/2.0
	fmt.Println(res.Header["Date"])         // [Wed, 31 May 2023 02:04:10 GMT]
	fmt.Println(res.Header["Content-Type"]) // [text/html; charset=UTF-8]
	fmt.Println(res.Request.Method)         // GET
	fmt.Println(res.Request.URL)            // https://www.pixiv.net/

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read the response body: ", err)
		return
	}
	fmt.Println(string(body))

	if err := ioutil.WriteFile("./README.md", body, 0644); err != nil {
		fmt.Println("Failed to write to file: ", err)
	}

	// POSTメソッド
	// vs := url.Values{}
	// // vs.Add("todos", "1")
	// // vs.Add("massage", "メッセージ")
	// fmt.Println(vs.Encode())

	// res2, _ := http.Post("https://jsonplaceholder.typicode.com/", "post/1")

	//応用
	//ヘッダーをつけたり、クエリをつけたり
	//Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")
	//クエリ の確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")
	//ResolveReference
	//クエリをくっつけたURLを生成する。
	//相対パスから絶対URLに変換する。
	//baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	//GET ver
	//リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	//まだリクエストはしていない。structを作っただけ。
	req, _ := http.NewRequest("GET", endpoint, nil)
	//requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)
	//URLのクエリを確認
	q := req.URL.Query()
	//クエリを追加
	q.Add("name", "test")
	//クエリを表示
	fmt.Println(q)
	//&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode())
	//encodeしてからURLに戻す
	//日本語などを変換する
	req.URL.RawQuery = q.Encode()
	//実際にアクセスする
	//クライアントを作る
	var client *http.Client = &http.Client{}
	//結果 アクセスする
	resp, _ := client.Do(req)
	//読み込み
	body1, _ := ioutil.ReadAll(resp.Body)
	//出力
	fmt.Println(string(body1))

}
