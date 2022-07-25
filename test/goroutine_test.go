package goroutine_test

import (
	"fmt"
	"testing"
)

// https://www.youtube.com/watch?v=dpJhTYqWOrE
// contoh penerapan concurrent adlh koneksi ke db, dibanding paralel maka concurrent lbh cpt
// concurrent berjln scra asynchronous
func concurrent() {
	fmt.Println("ini adalah fungsi yang akan dijlnkan lewat concurrent")
}
func TestCreateGoroutine(t *testing.T) {
	// untuk menjlnkan dg concurrent (goroutine) gunakan kata go sblm fungsi dipanggil
	go concurrent()
	fmt.Println("ini bukan goroutine")
	// untuk membuktikan bahwa bersifat asynchronou

}
