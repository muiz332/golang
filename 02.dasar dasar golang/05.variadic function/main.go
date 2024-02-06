/*

- variadic function

dimateri kali ini kita akan bahas mengenai variadic function
apa itu variadic function ?

variadic function adalah sebuah function yang memiliki kemampuan untuk
menangkap berapapun jumlah argument yang kita berikan saat functionnya dipanggil

jika kalian memiliki lebih dari satu parameter, maka variadic function harus berada
diposisi paling belakang sendiri

dengan nambahkan tanda ... serta type datanya kita bisa menangkap sisa parameter
yang ada

jadi kalo didalam javascript itu sama dengan rest parameter ya
oke langsung saja kita coba


*/

package main

import "fmt"

func sumAll(numbers ...int) int {
	var result int = 0

	for _, val := range numbers {
		result += val
	}
	return result
}

func sumAll2(from string, numbers ...int) (string, int) {
	var result int = 0

	for _, val := range numbers {
		result += val
	}
	return from, result
}

func main() {

	var result int = sumAll(10, 10, 10)
	fmt.Println(result) // 30

	/*

		jadi kita bisa mesukkan berapapun argument didalam function
		tersebut, dan kita bisa menangkapnya menggunakan variadic function

		kita coba menambahkan parameter biasa didalam functionnya

	*/
	var (
		nama  string
		hasil int
	)
	nama, hasil = sumAll2("hasan", 10, 10)
	fmt.Println(nama)  // hasan
	fmt.Println(hasil) // 20

	/*

		jadi parameter pertama akan menangkap argument pertama
		dan sisanya akan ditangkap oleh variadic tadi

		kalo misalkan kalian hanya memasukkan 1 argument saat functionnya
		dipanggil, yaitu berupa type data string

		maka variadicnya tidak akan error, karena nilainya optional
		boleh kalian berikan atau tidak kalian berikan

		dan juga variadic akan bernilai slice

		akan tetapi kalo misalkan datanya  yang kita kirimkan kedalam
		argument berupa slice, maka kita bisa pisahkan nilai menjadi
		single value dengan menambahkan ...

		kita coba

	*/

	var numbers []int = []int{1, 2, 3, 4}
	result = sumAll(numbers...)
	fmt.Println(result)

	/*

		jadi kita bisa pecah nilai yang ada didalam slice menjadi single value
		kalo didalam javascript ini sama dengan spread operator

		oke sampai disini materi kali ini mudah mudahan kalian paham

	*/

}
