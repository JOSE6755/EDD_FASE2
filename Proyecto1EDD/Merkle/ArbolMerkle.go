package Merkle

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoP struct {
	DPI       int
	Fecha     string
	numeroPed int
	hash      string
	derecha   *NodoP
	izquierda *NodoP
}

type Arbolito struct {
	Raiz *NodoP
}

func (n *NodoP) suma() string {

	if n.derecha != nil && n.izquierda != nil {
		datos := ""
		datos += strconv.Itoa(n.derecha.DPI) + n.derecha.Fecha + strconv.Itoa(n.derecha.numeroPed)
		return datos
	}
	return "-1"
}

func nuevoNodo(DPI int, fecha string, numero int, hash string, derecho *NodoP, izquierdo *NodoP) *NodoP {

	return &NodoP{DPI: DPI, Fecha: fecha, numeroPed: numero, derecha: derecho, izquierda: izquierdo, hash: hash}

}

func NuevoArbol() *Arbolito {
	return &Arbolito{}
}

func (a *Arbolito) Insertar(DPI int, fecha string, numero int) {
	hash := strconv.Itoa(DPI) + fecha + strconv.Itoa(numero)
	nuevo := nuevoNodo(DPI, fecha, numero, hash, nil, nil)

	if a.Raiz == nil {
		lista := list.New()
		lista.PushBack(nuevo)
		lista.PushBack(nuevoNodo(-1, "", 0, "-10", nil, nil))
		a.construccion(lista)
	} else {
		lista := a.lista()
		lista.PushBack(nuevo)
		a.construccion(lista)
	}
}

func (a *Arbolito) lista() *list.List {
	nueva := list.New()
	conseguirLista(nueva, a.Raiz.izquierda)
	conseguirLista(nueva, a.Raiz.derecha)
	return nueva
}

func conseguirLista(lista *list.List, actual *NodoP) {
	if actual != nil {
		conseguirLista(lista, actual.izquierda)
		if actual.derecha == nil && actual.DPI != -1 {
			lista.PushBack(actual)
		}
		conseguirLista(lista, actual.derecha)
	}
}

func (a *Arbolito) construccion(lista *list.List) {
	tamano := float64(lista.Len())
	cant := 1
	for (tamano / 2) > 1 {
		cant++
		tamano = tamano / 2
	}
	total := math.Pow(2, float64(cant))
	for lista.Len() < int(total) {
		lista.PushBack(nuevoNodo(-1, "", 0, "-10", nil, nil))
	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoP)
		nodo2 := segundo.Value.(*NodoP)
		hash := nodo1.hash + nodo2.hash
		fmt.Println(hash)
		nuevo := nuevoNodo(-1, "", 0, hash, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	a.Raiz = lista.Front().Value.(*NodoP)
}

func (a *Arbolito) codigo() {
	var dot strings.Builder
	fmt.Fprintf(&dot, "digraph G{\n")
	fmt.Fprintf(&dot, "node[shape=\"record\"];\n")
	if a.Raiz != nil {
		fmt.Println(a.Raiz)
		fmt.Fprintf(&dot, "node%p[label=\"<f0>|<f1>%v|<f2>\"];\n", &(*a.Raiz), a.Raiz.hash)
		a.generacion(&dot, a.Raiz, a.Raiz.izquierda, true)
		a.generacion(&dot, a.Raiz, a.Raiz.derecha, false)
	}
	fmt.Fprintf(&dot, "}\n")
	fmt.Println(dot.String())
}

func (a *Arbolito) generacion(dot *strings.Builder, padre *NodoP, actual *NodoP, izquierda bool) {
	if actual != nil {
		if actual.hash != "" {
			fmt.Fprintf(dot, "node%p[label=\"<f0>|<f1>%v|<f2>\"];\n", &(*actual), actual.hash)
		} else {
			fmt.Fprintf(dot, "node%p[label=\"<f0>|<f1>%v|<f2>\"];\n", &(*actual), actual.hash)

		}
		if izquierda {
			fmt.Fprintf(dot, "node%p:f0->node%p:f1\n", &(*padre), &(*actual))
		} else {
			fmt.Fprintf(dot, "node%p:f2->node%p:f1\n", &(*padre), &(*actual))
		}
		a.generacion(dot, actual, actual.izquierda, true)
		a.generacion(dot, actual, actual.derecha, false)
	}
}
