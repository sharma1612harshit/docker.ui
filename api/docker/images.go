package docker

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/filters"
	duiTypes "github.com/sharma1612harshit/docker.ui/api/types"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"github.com/sharma1612harshit/gomuf/logs"
)

// GetImages - return images data as json map
func GetImages(all, filter string) ([]duiTypes.ImageResponse, error) {
	allImages := false
	searchFilters := map[string]map[string]bool{}

	if all == "true" {
		allImages = true
	}

	_ = json.Unmarshal([]byte(filter), &searchFilters)

	// add filters to filter var
	var searchFilter = filters.NewArgs()
	
	for key, value := range searchFilters {
		if key == "label" || key == "created" {
			for name, boolean := range value {
				searchFilter.Add(key, fmt.Sprintf("{\"%s\":%v}",name, boolean))
			}
		} else {
			logs.Info("Invalid filter passed: " + key)
		}
	}

	images, err := docker.GetImages(allImages, searchFilter)

	var imageList = make([]duiTypes.ImageResponse, 0)

	if err != nil {
		logs.Warn(err)
		return imageList, err
	}

	for _, data := range images {
		imageList = append(imageList, duiTypes.ImageResponse{
			ID:         data.ID,
			Created:    data.Created,
			Containers: data.Containers,
			Labels:     data.Labels,
			RepoDigest: data.RepoDigests,
			RepoTags:   data.RepoTags,
			Size:       data.Size,
			ParentId:   data.ParentID,
		})
	}

	return imageList, err
}

// PullImage - pull an image from registry
func PullImage(all, reference, username, password string) (string, error) {
	allImages := false

	if all == "true" {
		allImages = true
	}

	err := docker.PullImage(reference, allImages, username, password)

	if err != nil {
		logs.Warn(err)
		return fmt.Sprintf("Error Downloading: %s", reference), err
	}

	return fmt.Sprintf("Downloaded: %s", reference), err
}
