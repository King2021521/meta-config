package portal

import (
	"configservice"
	"crypto/md5"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"storage"
	"time"
)

var mgo = storage.Mgo{"mongodb://localhost:27017", "meta-config", "user_space"}

//登陆业务处理
func Login(loginReq LoginRequest) CommonResponse {
	mongoTemplate := storage.NewMongoTemplate(&mgo)
	result := mongoTemplate.Query(bson.D{
		{"uid", loginReq.UserName},
		{"password", loginReq.Password}})

	var loginResponse CommonResponse
	if result == nil {
		loginResponse = CommonResponse{Code: http.StatusBadRequest, Timestamp: time.Now(), Message: "登陆失败，用户名或密码错误"}
	} else {
		var user configservice.Userspace
		result.Decode(&user)
		loginResponse = CommonResponse{Code: http.StatusOK, Timestamp: time.Now(), Message: "登陆成功", Data: &LoginDataResponse{Token: GetToken(user.Uid)}}
	}
	return loginResponse
}

//新增用户业务处理

func Registry(user configservice.Userspace) CommonResponse {
	user.Create = time.Now()
	user.Modify = time.Now()
	mongoTemplate := storage.NewMongoTemplate(&mgo)
	result, _ := mongoTemplate.Insert(user)
	return CommonResponse{Code: http.StatusOK, Timestamp: time.Now(), Message: "注册成功", Data: result}
}

func GetToken(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str
}
