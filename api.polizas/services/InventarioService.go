package services

import (
	"database/sql"
	"log"
	"main.go/entities"
)

func ConsultarInventario() (entities.Articulos, int) {
	log.Printf("Inicio InventarioService::consultarInventario")
	articulos := entities.Articulos{}

	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
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
