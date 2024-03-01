package config

import (
	"log"

	"github.com/chyuhung/my-dashboard/services"
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
	// 读取环境变量
	viper.AutomaticEnv()

	Username = viper.GetString("OS_USERNAME")
	Password = viper.GetString("OS_PASSWORD")
	ProjectName = viper.GetString("OS_PROJECT_NAME")
	DomainName = viper.GetString("OS_USER_DOMAIN_NAME")
	AuthURL = viper.GetString("OS_AUTH_URL")
	// 默认值为"ReginOne"
	viper.SetDefault("OS_REGION_NAME", "RegionOne")
	Region = viper.GetString("OS_REGION_NAME")

	// 环境变量值检查
	if Username == "" || Password == "" || ProjectName == "" || DomainName == "" || AuthURL == "" || Region == "" {
		// 缺少必要环境变量值
		log.Fatal("Missing required environment variables")
		return
	}
	log.Println("OpenStack configuration loading successful")
	log.Printf("Environment variables: Username:%s Password:%s ProjectName:%s Domain:%s AuthURL:%s Region:%s", Username, Password, ProjectName, DomainName, AuthURL, Region)

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: AuthURL,
		Username:         Username,
		Password:         Password,
		DomainName:       DomainName,
		AllowReauth:      true,
		Scope: &gophercloud.AuthScope{
			ProjectName: ProjectName,
			DomainName:  DomainName,
		},
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		log.Println("Failed to create OpenStack client:", err)
		return
	}

	// computeClient
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: Region,
		Name:   "nova",
		Type:   "compute",
	})
	if err != nil {
		log.Println("Failed to create compute service client:", err)
		return
	}
	// 将computeClient传递给需要使用的服务函数
	services.SetComputeClient(computeClient)

	// volumeClient
	volumeClient, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
		Region: Region,
		Name:   "cinderv2",
		Type:   "volumev2",
	})
	if err != nil {
		log.Println("Failed to create volume service client:", err)
		return
	}
	services.SetVolumeClient(volumeClient)

	// networkClient
	networkClient, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: Region,
		Name:   "neutron",
		Type:   "network",
	})
	if err != nil {
		log.Println("Failed to create network service client:", err)
		return
	}
	services.SetNetworkClient(networkClient)

}
