package client

const (
	NetworkPolicyIngressRuleType       = "networkPolicyIngressRule"
	NetworkPolicyIngressRuleFieldFrom  = "from"
	NetworkPolicyIngressRuleFieldPorts = "ports"
)

type NetworkPolicyIngressRule struct {
	From  []NetworkPolicyPeer `json:"from,omitempty" yaml:"from,omitempty"`
	Ports []NetworkPolicyPort `json:"ports,omitempty" yaml:"ports,omitempty"`
}
