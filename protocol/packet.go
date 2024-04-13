package protocol

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

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

var _packetHeaderSize int16

func InitPacketHeaderSize() {
	_packetHeaderSize = PacketHeaderSize()
}

/*
전체 패킷에서 총 크기를 제외한 다음 2바이트를 꺼내옴
*/
func PeekPacketID(rawData []byte) int16 {
	packetID := binary.LittleEndian.Uint16(rawData[2:])
	return int16(packetID)
}

/*
전체 패킷에서 헤더를 뺸 만큼 바디로 지정
*/
func PeekPacketBody(rawData []byte) (int16, []byte) {
	headerSize := _packetHeaderSize
	totalSize := int16(binary.LittleEndian.Uint16(rawData))
	bodySize := totalSize - headerSize

	if bodySize > 0 {
		return bodySize, rawData[headerSize:]
	}

	return bodySize, []byte{}
}

/*
패킷헤더의 크기를 사전에 구함
*/
func PacketHeaderSize() int16 {
	var header Header
	hSize := unsafe.Sizeof(reflect.TypeOf(header))
	return (int16)(hSize)
}
