package gorm

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/storyofhis/webrtc-go/repositories/models"
)

type RoomMap struct {
	Mutex sync.Mutex
	Map   map[string][]models.Participants
}

// initializes the RoomMap struct
// func (room *RoomMap) NewRoom(Map map[string][]models.Participants) {
// 	room.Map = make(map[string][]models.Participants)
// }

// func (room *RoomMap) Init() {
// 	room.Map = make(map[string][]models.Participants)
// }

// Get will return the array of participants in the room
func (room *RoomMap) Get(roomID string) []models.Participants {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	return room.Map[roomID]
}

// CreateRoom generate a unique room ID and return it -> insert it in the hashmap
func (room *RoomMap) CreateRoom() string {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
		log.Println(b[i])
	}

	roomID := string(b)
	room.Map[roomID] = []models.Participants{}

	return roomID
}

// InsertIntoRoom will create a participant and add it in the hashmap
func (room *RoomMap) InsertIntoRoom(roomID string, host bool, coon *websocket.Conn) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	p := models.Participants{
		Host: host,
		Conn: coon,
	}

	log.Println("Inserting into Room with RoomID: ", roomID)
	room.Map[roomID] = append(room.Map[roomID], p)
}

// DeleteRoom deletes the room with the roomID
func (room *RoomMap) DeleteRoom(roomID string) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	delete(room.Map, roomID)
}
