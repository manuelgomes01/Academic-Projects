package stack

const _CAPACIDAD_INICIAL = 10
const _REDIMENSIONAR_CAPACIDAD = 2
const _FACTOR_REDIMENSION = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := &pilaDinamica[T]{}
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) validarPilaNoVacia() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (pila *pilaDinamica[T]) VerTope() T {
	pila.validarPilaNoVacia()
	return pila.datos[pila.cantidad-1]
}

/*
 * Precondicion: 'nueva_capacidad' debe indicar la nueva capacidad que tendra el arreglo 'pila.datos'
 * Postcondicion: Cambia el tamaño de 'pila.datos' dependiendo del valor que recibe por 'nueva_capacidad' sin que se pierdan los datos
 */
func (pila *pilaDinamica[T]) redimensionarPila(nueva_capacidad int) {
	redimensionarDatos := make([]T, nueva_capacidad)
	copy(redimensionarDatos, pila.datos)
	pila.datos = redimensionarDatos
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	pila.datos[pila.cantidad] = elemento
	pila.cantidad += 1

	if pila.cantidad == len(pila.datos) {
		pila.redimensionarPila(pila.cantidad * _REDIMENSIONAR_CAPACIDAD)
	}
}

func (pila *pilaDinamica[T]) Desapilar() T {
	pila.validarPilaNoVacia()

	elemento := pila.datos[pila.cantidad-1]
	pila.cantidad -= 1

	if pila.cantidad*_FACTOR_REDIMENSION <= len(pila.datos) && len(pila.datos) > _CAPACIDAD_INICIAL {
		pila.redimensionarPila(len(pila.datos) / _REDIMENSIONAR_CAPACIDAD)
	}

	return elemento
}
