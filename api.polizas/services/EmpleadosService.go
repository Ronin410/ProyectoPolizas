package services

import (
	"database/sql"
	"log"
	"os"

	"main.go/entities"
)

type Empleado []entities.Empleado

var urlPostgress = os.Getenv("urlPostgress")

func ConsultarEmpleados() (entities.Empleados, int) {
	// Conectarse a la base de datos
	empleados := entities.Empleados{}

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos", err)
		return empleados, 0
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT id_emp,nombre_emp,apellido_emp,puesto_emp FROM fun_consultaempleados()")
	if err != nil {
		log.Printf("Error al consultar la funcion fun_consultaempleados", err)
		return empleados, 0
	}
	defer rows.Close()

	// Recorrer los resultados
	for rows.Next() {
		empleado := entities.Empleado{}
		err := rows.Scan(&empleado.Idempleado, &empleado.Nombre, &empleado.Apellido, &empleado.Puesto)
		empleados = append(empleados, empleado)

		if err != nil {
			log.Printf("Error al leer la respuesta de la funcion fun_consultaempleados", err)
			return empleados, 0
		}
	}

	return empleados, 1
}
