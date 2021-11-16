/*
	document: https://pkg.go.dev/net/http
*/
package main

import (
	"log"
	"net/http"
)

var body = []byte(
	`<http><head><title>hello</title></head>
<body><h1>Http Sample</h1><ul><li><a href="/ping">ping</a></li><li><a href="/panic">panic</a></li></ul>
</body></http>
	`)

// 標準ライブラリのみを使ったhttpサーバーの実装
// ルーティングはできてもMethodが何だろうと同じfunctionが呼ばれるため非常に使いづらい
// エンドポイントが一つだけみたいな本当に限定的な用途でないとサードパーティを使うかも
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(body)
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic(`コネクションごとにgoroutineをしているのでpanicが起こっても全体への影響はない\
			ただしそのままだとリクエストを投げた瞬間コネクションがぶつ切りされるのでrecoverはしたほうがいい`)
	})

	log.Println("Ctrl + Cで終了します")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Println(err)
	}

	// 起動してcurlなりブラウザでhttp://localhost:3000/にアクセスしてみよう！
}
