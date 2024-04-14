package main

import (
	"encoding/json"
	"fmt"
	"os"
	"server/network"
)

func main() {
	netConfig, err := parseNetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	startLifeGameServer(netConfig)
}

func parseNetConfig() (network.NetConfig, error) {
	var netConfig network.NetConfig
	file, err := os.Open("./config/net_config.json")
	if err != nil {
		return netConfig, err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(netConfig)

	return netConfig, err
}
