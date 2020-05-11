package main

import (
	"fmt"

	natt "github.com/0xBahamoot/go-libp2p-natt"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	privB, err := crypto.ConfigDecodeKey("CAESQFL9/AchgmVPFj9fD5piHXKVZsdNEAN8TXu7BAfR4sZJbxWBcJu3se8DDSENsY47C6HHdvumXYzarQVBUULRifg=")
	if err != nil {
		panic(err)
	}
	priv, err := crypto.UnmarshalPrivateKey(privB)
	if err != nil {
		panic(err)
	}
	h, err := natt.CreateHost(priv, 9000, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(h.GetBroadcastAddrInfo())
	_, err = peerInfoFromString(h.GetBroadcastAddrInfo())
	if err != nil {
		fmt.Println(err)
	}
	select {}
}

func peerInfoFromString(peerAddr string) (*peer.AddrInfo, error) {
	pAddr, err := ma.NewMultiaddr(peerAddr)
	if err != nil {
		return nil, err
	}

	pInfo, err := peer.AddrInfoFromP2pAddr(pAddr)
	if err != nil {
		return nil, err
	}
	return pInfo, nil
}
