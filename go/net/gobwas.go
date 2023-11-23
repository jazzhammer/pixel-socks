package net

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
)

func StartGobwasWs() {
	fmt.Println(fmt.Sprintf("pixel-socket-server(gobwas) listening on %d", wsPort))
	http.ListenAndServe(fmt.Sprintf(":%d", wsPort), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Println(fmt.Sprintf("error upgrading http -> ws: %s", err.Error()))
		}
		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Println(fmt.Sprintf("error reading client data: %s", err.Error()))
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					fmt.Println(fmt.Sprintf("error writing client data: %s", err.Error()))
				}
			}
		}()
	}))
}
