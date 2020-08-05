package portal

import "time"

//登陆请求入参
type LoginRequest struct {
	UserName string
	Password string
}

//登陆响应结果
type CommonResponse struct {
	Code      int
	Data      interface{}
	Timestamp time.Time
	Message   string
}

//登陆响应数据
type LoginDataResponse struct {
	Token string
}
