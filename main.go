package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)    //GET POST PUT PATCH DELETE
	server.GET("/events/:id", getEvent) //events/1, events/5
	server.POST("/events", createEvent)

	server.Run(":8081") //localhost during development
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)
	fmt.Println("error is--->", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	fmt.Println("error is---->", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.Try agin later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Parse Data!"})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Create Event.Please try agin later."})
		return
	}
	fmt.Println("event is--->", event)
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}
