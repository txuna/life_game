package packet

const (
	PACKET_TYPE_NORMAL   = 0
	PACKET_TYPE_COMPRESS = 1
	PACKET_TYPE_SECURE   = 2
)

const (
	MAX_USER_ID_BYTE_LENGTH      = 16
	MAX_USER_PW_BYTE_LENGTH      = 16
	MAX_CHAT_MESSAGE_BYTE_LENGTH = 126
)

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
