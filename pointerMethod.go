package pointerMethod

import "fmt"

type Person struct {
	Nama string
}

func (p Person) PassByValue() {
	p.Nama = "tidak berubah"
	fmt.Println("ini di dalam method, valuenya : ", p.Nama)
}

// cara ini lbh hemat memory
func (p *Person) ByPointer() {
	p.Nama = "berubah dengan pointer"
}
func main() {
	data := Person{"emprit"}
	data.PassByValue()
	fmt.Println(data)
	data.ByPointer()
	fmt.Println(data)
}
