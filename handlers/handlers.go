package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/erick-mondragon/gambit/auth"
	"github.com/erick-mondragon/gambit/routers"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validoAuthorization(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return procesoUsers(body, path, method, user, id, request)
	case "prod":
		return procesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return procesoStock(body, path, method, user, idn, request)
	case "addr":
		return procesoAddress(body, path, method, user, idn, request)
	case "cate":
		return procesoCategory(body, path, method, user, idn, request)
	case "orde":
		return procesoOrders(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid"
}

func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		path == "category" && method == "GET" {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido"
	}

	todoOk, err, msg := auth.ValidoToken(token)
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg
}

func procesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	if path == "user/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)
		}
	}
	if path == "users" {
		if method == "GET" {
			return routers.SelectUsers(body, user, request)
		}
	}

	return 400, "Method Invalid"
}

func procesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "PUT":
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		return routers.DeleteProduct(user, id)
	case "GET":
		return routers.SelectProduct(request)
	}
	return 400, "Method Invalid"
}

func procesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(body, user, id)
	case "GET":
		return routers.SelectCategories(body, request)
	}
	return 400, "Method Invalid"
}

func procesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateStock(body, user, id)
}

func procesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertAddress(body, user)
	case "PUT":
		return routers.UpdateAddress(body, user, id)
	case "DELETE":
		return routers.DeleteAddress(user, id)
	case "GET":
		return routers.SelectAddress(user)
	}
	return 400, "Method Invalid"
}

func procesoOrders(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertOrder(body, user)
	case "GET":
		return routers.SelectOrders(user, request)
	}
	return 400, "Method Invalid"
}
