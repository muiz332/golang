/*

- deadlock


sebelumnya kita sudah tahu bahwa ada masalah race condition ketika kita melakukan
concurrency atau paraller programming

nah sebenarnya ada lagi problem yang terjadi yaitu deadlock
ini biasanya terjadi kalo kita salah implementasi mutex atau locking

jadi hati hati saat kita membuat applikasi yang parller atau concurrent, masalah yang
sering kita hadapi adalah deadlock

deadlock adalah keadaan dimana sebuah prosess goroutine saling menunggu lock sehingga tidak
ada satupun goroutine yang bisa jalan

biar kalian kebayang kita coba simulasikan masalah deadlock



*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) ChangeAmount(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {

	user1.Mutex.Lock()
	fmt.Println("lock", user1.Name)
	user1.ChangeAmount(-amount)

	time.Sleep(1 * time.Second)

	user2.Mutex.Lock()
	fmt.Println("lock", user2.Name)
	user2.ChangeAmount(amount)

	time.Sleep(1 * time.Second)

	user1.Mutex.Unlock()
	user2.Mutex.Unlock()
}

func TransferFix(user1 *UserBalance, user2 *UserBalance, amount int) {

	user1.Mutex.Lock()
	fmt.Println("lock", user1.Name)
	user1.ChangeAmount(-amount)

	time.Sleep(1 * time.Second)
	user1.Mutex.Unlock()

	user2.Mutex.Lock()
	fmt.Println("lock", user2.Name)
	user2.ChangeAmount(amount)

	time.Sleep(1 * time.Second)
	user2.Mutex.Unlock()

}

func main() {

	user1 := UserBalance{
		Name:    "hasan",
		Balance: 10000,
	}

	user2 := UserBalance{
		Name:    "budi",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 1000)

	go Transfer(&user2, &user1, 2000)

	time.Sleep(5 * time.Second)
	fmt.Println("Balance user 1", user1.Balance)
	fmt.Println("Balance user 2", user2.Balance)

	/*

		pemanggilan function goroutine perbama
		hasan mentransfer 1 1000 ke budi

		maka hasilnya harusnya

		hasan balance = 9000
		budi balance = 11000

		pemanggilan function goroutine kedua
		budi mentransfer 2000 ke hasan

		maka hasilnya harusnya
		hasan balance = 11000
		budi = 9000

		akan tetapi kalo kita jalankan hasilnya seperti ni

		hasan balance = 9000
		budi balance 8000

		lock budi
		lock hasan
		Balance user 1 9000
		Balance user 2 8000

		kenapa seperti ini ?

		jadikan goroutine pertama itu melakukan lock terhadap user1 dan ingin melakukan lock
		terhadap user2

		akan tetapi sebelum goroutine melakukan lock terhadap user2, goroutine kedua sudah
		terburu melakukan locking terhadap user2 diparameter1

		dan user1 keburuh dilock oleh goroutine pertama, jadi saling menunggu\
		karena belum sempat melakukan unlock

		itu terjadi karena kalian tidak menulis unlocknya setelah melakukan perubahan
		unclock nya ditulis paling bawah sendiri

		maka terjadi deadlock seperti itu, saling menunggu
	*/

	user3 := UserBalance{
		Name:    "hasan",
		Balance: 10000,
	}

	user4 := UserBalance{
		Name:    "budi",
		Balance: 10000,
	}

	go TransferFix(&user3, &user4, 1000)

	go TransferFix(&user4, &user3, 2000)

	time.Sleep(5 * time.Second)
	fmt.Println("Balance user 3", user3.Balance)
	fmt.Println("Balance user 4", user4.Balance)

	/*

		akan tetapi kalo kalian melakukan mutex secara benar maka deadlock tidak akan
		terjadi

		jadi seperti itu ya itu terjadi karena code programnya belum sempat diunlock
		dan masih menunggu

		karena kita menggunakan time.Spleep, maka goroutinenya akan berhenti karena code
		program kita sudah selesai

		jadi seperti itu ya proses deadlock
		mudah mudahan kalian paham

	*/

}
