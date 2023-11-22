package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var db *sql.DB = nil

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "--http:home--")

}
func writePixel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "--http:writePixel--")
	db.Query("select count(*) from pixels")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
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
func setupRoutes() {
	http.HandleFunc("/write_pixel", writePixel)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	port := 8084
	fmt.Println(fmt.Sprintf("pixel-socket-server listening on %d", port))
	setupRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

	dburl := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
	database, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
	}
	db = database
}
