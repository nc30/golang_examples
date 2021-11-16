/*
	- https://pkg.go.dev/net/http
	- https://httpbin.org/
*/
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func get() {
	r, err := http.Get("https://httpbin.org/status/200")
	if err != nil {
		// この場合のエラーはL6辺りまでの通信障害などのケースで
		// 404や500などのL7エラーはここでは拾えない
		// 試しにURLの200を変えたり、LANを引っこ抜いて動かしてみよう
		log.Fatal(err)
	}

	log.Println(r.StatusCode) // 200
	for key, value := range r.Header {
		log.Printf("%s: %s", key, value)
	}
	/*
		Access-Control-Allow-Credentials: [true]
		Date: [Mon, 15 Nov 2021 09:56:22 GMT]
		Content-Type: [application/json]
		Content-Length: [272]
		Server: [gunicorn/19.9.0]
		Access-Control-Allow-Origin: [*]
	*/

	// レスポンスはio.ReadCloserで返ってくる
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
	/*
		2021/11/15 18:53:24 {
		  "args": {},
		  "headers": {
		    "Accept-Encoding": "gzip",
		    "Host": "httpbin.org",
		    "User-Agent": "Go-http-client/2.0",
		    "X-Amzn-Trace-Id": "Root=1-61922e14-3971ae515f0ac14f6a7cd8db"
		  },
		  "origin": "***.***.***.***",
		  "url": "https://httpbin.org/get"
		}
	*/
}

func post() {
	payload := []byte(`{"id": 12, "name": "東方仗助"}`)
	body := bytes.NewBuffer(payload)

	// bodyはio.Readerで渡す
	r, err := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r.StatusCode) // 200
	for key, value := range r.Header {
		log.Printf("%s: %s", key, value)
	}
	/*
		Date: [Mon, 15 Nov 2021 10:05:45 GMT]
		Content-Type: [application/json]
		Content-Length: [528]
		Server: [gunicorn/19.9.0]
		Access-Control-Allow-Origin: [*]
		Access-Control-Allow-Credentials: [true]
	*/

	// レスポンスはio.ReadCloserで返ってくる
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))
	/*
		{
		  "args": {},
		  "data": "{\"id\": 12, \"name\": \"\u6771\u65b9\u4ed7\u52a9\"}",
		  "files": {},
		  "form": {},
		  "headers": {
		    "Accept-Encoding": "gzip",
		    "Content-Length": "34",
		    "Content-Type": "application/json; charset=utf-8",
		    "Host": "httpbin.org",
		    "User-Agent": "Go-http-client/2.0",
		    "X-Amzn-Trace-Id": "Root=1-619230f9-330019b61045709e1ccf8686"
		  },
		  "json": {
		    "id": 12,
		    "name": "\u6771\u65b9\u4ed7\u52a9"
		  },
		  "origin": "***.***.***.***",
		  "url": "https://httpbin.org/post"
		}
	*/
}

func main() {
	get()
	post()
}
