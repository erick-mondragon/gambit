package routers

import (
	"encoding/json"

	"github.com/erick-mondragon/gambit/db"
	"github.com/erick-mondragon/gambit/models"
)

func InsertAddress(body string, User string) (int, string) {
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Debe especificar el Address "
	}
	if t.AddName == "" {
		return 400, "Debe especificar el Name "
	}
	if t.AddTitle == "" {
		return 400, "Debe especificar el Title "
	}
	if t.AddCity == "" {
		return 400, "Debe especificar el City "
	}
	if t.AddPhone == "" {
		return 400, "Debe especificar el Phone "
	}
	if t.AddPostalCode == "" {
		return 400, "Debe especificar el PostalCode "
	}

	err = db.InsertAddress(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del Address para el ID de usuario " + User + " > " + err.Error()
	}

	return 200, "Insert Address"
}

func UpdateAddress(body string, User string, id int) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	t.AddId = id
	var encontrado bool
	err, encontrado = db.AddressExists(User, t.AddId)
	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar Adress para el usuario " + User + " > " + err.Error()
		}
		return 400, "No se encuentra un registro de ID de Usuario asociado a ese ID de Address"
	}

	err = db.UpdateAddress(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la actualización del Address para el ID de usuario " + User + " > " + err.Error()
	}

	return 200, "UpdateAddress OK"
}
