package client

const (
	NetworkPolicySpecType             = "networkPolicySpec"
	NetworkPolicySpecFieldEgress      = "egress"
	NetworkPolicySpecFieldIngress     = "ingress"
	NetworkPolicySpecFieldPodSelector = "podSelector"
	NetworkPolicySpecFieldPolicyTypes = "policyTypes"
)

type NetworkPolicySpec struct {
	Egress      []NetworkPolicyEgressRule  `json:"egress,omitempty" yaml:"egress,omitempty"`
	Ingress     []NetworkPolicyIngressRule `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	PodSelector *LabelSelector             `json:"podSelector,omitempty" yaml:"podSelector,omitempty"`
	PolicyTypes []string                   `json:"policyTypes,omitempty" yaml:"policyTypes,omitempty"`
}
