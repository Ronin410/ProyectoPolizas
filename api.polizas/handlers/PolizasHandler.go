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

func ConsultarPolizasEmpleado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio PolizasHandler::ConsultarPolizasEmpleado")
	queryParams := r.URL.Query()
	var dataMessage = entities.PolizaMessage{}
	id := queryParams.Get("idempleado")
	idempleado, _ := strconv.Atoi(id)
	var mensaje = ""
	polizas, mensaje := services.ConsultarPolizasEmpleado(int32(idempleado))
	if mensaje == utilities.StatusOk {
		response.Meta.Status = utilities.StatusOk
		response.Data = polizas
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
	} else {
		dataMessage.Message = mensaje
		response.Meta.Status = utilities.StatusFail
		response.Data = polizas
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
	}

	log.Printf("Termina PolizasHandler::ConsultarPolizasEmpleado")

}

func AgregarPoliza(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio PolizasHandler::AgregarPoliza")
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

	var poliza entities.PolizaEntrada
	err = json.Unmarshal(body, &poliza)
	if err != nil {
		log.Printf("eror")
	}

	respuesta := services.AgregarPoliza(poliza)

	var data = entities.PolizaMessage{}
	var mensaje = ""

	if respuesta == 1 {
		log.Printf("Exito al guardar la poliza del cliente " + poliza.NombreCliente)
		response.Meta.Status = utilities.StatusOk
		data.Message = "Exito al guardar la poliza"
		response.Data = data
		responseMethod, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseMethod))
		log.Printf("Termina PolizasHandler::AgregarPoliza")
		return
	} else {
		if respuesta == 0 {
			mensaje = "Error al guardar la poliza"
		} else if respuesta == 2 {
			mensaje = "El empleado a guardar no existe"
		} else if respuesta == 3 {
			mensaje = "El sku a registrar no tiene inventario disponible"
		} else if respuesta == 4 {
			mensaje = "El sku a registrar no existe"
		}
		log.Printf("Error al guardar la poliza del cliente " + poliza.NombreCliente)
		response.Meta.Status = utilities.StatusFail
		data.Message = mensaje
		response.Data = data
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		return
	}
}

func EliminarPoliza(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio PolizasHandler::EliminarPoliza")

	queryParams := r.URL.Query()

	id := queryParams.Get("idpoliza")
	idpoliza, _ := strconv.Atoi(id)

	result := services.EliminarPoliza(idpoliza)

	var polizaMessage = entities.PolizaMessage{}

	if result == 0 {
		log.Printf("Error al eliminar la poliza con id" + id)
		polizaMessage.Message = utilities.ErrorEliminarPoliza
		response.Meta.Status = utilities.StatusFail
		response.Data = polizaMessage
		response, _ := json.Marshal(response)

		fmt.Fprintf(w, string(response))
		return
	}
	polizaMessage.Message = "Se ha eliminado la poliza correctamente"
	response.Meta.Status = utilities.StatusOk
	response.Data = polizaMessage
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Inicio PolizasHandler::EliminarPoliza")
}

func ActualizarPoliza(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio PolizasHandler::ActualizarPoliza")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	bodyString := string(body)

	fmt.Println("Valores del cuerpo:", bodyString)

	var poliza entities.Poliza
	err = json.Unmarshal(body, &poliza)
	if err != nil {
		log.Printf("error: ", err)
	}

	respuesta := services.ActualizarPoliza(poliza)

	var data = entities.PolizaMessage{}
	var mensaje = ""

	if respuesta == 1 {
		log.Printf("Exito al guardar la poliza del cliente " + poliza.NombreCliente)
		response.Meta.Status = utilities.StatusOk
		data.Message = "Exito al guardar la poliza"
		response.Data = data
		responseMethod, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseMethod))
		log.Printf("Termina PolizasHandler::AgregarPoliza")
		return
	} else {
		if respuesta == 0 {
			mensaje = "Error al guardar la poliza"
		} else if respuesta == 2 {
			mensaje = "El sku a registrar no tiene inventario disponible"
		} else if respuesta == 3 {
			mensaje = "El sku a registrar no existe"
		} else if respuesta == 4 {
			mensaje = "El empleado a guardar no existe"
		} else if respuesta == 5 {
			mensaje = "La poliza a actualizar no existe"
		}
		log.Printf("Error al actualizar la poliza " + string(poliza.IdPoliza))
		response.Meta.Status = utilities.StatusFail
		data.Message = mensaje
		response.Data = data
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		return
	}

	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Inicio PolizasHandler::ActualizarPoliza")
}
