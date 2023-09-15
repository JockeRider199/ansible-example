package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	stop := make(chan bool, 1)

	go PrintEachSec(stop)

	select {
	case <-sc:
		stop <- true
		<-sc
		os.Exit(0)
	}
}

func PrintEachSec(sc chan bool) {
	for {
		select {
		case <-sc:
			os.Exit(0)
		default:
			log.Println("Hello World")
			time.Sleep(1 * time.Second)
		}
	}
}
