package main


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example.com/rest-api/models"
)
func main(){
	server := gin.Default()

	server.GET("/events", getEvents)	//GET POST PUT PATCH DELETE
	server.POST("/events")

	server.Run(":8081")	//localhost during development
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}


func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not Parse Data!"})
		return
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message" : "Event Created!", "event" : event})
}