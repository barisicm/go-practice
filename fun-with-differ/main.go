package main

import (
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type user struct {
	ID          int
	Name        string
	StringArray []string
	IntArray    []int
}

func main() {
	oldUser := user{ID: 0, Name: "Stjepan", IntArray: []int{2, 3, 4, 5}, StringArray: []string{"burek", "ćevapi"}}
	newUser := user{ID: 0, Name: "Marin", IntArray: []int{2, 3, 4}, StringArray: []string{"burek", "ćevapii"}}

	_, _ = findStructDifferences(oldUser, newUser)
}

func findStructDifferences(old interface{}, new interface{}) (diff, error) {
	resDiff := diff{
		old: make(map[string]interface{}),
		new: make(map[string]interface{}),
	}

	err := checkDiffItemsValidity(old, new)
	if err != nil {
		return resDiff, err
	}

	//iterate over fields
	for fieldIndex := 0; fieldIndex < reflect.ValueOf(old).NumField(); fieldIndex++ {

		oldFieldValue := reflect.ValueOf(old).Field(fieldIndex)
		newFieldValue := reflect.ValueOf(new).Field(fieldIndex)
		fieldType := reflect.TypeOf(old).Field(fieldIndex).Type.Kind()
		fieldName := reflect.TypeOf(old).Field(fieldIndex).Name

		// check if field is an array
		if fieldType == reflect.Array || fieldType == reflect.Slice || fieldType == reflect.Map || fieldType == reflect.Struct || fieldType == reflect.Interface {
			if !reflect.DeepEqual(oldFieldValue.Interface(), newFieldValue.Interface()) {
				resDiff.old[fieldName] = oldFieldValue.Interface()
				resDiff.new[fieldName] = newFieldValue.Interface()
			}
			continue
		}

		if oldFieldValue.Interface() != newFieldValue.Interface() {
			resDiff.old[fieldName] = oldFieldValue.Interface()
			resDiff.new[fieldName] = newFieldValue.Interface()
		}

	}

	return resDiff, nil
}

func checkDiffItemsValidity(old interface{}, new interface{}) error {
	//check if diff items are structs
	if reflect.ValueOf(old).Kind() != reflect.Struct || reflect.ValueOf(new).Kind() != reflect.Struct {
		return status.Error(codes.Internal, "diff items are not structs")
	}
	//check if diff items types are same structs
	if reflect.TypeOf(old) != reflect.TypeOf(new) {
		return status.Error(codes.Internal, "diff items are not the same struct types")
	}

	return nil
}

type diff struct {
	old map[string]interface{}
	new map[string]interface{}
}
