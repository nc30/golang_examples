/*
   Context Key-Value Store Sample

   doc: https://pkg.go.dev/context
*/
package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	// contextに"silly"をキーとして"work"を入れる
	// contextは参照型ではないため戻り値を元の変数に書き換える
	ctx = context.WithValue(ctx, "silly", "work")

	// ctxから"silly"をキーとした中身を取り出す
	value := ctx.Value("silly")

	// Context.Value()の返り値はinterface{}なので値のキャストをする。
	str, ok := value.(string)
	if !ok {
		log.Println("value not found")
		return
	}
	log.Printf("silly: value=%+v", str) // silly: value=work

	// 上の省略形がこれ
	var ctxValue string
	if ctxValue, ok = ctx.Value("silly").(string); !ok {
		log.Println("value not found")
		return
	}
	log.Printf("silly: value=%+v", ctxValue) // silly: value=work

	// キーに該当するものがない場合はnilがインターフェイスとして戻ってくる
	v := ctx.Value("出鱈目なキー")
	log.Printf("出鱈目: value=%+v", v) // 出鱈目: value=<nil>
}
