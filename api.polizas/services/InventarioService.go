package services

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"main.go/entities"
)

func ConsultarInventario() entities.Articulos {
	log.Printf("Inicio InventarioService::consultarInventario")

	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT sku_articulo,nombre_articulo, descripcion_articulo,cantidad_articulo FROM fun_consultarinventario()")
	if err != nil {
		log.Printf("Error al conectarse a la base de datos", err)
		if err, ok := err.(*pq.Error); ok {
			log.Printf("", err.Message)
		} //log.Fatal(err)
	}
	defer rows.Close()
	articulos := entities.Articulos{}

	// Recorrer los resultados
	for rows.Next() {
		articulo := entities.Articulo{}
		err := rows.Scan(&articulo.Sku, &articulo.NombreArticulo, &articulo.DescripcionArticulo, &articulo.CantidadArticulo)
		articulos = append(articulos, articulo)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Termina InventarioService::consultarInventario")
	return articulos
}
