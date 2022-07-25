kita gunakan package (driver) go-sql-driver
go get github.com/go-sql-driver/mysql

KONEKSI KE DB
gunakan objek sql.DB dg func sql.Open(driver, dataSourceName)
driver yg dpt digunakan adlh mysql, postgresql
dataSourceName = url nya ex mysql => username:password@tcp(host:port)/nama_db

stlh beres koneksi ke db maka tutup koneksi dg func Close() karena db punya limit open connect jika penuh maka hrs kill db scra manual

DATABASE POOLING
(DB)SetMaxIdleConns(number) : jumlah koneksi min yg dibuat (baca konsep koneksi idle)
(DB)SetMaxOpenConns(number) : jumlah koneksi max yg dibuat. jika melbhi maka hrs menunggu koneksi lama selesai
(DB)SetConnMaxIdleTime(duration) : lama waktu koneksi tidak digunakan. jika koneksi tdk digunakan & telah melebihi durasi maka koneksi akan dipaksa putus (close)
(DB)SetConnMaxIdleLifetime(duration) : berapa lama koneksi dpt digunakan. misal kita set 1 jam maka koneksi akan diperbarui setiap 1 jam sekali meskipun sedang digunakan.

db pooling dilakukan agar tdk perlu melakukan open koneksi trs karena ini perlu resource yg bsr

QUERY
stlh berhsl konek ke db kita akan mengeksekusi query yg sdilakukan melalui (DB)ExecContext(context, sql, params)
dg context kita dpt cancel query  

query select tdk bisa memakai ExecContext karena tdk mengembalikan resource (data, tdk dpt mengambil data) sehingga ExecContext hanya digunakan untuk CUD

QUERY SELECT YG PERLU RETURN DATA (RESOURCE)
gunakan func (DB)QueryContext(context, sql, params)
SUMBER
https://www.youtube.com/watch?v=RM8mP6Tzolg