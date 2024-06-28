package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := gin.Default()

	router.GET("/stream", func(c *gin.Context) {
		tr := c.Query("track")
		if tr == "" {
			c.Writer.WriteHeader(http.StatusNotFound)
			panic("not found")
		}

		fstream, err := os.Open(tr)
		if err != nil {
			panic(err)
		}
		stream(fstream, c)
	})

	router.Run("127.0.0.1:8080")
}
