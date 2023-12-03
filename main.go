package main

import (
	"github.com/gin-gonic/gin"
	"/home/rohith/gogingorm/db"
	"log"
	"net/http"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "albums")
}

func main() {
	db, err := db.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8001")
}
