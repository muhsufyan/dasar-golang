package fungsi

import "fmt"

func fungsiBiasa() (string, string) {
	return "halo", "pagi"
}
func namedReturnValue() (angka int, kata string) {
	angka = 2
	kata = "dua"
	return
}

// variadic func adlh parameter terakhir dr suatu fungsi dpt menerima banyak data input(sprti array). ini sprti konsep ... di js. tandanya adlh ...tipe_data dan disimpan diakhir
func variadicFunc(angka ...int) int {
	total := 0
	// angka ketika masuk ke fungsi akan menjd slice
	for _, value := range angka {
		total += value
	}
	return total
}

// func di go bisa dijdkan as tipe data & bisa disimpan dlm variabel, dibuat independen (tanpa didlm class)
func func_value(waktu string) string {
	return "selamat " + waktu
}

// func as parameter for another func. func ini akan menerima param func lainnya yg yaitu func(string) string. dg memakai func as param kode jd lbh clean
func func_as_param(nama string, filter_nama func(string) string) {
	filter := filter_nama(nama)
	fmt.Println(nama, " ", filter)
}
func filtering(kata string) string {
	if kata == "kasar" {
		return "sensor"
	} else {
		return kata
	}
}

// jika func as param memiliki banyak input (param, misal pd func_as_param diatas filter_nama memiliki 5 inputan) maka gunakan func type declaration yg
// berfungsi as alias from a func
type Filter func(string) string

func func_as_param2(nama string, filter_nama Filter) {
	filter := filter_nama(nama)
	fmt.Println(nama, " ", filter)
}

// REKURSIF
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	hi, pagi := fungsiBiasa()
	fmt.Println(hi, " ", pagi)
	halo, _ := fungsiBiasa()
	fmt.Println(halo)
	fmt.Println(namedReturnValue())
	kosong := variadicFunc()
	fmt.Println(kosong)
	ada_isi := variadicFunc(1, 2, 3)
	fmt.Println(ada_isi)
	slice := []int{2, 3, 4, 5}
	// pecah jd variabel argumen dg cara tambahkan ... diakhir nama variabel slice
	slice_as_variadic := variadicFunc(slice...)
	fmt.Println(slice_as_variadic)
	// 1. func as tipe data & bisa disimpan dlm variabel. kita panggil siang as func
	siang := func_value
	fmt.Println(siang("Siang bos"))
	func_as_param("bolang", filtering)
	func_as_param("kasar", filtering)
	func_as_param2("tes", filtering)
	func_as_param2("kasar", filtering)
	// ANONYMOUS FUNC
	anonymous := func(kata string) bool {
		return kata == "kotor"
	}
	fmt.Println(anonymous("bersih"))
	fmt.Println(anonymous("kotor"))
	// REKURSIF
	loop := factorialLoop(3)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialRecursive(3)
	fmt.Println(recursive)
}
