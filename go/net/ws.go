package net

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var wsPort = 8084

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
