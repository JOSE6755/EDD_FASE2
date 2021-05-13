package Merkle

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NodoInv struct {
	Nombre    string
	Codigo    int
	Precio    float64 `json:"Precio"`
	hash      string
	derecha   *NodoInv
	izquierda *NodoInv
}

type ArbolitoInv struct {
	Raiz *NodoInv
}

/*func (n *NodoInv) suma() string {

	if n.derecha != nil && n.izquierda != nil {
		datos := ""
		datos += strconv.Itoa(n.derecha.DPI) + n.derecha.Fecha + strconv.Itoa(n.derecha.numeroPed)
		return datos
	}
	return "-1"
}
*/

func NuevoNodoInv(Nombre string, Codigo int, Precio float64, hash string, derecho *NodoInv, izquierdo *NodoInv) *NodoInv {

	return &NodoInv{Nombre: Nombre, Codigo: Codigo, Precio: Precio, derecha: derecho, izquierda: izquierdo, hash: hash}

}

func NuevoArbolInv() *ArbolitoInv {
	return &ArbolitoInv{}
}

func (a *ArbolitoInv) InsertarInv(Nombre string, Codigo int, Precio float64) {
	hash := Nombre + strconv.Itoa(Codigo) + fmt.Sprintf("%f", Precio)
	nuevo := NuevoNodoInv(Nombre, Codigo, Precio, hash, nil, nil)

	if a.Raiz == nil {
		lista := list.New()
		lista.PushBack(nuevo)
		lista.PushBack(NuevoNodoInv("", -1, 0, "-10", nil, nil))
		a.ConstruccionInv(lista)
	} else {
		lista := a.ListaInv()
		lista.PushBack(nuevo)
		a.ConstruccionInv(lista)
	}
}

func (a *ArbolitoInv) ListaInv() *list.List {
	nueva := list.New()
	ConseguirListaInv(nueva, a.Raiz.izquierda)
	ConseguirListaInv(nueva, a.Raiz.derecha)
	return nueva
}

func ConseguirListaInv(lista *list.List, actual *NodoInv) {
	if actual != nil {
		ConseguirListaInv(lista, actual.izquierda)
		if actual.derecha == nil && actual.Codigo != -1 {
			lista.PushBack(actual)
		}
		ConseguirListaInv(lista, actual.derecha)
	}
}

func (a *ArbolitoInv) ConstruccionInv(lista *list.List) {
	tamano := float64(lista.Len())
	cant := 1
	for (tamano / 2) > 1 {
		cant++
		tamano = tamano / 2
	}
	total := math.Pow(2, float64(cant))
	for lista.Len() < int(total) {
		lista.PushBack(NuevoNodoInv("", -1, 0, "-10", nil, nil))
	}
	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()
		lista.Remove(primero)
		lista.Remove(segundo)
		nodo1 := primero.Value.(*NodoInv)
		nodo2 := segundo.Value.(*NodoInv)
		hash := nodo1.hash + nodo2.hash
		fmt.Println(hash)
		nuevo := NuevoNodoInv("", -1, 0, hash, nodo2, nodo1)
		lista.PushBack(nuevo)
	}
	a.Raiz = lista.Front().Value.(*NodoInv)
}

func (a *ArbolitoInv) codigo() {
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

func (a *ArbolitoInv) generacion(dot *strings.Builder, padre *NodoInv, actual *NodoInv, izquierda bool) {
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
