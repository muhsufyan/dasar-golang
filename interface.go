package interfac

import (
	"errors"
	"fmt"
)

// interface di go adlh tipe data abstract. used as contract (kontrak). kontrak adlh implementasi dari method dlm interface
// interface can used as param func.
type IHasName interface {
	GetName() string
}

func halo(hasName IHasName) {
	fmt.Println("hi ", hasName.GetName())
}

// jika struct a mirip dg kontrak interface a maka otomatis struct a adlh interface a. ini membuat bingung jika menggunakan python karena berbeda pd python disebutkan
// contoh pd struct method Person dpt mengakses / punya method GetName, GetName sendiri juga mrpkn interface dr IHasName maka struct Person sdh mengikuti interface IHasName
type Person struct {
	Name string
}

// jika struct method ini dihilangkan maka akan error, karena Person tidak punya kontrak GetName
func (p Person) GetName() string {
	return p.Name
}

// kita buat struct/objek yg lain
type Pet struct {
	Name string
}

func (pet Pet) GetName() string {
	return pet.Name
}

// INTERFACE KOSONG
// pd typescript kita mengenal tipe data Any yg dpt menerima semua tipe data apapun. konsep interface kosong itu = super class object di java, any di kotlin.
// deklasarinya interface{}
func interface_kosong() interface{} {
	// return "ini string"
	// return 1
	return false
}

// ERROR
// error pd go dihandle oleh interface kosong Error. untuk membuat error baru tinggal errors.New("nama errornya")
func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

// ASSERTIONS
// change data type to data type we want. seringnya digunakan ketika berhub dg interface{}
func assertion() interface{} {
	return "OK"
	// return 1
}
func main() {
	person := Person{Name: "slamet"}
	halo(person)
	kucing := Pet{Name: "pus"}
	halo(kucing)
	// cara deklarasi interface kosong
	kosong := interface_kosong()
	fmt.Println("ini interface kosong datanya mengikuti (dinamis). ", kosong)
	// cara deklarasi interface kosong
	var kosong2 interface{} = interface_kosong()
	fmt.Println(kosong2)
	// ERROR INTERFACE
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
	// UBAH INTERFACE{} JD STRING
	result := assertion()
	res_string := result.(string)
	fmt.Println(res_string)
	// ASSERTION DG SWITCH UNTUK MENCEGAH PANIC KARENA TIPE DATANYA BERBEDA
	switch tipe := result.(type) {
	case string:
		fmt.Println("String ", tipe)
	case int:
		fmt.Println("int ", tipe)
	default:
		fmt.Println("unk")
	}
}
