package main

import (
	"encoding/json"
	"fmt"
	"messagingapp/common"
	"messagingapp/server"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const port = "8080"

func main() {

	// make internal messaging channel
	msgChan := make(chan string)
	defer close(msgChan)

	tss := &common.Database{
		MsgChan: msgChan,
	}

	// set up loop to read channel and send to each connection
	go func(tss *common.Database) {

		for {

			jsonMsg := <-msgChan

			// convert the msg which is json into a Message type
			var msg common.Message
			json.Unmarshal([]byte(jsonMsg), &msg)

			if msg.Message == "quit" {
				return
			}

			tss.Mut.RLock()
			// write data to database
			tss.Messages = append(tss.Messages, msg)

			strMsg := common.MessageToJson(msg)

			// now send messages to clients
			for _, conn := range tss.Conns {
				err := wsutil.WriteServerMessage(conn, ws.OpText, []byte(strMsg))
				if err != nil {
					fmt.Printf("error while writing on server: %v\n", err)
					tss.Mut.RUnlock()
					return
				}
			}
			tss.Mut.RUnlock()
		}

	}(tss)

	// setup routes to simple web page
	http.HandleFunc("/serverwebsocket", server.CreateServerHandler(tss))
	http.HandleFunc("/messaging", server.CreateMessagePageHandler(tss))
	http.HandleFunc("/", server.CreateMessagePageHandler(tss))

	// setup route to resources
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("resources"))))

	// start server
	http.ListenAndServe(":"+port, nil)
}
