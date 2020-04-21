package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	VirtualPrivateCloudClient *network.VirtualPrivateCloudClient
	SubnetClient              *network.SubnetClient

	// Network ACL
	AclClient      *network.ACLClient
	AclRuleClient  *network.RuleClient
	InboundClient  *network.InboundRuleClient
	OutboundClient *network.OutboundRuleClient
}

func NewClient(o *common.ClientOptions) *Client {
	VirtualPrivateCloudClient := network.NewVirtualPrivateCloudClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&VirtualPrivateCloudClient.Client, o.ResourceManagerAuthorizer)

	SubnetClient := network.NewSubnetClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&SubnetClient.Client, o.ResourceManagerAuthorizer)

	AclClient := network.NewACLClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&AclClient.Client, o.ResourceManagerAuthorizer)

	AclRuleClient := network.NewRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&AclRuleClient.Client, o.ResourceManagerAuthorizer)

	InboundClient := network.NewInboundRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&InboundClient.Client, o.ResourceManagerAuthorizer)

	OutboundClient := network.NewOutboundRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&OutboundClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		VirtualPrivateCloudClient: &VirtualPrivateCloudClient,
		SubnetClient:              &SubnetClient,

		// Network ACL
		AclClient:      &AclClient,
		AclRuleClient:  &AclRuleClient,
		InboundClient:  &InboundClient,
		OutboundClient: &OutboundClient,
	}
}