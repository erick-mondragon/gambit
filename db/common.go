package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/erick-mondragon/gambit/models"
	"github.com/erick-mondragon/gambit/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la base de datos")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndPoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true&parseTime=true",
		dbUser, authToken, dbEndPoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza UserIsAdmin...")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserIsAdmin > Ejecución éxitosa - valor devuelto " + valor)
	if valor == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}
