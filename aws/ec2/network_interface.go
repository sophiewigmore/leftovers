package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type NetworkInterface struct {
	client     networkInterfacesClient
	id         *string
	identifier string
}

func NewNetworkInterface(client networkInterfacesClient, id *string, tags []*awsec2.Tag) NetworkInterface {
	identifier := *id

	extra := []string{}
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		identifier = fmt.Sprintf("%s (%s)", *id, strings.Join(extra, ", "))
	}

	return NetworkInterface{
		client:     client,
		id:         id,
		identifier: identifier,
	}
}

func (n NetworkInterface) Delete() error {
	_, err := n.client.DeleteNetworkInterface(&awsec2.DeleteNetworkInterfaceInput{
		NetworkInterfaceId: n.id,
	})
	return err
}
