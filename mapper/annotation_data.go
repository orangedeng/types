package mapper

import (
	"encoding/json"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
	"github.com/sirupsen/logrus"
)

const prefix = "store.cattle.io/"

type AnnotationData struct {
	Field string
	List  bool
}

func (o AnnotationData) FromInternal(data map[string]interface{}) {
	value, ok := values.RemoveValue(data, "annotations", prefix+o.Field)
	if !ok {
		return
	}
	var rtn interface{}
	if o.List {
		tmpRTN := []map[string]interface{}{}
		if err := json.Unmarshal([]byte(value.(string)), &tmpRTN); err != nil {
			logrus.Errorf("failed to decode content %s, %v", value, err)
		}
		rtn = tmpRTN
	} else {
		tmpRTN := map[string]interface{}{}
		if err := json.Unmarshal([]byte(value.(string)), &tmpRTN); err != nil {
			logrus.Errorf("failed to decode content %s, %v", value, err)
		}
		rtn = tmpRTN
	}
	values.PutValue(data, rtn, o.Field)
}

func (o AnnotationData) ToInternal(data map[string]interface{}) error {
	targetValue, ok := values.GetValue(data, o.Field)
	if !ok {
		return nil
	}
	var annotationMap map[string]interface{}
	anno, ok := values.GetValue(data, "annotations")
	if !ok {
		annotationMap = map[string]interface{}{}
		values.PutValue(data, annotationMap, "annotations")
	} else {
		annotationMap = convert.ToMapInterface(anno)
	}

	rtnValue, err := json.Marshal(targetValue)
	if err != nil {
		return err
	}
	annotationMap[prefix+o.Field] = string(rtnValue)

	return nil
}

func (o AnnotationData) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
