package list

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// crearNodo devuelve un puntero a nodoLista con elemento como dato.
func (lista *listaEnlazada[T]) crearNodo(elemento T, proximo *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: elemento, siguiente: proximo}
}

// validarListaVacia levanta un panic con el mensaje "La lista esta vacia" si la lista no contiene elementos.
func (lista *listaEnlazada[T]) validarListaNoVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (iter *iteradorLista[T]) validarIteracionFinalizada() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil && lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := lista.crearNodo(elemento, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}

	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nodo := lista.crearNodo(elemento, nil)
	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.siguiente = nodo
	}
	lista.ultimo = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.validarListaNoVacia()

	nodoBorrado := lista.primero
	if nodoBorrado == lista.ultimo {
		lista.ultimo = nil
	}

	lista.primero = nodoBorrado.siguiente
	lista.largo--
	return nodoBorrado.dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.validarListaNoVacia()
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.validarListaNoVacia()
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodo := lista.primero
	for nodo != nil {
		if !visitar(nodo.dato) {
			return
		}
		nodo = nodo.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{actual: lista.primero, anterior: nil, lista: lista}
}

func (iter *iteradorLista[T]) VerActual() T {
	iter.validarIteracionFinalizada()
	return iter.actual.dato
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Siguiente() {
	iter.validarIteracionFinalizada()

	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(elemento T) {
	nuevoNodo := iter.lista.crearNodo(elemento, iter.actual)

	if iter.anterior == nil {
		iter.lista.primero = nuevoNodo
	} else {
		iter.anterior.siguiente = nuevoNodo
	}

	if iter.actual == nil {
		iter.lista.ultimo = nuevoNodo
	}

	iter.lista.largo++
	iter.actual = nuevoNodo
}

func (iter *iteradorLista[T]) Borrar() T {
	iter.validarIteracionFinalizada()

	nodoBorrado := iter.actual
	if iter.anterior == nil {
		iter.lista.primero = nodoBorrado.siguiente
	} else {
		iter.anterior.siguiente = nodoBorrado.siguiente
	}
	if nodoBorrado.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = nodoBorrado.siguiente
	iter.lista.largo--
	return nodoBorrado.dato
}
