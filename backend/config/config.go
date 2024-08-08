package config

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/spf13/viper"
)

// Config 结构体用于存储 OpenStack 配置
type Config struct {
	Username    string
	Password    string
	ProjectName string
	DomainName  string
	AuthURL     string
	Region      string
}

var config Config

// loadConfig 加载环境变量配置
func loadConfig() {
	viper.AutomaticEnv()

	config.Username = viper.GetString("OS_USERNAME")
	config.Password = viper.GetString("OS_PASSWORD")
	config.ProjectName = viper.GetString("OS_PROJECT_NAME")
	config.DomainName = viper.GetString("OS_USER_DOMAIN_NAME")
	config.AuthURL = viper.GetString("OS_AUTH_URL")

	viper.SetDefault("OS_REGION_NAME", "RegionOne")
	config.Region = viper.GetString("OS_REGION_NAME")

	if config.Username == "" || config.Password == "" || config.ProjectName == "" || config.DomainName == "" || config.AuthURL == "" || config.Region == "" {
		log.Fatal("Missing required environment variables")
	}
}

// getProviderClient 获取已认证的 ProviderClient
func getProviderClient() (*gophercloud.ProviderClient, error) {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: config.AuthURL,
		Username:         config.Username,
		Password:         config.Password,
		DomainName:       config.DomainName,
		AllowReauth:      true,
		Scope: &gophercloud.AuthScope{
			ProjectName: config.ProjectName,
			DomainName:  config.DomainName,
		},
	}

	return openstack.AuthenticatedClient(authOpts)
}

// createComputeClient 创建 Compute 客户端
func createComputeClient() (*gophercloud.ServiceClient, error) {
	provider, err := getProviderClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: config.Region,
	})
}

// createVolumeClientV1 创建 Volume V1 客户端
func createVolumeClientV1() (*gophercloud.ServiceClient, error) {
	provider, err := getProviderClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
		Region: config.Region,
	})
}

// createVolumeClientV2 创建 Volume V2 客户端
func createVolumeClientV2() (*gophercloud.ServiceClient, error) {
	provider, err := getProviderClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: config.Region,
	})
}

// createVolumeClientV3 创建 Volume V3 客户端
func createVolumeClientV3() (*gophercloud.ServiceClient, error) {
	provider, err := getProviderClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: config.Region,
	})
}

// createNetworkClient 创建 Network 客户端
func createNetworkClient() (*gophercloud.ServiceClient, error) {
	provider, err := getProviderClient()
	if err != nil {
		return nil, err
	}

	return openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: config.Region,
	})
}

// Init 初始化 OpenStack 客户端
func Init() {
	loadConfig()
}

// GetComputeClient 获取 Compute 客户端
func GetComputeClient() (*gophercloud.ServiceClient, error) {
	return createComputeClient()
}

// GetVolumeClientV1 获取 Volume V1 客户端
func GetVolumeClientV1() (*gophercloud.ServiceClient, error) {
	return createVolumeClientV1()
}

// GetVolumeClientV2 获取 Volume V2 客户端
func GetVolumeClientV2() (*gophercloud.ServiceClient, error) {
	return createVolumeClientV2()
}

// GetVolumeClientV3 获取 Volume V3 客户端
func GetVolumeClientV3() (*gophercloud.ServiceClient, error) {
	return createVolumeClientV3()
}

// GetNetworkClient 获取 Network 客户端
func GetNetworkClient() (*gophercloud.ServiceClient, error) {
	return createNetworkClient()
}
