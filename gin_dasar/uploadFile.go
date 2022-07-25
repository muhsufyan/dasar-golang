package gin_dasar

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	/*SINGLE FILE UPLOAD

	curl -X POST http://localhost:8080/upload \
		-F "file=@/Users/appleboy/test.zip" \
		-H "Content-Type: multipart/form-data"

	untuk mengecek ekstensi file yg diupload gunakan
	extension := filepath.Ext(file.Filename)
	sumber nya https://ramezanpour.net/post/2020/09/12/file-upload-using-go-gin
	*/
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	/*MULTIPLE FILE UPLOAD
	 Set a lower memory limit for multipart forms (default is 32 MiB)

	 curl -X POST http://localhost:8080/upload \
		-F "upload[]=@/Users/appleboy/test1.zip" \
		-F "upload[]=@/Users/appleboy/test2.zip" \
		-H "Content-Type: multipart/form-data"
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	*/
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}
