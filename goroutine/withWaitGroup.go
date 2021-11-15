/*
	- https://pkg.go.dev/sync#WaitGroup
*/
package main

import (
	"log"
	"time"
	"sync"
)

// httpリクエストの並列処理など、
// 同じ処理を複数並列にしたい場合はsync.WaitGroupが便利
// このままではエラー処理までハンドリングしづらいのが難点
func main() {
	var wg sync.WaitGroup

	list := []string {
		"https://httpbin.org/get",
		"https://yahoo.co.jp/",
		"https://google.com/",
	}

	for _, url := range list{

		// 開始前にAdd(1)する
		wg.Add(1)


		// 名前空間の関係でrangeでとったurlを直接渡すとurlの値が意図したものと変わってしまう
		// その場合このように引数としてurlを渡さないといけない
		// 試しにgo fund()の引数の変数名を変えてみよう
		go func(url string) {

			// goroutine終了時にDone()するようにする
			defer wg.Done()


			// http処理は今回は無関係なのでリクエストを投げてるフリ
			log.Printf("requesting to %s", url)
			time.Sleep(time.Second * 3)
			log.Println("done!")
		}(url)
	}

	// Add()した数字の合計 - Done()した回数が0になるまで待ち続ける。
	// 試しにforの手前辺りで一回wg.Add(1)してみよう。
	// 存在しない最後のDone()を待ち続けて永遠に待ち続けるはず。(Ctrl+Cで強制終了)
	wg.Wait()


	/*
		2021/11/15 19:21:29 waiting 3 seconds...
		2021/11/15 19:21:32 done!
	*/
}
