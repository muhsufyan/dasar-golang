package native

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnectionDBPool()
	defer db.Close()
	// disarankan memakai context agar dpt membatalkan query
	ctx := context.Background()

	script := "INSERT INTO category(name) VALUES('Joko')"
	// kita tdk perlu resource sehingga gunakan _ & juga ExecContext tdk mengembalikan resource (data, tdk dpt mengambil data) sehingga ExecContext hanya digunakan untuk CUD
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new category")
}

// gunakan func (DB)QueryContext(context, sql, params) untuk query select / R (read) kemudian tutup datanya dg Close()
func TestQueryRead(t *testing.T) {
	db := GetConnectionDBPool()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT * FROM category"
	data, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer data.Close()
}

// iterasi datanya

func TestQueryReadIterateData(t *testing.T) {
	db := GetConnectionDBPool()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT * FROM category"
	data, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	/* hasil query diatas adlh struct sql.Rows dimana Rows itu u/ melakukan iterasi pd hasil data query yg didpt
	func (Rows)Next()(boolean) akan melakukan iterasi data, if false mean sdh tdk ada data yg dpt diiterasi lagi(all data already get)
	for read setiap data use (Row)Scan(columns...)
	stlh selesai tutup Row dg (Row)Close()

	*/
	//  selama datanya ada maka lakukan iterasi
	for data.Next() {
		// menampung data id_category dan name
		var id_category int32
		var name string
		// read setiap data (get all data), gunakan pointer agar data dpt ditangkap
		err = data.Scan(&id_category, &name)
		if err != nil {
			panic(err)
		}
		// print data
		fmt.Println("Id:", id_category)
		fmt.Println("Name:", name)
	}
}
