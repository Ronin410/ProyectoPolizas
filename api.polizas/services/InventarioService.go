package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	"main.go/entities"
)

func ConsultarInventario() (entities.Articulos, int) {
	log.Printf("Inicio InventarioService::consultarInventario")
	articulos := entities.Articulos{}

	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		if err, ok := err.(*pq.Error); ok {
			log.Printf("", err.Message)
		}
		return articulos, 0
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT sku_articulo,nombre_articulo, descripcion_articulo,cantidad_articulo FROM fun_consultarinventario()")
	if err != nil {
		log.Printf("Error al conectarse a la base de datos", err)
		return articulos, 0
	}
	defer rows.Close()

	// Recorrer los resultados
	for rows.Next() {
		articulo := entities.Articulo{}
		err := rows.Scan(&articulo.Sku, &articulo.NombreArticulo, &articulo.DescripcionArticulo, &articulo.CantidadArticulo)
		articulos = append(articulos, articulo)

		if err != nil {
			log.Printf("Error al recorrer los resultados de la funcion fun_consultarinventario()")
			return articulos, 0
		}
	}

	log.Printf("Termina InventarioService::consultarInventario")
	return articulos, 1
}

func AgregarInventario(articulo entities.Articulo) int32 {
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_agregarinventario($1::INTEGER, $2::varchar(50), $3::varchar(50), $4::INTEGER)")
	if err != nil {
		log.Printf("Error al preparar el statement a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(articulo.Sku, articulo.NombreArticulo, articulo.DescripcionArticulo, articulo.CantidadArticulo).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y se leer los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)
	return resultado
}

func ModificarInventario(articulo entities.Articulo) int32 {
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_modificarinventario($1::INTEGER, $2::INTEGER)")
	if err != nil {
		log.Printf("Error al preparar el statement a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(articulo.Sku, articulo.CantidadArticulo).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y se leer los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)
	return resultado
}

func EliminarArticulo(idArticulo int) int32 {
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer db.Close()

	// preparamos el statement
	stmt, err := db.Prepare("SELECT fun_eliminarinventario($1::INTEGER)")
	if err != nil {
		log.Printf("Error al preparar el statement a la base de datos")
		log.Printf("", err.Error())
		return 0
	}
	defer stmt.Close()

	// se llama la funcion y se cachan los resultados
	var resultado int32
	err = stmt.QueryRow(idArticulo).Scan(&resultado)
	if err != nil {
		log.Printf("Error al llamar la funcion y se leer los resultados")
		log.Printf("", err)
		return 0
	}

	fmt.Println("Resultado:", resultado)
	return resultado
}
