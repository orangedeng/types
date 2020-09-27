package mapper

import (
	"strings"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/values"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type IDSelector struct {
	Key        string
	ID         string
	Selector   string
	TrimPrefix bool
}

func (o IDSelector) FromInternal(data map[string]interface{}) {

}

func (o IDSelector) ToInternal(data map[string]interface{}) error {
	v, ok := data[o.ID]
	if !ok {
		return nil
	}
	var id string
	if o.TrimPrefix {
		index := strings.LastIndexByte(v.(string), ':')
		if index < 0 {
			id = v.(string)
		} else {
			id = v.(string)[index+1:]
		}
	} else {
		id = strings.ReplaceAll(v.(string), ":", "-")
	}
	selector := metav1.SetAsLabelSelector(labels.Set(map[string]string{
		o.Key: id,
	}))
	values.PutValue(data, selector, o.Selector)
	return nil
}

func (o IDSelector) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
