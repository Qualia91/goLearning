package server

import (
	"fmt"
	"messagingapp/common"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// create a handler function for server
func CreateServerHandler(tss *common.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// upgrades connection to websocket connection
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Printf("error while upgrading connection: %v\n", err)
		}

		tss.Mut.Lock()
		tss.Conns = append(tss.Conns, conn)
		tss.Mut.Unlock()

		// Server actions
		go func() {
			// defer the close of connection here so however the connection exits, it always closes properly
			defer func() {
				conn.Close()

				// remove from tss
				tss.Mut.Lock()
				for index, currentConn := range tss.Conns {
					if conn == currentConn {
						tss.Conns = append(tss.Conns[:index], tss.Conns[index+1:]...)
						break
					}
				}
				tss.Mut.Unlock()
			}()

			for {
				// using the higher level library wsutil to manage the read and write of the server
				msg, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					fmt.Printf("error while reading on server: %v\n", err)
					break
				}

				tss.MsgChan <- string(msg)

			}
		}()

	}

}
