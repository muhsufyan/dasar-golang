package struk

import (
	"fmt"
)

type Person struct {
	Nama, Alamat string
	Umur         int
}

//STRUCT METHOD
// seolah" struct punya function (class memiliki method/func)
//dg kata lain func dpt mengakses struct
func (p Person) sayHi(nama string) {
	fmt.Println("Hello", nama, "My Name is", p.Nama)
}

func main() {
	var deklarasi_biasa Person
	deklarasi_biasa.Nama = "wawan"
	deklarasi_biasa.Alamat = "ina"
	deklarasi_biasa.Umur = 30
	fmt.Println(deklarasi_biasa)
	fmt.Println(deklarasi_biasa.Umur)
	literal := Person{
		Nama:   "wewen",
		Alamat: "ina",
		Umur:   31,
	}
	fmt.Println(literal)
	literal2 := Person{"qa", "ina", 76}
	fmt.Println(literal2)
	literal.sayHi("Joko")

}
