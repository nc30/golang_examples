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

	// Timezoneのデフォルトは/etc/localtime つまり動作環境のローカルタイムゾーンとなる
	// TZ=Australia/Sydney go run time/basic.go のようにタイムゾーンを指定して試してみよう
	now = time.Now()

	log.Println(now) // 2022-01-25 12:25:58.039466266 +0900 JST m=+0.000108249

	// UTCが欲しい場合はUTC()を使う
	log.Println(now.UTC()) // 2022-01-25 03:25:58.039466266 +0000 UTC

	// ローカルタイムゾーンが必要な場合はLocal()を使う
	// (別の方法でtime.Timeを作った場合など)
	log.Println(now.Local()) // 2022-01-25 12:25:58.039466266 +0900 JST m=+0.000108249

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
	var timezone *time.Location
	var err error

	log.Println("--- Get Timezone")
	// タイムゾーン(*time.Locale)を取得するにはtime.LoadLocation()を使用する
	// タイムゾーンは"UTC"でUTC、"Local"でローカルタイムゾーン
	// そしてIANAデータベースのものを引っ張ってくる。
	// ここにあるものなら取れるんじゃないでしょうか
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

	// ドキュメントを読む限りだと/usr/share/zoneinfo/にオリジナルのタイムゾーンを入れると
	// それも取れそうではありますが、基本的に悪手です

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
	tz, _ := time.LoadLocation("MST")
	time.Local = tz

	// MST -25200
	log.Println(time.Now().Zone())
}

func main() {
	// 現在日時の取得
	getNow()
	// タイムゾーンの取得
	getTimeZone()
}
