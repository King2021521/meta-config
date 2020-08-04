package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storage"
	"time"
)

var mgo = storage.Mgo{"127.0.0.1:27017", "meta-config", "user_space"}

func WebServer() {
	http.HandleFunc("/v1.0/api/login", login)
	http.HandleFunc("/v1.0/api/registry", registry)
	http.HandleFunc("/v1.0/api/namespaces", getNamespaces)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("127.0.0.1:8081", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	//获取客户端通过GET/POST方式传递的参数
	var loginReq LoginRequest

	// 将请求体中的 JSON 数据解析到结构体中
	// 发生错误，返回400 错误码
	err := json.NewDecoder(req.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var loginResponse LoginResponse
	if loginReq.UserName != "zhangsan" || loginReq.Password != "123456" {
		loginResponse = LoginResponse{Code: 401, Timestamp: time.Now(), Message: "登陆失败，用户名或密码错误"}
	} else {
		loginResponse = LoginResponse{Code: 200, Timestamp: time.Now(), Message: "登陆成功", Data: &LoginDataResponse{Token: "asdasdadadad"}}
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")
	error := json.NewEncoder(w).Encode(&loginResponse)
	if error != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func registry(w http.ResponseWriter, req *http.Request) {

}

func getNamespaces(w http.ResponseWriter, req *http.Request) {

}
