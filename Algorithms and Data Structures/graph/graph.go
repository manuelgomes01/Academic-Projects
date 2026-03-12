package graph

type Grafo[T comparable, W any] interface {
	EsDirigido() bool

	ExisteVertice(vertice T) bool

	AgregarVertice(v T)

	BorrarVertice(v T)

	AgregarArista(v, w T, peso W)

	ObtenerPeso(v, w T) W

	BorrarArista(v, w T)

	EstanUnidos(v, w T) bool

	Cantidad() int

	ObtenerVertices() []T

	Adyacentes(v T) []T
}
