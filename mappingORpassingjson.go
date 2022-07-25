package mappingORpassingjson

import (
	"encoding/json"
	"fmt"
)

// disini akan mapping data dari json menggunakan struct (objek) & mapping.
// INGAT DALAM GOLANG HURUF AWAL BESAR ARTINYA PUBLIC DAN HURUF AWAL KECIL ARTINYA PRIVATE
// 1. DG STRUCT (OBJEK)
// 1) buat dulu struct, berfungsi untuk menangkap(menampung) data json. data yang ditangkap(ditangkum/disimpan) adalah nama dan umur
// nama didapat dari json dg key name, umur didapat dari json dg key umur (karena nama key untuk umur dg nama structnya sama sama umur sehingga tidak perlu ditulis)
// dan kita akan lihat bagaimana jika nama key-nya tidak ditulis & berbeda dpt dilihat pd nama dg key alamat
type User struct {
	Nama   string `json:"name"`
	Umur   int
	Alamat string
}

type Pengguna struct {
	Nama   string `json:"name"`
	Umur   int    `json:"umur"`
	Alamat string `json:"address"`
}

func main() {
	// ini adalah data dummy jsonnya
	var jsonString = `{"name":"ujang","umur":12, "address":"indonesia"}`
	var jsonData = []byte(jsonString) // ubah data json jd byte
	// instansiasi struct (objek) User, nantinya digunakan untuk menampung data json
	var data User
	// 2) deocde json agar dapat dibaca oleh golang melalui json.Unmarshal().
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("====MAPPING DG STRUCT (OBJEK)====")
	fmt.Println("nama : ", data.Nama)
	fmt.Println("umur : ", data.Umur)
	fmt.Println("alamat : ", data.Alamat)
	fmt.Println("========")
	// 2. DG VARIABEL BERTIPE map[string]interface{} ATAU BISA JUGA DISEBUT DG MAPPING MELALUI INTERFACE KOSONG
	// 1) buat dulu map dg interface kosong
	var data1 map[string]interface{}
	// 2)decode json
	json.Unmarshal(jsonData, &data1)
	fmt.Println("==== DG VARIABEL BERTIPE map[string]interface{}====")
	fmt.Println("nama : ", data1["name"])
	fmt.Println("umur : ", data1["umur"])
	fmt.Println("alamat : ", data1["address"])
	fmt.Println("========")
	// 3. DG INTERFACE KOSONG
	// dlm cara ini untuk mengakses properti(atribut) perlu melakukan casting dulu melalui map[string]interface
	// 1) buat interface kosong
	var data2 interface{}
	// 2) decode json
	json.Unmarshal(jsonData, &data2)
	// 3) casting
	var castingData = data2.(map[string]interface{})
	fmt.Println("====DG INTERFACE KOSONG====")
	fmt.Println("nama : ", castingData["name"])
	fmt.Println("umur : ", castingData["umur"])
	fmt.Println("alamat : ", castingData["address"])
	fmt.Println("========")
	// SBLMNYA HANYA 1 DATA SAJA, NOW DLM BNTK ARRAY JSON MELALUI OBJEK ARRAY/SLICE
	// datanya
	var jsonStringArray = `[
		{"name": "john wick", "umur": 27, "address":"indonesia"},
		{"name": "ethan hunt", "umur": 32, "address":"indonesia"}
		]`
	// struct dlm bntk array, untuk menampung data
	var data3 []Pengguna
	var er = json.Unmarshal([]byte(jsonStringArray), &data3)
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	fmt.Println("====ARRAY OF JSON MELALUI OBJEK ARRAY/SLICE====")
	fmt.Println("nama 1 : ", data3[0].Nama)
	fmt.Println("umur 1 : ", data3[0].Umur)
	fmt.Println("nama 2 : ", data3[1].Nama)
	fmt.Println("umur 2 : ", data3[1].Umur)
	fmt.Println("alamat 1 : ", data3[0].Alamat)
	fmt.Println("========")
	// MAPPING NESTED DATA JSON DG MAP INTERFACE KOSONG
	jsonDataRaw := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null
		}
	`

	var data_map map[string]interface{}
	err2 := json.Unmarshal([]byte(jsonDataRaw), &data_map)
	if err2 != nil {
		fmt.Printf("could not unmarshal json: %s\n", err2)
		return
	}
	fmt.Println("====MAPPING NESTED DATA JSON DG MAP INTERFACE KOSONG====")
	fmt.Printf("json map: %v\n", data_map)
	fmt.Println("object value : ", data_map["objectValue"])
	fmt.Println("========")
	// NESTED DATA KE BENTUK JSON
	data_encode := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonDataJSON, err := json.Marshal(data_encode)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	fmt.Println("====NESTED DATA KE BENTUK JSON====")
	fmt.Printf("json data: %s\n", jsonDataJSON)
	fmt.Println("========")
}

// SUMBER
/*
https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go
https://dasarpemrogramangolang.novalagung.com/A-json.html
PELAJARI INI https://blog.serverbooter.com/post/parsing-nested-json-in-go/ UNTUK ACCESS NESTED JSON
*/
