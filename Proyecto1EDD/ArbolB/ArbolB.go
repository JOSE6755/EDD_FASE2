package ArbolB

type Arbol struct {
	k    int
	Raiz *Nodo
}

func Nuevoarbol(nivel int) *Arbol {
	arb := &Arbol{k: nivel, Raiz: nil}
	nodoraiz := NuevoNodo(nivel)
	arb.Raiz = nodoraiz
	return arb
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
