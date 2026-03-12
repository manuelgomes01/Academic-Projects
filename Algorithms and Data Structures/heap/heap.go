package heap

const (
	_CAPACIDAD_INICIAL       = 10
	_REDIMENSIONAR_CAPACIDAD = 2
	_FACTOR_REDIMENSION      = 4
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func crearHeap[T any](arr []T, cantidad int, funcionCmp func(T, T) int) *colaConPrioridad[T] {
	return &colaConPrioridad[T]{datos: arr, cant: cantidad, cmp: funcionCmp}
}

func CrearHeap[T any](funcionCmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, _CAPACIDAD_INICIAL)
	return crearHeap(datos, 0, funcionCmp)
}

func upHeap[T any](actual int, datos []T, cmp func(T, T) int) {
	if actual <= 0 {
		return
	}
	padre := (actual - 1) / 2
	if cmp(datos[padre], datos[actual]) >= 0 {
		return
	}
	aux := datos[actual]
	datos[actual] = datos[padre]
	datos[padre] = aux
	upHeap[T](padre, datos, cmp)
}

func downHeap[T any](actual int, datos []T, cmp func(T, T) int) {
	if actual >= len(datos) {
		return
	}
	hijoMayor := 0
	hijoIzq := (2 * actual) + 1
	hijoDer := (2 * actual) + 2

	if hijoIzq < len(datos) {
		if hijoDer < len(datos) && cmp(datos[hijoDer], datos[hijoIzq]) >= 0 {
			hijoMayor = hijoDer
		} else {
			hijoMayor = hijoIzq
		}
	}

	if hijoMayor == 0 || cmp(datos[actual], datos[hijoMayor]) >= 0 {
		return
	}

	aux := datos[actual]
	datos[actual] = datos[hijoMayor]
	datos[hijoMayor] = aux
	downHeap[T](hijoMayor, datos, cmp)
}

func heapify[T any](datos []T, cmp func(T, T) int) {
	for i := (len(datos) - 1) / 2; i >= 0; i-- {
		downHeap(i, datos, cmp)
	}
}

func CrearHeapArr[T any](arreglo []T, funcionCmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap(funcionCmp)
	}
	capacidad := max(len(arreglo), _CAPACIDAD_INICIAL)
	arrAux := make([]T, capacidad)
	copy(arrAux, arreglo)
	heapify(arrAux[:len(arreglo)], funcionCmp)
	heap := crearHeap(arrAux, len(arreglo), funcionCmp)
	return heap
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) validarHeapNoVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (heap *colaConPrioridad[T]) redimensionarHeap(nuevaCapacidad int) {
	redimensionarDatos := make([]T, nuevaCapacidad)
	copy(redimensionarDatos, heap.datos)
	heap.datos = redimensionarDatos
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	heap.datos[heap.cant] = elemento
	upHeap(heap.cant, heap.datos[:heap.cant+1], heap.cmp)
	heap.cant += 1
	if heap.cant == len(heap.datos) {
		heap.redimensionarHeap(heap.cant * _REDIMENSIONAR_CAPACIDAD)
	}
}

func (heap *colaConPrioridad[T]) VerMax() T {
	heap.validarHeapNoVacia()
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.validarHeapNoVacia()
	desencolado := heap.datos[0]

	heap.datos[0] = heap.datos[heap.cant-1]
	heap.cant -= 1
	downHeap(0, heap.datos[:heap.cant], heap.cmp)

	if heap.cant <= len(heap.datos)/_FACTOR_REDIMENSION && len(heap.datos) > _CAPACIDAD_INICIAL {
		heap.redimensionarHeap(len(heap.datos) / _REDIMENSIONAR_CAPACIDAD)
	}

	return desencolado
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func swap[T any](x, y *T) {
	*x, *y = *y, *x
}

func HeapSort[T any](elementos []T, funcionCmp func(T, T) int) {
	heapify(elementos, funcionCmp)
	for i := 0; i < len(elementos); i++ {
		swap(&elementos[0], &elementos[len(elementos)-1-i])
		downHeap(0, elementos[:len(elementos)-1-i], funcionCmp)
	}
}
