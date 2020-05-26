package main

import (
	"fmt"
	"time"

	"context"

	natt "github.com/0xBahamoot/go-libp2p-natt"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	h, err := natt.CreateHost(nil, 9001, "/ip4/50.0.0.103/tcp/9000/p2p/12D3KooWHHzSeKaY8xuZVzkLbKFfvNgPPeKhFBGrMbNzbm5akpqu", ctx)
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(h.GetNATType())
				fmt.Println(h.GetHost().Addrs())
				fmt.Println(h.GetBroadcastAddrInfo())
			}
		}
	}()
	select {}
}
