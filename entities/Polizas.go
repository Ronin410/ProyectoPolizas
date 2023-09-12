package entities

type Poliza struct {
	IdPoliza       int32
	EmpleadoGenero int32
	Sku            int32
	Cantidad       int32
	Fechamovto     string
	NombreCliente  string
}

type Polizas []Poliza

type PolizaDetalle struct {
	IdPoliza      int32
	Cantidad      int32
	Fechamovto    string
	NombreCliente string
}

type PolizaEntrada struct {
	EmpleadoGenero int32
	Sku            int32
	Cantidad       int32
	NombreCliente  string
}

type PolizasResponse struct {
	Poliza          PolizaDetalle
	Empleado        EmpleadoResponse
	DetalleArticulo DetalleArticulo
}

type PolizaMessage struct {
	Message string
}
