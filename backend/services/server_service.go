package services

import (
	"errors"
	"log"

	"github.com/chyuhung/my-dashboard/config"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

var computeClient *gophercloud.ServiceClient

func init() {
	var err error
	computeClient, err = config.GetComputeClient()
	if err != nil {
		log.Fatalf("Failed to create compute client: %v", err)
	}
}

func GetImage(name string) (images.Image, error) {
	opts := images.ListOpts{
		Limit: 99999,
	}
	allPages, err := images.ListDetail(computeClient, opts).AllPages()
	if err != nil {
		return images.Image{}, errors.New("failed to get all images")
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return images.Image{}, errors.New("failed to extract images")
	}

	for _, image := range allImages {
		if image.Name == name {
			return image, nil
		}
	}
	return images.Image{}, errors.New("image not found")
}

func GetFlavor(name string) (flavors.Flavor, error) {
	opts := flavors.ListOpts{
		Limit: 99999,
	}
	allPages, err := flavors.ListDetail(computeClient, opts).AllPages()
	if err != nil {
		return flavors.Flavor{}, errors.New("failed to get all flavors")
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return flavors.Flavor{}, errors.New("failed to extract flavors")
	}

	for _, flavor := range allFlavors {
		if flavor.Name == name {
			return flavor, nil
		}
	}
	return flavors.Flavor{}, errors.New("flavor not found")
}

func GetServers() ([]servers.Server, error) {
	opts := servers.ListOpts{
		Limit: 99999,
	}
	allPages, err := servers.List(computeClient, opts).AllPages()
	if err != nil {
		return []servers.Server{}, errors.New("failed to get all servers")
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		return []servers.Server{}, errors.New("failed to extract servers")
	}
	return allServers, nil
}

func GetFlavors() ([]flavors.Flavor, error) {
	opts := flavors.ListOpts{
		Limit: 99999,
	}
	allPages, err := flavors.ListDetail(computeClient, opts).AllPages()
	if err != nil {
		return []flavors.Flavor{}, errors.New("failed to get all flavors")
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return []flavors.Flavor{}, errors.New("failed to extract flavors")
	}
	return allFlavors, nil
}

func GetImages() ([]images.Image, error) {
	opts := images.ListOpts{
		Limit: 99999,
	}
	allPages, err := images.ListDetail(computeClient, opts).AllPages()
	if err != nil {
		return []images.Image{}, errors.New("failed to get all images")
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return []images.Image{}, errors.New("failed to extract images")
	}
	return allImages, nil
}
