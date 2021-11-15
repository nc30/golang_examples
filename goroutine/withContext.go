/*
	- https://pkg.go.dev/context
*/
package main

import (
	"context"
	"log"
	"time"
)

func basic() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		// これが重要
		defer cancel()


		log.Println("waiting 3 seconds...")
		time.Sleep(time.Second * 3)
		log.Println("done!")
	}()


	<-ctx.Done()

	/*
		2021/11/15 19:21:29 waiting 3 seconds...
		2021/11/15 19:21:32 done!
	*/
}


// 通信系の処理などタイムアウトが必要な場合はcontextのTimeout機能を使うといい
// WithTimeoutやSleepの時間を変えて挙動を調べてみよう
func withTimeout() {
	ctx := context.Background()

	// withCancelの変わりにWithTimeoutを使用する。
	ctx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	go func() {

		defer cancel()


		log.Println("waiting 3 seconds...")
		time.Sleep(time.Second * 3)
		log.Println("done!")
	}()


	<-ctx.Done()

	// contextがDoneになった理由はctx.Err()でわかる
	log.Println(ctx.Err())
	// 処理成功の可否はこれを比較することで行える
	log.Println(ctx.Err() == context.Canceled)

	/*
		2021/11/15 19:32:17 waiting 3 seconds...
		2021/11/15 19:32:20 done!
		2021/11/15 19:32:20 context canceled
		2021/11/15 19:32:20 true
	*/
}


func main() {
	basic()
	withTimeout()
}
