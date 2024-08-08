package services

import (
	"errors"
	"log"

	"github.com/chyuhung/my-dashboard/config"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
)

var volumeClientV2 *gophercloud.ServiceClient

func init() {
	var err error
	volumeClientV2, err = config.GetVolumeClientV2()
	if err != nil {
		log.Fatalf("Failed to create volume client: %v", err)
	}
}

func GetVolume(name string) (volumes.Volume, error) {
	opts := volumes.ListOpts{Name: name}
	allPages, err := volumes.List(volumeClientV2, opts).AllPages()
	if err != nil {
		return volumes.Volume{}, errors.New("failed to list volumes")
	}

	allVolumes, err := volumes.ExtractVolumes(allPages)
	if err != nil {
		return volumes.Volume{}, errors.New("failed to extract volumes from pages")
	}

	if len(allVolumes) == 0 {
		return volumes.Volume{}, errors.New("no volume found with the given name")
	}

	if len(allVolumes) > 1 {
		return volumes.Volume{}, errors.New("multiple volumes found with the same name")
	}

	return allVolumes[0], nil
}

func CreateVolume(opts volumes.CreateOpts) (*volumes.Volume, error) {
	volume, err := volumes.Create(volumeClientV2, opts).Extract()
	if err != nil {
		return &volumes.Volume{}, errors.New("failed to create volume")
	}
	return volume, nil
}
