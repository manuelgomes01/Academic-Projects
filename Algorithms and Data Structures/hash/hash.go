package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estado int

const (
	_VACIO = estado(iota)
	_OCUPADO
	_BORRADO
)

const _TAMANIO_INICIAL = 11
const _FACTOR_REDIMENSION = 2
const _FACTOR_CARGA_AUMENTAR = 0.7
const _FACTOR_CARGA_REDUCIR = 0.3
const _ERROR = -1

// ===== HASH CERRADO =====
type celdaHash[K any, V any] struct {
	clave  K
	valor  V
	estado estado
}

type hashCerrado[K any, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
	comparar func(K, K) bool
}

type iteradorHash[K any, V any] struct {
	posActual int
	tabla     []celdaHash[K, V]
	tam       int
}

func crearTabla[K any, V any](f func(K, K) bool, tam int) *hashCerrado[K, V] {
	return &hashCerrado[K, V]{tabla: make([]celdaHash[K, V], tam),
		cantidad: 0,
		tam:      tam,
		borrados: 0,
		comparar: f,
	}
}

func CrearHash[K any, V any](f func(K, K) bool) Diccionario[K, V] {
	return crearTabla[K, V](f, _TAMANIO_INICIAL)
}

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// funcion de hash del paquete hash de Go (https://pkg.go.dev/hash/fnv)
func (dicc *hashCerrado[K, V]) fnvHash(clave K) uint64 {
	bytesClave := convertirABytes(clave)
	hash := fnv.New64a()

	hash.Write(bytesClave)
	return hash.Sum64()
}

// existeClave devuelve true si la clave esta en el diccionario. En caso contrario, devuelve false.
func (dicc *hashCerrado[K, V]) existeClave(clave K, index int) bool {
	return dicc.comparar(dicc.tabla[index].clave, clave) && dicc.tabla[index].estado == _OCUPADO
}

// panicClaveNoExistente levanta un panic con el mensaje "La clave no pertenece al diccionario", en caso de que la clave no pertenezca al diccionario.
func (dicc *hashCerrado[K, V]) panicClaveNoExistente(clave K, index int) {
	if !dicc.existeClave(clave, index) {
		panic("La clave no pertenece al diccionario")
	}
}

// buscador devuelve el primero índice vacío o el índice de la clave buscada.
func (dicc *hashCerrado[K, V]) buscador(clave K) int {
	i := int(dicc.fnvHash(clave))
	index := i % dicc.tam
	if index < 0 { // validación necesaria porque la función siempre devuelve un negativo.
		index *= -1
	}
	celdaActual := dicc.tabla[index]
	for celdaActual.estado != _VACIO {
		if dicc.comparar(celdaActual.clave, clave) {
			return index
		}
		if index == dicc.tam-1 {
			index = 0
		} else {
			index++
		}
		celdaActual = dicc.tabla[index]
	}
	return index
}

// pre: nuevaCapacidad debe ser mayor a 0.
// redimensionarDicc modifica el tamaño del diccionario a una nuevaCapacidad restaurando la cantidad de borrados.
func (dicc *hashCerrado[K, V]) redimendionarDicc(nuevaCapacidad int) {
	diccAux := crearTabla[K, V](dicc.comparar, nuevaCapacidad)
	for _, celda := range dicc.tabla {
		if celda.estado == _OCUPADO {
			diccAux.Guardar(celda.clave, celda.valor)
		}
	}
	*dicc = *diccAux
}

func (dicc *hashCerrado[K, V]) Guardar(clave K, valor V) {
	index := dicc.buscador(clave)

	dicc.tabla[index].clave = clave
	dicc.tabla[index].valor = valor
	if dicc.tabla[index].estado != _OCUPADO {
		dicc.cantidad++
		dicc.tabla[index].estado = _OCUPADO
	}
	if float64(dicc.cantidad+dicc.borrados)/float64(dicc.tam) > _FACTOR_CARGA_AUMENTAR {
		dicc.redimendionarDicc(dicc.tam * _FACTOR_REDIMENSION)
	}
}

func (dicc *hashCerrado[K, V]) Pertenece(clave K) bool {
	index := dicc.buscador(clave)
	return dicc.existeClave(clave, index)
}

func (dicc *hashCerrado[K, V]) Obtener(clave K) V {
	index := dicc.buscador(clave)
	dicc.panicClaveNoExistente(clave, index)

	return dicc.tabla[index].valor
}

func (dicc *hashCerrado[K, V]) Borrar(clave K) V {
	index := dicc.buscador(clave)
	dicc.panicClaveNoExistente(clave, index)

	valor := dicc.tabla[index].valor
	dicc.tabla[index].estado = _BORRADO
	dicc.borrados++
	dicc.cantidad--
	if float64(dicc.cantidad)/float64(dicc.tam) < _FACTOR_CARGA_REDUCIR && dicc.tam > _TAMANIO_INICIAL {
		dicc.redimendionarDicc(dicc.tam / _FACTOR_REDIMENSION)
	}
	return valor
}

func (dicc *hashCerrado[K, V]) Cantidad() int {
	return dicc.cantidad
}

// Interador interno
func (dicc *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range dicc.tabla {
		if celda.estado == _OCUPADO && !visitar(celda.clave, celda.valor) {
			return
		}
	}
}

// // Iterador

// panicFinIteracion levanta un panic con el mensaje "El iterador termino de iterar" si no hay más elementos para iterar.
func (iter *iteradorHash[K, V]) panicFinIteracion() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

// posElementoValido devuelve la posición del siguiente elemento válido de la tabla. En caso de no haberlo, devuelve _ERROR
func posElementoValido[K any, V any](tabla []celdaHash[K, V], posActual int) int {
	for i := posActual; i < len(tabla); i++ {
		if tabla[i].estado == _OCUPADO {
			return i
		}
	}
	return _ERROR
}

func (dicc *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K, V]{posActual: posElementoValido(dicc.tabla, 0), tabla: dicc.tabla, tam: dicc.tam}
}

func (iter *iteradorHash[K, V]) HaySiguiente() bool {
	if iter.posActual != -1 && iter.posActual < iter.tam {
		return true
	}
	return false
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
	iter.panicFinIteracion()
	return iter.tabla[iter.posActual].clave, iter.tabla[iter.posActual].valor
}

func (iter *iteradorHash[K, V]) Siguiente() {
	iter.panicFinIteracion()
	iter.posActual = posElementoValido(iter.tabla, iter.posActual+1)
}
