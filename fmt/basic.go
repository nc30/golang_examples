/*
	docs:
		- https://go-tour-jp.appspot.com/welcome/1
		- https://pkg.go.dev/fmt
		- https://pkg.go.dev/io#Writer
*/

package main

import (
	"fmt"

	"io/ioutil"
	"log"
	"os"
)

func init() {
	// ひとまず覚えなくていいです
	log.SetFlags(0)
}

/*
fmtの出力関数は基本的にstringを作るものです。

fmt出力関数には入力方式として2種類、出力先として2種類の合計4種類を覚えておけばいい。
他にもいくつかありますが、あまり使う機会は無いのでgodocを見ればいいです

	入力方式はこの2つ
		- *Println
		- *Printf

	出力先はこの2つ
		- Sprint*系
		- Fprint*系

	入力方式を軸に試してみよう
*/
func main() {
	var st string

	//// *Println系
	// これは与えられた引数をstringに変換し、スペースをくっつける。
	// 最後に改行がくっつくため、開発中に変数の確認として使うと手っ取り早い。
	// スレッドセーフではないし本番運用時に消さないといけないのでlog.Printlnのほうがいいですけどね
	world := "world"

	st = fmt.Sprintln("Hello", world, 123, nil)
	log.Print(st) // Hello world 123 <nil>

	//// *Printf系
	// 最初にフォーマットを指定して、それ以降の引数を割り当てる
	// いろいろな言語でよく見るやつ
	// 改行も自分で入れないといけないが、応用の幅は大きい
	// %sとか%dとかの詳細はドキュメントを見たほうが確実です
	// https://pkg.go.dev/fmt#pkg-overview

	st = fmt.Sprintf("Hello %s %d %T\r\n", world, 123, nil)
	log.Print(st) // Hello world 123 <nil>

	////// ここから出力先の話

	//// Sprint*系
	// 上の例でやりましたが、これは作ったstringを返り値として戻すもの。

	//// Fprint*系
	// これは第一引数にio.Writerを指定して作ったstringをutf8のバイナリとして書き込むもの
	// 第一引数はio.Writerならなんでもいい。
	// https://pkg.go.dev/io#Writer

	/// 標準出力に書き出す
	fmt.Fprintln(os.Stdout, "Hello", world, 123, nil) // Hello world 123 <nil>

	/// ファイルに書き出す
	// テスト用に一時ファイルを作成
	file, _ := ioutil.TempFile("", "example")

	fmt.Fprintln(file, "Hello", world, 123, nil)
	log.Println(file.Name(), "ファイルの中身を見てみよう")

	// Fprint*系の戻り値は書き込んだバイト数と書き込み時のエラーが返ってくる。
	p, err := fmt.Fprintln(os.Stdout, "I'm LumberJack It's OK.")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p) // 24
}
