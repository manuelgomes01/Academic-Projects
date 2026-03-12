package operations

import (
	"errors"
	"math"
)

const _ERROR = -1

// CalcularTernario devuelve: 'operadores[0]' si 'operadores[2]' es igual a 0
// o 'operadores[1]' para cualquier otro valor de 'operadores[2]'
func CalcularTernario(operadores []int64) (int64, error) {
	if operadores[2] == 0 {
		return operadores[0], nil
	}
	return operadores[1], nil
}

// CalcularLogaritmo devuelve el resultado del logaritmo de 'a' en base 'b' o ERROR en caso de que:
//   - 'operadores[1]' sea menor o igual a 0
//   - 'operadores[0]' sea menor a 2
func CalcularLogaritmo(operadores []int64) (int64, error) {
	if operadores[0] < 2 {
		return _ERROR, errors.New("no se puede calcular el logaritmo en base menor a 2")
	}
	if operadores[1] <= 0 {
		return _ERROR, errors.New("el argumento debe ser mayor a 0")
	}

	base := math.Log(float64(operadores[0]))
	argumento := math.Log(float64(operadores[1]))

	return int64(argumento / base), nil
}

// CalcularExponencial devuelve el resultado de la exponencial o ERROR en caso de que 'operadores[0]' sea menor a 0
func CalcularExponencial(operadores []int64) (int64, error) {
	if operadores[0] < 0 {
		return _ERROR, errors.New("no se puede elevar por un numero negativo")
	}

	return int64(math.Pow(float64(operadores[1]), float64(operadores[0]))), nil
}

// CalcularRaiz devuelve el resultado de la raiz cuadrada o ERROR en caso de que 'operadores[0]' sea menor a 0
func CalcularRaiz(operadores []int64) (int64, error) {
	if operadores[0] < 0 {
		return _ERROR, errors.New("no se puede calcular la raiz cuadrada de un numero negativo")
	}

	return int64(math.Sqrt(float64(operadores[0]))), nil
}

// CalcularDivision devuelve el resultado de la division o ERROR en caso de que 'operadores[0]' sea igual a 0
func CalcularDivision(operadores []int64) (int64, error) {
	if operadores[0] == 0 {
		return _ERROR, errors.New("no se puede dividir por 0")
	}

	return operadores[1] / operadores[0], nil
}

// CalcularMultiplicacion devuelve el resultado de la multiplicacion
func CalcularMultiplicacion(operadores []int64) (int64, error) {
	return operadores[1] * operadores[0], nil
}

// CalcularResta devuelve el resultado de la resta
func CalcularResta(operadores []int64) (int64, error) {
	return operadores[1] - operadores[0], nil
}

// CalcularSuma devuelve el resultado de la suma
func CalcularSuma(operadores []int64) (int64, error) {
	return operadores[1] + operadores[0], nil
}
