package Merkle

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoUs struct {
	DPI       int
	Pass      string
	Correo    string
	hash      string
	derecha   *NodoUs
	izquierda *NodoUs
}

type ArbolitoUs struct {
	Raiz *NodoUs
}

func NuevoNodoUs(DPI int, pass string, Correo string, hash string, derecho *NodoUs, izquierdo *NodoUs) *NodoUs {

	return &NodoUs{DPI: DPI, Pass: pass, Correo: Correo, derecha: derecho, izquierda: izquierdo, hash: hash}

}

func NuevoArbolUs() *ArbolitoUs {
	return &ArbolitoUs{}
}

func (a *ArbolitoUs) InsertarUs(DPI int, pass string, correo string) {
	hash := strconv.Itoa(DPI) + pass + correo
	nuevo := NuevoNodoUs(DPI, pass, correo, hash, nil, nil)

	if a.Raiz == nil {
		lista := list.New()
		lista.PushBack(nuevo)
		lista.PushBack(NuevoNodoUs(-1, "", "", "-10", nil, nil))
		a.construccion(lista)
	} else {
		lista := a.lista()
		lista.PushBack(nuevo)
		a.construccion(lista)
	}
}

func (a *ArbolitoUs) lista() *list.List {
	nueva := list.New()
	conseguirListaUs(nueva, a.Raiz.izquierda)
	conseguirListaUs(nueva, a.Raiz.derecha)
	return nueva
}

func conseguirListaUs(lista *list.List, actual *NodoUs) {
	if actual != nil {
		conseguirListaUs(lista, actual.izquierda)
		if actual.derecha == nil && actual.DPI != -1 {
			lista.PushBack(actual)
		}
		conseguirListaUs(lista, actual.derecha)
	}
}

func (a *ArbolitoUs) construccion(lista *list.List) {
	tamano := float64(lista.Len())
	cant := 1
	for (tamano / 2) > 1 {
		cant++
		tamano = tamano / 2
	}
	total := math.Pow(2, float64(cant))
	for lista.Len() < int(total) {
		lista.PushBack(NuevoNodoUs(-1, "", "", "-10", nil, nil))
	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoUs)
		nodo2 := segundo.Value.(*NodoUs)
		hash := nodo1.hash + nodo2.hash
		fmt.Println(hash)
		nuevo := NuevoNodoUs(-1, "", "", hash, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	a.Raiz = lista.Front().Value.(*NodoUs)
}

func (a *ArbolitoUs) CodigoUs() {
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

func (a *ArbolitoUs) generacion(dot *strings.Builder, padre *NodoUs, actual *NodoUs, izquierda bool) {
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
