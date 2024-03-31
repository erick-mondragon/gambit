package routers

import (
	"encoding/json"
	"strconv"

	"github.com/erick-mondragon/gambit/db"
	"github.com/erick-mondragon/gambit/models"
)

func InsertProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el nombre (Title) del producto"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := db.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateProduct(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := db.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el UPDATE del producto " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK"

}

func DeleteProduct(User string, id int) (int, string) {

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := db.DeleteProduct(id)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el DELETE del producto " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Delete OK"
}
