package main

import (
	"reflect"
	"testing"
)

func Test_findStructDifferences(t *testing.T) {
	type car struct {
		ID    int
		Model string
	}
	type user struct {
		ID               int
		Name             string
		FavouriteFood    []string
		FavouriteNumbers []int
		Car              car
		Map              map[string]string
		Unknown          interface{}
	}

	test1UserOld := user{ID: 0, Name: "Stjepan", FavouriteNumbers: []int{2, 3, 4, 5}, FavouriteFood: []string{"burek", "ćevapi"}}
	test1UserNew := user{ID: 0, Name: "Marin", FavouriteNumbers: []int{2, 3, 4}, FavouriteFood: []string{"burek", "ćevapii"}}
	test1DiffRes := diff{
		old: map[string]interface{}{
			"Name":             "Stjepan",
			"FavouriteNumbers": []int{2, 3, 4, 5},
			"FavouriteFood":    []string{"burek", "ćevapi"},
		},
		new: map[string]interface{}{
			"Name":             "Marin",
			"FavouriteNumbers": []int{2, 3, 4},
			"FavouriteFood":    []string{"burek", "ćevapii"},
		},
	}

	test2UserOld := user{ID: 0, Name: "Stjepan", FavouriteNumbers: []int{2, 3, 4, 5}, FavouriteFood: []string{"burek", "ćevapi"}}
	test2UserNew := car{ID: 12, Model: "Mazda"}
	test2DiffRes := diff{
		old: make(map[string]interface{}),
		new: make(map[string]interface{}),
	}

	test3UserOld := user{
		FavouriteFood: []string{"Burek"},
		Car: car{
			ID:    0,
			Model: "Mazda",
		},
		Unknown: 34,
		Map: map[string]string{
			"case": "test",
		},
	}
	test3UserNew := user{
		FavouriteFood: []string{"Burek"},
		Car: car{
			ID:    0,
			Model: "Mazda",
		},
		Unknown: 34,
		Map: map[string]string{
			"case": "test2",
		},
	}
	test3DiffRes := diff{
		old: map[string]interface{}{
			"Map": map[string]string{
				"case": "test",
			},
		},
		new: map[string]interface{}{
			"Map": map[string]string{
				"case": "test2",
			},
		},
	}

	test4UserOld := user{
		FavouriteFood: []string{"Burekk"},
		Car: car{
			ID:    1,
			Model: "Mazda",
		},
		Unknown: 34,
		Map: map[string]string{
			"case": "test",
		},
	}
	test4UserNew := user{
		FavouriteFood: []string{"Burek"},
		Car: car{
			ID:    0,
			Model: "Mazda",
		},
		Unknown: "whatever",
		Map: map[string]string{
			"case": "test2",
		},
	}
	test4DiffRes := diff{
		old: map[string]interface{}{
			"FavouriteFood": []string{"Burekk"},
			"Car": car{
				ID:    1,
				Model: "Mazda",
			},
			"Unknown": 34,
			"Map": map[string]string{
				"case": "test",
			},
		},
		new: map[string]interface{}{
			"FavouriteFood": []string{"Burek"},
			"Car": car{
				ID:    0,
				Model: "Mazda",
			},
			"Unknown": "whatever",
			"Map": map[string]string{
				"case": "test2",
			},
		},
	}

	type args struct {
		old interface{}
		new interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    diff
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Correct structs with simple fields",
			args: args{
				old: test1UserOld,
				new: test1UserNew,
			},
			want:    test1DiffRes,
			wantErr: false,
		},
		{
			name: "Incorrect structs",
			args: args{
				old: test2UserOld,
				new: test2UserNew,
			},
			want:    test2DiffRes,
			wantErr: true,
		},
		{
			name: "Correct structs with complex fields",
			args: args{
				old: test3UserOld,
				new: test3UserNew,
			},
			want:    test3DiffRes,
			wantErr: false,
		},
		{
			name: "Correct structs with complex fields2",
			args: args{
				old: test4UserOld,
				new: test4UserNew,
			},
			want:    test4DiffRes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findStructDifferences(tt.args.old, tt.args.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("findStructDifferences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStructDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDiffItemsValidity(t *testing.T) {
	type user struct {
		ID               int
		Name             string
		FavouriteFood    []string
		FavouriteNumbers []int
	}

	type car struct {
		ID    int
		Model string
	}

	test1UserOld := user{ID: 0, Name: "Stjepan", FavouriteNumbers: []int{2, 3, 4, 5}, FavouriteFood: []string{"burek", "ćevapi"}}
	test1UserNew := user{ID: 0, Name: "Marin", FavouriteNumbers: []int{2, 3, 4}, FavouriteFood: []string{"burek", "ćevapii"}}

	test2UserOld := 32
	test2UserNew := user{ID: 0, Name: "Marin", FavouriteNumbers: []int{2, 3, 4}, FavouriteFood: []string{"burek", "ćevapii"}}

	test3UserOld := user{ID: 0, Name: "Marin", FavouriteNumbers: []int{2, 3, 4}, FavouriteFood: []string{"burek", "ćevapii"}}
	test3UserNew := "A string"

	test4UserOld := []string{"sample1", "sample2"}
	test4UserNew := []int{23, 14, 15}

	test5UserOld := user{ID: 0, Name: "Marin", FavouriteNumbers: []int{2, 3, 4}, FavouriteFood: []string{"burek", "ćevapii"}}
	test5CarNew := car{ID: 13, Model: "Subaru"}

	type args struct {
		old interface{}
		new interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Both args are valid",
			args: args{
				old: test1UserOld,
				new: test1UserNew,
			},
			wantErr: false,
		},
		{
			name: "First arg is invalid",
			args: args{
				old: test2UserOld,
				new: test2UserNew,
			},
			wantErr: true,
		},
		{
			name: "Second arg is invalid",
			args: args{
				old: test3UserOld,
				new: test3UserNew,
			},
			wantErr: true,
		},
		{
			name: "Both args are invalid",
			args: args{
				old: test4UserOld,
				new: test4UserNew,
			},
			wantErr: true,
		},

		{
			name: "Args are valid interfaces but different structs",
			args: args{
				old: test5UserOld,
				new: test5CarNew,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkDiffItemsValidity(tt.args.old, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("checkDiffItemValidity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
