package clientsession

import (
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
