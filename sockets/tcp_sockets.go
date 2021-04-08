package main

import (
	"fmt"
	"net"
	"sync"
)

func StartTCPClient(networkType, address, message string, wg *sync.WaitGroup) (returnString string, err error) {

	// defer a wait group done so that whatever happens, the wg is release
	defer func(wg *sync.WaitGroup) {
		wg.Done()
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		} else {
			fmt.Println(returnString)
		}
	}(wg)

	// the two variables in ResolveTCPAddr are:
	// 1) Network Type: Here its tcp
	// 2) Address: eg localhost:6000
	// This basically checks for possible connection point
	tcpAddr, err := net.ResolveTCPAddr(networkType, address)
	if err != nil {
		return "", err
	}

	// now we connect. The nil in the middle means it will use local network
	conn, err := net.DialTCP(networkType, nil, tcpAddr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send a message
	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}

	// receive a message
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return "", err
	}

	returnString = string(buf)

	return
}

func StartTCPServer(networkType, address, message string, wg *sync.WaitGroup) (returnString string, err error) {

	// defer a wait group done so that whatever happens, the wg is release
	defer func(wg *sync.WaitGroup) {
		wg.Done()
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		} else {
			fmt.Println(returnString)
		}
	}(wg)

	tcpAddr, err := net.ResolveTCPAddr(networkType, address)
	if err != nil {
		return "", err
	}

	// create a listener
	listener, err := net.ListenTCP(networkType, tcpAddr)
	if err != nil {
		return "", err
	}
	defer listener.Close()

	// listen for incoming messages
	conn, err := listener.Accept()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// receive message
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return "", err
	}

	responseMessage := message + string(buf)

	// send message back
	if _, err = conn.Write([]byte(responseMessage)); err != nil {
		return "", err
	}

	returnString = "Server closing down"

	return

}
