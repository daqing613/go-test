package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
)

func main() {
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}

	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Run()
	stats := pinger.Statistics()
	fmt.Println(stats)
}
