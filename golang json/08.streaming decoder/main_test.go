/*

- streaming decoder dan encoder

sebelumnya kita belajar package json dengan melakukan konversi data json
yang sudah dalam bentuk variable dan data string atau byte

pada kenyataannya, kadang data json itu berasal dari input berupa io.Reader (file,network, request body)

kita bisa saja membaca data tersebut menggunakan io.ReadAll() dan mengkonversinya
menggunakan string baru lakukan konversi ke json

akan tetapi hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur
membaca data dari stream

dan untuk membaca data dari stream kita bisa gunakan function
json.NewDecoder(reader)

oke langsung saja kita coba


*/

package streaming_decoder_encoder

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)



type Person struct{
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
}



func TestStreamDecoder(t *testing.T){
	

	reader, err := os.Open("data.json") // membaca file menggunakan reader
	if err != nil{
		panic(err)
	}

	decoder := json.NewDecoder(reader) // decoder
	muiz := Person{}

	err = decoder.Decode(&muiz)
	if err != nil{
		panic(err)
	}

	fmt.Println(muiz)

	/*
	
	jadi kalian tidak perlu membaca isi filenya dulu menggunakan io.ReadAll lalu kemudian 
	diubah ke struct menggunakan json.UnMarshal

	jadi bisa secara langsung seperti itu 

	dan ini lebih efisien karena tidak perlu meload data file kememory dulu
	bisa langsung kita baca menggunakan json.NewDecoder

	jadi ketika nanti kalian berinteraksi dengan file json
	maka gunakanlah stream decoder ini 



	
	*/
}


type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
}

func TestStreamEncoder(t *testing.T){

	/*
	
	nah sekarang kita coba ubah dari type data digolang menjadi file 
	json

	kita coba

	*/

	user := User{
		Username: "muiz",
		Email: "muiz@gmail.com",
	}

	writer, err := os.Create("data1.json") // buat filenya dulu
	if err != nil{
		panic(err)
	}

	encoder := json.NewEncoder(writer) // kita buat encodenya

	err = encoder.Encode(user) // kita masukkan datanya
	if err != nil{
		panic(err)
	}



}