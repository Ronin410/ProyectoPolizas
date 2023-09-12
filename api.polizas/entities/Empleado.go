package entities

type Empleado struct {
	Idempleado int64  `json:"idempleado"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Puesto     string `json:"puesto"`
}

type Empleados []Empleado

type EmpleadoResponse struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
}
