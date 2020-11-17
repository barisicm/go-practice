package main

import (
	"fmt"

	mk "github.com/DATA-DOG/go-sqlmock"
)

func main() {
	db, mock, err := mk.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	rows := mk.NewRows([]string{"id", "title"}).
		AddRow(1, "one").
		AddRow(2, "two")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	rs, _ := db.Query("SELECT")
	defer rs.Close()

	for rs.Next() {
		var id int
		var title string
		rs.Scan(&id, &title)
		fmt.Println("scanned id:", id, "and title:", title)
	}

	if rs.Err() != nil {
		fmt.Println("got rows error:", rs.Err())
	}
}
