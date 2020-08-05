package portal

import (
	"configservice"
	"encoding/json"
	"log"
	"net/http"
)

const HOST = "127.0.0.1:8081"
const (
	LOGIN      = "/v1.0/api/login"
	REGISTRY   = "/v1.0/api/users/add"
	NAMESPACES = "/v1.0/api/namespaces"
)

func init() {
	configservice.LoggerInit()
}

func WebServer() {
	http.HandleFunc(LOGIN, login)
	http.HandleFunc(REGISTRY, registry)
	http.HandleFunc(NAMESPACES, getNamespaces)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe(HOST, nil)

	if err != nil {
		log.Println("ListenAndServe error: ", err.Error())
	}
	log.Println("服务启动成功，正在监听")
}

//登陆接口
func login(w http.ResponseWriter, req *http.Request) {
	log.Println("登录接口调用")
	//获取客户端通过GET/POST方式传递的参数
	var loginReq LoginRequest

	// 将请求体中的 JSON 数据解析到结构体中
	// 发生错误，返回400 错误码
	err := json.NewDecoder(req.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loginResponse := Login(loginReq)
	handlerResponse(w, &loginResponse)
}

//注册用户接口
func registry(w http.ResponseWriter, req *http.Request) {
	log.Println("注册接口调用")
	var user configservice.Userspace
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Registry(user)
	handlerResponse(w, &response)
}

func getNamespaces(w http.ResponseWriter, req *http.Request) {

}

func handlerResponse(w http.ResponseWriter, response *CommonResponse) {
	w.Header().Set("Content-Type", "application/json")
	error := json.NewEncoder(w).Encode(response)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
}
