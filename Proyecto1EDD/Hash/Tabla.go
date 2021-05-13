package Hash

import (
	"fmt"
)

type Node struct {
	Hash  int    // dpi o llave
	Valor string //lo que se guardará

}
type Hashtable struct {
	Size                   int
	Carga                  int
	Porcentaje             int
	Porcentaje_crecimiento int
	Arreglo                []*Node
}

func NuevaTabla(size int, Porcentaje int, Porcentaje_crecimiento int) *Hashtable {
	arreglo := make([]*Node, size)
	return &Hashtable{size, 0, Porcentaje, Porcentaje_crecimiento, arreglo}
}

func (this *Hashtable) Insertar(nuevo int, valor string) {
	nuevo_nodo := Node{nuevo, valor}
	pos := this.posicion(nuevo, valor)
	this.Arreglo[pos] = &nuevo_nodo
	this.Carga++
	porcentaje_ocupado := (this.Carga * 100) / this.Size

	if porcentaje_ocupado > this.Porcentaje {
		sizenuevo := this.Size
		for {
			sizenuevo++
			porcentaje_ocupado = (this.Carga * 100) / sizenuevo
			if porcentaje_ocupado <= this.Porcentaje_crecimiento {
				break
			}
		}
		nuevo_array := make([]*Node, sizenuevo)
		antiguo := this.Arreglo
		this.Arreglo = nuevo_array
		this.Size = sizenuevo
		aux := 0
		for i := 0; i < len(antiguo); i++ {
			if antiguo[i] != nil {
				aux = this.posicion(antiguo[i].Hash, antiguo[i].Valor)
				nuevo_array[aux] = antiguo[i]
			}
		}
	}
}

func (this *Hashtable) posicion(clave int, valor string) int {
	i, p := 0, 0
	d := (0.6180334 * float64(clave)) - float64(int(0.6180334*float64(clave)))
	p = int(float64(this.Size) * d)
	fmt.Println("h(x):", p)
	for this.Arreglo[p] != nil && this.Arreglo[p].Valor != valor {
		i++
		i = i * i
		p = p + i
		fmt.Println("antes del ---if.... posicion:", p, "---i---", i)
		if p >= this.Size {
			p = p - this.Size
			fmt.Println("entra al if:", p)
		}
	}
	fmt.Println("posicion:", p)
	return p
}

/*func (this *Hashtable) imprimir() {
	data := make([][]string, this.Size)
	fmt.Println("dpi <---> valor")
	for i := 0; i < len(this.Arreglo); i++ {
		tmp := make([]string, 2)
		aux := this.Arreglo[i]
		if aux != nil {
			tmp[0] = strconv.Itoa(aux.Hash)
			tmp[1] = aux.Valor
		} else {
			tmp[0] = "-"
			tmp[1] = "-"
		}
		data[i] = tmp
		fmt.Println(tmp[0], "<--->", tmp[1])
	}

}
*/

/*func main() {
	ta := NewTableH(7, 60, 20)
	for i := 0; i < 20; i++ {
		ta.Insertar(i, "comentario:"+strconv.Itoa(i))
	}
	ta.imprimir()

}
*/
