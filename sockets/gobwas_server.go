package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// create a handler function for server
func EchoServerHandler(w http.ResponseWriter, r *http.Request) {

	// upgrades connection to websocket connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		fmt.Printf("error while upgrading connection: %v\n", err)
	}

	// Server actions
	go func() {
		// defer the close of connection here so however the connection exits, it always closes properly
		defer conn.Close()

		for {
			// using the higher level library wsutil to manage the read and write of the server
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				fmt.Printf("error while reading on server: %v\n", err)
				return
			}

			fmt.Printf("message received on client: %v\n", string(msg))

		}
	}()

	// server acting as client action
	go func() {
		defer conn.Close()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			msg := scanner.Text()

			if msg == "quit" {
				return
			}

			err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(msg))
			if err != nil {
				fmt.Printf("error while writing on server: %v\n", err)
				return
			}
		}
	}()

}
