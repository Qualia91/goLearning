package main

import (
	"fmt"
	"time"
)

func main() {

	// make done channel
	done := make(chan interface{})

	hb, rs := DoWork(done, 4*time.Second)

	for {
		select {
		case _, ok := <-hb:
			if !ok {
				fmt.Println("Heartbeat channel not ok")
			}
			fmt.Println("Heartbeat")
		case _, ok := <-rs:
			if !ok {
				fmt.Println("Work done channel not ok")
			}
			fmt.Println("Work done")
		// this case is run when a heartbeat isn't received within 2 seconds
		case <-time.After(2 * time.Second):
			fmt.Println("Go routine is not working, restarting it")
			close(done)
			done = make(chan interface{})
			hb, rs = DoWork(done, time.Second)
		}
	}

}

func DoWork(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {

	// create return channels
	heartbeat := make(chan interface{})
	results := make(chan time.Time)

	// create and run go routine that does the work
	go func() {
		defer close(heartbeat)
		defer close(results)

		// create tick channels for pulse, and an artificial for work done to show heartbeat working
		pulse := time.Tick(pulseInterval)
		workChan := time.Tick(5 * pulseInterval)

		// create send pulse function that sends heartbeat
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}

		// create function to send result
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		// create work look that sends pulses, listens to done channel, and does work
		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workChan:
				sendResult(r)
			}
		}

	}()

	return heartbeat, results
}
