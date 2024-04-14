package main

import "server/network"

func main() {
	netConfig := network.NetConfig{
		BindAdress: "0.0.0.0",
		Port:       8000,
	}

	startLifeGameServer(netConfig)
}
