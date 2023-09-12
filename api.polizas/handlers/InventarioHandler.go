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

	articulos := services.ConsultarInventario()
	response.Meta.Status = utilities.StatusOk
	response.Data = articulos
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Termina InventarioHandler::ConsultarInventario")
}
