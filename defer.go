package main

import (
	"log"
	"io/ioutil"
	"os"
	"errors"
)

func init() {
	log.SetFlags(0)
}

// ファイル処理の鉄則。ファイルを開いたら必ずcloseする。
// (今回は一時ファイルの削除も)
// ファイルを開いて処理した後にclose()してもいいが、
// ファイル処理の最中にエラーが起こった等でcloseの行に達する前に関数がreturnされる恐れがある。
// そのためtry{}finally{}なりで関数が終わるときに必ず実行されるようにしたほうがいい
// go言語でのtry-finallyに相当するのがdefer。
// 一見ややこしいように思えますが、インデントが深くならないのでコードが読みやすくなります
//
// ちなみにtry-finallyというのはphpで言うところのこういうの
//
// $fp = fopen("some/file.dat", "w");
// if($fp === false){return;}
// try {
//     $body = fread($fp);
// }finally{
//     fclose($fp);
// }
func basic_example() {
	//go言語ではこう書く
	fp, err := os.Open("some/file.dat")
	if err != nil {
		// ここでエラーが出たら touch README.mdしてね
		log.Fatal(err)
	}

	// このタイミングで宣言する
	defer fp.Close()

	// これ以降、関数が終了したタイミングでclose()が行われる
	_,_ = ioutil.ReadAll(fp)

	return
}

func basic(){
	// deferは新しいものから順番に処理される
	defer log.Println("1")
	defer log.Println("2")
	defer log.Println("3")

	log.Println("main process")

	/*
		main process
		3
		2
		1
	*/
}

// 例えば一時的なファイルを開いたとする。
// その場合ファイルクローズとファイル削除をしないといけない
// その場合こうする
func tempfile_function() {
	fp, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		log.Fatal(err)
	}

	// deferは無名関数でもOK
	defer func(){
		log.Println("doing close and remove")

		fp.Close()
		os.Remove(fp.Name())
	}()

	// 何かしらのファイル処理
	_, _ = ioutil.ReadAll(fp)

	log.Printf("this file is %s\r\n", fp.Name())

	/*
		this file is /tmp/454586748
		doing close and remove
	*/
}

// 関数が返す値をdeferで拾いたい場合はこのように書く
//              ⇓ここ重要
func cath_function() (err error) {
	// 無名関数で変数を参照させる
	defer func(){
		log.Println(err) // test error
	}()

	// こういう書き方だと、
	// Printlnの引数はその時点での変数の内容を参照するので拾えない
	defer log.Println(err) // <nil>



	err = errors.New("test error")

	return
	/*
		<nil>
		test error
	*/
}

func main() {
	log.Println("--- basic function")
	basic()

	log.Println("")
	log.Println("--- tempfile function")
	tempfile_function()

	log.Println("")
	log.Println("--- catch function")
	cath_function()
}