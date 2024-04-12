package main

import "server/network"

func main() {

	/* network callback function */
	snFunctor := network.SessionNetworkFunctor{
		OnConnect: OnConnect,
		OnClose:   OnClose,
		OnReceive: OnReceive,
	}

	network.StartServerBlock(snFunctor)
}
