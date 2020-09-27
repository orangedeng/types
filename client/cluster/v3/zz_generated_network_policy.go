package client

import (
	"github.com/rancher/norman/types"
)

const (
	NetworkPolicyType                  = "networkPolicy"
	NetworkPolicyFieldAnnotations      = "annotations"
	NetworkPolicyFieldCreated          = "created"
	NetworkPolicyFieldCreatorID        = "creatorId"
	NetworkPolicyFieldEgress           = "egress"
	NetworkPolicyFieldIngress          = "ingress"
	NetworkPolicyFieldLabels           = "labels"
	NetworkPolicyFieldName             = "name"
	NetworkPolicyFieldNamespaceId      = "namespaceId"
	NetworkPolicyFieldOwnerReferences  = "ownerReferences"
	NetworkPolicyFieldPodSelector      = "podSelector"
	NetworkPolicyFieldPolicyTypes      = "policyTypes"
	NetworkPolicyFieldRemoved          = "removed"
	NetworkPolicyFieldTargetWorkloadId = "targetWorkloadId"
	NetworkPolicyFieldUUID             = "uuid"
)

type NetworkPolicy struct {
	types.Resource
	Annotations      map[string]string          `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created          string                     `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID        string                     `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Egress           []NetworkPolicyEgressRule  `json:"egress,omitempty" yaml:"egress,omitempty"`
	Ingress          []NetworkPolicyIngressRule `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	Labels           map[string]string          `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name             string                     `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId      string                     `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences  []OwnerReference           `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	PodSelector      *LabelSelector             `json:"podSelector,omitempty" yaml:"podSelector,omitempty"`
	PolicyTypes      string                     `json:"policyTypes,omitempty" yaml:"policyTypes,omitempty"`
	Removed          string                     `json:"removed,omitempty" yaml:"removed,omitempty"`
	TargetWorkloadId string                     `json:"targetWorkloadId,omitempty" yaml:"targetWorkloadId,omitempty"`
	UUID             string                     `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type NetworkPolicyCollection struct {
	types.Collection
	Data   []NetworkPolicy `json:"data,omitempty"`
	client *NetworkPolicyClient
}

type NetworkPolicyClient struct {
	apiClient *Client
}

type NetworkPolicyOperations interface {
	List(opts *types.ListOpts) (*NetworkPolicyCollection, error)
	ListAll(opts *types.ListOpts) (*NetworkPolicyCollection, error)
	Create(opts *NetworkPolicy) (*NetworkPolicy, error)
	Update(existing *NetworkPolicy, updates interface{}) (*NetworkPolicy, error)
	Replace(existing *NetworkPolicy) (*NetworkPolicy, error)
	ByID(id string) (*NetworkPolicy, error)
	Delete(container *NetworkPolicy) error
}

func newNetworkPolicyClient(apiClient *Client) *NetworkPolicyClient {
	return &NetworkPolicyClient{
		apiClient: apiClient,
	}
}

func (c *NetworkPolicyClient) Create(container *NetworkPolicy) (*NetworkPolicy, error) {
	resp := &NetworkPolicy{}
	err := c.apiClient.Ops.DoCreate(NetworkPolicyType, container, resp)
	return resp, err
}

func (c *NetworkPolicyClient) Update(existing *NetworkPolicy, updates interface{}) (*NetworkPolicy, error) {
	resp := &NetworkPolicy{}
	err := c.apiClient.Ops.DoUpdate(NetworkPolicyType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkPolicyClient) Replace(obj *NetworkPolicy) (*NetworkPolicy, error) {
	resp := &NetworkPolicy{}
	err := c.apiClient.Ops.DoReplace(NetworkPolicyType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *NetworkPolicyClient) List(opts *types.ListOpts) (*NetworkPolicyCollection, error) {
	resp := &NetworkPolicyCollection{}
	err := c.apiClient.Ops.DoList(NetworkPolicyType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *NetworkPolicyClient) ListAll(opts *types.ListOpts) (*NetworkPolicyCollection, error) {
	resp := &NetworkPolicyCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *NetworkPolicyCollection) Next() (*NetworkPolicyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NetworkPolicyCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NetworkPolicyClient) ByID(id string) (*NetworkPolicy, error) {
	resp := &NetworkPolicy{}
	err := c.apiClient.Ops.DoByID(NetworkPolicyType, id, resp)
	return resp, err
}

func (c *NetworkPolicyClient) Delete(container *NetworkPolicy) error {
	return c.apiClient.Ops.DoResourceDelete(NetworkPolicyType, &container.Resource)
}
