/*
	- https://pkg.go.dev/context#WithCancel
*/
package main

import (
	"context"
	"time"
	"log"
)

// context.WithCancelはcontextのコピーと関数の組を返す
// 関数を呼び出すことで対となるcontextに通知を送ることができる
// 並列処理をする際の処理待ちとして使える
func basic() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()


	// 完了まで一秒かかる処理を並列作業させる
	go func() {
		log.Println("[routine1] processing.")
		time.Sleep(time.Second * 1)
		log.Println("[routine1] done.")
		cancel()
	}()

	// 上の処理が終了するのを待つ
	log.Println("[main] waiting routine1 done.")


	// cancel()が呼び出されるの待つ
	// context.Done()の正体は chan struct{}
	// チャンネルの戻り値自体には特に意味はない
	<-ctx.Done()


	log.Println("[main] routine1 done received.")

	/*
		2021/11/16 10:39:51 [main] waiting routine1 done.
		2021/11/16 10:39:51 [routine1] processing.
		2021/11/16 10:39:52 [routine1] done.
		2021/11/16 10:39:52 [main] routine1 done received.
	*/
}


// context.Done()はほぼ同時に伝播するので待ち側が増えても問題ない
// 今度はgoroutine1の終了をgoroutine2とgoroutine3で待ってみる
func multiWait() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)


	// 完了まで一秒かかる処理を並列作業させる
	go func() {
		log.Println("[routine1] processing.")
		time.Sleep(time.Second * 1)
		log.Println("[routine1] done.")
		cancel()
	}()


	// goroutine1を待つgoroutine
	go func() {
		log.Println("[routine2] waiting routine1 done.")
		<-ctx.Done()
		log.Println("[routine2] routine1 done received.")
	}()


	// goroutine1を待つgoroutineその2
	go func() {
		log.Println("[routine3] waiting routine1 done.")
		<-ctx.Done()
		log.Println("[routine3] routine1 done received.")
	}()


	// 決め打ちで上の処理が終わるまでブロック
	time.Sleep(time.Millisecond * 1100)

	/*
		2021/11/16 10:55:15 [routine1] processing.
		2021/11/16 10:55:15 [routine2] waiting routine1 done.
		2021/11/16 10:55:15 [routine3] waiting routine1 done.
		2021/11/16 10:55:16 [routine1] done.
		2021/11/16 10:55:16 [routine2] routine1 done received.
		2021/11/16 10:55:16 [routine3] routine1 done received.
	*/
}

func main() {
	basic()
	log.Println("**** FuncMultiWait ****")
	multiWait()
}
