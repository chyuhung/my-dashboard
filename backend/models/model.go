package models

type User struct {
	Username string `label:"用户名" json:"username"`
	Password string `label:"密码" json:"password"`
}

// CreateServerRequest 结构体表示创建服务器的请求数据
// 其他必要信息动态获取
type CreateServerRequest struct {
	OS       string    `json:"os"`
	Flavor   string    `json:"flavor"`
	Volumes  []Volume  `json:"volumes"`
	Hostname string    `json:"hostname"`
	Networks []Network `json:"networks"`
}

// Volume 结构体表示创建服务器时的卷信息
// 其他必要信息动态获取
type Volume struct {
	Size int `json:"size"`
}

// Network 结构体表示网络信息
// 其他必要信息动态获取
type Network struct {
	Vlan string `json:"vlan"`
	IP   string `json:"ip_address"`
}
