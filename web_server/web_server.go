package main

import (
	"bytes"
	"net"
	"net/http"
	"os"
	"time"
)

func hostinfo(w http.ResponseWriter, r *http.Request) {
	var addr string
	message := "捷越ECS测试"
	// Retrive the hostname.
	hostname, _ := os.Hostname()

	// Retrive the net info.
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
			// Don't use random as we have a real address
			addr = i.HardwareAddr.String()
			break
		}
	}

	// Retrive the current time.
	timeFormat := time.Now()
	currentTime := timeFormat.String()

	// Version info
	version := "The current version is: 0.1.0"

	message = message + "\n" + hostname + "\n" + addr + "\n" + currentTime + "\n" + version
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", hostinfo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
