/*
	docs:
		- https://pkg.go.dev/golang.org/x/sync/semaphore
*/
package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

// これまでの例では固定数の処理を並列処理していたが、
// これが不定数であった場合、その数がそのまま並列で動いてしまう。
// 例えば30の処理を並列でこなしたい場合、
// 一度に30を処理すると逆にパフォーマンスが落ちる可能性もある
//
// そこでWaitGroupの変わりにsemaphoreを使うことで同時に走る処理をコントロールする
func main() {
	var MaxWorkers int64 = 2

	// 同時処理数制御機能付きのWaitGroup的なもの
	// この引数(int64)で同時に動かせる数を定義する
	sem := semaphore.NewWeighted(MaxWorkers)

	ctx := context.Background()

	for i := 0; i < 30; i++ {
		// waitGroup.Add()に相当
		// semaphoreに処理の開始を宣言する
		// 第２引数の1は宣言する処理数。基本1固定でいい
		// 処理上限内で空きができるまでブロックされる
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(number int) {
			// waitGroup.Done()に相当
			// これが重要。
			// ここで処理が終わったことを通知すると、
			// 空きができたと判断され上のブロックが解除される。
			//
			// 整合性にズレが出ないよう、goroutine内のできるだけ早い段階で
			// 必ずdeferを使ってリリースする
			defer sem.Release(1)

			// 適当に待つ
			log.Println("starting number:", number)
			time.Sleep(time.Millisecond * 500)
			log.Println("end of number:", number)
		}(i)
	}

	// waitGroup.Wait()に相当
	// 同時処理上限分の空きが確保できる = 処理が全部終了したと判断できる
	sem.Acquire(ctx, MaxWorkers)
	log.Println("all done")

	/*
		2021/12/27 14:38:41 starting number: 1
		2021/12/27 14:38:41 starting number: 0
		2021/12/27 14:38:42 end of number: 0
		2021/12/27 14:38:42 end of number: 1
		2021/12/27 14:38:42 starting number: 2
		2021/12/27 14:38:42 starting number: 3
		2021/12/27 14:38:42 end of number: 2
		2021/12/27 14:38:42 starting number: 4
		2021/12/27 14:38:42 end of number: 3
		2021/12/27 14:38:42 starting number: 5
		2021/12/27 14:38:43 end of number: 4
		2021/12/27 14:38:43 starting number: 6
		2021/12/27 14:38:43 end of number: 5
		2021/12/27 14:38:43 starting number: 7
		2021/12/27 14:38:43 end of number: 6
		2021/12/27 14:38:43 end of number: 7
		2021/12/27 14:38:43 starting number: 8
		2021/12/27 14:38:43 starting number: 9
		2021/12/27 14:38:44 end of number: 8
		2021/12/27 14:38:44 starting number: 10
		2021/12/27 14:38:44 end of number: 9
		2021/12/27 14:38:44 starting number: 11
		2021/12/27 14:38:44 end of number: 10
		2021/12/27 14:38:44 starting number: 12
		2021/12/27 14:38:44 end of number: 11
		2021/12/27 14:38:44 starting number: 13
		2021/12/27 14:38:45 end of number: 12
		2021/12/27 14:38:45 starting number: 14
		2021/12/27 14:38:45 end of number: 13
		2021/12/27 14:38:45 starting number: 15
		2021/12/27 14:38:45 end of number: 14
		2021/12/27 14:38:45 starting number: 16
		2021/12/27 14:38:45 end of number: 15
		2021/12/27 14:38:45 starting number: 17
		2021/12/27 14:38:46 end of number: 16
		2021/12/27 14:38:46 starting number: 18
		2021/12/27 14:38:46 end of number: 17
		2021/12/27 14:38:46 starting number: 19
		2021/12/27 14:38:46 end of number: 18
		2021/12/27 14:38:46 starting number: 20
		2021/12/27 14:38:46 end of number: 19
		2021/12/27 14:38:46 starting number: 21
		2021/12/27 14:38:47 end of number: 20
		2021/12/27 14:38:47 starting number: 22
		2021/12/27 14:38:47 end of number: 21
		2021/12/27 14:38:47 starting number: 23
		2021/12/27 14:38:47 end of number: 22
		2021/12/27 14:38:47 starting number: 24
		2021/12/27 14:38:47 end of number: 23
		2021/12/27 14:38:47 starting number: 25
		2021/12/27 14:38:48 end of number: 25
		2021/12/27 14:38:48 end of number: 24
		2021/12/27 14:38:48 starting number: 26
		2021/12/27 14:38:48 starting number: 27
		2021/12/27 14:38:48 end of number: 27
		2021/12/27 14:38:48 starting number: 28
		2021/12/27 14:38:48 end of number: 26
		2021/12/27 14:38:48 starting number: 29
		2021/12/27 14:38:49 end of number: 28
		2021/12/27 14:38:49 end of number: 29
		2021/12/27 14:38:49 all done
	*/
}
