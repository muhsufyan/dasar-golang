sblm menggunakan golang-migrate kita harus install dulu scoop untuk windows, caranya dilink ini https://scoop.sh/


https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

https://www.youtube.com/watch?v=i8sJPbL6K0M

LANGKAH
A. buat file migrate dulu scra auto generate dg perintah
    migrate create -ext sql -dir file_auto_generate_migration -seq file_generate_create_dbandtable
pd perintah tsb terdpt beberapa param yaitu
1. create = membuat file migrate baru
2. -ext sql= ditulis dan dieksekusi dlm sql & berformat .sql
3. -dir file_auto_generate_migration = tempat (folder) dimana file hasil generate akan di simpan, kita simpan di folder file_auto_generate_migration 
4. -seq file_generate_create_dbandtable = nama file generate-nya, kasus ini nama file nya file_generate_create_dbandtable

Perintah membuat file generate untuk migrasi di documentationnya
    migrate create -ext sql -dir db/migrations -seq create_users_table

untuk membuat file migrasi mongodb blm diketahui

B. QUERY MEMBUAT TABEL, DATABASE, SCHEMA
Hsl generate akan ada 2 file yaitu up (u/ make tabel, schema) & down (rollback [menghapus tabel, schema]). untuk membuat db baru dg golang-migrate ini tidak bisa jd hrs buat dulu db nya di server dg query 
CREATE DATABASE IF NOT EXISTS migrasi_test;

pd up scriptnya

CREATE TABLE category(
    id_category int AUTO_INCREMENT PRIMARY KEY, 
    name varchar(200)
);
CREATE TABLE barang(
    id_barang int AUTO_INCREMENT PRIMARY KEY,
    nama varchar(200),
    jumlah_stok int
);

untuk down scriptnya

DROP TABLE IF EXISTS category, barang;

untuk run query migrasinya (up) di documentasi
    migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2
2 artinya run 2 file migrasi pertama (run the first two migrations)

sedangkan pd kasus kita karena memakai windows jd
    migrate -path file_auto_generate_migration -database "mysql://root:@tcp(localhost:3306)/migrasi_test" up

menghapus tabel pd query up 
    migrate -path file_auto_generate_migration -database "mysql://root:@tcp(localhost:3306)/migrasi_test" down


sumber
https://www.youtube.com/watch?v=i8sJPbL6K0M


DB NYA DARI DOCKER
https://www.youtube.com/watch?v=0CYkrGIJkpw
https://www.youtube.com/watch?v=6_CH8Gx414A

