package models

import "gorm.io/gorm"

type Instance struct {
	gorm.Model
	Name     string    `label:"虚拟机名称" json:"name"`
	Flavor   string    `label:"规格名称" json:"flavor"`
	Image    string    `laebl:"镜像" json:"image"`
	Volumes  []Volume  `label:"卷" json:"volumes"`
	Networks []Network `label:"网络" json:"networks"`
	Host     string    `label:"宿主机名称" json:"host"`
}

type User struct {
	gorm.Model
	Username string `label:"用户名" json:"username"`
	Password string `label:"密码" json:"password"`
}
type Network struct {
	Vlan string `label:"VLAN" json:"vlan"`
	Ip   string `label:"地址" json:"ip"`
}

type Volume struct {
	Name string `label:"名称" json:"name"`
	Size int    `label:"大小" json:"size"`
	Type string `label:"类型" json:"type"`
}
