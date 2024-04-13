package network

import "encoding/binary"

/*
패킷 헤더의 구조는 아래와 같음
패킷전체크기(2byte) + 패킷ID(2byte) + 패킷 Type(1byte)
패킷의 2바이트를 꺼내서 패킷의 전체 크기를 구한다.
*/
func PacketTotalSize(data []byte) int16 {
	totalSize := binary.LittleEndian.Uint16(data)
	return int16(totalSize)
}
