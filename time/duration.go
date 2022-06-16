/*
	timeパッケージを使った現在時刻の取得・タイムゾーン処理

	doc:
	  - https://pkg.go.dev/time#Duration
	  - https://pkg.go.dev/time
	  - https://xn--go-hh0g6u.com/pkg/time/
*/
package main

import (
	"time"

	"log"
)

func init() { log.SetFlags(0) }

func basic() {
	log.Println("## 基本的なtime.Durationの作成")

	// 時刻を表現するのは`time.Time`ですが、時間を表現するのは`time.Duration`です
	// pythonにおける`datetime.timedelta`
	var duration time.Duration

	// durationの作り方は、基本的にtimeで用意されている定数を
	// 掛けたり引いたりして作成する。

	duration = time.Second * 3
	log.Println(duration) // 3s

	duration = time.Minute * 3
	log.Println(duration) // 3m0s

	duration = time.Hour * 3
	log.Println(duration) // 3h0m0s

	duration = time.Hour*3 + time.Minute*5 + time.Second*32
	log.Println(duration) // 3h5m32s

	// よく躓くのがtime.Duration * 変数の場合。
	// コメントアウトして試してもらうとわかりますが、
	// こんなエラーがでて起動すらできません
	// time/duration.go:47:25: invalid operation: time.Second * d (mismatched types time.Duration and int)

	var d int = 3
	// duration = time.Second * d

	// これを回避するには倍率をtime.Durationにキャストして計算します
	duration = time.Second * time.Duration(d)
	log.Println(duration) // 3s

	// もしくは最初から倍率変数をtime.Durationとして扱うか
	// でもこのやり方だと謎の3ナノ秒の変数が存在することになるので混乱するかもしれませんね
	var dd time.Duration = 3

	duration = time.Second * dd
	log.Println(duration) // 3s
}

func fromString() {
	log.Println("## stringからの作成")
	var dur time.Duration
	var err error

	dur, _ = time.ParseDuration("1h10m10s1ms32µs54ns")
	log.Println(dur) // 1h10m10.001032054s

	// 表現できる単位は h, m, s, ms, µs, ns のみ
	// つまり30dとかを作ろうとするとエラーになる
	dur, err = time.ParseDuration("30d")
	log.Println(err.Error()) // time: unknown unit d in duration 30d
	log.Println(dur)         // 0s

	// そういう場合は地道にhourに計算し直すしかないですね
	dur, _ = time.ParseDuration("720h")
	log.Println(dur) // 720h0m0s
}

func main() {
	basic()
	fromString()
}
