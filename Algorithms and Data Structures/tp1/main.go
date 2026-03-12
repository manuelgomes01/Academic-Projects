package tp1

import (
	"bufio"
	"dc/calculadora"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

const _ERROR = -1

// HacerOperaciones devuelve el resultado de la operacion correspondiente o ERROR en caso de que:
//   - no se hayan ingresadon la cantidad de argumentos para realizar operaciones
//   - se hayan ingresado argumentos de mas para realizar operaciones
//   - se ingrese un simbolo invalido
func HacerOperaciones(entrada []string) (int64, error) {
	if len(entrada) == 1 {
		return _ERROR, errors.New("cantidad insuficiente de argumentos")
	}
	pila := TDAPila.CrearPilaDinamica[int64]()
	for _, valor := range entrada {
		if valor != "" {
			if calculadora.EsOperadorValido(valor) {
				errorCalculo := calculadora.RealizarCalculo(valor, pila)
				if errorCalculo != nil {
					return _ERROR, errors.New("error al realizar el calculo")
				}
			} else {
				digito, errorDigito := strconv.Atoi(valor)
				if errorDigito != nil {
					return _ERROR, errors.New("digito invalido")
				} else {
					pila.Apilar(int64(digito))
				}
			}
		}
	}

	resultado := pila.Desapilar()
	if !pila.EstaVacia() {
		return _ERROR, errors.New("cantidad de argumentos superior a la requerida")
	}

	return resultado, nil
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := strings.Fields(s.Text())
		resultado, error := HacerOperaciones(entrada)
		if error != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(resultado)
		}
	}
}
