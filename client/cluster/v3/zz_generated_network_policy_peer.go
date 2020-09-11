package client

const (
	NetworkPolicyPeerType                   = "networkPolicyPeer"
	NetworkPolicyPeerFieldIPBlock           = "ipBlock"
	NetworkPolicyPeerFieldNamespaceSelector = "namespaceSelector"
	NetworkPolicyPeerFieldPodSelector       = "podSelector"
	NetworkPolicyPeerFieldProjectID         = "projectId"
	NetworkPolicyPeerFieldTargetWorkloadId  = "targetWorkloadId"
)

type NetworkPolicyPeer struct {
	IPBlock           *IPBlock       `json:"ipBlock,omitempty" yaml:"ipBlock,omitempty"`
	NamespaceSelector *LabelSelector `json:"namespaceSelector,omitempty" yaml:"namespaceSelector,omitempty"`
	PodSelector       *LabelSelector `json:"podSelector,omitempty" yaml:"podSelector,omitempty"`
	ProjectID         string         `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	TargetWorkloadId  string         `json:"targetWorkloadId,omitempty" yaml:"targetWorkloadId,omitempty"`
}
