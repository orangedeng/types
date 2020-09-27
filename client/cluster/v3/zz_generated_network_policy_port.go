package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	NetworkPolicyPortType          = "networkPolicyPort"
	NetworkPolicyPortFieldPort     = "port"
	NetworkPolicyPortFieldProtocol = "protocol"
)

type NetworkPolicyPort struct {
	Port     intstr.IntOrString `json:"port,omitempty" yaml:"port,omitempty"`
	Protocol string             `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
