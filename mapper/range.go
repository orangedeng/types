package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/values"
)

type Range struct {
	List   string
	Mapper types.Mapper
}

func (r Range) FromInternal(data map[string]interface{}) {
	list, ok := values.GetSlice(data, r.List)
	if !ok {
		return
	}
	for _, item := range list {
		if r.Mapper != nil {
			r.Mapper.FromInternal(item)
		}
	}
}

func (r Range) ToInternal(data map[string]interface{}) error {
	list, ok := values.GetSlice(data, r.List)
	if !ok {
		return nil
	}
	for _, item := range list {
		if r.Mapper != nil {
			if err := r.Mapper.ToInternal(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r Range) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
