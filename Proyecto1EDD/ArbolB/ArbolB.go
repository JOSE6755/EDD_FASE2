package ArbolB

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"../MatrizD"
)

type Arbol struct {
	k    int
	Raiz *Nodo
}

type Usu struct {
	DPI    int    `json:"DPI"`
	Pass   string `json:"Password"`
	Correo string `json:"Correo"`
}

type Aux struct {
	Tipo     string
	Correcto bool
}

func Nuevoarbol(nivel int) *Arbol {
	arb := &Arbol{k: nivel, Raiz: nil}
	nodoraiz := NuevoNodo(nivel)
	arb.Raiz = nodoraiz
	return arb
}

func (a *Arbol) Gpedido(DPI int, pedido MatrizD.Pedidos, w http.ResponseWriter) {

	inicio := a.Raiz
	encontrado := false
	for inicio != nil {
		for i := 0; i < inicio.Max; i++ {
			if inicio.Keys[i] != nil {
				if DPI < inicio.Keys[i].DPI {
					inicio = inicio.Keys[i].Izquierdo
					break
				} else if DPI == inicio.Keys[i].DPI {
					if w == nil {
						inicio.Keys[i].Pedidos = pedido
					} else {
						json.NewEncoder(w).Encode(inicio.Keys[i].Pedidos)
					}
					encontrado = true
					break
				}
			} else {
				inicio = inicio.Keys[i-1].Derecho
				break
			}
		}
		if encontrado == true {
			break
		}
	}
}
func (a *Arbol) Buscar(DPI int, pass string, correo string, w http.ResponseWriter) {

	inicio := a.Raiz
	encontrado := false
	for inicio != nil {
		for i := 0; i < inicio.Max; i++ {
			if inicio.Keys[i] != nil {
				if DPI < inicio.Keys[i].DPI {
					inicio = inicio.Keys[i].Izquierdo
					break
				} else if DPI == inicio.Keys[i].DPI {
					if inicio.Keys[i].Password == pass {
						encontrado = true
						aux := Aux{Tipo: inicio.Keys[i].Cuenta, Correcto: true}
						json.NewEncoder(w).Encode(aux)
					} else {
						encontrado = true
						aux := Aux{Tipo: inicio.Keys[i].Cuenta, Correcto: false}
						json.NewEncoder(w).Encode(aux)
					}
					break
				}
			} else {
				inicio = inicio.Keys[i-1].Derecho
				break
			}
		}
		if encontrado == true {
			break
		}
	}
}

func (a *Arbol) Insertar(k *Key) {
	if a.Raiz.Keys[0] == nil {
		a.Raiz.Colocar(0, k)
	} else if a.Raiz.Keys[0].Izquierdo == nil {
		inser := -1
		aux := a.Raiz
		inser = a.colocarNodo(aux, k)
		if inser != -1 {
			if inser == aux.Max-1 {
				medio := aux.Max / 2
				centro := aux.Keys[medio]
				der := NuevoNodo(a.k)
				izq := NuevoNodo(a.k)
				indiceizq := 0
				indiceder := 0
				for j := 0; j < aux.Max; j++ {
					if aux.Keys[j].DPI < centro.DPI {
						izq.Colocar(indiceizq, aux.Keys[j])
						indiceizq++
						aux.Colocar(j, nil)

					} else if aux.Keys[j].DPI > centro.DPI {
						der.Colocar(indiceder, aux.Keys[j])
						indiceder++
						aux.Colocar(j, nil)
					}
				}
				aux.Colocar(medio, nil)
				a.Raiz = aux
				a.Raiz.Colocar(0, centro)
				izq.Padre = a.Raiz
				der.Padre = a.Raiz
				centro.Izquierdo = izq
				centro.Derecho = der

			}
		}

	} else if a.Raiz.Keys[0].Izquierdo != nil {
		aux := a.Raiz
		for aux.Keys[0].Izquierdo != nil {
			l := 0
			for i := 0; i < aux.Max; i, l = i+1, l+1 {
				if aux.Keys[i] != nil {
					if aux.Keys[i].DPI > k.DPI {
						aux = aux.Keys[i].Izquierdo
						break
					}
				} else {
					aux = aux.Keys[i-1].Derecho
					break
				}
			}
			if l == aux.Max {
				aux = aux.Keys[l-1].Derecho

			}
		}
		indice := a.colocarNodo(aux, k)
		if indice == aux.Max-1 {
			for aux.Padre != nil {
				medio := aux.Max / 2
				centro := aux.Keys[medio]
				izq := NuevoNodo(a.k)
				der := NuevoNodo(a.k)
				indiceizq := 0
				indiceder := 0
				for i := 0; i < aux.Max; i++ {
					if aux.Keys[i].DPI < centro.DPI {
						izq.Colocar(indiceizq, aux.Keys[i])
						indiceizq++
						aux.Colocar(i, nil)
					} else if aux.Keys[i].DPI > centro.DPI {
						der.Colocar(indiceder, aux.Keys[i])
						indiceder++
						aux.Colocar(i, nil)
					}
				}
				aux.Colocar(medio, nil)
				centro.Izquierdo = izq
				centro.Derecho = der
				aux = aux.Padre
				izq.Padre = aux
				der.Padre = aux
				for i := 0; i < izq.Max; i++ {

					if izq.Keys[i] != nil {
						if izq.Keys[i].Izquierdo != nil {
							izq.Keys[i].Izquierdo.Padre = izq
						}
						if izq.Keys[i].Derecho != nil {
							izq.Keys[i].Derecho.Padre = izq
						}

					}

				}
				for i := 0; i < der.Max; i++ {
					if der.Keys[i] != nil {
						if der.Keys[i].Izquierdo != nil {
							der.Keys[i].Izquierdo.Padre = der
						}
						if der.Keys[i].Derecho != nil {
							der.Keys[i].Derecho.Padre = der
						}
					}
				}
				colocar := a.colocarNodo(aux, centro)
				if colocar == aux.Max-1 {
					if aux.Padre == nil {
						medio := aux.Max / 2
						centro := aux.Keys[medio]
						izq := NuevoNodo(a.k)
						der := NuevoNodo(a.k)
						indiceizq := 0
						indiceder := 0
						for i := 0; i < aux.Max; i++ {
							if aux.Keys[i].DPI < centro.DPI {
								izq.Colocar(indiceizq, aux.Keys[i])
								indiceizq++
								aux.Colocar(i, nil)
							} else if aux.Keys[i].DPI > centro.DPI {
								der.Colocar(indiceder, aux.Keys[i])
								indiceder++
								aux.Colocar(i, nil)
							}
						}
						aux.Colocar(medio, nil)
						aux.Colocar(0, centro)
						for i := 0; i < a.k; i++ {
							if izq.Keys[i] != nil {
								izq.Keys[i].Izquierdo.Padre = izq
								izq.Keys[i].Derecho.Padre = izq
							}
						}
						for i := 0; i < a.k; i++ {
							if der.Keys[i] != nil {
								der.Keys[i].Izquierdo.Padre = der
								der.Keys[i].Derecho.Padre = der
							}
						}
						centro.Izquierdo = izq
						centro.Derecho = der
						izq.Padre = aux
						der.Padre = aux
						a.Raiz = aux

					}
					continue
				} else {
					break
				}

			}
		}
	}
}

func (a *Arbol) colocarNodo(node *Nodo, k *Key) int {
	indice := -1
	for i := 0; i < node.Max; i++ {
		if node.Keys[i] == nil {
			colocado := false
			for j := i - 1; j >= 0; j-- {
				if node.Keys[j].DPI > k.DPI {
					node.Colocar(j+1, node.Keys[j])
				} else {
					node.Colocar(j+1, k)
					node.Keys[j].Derecho = k.Izquierdo
					if j+2 < a.k && node.Keys[j+2] != nil {
						node.Keys[j+2].Izquierdo = k.Derecho
					}
					colocado = true
					break
				}
			}
			if colocado == false {
				node.Colocar(0, k)
				node.Keys[1].Izquierdo = k.Derecho

			}
			indice = i
			break
		}
	}
	return indice
}

func (a *Arbol) Graficar(nombre string) {
	var builder strings.Builder
	var build2 strings.Builder
	var build3 strings.Builder
	nombre1 := nombre + "B"
	nombre2 := nombre + "C"
	nombre3 := nombre + "CS"
	fmt.Fprintf(&builder, "digraph G{\n node[shape=record]\n")
	fmt.Fprintf(&build2, "digraph G{\n node[shape=record]\n")
	fmt.Fprintf(&build3, "digraph G{\n node[shape=record]\n")
	m := make(map[string]*Nodo)
	n := make(map[string]*Nodo)
	c := make(map[string]*Nodo)
	graficar(a.Raiz, &builder, m, nil, 0)
	graficar2(a.Raiz, &build2, n, nil, 0)
	graficar3(a.Raiz, &build3, c, nil, 0)
	fmt.Fprintf(&builder, "}")
	fmt.Fprintf(&build2, "}")
	fmt.Fprintf(&build3, "}")
	estructuraDot := builder.String()
	dot2 := build2.String()
	dot3 := build3.String()

	generarArchivo(estructuraDot, nombre1)
	generarImagen("arbolB.png", nombre1)

	generarArchivo(dot2, nombre2)
	generarImagen("arbolC.png", nombre2)

	generarArchivo(dot3, nombre3)
	generarImagen("arbolCS.png", nombre3)
}
func graficar(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	fmt.Println(contiene)
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	if actual.Keys[0] != nil {
		fmt.Fprintf(cad, "node%p[label=\"", &(*actual))
	}
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++
			} else {
				cadi := strconv.Itoa(actual.Keys[i].DPI) + "|"
				fmt.Fprintf(cad, cadi)
				j++
				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	if actual.Keys[0] != nil {
		fmt.Fprintf(cad, "\"]\n")
	}
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficar(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficar(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p : f%d -> node%p\n", &(*padre), pos, &(*actual))
	}
}
func graficar2(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	fmt.Println(contiene)
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	if actual.Keys[0] != nil {
		fmt.Fprintf(cad, "node%p[label=\"", &(*actual))
	}
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++
			} else {
				cadi := actual.Keys[i].Password + "|"
				fmt.Fprintf(cad, cadi)
				j++
				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	if actual.Keys[0] != nil {
		fmt.Fprintf(cad, "\"]\n")
	}
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficar2(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficar2(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p : f%d -> node%p\n", &(*padre), pos, &(*actual))
	}
}
func graficar3(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	fmt.Println(contiene)
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	if actual.Keys[0] != nil {
		fmt.Fprintf(cad, "node%p[label=\"", &(*actual))
	}
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++
			} else {
				contra := sha256.Sum256([]byte(actual.Keys[i].Correo))
				cadi := fmt.Sprintf("%x", contra) + "|"
				fmt.Fprintf(cad, cadi)
				j++
				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	if actual.Keys[0] != nil {
		fmt.Fprintf(cad, "\"]\n")
	}
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficar3(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficar3(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p : f%d -> node%p\n", &(*padre), pos, &(*actual))
	}
}
func generarArchivo(cadena string, nombre string) {
	f, err := os.Create(nombre + ".dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func generarImagen(nombre string, nombre2 string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./"+nombre2+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile(nombre, cmd, os.FileMode(mode))
}

func BuscarImg(nombre string, w http.ResponseWriter) {
	imgFile, _ := os.Open(nombre + ".png")

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
