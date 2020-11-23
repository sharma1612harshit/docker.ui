package docker

import (
	"io"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

// List local docker images
func GetImages(all bool, filters filters.Args) ([]types.ImageSummary, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     all,		// true
		Filters: filters, 	// {}
	})

	return images, err
}

// Pull specified docker image
func PullImage(imageRef string, all bool, registryAuth string) (io.ReadCloser, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	pullResponse, err := cli.ImagePull(ctx, imageRef, types.ImagePullOptions{
		All:           all,
		RegistryAuth:  registryAuth,
		PrivilegeFunc: nil,
	})

	return pullResponse, err
}

// Delete specified docker image
func DeleteImage(imageId string, force bool, pruneChildren bool) ([]types.ImageDelete, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	dResponse, err := cli.ImageRemove(ctx, imageId, types.ImageRemoveOptions{
		Force:         force,			// false
		PruneChildren: pruneChildren,	// false
	})

	return dResponse, err
}
