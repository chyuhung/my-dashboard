package models

// type Instance struct {
// 	Name     string    `label:"虚拟机名称" json:"name"`
// 	Flavor   string    `label:"规格名称" json:"flavor"`
// 	Volumes  []Volume  `label:"卷" json:"volumes"`
// 	Networks []Network `label:"网络" json:"networks"`
// 	Host     string    `label:"宿主机名称" json:"host"`
// }
// type Network struct {
// 	Vlan string `label:"VLAN" json:"vlan"`
// 	Ip   string `label:"地址" json:"ip"`
// }

// type Volume struct {
// 	Name string `label:"名称" json:"name"`
// 	Size int    `label:"大小" json:"size"`
// 	Type string `label:"类型" json:"type"`
// 	Image
// }
// type Image struct {
// 	Name string `label:"名称" json:"name"`
// 	Id   string `label:"id" json:"id"`
// }

type User struct {
	Username string `label:"用户名" json:"username"`
	Password string `label:"密码" json:"password"`
}
