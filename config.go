package main

import (
	"os"
)

// OpenStackConfig 包含OpenStack认证参数和其他相关配置
type OpenStackConfig struct {
	AuthURL   string
	Username  string
	Password  string
	ProjectID string
	Region    string
}

// LoadConfig 从环境变量中加载OpenStack配置
func LoadConfig() OpenStackConfig {
	config := OpenStackConfig{
		AuthURL:   os.Getenv("OS_AUTH_URL"),
		Username:  os.Getenv("OS_USERNAME"),
		Password:  os.Getenv("OS_PASSWORD"),
		ProjectID: os.Getenv("OS_PROJECT_ID"),
		Region:    os.Getenv("OS_REGION_NAME"),
	}

	return config
}
