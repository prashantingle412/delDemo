package main

import (
	// "MyDemo/handlers"

	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"io/ioutil"
	// "MyDemo/handlers/first"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./f.json")
		f, _ := os.Open("./f.json")
		b, _ := ioutil.ReadAll(f)
		fmt.Print(b)
		var result map[string]interface{}
		json.Unmarshal([]byte(b), &result)
		fmt.Println("after marshal ", result)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
