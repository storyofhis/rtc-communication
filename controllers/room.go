package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/storyofhis/webrtc-go/controllers/views"
	"github.com/storyofhis/webrtc-go/repositories/gorm"
)

var AllRooms gorm.RoomMap

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()

	log.Println(AllRooms.Map)
	json.NewEncoder(w).Encode(views.Response{
		RoomID: roomID,
	})
}

var broadcast = make(chan BroadcastMsg)

func broadcaster() {
	// msg := <-broadcast
	// log.Println(msg.RoomID)
	for {
		msg := <-broadcast
		for _, client := range AllRooms.Map[msg.RoomID] {
			// log.Println(client.Conn)
			// log.Println(msg.Client)
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
}

// JoinRoomRequestHandler will join the client in particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID, ok := r.URL.Query()["roomID"]

	if !ok {
		log.Println("roomID missing in URL Parameters")
		return
	}

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	AllRooms.InsertIntoRoom(roomID[0], false, ws)
	go broadcaster()

	for {
		var msg BroadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error :", err)
		}

		msg = BroadcastMsg{
			Client: ws,
			RoomID: roomID[0],
		}

		log.Println("this is messagee :", msg.Message)

		broadcast <- msg
	}
}
