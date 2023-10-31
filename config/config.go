package config

import (
	"os"
)

// OpenStackConfig 包含OpenStack认证参数和其他相关配置
type OpenStackConfig struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	ProjectName string `json:"project_name"`
	DomainName  string `json:"domain_name"`

	AuthURL string `json:"auth_url"`
	Region  string `json:"region"  default:"ReginOne"`
}

// LoadConfig 从环境变量中加载OpenStack配置
func LoadConfig() OpenStackConfig {
	config := OpenStackConfig{
		Username:    os.Getenv("OS_USERNAME"),
		Password:    os.Getenv("OS_PASSWORD"),
		ProjectName: os.Getenv("OS_PROJECT_NAME"),
		DomainName:  os.Getenv("OS_USER_DOMAIN_NAME"),
		AuthURL:     os.Getenv("OS_AUTH_URL"),
		Region:      os.Getenv("OS_REGION_NAME"),
	}

	return config
}
