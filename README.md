# Client - Server Architecture 

### Packet

- Packet Header
패킷 헤더의 크기는 5바이트 

패킷의 총 크기(2바이트 header + body) + 패킷ID(2 바이트) + 패킷 Type(1 바이트) 

### Protocol 

1. Login Request 


2. Login Response 


3. Join Request 


4. Join Request

### Client Connect To Server
클라이언트의 접속 요청이 오면 서버의 네트워크 모듈은 TcpSession을 만들고 Network 관련 콜백 함수 등록 및 TcpSessionManager에 추가 
이때 클라이언트 동시접속자 수는 POOL_SIZE로 지정하고 Deque로 관리  

클라이언트 TcpSession이 만들어지면 Deque로 부터 인덱스를 부여받고 세션할당

패킷의 처리는 msgpack을 사용, RING BUFFER를 사용?
