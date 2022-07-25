package pointer

import "fmt"

// all variable in go default is passing by value not by reference.
// yg dikirim adalh duplikasi nilainya bukan alamat memory dari variabel-nya
type Address struct {
	City, Province, Country string
}

func main() {
	// maslah pd passing by value
	data := Address{"ab", "cd", "ef"}
	data2 := data
	data2.City = "Lampung"
	fmt.Println(data)
	fmt.Println(data2)
	// dg passing by value (bkn by reference) data & data2 isinya beda pdahal data2 itu mereference ke data. ini mslh passing by value
	// POINTER
	// pointer adlh passing by reference (alamat memory). data3 akan reference ke memory variabel data. jika data3/data berubah maka data/data3 juga ikut berubah
	// tp ada tanda & ketika diprint yg menunjukkan bahwa data3 adlh pointer
	data3 := &data
	data3.City = "langit"
	fmt.Println(data)
	fmt.Println(data3)
	// penulisan pointer (pass by reference)
	var data4 *Address = &data
	fmt.Println(data)
	fmt.Println(data4)
	data4.City = "sky"
	fmt.Println(data)
	fmt.Println(data4)
	// dg melakukan ini maka variabel data tdk akan berubah tp hanya data4 saja yg berubah (mengubah data4 keseluruhan)
	data4 = &Address{"zak", "tak", "wak"}
	fmt.Println(data)
	fmt.Println(data4)
	// mengubah keseluruhan data
	var data5 *Address = &data
	*data3 = Address{"ini", "di", "mana"}
	// fmt.Println(data)
	fmt.Println(data3)
	fmt.Println(data4)
	fmt.Println(data5)
	// POINTER KOSONG
	var data6 *Address = new(Address)
	data6.Province = "pulau ikan"
	fmt.Println(data6)
}
