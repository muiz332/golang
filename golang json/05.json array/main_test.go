/*

- json array

digolang json array direpresentasikan dalam bentuk slice
oke langsung saja kita coba



*/

package json_array

import (
	"encoding/json"
	"fmt"
	"testing"
)


type Orang struct{
	Nama string
	Alamat string
	Sosmed []string
}

func TestJsonArray(t *testing.T){


	budi := Orang{
		Nama: "budi",
		Alamat: "banyuwangi",
		Sosmed: []string{"ig","fb"},
	}

	bytes, err := json.Marshal(budi)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(bytes))

	/*
	
	kalo saya jalankan maka hasilnya seperti ini 
	{"Nama":"budi","Alamat":"banyuwangi","Sosmed":["ig","fb"]}
	
	*/

}