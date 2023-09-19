package routers

import (
	"encoding/json"
	"strconv"

	"github.com/eplacenciatz/gambit/db"
	"github.com/eplacenciatz/gambit/models"
)

func InsertProduct(body string, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el Nombre (Title) del Producto"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := db.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ ProductId: " + strconv.Itoa(int(result)) + " }"
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

	err = db.UpdateProduct(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la actualización de producto " + strconv.Itoa(t.ProdId) + " > " + err.Error()
	}

	return 200, "Update OK"
}
