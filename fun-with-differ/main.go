package main

import (
	"fmt"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	oldUser := user{ID: 0, Name: "Stjepan", IntArray: []int{2, 3, 4, 5}, StringArray: []string{"burek", "ćevapi"}}
	newUser := user{ID: 0, Name: "Marin", IntArray: []int{2, 3, 4}, StringArray: []string{"burek", "ćevapii"}}

	diffStuff, _ := doMeADiff(oldUser, newUser)

	fmt.Println(diffStuff.old["ID"])
	fmt.Println(diffStuff.old["Name"])
	fmt.Println(diffStuff.old["IntArray"])
	fmt.Println(diffStuff.old["StringArray"])

	fmt.Println(diffStuff.new["ID"])
	fmt.Println(diffStuff.new["Name"])
	fmt.Println(diffStuff.new["IntArray"])
	fmt.Println(diffStuff.new["StringArray"])
}

func doMeADiff(old interface{}, new interface{}) (diff, error) {
	resDiff := diff{
		old: make(map[string]interface{}),
		new: make(map[string]interface{}),
	}

	err := checkDiffItemValidity(old, new)
	if err != nil {
		return resDiff, err
	}

	//iterate over struct fields
	for fieldIndex := 0; fieldIndex < reflect.ValueOf(old).NumField(); fieldIndex++ {

		oldFieldValue := reflect.ValueOf(old).Field(fieldIndex)
		newFieldValue := reflect.ValueOf(new).Field(fieldIndex)
		fieldType := reflect.TypeOf(old).Field(fieldIndex).Type.Kind()
		fieldName := reflect.TypeOf(old).Field(fieldIndex).Name

		// check if field is an array
		if fieldType == reflect.Array || fieldType == reflect.Slice || fieldType == reflect.Map || fieldType == reflect.Struct || fieldType == reflect.Interface {
			if !reflect.DeepEqual(oldFieldValue.Interface(), newFieldValue.Interface()) {
				//write changes into old and new map
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

func checkDiffItemValidity(old interface{}, new interface{}) error {
	//check if diff items are structs
	if reflect.ValueOf(old).Kind() != reflect.Struct || reflect.ValueOf(new).Kind() != reflect.Struct {
		return status.Error(codes.Internal, "diff items are not structs")
	}
	//check if diff items types are same structs
	if reflect.TypeOf(old) != reflect.TypeOf(new) {
		return status.Error(codes.Internal, "diff items are not the same structs")
	}
	//check if objects have the same number of fields
	if reflect.ValueOf(old).NumField() != reflect.ValueOf(new).NumField() {
		return status.Error(codes.Internal, "diff items have different amount of fields")
	}

	return nil
}

type user struct {
	ID          int      ``
	Name        string   ``
	IntArray    []int    ``
	StringArray []string ``
}

type diff struct {
	old map[string]interface{}
	new map[string]interface{}
}
