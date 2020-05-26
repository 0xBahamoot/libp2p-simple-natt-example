package main

import (
	"context"
	"fmt"

	natt "github.com/0xBahamoot/go-libp2p-natt"
	"github.com/libp2p/go-libp2p-core/crypto"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	privB, err := crypto.ConfigDecodeKey("CAESQFL9/AchgmVPFj9fD5piHXKVZsdNEAN8TXu7BAfR4sZJbxWBcJu3se8DDSENsY47C6HHdvumXYzarQVBUULRifg=")
	if err != nil {
		panic(err)
	}
	priv, err := crypto.UnmarshalPrivateKey(privB)
	if err != nil {
		panic(err)
	}
	h, err := natt.CreateHost(priv, 9000, "", ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(h.GetBroadcastAddrInfo())

	select {}
}
