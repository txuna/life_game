# Client - Server Architecture 

### Protocol 

#### Register Request 

#### Register Response 

#### Login Request 

#### Login Response

### Client Connect To Server
클라이언트의 접속 요청이 오면 서버의 네트워크 모듈은 TcpSession을 만들고 Network 관련 콜백 함수 등록 및 TcpSessionManager에 추가 
이때 클라이언트 동시접속자 수는 POOL_SIZE로 지정하고 Deque로 관리  

클라이언트 TcpSession이 만들어지면 Deque로 부터 인덱스를 부여받고 세션할당

패킷의 처리는 msgpack을 사용, RING BUFFER를 사용?
