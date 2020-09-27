package client

const (
	IPBlockType        = "ipBlock"
	IPBlockFieldCIDR   = "cidr"
	IPBlockFieldExcept = "except"
)

type IPBlock struct {
	CIDR   string   `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	Except []string `json:"except,omitempty" yaml:"except,omitempty"`
}
