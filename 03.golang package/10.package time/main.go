/*

- package time

disini kita akan belajar mengenai package time
bukan hanya mengetahui waktu tetapi kita juga dapat mengetahui tanggal

lansung saja kita coba



*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// mengetahui waktu sekarang
	var waktuSekarang time.Time = time.Now()
	fmt.Println(waktuSekarang)

	// mengetahui waktu tertentu
	fmt.Println(waktuSekarang.Nanosecond())
	fmt.Println(waktuSekarang.Second())
	fmt.Println(waktuSekarang.Minute())
	fmt.Println(waktuSekarang.Hour())
	fmt.Println(waktuSekarang.Month())
	fmt.Println(waktuSekarang.Year())
	fmt.Println(waktuSekarang.Date())

	// membuat waktu

	var membuatWaktu = time.Date(2021, 9, 2, 3, 0, 0, 0, time.UTC)
	fmt.Println(membuatWaktu)

	// mengubah string tanggal ke format tanggal
	var layout string = time.RFC3339
	var parse, _ = time.Parse(layout, "2022-10-01")
	fmt.Println(parse)

	/*

		nah cara membuat layout digolang itu berbeda, biasanya layoutnya seperti ini

		YY-MM-dd

		akan tetapi digolang tidak begitu, digolang layoutnya sudah ada format yang standart  yaitu RFC3339
		kalo kalian ctrl + cick tulisan RFC3339

		maka itu ada sebuah constanta yang isinya 2006-01-02T15:04:05Z07:00
		jadi formatnya sudah ada seperti itu ya

		formatnya adalah tanggal,
		atau kalian bisa menuliskannya seperti ini

		2006-01-02

		kalian tuliskan layoutnya seperti itu juga bisa
		jadi digolang agak sedikit unik ya

		oke jadi seperti itu ya package time digolang
		mudah muahan kalian paham



	*/

}
