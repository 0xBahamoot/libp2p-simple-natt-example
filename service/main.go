package main

import (
	"fmt"

	natt "github.com/0xBahamoot/go-libp2p-natt"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	h, _ := natt.CreateHost(9000, "")
	fmt.Println(h.GetBroadcastAddrInfo())
	_, err := peerInfoFromString(h.GetBroadcastAddrInfo())
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
