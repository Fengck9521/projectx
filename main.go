package main

import (
	"projectx/network"
	"time"
)

func main() {

	trLocal := network.NewLocalTransport("local")
	trRemote := network.NewLocalTransport("remote")

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello"))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}
	server := network.NewServer(opts)
	server.Start()
}
