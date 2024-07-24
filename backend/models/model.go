package models

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

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

type CustomVolumeCreateOpts struct {
	BootIndex  int `label:"启动顺序" json:"bootindex"`
	CreateOpts volumes.CreateOpts
}

// CustomCreateOpts 结构体包含创建服务器和创建卷的选项
type CustomCreateOpts struct {
	ServerCreateOpts servers.CreateOpts       // 服务器创建选项
	VolumeCreateOpts []CustomVolumeCreateOpts // 卷创建选项列表
}
