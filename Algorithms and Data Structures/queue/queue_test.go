package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _CANTIDAD_ELEMENTOS_ENCOLAR = 10
const _CANTIDAD_TEST_VOLUMEN = 10000

// Prueba el comportamiento de la cola y que 'EstaVacia' responda bien mientras se hacen las operaciones basicas de la cola
func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "se quiere desencolar un elemento de la cola. como esta vacia entra en panic")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "se quiere ver el primer elemento de la cola. como esta vacia entra en panic")

	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Desencolar())

	require.True(t, cola.EstaVacia())

}

// Prueba el comportamiento de la cola cuando se encola y desencola un unico elemento
func TestColaUnElemento(t *testing.T) {
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaBool.Encolar(true)
	require.False(t, colaBool.EstaVacia())
	require.Equal(t, true, colaBool.VerPrimero())
	require.Equal(t, true, colaBool.Desencolar())
	require.True(t, colaBool.EstaVacia())
}

// Prueba el comportamiento de la cola cuando de encolan y desencolan '_CANTIDAD_ELEMENTOS_APILAR' elementos tipo struct
func TestEncolarVariosElementos(t *testing.T) {
	type test_t struct {
		x int
		y float64
		a string
		k bool
	}

	estructura := test_t{x: 23, y: 2, a: "hola", k: true}

	colaStruct := TDACola.CrearColaEnlazada[test_t]()
	require.True(t, colaStruct.EstaVacia())

	for _ = range _CANTIDAD_ELEMENTOS_ENCOLAR {
		colaStruct.Encolar(estructura)
		require.False(t, colaStruct.EstaVacia())
	}

	for _ = range _CANTIDAD_ELEMENTOS_ENCOLAR {
		require.Equal(t, estructura, colaStruct.VerPrimero())
		require.Equal(t, estructura, colaStruct.Desencolar())
	}

	require.True(t, colaStruct.EstaVacia())
}

// Prueba el comportamiento de la cola con 'string' como tipo de dato
func TestComportamientoColaString(t *testing.T) {
	var (
		elemento1 string = "hola mundo"
		elemento2 string = "chau universo"
		elemento3 string = "test de cola"
	)

	colaString := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaString.EstaVacia())
	colaString.Encolar(elemento1)
	require.False(t, colaString.EstaVacia())
	require.Equal(t, elemento1, colaString.VerPrimero())
	colaString.Encolar(elemento2)
	require.False(t, colaString.EstaVacia())
	require.Equal(t, elemento1, colaString.VerPrimero())
	colaString.Encolar(elemento3)
	require.False(t, colaString.EstaVacia())
	require.Equal(t, elemento1, colaString.VerPrimero())

	require.Equal(t, elemento1, colaString.Desencolar())
	require.False(t, colaString.EstaVacia())
	require.Equal(t, elemento2, colaString.VerPrimero())
	require.Equal(t, elemento2, colaString.Desencolar())
	require.False(t, colaString.EstaVacia())
	require.Equal(t, elemento3, colaString.VerPrimero())
	require.Equal(t, elemento3, colaString.Desencolar())

	require.PanicsWithValue(t, "La cola esta vacia", func() { colaString.Desencolar() }, "se quiere desencolar un elemento de la cola. como esta vacia entra en panic")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaString.VerPrimero() }, "se quiere ver el primer elemento de la cola. como esta vacia entra en panic")

	require.True(t, colaString.EstaVacia())
}

// Prueba el comportamiento de la cola con 'float' como tipo de dato
func TestComportamientoColaFloat(t *testing.T) {
	var (
		elemento4 float64 = 3.14
		elemento5 float64 = 2.718
		elemento6 float64 = 5.89
		elemento7 float64 = 7.3124
	)

	colaFloat := TDACola.CrearColaEnlazada[float64]()
	require.True(t, colaFloat.EstaVacia())

	colaFloat.Encolar(elemento4)
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento4, colaFloat.VerPrimero())
	colaFloat.Encolar(elemento5)
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento4, colaFloat.VerPrimero())
	colaFloat.Encolar(elemento6)
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento4, colaFloat.VerPrimero())
	colaFloat.Encolar(elemento7)
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento4, colaFloat.VerPrimero())

	require.Equal(t, elemento4, colaFloat.Desencolar())
	require.Equal(t, elemento5, colaFloat.VerPrimero())
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento5, colaFloat.Desencolar())
	require.Equal(t, elemento6, colaFloat.VerPrimero())
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento6, colaFloat.Desencolar())
	require.Equal(t, elemento7, colaFloat.VerPrimero())
	require.False(t, colaFloat.EstaVacia())
	require.Equal(t, elemento7, colaFloat.Desencolar())

	require.PanicsWithValue(t, "La cola esta vacia", func() { colaFloat.VerPrimero() }, "se quiere ver el primer elemento de la cola. como esta vacia entra en panic")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaFloat.Desencolar() }, "se quiere desencolar un elemento de la cola. como esta vacia entra en panic")

	require.True(t, colaFloat.EstaVacia())

}

// Prueba encolar y desencolar '_CANTIDAD_TEST_VOLUMEN' elementos para ver que el comportamiento de la cola sea el correcto
func TestVolumen(t *testing.T) {

	colaInt := TDACola.CrearColaEnlazada[int]()
	require.True(t, colaInt.EstaVacia())
	for i := range _CANTIDAD_TEST_VOLUMEN {
		colaInt.Encolar(i)
		require.Equal(t, 0, colaInt.VerPrimero())
		require.False(t, colaInt.EstaVacia())
	}

	for i := range _CANTIDAD_TEST_VOLUMEN {
		require.False(t, colaInt.EstaVacia())
		require.Equal(t, i, colaInt.Desencolar())
	}
	require.True(t, colaInt.EstaVacia())

}
