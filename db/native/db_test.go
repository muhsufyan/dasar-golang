package native

// kita hanya perlu init (constructor) dr driver go-sql-driver jd gunakan _ saja
import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnectDb(t *testing.T) {
	// buat koneksi ke db
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/migrasi_test")
	if err != nil {
		panic("koneksi db error")
	}
	// tutup koneksi
	defer db.Close()
	// gunakan DB u/ konek ke db
}

// saat menggunakan open sql.DB kita menggunakan db pool yg berisi kumpulan koneksi ke db untuk melakukan management koneksi scra otomatis
// karena setiap db punya maks koneksi maka kita perlu membatasi koneksi (misal ada 100 request ke db maka perlu 100 koneksi ke db ini tdk berbahaya)
