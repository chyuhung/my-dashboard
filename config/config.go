package config

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/spf13/viper"
)

var (
	Username    string
	Password    string
	ProjectName string
	DomainName  string

	AuthURL string
	Region  string
)

func Init() {
	viper.AutomaticEnv()

	Username = viper.GetString("OS_USERNAME")
	Password = viper.GetString("OS_PASSWORD")
	ProjectName = viper.GetString("OS_PROJECT_NAME")
	DomainName = viper.GetString("OS_USER_DOMAIN_NAME")
	AuthURL = viper.GetString("OS_AUTH_URL")
	Region = viper.GetString("OS_REGION_NAME")

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: AuthURL,
		Username:         Username,
		Password:         Password,
		DomainName:       DomainName,
		AllowReauth:      true,
		Scope: &gophercloud.AuthScope{
			ProjectName: ProjectName,
			DomainName:  DomainName},
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		log.Println("Failed to create OpenStack client:", err)
		return
	}

	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: Region,
		Name:   "nova",
		Type:   "compute",
	})
	if err != nil {
		// 错误处理
	}

	// 将computeClient传递给需要使用的服务函数
}
