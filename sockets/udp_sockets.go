package main

import (
	"fmt"
	"net"
	"sync"
)

func StartUDPClient(networkType, address, message string, wg *sync.WaitGroup) (returnString string, err error) {

	// defer a wait group done so that whatever happens, the wg is release
	defer func(wg *sync.WaitGroup) {
		wg.Done()
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		} else {
			fmt.Println(returnString)
		}
	}(wg)

	// the two variables in ResolveUDPAddr are:
	// 1) Network Type: Here its udp
	// 2) Address: eg localhost:6000
	// This basically checks for possible connection point
	udpAddr, err := net.ResolveUDPAddr(networkType, address)
	if err != nil {
		return "", err
	}

	// now we connect. The nil in the middle means it will use local network
	conn, err := net.DialUDP(networkType, nil, udpAddr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send a message
	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return
}

func StartUDPServer(networkType, address, message string, wg *sync.WaitGroup) (returnString string, err error) {

	// defer a wait group done so that whatever happens, the wg is release
	defer func(wg *sync.WaitGroup) {
		wg.Done()
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		} else {
			fmt.Println(returnString)
		}
	}(wg)

	udpAddr, err := net.ResolveUDPAddr(networkType, address)
	if err != nil {
		return "", err
	}

	// create a listener
	conn, err := net.ListenUDP("udp", udpAddr)
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

	returnString = message + string(buf)

	return

}
