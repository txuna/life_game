package protocol

type Header struct {
	TotalSize  int16
	ID         int16
	PacketType int8
}

type Packet struct {
	UserSessionIndex       int32
	UserSessionUniqueIndex uint64
	Id                     int16
	DataSize               int16
	Data                   []byte
}

/* 로그인 요청 */
type LoginReqPacket struct {
	UserID []byte
	UserPW []byte
}

/* 로그인 응답 */
type LoginResPacket struct {
	ErrorCode int16
}

/* 회원가입 요청 */
type JoinReqPacket struct {
	UserID   []byte
	UserPW   []byte
	UserName []byte
}

/* 회원가입 응답 */
type JoinResPacket struct {
	ErrorCode int16
}

/* 핑 요청 */
type PingReqPacket struct {
	Ping int8 `msgpack:"ping"`
}

/* 핑 응답 */
type PingResPacket struct {
	Pong int8 `msgpack:"pong"`
}
