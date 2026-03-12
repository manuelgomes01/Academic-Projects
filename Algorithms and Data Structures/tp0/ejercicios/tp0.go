package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x

}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}

	var posMaximo int
	for j := 1; j < len(vector); j++ {
		if vector[j] > vector[posMaximo] {
			posMaximo = j
		}
	}

	return posMaximo
}

/*
 * Precondicion:
 * Postcondicion: Devuelve el largo del vector mas corto
 */
func cualEsMasCorto(largoVector1 int, largoVector2 int) int {
	if largoVector1 < largoVector2 {
		return largoVector1
	} else {
		return largoVector2
	}
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	menorLongitud := cualEsMasCorto(len(vector1), len(vector2))

	for i := range menorLongitud {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}

	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
	}

	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector) - 1; i > 0; i-- {
		posicionMax := Maximo(vector[:i+1])
		Swap(&vector[i], &vector[posicionMax])
	}
}

/*
 * Precondicion: "vector", "indice" y "contador" deben estar previamente inicializados
 * Postcondicion: Devuelve la suma de los elementos del arreglo
 */
func sumatoria(vector []int, indice int, contador int) int {
	if indice == len(vector) {
		return contador
	}

	contador += vector[indice]

	return sumatoria(vector, indice+1, contador)
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	var indice, contador int
	return sumatoria(vector, indice, contador)
}

/*
 * Precondicion: "cadena", "posicionInicio" y "posicionFin" deben estar previamente inicializados
 * Postcondicion: Devuelve "0" es caso de serlo o "-1" en caso contrario
 */
func palabraCapicua(cadena string, posicionInicio int, posicionFin int) bool {
	if (posicionInicio == posicionFin) || (posicionInicio > posicionFin) {
		return true
	}

	if cadena[posicionInicio] != cadena[posicionFin] {
		return false
	}

	return palabraCapicua(cadena, posicionInicio+1, posicionFin-1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	var posicionInicio int
	posicionFin := len(cadena) - 1

	return palabraCapicua(cadena, posicionInicio, posicionFin)
}
