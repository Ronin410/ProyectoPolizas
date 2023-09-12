package services

import (
	"database/sql"
	"fmt"
	"log"
	"main.go/entities"
)

func ConsultarPolizasEmpleado(idempleado int32) entities.Polizas {
	log.Printf("Inicio PolizasService::ConsultarPolizasEmpleado")

	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT id_poliza,empleado_genero,sku_pol,cantidad_pol,fecha_pol,nombre_cliente FROM fun_consultapolizaempleado($1::INTEGER)", idempleado)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_consultapolizas()")
		log.Printf("", err)
	}
	defer rows.Close()

	polizas := entities.Polizas{}
	poliza := entities.Poliza{}

	// Recorrer los resultados
	for rows.Next() {

		err := rows.Scan(&poliza.IdPoliza, &poliza.EmpleadoGenero, &poliza.Sku, &poliza.Cantidad, &poliza.Fechamovto, &poliza.NombreCliente)
		if err != nil {
			log.Printf("Error al escanear los resultados de la funcion fun_consultapolizas()")
			log.Printf("", err)
		}
		polizas = append(polizas, poliza)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error al recorrer los resultados de la funcion fun_consultapolizas()")
		log.Printf("", err)
	}
	log.Printf("Termina PolizasService::ConsultarPolizasEmpleado")
	return polizas
}

func AgregarPoliza(poliza entities.PolizaEntrada) int32 {
	log.Printf("Inicia PolizasService::AgregarPoliza")

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_agregarpoliza($1::integer, $2::integer, $3::integer, $4::varchar(50))")
	if err != nil {
		log.Printf("Error al preparar el statement")
		log.Printf("", err)
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(poliza.EmpleadoGenero, poliza.Sku, poliza.Cantidad, poliza.NombreCliente).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y resivir los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)

	log.Printf("Termina PolizasService::AgregarPoliza")
	return resultado
}

func ActualizarPoliza(poliza entities.PolizaDetalle) int32 {
	log.Printf("Inicia PolizasService::ActualizarPoliza")

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_actualizarnombreclientepoliza($1::integer, $2::character varying)")
	if err != nil {
		log.Printf("Error al preparar el statement")
		log.Printf("", err)
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(poliza.IdPoliza, poliza.NombreCliente).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y resivir los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)

	log.Printf("Termina PolizasService::ActualizarPoliza")
	return resultado
}

func EliminarPoliza(idpoliza int) int32 {
	log.Printf("Inicia PolizasService::EliminarPoliza")

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_eliminarpoliza($1::integer)")
	if err != nil {
		log.Printf("Error al preparar el statement")
		log.Printf("", err)
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(idpoliza).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y resivir los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)

	log.Printf("Termina PolizasService::EliminarPoliza")
	return resultado
}

func ActualizarPoliza2(poliza entities.Poliza) int32 {
	log.Printf("Inicia PolizasService::ActualizarPoliza")

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_actualizarpoliza($1::integer, $2::integer,$3::integer,$4::integer, $5::character varying)")
	if err != nil {
		log.Printf("Error al preparar el statement")
		log.Printf("", err)
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(poliza.IdPoliza, poliza.EmpleadoGenero, poliza.Sku, poliza.Cantidad, poliza.NombreCliente).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y resivir los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)

	log.Printf("Termina PolizasService::ActualizarPoliza")
	return resultado
}
