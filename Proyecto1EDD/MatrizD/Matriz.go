package MatrizD

import (
	"reflect"
)

type NodoInfo struct {
	ESTE      interface{}
	NORTE     interface{}
	SUR       interface{}
	OESTE     interface{}
	Cantida   int
	Producto  int
	Precio    float64
	Dia       int
	Categoria string
}

type NodoV struct {
	ESTE      interface{}
	NORTE     interface{}
	SUR       interface{}
	OESTE     interface{}
	Categoria string
}

type NodoH struct {
	ESTE  interface{}
	NORTE interface{}
	SUR   interface{}
	OESTE interface{}
	dia   int
}

type Matriz struct {
	CabH *NodoH
	CabV *NodoV
}

type Nodoaño struct {
	Año       int
	Lista     *Lista_doble
	Siguiente *Nodoaño
}

type Lista_Simple struct {
	inico    *Nodoaño
	cantidad int
}

type NodoM struct {
	Mes       int
	pedidosM  *Matriz
	siguiente *NodoM
	anterior  *NodoM
}

type Lista_doble struct {
	inicio   *NodoM
	fin      *NodoM
	cantidad int
}

type Pedidos struct {
	Ped []DP `json:"Pedidos"`
}

type DP struct {
	Fecha        string `json:"Fecha"`
	Tienda       string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int    `json:"Calificacion"`
	Productos    []Prod `json:"Productos"`
}

type Prod struct {
	Codigo int `json:"Codigo"`
}

var ped Pedidos

func (m *Matriz) getV(categoria string) interface{} {
	if m.CabV == nil {
		return nil
	}
	var aux interface{} = m.CabV
	for aux != nil {
		if aux.(*NodoV).Categoria == categoria {
			return aux
		}
		aux = aux.(*NodoV).SUR
	}
	return nil
}

func (m *Matriz) getH(dia int) interface{} {
	if m.CabV == nil {
		return nil
	}
	var aux interface{} = m.CabH
	for aux != nil {
		if aux.(*NodoH).dia == dia {
			return aux
		}
		aux = aux.(*NodoH).ESTE
	}
	return nil
}

func (m *Matriz) crearH(dia int) *NodoH {
	if m.CabH == nil {
		nueva := &NodoH{dia: dia, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
		m.CabH = nueva
		return nueva
	}
	var auxi interface{} = m.CabH
	if dia < auxi.(*NodoH).dia {
		nueva := &NodoH{dia: dia, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
		nueva.ESTE = m.CabH
		m.CabH.OESTE = nueva
		m.CabH = nueva
		return nueva

	}
	for auxi.(*NodoH).SUR != nil {
		if dia > auxi.(*NodoH).dia && dia < auxi.(*NodoH).SUR.(*NodoH).dia {
			nueva := &NodoH{dia: dia, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
			temp := auxi.(*NodoH).ESTE
			temp.(*NodoH).OESTE = nueva
			nueva.ESTE = temp
			auxi.(*NodoH).ESTE = nueva
			nueva.OESTE = auxi
			return nueva

		}
		auxi = auxi.(*NodoH).ESTE
	}
	nueva := &NodoH{dia: dia, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
	auxi.(*NodoH).ESTE = nueva
	nueva.OESTE = auxi
	return nueva

}

func (m *Matriz) crearV(categoria string) *NodoV {
	if m.CabV == nil {
		nueva := &NodoV{Categoria: categoria, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
		m.CabV = nueva
		return nueva
	}
	var auxi interface{} = m.CabV
	if categoria < auxi.(*NodoV).Categoria {
		nueva := &NodoV{Categoria: categoria, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
		nueva.SUR = m.CabV
		m.CabV.NORTE = nueva
		m.CabV = nueva
		return nueva

	}
	for auxi.(*NodoV).SUR != nil {
		if categoria > auxi.(*NodoV).Categoria && categoria < auxi.(*NodoV).SUR.(*NodoV).Categoria {
			nueva := &NodoV{Categoria: categoria, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
			temp := auxi.(*NodoV).SUR
			temp.(*NodoV).NORTE = nueva
			nueva.SUR = temp
			auxi.(*NodoV).SUR = nueva
			nueva.NORTE = auxi
			return nueva

		}
		auxi = auxi.(*NodoV).SUR
	}
	nueva := &NodoV{Categoria: categoria, ESTE: nil, NORTE: nil, SUR: nil, OESTE: nil}
	auxi.(*NodoV).SUR = nueva
	nueva.NORTE = auxi
	return nueva

}

func (m *Matriz) getUltimoV(cab *NodoH, dia int) interface{} {
	if cab.SUR == nil {
		return cab
	}
	aux := cab.SUR
	if dia <= aux.(*NodoInfo).Dia {
		return cab
	}

	for aux.(*NodoInfo).SUR != nil {
		if dia > aux.(*NodoInfo).Dia && dia <= aux.(*NodoInfo).SUR.(*NodoInfo).Dia {
			return aux
		}
		aux = aux.(*NodoInfo).SUR
	}
	if dia <= aux.(*NodoInfo).Dia {
		return aux.(*NodoInfo).NORTE
	}
	return aux
}

func (m *Matriz) getUltimoH(cab *NodoV, categoria string) interface{} {
	if cab.ESTE == nil {
		return cab
	}
	aux := cab.ESTE
	if categoria <= aux.(*NodoInfo).Categoria {
		return cab
	}
	for aux.(*NodoInfo).ESTE != nil {
		if categoria > aux.(*NodoInfo).Categoria && categoria <= aux.(*NodoInfo).ESTE.(*NodoInfo).Categoria {
			return aux
		}
		aux = aux.(*NodoInfo).ESTE
	}
	if categoria <= aux.(*NodoInfo).Categoria {
		return aux.(*NodoInfo).OESTE
	}
	return aux
}

func (m *Matriz) Inser(nuevo *NodoInfo) {
	vert := m.getV(nuevo.Categoria)
	hor := m.getH(nuevo.Dia)
	if vert == nil {
		vert = m.crearV(nuevo.Categoria)
	}
	if hor == nil {
		hor = m.crearH(nuevo.Dia)
	}

	der := m.getUltimoH(vert.(*NodoV), nuevo.Categoria)
	sup := m.getUltimoV(hor.(*NodoH), nuevo.Dia)

	if reflect.TypeOf(der).String() == "*MatrizD.NodoInfo" {
		if der.(*NodoInfo).ESTE == nil {
			der.(*NodoInfo).ESTE = nuevo
			nuevo.OESTE = der
		} else {
			temp := der.(*NodoInfo).ESTE
			der.(*NodoInfo).ESTE = nuevo
			nuevo.OESTE = der
			temp.(*NodoInfo).OESTE = nuevo
			nuevo.ESTE = temp
		}

	} else {
		if der.(*NodoV).ESTE == nil {
			der.(*NodoV).ESTE = nuevo
			nuevo.OESTE = der
		} else {
			temp := der.(*NodoV).ESTE
			der.(*NodoV).ESTE = nuevo
			nuevo.OESTE = der
			temp.(*NodoInfo).OESTE = nuevo
			nuevo.OESTE = temp
		}
	}

	if reflect.TypeOf(sup).String() == "*MatrizD.NodoInfo" {
		if sup.(*NodoInfo).SUR == nil {
			sup.(*NodoInfo).SUR = nuevo
			nuevo.NORTE = sup
		} else {
			temp := sup.(*NodoInfo).SUR
			sup.(*NodoInfo).SUR = nuevo
			nuevo.NORTE = sup
			temp.(*NodoInfo).NORTE = nuevo
			nuevo.SUR = temp
		}

	} else {
		if sup.(*NodoH).SUR == nil {
			sup.(*NodoH).SUR = nuevo
			nuevo.NORTE = sup
		} else {
			temp := sup.(*NodoH).SUR
			sup.(*NodoH).SUR = nuevo
			nuevo.NORTE = sup
			temp.(*NodoInfo).NORTE = nuevo
			nuevo.SUR = temp
		}
	}

}

var precio float64

func getPrecio() float64 {
	return precio
}
func SetPrecio(pre float64) {
	precio = pre
}

func (l *Lista_Simple) InserSimple(doble *Lista_doble, año int) {
	nuevo := &Nodoaño{Año: año, Lista: doble}
	if l.inico == nil {
		l.inico = nuevo
	}
	inicio := l.inico
	for inicio.Siguiente != nil {
		inicio = inicio.Siguiente
	}
	inicio.Siguiente = nuevo
}

func (l *Lista_doble) InserDoble(m *Matriz, mes int) {
	nuevo := &NodoM{Mes: mes, pedidosM: m}
	if l.inicio == nil {
		l.inicio = nuevo
	}
	inicio := l.inicio
	for inicio.siguiente != nil {
		inicio = inicio.siguiente
	}
	inicio.siguiente = nuevo
}
