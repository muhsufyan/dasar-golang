package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
BINDING
adlh proses membangun koneksi antara antarmuka aplikasi (application UI) dan business logic. Data binding membuat kontrol kontrol dalam form terkoneksi (terkait) dengan sumber sumber data.
binding berfungsi agar data yg diinputkan oleh user (ex via json api, dlm hal ini api disbt UI) dpt terhubung (catched) oleh backend untuk selanjutnya diproses lbh lanjut disisi server (backend)

UNTUK mem-binding body request dari user ke server(backend) agar memiliki format(type) yg sesuai dg keperluan server maka gunakan MODEL BINDING.
& UNTUK MEMVALIDASI DATA-NYA GUNAKAN validator dari https://github.com/go-playground/validator
UNTUK MELAKUKAN BINDING (CATCH DATA) MAKA KITA HARUS MEMBUAT FIELD YG SESUAI DG YG DIINPUTKAN,
MISALNYA DATA YG DIINPUT BERUPA json DG KEY age MAKA UNTUK BINDING-NYA => json:"age"

PD GIN TERDPT 2 METHOD UNTUK BINDING
1. Type-nya berupa Must bind (MustBindWith), method yg bisa digunakan Bind, BindJSON, BindQuery
	jika saat binding terdpt error maka dpt mengembalikan status kode & header ex c.AbortWithError(400, err).SetType(ErrorTypeBind)
	jika kita mencoba mengatur kode respons setelah hal diatas, itu akan menghasilkan peringatan [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422
	jika ingin memiliki control yg lbh besar maka gunakan ShouldBinding
2. Type-nya berupa Should bind (Should bind), method yg bisa digunakan ShouldBind, ShouldBindJSON, ShouldBindQuery
	jika saat binding terdpt error maka error akan dikembalikan selain itu developer dpt menghandle response selanjutnya
*/

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// BINDING FROM URL PARAM
// path paramter with name details will mapped to Details
type URI struct {
	Details string `json:"name" uri:"details"`
}

// QUERY STRING PARAM BINDING
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

// BIND QUERY STRING / POST DATA
//  curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"
type Person2 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	// BINDING FROM URL
	engine := gin.New()
	// adding path params to router
	engine.GET("/test/:details", func(context *gin.Context) {
		uri := URI{}
		// binding to URI
		if err := context.BindUri(&uri); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(uri)
		context.JSON(http.StatusAccepted, &uri)
	})
	// QUERY STRING PARAM BINDING (?key=value)
	router.Any("/testing", startPage)
	// BIND QUERY STRING / POST DATA
	router.GET("/testing2", startPage2)
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

// QUERY STRING PARAM BINDING
func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

// BIND QUERY STRING / POST DATA
func startPage2(c *gin.Context) {
	var person Person2
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}
