package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main.go/services"
	"main.go/utilities"
	"net/http"
)

func ConsultarInventario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio InventarioHandler::ConsultarInventario")
	var error = 0
	articulos, error := services.ConsultarInventario()
	if error == 1 {
		response.Meta.Status = utilities.StatusOk
		response.Data = articulos
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
	} else {
		response.Meta.Status = utilities.StatusFail
		response.Data = articulos
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
	}
	log.Printf("Termina InventarioHandler::ConsultarInventario")
}
