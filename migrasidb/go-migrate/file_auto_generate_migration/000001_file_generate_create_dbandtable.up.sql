CREATE TABLE category(
    id_category int AUTO_INCREMENT PRIMARY KEY, 
    name varchar(200)
);
CREATE TABLE barang(
    id_barang int AUTO_INCREMENT PRIMARY KEY,
    nama varchar(200),
    jumlah_stok int
);