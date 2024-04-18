package clientsession

import (
	"server/service"
	"sync"
)

type SessionManager struct {
	SessionMap sync.Map
}

var _sessionManager *SessionManager

func Init() {
	_sessionManager = createSessionManager()
}

func createSessionManager() *SessionManager {
	sessionMgr := &SessionManager{
		SessionMap: sync.Map{},
	}

	return sessionMgr
}

/*
Redis에 저장
user_[unique_id]
*/
func AddSession(sessionUniqueId uint64, sessionId int32) bool {
	_, result := FindSession(sessionUniqueId)
	if result {
		return false
	}

	session := &ClientSession{
		SessionUniqueID: sessionUniqueId,
		SessionID:       sessionId,
		IsAuth:          false,
	}

	_sessionManager.SessionMap.Store(sessionUniqueId, session)
	err := service.StoreUserInfo(sessionUniqueId, sessionId, false)
	if err != nil {
		/* 롤백 */
		RemoveSession(sessionUniqueId)
		return false
	}

	return true
}

func FindSession(sessionUniqueId uint64) (*ClientSession, bool) {
	if session, ok := _sessionManager.SessionMap.Load(sessionUniqueId); ok {
		return session.(*ClientSession), true
	}

	return nil, false
}

func RemoveSession(sessionUniqueId uint64) bool {
	err := service.RemoveUserInfo(sessionUniqueId)
	if err != nil {
		return false
	}

	_sessionManager.SessionMap.Delete(sessionUniqueId)
	return true
}
