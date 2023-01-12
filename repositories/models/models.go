package models

import "github.com/gorilla/websocket"

type Participants struct {
	Host bool
	Conn *websocket.Conn
}
