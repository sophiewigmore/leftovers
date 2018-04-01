package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevieve/leftovers/aws/common"
)

type imagesClient interface {
	DescribeImages(*awsec2.DescribeImagesInput) (*awsec2.DescribeImagesOutput, error)
	DeregisterImage(*awsec2.DeregisterImageInput) (*awsec2.DeregisterImageOutput, error)
}

type Images struct {
	client       imagesClient
	logger       logger
	resourceTags resourceTags
}

func NewImages(client imagesClient, logger logger, resourceTags resourceTags) Images {
	return Images{
		client:       client,
		logger:       logger,
		resourceTags: resourceTags,
	}
}

func (i Images) ListOnly(filter string) ([]common.Deletable, error) {
	return i.get(filter, true)
}

func (i Images) List(filter string) ([]common.Deletable, error) {
	return i.get(filter, true)
}

func (i Images) get(filter string, prompt bool) ([]common.Deletable, error) {
	images, err := i.client.DescribeImages(&awsec2.DescribeImagesInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing EC2 Images: %s", err)
	}

	var resources []common.Deletable
	for _, image := range images.Images {
		r := NewImage(i.client, image.ImageId, i.resourceTags)

		if !strings.Contains(r.Name(), filter) {
			continue
		}

		if prompt {
			proceed := i.logger.PromptWithDetails(r.Type(), r.Name())
			if !proceed {
				continue
			}
		}

		resources = append(resources, r)
	}

	return resources, nil
}
