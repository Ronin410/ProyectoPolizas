package services

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"main.go/entities"
)

type Empleado []entities.Empleado

var urlPostgress = os.Getenv("urlPostgress")

func ConsultarEmpleados() entities.Empleados {
	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT id_emp,nombre_emp,apellido_emp,puesto_emp FROM fun_consultaempleados()")
	if err != nil {
		log.Printf("Error al conectarse a la base de datos", err)
		if err, ok := err.(*pq.Error); ok {
			log.Printf("", err.Message)
		} //log.Fatal(err)
	}
	defer rows.Close()
	empleados := entities.Empleados{}

	// Recorrer los resultados
	for rows.Next() {
		empleado := entities.Empleado{}
		err := rows.Scan(&empleado.Idempleado, &empleado.Nombre, &empleado.Apellido, &empleado.Puesto)
		empleados = append(empleados, empleado)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return empleados
}
