package main

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	address     = "IP_ADDRESS:PORT"
	logFileName = "ping.log"
	timeout     = 20 * time.Second
	interval    = 10 * time.Second
)

func main() {
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("err log file open: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	for {
		checkAddress(logger)
		time.Sleep(interval)
	}
}

func checkAddress(logger *log.Logger) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		logger.Printf("DOWN - %s", err)
		return
	}
	conn.Close()
	logger.Println("UP")
}



