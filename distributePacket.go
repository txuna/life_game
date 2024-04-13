package main

import (
	"fmt"
	"server/protocol"
)

/*
실제 패킷에서 아이디, 바디를 가지고옴
*/
func (server *LifeGameServer) DistributePacket(sessionUniqueId uint64, sessionId int32, packetData []byte) {
	packetID := protocol.PeekPacketID(packetData)
	bodySize, packetBody := protocol.PeekPacketBody(packetData)

	packet := protocol.Packet{
		UserSessionIndex:       sessionId,
		UserSessionUniqueIndex: sessionUniqueId,
		Id:                     packetID,
		DataSize:               bodySize,
		Data:                   make([]byte, bodySize),
	}

	copy(packet.Data, packetBody)

	/*
		수신한 패킷을 처리하는 채널로 보냄
	*/
	server.PacketChan <- packet

}

/*
DistributePacket함수에서 채널형식으로 넘겨줌
실질적인 패킷 처리 함수
*/
func (server *LifeGameServer) PacketProcessGoroutine() {
	for {
		select {
		case packet := <-server.PacketChan:
			{
				fmt.Println(packet)
			}
		}
	}
}
