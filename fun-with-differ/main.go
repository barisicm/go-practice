package main

import (
	"fmt"
	"reflect"
)

func main() {
	oldUser := user{ID: 0, Name: "Stjepan"}
	newUser := user{ID: 0, Name: "Marin"}

	_, _ = doMeADiff(oldUser, newUser)
}

func doMeADiff(old interface{}, new interface{}) (diff, error) {
	//check if diff items are structs
	if reflect.ValueOf(old).Kind() != reflect.Struct || reflect.ValueOf(new).Kind() != reflect.Struct {
		fmt.Println("everything went to shit")
		fmt.Println("diff items are not structs, go away")
	}
	//check if diff items contain same structs
	if reflect.TypeOf(old) != reflect.TypeOf(new) {
		fmt.Println("everything went to shit")
		fmt.Println("diff items are not the same structs, go away")
	}
	//check if objects have the same number of fields
	if reflect.ValueOf(old).NumField() != reflect.ValueOf(new).NumField() {
		fmt.Println("everything went to shit")
		fmt.Println("diff items have different amount of fields, go away")
	}

	//iterate over struct fields
	for fieldIndex := 0; fieldIndex < reflect.ValueOf(old).NumField(); fieldIndex++ {
		//is this check really neccesary ?
		if reflect.TypeOf(reflect.ValueOf(old).Field(fieldIndex)) != reflect.TypeOf(reflect.ValueOf(new).Field(fieldIndex)) {
			fmt.Println("everything went to shit")
			fmt.Printf("\n field type at element %v, is not the same in the old and new struct \n", fieldIndex)
		}

		//TODO: Check for unexported fields ??
		if reflect.ValueOf(old).Field(fieldIndex).Interface() != reflect.ValueOf(new).Field(fieldIndex).Interface() {
			oldFieldVal := reflect.ValueOf(old).Field(fieldIndex).Interface()
			newFieldVal := reflect.ValueOf(new).Field(fieldIndex).Interface()
			fmt.Printf("\nvalues are not the same old: %v new: %v", oldFieldVal, newFieldVal)
		}

	}

	//check if values are base values(int, string, float, bool)
	//

	//check if values are structs

	//check if values are arrays

	resDiff := diff{
		old: make(map[string]interface{}),
		new: make(map[string]interface{}),
	}

	return resDiff, nil
}

func checkDiffItemValidity(old interface{}, new interface{}) error {
	return nil
}

type user struct {
	ID   int    ``
	Name string ``
}

type diff struct {
	old map[string]interface{}
	new map[string]interface{}
}
