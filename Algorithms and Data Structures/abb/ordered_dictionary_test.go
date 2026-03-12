package diccionario_test

import (
	"fmt"
	"math/rand/v2"

	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAM_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func compararString(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}

func compararInts(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}

func verificarClaveNoExistente[K any, V any](t *testing.T, abb TDADiccionario.DiccionarioOrdenado[K, V], claveNoExistente K) {
	require.False(t, abb.Pertenece(claveNoExistente))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claveNoExistente) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claveNoExistente) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claveNoExistente) })
}

func verificarGuardarElemento[K any, V any](t *testing.T, abb TDADiccionario.DiccionarioOrdenado[K, V], clave K, valor V, cant int) {
	abb.Guardar(clave, valor)
	if !abb.Pertenece(clave) {
		require.EqualValues(t, cant+1, abb.Cantidad())
	}
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
}

func verificarBorrarElemento[K any, V any](t *testing.T, abb TDADiccionario.DiccionarioOrdenado[K, V], clave K, valor V, cant int) {
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.EqualValues(t, cant-1, abb.Cantidad())
	verificarClaveNoExistente(t, abb, clave)
}

// ################## test abb ###################
func TestAbbVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	require.EqualValues(t, 0, abb.Cantidad())
	verificarClaveNoExistente(t, abb, 1)
}

func TestAbbClaveDefault(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](compararString)
	require.EqualValues(t, 0, abb.Cantidad())
	verificarClaveNoExistente(t, abb, "")

	abbNum := TDADiccionario.CrearABB[int, string](compararInts)
	require.EqualValues(t, 0, abbNum.Cantidad())
	verificarClaveNoExistente(t, abbNum, 0)
}

func TestAgregarRaiz(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	require.EqualValues(t, 0, abb.Cantidad())
	verificarClaveNoExistente(t, abb, 1)

	verificarGuardarElemento(t, abb, 5, 3, abb.Cantidad())
}

func TestAbbGuardar(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](compararString)
	require.EqualValues(t, 0, abb.Cantidad())
	verificarClaveNoExistente(t, abb, claves[0])
	verificarGuardarElemento(t, abb, claves[0], valores[0], abb.Cantidad())

	require.False(t, abb.Pertenece(claves[1]))
	verificarGuardarElemento(t, abb, claves[1], valores[1], abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	verificarGuardarElemento(t, abb, claves[2], valores[2], abb.Cantidad())
}

func TestReemplazoDatoAbb(t *testing.T) {
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](compararString)
	verificarGuardarElemento(t, abb, clave, "miau", abb.Cantidad())
	verificarGuardarElemento(t, abb, clave2, "guau", abb.Cantidad())

	verificarGuardarElemento(t, abb, clave, "grrr", abb.Cantidad())
	verificarGuardarElemento(t, abb, clave2, "ruf ruf", abb.Cantidad())
	require.EqualValues(t, "grrr", abb.Obtener(clave))
	require.EqualValues(t, "ruf ruf", abb.Obtener(clave2))
}

func TestAbbBorrar(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](compararString)

	require.EqualValues(t, 0, abb.Cantidad())
	verificarClaveNoExistente(t, abb, claves[0])

	verificarGuardarElemento(t, abb, claves[0], valores[0], abb.Cantidad())
	verificarGuardarElemento(t, abb, claves[1], valores[1], abb.Cantidad())
	verificarGuardarElemento(t, abb, claves[2], valores[2], abb.Cantidad())

	verificarBorrarElemento(t, abb, claves[2], valores[2], abb.Cantidad())

	verificarBorrarElemento(t, abb, claves[0], valores[0], abb.Cantidad())

	verificarBorrarElemento(t, abb, claves[1], valores[1], abb.Cantidad())
}

func TestGuardarYBorrarRepetido(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	for i := 0; i < 1000; i++ {
		abb.Guardar(i, i)
		require.True(t, abb.Pertenece(i))
		abb.Borrar(i)
		require.False(t, abb.Pertenece(i))
	}
}

func buscarAbb(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoString(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, *int](compararString)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarAbb(cs[0], claves))
	require.NotEqualValues(t, -1, buscarAbb(cs[1], claves))
	require.NotEqualValues(t, -1, buscarAbb(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func ejecutarPruebaVolumenAbb(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](compararString)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el abb */
	// rand.Perm(n) crea un slice de 0 hasta n-1, donde los elementos estan desordenados
	valoresRand := rand.Perm(n)
	for i, j := range valoresRand {
		valores[i] = j
		claves[i] = fmt.Sprintf("%08d", j)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := range len(valoresRand) {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := range len(valoresRand) {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func Benchmark(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAM_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenAbb(b, n)
			}
		})
	}
}

func TestIterarCorteAbb(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{108, 210, 150, 403, 103, 107}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	suma := 0
	abb.Iterar(func(clave int, dato int) bool {
		if clave == 108 {
			return false
		}
		suma += clave
		return true
	})

	require.EqualValues(t, 210, suma)
}

func TestIterarHastaCortar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{1, 2, 3, 4, 5, 6, 7}
	abb.Guardar(claves[3], 30)
	abb.Guardar(claves[1], 10)
	abb.Guardar(claves[4], 40)
	abb.Guardar(claves[2], 20)
	abb.Guardar(claves[5], 50)
	abb.Guardar(claves[0], 0)

	clavesRecorridas := []int{}
	valoresRecorridos := []int{}
	abb.Iterar(func(clave int, dato int) bool {
		if dato > 30 {
			return false
		}
		clavesRecorridas = append(clavesRecorridas, clave)
		valoresRecorridos = append(valoresRecorridos, dato)
		return true
	})

	require.EqualValues(t, []int{1, 2, 3, 4}, clavesRecorridas)
	require.EqualValues(t, []int{0, 10, 20, 30}, valoresRecorridos)
}

func TestVolumenIteradorCorteAbb(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](compararInts)

	/* Inserta 'n' parejas en el abb */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarPorRangos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 43, 17}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	desde := 4
	hasta := 20
	suma := 0
	abb.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		suma += clave
		return true
	})

	require.EqualValues(t, 43, suma)
}

func TestIterarRangoDesdeNil(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 43, 17}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	hasta := 21
	suma := 0
	abb.IterarRango(nil, &hasta, func(clave, dato int) bool {
		suma += clave
		return true
	})

	require.EqualValues(t, 64, suma)
}

func TestIterarRangoConCorte(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 43, 17}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	hasta := 21
	suma := 0
	abb.IterarRango(nil, &hasta, func(clave, dato int) bool {
		if clave == 17 {
			return false
		}
		suma += clave
		return true
	})

	require.EqualValues(t, 26, suma)
}

func TestIteradorDiccionarioVacioABB(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](compararString)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIteradorPorRangos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 43, 17}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	desde := 6
	hasta := 22
	iter := abb.IteradorRango(&desde, &hasta)

	sumaClaves := 0
	sumaDatos := 0
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		sumaClaves += clave
		sumaDatos += dato
		iter.Siguiente()
	}

	require.EqualValues(t, 59, sumaClaves)
	require.EqualValues(t, 4, sumaDatos)
}

func TestIteradorRangoDesdeNil(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 43, 17}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	hasta := 22
	iter := abb.IteradorRango(nil, &hasta)

	sumaClaves := 0
	sumaDatos := 0
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		sumaClaves += clave
		sumaDatos += dato
		iter.Siguiente()
	}

	require.EqualValues(t, 64, sumaClaves)
	require.EqualValues(t, 5, sumaDatos)
}

func TestIteradorRangoHastaNil(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 43, 17}
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 1)

	desde := 6
	iter := abb.IteradorRango(&desde, nil)

	sumaClaves := 0
	sumaDatos := 0
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		sumaClaves += clave
		sumaDatos += dato
		iter.Siguiente()
	}

	require.EqualValues(t, 102, sumaClaves)
	require.EqualValues(t, 5, sumaDatos)
}

func TestInOrderIteradorRango(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{1, 2, 3, 4, 5, 6, 7}
	abb.Guardar(claves[5], 1)
	abb.Guardar(claves[3], 1)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[1], 1)
	abb.Guardar(claves[0], 1)
	abb.Guardar(claves[2], 1)
	abb.Guardar(claves[6], 1)

	iter := abb.IteradorRango(nil, nil)
	clavesAbb := make([]int, 0)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesAbb = append(clavesAbb, clave)
		iter.Siguiente()
	}
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6, 7}, clavesAbb)
}

func TestIteradorRangoDesdeHastaNil(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInts)
	claves := []int{13, 21, 5, 8, 36, 17}
	abb.Guardar(claves[0], 80)
	abb.Guardar(claves[1], 5)
	abb.Guardar(claves[2], 4)
	abb.Guardar(claves[3], 100)
	abb.Guardar(claves[4], 1)
	abb.Guardar(claves[5], 10)

	iter := abb.IteradorRango(nil, nil)

	sumaClaves := 0
	sumaDatos := 0
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		sumaClaves += clave
		sumaDatos += dato
		iter.Siguiente()
	}

	require.EqualValues(t, 100, sumaClaves)
	require.EqualValues(t, 200, sumaDatos)
}
