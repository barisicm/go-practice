package main

import (
	"fmt"
	"math"
)

func main() {
	guitarManufacturers := []string{"Gibson", "Fender", "Ibanez", "Cort", "Yamaha", "Music Man", "Dean Martin"}
	perPage := 2
	page := 4
	pages := math.Ceil(float64(len(guitarManufacturers)) / float64(perPage))

	if page > int(pages) {
		fmt.Println("This is not gud")
	}

	page--

	skip := perPage * page
	end := skip * perPage

	if end > len(guitarManufacturers) {
		end = len(guitarManufacturers)
	}

	guitarManufacturers = append(guitarManufacturers[skip:end])

	fmt.Println(guitarManufacturers)
	fmt.Println("Pages : ", pages)
}

func paginate(page int, perPage int, pages int, array []string) ([]string, error) {
	var err error

	if page > int(pages) {
		return nil, err
	}

	page--

	skip := perPage * page
	end := skip * perPage

	if end > len(array) {
		end = len(array)
	}

	array = append(array[skip:end])

	return array, nil
}
