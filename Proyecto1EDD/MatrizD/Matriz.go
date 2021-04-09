package MatrizD

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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
}
type infor struct {
	Cantida  int
	Producto int
	Precio   float64
}

type NodoH struct {
	ESTE  interface{}
	NORTE interface{}
	SUR   interface{}
	OESTE interface{}
	dia   int
}
type datos struct {
	Nodos []NodoInfo
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
type Infoaño struct {
	Datos []Listaño `json:"Datos"`
}
type Listaño struct {
	Año   int   `json:"Año"`
	Meses []int `json:"Meses"`
}
type Imagenes struct {
	Nombre string `json:"Nombre"`
	Año    int    `json:"Año"`
	Mes    int    `json:"Mes"`
	Tipo   string `json:"Tipo"`
}

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

func (m *Matriz) getUltimoV(cab *NodoH, categoria string, n *bool) interface{} {
	if cab.SUR == nil {
		return cab
	}
	aux := cab.SUR
	if categoria <= aux.(*NodoInfo).Categoria {
		*n = true
		return cab
	}

	for aux.(*NodoInfo).SUR != nil {
		if categoria > aux.(*NodoInfo).Categoria && categoria <= aux.(*NodoInfo).SUR.(*NodoInfo).Categoria {
			return aux
		}
		aux = aux.(*NodoInfo).SUR
	}
	if categoria <= aux.(*NodoInfo).Categoria {
		return aux.(*NodoInfo).NORTE
	}
	return aux
}

func (m *Matriz) getUltimoH(cab *NodoV, dia int, n *bool) interface{} {
	if cab.ESTE == nil {
		return cab
	}
	aux := cab.ESTE
	if dia <= aux.(*NodoInfo).Dia {
		*n = true
		return cab
	}
	for aux.(*NodoInfo).ESTE != nil {
		if dia > aux.(*NodoInfo).Dia && dia <= aux.(*NodoInfo).ESTE.(*NodoInfo).Dia {
			return aux
		}
		aux = aux.(*NodoInfo).ESTE
	}
	if dia <= aux.(*NodoInfo).Dia {
		return aux.(*NodoInfo).OESTE
	}
	return aux
}

func (m *Matriz) Inser(nuevo *NodoInfo) {
	existe := false
	vert := m.getV(nuevo.Categoria)
	hor := m.getH(nuevo.Dia)

	if vert != nil && hor != nil {
		der := m.getUltimoH(vert.(*NodoV), nuevo.Dia, &existe)
		sup := m.getUltimoV(hor.(*NodoH), nuevo.Categoria, &existe)
		if (reflect.TypeOf(der).String() == "*MatrizD.NodoInfo" && reflect.TypeOf(sup).String() == "*MatrizD.NodoInfo") || existe == true {
			nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
			aux := der.(*NodoV).ESTE

			aux.(*NodoInfo).datos2 = append(aux.(*NodoInfo).datos2, nuevo2)
		}
	} else {
		if vert == nil {
			vert = m.crearV(nuevo.Categoria)
		}
		if hor == nil {
			hor = m.crearH(nuevo.Dia)
		}
		der := m.getUltimoH(vert.(*NodoV), nuevo.Dia, &existe)
		sup := m.getUltimoV(hor.(*NodoH), nuevo.Categoria, &existe)
		if reflect.TypeOf(der).String() == "*MatrizD.NodoInfo" {
			if der.(*NodoInfo).ESTE == nil {
				der.(*NodoInfo).ESTE = nuevo
				nuevo.OESTE = der
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
			} else {
				temp := der.(*NodoInfo).ESTE
				der.(*NodoInfo).ESTE = nuevo
				nuevo.OESTE = der
				temp.(*NodoInfo).OESTE = nuevo
				nuevo.ESTE = temp
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
			}

		} else {
			if der.(*NodoV).ESTE == nil {
				der.(*NodoV).ESTE = nuevo
				nuevo.OESTE = der
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
			} else {
				temp := der.(*NodoV).ESTE
				der.(*NodoV).ESTE = nuevo
				nuevo.OESTE = der
				temp.(*NodoInfo).OESTE = nuevo
				nuevo.ESTE = temp
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
			}
		}

		if reflect.TypeOf(sup).String() == "*MatrizD.NodoInfo" {
			if sup.(*NodoInfo).SUR == nil {
				sup.(*NodoInfo).SUR = nuevo
				nuevo.NORTE = sup
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
			} else {
				temp := sup.(*NodoInfo).SUR
				sup.(*NodoInfo).SUR = nuevo
				nuevo.NORTE = sup
				temp.(*NodoInfo).NORTE = nuevo
				nuevo.SUR = temp
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
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
				nuevo2 := infor{Cantida: nuevo.Cantida, Producto: nuevo.Producto, Precio: nuevo.Precio}
				nuevo.datos2 = append(nuevo.datos2, nuevo2)
			}
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
		l.Cantidad++
	} else {
		inicio := l.inico
		for inicio.Siguiente != nil {
			inicio = inicio.Siguiente
		}

		inicio.Siguiente = nuevo
		l.Cantidad++
	}
}
func (l *Lista_Simple) Esnul() bool {
	if l.inico == nil {
		return true
	}
	return false
}
func (l *Lista_Simple) Nuevomes(nuevaM *Matriz, mes int) {
	inicio := l.inico
	for inicio != nil {
		inicio = inicio.Siguiente
	}
	inicio.Lista.InserDoble(nuevaM, mes)
}
func (l *Lista_Simple) Buscar(dia int, año int, mes int, nuevo *NodoInfo, m *bool, nombre string) bool {
	inicio := l.inico
	encontrado := false
	for inicio != nil {
		if inicio.Año == año {
			inicio.Lista.buscar(nombre, dia, mes, año, nuevo, m, l)
			encontrado = true
		}
		inicio = inicio.Siguiente
	}
	return encontrado
}

func (l *Lista_doble) InserDoble(m *Matriz, mes int) {
	nuevo := &NodoM{Mes: mes, pedidosM: m}
	if l.inicio == nil {
		l.inicio = nuevo
		l.Cantidad++
	} else {
		inicio := l.inicio
		for inicio.siguiente != nil {
			inicio = inicio.siguiente
		}
		inicio.siguiente = nuevo
		l.Cantidad++
	}
}
func (l *Lista_doble) buscar(nombre string, dia int, mes int, año int, nuevo *NodoInfo, m *bool, lis *Lista_Simple) bool {
	inicio := l.inicio
	for inicio != nil {
		if inicio.Mes == mes {
			if nuevo != nil {
				inicio.pedidosM.Inser(nuevo)
				inicio.pedidosM.Graficar(dia, mes, año, nombre)
				lis.Añitos(nombre)
			}
			if m != nil {
				*m = true
			}
			return true
		}
		inicio = inicio.siguiente
	}
	*m = false
	return false
}

func (m *Matriz) Graficar(dia int, mes int, año int, nombre string) string {
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
	dot += "\n}"
	fmt.Println(dot)
	m.GraficarPedidos(dia, mes, año, nombre)
	err := ioutil.WriteFile(nombre+"-"+strconv.Itoa(mes)+"-"+strconv.Itoa(año)+".dot", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", nombre+"-"+strconv.Itoa(mes)+"-"+strconv.Itoa(año)+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre+"-"+strconv.Itoa(mes)+"-"+strconv.Itoa(año)+".png", cmd, os.FileMode(mode))

	return dot
}

func (m *Matriz) GraficarPedidos(dia int, mes int, año int, nombre string) string {
	dot := "digraph {\n    tbl[\n     shape=plaintext\n     label=<\n     <table border='0' cellborder='1' color='blue' cellspacing='0'>\n" + "<tr> <td>Departamento</td> <td> Dia </td> <td>Producto</td></tr>\n"
	var aux interface{} = m.CabV

	for aux != nil {

		dot += "<tr>\n<td>" + aux.(*NodoV).Categoria + "</td>\n"

		temp := aux.(*NodoV).ESTE
		temp2 := temp.(*NodoInfo).NORTE
		categoria := aux.(*NodoV).Categoria
		var dia string
		if reflect.TypeOf(temp2).String() == "*MatrizD.NodoH" {
			dot += "<td>" + strconv.Itoa(temp2.(*NodoH).dia) + "</td>\n"
			dia = strconv.Itoa(temp2.(*NodoH).dia)

		} else {
			for temp2 != nil {
				if reflect.TypeOf(temp2).String() == "*MatrizD.NodoH" {
					dot += "<td>" + strconv.Itoa(temp2.(*NodoH).dia) + "</td>\n"
					dia = strconv.Itoa(temp2.(*NodoH).dia)
					break

				}
				temp2 = temp2.(*NodoInfo).NORTE
			}
		}
		for temp != nil {
			for i := 0; i < len(temp.(*NodoInfo).datos2); i++ {
				if i == 0 {
					dot += "<td>Producto: " + strconv.Itoa(temp.(*NodoInfo).datos2[i].Producto) + "\nCantidad: " + strconv.Itoa(temp.(*NodoInfo).datos2[i].Cantida) + "\nPrecio: " + strconv.Itoa(int(temp.(*NodoInfo).datos2[i].Precio)) + "</td>\n</tr>"
				} else {
					dot += "<tr>\n<td>" + categoria + "</td>\n<td>" + dia + "</td>\n" + "<td>Codigo: " + strconv.Itoa(temp.(*NodoInfo).datos2[i].Producto) + " \nCantidad: " + strconv.Itoa(temp.(*NodoInfo).datos2[i].Cantida) + " \nPrecio: " + strconv.Itoa(int(temp.(*NodoInfo).datos2[i].Precio)) + "</td>\n</tr>\n"
				}
			}
			temp = temp.(*NodoInfo).ESTE
			if temp != nil {
				temp2 = temp.(*NodoInfo).NORTE
			}

			if reflect.TypeOf(temp2).String() == "*MatrizD.NodoH" && temp != nil {

				dia = strconv.Itoa(temp2.(*NodoH).dia)

			} else {
				if temp != nil {
					for temp2 != nil {
						if reflect.TypeOf(temp2).String() == "*MatrizD.NodoH" {

							dia = strconv.Itoa(temp2.(*NodoH).dia)
							break

						}
						temp2 = temp2.(*NodoInfo).NORTE
					}
				}
			}

		}
		aux = aux.(*NodoV).SUR

	}
	dot += "</table>\n    >];\n}"
	fmt.Println(dot)
	err := ioutil.WriteFile(nombre+"-"+strconv.Itoa(mes)+"-"+strconv.Itoa(año)+"-Pedidos.dot", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", nombre+"-"+strconv.Itoa(mes)+"-"+strconv.Itoa(año)+"-Pedidos.dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre+"-"+strconv.Itoa(mes)+"-"+strconv.Itoa(año)+"-Pedidos.png", cmd, os.FileMode(mode))
	return ""
}
func (l *Lista_Simple) Añitos(nombre string) {
	dot := "digraph G{\n node[shape=circle]\n"
	inicio := l.inico
	for inicio != nil {
		dot += strconv.Itoa(inicio.Año) + "\n"
		año := strconv.Itoa(inicio.Año)
		inicio2 := inicio.Lista.inicio
		dot += año + "->" + strconv.Itoa(inicio2.Mes) + "\n"
		for inicio2 != nil {
			if inicio2.siguiente != nil {
				dot += strconv.Itoa(inicio2.Mes) + "->" + strconv.Itoa(inicio2.siguiente.Mes) + "\n"
			}
			inicio2 = inicio2.siguiente
		}
		inicio.Lista.Mesesitos(nombre, inicio.Año)
		inicio = inicio.Siguiente
	}
	dot += "\n}"
	err := ioutil.WriteFile(nombre+"-Años"+".dot", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", nombre+"-Años"+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre+"-Años"+".png", cmd, os.FileMode(mode))

}
func (l *Lista_doble) Mesesitos(nombre string, año int) {
	dot := "digraph G{\n node[shape=circle]\n"
	inicio := l.inicio

	for inicio != nil {
		dot += strconv.Itoa(inicio.Mes) + "->" + strconv.Itoa(inicio.pedidosM.CabH.dia) + "\n"
		temp := inicio.pedidosM.CabH
		for temp != nil {
			if temp.ESTE != nil {
				dot += strconv.Itoa(temp.dia) + "->" + strconv.Itoa(temp.ESTE.(*NodoH).dia) + "\n"
				temp = temp.ESTE.(*NodoH)
			} else {
				break
			}

		}
		inicio = inicio.siguiente

	}
	dot += "\n}"
	err := ioutil.WriteFile(nombre+"-"+strconv.Itoa(año)+"-Meses.dot", []byte(dot), 0644)
	if err != nil {
		log.Fatal(err)
	}
	ruta, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(ruta, "-Tpng", nombre+"-"+strconv.Itoa(año)+"-Meses.dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre+"-"+strconv.Itoa(año)+"-Meses.png", cmd, os.FileMode(mode))

}

func (l *Lista_Simple) Listaaños(w http.ResponseWriter) {
	aux := make([]Listaño, l.Cantidad)
	aux2 := Infoaño{}
	inicio := l.inico

	cont := 0
	for inicio != nil {
		meses := make([]int, inicio.Lista.Cantidad)
		aux[cont].Año = inicio.Año
		inicio.Lista.listameses(&meses)
		aux[cont].Meses = meses
		cont++
		inicio = inicio.Siguiente
	}
	aux2.Datos = aux
	json.NewEncoder(w).Encode(aux2)

}
func (l *Lista_doble) listameses(meses *[]int) {
	inicio := l.inicio
	cont := 0
	for inicio != nil {
		(*meses)[cont] = inicio.Mes
		cont++
		inicio = inicio.siguiente
	}
}

func Img(nombre string, año int, mes int, w http.ResponseWriter) {
	imgFile, _ := os.Open(nombre + "-" + strconv.Itoa(mes) + "-" + strconv.Itoa(año) + ".png")

	defer imgFile.Close()
	info, _ := imgFile.Stat()
	var size int64 = info.Size()
	buf := make([]byte, size)

	lector := bufio.NewReader(imgFile)
	lector.Read(buf)
	imgBase64 := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(imgBase64)
	json.NewEncoder(w).Encode(imgBase64)

}

func Años(nombre string, año int, mes int, w http.ResponseWriter) {
	imgFile, _ := os.Open(nombre + "-Años" + ".png")

	defer imgFile.Close()
	info, _ := imgFile.Stat()
	var size int64 = info.Size()
	buf := make([]byte, size)

	lector := bufio.NewReader(imgFile)
	lector.Read(buf)
	imgBase64 := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(imgBase64)
	json.NewEncoder(w).Encode(imgBase64)
}
func Meses(nombre string, año int, mes int, w http.ResponseWriter) {
	imgFile, _ := os.Open(nombre + "-" + strconv.Itoa(año) + "-Meses.png")

	defer imgFile.Close()
	info, _ := imgFile.Stat()
	var size int64 = info.Size()
	buf := make([]byte, size)

	lector := bufio.NewReader(imgFile)
	lector.Read(buf)
	imgBase64 := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(imgBase64)
	json.NewEncoder(w).Encode(imgBase64)
}
