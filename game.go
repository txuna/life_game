package main

import (
	"fmt"
	clientsession "server/clientSession"
	"server/network"
	"server/protocol"
)

type LifeGameServer struct {
	ServerIndex int
	netConfig   network.NetConfig
	PacketChan  chan protocol.Packet
}

func startLifeGameServer(netConfig network.NetConfig) {
	server := LifeGameServer{
		ServerIndex: 1,
		netConfig:   netConfig,
	}

	protocol.InitPacketHeaderSize()

	/*
		클라이언트 세션매니저 초기화
	*/
	clientsession.Init()

	/*
		채널 버퍼 256으로 설정
	*/
	server.PacketChan = make(chan protocol.Packet, 256)

	/*
		패킷을 처리하는 고루틴
	*/
	go server.PacketProcessGoroutine()

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

	network.StartServerBlock(netConfig, snFunctor)
}

/*
클라이언트의 접속이 끊어졌을 때 콜백
*/
func (server *LifeGameServer) OnClose(sessionUniqueId uint64, sessionId int32) {
	fmt.Printf("Client Disconnected:%d - %d\n", sessionUniqueId, sessionId)
	result := clientsession.RemoveSession(sessionUniqueId)
	if !result {
		fmt.Println("remove client session failed")
	} else {
		fmt.Println("remove client session")
	}
}

/*
클라이언트가 접속일 하였을 때 콜백
게임 세션 만들어 줘야함
*/
func (server *LifeGameServer) OnConnect(sessionUniqueId uint64, sessionId int32) {
	fmt.Printf("New Client Connected:%d - %d\n", sessionUniqueId, sessionId)
	result := clientsession.AddSession(sessionUniqueId, sessionId)
	if !result {
		fmt.Println("append client session failed")
	} else {
		fmt.Println("append client session")
	}
}

/*
클라이언트가 데이터를 주었을 때 콜백
인/디코딩 작업 들어가야함
*/
func (server *LifeGameServer) OnReceive(sessionUniqueId uint64, sessionId int32, packet []byte) {
	fmt.Printf("Client Send Message:%d-%d: %s\n", sessionUniqueId, sessionId, packet)
	server.DistributePacket(sessionUniqueId, sessionId, packet)
}
