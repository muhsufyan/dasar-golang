package gin_dasar

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	/* URL PARAMETER
	This handler will match /halo/john but will not match neither /user/ or /user
	*/
	router.GET("/halo/:name", func(c *gin.Context) {
		// get data from url param :name
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send /user/john/update /user/john/delete etc.
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		// get data from url param :name
		name := c.Param("name")
		// get everyting data from *action
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	/* QUERY STRING => ?data=x&data2=y
	Query string parameters are parsed using the existing underlying request object.
	The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	*/
	router.GET("/welcome", func(c *gin.Context) {
		// set default query string with the variable/parameter is firstname & the value is Guest
		firstname := c.DefaultQuery("firstname", "Guest")
		// get data from query string with parameter is lastname, ex /welcome?firstname=Jane&lastname=Doe so data lastname is Doe
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello, query string for parameter firstname is %s & lastname is %s", firstname, lastname)
	})
	/* MULTIPART/URLENCODED FORM
	Di postman pilih POST, pd bagian body pilih form-data ataupun x-www-urlencoded-form lalu masukkan parameter value-nya berupa
	message, status, dan nick isi dg data input
	*/
	router.POST("/form_post", func(c *gin.Context) {
		// catch data from form-data ataupun x-www-urlencoded-form with key message
		message := c.PostForm("message")
		// catch data from form-data ataupun x-www-urlencoded-form with key nick but if no filling the data(value) will set default is anonymous
		nick := c.DefaultPostForm("nick", "anonymous")
		// key status will filled hardcode with value posted
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	/*GABUNGAN QUERY STRING (PARAMETER URL) DG POST FORM (URLENCODED FORM)

	URL => POST /post?id=1234&page=1 HTTP/1.1
	HEADER => Content-Type: application/x-www-form-urlencoded

	form-data ataupun x-www-urlencoded-form => name=manu&message=this_is_great
	*/
	router.POST("/post", func(c *gin.Context) {
		// catch data from query string with param is id
		id := c.Query("id")
		// catch data from query string with param is page and if not filling/set will contain default value yaitu 0
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
