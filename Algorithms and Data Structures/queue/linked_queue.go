package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

/*
 * Precondicion:
 * Postcondicion: Si la cola esta vacia, hace un llamado a panic con el mensaje "La cola esta vacia""
 */
func (cola *colaEnlazada[T]) validarColaNoVacia() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.validarColaNoVacia()
	return cola.primero.dato
}

/*
 * Precondicion:
 * Postcondicion: Devuelve un 'nodoCola' con 'elemento' como dato
 */
func (cola *colaEnlazada[T]) crearNodo(elemento T) nodoCola[T] {
	return nodoCola[T]{dato: elemento, prox: nil}
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nodo := cola.crearNodo(elemento)

	if cola.EstaVacia() {
		cola.primero = &nodo
	} else {
		cola.ultimo.prox = &nodo
	}

	cola.ultimo = &nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.validarColaNoVacia()

	nodo := cola.primero

	if cola.primero == cola.ultimo {
		cola.primero = nil
		cola.ultimo = nil
	} else {
		cola.primero = nodo.prox
	}

	return nodo.dato
}
