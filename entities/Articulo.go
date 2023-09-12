package entities

type Articulo struct {
	Sku                 int
	NombreArticulo      string
	DescripcionArticulo string
	CantidadArticulo    int
	fechamovto          string
}

type Articulos []Articulo

type DetalleArticulo struct {
	sku                 int32
	nombreArticulo      string
	descripcionArticulo string
}
