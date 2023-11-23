package main

import (
	_ "github.com/lib/pq"
	. "main/go/db"
	. "main/go/net"
)

func main() {
	StartDb()
	StartGobwasWs()
}
