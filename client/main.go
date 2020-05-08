package main

import (
	"fmt"
	"time"

	natt "github.com/0xBahamoot/go-libp2p-natt"
)

func main() {
	h, err := natt.CreateHost(9001, "/ip4/172.0.0.2/tcp/9000/p2p/QmWqKXPcCxRHVC48UkqP7aqd2tqc2chhRd46ydVyhLdGoa")
	if err != nil {
		panic(err)
	}
	go func() {
		time.Sleep(20 * time.Second)
		fmt.Println(h.GetNATType())
		fmt.Println(h.GetHost().Addrs())
	}()
	select {}
}
