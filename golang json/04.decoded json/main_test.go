/*

- decoded json

sebelumnya kita sudah tahu bagaimana melakukan encoded json
nah sekarang kita coba decoded json

jadi data json kita ubah ke structnya
jadi kita tinggal gunakan function UnMarshal ya

oke langsung saja kita coba

*/

package decoded_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct{
	FirstName string
	LastName string
}


func TestDecoded(t *testing.T){
	dataJson := `{
		"FirstName" : "muiz",
		"LastName" : "kuy"
	}`

	dataJsonByte := []byte(dataJson)

	customer := Customer{}

	err := json.Unmarshal(dataJsonByte,&customer)
	if err != nil{
		panic(err)
	}

	fmt.Println(customer.FirstName) // muiz
	fmt.Println(customer.LastName) // kuy
}