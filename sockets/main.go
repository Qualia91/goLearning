package main

import "net/http"

//TODO: These are both horrible implementations because i haven't properly done the
// control of creation. Need to make this better. Implement it via channels and messaging
// func main() {

// 	networkType := "tcp"
// 	networkAddress := "localhost:6000"

// 	// create wait group to make sure client and server both finish before app is closed
// 	wg := &sync.WaitGroup{}

// 	wg.Add(1)
// 	go StartTCPServer(networkType, networkAddress, "The server heard from the client: ", wg)
// 	wg.Add(1)
// 	go StartTCPClient(networkType, networkAddress, "This is the message from the client", wg)

// 	// now wait on wg to be finished
// 	wg.Wait()

// 	networkType = "udp4"
// 	networkAddress = "127.0.0.1:6001"

// 	wg.Add(1)
// 	go StartUDPServer(networkType, networkAddress, "The server heard from the client: ", wg)
// 	wg.Add(1)
// 	go StartUDPClient(networkType, networkAddress, "This is the message from the client", wg)

// 	// now wait on wg to be finished
// 	wg.Wait()
// }

func main() {

	// creating a websocket:
	// 1) Initiate a handshake
	// 2) Receive data frames from the client
	// 3) Send data frames to the client
	// 4) Close the handshake

	// in this example, im using the Gobwas library (https://github.com/gobwas/ws)
	// as its generally more featured and faster than others like Gorilla.
	http.ListenAndServe(":8080", http.HandlerFunc(EchoServerHandler))
}
