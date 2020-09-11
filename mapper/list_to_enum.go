package mapper

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/mapper"
	"github.com/rancher/norman/types/values"
)

type ListToEnum struct {
	Field   string
	Options []string
}

func (e ListToEnum) FromInternal(data map[string]interface{}) {
	contents, ok := values.GetStringSlice(data, e.Field)
	if !ok {
		return
	}
	if len(contents) == 0 {
		values.RemoveValue(data, e.Field)
		return
	}
	sort.Strings(contents)
	if reflect.DeepEqual(contents, e.Options) {
		values.PutValue(data, "Both", e.Field)
		return
	}
	values.PutValue(data, contents[0], e.Field)
}

func (e ListToEnum) ToInternal(data map[string]interface{}) error {
	option, ok := values.GetValue(data, e.Field)
	if !ok {
		return fmt.Errorf("field %s is required", e.Field)
	}

	if option == "Both" {
		values.PutValue(data, e.Options, e.Field)
	} else {
		values.PutValue(data, []string{option.(string)}, e.Field)
	}

	return nil
}

func (e ListToEnum) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	field := schema.ResourceFields[e.Field]
	if !strings.HasPrefix(field.Type, "array[") {
		return fmt.Errorf("field %s of schema %s is not an array", e.Field, schema.ID)
	}
	sort.Strings(e.Options)
	return mapper.Enum{Field: e.Field, Options: append(e.Options, "Both")}.ModifySchema(schema, schemas)
}
