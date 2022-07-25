// www.youtube.com/watch?v=t9QJPE5vwhs
// untuk unit test gunakan parm *testing.T
// tdk boleh return value

/*fungsi untuk menggagalkan unit test
t.Fail() & t.FailNow()
t.Error(argument) => menampilkan log error lalu jlnkan Fail()
t.Fatal(argument) => mirip dg t.Error() bedanya akan menjlnkan FailNow()
lbh disarankan Fail() & Error()
*/
package utils

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelamatWaktuFail(t *testing.T) {
	result := SelamatWaktu("siang")
	// hasilnya unit test gagal (tdk sesuai ekspektasi)
	if result != "selama siang" {
		t.Fail()
	}
	fmt.Println("program ini tetap di eksekusi")
}

func TestSelamatWaktuFailNow(t *testing.T) {
	result := SelamatWaktu("siang")
	// hasilnya unit test gagal (tdk sesuai ekspektasi)
	if result != "selama siang" {
		t.FailNow()
	}
	fmt.Println("program ini tidak di eksekusi")
}

func TestSelamatWaktuError(t *testing.T) {
	result := SelamatWaktu("siang")
	// hasilnya unit test gagal (tdk sesuai ekspektasi)
	if result != "selama siang" {
		t.Error("Result must selamat siang")
	}
	fmt.Println("program ini tetap di eksekusi")
}

//PACKAGE testify UNTUK ASSERTION(menginput data yg banyak untuk testing)
// go get github.com/stretchr/testify
func TestSelamatWaktuAssertion(t *testing.T) {
	result := SelamatWaktu("siang")
	// param 1 adlh objek test, param 2 expected value, param 3 input data, param 4 keterangan
	assert.Equal(t, "selamat siang", result, "equal")
	fmt.Println("program ini tetap di eksekusi")
}

func TestSelamatWaktuAssertionFail(t *testing.T) {
	result := SelamatWaktu("siang")
	// param 1 adlh objek test, param 2 expected value, param 3 input data, param 4 keterangan
	assert.Equal(t, "selama siang", result, "not equal")
	fmt.Println("program ini tetap di eksekusi")
}

// di testify assert jika gagal akan memanggil Fail sedangkan require akan memanggil FailNow
func TestSelamatWaktuRequire(t *testing.T) {
	result := SelamatWaktu("siang")
	// param 1 adlh objek test, param 2 expected value, param 3 input data, param 4 keterangan
	require.Equal(t, "selamat siang", result, "equal")
	fmt.Println("program ini tetap di eksekusi")
}

func TestSelamatWaktuRequireFail(t *testing.T) {
	result := SelamatWaktu("siang")
	// param 1 adlh objek test, param 2 expected value, param 3 input data, param 4 keterangan
	require.Equal(t, "selama siang", result, "not equal")
	fmt.Println("program ini tetap di eksekusi")
}

/*SKIP TEST (MEMBATALKAN TEST, DITOLAK)
misal unit test hanya untuk di window maka ketika di linux akan di skip dg fungsi skip()
*/
func TestSkip(t *testing.T) {
	// os nya mac os
	if runtime.GOOS == "darwin" {
		t.Skip("unit test cant run on mac os")
	}
	result := SelamatWaktu("siang")
	// param 1 adlh objek test, param 2 expected value, param 3 input data, param 4 keterangan
	require.Equal(t, "selamat siang", result, "equal")
}

/*Before dan After Test
gunakan *testing.M yg akan dirun disemua unit test dlm 1 package yg sama
*/
func TestMain(m *testing.M) {
	// BEFORE
	// misalnya koneksi ke db, buat dulu token jwt. kasus ini buat dulu.
	fmt.Println("KONEKSI DB (before unit test")
	//UNIT TEST
	m.Run()
	// AFTER
	fmt.Println("CLOSE DB (after unit test)")
}

/* Sub Test
func unit test dalam func unit test.
gunakan t.Run()
*/
func TestSubTest(t *testing.T) {
	// nama sub unit testnya adlh pagi, lalu func testnya apa. untuk run unit test pagi ini saja (1 sub test) perintahnya go test -run TestNamaFunc/NamaSubTest
	// contoh go test -run TestSubTest/pagi
	// untuk semua sub unit test
	t.Run("pagi", func(t *testing.T) {
		result := SelamatWaktu("pagi")
		require.Equal(t, "selamat pagi", result, "Result must be 'selamat pagi'")
	})
	t.Run("siang", func(t *testing.T) {
		result := SelamatWaktu("siang")
		require.Equal(t, "selamat siang", result, "Result must be 'selamat siang'")
	})
	t.Run("malam", func(t *testing.T) {
		result := SelamatWaktu("malam")
		require.Equal(t, "selamat malam", result, "Result must be 'selamat malam'")
	})
}

/* Table Test
sub test biasanya digunakan untuk table test (data berupa slice yg berisi ekspektasi dan data input). table test akan diiterasi oleh sub test
*/
func TestTableSelamatWaktu(t *testing.T) {
	// buat struct slice/array scra langsung
	tests := []struct {
		name_test string
		// request/input
		input    string
		expected string
		adverb   string
	}{
		// data slice nya (tabel slice)
		{
			name_test: "siang",
			input:     "siang",
			expected:  "selamat siang",
			adverb:    "unit test sukses",
		},
		// data slice nya (tabel slice)
		{
			name_test: "malam",
			input:     "malam",
			expected:  "selamat malam",
			adverb:    "unit test sukses",
		},
	}
	// looping untuk melakukan test
	for _, value_test := range tests {
		t.Run(value_test.name_test, func(t *testing.T) {
			result := SelamatWaktu(value_test.input)
			require.Equal(t, value_test.expected, result, value_test.adverb)
		})
	}
}

/* MOCK
adlh objek yg kita siapkan. digunakan untuk membuat objek yg sulit untuk dilakukakun test
contohnya call API third party, response-nya tidak dpt kita tentukan ekspektasinya. konek ke db, query ke db
gunakan assertion, kode program harus benar" baik
CONTOH KASUS KALI INI BUAT QUERY KE MOCK DB, KITA BUAT SERVICE LAYER as bussiness logic & repository as bridge to db
IDEALNYA KITA BUAT INTERFACE UNTUK SERVICE & REPOSITORY LAYER
(lbh dlm https://threedots.tech/post/repository-pattern-in-go/
https://medium.com/golangid/mencoba-golang-clean-architecture-c2462f355f41
)
PROJEK UNTUK MOCK IMPLEMENTASINYA ADA DI FOLDER/PACKAGE entity, repository & service
*/

/*
testing.B => untuk menghitung benchmark (speed code program kita). fungsinya sama dg testing.T
nama func untuk uji benchmark hrs diawali BenchmarkNamaFunc & ada param *testing.B, tdk boleh mengembalikan return value. dpt digabung dg unit test
untuk run semua benchmark di 1 package (termasuk unit test) go test -v -bench=.
run benchmark saja (tdk ada unit test) go test -v -run=TestTidak -bench=.
run benchmark tertentu saja (tdk ada unit test) go test -v -run=TestTidakAda -bench=BenchmarkSelamatWaktuMalam
*/
func BenchmarkSelamatWaktuMalam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelamatWaktu("malam")
	}

}
func BenchmarkSelamatWaktuSiang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelamatWaktu("siang")
	}

}

/*
Sub benchmark
*/
func BenchmarkSub(b *testing.B) {
	b.Run("pagi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelamatWaktu("pagi")
		}
	})
	b.Run("sore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SelamatWaktu("sore")
		}
	})
}

// TABLE BENCHMARK
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "siang",
			request: "siang",
		},
		{
			name:    "malam",
			request: "malam",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SelamatWaktu(benchmark.request)
			}
		})
	}
}
