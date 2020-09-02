package main

import (
	"fmt"
	"os"

	"github.com/markbates/pkger"
)

func main() {

	// connection := "host=localhost port=7654 user=miyagi dbname=pkger_fun password=miyagi sslmode=disable"

	// db, err := gorm.Open("postgres", connection)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()

	// _, err = db.DB().Exec("CREATE SCHEMA IF NOT EXISTS pkger_fun")
	// if err != nil {
	// }

	// db.DB().Exec("SET search_path TO pkger_fun,public")

	// instance, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	// if err != nil {
	// }

	// get bundeled resources with pkger
	// bindata or should i say pkger package and do migrations
	// s2 := pkger.Include("migrations")
	// fmt.Println(s2)
	var names []string
	pkger.Walk("/migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		names = append(names, info.Name())

		// f, _ := pkger.Open(info.Name())
		// f.Read()

		return nil
	})

	fmt.Println(names)

	// d, err := bindata.WithInstance(s)
	// if err != nil {
	// }

	// migration, err := migrate.NewWithInstance("go-bindata", d, "postgres", instance)
	// if err != nil {
	// }

	// err = migration.Up()
	// if err != nil {

	// }
}
