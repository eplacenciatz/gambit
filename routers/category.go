package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/eplacenciatz/gambit/db"
	"github.com/eplacenciatz/gambit/models"
)

func InserCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}
	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la Categoría"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err := db.InsertCategory(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoría " + t.CategName + " > " + err.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + " }"
}

func UpdateCategory(body string, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar el Nombre (Title) y Path (Ruta) de la Categoría para actualizar"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID = id

	err = db.UpdateCategory(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la actualización de la categoría " + strconv.Itoa(t.CategID) + " > " + err.Error()
	}

	return 200, "Update OK"
}

func DeleteCategory(body string, User string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar ID de la Categoría a borrar"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := db.DeleteCategory(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la eliminación de la categoría " + strconv.Itoa(id) + " > " + err.Error()
	}
	return 200, "Delete OK"
}

func SelectCategories(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var CategId int
	var Slug string

	if len(request.QueryStringParameters["categId"]) > 0 {
		CategId, err = strconv.Atoi(request.QueryStringParameters["categId"])
		if err != nil {
			return 500, "Ocurrió un error al intentar convertir en entero el valor " + request.QueryStringParameters["categId"]
		}

	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	lista, err2 := db.SelectCategories(CategId, Slug)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar capturar Categoría/s > " + err2.Error()
	}

	Categ, err3 := json.Marshal(lista)
	if err3 != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON Categoría/s > " + err3.Error()
	}

	return 200, string(Categ)
}
