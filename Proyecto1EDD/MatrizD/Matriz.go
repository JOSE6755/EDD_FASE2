package MatrizD

import (
	"fmt"
	"reflect"
	"strconv"
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
	datos2    []infor
}

type NodoV struct {
	ESTE      interface{}
	NORTE     interface{}
	SUR       interface{}
	OESTE     interface{}
	Categoria string
	Pos       int
}
type infor struct {
}

type NodoH struct {
	ESTE  interface{}
	NORTE interface{}
	SUR   interface{}
	OESTE interface{}
	dia   int
	Pos   int
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
	Cantidad int
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
	Cantidad int
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
			nuevo.ESTE = temp
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
	} else {
		inicio := l.inico
		for inicio.Siguiente != nil {
			inicio = inicio.Siguiente
		}

		inicio.Siguiente = nuevo
	}
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

func (m *Matriz) Graficar() string {
	dot := "digraph Sparce_Matrix {\nnode [shape=box]\nMT[label=\"Matrix\",width=1.5,style=filled,fillcolor=firebrick1,group=1];\ne0[ shape = point, width = 0 ];\ne1[ shape = point, width = 0 ];\n"
	var aux interface{} = m.CabV

	contV := 0
	relaciones := ""
	relfin := "MT->V0\n"
	for aux != nil {
		dot += "V" + strconv.Itoa(contV) + "[label=\"" + aux.(*NodoV).Categoria + "\"" + "width = 1.5 style = filled, fillcolor = bisque1, group = 1];\n"
		if aux.(*NodoV).SUR != nil {
			relaciones += "V" + strconv.Itoa(contV) + "-> V" + strconv.Itoa(contV+1) + "\n"
			relaciones += "V" + strconv.Itoa(contV+1) + "-> V" + strconv.Itoa(contV) + "\n"
		}
		contV++
		aux = aux.(*NodoV).SUR
	}
	dot += relaciones
	dot += relfin
	relaciones = ""
	var aux2 interface{} = m.CabH
	contH := 0
	contG := 2
	relfin += "MT->H0\n"
	resame := "{rank=same; MT;"
	for aux2 != nil {
		dot += "H" + strconv.Itoa(contH) + "[label=\"" + strconv.Itoa(aux2.(*NodoH).dia) + "\"" + "width = 1.5 style = filled, fillcolor = lightskyblue, group =" + strconv.Itoa(contG) + "];\n"
		if aux2.(*NodoH).ESTE != nil {
			relaciones += "H" + strconv.Itoa(contH) + "-> H" + strconv.Itoa(contH+1) + "\n"
			relaciones += "H" + strconv.Itoa(contH+1) + "-> H" + strconv.Itoa(contH) + "\n"

		}
		resame += "H" + strconv.Itoa(contH) + ";"
		contH++
		contG++
		aux2 = aux2.(*NodoH).ESTE
	}
	resame += "}\n"
	dot += relfin
	dot += relaciones
	dot += resame
	aux = m.CabV
	aux2 = m.CabH
	contG = 2
	for aux2 != nil {
		temp := aux2.(*NodoH).SUR
		//temp2 := temp.(*NodoInfo)
		for temp != nil {
			dot += "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "[label=\"Pedidos\" width=1.5,group=" + strconv.Itoa(contG) + "];\n"
			temp = temp.(*NodoInfo).SUR

		}
		contG++
		aux2 = aux2.(*NodoH).ESTE
	}
	aux2 = m.CabH
	aux = m.CabV
	contV = 0
	contH = 0

	for aux != nil {
		resame = "{rank=same "
		temp := aux.(*NodoV).ESTE
		//temp2 := temp.(*NodoInfo)
		//temp3 := temp2.ESTE.(*NodoInfo)
		if temp != nil {
			resame += "V" + strconv.Itoa(contV) + ";"
			dot += "V" + strconv.Itoa(contV) + "->" + "\"" + fmt.Sprintf("%p", *&(temp)) + "\"" + "\n"
			dot += "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "->" + "V" + strconv.Itoa(contV) + "\n"

		}
		for temp != nil {
			if temp.(*NodoInfo).ESTE != nil {
				dot += "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "->" + "\"" + fmt.Sprintf("%p", *&temp.(*NodoInfo).ESTE) + "\"" + "\n"
				dot += "\"" + fmt.Sprintf("%p", *&temp.(*NodoInfo).ESTE) + "\"" + "->" + "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "\n"

			}
			resame += "\"" + fmt.Sprintf("%p", *&temp) + "\"" + ";"
			temp = temp.(*NodoInfo).ESTE
		}
		resame += "}\n"
		contV++
		dot += resame
		aux = aux.(*NodoV).SUR
	}

	for aux2 != nil {
		temp := aux2.(*NodoH).SUR
		//temp2 := temp.(*NodoInfo)
		//temp3 := temp2.SUR.(*NodoInfo)
		if temp != nil {
			dot += "H" + strconv.Itoa(contH) + "->" + "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "\n"
			dot += "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "->" + "H" + strconv.Itoa(contH) + "\n"
		}
		for temp != nil {
			if temp.(*NodoInfo).SUR != nil {
				dot += "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "->" + "\"" + fmt.Sprintf("%p", *&temp.(*NodoInfo).SUR) + "\"" + "\n"
				dot += "\"" + fmt.Sprintf("%p", *&temp.(*NodoInfo).SUR) + "\"" + "->" + "\"" + fmt.Sprintf("%p", *&temp) + "\"" + "\n"
			}
			temp = temp.(*NodoInfo).SUR
		}
		contH++
		aux2 = aux2.(*NodoH).ESTE
	}
	fmt.Println(dot)

	return dot
}
