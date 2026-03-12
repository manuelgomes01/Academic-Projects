package graph

import (
	"fmt"
	d "tdas/diccionario"
)

type grafo[T comparable, W any] struct {
	vertices d.Diccionario[T, d.Diccionario[T, W]]
	dirigido bool
}

func (g *grafo[T, W]) panicVerticeNoExistente(v T) {
	if !g.ExisteVertice(v) {
		panic(fmt.Sprintf("No existe el vertice: %v", v))
	}
}

func CrearGrafo[T comparable, W any](dirigido bool, vertices []T) Grafo[T, W] {
	verticesDicc := d.CrearHash[T, d.Diccionario[T, W]](func(a, b T) bool { return a == b })
	for _, v := range vertices {
		verticesDicc.Guardar(v, d.CrearHash[T, W](func(a, b T) bool { return a == b }))
	}

	return &grafo[T, W]{
		vertices: verticesDicc,
		dirigido: dirigido,
	}
}

func (g *grafo[T, W]) EsDirigido() bool { return g.dirigido }

func (g *grafo[T, W]) ExisteVertice(v T) bool {
	return g.vertices.Pertenece(v)
}

func (g *grafo[T, W]) AgregarVertice(v T) {
	if g.ExisteVertice(v) {
		panic("Vertice ya existente.")
	}
	g.vertices.Guardar(v, d.CrearHash[T, W](func(a, b T) bool { return a == b }))
}

func (g *grafo[T, W]) BorrarVertice(v T) {
	g.panicVerticeNoExistente(v)

	ady := g.vertices.Obtener(v)
	if !g.dirigido {
		ady.Iterar(func(ady T, _ W) bool {
			g.vertices.Obtener(ady).Borrar(v)
			return true
		})
		return
	}

	g.vertices.Iterar(func(_ T, ady d.Diccionario[T, W]) bool {
		if ady.Pertenece(v) {
			ady.Borrar(v)
		}
		return true
	})
	g.vertices.Borrar(v)
}

func (g *grafo[T, W]) AgregarArista(v T, w T, peso W) {
	g.panicVerticeNoExistente(v)
	g.panicVerticeNoExistente(w)

	g.vertices.Obtener(v).Guardar(w, peso)

	if !g.dirigido {
		g.vertices.Obtener(w).Guardar(v, peso)
	}
}

func (g *grafo[T, W]) ObtenerPeso(v, w T) W {
	if !g.EstanUnidos(v, w) {
		panic("Arista no existente")
	}
	return g.vertices.Obtener(v).Obtener(w)
}

func (g *grafo[T, W]) BorrarArista(v T, w T) {
	g.panicVerticeNoExistente(v)
	g.panicVerticeNoExistente(w)

	ady1 := g.vertices.Obtener(v)
	if ady1.Pertenece(w) {
		ady1.Borrar(w)
	}

	if !g.dirigido {
		g.vertices.Obtener(w).Borrar(v)
	}
}

func (g *grafo[T, W]) EstanUnidos(v T, w T) bool {
	g.panicVerticeNoExistente(v)
	g.panicVerticeNoExistente(w)
	return g.vertices.Obtener(v).Pertenece(w)
}

func (g *grafo[T, W]) Cantidad() int {
	return g.vertices.Cantidad()
}

func (g *grafo[T, W]) ObtenerVertices() []T {
	vertices := make([]T, 0, g.Cantidad())
	g.vertices.Iterar(func(clave T, _ d.Diccionario[T, W]) bool {
		vertices = append(vertices, clave)
		return true
	})
	return vertices
}

func (g *grafo[T, W]) Adyacentes(v T) []T {
	g.panicVerticeNoExistente(v)
	ady := g.vertices.Obtener(v)
	adyacentes := make([]T, 0, ady.Cantidad())
	ady.Iterar(func(vertice T, _ W) bool {
		adyacentes = append(adyacentes, vertice)
		return true
	})
	return adyacentes
}
