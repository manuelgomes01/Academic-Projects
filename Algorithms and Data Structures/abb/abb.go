package abb

import (
	TDAPila "tdas/pila"
)

type funcCmp[K any] func(K, K) int

type nodoAbb[K any, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K any, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iteradorAbb[K any, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   funcCmp[K]
}

// crearNodo devuelve un puntero a un nodo sin hijos que contiene clave y dato como clave - valor.
func (abb *abb[K, V]) crearNodo(clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		izquierdo: nil,
		derecho:   nil,
		clave:     clave,
		dato:      dato,
	}
}

func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cantidad: 0,
		cmp:      funcionCmp,
	}
}

// panicClavesNoExistente levanta un panic en caso de que el nodo sea nil.
func (nodo *nodoAbb[K, V]) panicClaveNoExistente() {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
}

// buscador busca la referencia al nodo que contiene la clave pasada por parametro y devuelve un puntero a la misma.
func buscador[K any, V any](clave K, comparar funcCmp[K], ref **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*ref) == nil {
		return ref
	}
	cmp := comparar(clave, (*ref).clave)
	if cmp == 0 {
		return ref
	} else if cmp > 0 {
		return buscador(clave, comparar, &(*ref).derecho)
	}
	return buscador(clave, comparar, &(*ref).izquierdo)
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	ref := buscador(clave, abb.cmp, &abb.raiz)
	if *ref == nil {
		abb.cantidad++
		*ref = abb.crearNodo(clave, dato)
	} else {
		(*ref).dato = dato
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo := buscador(clave, abb.cmp, &abb.raiz)
	return *nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := buscador(clave, abb.cmp, &abb.raiz)
	(*nodo).panicClaveNoExistente()

	return (*nodo).dato
}

// buscarMaximo devuelve una referencia al nodo con la clave de mayor valor del ABB.
func (nodo *nodoAbb[K, V]) buscarMaximo() **nodoAbb[K, V] {
	if nodo.derecho == nil {
		return &nodo
	}
	return nodo.derecho.buscarMaximo()
}

func (abb *abb[K, V]) Borrar(clave K) V {
	ref := buscador(clave, abb.cmp, &abb.raiz)
	(*ref).panicClaveNoExistente()

	valor := (*ref).dato

	if (*ref).izquierdo == nil {
		*ref = (*ref).derecho
	} else if (*ref).derecho == nil {
		*ref = (*ref).izquierdo
	} else {
		refReemplazo := (*ref).izquierdo.buscarMaximo()
		reemplazo := *refReemplazo
		abb.Borrar(reemplazo.clave)
		(*ref).clave = reemplazo.clave
		(*ref).dato = reemplazo.dato
		abb.cantidad++
	}

	abb.cantidad--
	return valor
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

// ################ Iterar dicc #################
func (abb *abb[K, V]) iterarRangoAux(actual *nodoAbb[K, V], desde, hasta *K, visitar func(K, V) bool) bool {
	if actual == nil {
		return true
	}

	if desde == nil || abb.cmp(actual.clave, *desde) > 0 {
		if !abb.iterarRangoAux(actual.izquierdo, desde, hasta, visitar) {
			return false
		}
	}

	if (desde == nil || abb.cmp(actual.clave, *desde) >= 0) && (hasta == nil || abb.cmp(actual.clave, *hasta) <= 0) {
		if !visitar(actual.clave, actual.dato) {
			return false
		}
	}

	if hasta == nil || abb.cmp(actual.clave, *hasta) < 0 {
		if !abb.iterarRangoAux(actual.derecho, desde, hasta, visitar) {
			return false
		}
	}
	return true
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.iterarRangoAux(abb.raiz, nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde, hasta *K, visitar func(K, V) bool) {
	abb.iterarRangoAux(abb.raiz, desde, hasta, visitar)
}

// ############## Iterador dicc #####################
// apilarPorRango apila los nodos que se encuentran dentro del rango mencionado.
func (nodo *nodoAbb[K, V]) apilarPorRango(desde, hasta *K, cmp funcCmp[K], pila TDAPila.Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}

	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		pila.Apilar(nodo)
		nodo.izquierdo.apilarPorRango(desde, hasta, cmp, pila)
	} else if desde == nil || cmp(nodo.clave, *desde) < 0 {
		nodo.derecho.apilarPorRango(desde, hasta, cmp, pila)
	} else {
		nodo.izquierdo.apilarPorRango(desde, hasta, cmp, pila)
	}
}

// panicFinalizoIteracion levanta un panic con el mesaje "El iterador termino de iterar" si la pila se encuentra vacia.
func (iter *iteradorAbb[K, V]) panicFinalizoIteracion() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (abb *abb[K, V]) IteradorRango(desde, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	abb.raiz.apilarPorRango(desde, hasta, abb.cmp, pila)
	return &iteradorAbb[K, V]{pila: pila, desde: desde, hasta: hasta, cmp: abb.cmp}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (iter *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorAbb[K, V]) VerActual() (K, V) {
	iter.panicFinalizoIteracion()
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iteradorAbb[K, V]) Siguiente() {
	iter.panicFinalizoIteracion()
	nodo := iter.pila.Desapilar()
	if nodo.derecho != nil {
		nodo.derecho.apilarPorRango(iter.desde, iter.hasta, iter.cmp, iter.pila)
	}
}
