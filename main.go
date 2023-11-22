package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB = nil

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type RGBColor struct {
	R int
	G int
	B int
}

type Pixel struct {
	CanvasX int
	CanvasY int
	Color   RGBColor
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
		var pixel Pixel
		parseErr := json.Unmarshal(p, &pixel)
		if parseErr != nil {
			fmt.Println(err)
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		savePixel(pixel)
	}
}

func savePixel(pixel Pixel) {
	_, err := db.Exec(`INSERT INTO pixel (x, y, r, g, b) values ($1, $2, $3, $4, $5)`, pixel.CanvasX, pixel.CanvasY, pixel.Color.R, pixel.Color.G, pixel.Color.B)
	if err != nil {
		log.Fatalf("An error occured while executing insert: %v", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "--http:home--")

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
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	dburl := "postgresql://postgres@localhost/pixel_socks?sslmode=disable"
	database, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
		fmt.Println(fmt.Sprintf("database connection: failed: %s", err))
	} else {
		fmt.Println("database connection: ok")
	}
	db = database

	port := 8084
	fmt.Println(fmt.Sprintf("pixel-socket-server listening on %d", port))
	setupRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
