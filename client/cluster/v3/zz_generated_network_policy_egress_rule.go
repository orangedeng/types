package client

const (
	NetworkPolicyEgressRuleType       = "networkPolicyEgressRule"
	NetworkPolicyEgressRuleFieldPorts = "ports"
	NetworkPolicyEgressRuleFieldTo    = "to"
)

type NetworkPolicyEgressRule struct {
	Ports []NetworkPolicyPort `json:"ports,omitempty" yaml:"ports,omitempty"`
	To    []NetworkPolicyPeer `json:"to,omitempty" yaml:"to,omitempty"`
}
