package main

import (
	"bytes"
	"fmt"
	"server/network"
	"server/protocol"
	"time"
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

	roomUpdateTimerTicker := time.NewTicker(time.Second)
	defer roomUpdateTimerTicker.Stop()

	for {
		select {
		case packet := <-server.PacketChan:
			{
				sessionId := packet.UserSessionIndex
				sessionUniqueId := packet.UserSessionUniqueIndex
				bodySize := packet.DataSize
				bodyData := packet.Data

				if packet.Id == protocol.PACKET_ID_LOGIN_REQ {
					ProcessPacketLogin(sessionUniqueId, sessionId, bodySize, bodyData)
				} else if packet.Id == protocol.PACKET_ID_JOIN_REQ {

				} else if packet.Id == protocol.PACKET_ID_PING_REQ {
					ProcessPacketPing(sessionUniqueId, sessionId, bodySize, bodyData)
				} else {
					fmt.Println("Invalid Packet ID")
				}
			}

		/*
			초당 호출되는 로직
			이때 방을 업데이트를 진행한다.
		*/
		case curTime := <-roomUpdateTimerTicker.C:
			{
				_ = curTime
				//fmt.Println("Update Room")
				//fmt.Println(curTime)
			}
		}
	}
}

func ProcessPacketPing(sessionUniqueId uint64, sessionId int32, bodySize int16, bodyData []byte) {
	var request protocol.PingReqPacket
	result := (&request).Decoding(bodyData)
	if !result {
		return
	}

	fmt.Printf("[%d]-[%d] Client Send PING\n", sessionUniqueId, sessionId)
}

func ProcessPacketJoin(sessionUniqueId uint64, sessionId int32, bodySize int16, bodyData []byte) {
	var request protocol.JoinReqPacket
	result := (&request).Decoding(bodyData)
	if !result {
		return
	}

	userID := bytes.Trim(request.UserID[:], "\x00")
	userPW := bytes.Trim(request.UserPW[:], "\x00")

	if len(userID) <= 0 || len(userPW) <= 0 {
		return
	}

	fmt.Printf("[%d]request join: %s-%s\n", sessionUniqueId, string(userID), string(userPW))
	/* 회원가입 */
	sendJoinResult(sessionUniqueId, sessionId, protocol.ERROR_CODE_NONE)
}

func ProcessPacketLogin(sessionUniqueId uint64, sessionId int32, bodySize int16, bodyData []byte) {
	var request protocol.LoginReqPacket
	result := (&request).Decoding(bodyData)
	if !result {
		return
	}

	userID := bytes.Trim(request.UserID[:], "\x00")
	userPW := bytes.Trim(request.UserPW[:], "\x00")

	if len(userID) <= 0 || len(userPW) <= 0 {
		return
	}

	fmt.Printf("[%d]request login: %s - %s\n", sessionUniqueId, string(userID), string(userPW))
	/* 로그인 */
	sendLoginResult(sessionUniqueId, sessionId, protocol.ERROR_CODE_NONE)
}

func sendLoginResult(sessionUniqueId uint64, sessionId int32, result int16) {
	loginRes := protocol.LoginResPacket{
		ErrorCode: result,
	}

	packet, _ := loginRes.EncodingPacket()
	network.SendToClient(sessionUniqueId, sessionId, packet)
}

func sendJoinResult(sessionUniqueId uint64, sessionId int32, result int16) {
	joinRes := protocol.JoinResPacket{
		ErrorCode: result,
	}

	packet, _ := joinRes.EncodingPacket()
	network.SendToClient(sessionUniqueId, sessionId, packet)
}
