package services

import (
	"errors"
	"log"

	"github.com/chyuhung/my-dashboard/config"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
)

var networkClient *gophercloud.ServiceClient

func init() {
	var err error
	networkClient, err = config.GetNetworkClient()
	if err != nil {
		log.Fatalf("Failed to create compute client: %v", err)
	}
}

func GetNetwork(name string) (networks.Network, error) {
	opts := networks.ListOpts{Name: name}
	allPages, err := networks.List(networkClient, opts).AllPages()
	if err != nil {
		return networks.Network{}, errors.New("failed to list networks")
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return networks.Network{}, errors.New("failed to extract networks from pages")
	}

	if len(allNetworks) == 0 {
		return networks.Network{}, errors.New("no network found with the given name")
	}

	if len(allNetworks) > 1 {
		return networks.Network{}, errors.New("multiple networks found with the same name")
	}

	return allNetworks[0], nil
}
func GetNetworks() ([]networks.Network, error) {
	opts := networks.ListOpts{
		Limit: 99999,
	}
	allPages, err := networks.List(networkClient, opts).AllPages()
	if err != nil {
		return []networks.Network{}, err
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return []networks.Network{}, errors.New("failed to extract networks")
	}
	return allNetworks, nil
}
