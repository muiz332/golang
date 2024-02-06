/*

- json tag


secara default attribute yang terdapat distruct dan json akan dimapping
dengan attribute yang sama atau case sensitive

kadang ada style yang berbeda antara distruct dan dijson
kalao distruct itu menggunakan awalan huruf besar sedangkan dijson itu
menggunakan awalan huruf kecil

nah kita bisa menambahkan tag reflection dengan nama json lalu diikuti dengan attribute
yang kita ingin kan ketika dikonversi dari json atau ke json

jadi simplenya json tag ini sebagai penanda
oke langsung saja kita coba




*/

package json_tag

import (
	"encoding/json"
	"fmt"
	"testing"
)


type Product struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Price int64 `json:"price"`
}


func TestJsonTagFrom(t *testing.T) {
	
	/*
	
	jadi nanti hasil dari converti dari struct ke json akan
	mengikuti tag json yang ada setelah type datanya yang ada
	didalam kutip 2 tersebut
	
	*/


	rinso := Product{
		Id: 1,
		Name: "rinso",
		Price: 3000,
	}

	bytes, err := json.Marshal(rinso)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(bytes))

	/*
	
	maka kalo saya jalankan hasilnya seperti ini 
	{"id":1,"name":"rinso","price":3000}
	
	*/
}
func TestJsonTagTo(t *testing.T) {

	jsonBytes := []byte(`{"id":1,"name":"rinso","price":3000}`)

	rinso := Product{}
	err := json.Unmarshal(jsonBytes,&rinso)
	if err != nil{
		panic(err)
	}

	fmt.Println(rinso)

	fmt.Println(rinso.Id)
	fmt.Println(rinso.Name)
	fmt.Println(rinso.Price)

	/*
	
	jadi itulah fungsi json tag yaitu sebagai penanda sebuah attribute akan dimasukkan 
	ke nama attribute yang mana
	
	*/
}