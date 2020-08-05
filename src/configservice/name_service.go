package configservice

import "time"

//命名空间
type Namespace struct {
	Appid       string
	Appsecret   string
	Appname     string
	Owner       string
	Contact     string
	Description string
	Create      time.Time
	Modify      time.Time
}

//元数据空间
type Metaspace struct {
	Appid string
	//元数据，json字符串
	Properties     string
	LastAccessTime time.Time
	LastModifyBy   string
	Create         time.Time
	Modify         time.Time
}

//用户空间
type Userspace struct {
	Uid      string
	Password string
	//角色：SuperStar、Admin、Leader、Follower
	Role string
	//readOnly、read/write
	Authority string
	Create    time.Time
	Modify    time.Time
}

//用户命名空间绑定关系
type Binding struct {
	Uid    string
	Appid  string
	Create time.Time
	Modify time.Time
}
