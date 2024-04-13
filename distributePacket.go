package main

/*
실제 패킷에서 아이디, 바디를 가지고옴
*/
func (server *LifeGameServer) DistributePacket(sessionUniqueId uint64, sessionId int32, packet []byte) {

}

/*
DistributePacket함수에서 채널형식으로 넘겨줌
*/
func (server *LifeGameServer) PacketProcessGoroutine() {

}
