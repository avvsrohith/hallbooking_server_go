package routes

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/avvsrohith/hallbooking_server_go/models"
	"log"
    "net/http"
)

type NewHall struct {
	Name        string `json:"name"				binding:"required"`
	Description string `json:"description"		binding:"required"`
	Capacity    int32  `json:"capacity"			binding:"required"`
	Rent        int32  `json:"rent"				binding:"required"`
	Equipment   string `json:"equipment"		binding:"required"`
}

type Hall struct {
	Name        string `json:"name"`	
	Description string `json:"description"`
	Capacity    int32  `json:"capacity"`
	Rent        int32  `json:"rent"`
	Equipment   string `json:"equipment"`
}


func PostHall(c *gin.Context) {

    var hall NewHall

    if err := c.ShouldBindJSON(&hall); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newhall := models.Hall{Name: hall.Name, Description: hall.Description, Capacity: hall.Capacity, Rent: hall.Rent, Equipment: hall.Equipment}

    db, err := models.Database()
    if err != nil {
        log.Println(err)
    }

    if err := db.Table("hall_details").Create(&newhall).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, newhall)
}


func GetHall(c *gin.Context) {

    var halls []models.Hall

    db, err := models.Database()
    if err != nil {
        log.Println(err)
    }

    // var count int64
    if err := db.Table("hall_details").Find(&halls).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Hall not found"})
        return
    }

    c.JSON(http.StatusOK, halls)

}