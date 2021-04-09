package ArbolB

type Nodo struct {
	Max   int
	Padre *Nodo
	Keys  []*Key
}

func NuevoNodo(max int) *Nodo {
	keys := make([]*Key, max)
	n := &Nodo{Keys: keys, Padre: nil, Max: max}
	return n
}

func (n *Nodo) Colocar(i int, k *Key) {
	n.Keys[i] = k

}
