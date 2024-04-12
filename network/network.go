package network

import (
	"fmt"
	"net"
	"sync"
)

const (
	POOL_SIZE = 512
)

/*
* 	RING BUFFER?
 */
type SessionNetworkFunctor struct {
	OnConnect func()
	OnClose   func()
	OnReceive func() /* RING BUFFER? */
}

type TcpSession struct {
	Index          int
	Conn           net.Conn
	NetworkFunctor SessionNetworkFunctor
}

/*
*	NEED TO LOCK!
 */
type TcpSessionManager struct {
	NetworkFunctor SessionNetworkFunctor
	SessionMap     sync.Map
	/* Queue */
	SessionIndexPool *Deque
}

func StartServerBlock(networkFunctor SessionNetworkFunctor) {

	tcpSessionManager := createSessionManager(networkFunctor)

	l, err := net.Listen("tcp", ":7000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		/*
			Create TCP Session
			Index : Allocate from Session Index Pool (Deque)
		*/
		newTcpSession := &TcpSession{
			Index:/* INDEX POOL */ 1,
			Conn:           conn,
			NetworkFunctor: tcpSessionManager.NetworkFunctor,
		}

		/*
			Append TCP Session at TCP Session Manager
		*/

		/*
			Handle Tcp Read
		*/
		go newTcpSession.handleTcpRead()
	}
}

func (session *TcpSession) handleTcpRead() {
	/* 클라이언트의 접속 처리 */
	session.NetworkFunctor.OnConnect()
}
