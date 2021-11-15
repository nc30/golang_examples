package main

import (
	"log"
	"time"
)

func main() {
	end := make(chan interface{}, 0)

	go func() {
		log.Println("waiting 3 seconds...")
		time.Sleep(time.Second * 3)
		log.Println("done!")

		end <- nil
	}() // 最後の()が重要


	// goroutineが途中だろうがメイン関数が終了すれば容赦なく途中で終了してしまう
	// そのためwait処理は必ず必要になる。
	// 一番簡単なのがこのchannelを使う方法。
	// 試しに<-endをコメントアウトして動かしてみよう
	// きっと"waiting ~"すら表示されずに終わるはず
	<-end

	/*
		2021/11/15 19:21:29 waiting 3 seconds...
		2021/11/15 19:21:32 done!
	*/
}
