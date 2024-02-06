/*

- endoced json


golang telah menyediakan sebuah function untuk melakukan conversi data ke json
dengan menggunakan function Mershal(any)

karena parameternya adalah interface kosong, maka kita bisa memasukkan type data apapun
kedalam function tersebut

kita coba




*/

package golang_encoded_json

import (
	"encoding/json"
	"fmt"
	"testing"
)


func covertToJson(data any){

	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	/*
	
	jadi karena json itu string maka yang dikembalikan ada bytes
	jadi kita harus konversi ke type data string untuk membacanya
	
	*/
}

type Person struct{
	Nama string
	Alamat string
}

func TestJSONEncoded(t *testing.T){


	covertToJson("muiz")
	covertToJson([]string{"hasan","husain"})
	covertToJson(Person{
		Nama: "muiz",
		Alamat: "banyuwangi",
	})

	/*
	
	maka kalo kita jalankan hasilnya akan seperti ini 
	"muiz"
	["hasan","husain"]
	{"Nama":"muiz","Alamat":"banyuwangi"}
		
	*/


}