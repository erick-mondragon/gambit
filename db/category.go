package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/erick-mondragon/gambit/models"
	"github.com/erick-mondragon/gambit/tools"
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

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza registro UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE category SET "
	if len(c.CategName) > 0 {
		sentencia += "Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}
		sentencia += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución éxitosa")
	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Comienza proceso DeleteCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecución éxitosa")
	return nil
}

func SelectCategories(CategId int, Slug string) ([]models.Category, error) {
	fmt.Println("Comienza proceso SelectCategories")

	var Categ []models.Category

	err := DbConnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()

	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "

	if CategId > 0 {
		sentencia += "WHERE Categ_Id = " + strconv.Itoa(CategId)
	} else {
		if len(Slug) > 0 {
			sentencia += "WHERE Categ_Path LIKE '%" + Slug + "%'"
		}
	}

	fmt.Println(sentencia)

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return Categ, err
	}

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Ejecución éxitosa")
	return Categ, nil
}
