package models

import "gorm.io/gorm"

type Instance struct {
	gorm.Model
	ImageName      string            `label:"镜像名称" json:"image_name"`
	FlavorName     string            `label:"规格名称" json:"flavor_name"`
	Volumes        map[string]int    `label:"卷" json:"volumes"`
	VmName         string            `label:"虚拟机名称" json:"vm_name"`
	Networks       map[string]string `label:"网络" json:"networks"`
	HostName       string            `label:"宿主机名称" json:"host_name"`
	VolumeTypeName string            `label:"卷类型名称" json:"volume_type_name"`
}

type User struct {
	gorm.Model
	Username string `lable:"用户名" json:"username"`
	Password string `lable:"密码" json:"password"`
}
