package db

import (
	"database/sql"
	"fmt"

	"github.com/erick-mondragon/gambit/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Category > Ejecución Éxitosa")
	return LastInsertId, err2
}
