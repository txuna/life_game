package protocol

import (
	"encoding/binary"
	"reflect"
	"server/network"

	"github.com/vmihailenco/msgpack"
)

const (
	PACKET_TYPE_NORMAL   = 0
	PACKET_TYPE_COMPRESS = 1
	PACKET_TYPE_SECURE   = 2
)

const (
	MAX_USER_ID_BYTE_LENGTH      = 16
	MAX_USER_PW_BYTE_LENGTH      = 16
	MAX_USER_NAME_BYTE_LENGTH    = 16
	MAX_CHAT_MESSAGE_BYTE_LENGTH = 126
)

var _packetHeaderSize int16

func InitPacketHeaderSize() {
	_packetHeaderSize = PacketHeaderSize()
}

func GetPacketHeaderSize() int16 {
	return _packetHeaderSize
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
패킷 헤더를 추가한다.
*/
func EncodingPacketHeader(writer *network.RawPacketData, totalSize int16, pktId int16, pktType int8) {
	writer.WriteS16(totalSize)
	writer.WriteS16(pktId)
	writer.WriteS8(pktType)
}

/*
패킷 헤더를 분석한다.
*/
func DecodingPacketHeader(header *Header, data []byte) {
	reader := network.MakeReader(data, true)
	header.TotalSize, _ = reader.ReadS16()
	header.ID, _ = reader.ReadS16()
	header.PacketType, _ = reader.ReadS8()
}

/*
패킷헤더의 크기를 사전에 구함bi
*/
func PacketHeaderSize() int16 {
	var header Header
	hSize := network.Sizeof(reflect.TypeOf(header))
	return (int16)(hSize)
}

func EncodingPacket(pkId int16, pkType int8, v interface{}) ([]byte, int16) {
	raw, _ := msgpack.Marshal(v)
	totalSize := _packetHeaderSize + int16(len(raw))
	sendBuff := make([]byte, totalSize)

	writer := network.MakeWrite(sendBuff, true)
	EncodingPacketHeader(&writer, int16(totalSize), pkId, pkType)
	writer.WriteBytes(raw)
	return sendBuff, int16(totalSize)
}

func DecodingPacket(bodyData []byte, v interface{}) bool {
	err := msgpack.Unmarshal(bodyData, v)
	return err == nil
}

/* 로그인 요청 */
func (loginReq LoginReqPacket) EncodingPacket() ([]byte, int16) {
	totalSize := _packetHeaderSize + MAX_USER_ID_BYTE_LENGTH + MAX_USER_PW_BYTE_LENGTH
	sendBuf := make([]byte, totalSize)
	writer := network.MakeWrite(sendBuf, true)
	EncodingPacketHeader(&writer, totalSize, PACKET_ID_LOGIN_REQ, 0)
	writer.WriteBytes(loginReq.UserID[:])
	writer.WriteBytes(loginReq.UserPW[:])
	return sendBuf, totalSize
}

func (loginReq *LoginReqPacket) Decoding(bodyData []byte) bool {
	bodySize := MAX_USER_ID_BYTE_LENGTH + MAX_USER_PW_BYTE_LENGTH
	if len(bodyData) != bodySize {
		return false
	}

	reader := network.MakeReader(bodyData, true)

	var err error
	loginReq.UserID, err = reader.ReadBytes(MAX_USER_ID_BYTE_LENGTH)
	if err != nil {
		return false
	}

	loginReq.UserPW, err = reader.ReadBytes(MAX_USER_PW_BYTE_LENGTH)
	return err == nil
}

func (loginRes LoginResPacket) EncodingPacket() ([]byte, int16) {
	totalSize := _packetHeaderSize + 2
	sendBuf := make([]byte, totalSize)
	writer := network.MakeWrite(sendBuf, true)
	EncodingPacketHeader(&writer, totalSize, PACKET_ID_LOGIN_RES, 0)
	writer.WriteS16(loginRes.ErrorCode)
	return sendBuf, totalSize
}

/* 로그인 응답 */
func (loginRes *LoginResPacket) Decoding(bodyData []byte) bool {
	bodySize := 2
	if len(bodyData) != bodySize {
		return false
	}

	reader := network.MakeReader(bodyData, true)

	var err error
	loginRes.ErrorCode, err = reader.ReadS16()
	return err == nil
}

func (joinReq JoinReqPacket) EncodingPacket() ([]byte, int16) {
	totalSize := _packetHeaderSize + MAX_USER_ID_BYTE_LENGTH + MAX_USER_PW_BYTE_LENGTH + MAX_USER_NAME_BYTE_LENGTH
	sendBuf := make([]byte, totalSize)
	writer := network.MakeWrite(sendBuf, true)
	EncodingPacketHeader(&writer, totalSize, PACKET_ID_LOGIN_REQ, 0)
	writer.WriteBytes(joinReq.UserID[:])
	writer.WriteBytes(joinReq.UserPW[:])
	writer.WriteBytes(joinReq.UserName[:])
	return sendBuf, totalSize
}

/* 회원가입 요청 */
func (joinReq *JoinReqPacket) Decoding(bodyData []byte) bool {
	bodySize := MAX_USER_ID_BYTE_LENGTH + MAX_USER_PW_BYTE_LENGTH + MAX_USER_NAME_BYTE_LENGTH
	if len(bodyData) != bodySize {
		return false
	}
	reader := network.MakeReader(bodyData, true)

	var err error
	joinReq.UserID, err = reader.ReadBytes(MAX_USER_ID_BYTE_LENGTH)
	if err != nil {
		return false
	}

	joinReq.UserPW, err = reader.ReadBytes(MAX_USER_PW_BYTE_LENGTH)
	if err != nil {
		return false
	}

	joinReq.UserName, err = reader.ReadBytes(MAX_USER_NAME_BYTE_LENGTH)
	return err == nil
}

/* 회원가입 응답 */
func (joinRes JoinResPacket) EncodingPacket() ([]byte, int16) {
	totalSize := _packetHeaderSize + 2
	sendBuf := make([]byte, totalSize)
	writer := network.MakeWrite(sendBuf, true)
	EncodingPacketHeader(&writer, totalSize, PACKET_ID_JOIN_RES, 0)
	writer.WriteS16(joinRes.ErrorCode)
	return sendBuf, totalSize
}

func (joinRes *JoinResPacket) Decoding(bodyData []byte) bool {
	bodySize := 2
	if len(bodyData) != bodySize {
		return false
	}

	reader := network.MakeReader(bodyData, true)

	var err error
	joinRes.ErrorCode, err = reader.ReadS16()
	return err == nil
}

/* 핑 요청 */
func (pingReq PingReqPacket) EncodingPacket() ([]byte, int16) {
	totalSize := _packetHeaderSize + int16(network.Sizeof(reflect.TypeOf(int8(0))))
	sendBuf := make([]byte, totalSize)
	writer := network.MakeWrite(sendBuf, true)
	EncodingPacketHeader(&writer, totalSize, PACKET_ID_PING_REQ, 0)
	writer.WriteS8(pingReq.Ping)
	return sendBuf, totalSize
}

func (pingReq *PingReqPacket) Decoding(bodyData []byte) bool {
	bodySize := network.Sizeof(reflect.TypeOf(int8(0)))
	if len(bodyData) != bodySize {
		return false
	}

	reader := network.MakeReader(bodyData, true)
	var err error
	pingReq.Ping, err = reader.ReadS8()
	return err == nil
}

/* 핑 응답 */
func (pingRes PingResPacket) EncodingPacket() ([]byte, int16) {
	totalSize := _packetHeaderSize + int16(network.Sizeof(reflect.TypeOf(int8(0))))
	sendBuf := make([]byte, totalSize)
	writer := network.MakeWrite(sendBuf, true)
	EncodingPacketHeader(&writer, totalSize, PACKET_ID_PING_RES, 0)
	writer.WriteS8(pingRes.Pong)
	return sendBuf, totalSize
}

func (pingRes *PingResPacket) Decoding(bodyData []byte) bool {
	bodySize := network.Sizeof(reflect.TypeOf(int8(0)))
	if len(bodyData) != bodySize {
		return false
	}

	reader := network.MakeReader(bodyData, true)

	var err error
	pingRes.Pong, err = reader.ReadS8()

	return err == nil
}
