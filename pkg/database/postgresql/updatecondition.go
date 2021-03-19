package postgresql

import (
	"reflect"
	"strings"
)

// UpdateConditionFromStruct
// Creates condition string (valueOfTag = :valueOfTag) for updating record in database.
func UpdateConditionFromStruct(i interface{}) string {
	var updateConditionString string
	valuesOfStruct := reflect.ValueOf(i).Elem()
	typeOfStruct := reflect.TypeOf(i).Elem()
	for i := 0; i < valuesOfStruct.NumField(); i++ {
		f1 := valuesOfStruct.Field(i)
		if f1.IsZero() == false {
			t := reflect.TypeOf(i).Elem()
			fieldOfStruct, _ := t.FieldByName(typeOfStruct.Field(i).Name)
			valueOfTag, _ := fieldOfStruct.Tag.Lookup("db")
			updateConditionString += valueOfTag + " = :" + valueOfTag + ", "
		}
	}
	updateConditionString = strings.TrimSuffix(updateConditionString, ", ")
	return updateConditionString
}
