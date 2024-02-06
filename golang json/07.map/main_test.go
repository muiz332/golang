/*


- map


kadang ada kasus dimana kita menerima data json secara dynamic
artinya bisa berubah ubah type data ataupun keynya

nah kalo kita menggunakan struct itu akan susah, karena
structur struct itu sudah pasti

nah untuk mengatasi hal tersebut kita bisa gunakan map
dimana properti akan menjadi key didalam map

dan valuenya didalam map kita bisa kasih interface kosong
kita coba


*/

package json_map

import (
	"encoding/json"
	"fmt"
	"testing"
)


func TestJsonDynamic(t *testing.T){
	jsonByte := []byte(`{"id":1,"name":"rinso","price":3000}`)

	data := map[string]any{}
	err := json.Unmarshal(jsonByte,&data)
	if err != nil{
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data["id"])
	fmt.Println(data["name"])
	fmt.Println(data["price"])
	fmt.Println(data["rins"])

	/*
	
	akan tetapi kalian harus konversi type datanya secara manual ya
	karena dia interface kosong

	dan kalian juga bisa mengkonverti type data map menjadi json
	
	*/
}