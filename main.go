package main

import (
	"log"
	"net/http"

	"github.com/storyofhis/webrtc-go/controllers"
	"github.com/storyofhis/webrtc-go/repositories/models"
)

func init() {
	Map := make(map[string][]models.Participants)
	controllers.AllRooms.Map = Map
}

func main() {
	http.HandleFunc("/create", controllers.CreateRoomRequestHandler)
	http.HandleFunc("/join", controllers.JoinRoomRequestHandler)
	log.Println("Starting Server on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
