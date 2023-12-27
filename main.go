package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/avvsrohith/hallbooking_backend/db"
	"github.com/avvsrohith/hallbooking_server_go/models"
	"github.com/avvsrohith/hallbooking_server_go/routes"
	"log"
	// "net/http"
)

func main() {
	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()
	router.POST("/hall", routes.PostHall)
	router.GET("/hall", routes.GetHall)
	router.POST("/event", routes.PostEvent)
	router.GET("/event", routes.GetEvent)

	router.Run("localhost:8001")
}
