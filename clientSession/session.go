package clientsession

import "server/protocol"

type ClientSession struct {
	SessionUniqueID uint64
	SessionID       int32
	UserID          [protocol.MAX_USER_ID_BYTE_LENGTH]byte
	UserIDLength    int8
	IsAuth          bool
}
