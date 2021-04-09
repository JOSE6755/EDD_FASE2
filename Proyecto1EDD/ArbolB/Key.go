package ArbolB

type Key struct {
	DPI       int    `json:"Dpi"`
	Nombre    string `json:"Nombre"`
	Correo    string `json:"Correo"`
	Password  string `json:"Password"`
	Cuenta    string `json:"Cuenta"`
	Izquierdo *Nodo
	Derecho   *Nodo
}

func nuevaK(DPI int, Nombre string, Correo string, Password string, Cuenta string) *Key {
	k := &Key{DPI: DPI, Nombre: Nombre, Correo: Correo, Password: Password, Cuenta: Cuenta}
	return k

}
