/*

- sync rwmutex

jadi rw itu read dan write mutual exclusion

kadang ada kasus dimana kita melakukan locking tidak hanya pada
proses mengubah data tetapi juga membaca data

kita sebenarnya bisa menggunakan mutex saja,namun masalahnya nanti akan rebutan dalam proses
membaca dan mengubah

digolang disediakan struct Rwmutex untuk menangani hal ini, dimana mutex jenis ini memiliki
2 lock, lock untuk read dan lock untuk write



*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // lock untuk melakukan perubahan
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock() // unclokc untuk melakukan perubahan
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock() // lock untuk melakukan read
	var bal int = account.Balance
	account.RWMutex.RUnlock() // unlock untuk melakukan read
	return bal
}

func main() {

	var account BankAccount = BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(account.GetBalance())

	/*

		jadi kalian harus melakukan locking terhadap method yang melakukan perubahan data
		dan method yang melakukan membaca data

		agar kemungkinan untuk terjadi race condition itu tidak ada
		jadi data kita akan aman ya

		jadi seperti itulah penggunaan dari read and write mutation exclusion

		ingat bukan berarti kalian membuat struct itu harus menggunakan rwmutex ya
		ini kalo kasusnya kalian membuat struct dan ternyata dia bakal diakses ke beberapa
		goroutine

		tapi kalo tidak kalian tidak perlu menggunakan rwmutex ini ya
		oke jadi seperti itu mudah mudahan kalian paham

	*/
}
