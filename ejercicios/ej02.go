package ejercicios

import (
	"sort"
)

// Problema de la mochila 0-1
//Consiste en llenar una "mochila" con un peso limitado, por una cantidad de objetos,
// cada uno con un peso y valor específico, máximizando el valor total almacenado.
// Hay una simplificación del mismo, denominado problema de la mochila 0-1 donde un objeto
// puede estar o no dentro de la mochila, por completo, es decir no se puede fraccionar

type Objeto struct {
	Peso, Valor int
}

func Mochila01(objetos []Objeto, capacidad int) int {
	sort.Slice(objetos, func(i, j int) bool {
		return objetos[i].Valor/objetos[i].Peso > objetos[j].Valor/objetos[j].Peso
	})

	pesoActual := 0
	v_total := 0
	for _, objeto := range objetos {
		if pesoActual+objeto.Peso <= capacidad {
			v_total += objeto.Valor
			pesoActual += objeto.Peso
		} else {
			break
		}
	}
	return v_total
}
