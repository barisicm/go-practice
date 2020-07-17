package fun_with_adapter

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "./db/testDb.db")
	defer db.Close()

	return
}
