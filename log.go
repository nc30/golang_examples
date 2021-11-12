/*
   log package Sample

   document:
    - https://pkg.go.dev/log
    - https://pkg.go.dev/fmt
*/
package main

import (
	"errors"
	"log"
	"os"
)

// logパッケージのサンプル
// Goの技術書などではよくfmt.Println()を使っているが
// あれはただ単純に書き出しているだけのため、
// 並列処理中に同時に書き込んだ場合にそれぞれが混ざり合ってしまう
// logは一つのハンドルに対してはスレッドセーフとなっているためこっちを使ったほうが安心
// あと時間が表示されるのはデバッグ中にとても助かる
func basic() {
	log.Println("spam")                     // 2021/01/19 18:47:35 spam
	log.Println("egg", "bacon", "and spam") // 2021/01/19 18:51:04 egg bacon and spam

	// Printfはfmtと同じ
	log.Printf("I'm a %s It's Ok", "Lumberjack") // 2021/01/19 19:53:54 I'm a Lumberjack It's Ok
	log.Printf("int: %d", 253)                   // 2021/01/19 19:51:39 int: 253
	log.Printf("hex: 0x%x", 253)                 // 2021/01/19 19:51:39 hex: 0xfd
	log.Printf("oct: 0o%o", 253)                 // 2021/01/19 19:51:39 oct: 0o375
	log.Printf("bin: 0b%b", 253)                 // 2021/01/19 19:51:39 bin: 0b11111101

	s := struct {
		ID   int
		Name string
	}{123, "Graham"}

	// 構造体のダンプ時に便利
	log.Printf("%+v", s) // 2021/01/19 19:50:00 {ID:123 Name:Graham}

	log.SetPrefix("[log] ")
	log.Println("プレフィックスをつける") // [log] 2021/01/19 18:50:07 プレフィックスをつける
	log.SetPrefix("")

	log.SetFlags(log.Flags() | log.LUTC)
	log.Println("時刻タイムゾーンはデフォルトを使用するが、フラグを追加することによりUTCにできる") // 2021/01/19 09:57:09 時刻タイムゾーンはデフォルトを使用するが、フラグを追加することによりUTCにできる

	log.SetFlags(0)
	log.Println("フラグをすべて外すと時間出力をオフにできる") // フラグを設定すると時間表示をオフにできる

	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("マイクロ秒まで表記") // 2021/01/19 18:57:09.086480 マイクロ秒まで表記

	// このロガーのデフォルト出力先は標準エラー出力(os.Stderr)である。
	// この関数で出力先を変える。
	// 引数はio.Writerならなんでもいい
	log.SetOutput(os.Stdout)

	var err error = nil
	if err != nil {
		// エラー内容を出力してstatus code 1で終了する。
		// 正直使う機会はない
		log.Fatal(errors.New("何かしらのエラー"))
	}
}

// pythonで言うところのNullHandler的な
// https://docs.python.org/ja/3/library/logging.handlers.html#logging.NullHandler
type NoneWriter struct{}

func (n *NoneWriter) Write(o []byte) (int, error) { return len(o), nil }

var DEBUG = false

// 独自のハンドラを作成する。
// 上で呼び出していたのはlogパッケージが自動で作成したデフォルトのロガーである
// 新たにハンドラを作成することにより、例えばエラーレベルによってハンドラを変えたり、
// フラグを全てオフにしてファイル書き込みに使ったりと大抵のアウトプット処理はなんとかなる
func handlers() {
	logger := log.New(os.Stdout, "[stdout] ", log.LstdFlags)

	logger.Println("I sleep all night and I work all day") // [stdout] 2021/11/11 11:42:38 I sleep all night and I work all day
	log.Println("I cut down trees, I skip and jump")       // 2021/11/11 11:42:38.997646 I cut down trees, I skip and jump

	// デバッグ用のハンドラを予め分けておいて、
	// 運用モード時は出力先を変えるようにすればいちいちデバッグコードを消さなくて済む
	var debug *log.Logger
	if DEBUG {
		debug = log.New(os.Stderr, "[DEBUG] ", log.LstdFlags)
	} else {
		debug = log.New(&NoneWriter{}, "[DEBUG] ", log.LstdFlags)
	}
	debug.Println("何かしらのデバッグ文言")
}

func main() {
	basic()
	handlers()
}
