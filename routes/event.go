package routes

import (
	// "fmt"
	"github.com/avvsrohith/eventbooking_server_go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type NewEvent struct {
	Person_name    string `json:"person_name"		binding:"required"`
	Person_contact string `json:"person_contact"	binding:"required"`
	Date           int32  `json:"date"			binding:"required"`
	Event_name     int32  `json:"event_name"		binding:"required"`
	Start_time     string `json:"start_time" 	    binding:"required"`
	End_time       string `json:"end_time"		binding:"required"`
	Hall_id        string `json:"hall_id"		    binding:"required"`
}

type Event struct {
	Person_name    string `json:"person_name"`
	Person_contact string `json:"person_contact"`
	Date           int32  `json:"date"`
	Event_name     int32  `json:"event_name"`
	Start_time     string `json:"start_time"`
	End_time       string `json:"end_time"`
	Hall_id        string `json:"hall_id"`
}

func PostEvent(c *gin.Context) {

	var event NewEvent

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newevent := models.Event{Name: event.Name, Description: event.Description, Capacity: event.Capacity, Rent: event.Rent, Equipment: event.Equipment}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Table("event_details").Create(&newevent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newevent)
}

func Getevent(c *gin.Context) {

	var events []models.Event

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Table("event_details").Find(&events).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, events)

}
