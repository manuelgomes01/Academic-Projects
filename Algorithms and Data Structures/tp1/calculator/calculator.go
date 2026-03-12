package calculator

import (
	"dc/calculadora/operaciones"
	"errors"
	TDAPila "tdas/pila"
)

const _OPERADOR_SUMA = "+"
const _OPERADOR_RESTA = "-"
const _OPERADOR_MULTIPLICACION = "*"
const _OPERADOR_DIVISION = "/"
const _OPERADOR_RAIZ_CUADRADA = "sqrt"
const _OPERADOR_EXPONENCIAL = "^"
const _OPERADOR_LOGARITMO = "log"
const _OPERADOR_TERNARIO = "?"

type operacion struct {
	simbolo string
	aridad  int
	operar  func(operadores []int64) (int64, error)
}

// EsOperadorValido devuelve true si se ingreso un operador valido. En caso contrario devuelve false
func EsOperadorValido(operador string) bool {
	return operador == _OPERADOR_SUMA || operador == _OPERADOR_RESTA || operador == _OPERADOR_MULTIPLICACION || operador == _OPERADOR_DIVISION || operador == _OPERADOR_RAIZ_CUADRADA || operador == _OPERADOR_EXPONENCIAL || operador == _OPERADOR_LOGARITMO || operador == _OPERADOR_TERNARIO
}

func llenarArreglo(aridad int, pila TDAPila.Pila[int64]) []int64 {
	resultado := []int64{}
	i := 0
	for i < aridad && !pila.EstaVacia() {
		resultado = append(resultado, pila.Desapilar())
		i++
	}

	return resultado
}

// validarResultado devuelve el resultado de la operacion correspondiente, o ERROR en caso que 'op.operar(operadores)'
// devuelva ERROR distinto de 'nil'
func validarResultado(op operacion, pila TDAPila.Pila[int64]) error {
	operadores := llenarArreglo(op.aridad, pila)

	if len(operadores) != op.aridad {
		return errors.New("cantidad de operadores invalida para la operacion")
	}

	resultado, errorOperacion := op.operar(operadores)
	if errorOperacion != nil {
		return errorOperacion
	} else {
		pila.Apilar(resultado)
	}
	return nil
}

func RealizarCalculo(operador string, pila TDAPila.Pila[int64]) error {
	operandos := []operacion{
		{simbolo: _OPERADOR_SUMA, aridad: 2, operar: operaciones.CalcularSuma},
		{simbolo: _OPERADOR_RESTA, aridad: 2, operar: operaciones.CalcularResta},
		{simbolo: _OPERADOR_MULTIPLICACION, aridad: 2, operar: operaciones.CalcularMultiplicacion},
		{simbolo: _OPERADOR_DIVISION, aridad: 2, operar: operaciones.CalcularDivision},
		{simbolo: _OPERADOR_RAIZ_CUADRADA, aridad: 1, operar: operaciones.CalcularRaiz},
		{simbolo: _OPERADOR_EXPONENCIAL, aridad: 2, operar: operaciones.CalcularExponencial},
		{simbolo: _OPERADOR_LOGARITMO, aridad: 2, operar: operaciones.CalcularLogaritmo},
		{simbolo: _OPERADOR_TERNARIO, aridad: 3, operar: operaciones.CalcularTernario},
	}

	for _, op := range operandos {
		if op.simbolo == operador {
			return validarResultado(op, pila)
		}
	}

	return errors.New("error")
}
