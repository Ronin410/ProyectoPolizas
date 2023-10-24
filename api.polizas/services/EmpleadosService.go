package services

import (
	"database/sql"
	"fmt"
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

func AgregarEmpleado(empleado entities.Empleado) int32 {
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_agregarempleado($1::INTEGER, $2::varchar(50), $3::varchar(50), $4::varchar(50))")
	if err != nil {
		log.Printf("Error al preparar el statement a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(empleado.Idempleado, empleado.Nombre, empleado.Apellido, empleado.Puesto).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y se leer los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)
	return resultado
}

func ModificarEmpleado(empleado entities.Empleado) int32 {
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_modificarempleado($1::INTEGER, $2::varchar(50), $3::varchar(50), $4::varchar(50))")
	if err != nil {
		log.Printf("Error al preparar el statement a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(empleado.Idempleado, empleado.Nombre, empleado.Apellido, empleado.Puesto).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y se leer los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)
	return resultado
}

func EliminarEmpleado(idEmpleado int) int32 {
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_eliminarempleado($1::integer)")
	if err != nil {
		log.Printf("Error al preparar el statement a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(idEmpleado).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y se leer los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)
	return resultado
}
