package services

import (
	"errors"
	"log"

	"github.com/chyuhung/my-dashboard/config"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumetypes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
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

// GetVolumeTypes 获取可用的卷类型
func GetVolumeTypes() ([]volumetypes.VolumeType, error) {
	allPages, err := volumetypes.List(volumeClientV2).AllPages()
	if err != nil {
		return nil, errors.New("failed to list volume types")
	}
	allTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		return nil, errors.New("failed to extract volume types from pages")
	}
	return allTypes, nil
}

func GetVolumes() ([]volumes.Volume, error) {
	opts := &volumes.ListOpts{
		AllTenants: true,
		Limit:      99999,
	}

	allPages, err := volumes.List(volumeClientV2, opts).AllPages()
	if err != nil {

		return nil, errors.New("failed to list volumes")

	}
	allVolumes, err := volumes.ExtractVolumes(allPages)
	if err != nil {

		return nil, errors.New("failed to extract volumes from pages")

	}
	return allVolumes, nil
}

func SearchVolumeTypes(name string) ([]volumetypes.VolumeType, error) {
	vts, err := GetVolumeTypes()
	if err != nil {
		return nil, err
	}
	var filteredVolumeTypes []volumetypes.VolumeType
	for _, vt := range vts {
		if vt.Name == name {
			filteredVolumeTypes = append(filteredVolumeTypes, vt)
		}
	}
	return filteredVolumeTypes, nil
}
