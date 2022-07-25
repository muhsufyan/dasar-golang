package cabang_loop

import "fmt"

func main_func() {
	// SHORT IF
	fmt.Println("====SHORT IF====")
	if uji := "uji"; uji != "uji" {
		fmt.Println("ini bukan uji")
	} else {
		fmt.Println("ini uji")
	}
	fmt.Println("========")
	// SWITCH
	fmt.Println("====SWITCH====")
	// switch nama_variabel_yg_dicek
	// case kondisi_yg_akan_dieksekusi
	hari := "minggu"
	switch hari {
	case "minggu":
		fmt.Println("libur")
	case "senin":
		fmt.Println("pertama kerja")
	default:
		fmt.Println("kerja")
	}
	fmt.Println("========")
	// SHORT SWITCH
	fmt.Println("====SHORT SWITCH====")
	switch hari2 := "minggu"; hari2 == "minggu" {
	case true:
		fmt.Println("libur")
	case false:
		fmt.Println("kerja")
	}
	fmt.Println("========")
	// SHORT SWITCH
	fmt.Println("====SHORT SWITCH 2====")
	switch hari2 := "minggu"; hari2 {
	case "minggu":
		fmt.Println("libur")
	default:
		fmt.Println("kerja")
	}
	fmt.Println("========")
	// SWITCH TANPA VARIABEL
	fmt.Println("====SWITCH TANPA VARIABE====")
	hari3 := "minggu"
	switch {
	case hari3 == "minggu":
		fmt.Println("libur")
	default:
		fmt.Println("kerja")
	}
	fmt.Println("========")
	// FOR LOOP
	fmt.Println("====FOR LOOP====")
	prima := []int{1, 3, 5, 7, 9}
	for index, angka := range prima {
		fmt.Println("index ", index, " = ", angka)
	}
	person := map[string]string{
		"nama":   "saya",
		"alamat": "indo",
	}
	for key, value := range person {
		fmt.Println("key = ", key, " : value ", value)
	}
	fmt.Println("========")
}
