/*


- context with cancel

selain menambahkan value,kita juga bisa menambahkan sinya cancel ke context
kapan sinya cancel diperlukan dalam context?

bisanya kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke
proses tersebut

biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan
eksekusi goroutine, kita tinggal mengirim sinyal cancel ke contextnya

tetapi ingat goroutine yang menggunakan context tetap harus melakukan pengecheckan terhadap
contextnya

untuk membuat context dengan cancel sinyal kita bisa gunakan function
withCancel(parent)

dan ini akan menjadi child baru lagi ya, bukan berarti kita menambah sinyal cancelnya
tetapi beneran membuat child context baru

nah sekarang kita buat simulasinya dengan sekenario goroutine leak

jadi goroutine leak itu adalah goroutine yang berjalan terus dan tidak akan pernah berhenti
oke kita coba


*/

package context_with_cancel

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func CreateCounterWithGoroutineLeak() chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		counter := 1
		for {
			channel <- counter
			counter++
		}
	}()

	return channel
}

func TestGoroutineLeak(t *testing.T) {

	fmt.Println("goroutine yang berjalan diawal", runtime.NumGoroutine())
	dataChannel := CreateCounterWithGoroutineLeak()
	fmt.Println("goroutine yang berjalan", runtime.NumGoroutine())

	for data := range dataChannel {
		if data == 10 {
			break
		}
		fmt.Println(data)
	}

	fmt.Println("goroutine yang berjalan", runtime.NumGoroutine())

	/*

		kalo kita jalankan maka akan muncul seperti ini

		goroutine yang berjalan diawal 2
		goroutine yang berjalan 3
		1
		2
		3
		4
		5
		6
		7
		8
		9
		goroutine yang berjalan 3

		nah kenapa yang awalnya goroutine 2 setelah program selesai dijalankan
		kenapa 3? harusnya kan harus kembali ke 2 lagi

		nah karena kita menggunakan looping tanpa henti pada function CreateCounterWithGoroutineLeak
		maka sebenarnya goroutine tersebut tetap menjalankan loopingnya

		meskipun kita break pada looping penerima data channelnya
		sebenarnya breaknya hanya berlaku pada looping itu saja

		looping yang didalam createcounterWithGoroutineLeak tetap berjalan
		makanya diakhir goroutine yang berjalan tetap 3

		nah ini berbahaya, bayangkan kalian bikin applikasinya berupa website
		dan tiap request ternyata ada goroutine leak satu aja

		kalo misalkan ada 1000 request maka akan ada 1000 goroutine juga
		dan makin banyak goroutine leak makin lambat juga applikasi kita
		dan itu bisa memakan banyak memory

		nah ini kita bisa atasi dengan menggunakan context with cancel



	*/

}

/*

jadi didalam context itu ada sebuah function Done
lalu dia akan mengembalikan channel yang isinya struct kosong

nah kita bisa gunakan function Done untuk menangkap sinyal cancel
kita bisa gunakan select

karena kan done itu mengembalikan channel,
jadi kalo misalkan done tersebut mengembalikan data channel

artinya sinyal cancel itu dikirim dan proses kita harus berhenti
nah kita coba

*/

func CreateCounterWithGoroutineLeakFix(ctx context.Context, wg *sync.WaitGroup) chan int {
	channel := make(chan int)
	wg.Add(1)

	go func() {

		defer close(channel)
		counter := 1

		// jadi seperti ini ya, kita tinggal tambahkan select
		// kalo misalkan function done mengembalikan data channel maka proses kita harus berhenti
		// kita coba return seperti itu agar goroutinenya berhenti
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				channel <- counter
				counter++
			}
		}
	}()

	return channel
}

func TestGoroutineLeakFix(t *testing.T) {

	/*

		nah untuk mengirimkan sinyal cancel ke function done kalian bisa
		tulis seperti ini

		jadi kalian harus bikin parent context dulu kemudian bikin
		context with cancel

		dimana dia mengembalikan 2 buah return value, ada contextnya
		dan function cancel

		ketika function cancel itu dipanggil maka secara otomatis akan mengirimkan
		struct kosong ke function done

		dan function done akan mereturn channel isinya struct kosong juga

	*/

	var parent = context.Background()
	ctx, cancel := context.WithCancel(parent)
	var wg sync.WaitGroup

	fmt.Println("goroutine yang berjalan diawal", runtime.NumGoroutine())
	dataChannel := CreateCounterWithGoroutineLeakFix(ctx, &wg)
	fmt.Println("goroutine yang berjalan", runtime.NumGoroutine())

	for data := range dataChannel {
		if data == 10 {
			break
		}
		fmt.Println(data)
	}
	cancel()
	wg.Wait()

	fmt.Println("goroutine yang berjalan", runtime.NumGoroutine())

	/*

		perlu kalian ingat bahwa kita menggunakan goroutine, jadi kalo misalkan kalian cancel saja
		tanpa menggunakan wait grub, maka program akan terlanjut untuk berhenti dan tidak sempat
		melakukan cancel

		jadi kalian tetap menggunakan wait grub ya untuk menunggu goroutinenya selesai dijalankan
		dan untuk menunggu juga function cancel itu selesai mengirim sinyal cancel ke function done

		kalo kita jalankan maka akan muncul seperti ini

		goroutine yang berjalan diawal 2
		goroutine yang berjalan 3
		1
		2
		3
		4
		5
		6
		7
		8
		9
		goroutine yang berjalan 2

		jadi awalnya 2, lalu ketika goroutinenya berjalan mejadi 3
		dan ketika goroutinenya selesai dijalankan dengan mengirimkan sinyal cancel
		maka goroutinenya akan mati dan tidak akan menyebabkan goroutineleak

		itu itulah penggunakan dari context with cancel


		jadi nanti kalo kalian belajar tentang golang database, golang web
		itu semuanya hampir ada contextnya

		artinya kalo misalkan kalian mengirim query ke database dan ditengah jalan
		kalian ingin membatalkan karena ada masalah

		kalian tinggal kirim sinyal cancelnya, nanti secara otomatis querynya akan dibatalkan
		oleh golang databasenya

		jadi dengan context with cancel, kalian bisa memastikan bahwa tidak ada goroutine leak
		pada program kalian

		jadi seperti itu
		mudah mudahan kalian paham

	*/
}
