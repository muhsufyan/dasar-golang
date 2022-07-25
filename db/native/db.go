package native

import (
	"database/sql"
	"time"
)

func GetConnectionDBPool() *sql.DB {
	// buat koneksi ke db
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/migrasi_test")
	if err != nil {
		panic(err)
	}
	// close koneksi dilakukan saat test, karena ini menggunakan db pooling
	// set db pool
	db.SetMaxIdleConns(10)                  //min koneksi
	db.SetMaxOpenConns(100)                 //max koneksi
	db.SetConnMaxIdleTime(5 * time.Minute)  //dlm 5 menit jika koneksi tdk digunakan maka koneksi akan diputus (closed)
	db.SetConnMaxLifetime(60 * time.Minute) //stlh 60 menit koneksi akan renew
	return db
}
