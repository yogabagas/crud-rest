package data

import (
	"database/sql"
	"day14gaji/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ItemRepositori struct {
	DB *sql.DB
}

func GetAll(db *ItemRepositori) []models.Item {
	fmt.Println(db.DB)
	result, err := db.DB.Query("Select tblItemID, ItemCode, ItemName From Item")
	if err != nil {
		return nil
	}
	defer result.Close()
	fmt.Println(result)
	item := []models.Item{}
	for result.Next() {
		var i models.Item
		if err := result.Scan(&i.tblItemID, &i.ItemCode, &i.ItemName); err != nil {
			return nil
		}
		item = append(item, i)
	}
	return item
}
