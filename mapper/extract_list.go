package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/values"
)

type ExtractListField struct {
	List   string
	Fields []string
	ToList string
}

func (o ExtractListField) FromInternal(data map[string]interface{}) {
	toList, ok := values.GetSlice(data, o.ToList)
	if !ok {
		return
	}
	fromList, ok := values.GetSlice(data, o.List)
	if !ok {
		return
	}
	if len(fromList) != len(toList) {
		return
	}
	for i := 0; i < len(fromList); i++ {
		toData := toList[i]
		fromData := fromList[i]
		for _, field := range o.Fields {
			if toData[field] == nil {
				continue
			}
			fromData[field] = toData[field]
		}
	}
}

func (o ExtractListField) ToInternal(data map[string]interface{}) error {
	array, ok := values.GetSlice(data, o.List)
	if !ok {
		return nil
	}

	if len(array) == 0 {
		return nil
	}

	newList := make([]map[string]interface{}, len(array))
	for i, v := range array {
		listContent := map[string]interface{}{}
		for _, field := range o.Fields {
			inner, _ := values.GetValue(v, field)
			listContent[field] = inner
		}
		newList[i] = listContent
	}

	values.PutValue(data, newList, o.ToList)
	return nil
}

func (o ExtractListField) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
