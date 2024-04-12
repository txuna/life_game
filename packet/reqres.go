package packet

type RegisterRequest struct {
	Email           string `msgpack:"email"`
	NickName        string `msgpack:"nickname"`
	Password        string `msgpack:"password"`
	ConfirmPassword string `msgpack:"confirm_password"`
}

type RegisterResponse struct {
	ErrorCode int64 `msgpack:"error_code"`
}

type LoginRequest struct {
	Email    string `msgpack:"email"`
	Password string `msgpack:"password"`
}

type LoginResponse struct {
	ErrorCode int64  `msgpack:"error_code"`
	Token     string `msgpack:"token"`
}

type UserInfoRequest struct {
	Token string `msgpack:"token"`
}

type UserInfoResponse struct {
	NickName string `msgpack:"nickname"`
	Email    string `msgpack:"email"`
}

func CreateRegisterRequest() {

}

func ParseRegisterRequest() {

}

func CreateRegisterResponse() {

}

func ParseRegisterResponse() {

}

func CreateLoginRequest() {

}

func ParseLoginRequest() {

}

func CreateLoginResponse() {

}

func ParseLoginResponse() {

}
