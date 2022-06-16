/*
	timeパッケージを使った現在時刻の取得・タイムゾーン処理

	doc:
	  - https://pkg.go.dev/time
	  - https://xn--go-hh0g6u.com/pkg/time/
*/

package main

import (
	"time"

	"log"
)

// 現在日時の取得
func getNow() {
	var now time.Time

	log.Println("--- Get now")

	// goにおける時刻はtime.Timeで表現する。
	// pythonにおけるdatetime.datetimeで、jsで言うところのDate

	// Timezoneのデフォルトは/etc/localtime つまり動作環境のローカルタイムゾーンとなる
	// TZ=Australia/Sydney go run time/basic.go のようにタイムゾーンを指定して試してみよう
	now = time.Now()

	log.Println(now) // 2022-01-25 12:25:58.039466266 +0900 JST m=+0.000108249

	// UTCが欲しい場合はUTC()を使う
	log.Println(now.UTC()) // 2022-01-25 03:25:58.039466266 +0000 UTC

	// ローカルタイムゾーンでのオブジェクトが必要な場合はLocal()を使う
	// (別の方法でtime.Timeを作った場合など)
	// 試しにUTC版を作ってローカルに戻してみる
	log.Println(now.UTC().Local()) // 2022-01-25 12:25:58.039466266 +0900 JST

	// オブジェクトのタムゾーン取得
	log.Println(now.Zone()) // JST 32400

	// UnixTimeの取得 (int64であることに注意)
	log.Println(now.Unix()) // 1643085657

	// 年の取得
	log.Println(now.Year()) // 2022

	// 月の取得
	month := now.Month()
	log.Println(month, int(month)) // January 1

	// 日の取得
	log.Println(now.Day()) // 25

	// 時の取得
	log.Println(now.Hour()) // 12

	// 秒の取得 (not unixtime)
	log.Println(now.Second()) // 25
}

// タイムゾーンの取得
func getTimeZone() {
	log.Println("--- Get Timezone")

	var timezone *time.Location
	var err error

	// タイムゾーン(*time.Locale)を取得するにはtime.LoadLocation()を使用する
	// タイムゾーンは"UTC"でUTC、"Local"でローカルタイムゾーン
	// そしてIANAデータベースのものを引っ張ってくる。
	// ここにあるものなら取れるんじゃないでしょうか
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

	// ドキュメントを読む限りだと/usr/share/zoneinfo/にオリジナルのタイムゾーンを入れると
	// それも取れそうではありますが、どう考えても悪手

	timezone, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(timezone) // Asia/Tokyo

	timezone, err = time.LoadLocation("Egypt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(timezone) // Egypt

	// timeパッケージのデフォルトタイムゾーンはtime.Localで変更できる
	// 影響範囲が大きいため使用する場合は要注意
	timezone, _ = time.LoadLocation("MST")
	time.Local = timezone

	// MST -25200
	log.Println(time.Now().Zone())


	// time.Timeオブジェクトのタイムゾーンを変える場合は
	// time.Time.In()を使う
	now := time.Now()

	log.Println(now) // 2022-06-15 19:38:55.117055228 -0700 MST m=+0.000226941

	timezone, _ = time.LoadLocation("Australia/Sydney")
	log.Println(now.In(timezone)) // 2022-06-16 12:38:55.117055228 +1000 AEST
}

func main() {
	// 現在日時の取得
	getNow()
	// タイムゾーンの取得
	getTimeZone()
}
