/*
- json object

nah biasanya ketika kita membuat json isi datanya kalo tidak object adalah array
nah baru didalamnya kita bisa memberikan type data primitif

jadi dijson object digolang direpresentasikan dengan type data struct
dimana tiap attribute atau properti didalam json itu ada properti didalam
struct

oke langsung saja kita coba
*/
package json_object

import (
	"encoding/json"
	"fmt"
	"testing"
)



type CUstomer struct{
	FirstName string
	LastName string
}

func convertToJson(data any){
	bytes, err := json.Marshal(data)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(bytes))
}


func TestJsonObject(t *testing.T){

	convertToJson(CUstomer{
		FirstName: "muiz",
		LastName: "kuy",
	})

	/*
	
	kalo kita jalankan hasilnya seperti ini
	{"FirstName":"muiz","LastName":"kuy"}
	
	*/
}


