package feed

import (
	"errors"
	h "tdas/cola_prioridad"
)

type feed[T any] struct {
	heap h.ColaPrioridad[T]
}

func CrearFeed[T any](cmp func(T, T) int) Feed[T] {
	return &feed[T]{heap: h.CrearHeap(cmp)}
}

func (f *feed[T]) AgregarPost(post T) {
	f.heap.Encolar(post)
}

func (f *feed[T]) Proximo() (T, error) {
	var zero T // Valor cero del tipo genérico T
	if f.heap.EstaVacia() {
		return zero, errors.New("Usuario no loggeado o no hay mas posts para ver")
	}
	return f.heap.Desencolar(), nil
}
