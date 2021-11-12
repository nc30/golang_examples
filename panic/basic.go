/*
   panic basic sample
*/
package main

import (
	"errors"
	"log"
)

// Panicは他言語で言うところの例外ではないが、一応panicを拾うことはできる。
func basicPanic() {
	defer func() {
		// defer内でrecover()を呼ぶことによってpanicを握りつぶすことができる。
		e := recover()
		if e != nil {
			str, _ := e.(string)
			log.Println("panic:", str) // 2021/11/11 12:00:55 panic: example panic
		}
	}()
	defer func() {
		log.Println("panicが起こってもdeferは動く")
	}()

	panic("example panic")

	log.Println("panicが起こった時点でdeferの消化が始まり、recoverで拾われない限りその後の処理は行われない")

	/*
	2021/11/11 12:27:17 *****basicPanic()*****
	2021/11/11 12:27:17 panicが起こってもdeferは動く
	2021/11/11 12:27:17 panic: example panic
	*/
}

func panicTiming() {
	defer func() {
		e := recover()
		if e != nil {
		} else {
			log.Println("一度recoverされるともうrecoverされない")
		}
	}()

	defer func() {
		e := recover()
		if e != nil {
			log.Println("呼び出した関数内でのpanicも受け取れる")
		}
	}()

	func() {
		panic("panic!")
	}()

	/*
	2021/11/11 12:30:37 *****panicTiming()*****
	2021/11/11 12:30:37 呼び出した関数内でのpanicも受け取れる
	2021/11/11 12:30:37 一度recoverされるともうrecoverされない
	*/
}

// recoverで戻ってくるものはpanicの引数によって変わる
func panicValue() {
	re := func() {
		e := recover()
		if e != nil {
			switch typ := e.(type) {
			case error:
				log.Println("type: error,", typ)
			case int:
				log.Println("type: int,", typ)
			case string:
				log.Println("type: string,", typ)
			default:
				log.Printf("unknown type of, %+v", typ)
			}
		}
	}

	func() {
		defer re()
		panic("some Error")
	}()

	func() {
		defer re()
		panic(errors.New("some Error"))
	}()

	func() {
		defer re()
		panic(12345)
	}()

	func() {
		defer re()

		var someValues []string
		_ = someValues[1]
	}()

	/*
	2021/11/11 12:27:17 *****panicValue()*****
	2021/11/11 12:27:17 type: string, some Error
	2021/11/11 12:27:17 type: error, some Error
	2021/11/11 12:27:17 type: int, 12345
	2021/11/11 12:27:17 type: error, runtime error: index out of range [1] with length 0
	*/
}

func main() {
	log.Println("*****basicPanic()*****")
	basicPanic()
	log.Println("*****panicTiming()*****")
	panicTiming()
	log.Println("*****panicValue()*****")
	panicValue()
}
