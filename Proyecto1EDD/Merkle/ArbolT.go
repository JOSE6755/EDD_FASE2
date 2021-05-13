package Merkle

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoT struct {
	Nombre       string
	Contacto     string
	Calificacion int
	hash         string
	derecha      *NodoT
	izquierda    *NodoT
}

type ArbolitoT struct {
	Raiz *NodoT
}

func NuevoNodoT(Nombre string, Contacto string, Calificacion int, hash string, derecho *NodoT, izquierdo *NodoT) *NodoT {

	return &NodoT{Nombre: Nombre, Contacto: Contacto, Calificacion: Calificacion, derecha: derecho, izquierda: izquierdo, hash: hash}

}

func NuevoArbolT() *ArbolitoT {
	return &ArbolitoT{}
}

func (a *ArbolitoT) Insertar(Nombre string, Contacto string, Calificacion int) {
	hash := Nombre + Contacto + strconv.Itoa(Calificacion)
	nuevo := NuevoNodoT(Nombre, Contacto, Calificacion, hash, nil, nil)

	if a.Raiz == nil {
		lista := list.New()
		lista.PushBack(nuevo)
		lista.PushBack(NuevoNodoT("", "", -1, "-10", nil, nil))
		a.construccion(lista)
	} else {
		lista := a.lista()
		lista.PushBack(nuevo)
		a.construccion(lista)
	}
}

func (a *ArbolitoT) lista() *list.List {
	nueva := list.New()
	ConseguirListaT(nueva, a.Raiz.izquierda)
	ConseguirListaT(nueva, a.Raiz.derecha)
	return nueva
}

func ConseguirListaT(lista *list.List, actual *NodoT) {
	if actual != nil {
		ConseguirListaT(lista, actual.izquierda)
		if actual.derecha == nil && actual.Calificacion != -1 {
			lista.PushBack(actual)
		}
		ConseguirListaT(lista, actual.derecha)
	}
}

func (a *ArbolitoT) construccion(lista *list.List) {
	tamano := float64(lista.Len())
	cant := 1
	for (tamano / 2) > 1 {
		cant++
		tamano = tamano / 2
	}
	total := math.Pow(2, float64(cant))
	for lista.Len() < int(total) {
		lista.PushBack(NuevoNodoT("", "", -1, "-10", nil, nil))
	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoT)
		nodo2 := segundo.Value.(*NodoT)
		hash := nodo1.hash + nodo2.hash
		fmt.Println(hash)
		nuevo := NuevoNodoT("", "", -1, hash, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	a.Raiz = lista.Front().Value.(*NodoT)
}

func (a *ArbolitoT) CodigoT() {
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

func (a *ArbolitoT) generacion(dot *strings.Builder, padre *NodoT, actual *NodoT, izquierda bool) {
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
