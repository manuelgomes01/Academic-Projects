package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const RUTA_ARCHIVO_1 = "archivo1.in"
const RUTA_ARCHIVO_2 = "archivo2.in"

/*
 * Precondicion:
 * Postcondicion: Muestra por pantalla el vector correspondiente
 */
func mostrarArreglo(slice []int) {
	ejercicios.Seleccion(slice)
	for i := range len(slice) {
		fmt.Printf("%d\n", slice[i])
	}
}

/*
 * Precondicion:
 * Postcondicion: Determina cual de los dos vectores es el mayor para mostrarlo por pantalla
 */
func queArregloMostrar(slice1 []int, slice2 []int) {
	resultado := ejercicios.Comparar(slice1, slice2)
	if resultado == 1 {
		mostrarArreglo(slice1)
	} else {
		mostrarArreglo(slice2)
	}
}

func llenarSlice(rutaArchvio string) []int {
	archivo, _ := os.Open(rutaArchvio)
	defer archivo.Close()
	slice := []int{}
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		i, _ := strconv.Atoi(lectura.Text())
		slice = append(slice, i)
	}

	return slice
}

func main() {

	slice1 := llenarSlice(RUTA_ARCHIVO_1)
	slice2 := llenarSlice(RUTA_ARCHIVO_2)

	queArregloMostrar(slice1, slice2)
}
