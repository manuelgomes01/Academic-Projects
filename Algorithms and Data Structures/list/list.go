package list

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un nuevo elemento al principio de la lista
	InsertarPrimero(T)

	// InsertarUltimo inserta un nuevo elemento al final de la lista
	InsertarUltimo(T)

	// BorrarPrimero saca el elemento del principio de la lista. Si la lista tiene elementos, se quita el primero de la lista, y
	// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si la lista tiene elementos se devuelve el valor del primero.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerPrimero obtiene el valor del ultimo elemento de la lista. Si la lista tiene elementos se devuelve el valor del ultimo.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que hay en la lista
	Largo() int

	// Iterar aplica la función visitar a cada elemento de la lista hasta finalizar su recorrido o que la función de false al aplicarsela a algún elemento de la lista.
	Iterar(visitar func(T) bool)

	// Iterador crea un IteradorLista que permite utilizar sus respectivas primitivas:
	// * VerActual()
	// * HaySiguiente()
	// * Siguiente()
	// * Insertar()
	// * Borrar()
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual obtiene el valor del elemento actual de la lista. Si la lista tiene elementos se devuelve el valor del actual.
	VerActual() T

	// HaySiguiente devuelve true si la lista tiene elementos, false en caso contrario.
	HaySiguiente() bool

	// Siguiente obtiene el valor del próximo elemento de la lista.
	Siguiente()

	// Insertar inserta un nuevo elemento antes de VerActual. El elemento insertado pasa a ser el dato que devolvera VerActual()
	Insertar(T)

	// Borrar saca el elemento que hay en VerActual y avanza el iterador
	Borrar() T
}
