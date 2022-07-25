package strukturdata

import "fmt"

// RUNNING
// untuk buat file kompilasinya jlnkan perintah "go build nama_file.go" maka akan dibuat file (windows : exe, linux)
// untuk menjlnkan file yg telah dikompilasi caranya jlnkan perintah ""./nama_file_hasil_kompilasi"
// cara langsung lainnya "go run nama_file.go"

// KONSTANTA
const kons_a string = "halo"
const kons_b = 10
const (
	multi = "hi"
	var_a = "aku"
)

func main() {
	fmt.Println("====const====")
	fmt.Println(kons_a + " " + multi + " " + var_a)
	fmt.Println(kons_b)
	fmt.Println("=======")
	// CONVERT INT
	// konversi dari tipe data. hati" int32 jd int8 akan menjd masalah karena int8 cakupannya lbh kecil dari int32.
	// dimana int8 range-nya 128 sampai -127. lebih detail https://www.callicoder.com/golang-basic-types-operators-type-conversion/
	fmt.Println("====convert int====")
	var nilai32 int32 = 100000
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai8)
	fmt.Println("=======")
	// CONVERT byte to string
	// di go, byte itu = uint8
	var nama = "Hylos"
	var huruf_y byte = nama[1]
	var toStr string = string(huruf_y)
	fmt.Println("====convert string====")
	fmt.Println(nama)
	// yg dicetak adlh byte untuk karakter y
	fmt.Println(huruf_y)
	fmt.Println(toStr)
	fmt.Println("=======")
	// TYPE
	// used for make new data type from native data type, misal kita buat tipe data baru yaitu nik yang berasal dari string
	type nik string
	var ujang nik = "32111111111"
	fmt.Println("====type====")
	fmt.Println(ujang)
	fmt.Println(nik("3222222222"))
	fmt.Println("=======")
	// ARRAY
	// deklarasi dulu baru assign value
	var abc [3]string
	abc[0] = "a"
	abc[1] = "b"
	abc[2] = "c"
	// deklarasi dan assign langsung
	var angka = [3]int{
		1, 2, 3,
	}
	// array kosong
	var kosong [5]string
	fmt.Println("====ARRAY====")
	fmt.Println(abc[2])
	fmt.Println(angka)
	fmt.Println(angka[0])
	fmt.Println(len(abc))
	fmt.Println(len(angka))
	fmt.Println(len(kosong))
	// ubah data array
	angka[1] = 10
	fmt.Println(angka[1])
	fmt.Println("========")
	// SLICE
	// ada 3 atribut pada slice yaitu pointer (data awal yang ingin didpt), capacity (length total slice dikurangi pointer), length (high - low atau pointer)
	// buat array tanpa tahu lengthnya
	var alfa = [...]string{"ab", "cd", "ef", "gh", "ij"}
	fmt.Println("====SLICE====")
	fmt.Println(alfa)
	fmt.Println(alfa[1:3])
	fmt.Println(len(alfa[1:3]))
	fmt.Println(cap(alfa[1:3])) //capacity
	// ubah slice maka array akan ikut berubah
	var alfa_slice = alfa[1:3]
	// tambah data baru dg append maka otomatis menambah 1 length. data pada alfa(data asal) akan ikut berubah
	alfa_slice2 := append(alfa_slice, "data_baru")
	// membuat slice baru dg make(); param 1 = tipe datanya dlm bntk array, param 2 = length, param 3 = capacity. cara ini lbh aman
	newSlice := make([]int, 6, 10)
	// copy slice, param 1 tujuan copy, param 2 asal data yg dicopy .
	copySlice := make([]int, len(newSlice), cap(newSlice))
	// when copy lengthnya harus sama agar data tidak terpotong
	// copySlice := make([]int, 1, cap(newSlice))
	copy(copySlice, newSlice)
	/*
		deklarasi array [...]int{1,2,3}
		deklarasi slice []int{1,2,3}
	*/
	fmt.Println(alfa_slice)
	alfa_slice[1] = "diubah"
	fmt.Println(alfa_slice)
	fmt.Println(alfa_slice2)
	fmt.Println(alfa)
	fmt.Println(newSlice)
	newSlice[2] = 4
	fmt.Println(newSlice)
	fmt.Println(copySlice)
	fmt.Println("========")
	// MAP
	// jika keynya sama maka data lama akan ditimpa dg data baru,
	// format membuatnya map[tipe_data_keynya]tipe_data_valuenya
	// deklarasi & langsung assign
	person := map[string]string{
		"nama":   "saya",
		"alamat": "indo",
	}
	// tambah atribut baru
	person["data_baru"] = "ini data baru"
	// buat map dg make(map[tipe_data_key]tipe_data_value)
	film := make(map[string]string)
	film["judul"] = "durjana"
	film["sutradara"] = "alien"
	// buat map dg deklarasi
	var movie map[string]string = make(map[string]string)
	movie["title"] = "sky"
	movie["from"] = "canada"
	fmt.Println("====MAP====")
	fmt.Println(person)
	fmt.Println(person["alamat"])
	fmt.Println(person["data_baru"])
	fmt.Println(len(person))
	fmt.Println(film)
	fmt.Println(len(film))
	// menghapus => delete(map, key_yg_dihapus)
	// menghapus data dg key sutradara
	delete(film, "sutradara")
	fmt.Println(film)
	fmt.Println(len(film))
	fmt.Println(movie)
	fmt.Println("========")
}
