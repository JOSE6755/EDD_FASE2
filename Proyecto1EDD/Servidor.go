package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"crypto/sha256"

	"./ArbolAVL"
	"./ArbolB"
	"./Grafo"
	"./Hash"
	"./MatrizD"
	"./Merkle"
	"github.com/gorilla/mux"
)

type Datos_fin struct {
	Datos []Datoss `json:"Datos"`
}

type Datoss struct {
	Indice        string          `json:"Indice"`
	Departamentos []Departamentos `json:"Departamentos"`
}

type Departamentos struct {
	Nombre  string    `json:"Nombre"`
	Tiendas []Tiendas `json:"Tiendas"`
}

type Tiendas struct {
	Nombre       string `json:"Nombre"`
	Descripcion  string `json:"Descripcion"`
	Contacto     string `json:"Contacto"`
	Calificacion int    `json:"Calificacion"`
	Logo         string `json:"Logo"`
	Comentario   *Hash.Hashtable
	Productos    *Merkle.Arbolito
	Pedidos      *Merkle.Arbolito
}

type busqueda struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
	Comentarios  Hash.Node
}

type eliminacion struct {
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	Calificacion int    `json:"Calificacion"`
}
type Inventarios struct {
	Tienda []DT `json:"Inventarios"`
}
type DT struct {
	Tienda       string      `json:"Tienda"`
	Departameto  string      `json:"Departamento"`
	Calificacion int         `json:"Calificacion"`
	Productos    []Productos `json:"Productos"`
}
type Productos struct {
	Nombre         string  `json:"Nombre"`
	Codigo         int     `json:"Codigo"`
	Descripcion    string  `json:"Descripcion"`
	Precio         float64 `json:"Precio"`
	Cantidad       int     `json:"Cantidad"`
	Imagen         string  `json:"Imagen"`
	Almacenamiento string  `json:"Almacenamiento"`
}
type auxiliar struct {
	Prod []Productos `json:"Productos"`
}

var indices []string
var depas []string
var vector []Lista_doble
var datos Datos_fin
var inv Inventarios
var tempo busqueda
var usuarios ArbolB.Arbol
var distancias Grafo.Total
var grafo Grafo.ListaDoble
var MerkleUs *Merkle.ArbolitoUs
var MerkleTi *Merkle.ArbolitoT

func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola")

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func list(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datos)
}

func crear(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")

	}
	json.Unmarshal(reqbody, &datos)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datos)

	llenar(datos)

}

func buscar(w http.ResponseWriter, r *http.Request) {
	var dat busqueda
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &dat)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	//encontrado(dat, w)
}

func invent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")

	}
	json.Unmarshal(reqbody, &inv)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	inven(inv)

}

func prod(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	Encontrado(tempo.Nombre, tempo.Departamento, tempo.Calificacion, nil, w, 0, nil, nil, 0, 0)

}

func temporal(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &tempo)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(tempo.Departamento)

}

func Pedidos(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var prueba MatrizD.Pedidos
	pila := Grafo.Pila{}
	numP := 0

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &prueba)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	Grafo.Tipos(prueba, &numP)
	tipos := make([]string, numP)

	mPedidos(prueba, &tipos)
	aux := 0
	for i := 0; i < len(prueba.Ped); i++ {
		for j := 0; j < len(prueba.Ped[i].Productos); j++ {
			tipos[aux] = prueba.Ped[i].Productos[j].Almacenamiento
			aux++
		}
	}
	distancia := 0.0
	tabla := "digraph {\n    tbl[\n     shape=plaintext\n     label=<\n     <table border='0' cellborder='1' color='blue' cellspacing='0'>\n" + "<tr> <td>Inicio</td> <td> Final </td> <td>Distancia</td></tr>\n"
	inicio := tipos[0]
	pila.Push(grafo.PosIn)
	pila.Push(inicio)

	grafo.Dikstra(grafo.PosIn, tipos[0], len(distancias.Datos), &distancia)
	tabla += "<tr>\n<td>" + grafo.PosIn + "</td>\n<td>" + inicio + "</td>\n<td>" + fmt.Sprint(distancia) + "</td>\n</tr>\n"

	for i := 1; i < len(tipos); i++ {
		if Grafo.Existe(pila, tipos[i]) == false {

			grafo.Dikstra(inicio, tipos[i], len(distancias.Datos), &distancia)
			pila.Push(tipos[i])
			tabla += "<tr>\n<td>" + inicio + "</td>\n<td>" + tipos[i] + "</td>\n<td>" + fmt.Sprint(distancia) + "</td>\n</tr>\n"
			inicio = tipos[i]

		}
	}
	grafo.Dikstra(inicio, grafo.PosFin, len(distancias.Datos), &distancia)
	tabla += "<tr>\n<td>" + inicio + "</td>\n<td>" + grafo.PosFin + "</td>\n<td>" + fmt.Sprint(distancia) + "</td>\n</tr>"
	tabla += "</table>\n>]\n}"
	err2 := ioutil.WriteFile("Camino.dot", []byte(tabla), 0644)
	if err2 != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", "Camino.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("Camino.png", cmd, os.FileMode(mode))
	fmt.Println(tabla)

}

func NuevoP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var prueba MatrizD.Pedidos

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &prueba)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	usuarios.Gpedido(prueba.Usuario, prueba, nil)

}
func CargarP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var prueba int
	var aux MatrizD.Pedidos
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &prueba)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	usuarios.Gpedido(prueba, aux, w)

}
func listaPedidos(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var prueba busqueda

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &prueba)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	Encontrado(prueba.Nombre, prueba.Departamento, prueba.Calificacion, nil, w, 1, nil, nil, 0, 0)

}
func matrices(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima MatrizD.Imagenes

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	MatrizD.Img(ima.Nombre, ima.Año, ima.Mes, w)
}
func años(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima MatrizD.Imagenes

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	MatrizD.Años(ima.Nombre, ima.Año, ima.Mes, w)
}
func meses(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima MatrizD.Imagenes

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	MatrizD.Meses(ima.Nombre, ima.Año, ima.Mes, w)
}
func usuario(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var datosUs ArbolB.Usuas
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &datosUs)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	for i := 0; i < len(datosUs.Usuarios); i++ {
		contra := sha256.Sum256([]byte(datosUs.Usuarios[i].Password))
		datosUs.Usuarios[i].Password = fmt.Sprintf("%x", contra)
	}

	if usuarios.Raiz == nil {
		MerkleUs = Merkle.NuevoArbolUs()
		usuarios = *ArbolB.Nuevoarbol(5)
		for i := 0; i < len(datosUs.Usuarios); i++ {
			usuarios.Insertar(datosUs.Usuarios[i])
			MerkleUs.InsertarUs(datosUs.Usuarios[i].DPI, datosUs.Usuarios[i].Password, datosUs.Usuarios[i].Correo)
		}
	} else {
		for i := 0; i < len(datosUs.Usuarios); i++ {
			usuarios.Insertar(datosUs.Usuarios[i])
			MerkleUs.InsertarUs(datosUs.Usuarios[i].DPI, datosUs.Usuarios[i].Password, datosUs.Usuarios[i].Correo)
		}
	}
	usuarios.Graficar("ArbolB")
	usuarios.Graficar("ArbolB")
	usuarios.Graficar("ArbolB")
	MerkleUs.CodigoUs()
}

func buscarUsu(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima ArbolB.Usu

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	contra := sha256.Sum256([]byte(ima.Pass))

	ima.Pass = fmt.Sprintf("%x", contra)
	usuarios.Buscar(ima.DPI, ima.Pass, ima.Correo, w)
}

func crearGrafo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &distancias)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)

	if grafo.Inicio == nil {
		grafo = *Grafo.Nuevosvertices(distancias.PosicionI, distancias.Entrega)
	}
	Grafo.CrearMadya(len(distancias.Datos))
	for i := 0; i < len(distancias.Datos); i++ {
		grafo.Insertar(distancias.Datos[i].Nombre, i)

	}
	for i := 0; i < len(distancias.Datos); i++ {
		for j := 0; j < len(distancias.Datos[i].Enlaces); j++ {
			grafo.Enlazar(distancias.Datos[i].Nombre, distancias.Datos[i].Enlaces[j].Nombre, distancias.Datos[i].Enlaces[j].Distancia)

		}
	}
	grafo.Graficar()

}

func arbolB(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima string

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	ArbolB.BuscarImg("arbolB", w)
}
func arbolC(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima string

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	ArbolB.BuscarImg("arbolC", w)
}
func arbolCS(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima string

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	ArbolB.BuscarImg("arbolCS", w)
}
func ImgGrafo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima string

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	ArbolB.BuscarImg("Grafo", w)
}
func Caminito(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima string

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	ArbolB.BuscarImg("Camino", w)
}
func TablaHash(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima busqueda

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	//Falta mandarlo a buscar e insertar en esa tienda

}
func TablaHashProd(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var ima busqueda

	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &ima)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	//Falta mandarlo a buscar e insertar en esa tienda

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", inicio)

	router.HandleFunc("/Cargar", crear).Methods("POST")
	router.HandleFunc("/CrearU", usuario).Methods("POST")
	router.HandleFunc("/Caminito", Caminito).Methods("GET")
	router.HandleFunc("/ImgGrafo", ImgGrafo).Methods("GET")
	router.HandleFunc("/ArbolB", arbolB).Methods("GET")
	router.HandleFunc("/ArbolC", arbolC).Methods("GET")
	router.HandleFunc("/ArbolCS", arbolCS).Methods("GET")
	router.HandleFunc("/Gpedido", NuevoP).Methods("POST")
	router.HandleFunc("/CargarP", CargarP).Methods("POST")
	router.HandleFunc("/BuscarU", buscarUsu).Methods("POST")
	router.HandleFunc("/Grafo", crearGrafo).Methods("POST")
	router.HandleFunc("/Calendario", listaPedidos).Methods("POST")
	router.HandleFunc("/Imagenes", matrices).Methods("POST")
	router.HandleFunc("/ImagenesAños", años).Methods("POST")
	router.HandleFunc("/ImagenesMes", meses).Methods("POST")
	router.HandleFunc("/Inventario", invent).Methods("POST")
	router.HandleFunc("/Pedidos", Pedidos).Methods("POST")
	router.HandleFunc("/Datos", temporal).Methods("POST")
	router.HandleFunc("/Tiendas/productos", prod).Methods("GET")
	router.HandleFunc("/Tiendas", list).Methods("GET")
	router.HandleFunc("/id/{id}", pornum).Methods("GET")
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	router.HandleFunc("/Guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", graficar).Methods("GET")
	http.ListenAndServe(":3000", router)
	log.Fatal(http.ListenAndServe(":3000", router))
}

type nodo struct {
	/*nombre       string `json:"Nombre"`
	descripcion  string `json:"Descripcion"`
	contacto     string `json:"Contacto"`
	calificacion int    `json:"Calificacion"`
	*/
	Tiendas   Tiendas
	Productos int
	arbol     *ArbolAVL.Arbolavl
	pedidos   *MatrizD.Lista_Simple
	siguiente *nodo
	anterior  *nodo
}

type Lista_doble struct {
	inicio   *nodo
	fin      *nodo
	cantidad int
}

func (l *Lista_doble) insertar(n Tiendas) {
	nuevo := &nodo{Tiendas: n}
	comentarios := Hash.NuevaTabla(7, 50, 50)
	nuevo.Tiendas.Comentario = comentarios

	ArbolP := Merkle.NuevoArbol()

	nuevo.Tiendas.Pedidos = ArbolP
	if l.inicio == nil {

		l.inicio = nuevo
		l.fin = nuevo
		l.cantidad++
	} else {
		fin := l.fin
		fin.siguiente = nuevo
		fin.siguiente.anterior = fin
		l.fin = nuevo
		l.cantidad++
	}

}
func (l *Lista_doble) listar(w http.ResponseWriter) {
	inicio := l.inicio

	for inicio != nil {
		json.NewEncoder(w).Encode(inicio.Tiendas)
		inicio = inicio.siguiente
	}
}

var ex bool

func find(nombre string, c Lista_doble, arbol *ArbolAVL.Arbolavl, w http.ResponseWriter, tiendas int, pedidos *MatrizD.Lista_Simple, nuevo *MatrizD.NodoInfo, año int, mes int) bool {
	encontrado := false
	inicio := c.inicio

	for inicio != nil {
		if inicio.Tiendas.Nombre == nombre {
			encontrado = true
			fmt.Println("Si entre")
			if arbol != nil {
				inicio.arbol = arbol
				inicio.Productos = tiendas
			} else if arbol == nil && nuevo == nil && año == 0 && pedidos == nil && tiendas == 0 {
				produ := make([][]string, inicio.Productos)
				for i := 0; i < inicio.Productos; i++ {
					produ[i] = make([]string, 7)

				}
				ArbolAVL.Matz(inicio.Productos)
				inicio.arbol.Raiz.DisplayNodesInOrder()
				produ = ArbolAVL.Regres()
				aux := auxiliar{}
				for i := 0; i < inicio.Productos; i++ {
					codigo, _ := strconv.Atoi(produ[i][1])
					precio, _ := strconv.Atoi(produ[i][3])
					cantidad, _ := strconv.Atoi(produ[i][4])
					aux2 := Productos{Nombre: produ[i][0], Codigo: codigo, Descripcion: produ[i][2], Precio: float64(precio), Cantidad: cantidad, Imagen: produ[i][5], Almacenamiento: produ[i][6]}
					aux.Prod = append(aux.Prod, aux2)
				}

				//fmt.Println(b, "\n", aux)
				json.NewEncoder(w).Encode(aux)
			} else if pedidos != nil {
				inicio.pedidos = pedidos
				inicio.pedidos.Añitos(nombre)
				tiendasM(nombre, pedidos)

			} else if w != nil {
				inicio.pedidos.Listaaños(w)
			}

			break
		} else {
			inicio = inicio.siguiente
		}
	}
	return encontrado

}

func pornum(w http.ResponseWriter, r *http.Request) {
	ind := mux.Vars(r)
	id, err := strconv.Atoi(ind["id"])

	if err != nil {
		fmt.Fprint(w, "Ingrese un dato valido")
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	vector[id-1].listar(w)
}

func indi(a Datos_fin) int {
	var n = 0
	indices = make([]string, len(a.Datos))
	for i := 0; i < len(indices); i++ {
		indices[i] = a.Datos[i].Indice
	}
	fmt.Println(indices)

	return n
}

func llenar(a Datos_fin) {

	depas = make([]string, len(a.Datos[0].Departamentos))
	for i := 0; i < len(depas); i++ {
		depas[i] = a.Datos[0].Departamentos[i].Nombre
	}
	indi(a)
	fmt.Println(depas)

	vector = make([]Lista_doble, (len(a.Datos) * len(a.Datos[0].Departamentos) * 5))

	ingresar(a)
	MerkleTi.CodigoT()

}

func Encontrado(nombre string, departamento string, calificacion int, arbol *ArbolAVL.Arbolavl, w http.ResponseWriter, produ int, pedidos *MatrizD.Lista_Simple, nuevo *MatrizD.NodoInfo, año int, mes int) {
	n := 0
	d := 0

	for i := 0; i < len(depas); i++ {
		fmt.Println(nombre)
		if depas[i] == departamento {
			d = i

			break
		}
	}
	fmt.Println(len(indices))
	for i := 0; i < len(indices); i++ {

		n = (i*len(depas)+d)*5 + (calificacion - 1)
		fmt.Println(n)
		if find(nombre, vector[n], arbol, w, produ, pedidos, nuevo, año, mes) == true {
			fmt.Println("asdasdasd")
			break
		}
	}

}

func ingresar(datos Datos_fin) {
	MerkleTi = Merkle.NuevoArbolT()
	for i := 0; i < len(datos.Datos); i++ {
		for j := 0; j < len(datos.Datos[i].Departamentos); j++ {
			for k := 0; k < len(datos.Datos[i].Departamentos[j].Tiendas); k++ {
				MerkleTi.Insertar(datos.Datos[i].Departamentos[j].Tiendas[k].Nombre, datos.Datos[i].Departamentos[j].Tiendas[k].Contacto, datos.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
				vector[((i*len(datos.Datos[i].Departamentos)+j)*5 + (datos.Datos[i].Departamentos[j].Tiendas[k].Calificacion - 1))].insertar(datos.Datos[i].Departamentos[j].Tiendas[k])
			}
		}

	}

}

func eliminar(w http.ResponseWriter, r *http.Request) {
	var dat eliminacion
	d := 0
	n := 0
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Inserte datos validos")
	}
	json.Unmarshal(reqbody, &dat)
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	for i := 0; i < len(depas); i++ {
		if depas[i] == dat.Categoria {
			d = i
			break
		}
	}
	for i := 0; i < len(indices); i++ {
		n = (i*len(depas)+d)*5 + (dat.Calificacion - 1)
		if vector[n].elimencon(dat, w) == true {
			break
		}
	}
	z := vector[n].inicio
	for z != nil {
		fmt.Println(z.Tiendas)
		z = z.siguiente

	}
}

func (l *Lista_doble) elimencon(a eliminacion, w http.ResponseWriter) bool {
	inicio := l.inicio
	echo := false
	if inicio != nil {
		if inicio.Tiendas.Nombre == a.Nombre && inicio.Tiendas.Calificacion == a.Calificacion {

			l.inicio = inicio.siguiente
			json.NewEncoder(w).Encode(l.inicio.Tiendas)
			echo = true
			l.cantidad--
		} else {
			for inicio != nil {
				if inicio.Tiendas.Nombre == a.Nombre && inicio.Tiendas.Calificacion == a.Calificacion {
					inicio.anterior.siguiente = inicio.siguiente
					if inicio.siguiente != nil {
						inicio.siguiente.anterior = inicio.anterior
					}
				}
				inicio = inicio.siguiente

			}
			l.cantidad--
			echo = true
		}
	}

	return echo
}

// Costo mucho :(
func graficar(w http.ResponseWriter, r *http.Request) {
	n := 0
	ayu := 0
	ayu2 := 0
	nodos := "{rank=same;"
	lista := ""
	doc := "digraph G {\nnode[shape=record]\n" + `graph[splines="ortho"]` + "\n"
	aux := make([]string, (len(indices) * len(depas) * 5))
	aux2 := make([]string, len(aux)/5)
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(depas); j++ {
			for k := 0; k < 5; k++ {
				n = (i*len(depas)+j)*5 + k

				aux[n] = "nodo" + strconv.Itoa(n) + `[label="` + indices[i] + "|" + depas[j] + "|" + "POS:" + strconv.Itoa(n+1) + `"]`

			}
		}

	}

	for i := 0; i < len(aux); i++ {
		aux2[ayu] += aux[i] + "\n"
		if ayu2 <= 4 {
			nodos += "nodo" + strconv.Itoa(i) + ";"
		}

		ayu2++
		if ayu2 == 5 {
			nodos += "}"
			aux2[ayu] += nodos
			nodos = "{rank=same;"
			ayu2 = 0
			ayu++
		}
	}
	ayu = 0
	ayu2 = 0
	nodos = ""
	dots := 0
	for i := 0; i < len(aux); i++ {
		if i != len(aux) && ayu2 < 4 {
			nodos += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(i+1) + "\n"
		}
		lista += vector[i].listar2(i)
		ayu2++

		if ayu2 == 5 {

			ayu2 = 0
			doc += aux2[ayu] + "\n" + nodos
			doc += lista + "\n}"

			fmt.Println(doc)
			err := ioutil.WriteFile("Tiendas"+strconv.Itoa(dots+1)+".dot", []byte(doc), 0644)
			if err != nil {
				log.Fatal(err)
			}
			ruta, _ := exec.LookPath("dot")
			cmd, _ := exec.Command(ruta, "-Tpng", "./Tiendas"+strconv.Itoa(dots+1)+".dot").Output()
			mode := int(0777)
			ioutil.WriteFile("Tiendas"+strconv.Itoa(dots+1)+".png", cmd, os.FileMode(mode))
			doc = "digraph G {\nnode[shape=record]\n" + `graph[splines="ortho"]` + "\n"
			nodos = ""
			lista = ""
			ayu++
			dots++
		}
	}

}

func (l *Lista_doble) listar2(n int) string {
	inicio := l.inicio
	nodos := "nodo" + strconv.Itoa(n) + "->"
	datos := ""
	if inicio != nil {
		for inicio != nil {
			datos += inicio.Tiendas.Nombre + `[label="` + inicio.Tiendas.Nombre + "|" + inicio.Tiendas.Contacto + "|" + strconv.Itoa(inicio.Tiendas.Calificacion) + `"]` + "\n"
			if inicio.siguiente != nil {
				nodos += inicio.Tiendas.Nombre + "->" + inicio.siguiente.Tiendas.Nombre
			}
			if l.inicio == l.fin {
				nodos += inicio.Tiendas.Nombre
			}
			inicio = inicio.siguiente
		}
		datos += nodos + "\n"
		return datos
	}
	return datos
}

//No me salio :()

func guardar(w http.ResponseWriter, r *http.Request) {

	var datos2 Datos_fin
	tam := 0
	datos2.Datos = make([]Datoss, len(indices))

	for i := 0; i < len(indices); i++ {
		datos2.Datos[i].Departamentos = make([]Departamentos, len(depas))
	}
	for i := 0; i < len(indices); i++ {
		datos2.Datos[i].Indice = indices[i]
		for j := 0; j < len(depas); j++ {
			datos2.Datos[i].Departamentos[j].Nombre = depas[j]
			for k := 0; k < 5; k++ {
				n := (i*len(depas)+j)*5 + k
				tam += vector[n].cantidad
				if k == 4 {
					datos2.Datos[i].Departamentos[j].Tiendas = make([]Tiendas, tam)
					tam = 0
					inicio := vector[n].inicio
					for z := 0; z < len(datos2.Datos[i].Departamentos[j].Tiendas); z++ {
						datos2.Datos[i].Departamentos[j].Tiendas[z].Nombre = inicio.Tiendas.Nombre
						datos2.Datos[i].Departamentos[j].Tiendas[z].Descripcion = inicio.Tiendas.Descripcion
						datos2.Datos[i].Departamentos[j].Tiendas[z].Contacto = inicio.Tiendas.Contacto
						datos2.Datos[i].Departamentos[j].Tiendas[z].Calificacion = inicio.Tiendas.Calificacion
						if inicio.siguiente != nil {
							inicio = inicio.siguiente
						} else {
							break
						}
					}
				}

			}

		}
	}

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(datos2)
	file, err := os.Create("datos2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	io.Copy(file, buf)

}

func listarT(a Datos_fin, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(a)
}

func inven(t Inventarios) {
	for i := 0; i < len(t.Tienda); i++ {
		prueba := &ArbolAVL.Arbolavl{}
		//ArbolM := Merkle.NuevoArbol()
		for j := 0; j < len(t.Tienda[i].Productos); j++ {
			prueba.Insertar(t.Tienda[i].Productos[j].Nombre, t.Tienda[i].Productos[j].Codigo, t.Tienda[i].Productos[j].Descripcion, t.Tienda[i].Productos[j].Precio, t.Tienda[i].Productos[j].Cantidad, t.Tienda[i].Productos[j].Imagen, t.Tienda[i].Productos[j].Almacenamiento)
			//ArbolM.Insertar()

		}
		Encontrado(t.Tienda[i].Tienda, t.Tienda[i].Departameto, t.Tienda[i].Calificacion, prueba, nil, len(t.Tienda[i].Productos), nil, nil, 0, 0)
	}

}

func BuscarAVL(nombre string, departamento string, calificacion int, codigo int, cantidad int, l *MatrizD.Lista_Simple) bool {
	n := 0
	d := 0
	b := false

	for i := 0; i < len(depas); i++ {
		fmt.Println(nombre)
		if depas[i] == departamento {
			d = i
			fmt.Println("Depas entro")

			break
		}
	}
	fmt.Println(len(indices), " ", len(depas))
	for i := 0; i < len(indices); i++ {

		n = (i*len(depas)+d)*5 + (calificacion - 1)
		fmt.Println(n)
		if AV(nombre, vector[n], codigo, cantidad, l) == true {
			fmt.Println("asdasdasd")
			b = true
			break
		}
	}
	return b

}
func AV(nombre string, c Lista_doble, codigo int, cantidad int, l *MatrizD.Lista_Simple) bool {
	encontrado := false
	inicio := c.inicio

	for inicio != nil {
		if inicio.Tiendas.Nombre == nombre {
			if l == nil {
				ArbolAVL.Inserencon()
				inicio.arbol.Raiz.DisplayNodesInOrder2(codigo, cantidad)

				if ArbolAVL.Getencon() == true {
					encontrado = true
					break
				}
			} else {
				encontrado = true
				inicio.pedidos = l

				break
			}

		} else {
			inicio = inicio.siguiente
		}
	}
	return encontrado
}

func mPedidos(p MatrizD.Pedidos, grafo *[]string) {
	mex := false
	mex3 := false

	for i := 0; i < len(p.Ped); i++ {

		a := strings.Split(p.Ped[i].Fecha, "-")
		d, _ := strconv.Atoi(a[0])
		m, _ := strconv.Atoi(a[1])
		año, _ := strconv.Atoi(a[2])
		mex2 := Buscmat(p.Ped[i].Tienda, p.Ped[i].Departamento, p.Ped[i].Calificacion, d, año, m, &mex, nil, nil, &mex3, nil)
		if mex == false {

			nombre := ""
			if mex2 == false && mex3 == false {
				nuevaM := &MatrizD.Matriz{}
				nuevaLD := &MatrizD.Lista_doble{}
				listaP := &MatrizD.Lista_Simple{}
				for j := 0; j < len(p.Ped[i].Productos); j++ {

					if BuscarAVL(p.Ped[i].Tienda, p.Ped[i].Departamento, p.Ped[i].Calificacion, p.Ped[i].Productos[j].Codigo, 1, nil) == true {
						(*grafo)[i] = p.Ped[i].Productos[j].Almacenamiento
						nuevoND := &MatrizD.NodoInfo{ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil, Cantida: 1, Producto: p.Ped[i].Productos[j].Codigo, Precio: ArbolAVL.Getprec(), Dia: d, Categoria: p.Ped[i].Departamento}

						nuevaM.Inser(nuevoND)
						nombre = p.Ped[i].Tienda

					}
				}
				nuevaLD.InserDoble(nuevaM, m)
				listaP.InserSimple(nuevaLD, año)
				nuevaM.Graficar(d, m, año, nombre)
				Encontrado(nombre, p.Ped[i].Departamento, p.Ped[i].Calificacion, nil, nil, 0, listaP, nil, 0, 0)

			} else if mex2 == true {
				nuevaM := &MatrizD.Matriz{}
				for j := 0; j < len(p.Ped[i].Productos); j++ {

					if BuscarAVL(p.Ped[i].Tienda, p.Ped[i].Departamento, p.Ped[i].Calificacion, p.Ped[i].Productos[j].Codigo, 1, nil) == true {
						(*grafo)[i] = p.Ped[i].Productos[j].Almacenamiento
						nuevoND := &MatrizD.NodoInfo{ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil, Cantida: 1, Producto: p.Ped[i].Productos[j].Codigo, Precio: ArbolAVL.Getprec(), Dia: d, Categoria: p.Ped[i].Departamento}

						nuevaM.Inser(nuevoND)
						nombre = p.Ped[i].Tienda

					}
				}
				nuevaM.Graficar(d, m, año, nombre)
				Buscmat(p.Ped[i].Tienda, p.Ped[i].Departamento, p.Ped[i].Calificacion, d, año, m, nil, nuevaM, nil, nil, nil)
			} else if mex3 == true {
				mex3 = false
				nuevaM := &MatrizD.Matriz{}
				nuevaLD := &MatrizD.Lista_doble{}

				for j := 0; j < len(p.Ped[i].Productos); j++ {

					if BuscarAVL(p.Ped[i].Tienda, p.Ped[i].Departamento, p.Ped[i].Calificacion, p.Ped[i].Productos[j].Codigo, 1, nil) == true {
						(*grafo)[i] = p.Ped[i].Productos[j].Almacenamiento
						nuevoND := &MatrizD.NodoInfo{ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil, Cantida: 1, Producto: p.Ped[i].Productos[j].Codigo, Precio: ArbolAVL.Getprec(), Dia: d, Categoria: p.Ped[i].Departamento}

						nuevaM.Inser(nuevoND)
						nombre = p.Ped[i].Tienda

					}
				}
				nuevaLD.InserDoble(nuevaM, m)

				Buscmat(nombre, p.Ped[i].Departamento, p.Ped[i].Calificacion, d, año, m, nil, nil, nil, nil, nuevaLD)
				nuevaM.Graficar(d, m, año, nombre)
			}

		} else {
			for j := 0; j < len(p.Ped[i].Productos); j++ {
				(*grafo)[i] = p.Ped[i].Productos[j].Almacenamiento
				if BuscarAVL(p.Ped[i].Tienda, p.Ped[i].Departamento, p.Ped[i].Calificacion, p.Ped[i].Productos[j].Codigo, 1, nil) == true {
					nuevoND := &MatrizD.NodoInfo{ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil, Cantida: 1, Producto: p.Ped[i].Productos[j].Codigo, Precio: ArbolAVL.Getprec(), Dia: d, Categoria: p.Ped[i].Departamento}
					nombre := p.Ped[i].Tienda
					Buscmat(nombre, p.Ped[i].Departamento, p.Ped[i].Calificacion, d, año, m, nil, nil, nuevoND, nil, nil)

				}

			}

		}

	}
}

func Buscmat(nombre string, departamento string, calificacion int, dia int, año int, mes int, m *bool, nuevaM *MatrizD.Matriz, nuevoP *MatrizD.NodoInfo, j *bool, LD *MatrizD.Lista_doble) bool {
	n := 0
	d := 0
	encontrado := false

	for i := 0; i < len(depas); i++ {
		fmt.Println(nombre)
		if depas[i] == departamento {
			d = i

			break
		}
	}
	fmt.Println(len(indices))
	for i := 0; i < len(indices); i++ {

		n = (i*len(depas)+d)*5 + (calificacion - 1)
		fmt.Println(n)
		if findmatriz(nombre, vector[n], dia, año, mes, m, nuevaM, nuevoP, j, LD) == true {
			fmt.Println("asdasdasd")
			encontrado = true
			break
		}
	}
	return encontrado
}
func findmatriz(nombre string, c Lista_doble, dia int, año int, mes int, m *bool, nuevaM *MatrizD.Matriz, nuevoP *MatrizD.NodoInfo, n *bool, LD *MatrizD.Lista_doble) bool {
	encontrado := false
	inicio := c.inicio

	for inicio != nil {
		if inicio.Tiendas.Nombre == nombre {
			if nuevaM == nil && inicio.pedidos != nil {
				if n != nil {
					*n = true
				}
				if LD != nil {
					inicio.pedidos.InserSimple(LD, año)
					tiendasM(nombre, inicio.pedidos)
					inicio.pedidos.Añitos(nombre)
					break
				}
				if inicio.pedidos.Esnul() == false {

					encontrado = inicio.pedidos.Buscar(dia, año, mes, nuevoP, m, nombre)
					break
				} else {
					encontrado = false
					break
				}
			} else if nuevaM != nil {
				inicio.pedidos.Nuevomes(nuevaM, mes)
				tiendasM(nombre, inicio.pedidos)
				inicio.pedidos.Añitos(nombre)
			}
			break
		} else {
			inicio = inicio.siguiente
		}
	}
	return encontrado

}

func tiendasM(nombre string, m *MatrizD.Lista_Simple) {
	n := 0
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(depas); j++ {
			for z := 0; z < 5; z++ {
				n = (i*len(depas)+j)*5 + z
				mismatienda(vector[n], nombre, m)

			}
		}
	}

}
func mismatienda(c Lista_doble, nombre string, m *MatrizD.Lista_Simple) {
	inicio := c.inicio
	for inicio != nil {
		if inicio.Tiendas.Nombre == nombre {
			inicio.pedidos = m
		}
		inicio = inicio.siguiente
	}
}

func popaños(nombre string, departamento string, calificacion int, w http.ResponseWriter) {
	//n := 0
	//d := 0

	for i := 0; i < len(depas); i++ {
		fmt.Println(nombre)
		if depas[i] == departamento {
			//d = i
			fmt.Println("Depas entro")

			break
		}
	}
	fmt.Println(len(indices), " ", len(depas))
	for i := 0; i < len(indices); i++ {

		//n = (i*len(depas)+d)*5 + (calificacion - 1)

	}

}

/*func mandaraños(c Lista_doble, nombre string, w http.ResponseWriter) bool {
	inicio := c.inicio
	encontrado := false
	for inicio != nil {
		if inicio.Tiendas.Nombre == nombre {
			encontrado=true

		} else {
			inicio = inicio.siguiente
		}
	}

}
*/
