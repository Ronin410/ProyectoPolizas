package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"main.go/entities"
	"main.go/services"
	"main.go/utilities"
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

func AgregarInventario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio InventarioHandler::AgregarInventario")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	bodyString := string(body)
	fmt.Println("Valores del cuerpo:", bodyString)

	var gson, _ = json.Marshal(body)
	fmt.Println("Valores del cuerpo:", gson)

	var articulo entities.Articulo
	err = json.Unmarshal(body, &articulo)
	if err != nil {
		log.Printf("error")
	}

	respuesta := services.AgregarInventario(articulo)

	var data = entities.PolizaMessage{}
	var mensaje = ""

	if respuesta == 1 {
		log.Printf("Exito al guardar el articulo " + articulo.NombreArticulo)
		response.Meta.Status = utilities.StatusOk
		data.Message = "Exito al guardar el articulo " + articulo.NombreArticulo
		response.Data = data
		responseMethod, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseMethod))
		log.Printf("Termina PolizasHandler::AgregarPoliza")
		return
	} else {
		if respuesta == 0 {
			mensaje = "Error al guardar el articulo"
		} else if respuesta == 2 {
			mensaje = "El id de articulo a guardar ya existe"
		}
		log.Printf("Error al guardar el articulo " + articulo.NombreArticulo)
		response.Meta.Status = utilities.StatusFail
		data.Message = mensaje
		response.Data = data
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		log.Printf("Termina InventarioHandler::AgregarInventario")

		return
	}
}

func ModificarInventario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio InventarioHandler::ModificarInventario")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	bodyString := string(body)
	fmt.Println("Valores del cuerpo:", bodyString)

	var gson, _ = json.Marshal(body)
	fmt.Println("Valores del cuerpo:", gson)

	var articulo entities.Articulo
	err = json.Unmarshal(body, &articulo)
	if err != nil {
		log.Printf("error")
	}

	respuesta := services.ModificarInventario(articulo)

	var data = entities.PolizaMessage{}
	var mensaje = ""

	if respuesta == 1 {
		log.Printf("Exito al modificar el articulo ")
		response.Meta.Status = utilities.StatusOk
		data.Message = "Exito al modificar el articulo "
		response.Data = data
		responseMethod, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseMethod))
		log.Printf("Termina PolizasHandler::ModificarInventario")
		return
	} else {
		if respuesta == 0 {
			mensaje = "Error al modificar el articulo "
		} else if respuesta == 2 {
			mensaje = "El id de articulo a modificar ya existe"
		}
		log.Printf("Error al modificar el articulo ")
		response.Meta.Status = utilities.StatusFail
		data.Message = mensaje
		response.Data = data
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		log.Printf("Termina InventarioHandler::ModificarInventario")

		return
	}
}

func EliminarInventario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio InventarioHandler::EliminarInventario")

	queryParams := r.URL.Query()

	id := queryParams.Get("idarticulo")
	idarticulo, _ := strconv.Atoi(id)

	result := services.EliminarArticulo(idarticulo)

	var polizaMessage = entities.PolizaMessage{}

	if result != 1 {
		if result == 0 {
			log.Printf("Error al eliminar el articulo con id" + id)
			polizaMessage.Message = "Error al eliminar el articulo con id" + id
		} else if result == 2 {
			log.Printf("Error el id a eliminar no existe")
			polizaMessage.Message = "Error el id a eliminar no existe"
		}

		response.Meta.Status = utilities.StatusFail
		response.Data = polizaMessage
		response, _ := json.Marshal(response)

		fmt.Fprintf(w, string(response))
		return

	} else if result == 1 {
		polizaMessage.Message = "Se ha eliminado el articulo correctamente"
		response.Meta.Status = utilities.StatusOk
		response.Data = polizaMessage
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		log.Printf("Inicio InventarioHandler::EliminarInventario")
	}

}
