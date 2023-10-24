package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"main.go/utilities"

	"main.go/entities"
	"main.go/services"
)

var response entities.Response

func ConsultarEmpleados(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio PolizasHandler::ConsultarEmpleados")
	var error = 0
	empleados, error := services.ConsultarEmpleados()

	if error == 1 {
		response.Meta.Status = utilities.StatusOk
		response.Data = empleados
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
	} else {
		response.Meta.Status = utilities.StatusFail
		response.Data = empleados
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
	}
	log.Printf("Termina PolizasHandler::ConsultarEmpleados")

}

func AgregarEmpleado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio EmpleadosHandler::AgregarEmpleado")
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

	var empleado entities.Empleado
	err = json.Unmarshal(body, &empleado)
	if err != nil {
		log.Printf("error")
	}

	respuesta := services.AgregarEmpleado(empleado)

	var data = entities.PolizaMessage{}
	var mensaje = ""

	if respuesta == 1 {
		log.Printf("Exito al guardar al empleado " + empleado.Nombre + empleado.Apellido)
		response.Meta.Status = utilities.StatusOk
		data.Message = "Exito al guardar al empleado " + empleado.Nombre + " " + empleado.Apellido
		response.Data = data
		responseMethod, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseMethod))
		log.Printf("Termina PolizasHandler::AgregarPoliza")
		return
	} else {
		if respuesta == 0 {
			mensaje = "Error al guardar al empleado"
		} else if respuesta == 2 {
			mensaje = "El id de empleado a guardar ya existe"
		}
		log.Printf("Error al guardar al empleado " + empleado.Nombre + " " + empleado.Apellido)
		response.Meta.Status = utilities.StatusFail
		data.Message = mensaje
		response.Data = data
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		log.Printf("Termina EmpleadosHandler::AgregarEmpleados")

		return
	}
}

func ModificarEmpleado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio EmpleadosHandler::ModificarEmpleado")
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

	var empleado entities.Empleado
	err = json.Unmarshal(body, &empleado)
	if err != nil {
		log.Printf("error")
	}

	respuesta := services.ModificarEmpleado(empleado)

	var data = entities.PolizaMessage{}
	var mensaje = ""

	if respuesta == 1 {
		log.Printf("Exito al modificar al empleado " + empleado.Nombre + empleado.Apellido)
		response.Meta.Status = utilities.StatusOk
		data.Message = "Exito al modificar al empleado " + empleado.Nombre + " " + empleado.Apellido
		response.Data = data
		responseMethod, _ := json.Marshal(response)
		fmt.Fprintf(w, string(responseMethod))
		log.Printf("Termina PolizasHandler::AgregarPoliza")
		return
	} else {
		if respuesta == 0 {
			mensaje = "Error al modificar al empleado"
		} else if respuesta == 2 {
			mensaje = "El id de empleado a guardar ya existe"
		}
		log.Printf("Error al modificar al empleado " + empleado.Nombre + " " + empleado.Apellido)
		response.Meta.Status = utilities.StatusFail
		data.Message = mensaje
		response.Data = data
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		log.Printf("Termina EmpleadosHandler::ModificarEmpleado")

		return
	}
}

func EliminarEmpleado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}
	log.Printf("Inicio EmpleadosHandler::EliminarEmpleado")

	queryParams := r.URL.Query()

	id := queryParams.Get("idaEmpleado")
	idaEmpleado, _ := strconv.Atoi(id)

	result := services.EliminarEmpleado(idaEmpleado)

	var polizaMessage = entities.PolizaMessage{}

	if result != 1 {
		if result == 0 {
			log.Printf("Error al eliminar al empleado con id" + id)
			polizaMessage.Message = "Error al eliminar al empleado con id" + id
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
		polizaMessage.Message = "Se ha eliminado al empleado correctamente"
		response.Meta.Status = utilities.StatusOk
		response.Data = polizaMessage
		response, _ := json.Marshal(response)
		fmt.Fprintf(w, string(response))
		log.Printf("Inicio EmpleadosHandler::EliminarEmpleado")
	}
}
