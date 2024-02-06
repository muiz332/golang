/*



- name return value

digolang kita bisa membuat sebuah variable  untuk dijadikan sebagai
return valuenya

langsung saja kita coba


*/

package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "muiz"
	lastName = "zuddin"
	return
}

func main() {

	firstName, lastName := getFullName()

	fmt.Println(firstName, lastName)

	/*

		jadi kita bisa buat variablenya secara langsung didalam kurung beserta type data
		untuk return variablenya

		kemudian kalian bisa assigment variablenya didalam functionnya
		untuk mereturn kalian hanya tulis

		return saja, karena kita sudah membuat variable didalam kurung buka dan tutup
		untuk return valuenya
	*/
}
