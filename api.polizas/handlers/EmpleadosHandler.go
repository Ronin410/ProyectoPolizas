package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main.go/utilities"
	"net/http"

	"main.go/entities"
	"main.go/services"
)

var response entities.Response

func ConsultarEmpleados(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio PolizasHandler::ConsultarEmpleados")

	empleados := services.ConsultarEmpleados()

	response.Meta.Status = utilities.StatusOk
	response.Data = empleados
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Termina PolizasHandler::ConsultarEmpleados")

}