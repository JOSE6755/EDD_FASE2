package Grafo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"../MatrizD"
)

type Total struct {
	Datos     []Nodos `json:"Nodos"`
	PosicionI string  `json:"PosicionInicialRobot"`
	Entrega   string  `json:"Entrega"`
}

type Nodos struct {
	Nombre  string    `json:"Nombre"`
	Enlaces []Enlaces `json:"Enlaces"`
}
type Enlaces struct {
	Nombre    string  `json:"Nombre"`
	Distancia float64 `json:"Distancia"`
}

type ListaDoble struct {
	Inicio *Nodo
	PosIn  string
	PosFin string
}
type Nodo struct {
	Siguiente *Nodo
	Anterior  *Nodo
	TipoProd  string
	Peso      float64
	id        int
	Enlaces   *ListaDoble2
}

type ListaDoble2 struct {
	inicio *Nodo2
	fin    *Nodo2
}

type Nodo2 struct {
	Siguiente *Nodo2
	Anterior  *Nodo2
	TipoP     string
	Peso      float64
	id        int
}
type Pila struct {
	TipoProd []string
}

var matrizAd [][]float64

func CrearMadya(len int) {
	matrizAd = make([][]float64, len)
	for i := 0; i < len; i++ {
		matrizAd[i] = make([]float64, len)
	}
}

func Tipos(p MatrizD.Pedidos, numP *int) {
	for i := 0; i < len(p.Ped); i++ {
		*numP += len(p.Ped[i].Productos)
	}
}

func (p *Pila) Push(Tipo string) {
	p.TipoProd = append(p.TipoProd, Tipo)
}

func (p *Pila) Pop() string {
	fuera := p.TipoProd[len(p.TipoProd)-1]
	p.TipoProd = p.TipoProd[:len(p.TipoProd)-1]
	return fuera
}

func Nuevosvertices(inicio string, final string) *ListaDoble {
	nueva := &ListaDoble{PosIn: inicio, PosFin: final}
	return nueva
}

func listaAdya() *ListaDoble2 {
	nueva := &ListaDoble2{}
	return nueva
}

func (l *ListaDoble) getVertice(Tipo string) *Nodo {
	inicio := l.Inicio
	for inicio != nil {
		if inicio.TipoProd == Tipo {
			return inicio
		}
		inicio = inicio.Siguiente
	}
	return nil
}

var inicio int
var final int

func index() {

}
func (l *ListaDoble) Insertar(Tipo string, i int) {
	if l.getVertice(Tipo) == nil {
		n := &Nodo{TipoProd: Tipo, Enlaces: listaAdya(), id: i}
		if l.Inicio == nil {
			l.Inicio = n
		} else {
			inicio := l.Inicio
			for inicio.Siguiente != nil {
				inicio = inicio.Siguiente
			}
			inicio.Siguiente = n
		}
	} else {
		fmt.Println("Ya existe ese vertice")
	}

}

func (l *ListaDoble2) InsertarEnlace(TipoP string, peso float64, i int, j int) {

	existe := false
	if l.inicio == nil {
		l.inicio = &Nodo2{TipoP: TipoP, Peso: peso, id: j}
		l.fin = l.inicio
	} else {
		inicio := l.inicio
		for inicio.Siguiente != nil {
			if inicio.TipoP == TipoP {
				existe = true
			}
			inicio = inicio.Siguiente
		}
		if existe == false {
			nuevo := &Nodo2{TipoP: TipoP, Peso: peso, id: j}
			fin := l.fin
			fin.Siguiente = nuevo
			fin.Siguiente.Anterior = fin
			l.fin = nuevo
		} else {
			fmt.Println("Ya existe ese enlace")
		}
	}
	matrizAd[i][j] = peso
}
func (l *ListaDoble) Enlazar(TipoA string, TipoB string, peso float64) {
	origen := l.getVertice(TipoA)
	destino := l.getVertice(TipoB)
	if origen == nil || destino == nil {
		fmt.Println("No existe el vertice")
		return
	}

	origen.Enlaces.InsertarEnlace(destino.TipoProd, peso, origen.id, destino.id)
	destino.Enlaces.InsertarEnlace(origen.TipoProd, peso, destino.id, origen.id)
}

func (l *ListaDoble) corto() {
	inicio := l.Inicio
	dot := "digraph G {\n "
	for inicio != nil {
		dot += inicio.TipoProd + "\n"
		temp := inicio.Enlaces.inicio
		for temp != nil {

			//dot += inicio.TipoProd + "->" + temp.Enlace.TipoProd + "[label=\"" + strconv.Itoa(temp.Enlace.Peso) + "\"]\n"
			temp = temp.Siguiente
		}
		inicio = inicio.Siguiente
	}
}

func minimo(pasado []int, matriz [][]float64, len int, actual int) int {
	//aux := mucho
	//regreso := 0
	for i := 0; i < len; i++ {
		//if pasado[i] != 1 && (aux > matriz[actual][i]) {
		//aux = matriz[actual][i]
		//regreso = i

	}
	return 0
}

func (l *ListaDoble) Dikstra(inicioRobot string, inicioPed string, len int, total *float64) {
	inicio := l.Inicio
	indexIn := 0
	indexFin := 0
	encontrado1 := false
	encontrado2 := false
	Recorrido := make([]int, len)
	distancia := *total
	primerNodo := false
	siguiente := 0
	for inicio != nil {
		if inicio.TipoProd == inicioRobot {
			indexIn = inicio.id
			encontrado1 = true

		}
		if inicio.TipoProd == inicioPed {
			indexFin = inicio.id
			encontrado2 = true
		}
		if encontrado1 == true && encontrado2 == true {
			break
		}

		inicio = inicio.Siguiente

	}

	aux := 0.0
	fmt.Println(matrizAd)
	for i := 0; i < len; i++ {

		if Recorrido[i] != 1 {
			if primerNodo == false && i != indexIn && matrizAd[indexIn][i] != 0 {

				aux = distancia + matrizAd[indexIn][i]
				primerNodo = true
				siguiente = i

			} else {
				if primerNodo == true && matrizAd[indexIn][i] != 0 {
					if distancia+matrizAd[indexIn][i] < aux {
						aux = distancia + matrizAd[indexIn][i]
						siguiente = i
					}

				}
			}

		}
		if i+1 == len {
			Recorrido[indexIn] = 1
			distancia = aux
			aux = 0
			if indexIn == 34 {
				fmt.Println(matrizAd[34])
			}
			indexIn = siguiente
			i = -1
			Recorrido[siguiente] = 1
			primerNodo = false
			if siguiente == indexFin {

				*total = distancia
				break
			}

		}

	}

}

func Existe(p Pila, tipo string) bool {
	aux := p
	for i := 0; i < len(p.TipoProd); i++ {
		sacado := aux.Pop()
		if sacado == tipo {
			return true
		}
	}
	return false
}

func (l *ListaDoble) Graficar() {
	aux := Pila{}
	var sc strings.Builder
	dot := "digraph G{\n"
	fmt.Fprintf(&sc, "digraph G{\n")
	inicio := l.Inicio
	for inicio != nil {

		if Existe(aux, inicio.TipoProd) == false {
			aux.Push(inicio.TipoProd)
			fmt.Fprintf(&sc, "node%p[label=\"%v\"]\n", inicio.id, inicio.TipoProd)
			dot += "node" + strconv.Itoa(inicio.id) + "[label=\"" + inicio.TipoProd + "\"]\n"
		}
		inicio2 := inicio.Enlaces.inicio
		for inicio2 != nil {

			fmt.Fprintf(&sc, "node%p->node%p\n", inicio.id, inicio2.id)
			dot += "node" + strconv.Itoa(inicio.id) + "->node" + strconv.Itoa(inicio2.id) + "\n"
			if Existe(aux, inicio2.TipoP) == false {
				aux.Push(inicio2.TipoP)
				fmt.Fprintf(&sc, "node%p[label=\"%v\"]\n", inicio2.id, inicio2.TipoP)
				dot += "node" + strconv.Itoa(inicio2.id) + "[label=\"" + inicio2.TipoP + "\"]\n"
			}
			inicio2 = inicio2.Siguiente
		}
		inicio = inicio.Siguiente
	}
	fmt.Fprintf(&sc, "}")
	dot += "}\n"
	GeneraGrafo(dot)
	fmt.Println(dot)
}

func GeneraGrafo(dot string) {
	err := ioutil.WriteFile("Grafo.dot", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", "Grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("Grafo.png", cmd, os.FileMode(mode))
}

func GenerarTabla(dot string) {
	err := ioutil.WriteFile("Grafo.dot", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", "Grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("Grafo.png", cmd, os.FileMode(mode))
}
