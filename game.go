package main

import (
	"fmt"
	"server/network"
)

type LifeGameServer struct {
}

func startLifeGameServer() {
	server := LifeGameServer{}

	/*
		네트워크의 콜백함수를 지정한다.
		네트워크 모듈에서 패킷이 도착하면 지정한 콜백함수를 호출하여 처리한다.
	*/
	snFunctor := network.SessionNetworkFunctor{
		OnConnect:           server.OnConnect,
		OnClose:             server.OnClose,
		OnReceive:           server.OnReceive,
		PacketTotalSizeFunc: network.PacketTotalSize,
		PacketHeaderSize:    network.PACKET_HEADER_SIZE,
	}

	network.StartServerBlock(snFunctor)
}

/*
클라이언트의 접속이 끊어졌을 때 콜백
*/
func (server *LifeGameServer) OnClose(sessionUniqueId uint64, sessionId int32) {
	fmt.Printf("Client Disconnected:%d - %d\n", sessionUniqueId, sessionId)
}

/*
클라이언트가 접속일 하였을 때 콜백
게임 세션 만들어 줘야함
*/
func (server *LifeGameServer) OnConnect(sessionUniqueId uint64, sessionId int32) {
	fmt.Printf("New Client Connected:%d - %d\n", sessionUniqueId, sessionId)
}

/*
클라이언트가 데이터를 주었을 때 콜백
인/디코딩 작업 들어가야함
*/
func (server *LifeGameServer) OnReceive(sessionUniqueId uint64, sessionId int32, packet []byte) {
	fmt.Printf("Client Send Message:%d-%d: %s\n", sessionUniqueId, sessionId, packet)
	server.DistributePacket(sessionUniqueId, sessionId, packet)
}
