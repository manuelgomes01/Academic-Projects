package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const _CANTIDAD_ELEMENTOS_APILAR = 10
const _CANTIDAD_TEST_VOLUMEN = 10000

// Prueba el comportamiento de la pila y que 'EstaVacia' responda bien mientras se hacen las operaciones basicas de la pila
func TestPilaVacia(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.VerTope() }, "se quiere ver el tope de la pila. como la pila esta vacia, entrane panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaInt.Desapilar() }, "se quiere desapilar un elemento. como la pila esta vacia, entrane panic")

	pilaInt.Apilar(1)
	require.False(t, pilaInt.EstaVacia())
	require.Equal(t, 1, pilaInt.VerTope())
	require.Equal(t, 1, pilaInt.Desapilar())
	require.True(t, pilaInt.EstaVacia())

	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaString.Apilar("a")
	require.Equal(t, "a", pilaString.VerTope())
	require.False(t, pilaString.EstaVacia())
}

// Prueba el comportamiento de la pila cuando de apila y desapila un unico elemento
func TestApilarUnElemento(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())
	pilaInt.Apilar(123456)
	require.False(t, pilaInt.EstaVacia())
	require.Equal(t, 123456, pilaInt.VerTope())
	require.Equal(t, 123456, pilaInt.Desapilar())
	require.True(t, pilaInt.EstaVacia())
}

// Prueba el comportamiento de la pila cuando de apilan y desapilan '_CANTIDAD_ELEMENTOS_APILAR' elementos
func TestApilarVariosElementos(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())

	for i := range _CANTIDAD_ELEMENTOS_APILAR {
		pilaInt.Apilar(i)
		require.False(t, pilaInt.EstaVacia())
		require.Equal(t, i, pilaInt.VerTope())
	}

	for i := _CANTIDAD_ELEMENTOS_APILAR - 1; i >= 0; i-- {
		require.Equal(t, i, pilaInt.VerTope())
		require.Equal(t, i, pilaInt.Desapilar())
	}

	require.True(t, pilaInt.EstaVacia())
}

// Prueba el comportamiento de la pila con 'string' como tipo de dato
func TestComportamientoPilaString(t *testing.T) {
	var (
		elemento1 string = "a"
		elemento2 string = "abc"
		elemento3 string = "abcd"
	)

	pilaString := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaString.EstaVacia())
	pilaString.Apilar(elemento1)
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, elemento1, pilaString.VerTope())
	pilaString.Apilar(elemento2)
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, elemento2, pilaString.VerTope())
	pilaString.Apilar(elemento3)
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, elemento3, pilaString.VerTope())

	require.Equal(t, elemento3, pilaString.Desapilar())
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, elemento2, pilaString.VerTope())
	require.Equal(t, elemento2, pilaString.Desapilar())
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, elemento1, pilaString.VerTope())
	require.Equal(t, elemento1, pilaString.Desapilar())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.VerTope() }, "se quiere ver el tope de la pila. como la pila esta vacia, entrane panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.Desapilar() }, "se quiere desapilar un elemento. como la pila esta vacia, entrane panic")

	require.True(t, pilaString.EstaVacia())
}

// Prueba el comportamiento de la pila con 'bool' como tipo de dato y '_CANTIDAD_ELEMENTOS_APILAR' de elementos apilados y desapilados
func TestComportamientoPilaBool(t *testing.T) {
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	require.True(t, pilaBool.EstaVacia())
	pilaBool.Apilar(true)
	require.Equal(t, true, pilaBool.VerTope())
	require.False(t, pilaBool.EstaVacia())
	require.Equal(t, true, pilaBool.Desapilar())
	require.True(t, pilaBool.EstaVacia())

	for _ = range _CANTIDAD_ELEMENTOS_APILAR {
		pilaBool.Apilar(true)
		require.False(t, pilaBool.EstaVacia())
		require.Equal(t, true, pilaBool.VerTope())
	}

	for !pilaBool.EstaVacia() {
		require.Equal(t, true, pilaBool.VerTope())
		require.Equal(t, true, pilaBool.Desapilar())
	}

	require.True(t, pilaBool.EstaVacia())
}

// Prueba el comportamiento de la pila con 'float' como tipo de dato
func TestComportamientoPilaFloat(t *testing.T) {
	var (
		elemento4 float64 = 3.14
		elemento5 float64 = 2.718
		elemento6 float64 = 5.89
	)

	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pilaFloat.EstaVacia())
	pilaFloat.Apilar(elemento4)
	require.False(t, pilaFloat.EstaVacia())
	require.Equal(t, elemento4, pilaFloat.VerTope())
	pilaFloat.Apilar(elemento5)
	require.False(t, pilaFloat.EstaVacia())
	require.Equal(t, elemento5, pilaFloat.VerTope())
	pilaFloat.Apilar(elemento6)
	require.False(t, pilaFloat.EstaVacia())
	require.Equal(t, elemento6, pilaFloat.VerTope())

	require.Equal(t, elemento6, pilaFloat.Desapilar())
	require.False(t, pilaFloat.EstaVacia())
	require.Equal(t, elemento5, pilaFloat.VerTope())
	require.Equal(t, elemento5, pilaFloat.Desapilar())
	require.False(t, pilaFloat.EstaVacia())
	require.Equal(t, elemento4, pilaFloat.VerTope())
	require.Equal(t, elemento4, pilaFloat.Desapilar())

	require.True(t, pilaFloat.EstaVacia())

}

// Prueba apilar y desapilar '_CANTIDAD_TEST_VOLUMEN' elementos para ver que el comportamiento de la pila sea el correcto
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia())

	for i := range _CANTIDAD_TEST_VOLUMEN {
		pila.Apilar(i)
		require.False(t, pila.EstaVacia())
		require.Equal(t, i, pila.VerTope())
	}

	for i := _CANTIDAD_TEST_VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
}
