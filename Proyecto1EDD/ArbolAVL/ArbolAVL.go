package ArbolAVL

//Si me da tiempo cambio el codigo a mi manera
import (
	"fmt"
	"strconv"
)

type Nodo struct {
	Nombre      string  `json:Nombre`
	Codigo      int     `json:Codigo`
	Descripcion string  `json:Descripcion`
	Precio      float64 `json:Precio`
	Cantidad    int     `json:Cantidad`
	imagen      string
	altura      int
	izq         *Nodo
	der         *Nodo
}

type Arbolavl struct {
	Raiz *Nodo
}

type nodo struct {
	Arbol        *Arbolavl
	Tienda       string
	Departamento string
	Calificacion int
	siguiente    *nodo
}

type Lista_simple struct {
	inicio   *nodo
	cantidad int
}

func rotarII(n *Nodo, n1 *Nodo) *Nodo {
	n.izq = n1.der
	n1.der = n
	if n1.altura == -1 {
		n.altura = 0
		n1.altura = 0
	} else {
		n.altura = -1
		n1.altura = 1
	}
	return n1

}

func rotarDD(n *Nodo, n1 *Nodo) *Nodo {
	n.der = n1.izq
	n1.izq = n
	if n1.altura == 1 {
		n.altura = 0
		n1.altura = 0
	} else {
		n.altura = 1
		n1.altura = -1
	}
	return n1
}

func rotaDI(n *Nodo, n1 *Nodo) *Nodo {
	n2 := n1.izq
	n.der = n2.izq
	n2.izq = n
	n1.izq = n2.der
	n2.der = n1
	if n2.altura == 1 {
		n.altura = -1

	} else {
		n.altura = 0
	}
	if n2.altura == -1 {
		n1.altura = 1

	} else {
		n1.altura = 0
	}
	n2.altura = 0
	return n2
}

func rodaID(n *Nodo, n1 *Nodo) *Nodo {
	n2 := n1.der
	n.izq = n2.der
	n2.der = n
	n1.der = n2.izq
	n2.izq = n1
	if n2.altura == 1 {
		n1.altura = -1
	} else {
		n1.altura = 0
	}
	if n2.altura == -1 {
		n.altura = 1
	} else {
		n.altura = 0
	}
	n2.altura = 0
	return n2
}

func insertar(raiz *Nodo, nombre string, codigo int, des string, precio float64, cantidad int, imagen string, hc *bool) *Nodo {
	var n1 *Nodo
	if raiz == nil {
		raiz = &Nodo{Nombre: nombre, Codigo: codigo, Descripcion: des, Precio: precio, Cantidad: cantidad, imagen: imagen, altura: 0}
		*hc = true

	} else if codigo < raiz.Codigo {
		izq := insertar(raiz.izq, nombre, codigo, des, precio, cantidad, imagen, hc)
		raiz.izq = izq
		if *hc == true {
			switch raiz.altura {
			case 1:
				raiz.altura = 0
				*hc = false
				break
			case 0:
				raiz.altura = -1
				break
			case -1:
				n1 := raiz.izq
				if n1.altura == -1 {
					raiz = rotarII(raiz, n1)
				} else {
					raiz = rodaID(raiz, n1)
				}
				*hc = false
				break
			}
		}

	} else if codigo > raiz.Codigo {
		der := insertar(raiz.der, nombre, codigo, des, precio, cantidad, imagen, hc)
		raiz.der = der
		if *hc == true {
			switch raiz.altura {
			case 1:
				n1 = raiz.der
				if n1.altura == 1 {
					raiz = rotarDD(raiz, n1)
				} else {
					raiz = rotaDI(raiz, n1)
				}
				*hc = false
				break
			case 0:
				raiz.altura = 1
				break
			case -1:
				raiz.altura = 0
				*hc = false
			}
		}
	}
	return raiz
}

func (a *Arbolavl) Insertar(nombre string, codigo int, des string, precio float64, cantidad int, imagen string) {
	b := false
	c := &b
	a.Raiz = insertar(a.Raiz, nombre, codigo, des, precio, cantidad, imagen, c)
}

var tiendas string

func (n *Nodo) balancear() *Nodo {
	if n == nil {
		return n
	}
	n.realtura()
	FE := n.der.alt() - n.izq.alt()

	if FE == 2 {
		if n.der.izq.alt() > n.der.der.alt() {
			n.der = n.der.rotarDer()
		}
		return n.rotarIzq()

	} else if FE == -2 {
		if n.izq.der.alt() > n.izq.izq.alt() {
			n.izq = n.der.rotarIzq()
		}
		return n.rotarDer()
	}
	return n
}

func (n *Nodo) realtura() {
	n.altura = 1 + mayor(n.izq.alt(), n.der.alt())
}
func (n *Nodo) alt() int {
	if n == nil {
		return 0
	}
	return n.altura
}

func mayor(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *Nodo) rotarDer() *Nodo {
	auxi := n.izq
	n.izq = auxi.der
	auxi.der = n
	n.realtura()
	auxi.realtura()
	return auxi

}

func (n *Nodo) rotarIzq() *Nodo {
	auxi := n.der
	n.der = auxi.izq
	auxi.izq = n
	n.realtura()
	auxi.realtura()
	return auxi
}

func Mietras() {
	fmt.Println("Tar chingando mrd")

}

var l int

func (n *Nodo) DisplayNodesInOrder() {

	if n.izq != nil {

		n.izq.DisplayNodesInOrder()

	}

	pru[l] = append(pru[l], n.Nombre, strconv.Itoa(n.Codigo), n.Descripcion, strconv.Itoa(int(n.Precio)), strconv.Itoa(n.Cantidad), n.imagen)
	l++
	fmt.Println(l)

	if n.der != nil {

		n.der.DisplayNodesInOrder()
	}

}

var prec float64
var encon bool

func Getprec() float64 {
	return prec
}
func Getencon() bool {
	return encon
}
func Inserencon() {
	encon = false
}
func (n *Nodo) DisplayNodesInOrder2(codigo int, cantidad int) {

	if n.izq != nil {

		n.izq.DisplayNodesInOrder2(codigo, cantidad)

	}

	if n.Codigo == codigo {
		if cantidad < n.Cantidad {
			prec = n.Precio
			encon = true

		}
	}

	if n.der != nil {

		n.der.DisplayNodesInOrder2(codigo, cantidad)
	}

}

func (l *Lista_simple) Met(Arbol *Arbolavl, tienda string, departamento string, calificacion int) {
	nuevo := &nodo{Arbol: Arbol, Tienda: tienda, Departamento: departamento, Calificacion: calificacion}

	if l.inicio == nil {
		l.inicio = nuevo
	}
	aux := l.inicio

	for aux.siguiente != nil {
		aux = aux.siguiente
	}
	aux.siguiente = nuevo

}
func (l *Lista_simple) busc(nombre string, departamento string, calificacion int) *nodo {
	aux := l.inicio
	for aux != nil {
		if aux.Tienda == nombre && aux.Departamento == departamento && aux.Calificacion == calificacion {
			return aux
		}
		aux = aux.siguiente
	}
	return nil
}

var pru [][]string

func Matz(productos int) {
	pru = make([][]string, productos)
	l = 0

}
func Regres() [][]string {
	return pru
}
