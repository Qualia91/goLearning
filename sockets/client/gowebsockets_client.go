package main

import (
	"bufio"
	"log"
	"os"

	"github.com/sacOO7/gowebsocket"
)

// create a handler function for server
func main() {

	socket := gowebsocket.New("ws://localhost:8080/")

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Received connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Received message " + message)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		log.Println("Received binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received pong " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
	}

	socket.Connect()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		msg := scanner.Text()

		if msg == "quit" {
			return
		}

		socket.SendText(msg)

	}

}
