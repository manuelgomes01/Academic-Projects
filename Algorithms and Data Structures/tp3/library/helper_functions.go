package library

import (
	q "tdas/cola_prioridad"
	d "tdas/diccionario"
)

type auxTopK struct {
	vertice   string
	prioridad float64
}

// Invertir invierte el orden del arreglo pasado.
func Invertir[K comparable](camino []K) []K {
	inicio := 0
	fin := len(camino) - 1
	for inicio < fin {
		camino[inicio], camino[fin] = camino[fin], camino[inicio]
		inicio++
		fin--
	}
	return camino
}

// topK devuelve los vertices con los pageRanks mas altos.
func TopK(k int, pageRank d.Diccionario[string, float64]) []string {
	heap := q.CrearHeap(func(a, b auxTopK) int {
		if b.prioridad > a.prioridad {
			return 1
		} else if b.prioridad < a.prioridad {
			return -1
		}
		return 0
	})

	elementos := make([]string, 0)
	pageRank.Iterar(func(clave string, valor float64) bool {
		if heap.Cantidad() < k {
			heap.Encolar(auxTopK{clave, valor})
		} else if valor > heap.VerMax().prioridad { // es de minimos VerMin()
			heap.Desencolar()
			heap.Encolar(auxTopK{clave, valor})
		}
		return true
	})

	for !heap.EstaVacia() {
		elementos = append(elementos, heap.Desencolar().vertice)
	}
	return Invertir(elementos)
}
