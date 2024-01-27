package strcgo

import "reflect"

// GetTagValueFromField gets tag value by tag name from struct field name
func GetTagValueFromField(s interface{}, fieldName string, tagName string) string {
	field, found := reflect.TypeOf(s).FieldByName(fieldName)
	if !found {
		return ""
	}

	return field.Tag.Get(tagName)
}
