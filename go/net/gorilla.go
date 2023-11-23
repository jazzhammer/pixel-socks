package net

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	. "main/go/db"
	. "main/go/model"
	"net/http"
)

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		var pixel Pixel
		parseErr := json.Unmarshal(p, &pixel)
		if parseErr != nil {
			fmt.Println(err)
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		SavePixel(pixel)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("connected")
	err = ws.WriteMessage(1, []byte("helo"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func StartGorillaWs() {
	fmt.Println(fmt.Sprintf("pixel-socket-server(gorilla) listening on %d", wsPort))
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", wsPort), nil))

}
