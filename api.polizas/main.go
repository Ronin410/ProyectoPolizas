package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"main.go/entities"

	"github.com/gorilla/mux"
	"main.go/handlers"
)

func main() {
	file, err := os.OpenFile("log-api-polizas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("%s Fatal: Fatal Error Signal")
	}
	defer file.Close()
	log.SetOutput(file)

	log.Printf("Inicio del api.polizas")

	mux := mux.NewRouter()
	mux.Use(CORS)
	//mux.Use(Token)
	path := os.Getenv("urlpath")
	log.Printf(path)
	path_api := "/api.polizas"

	mux.HandleFunc(path_api+"/Empleados", handlers.ConsultarEmpleados).Methods("GET", "OPTIONS")
	mux.HandleFunc(path_api+"/ConsultarPolizasEmpleado", handlers.ConsultarPolizasEmpleado).Methods("GET", "OPTIONS")
	mux.HandleFunc(path_api+"/AgregarPoliza", handlers.AgregarPoliza).Methods("POST", "OPTIONS")
	mux.HandleFunc(path_api+"/EliminarPolizas", handlers.EliminarPoliza).Methods("POST", "OPTIONS")
	mux.HandleFunc(path_api+"/ConsultarInventario", handlers.ConsultarInventario).Methods("GET", "OPTIONS")
	mux.HandleFunc(path_api+"/ActualizarPoliza", handlers.ActualizarPoliza).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":3000", mux))

}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		(w).Header().Set("Content-Type", "application/json")
		(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		(w).Header().Set("Access-Control-Allow-Methods", "*")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		(w).Header().Set("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)

		return
	})
}

func Token(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}
		log.Printf("Inicio Main::Token")

		var responseToken entities.Response

		var urlsso = os.Getenv("urlsso")
		method := "POST"
		payload := strings.NewReader(``)

		req, err := http.NewRequest(method, urlsso, payload)
		if err != nil {
			log.Printf("Error al formar request", err)
		}

		req.Header.Add("Authorization", r.Header.Get("Authorization"))

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Printf("Error al ejecutar el request", err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error al leer el cuerpo de la respuesta:", err)
		}

		var response entities.ResponseSSO
		json.Unmarshal(body, &response)

		responseToken.Meta.Status = response.Meta.Status
		responseToken.Data = response.Meta.Error

		status, _ := json.Marshal(responseToken)

		if responseToken.Meta.Status == "FAIL" || responseToken.Meta.Status == "" {
			log.Printf("El estatus de la peticion fue", status)
			return
		}
		next.ServeHTTP(w, r)
		log.Printf("Termina Main::Token")
		return
	})
}
