package pointerDiFunc

import "fmt"

type Person struct {
	Nama, Jk string
}

// hanya mengubah data duplikatnya sedangkan data pd alamat memory dr variabel tdk berubah
func dataNotChange(data Person) {
	data.Nama = "ujang"
}

// slbmnya pointer defaultnya pass by value sehingga data akan aman (tidak mudah untuk diubah)
// LALU BAGAIMANA JIKA INGIN MEMBUAT FUNC YG BISA MENGUBAH DATA PADA SUATU POINTER DIMANA POINTER INI DIJDKAN SBG PARAMETER
// CARANYA DG MENGGUNAKAN TANDA * PADA PARAMETER NYA
func dataChange(data *Person) {
	data.Nama = "ujang"
}
func dataChangeJk(data *Person) {
	data.Jk = "perempuan"
}
func main() {
	data := Person{"firman", "laki"}
	dataNotChange(data)
	fmt.Println(data) //nama tidak berubah jd ujang
	// agar berubah juga perlu menggunakan tanda &
	dataChange(&data)
	fmt.Println(data)
	// atau bisa juga
	var jk *Person = &data
	dataChangeJk(jk)
	fmt.Println(jk)
}

// NOTE JIKA PARAMETERNYA BANYAK MAKA GUNAKAN POINTER(*) DIBANDINGKAN PASS BY VALUE, AGAR MEMORY TIDAK BOROS. INI AKAN DIBAHAS PD POINTER METHOD
