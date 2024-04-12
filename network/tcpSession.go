package network

import "sync"

func (sessionMgr *TcpSessionManager) appendSession() {

}

func (sessionMgr *TcpSessionManager) removeSession() {

}

func (sessionMgr *TcpSessionManager) allocSession() {

}

func (sessionMgr *TcpSessionManager) createSessionPool(poolSize int) {
	sessionMgr.SessionIndexPool = NewCappedDeque(poolSize)

	for index := 0; index < poolSize; index++ {
		sessionMgr.SessionIndexPool.Append(index)
	}
}

func createSessionManager(networkFunctor SessionNetworkFunctor) *TcpSessionManager {
	sessionMgr := &TcpSessionManager{
		SessionMap: sync.Map{},
	}

	sessionMgr.NetworkFunctor = networkFunctor
	sessionMgr.createSessionPool(POOL_SIZE)

	return sessionMgr
}
