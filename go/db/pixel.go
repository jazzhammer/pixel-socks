package db

import (
	"log"
	. "main/go/model"
)

func SavePixel(pixel Pixel) {
	_, err := Db.Exec(`INSERT INTO pixel (x, y, r, g, b) values ($1, $2, $3, $4, $5)`, pixel.CanvasX, pixel.CanvasY, pixel.Color.R, pixel.Color.G, pixel.Color.B)
	if err != nil {
		log.Fatalf("An error occured while executing insert: %v", err)
	}
}
