package key

import "fmt"

func func_defer() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer func_defer()
	fmt.Println("Run application")
	// ini akan error tp func_defer akan tetap dijlnkan dg memakai defer
	result := 10 / value
	fmt.Println("Result", result)
}

// PANIC
// dg panic maka code program akan berhenti dieksekusi. panic is for stop run code
// RECOVER
// dg recover maka kode program akan terus dijlnkan dan kode yang error (dari panic) akan ditampung oleh recover ini.
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runAppErr(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
func main() {
	// CLOSURE (BLOCK CODE)
	nama := "nama 1"
	fmt.Println(nama)
	anony := func() {
		nama := "dalam dalam block anonymous func"
		fmt.Println(nama)
	}
	anony2 := func() string {
		nama := "block anony2"
		return nama
	}
	anony()
	fmt.Println(anony2())
	fmt.Println(nama)
	// DEFER
	// dieksekusi stlh suatu func selesai dijlnkan meskipun func tsb terdpt error. urutan run func kemudian defer. contoh loggin akan dirun stlh semua kode dlm runApp selesai running
	loggin := func() {
		fmt.Println("selesai")
	}
	runApp := func() {
		defer loggin()
		fmt.Println("run")
	}
	// runApp()
	runApp()
	// runApplication(0)
	runAppErr(false)
}
