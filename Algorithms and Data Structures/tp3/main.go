package tp3

import (
	"bufio"
	"errors"
	"fmt"
	i "netstats/internet"
	"os"
	"strconv"
	"strings"
	d "tdas/diccionario"
	g "tdas/grafo"
)

const (
	_SEPARADOR          string = "\t"
	_SEPARADOR_FLECHA   string = " -> "
	_SEPARADOR_COMA     string = ", "
	_LISTAR_OPERACIONES string = "listar_operaciones"
	_CAMINO             string = "camino"
	_MAS_IMPORTANTE     string = "mas_importantes"
	_CONECTADOS         string = "conectados"
	_CICLO              string = "ciclo"
	_LECTURA            string = "lectura"
	_DIAMETRO           string = "diametro"
	_EN_RANGO           string = "rango"
	_COMUNIDAD          string = "comunidad"
	_NAVEGACION         string = "navegacion"
	_CLUSTERING         string = "clustering"
	_PARTES_COMANDO     int    = 2
	_ERROR_PARAMETROS   string = "Parámetro erróneo"
)

// obtenerVerticesDeArchivo devuelve un arreglo con las primeras paginas leidas del archivo en cada linea.
func obtenerVerticesDeArchivo(file *os.File) []string {
	vertices := make([]string, 0)
	s := bufio.NewScanner(file)
	for s.Scan() {
		vertices = append(vertices, strings.Split(s.Text(), _SEPARADOR)[0])
	}
	return vertices
}

// agregarConexionesDelArchivo carga los enlaces de cada página en el grafo y devuelve un diccionario con la primer conexión del archivo de aquellas páginas que tengan al menos una conexión con otra.
func agregarConexionesDelArchivo(file *os.File, grafo g.Grafo[string, int]) d.Diccionario[string, string] {
	primerosLinks := d.CrearHash[string, string](func(a, b string) bool { return a == b })
	s := bufio.NewScanner(file)
	for s.Scan() {
		linea := strings.Split(s.Text(), _SEPARADOR)
		v := linea[0]
		enlaces := linea[1:]
		if len(enlaces) == 0 {
			continue
		}
		primerosLinks.Guardar(v, enlaces[0])
		for _, w := range enlaces {
			grafo.AgregarArista(v, w, 1)
		}
	}
	return primerosLinks
}

// cargarConexiones devuelve un grafo  con las conexiones entre links cargadas desde el archivo pasado por parámetro y un diccionario con los primeros enlaces de cada página.
// Devuelve un error si no se logró abrir el archivo.
func obtenerConexiones(rutaArchivo string) (g.Grafo[string, int], d.Diccionario[string, string], error) {
	file, err := os.Open(rutaArchivo)
	if err != nil {
		return nil, nil, errors.New("error al abrir el archivo")
	}
	defer file.Close()

	vertices := obtenerVerticesDeArchivo(file)
	grafo := g.CrearGrafo[string, int](true, vertices)

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, nil, errors.New("error al reiniciar el archivo")
	}

	primerosLinks := agregarConexionesDelArchivo(file, grafo)
	return grafo, primerosLinks, nil
}

func imprimirArticulos(articulos []string, separador string) string {
	return strings.Join(articulos, separador)
}

func EjecutarComando(comando string, parametros []string, internet i.Internet) string {
	var msj string
	switch comando {
	case _LISTAR_OPERACIONES:
		msj = strings.Join([]string{_CAMINO, _MAS_IMPORTANTE, _CONECTADOS, _CICLO, _LECTURA, _DIAMETRO, _EN_RANGO, _NAVEGACION, _CLUSTERING, _COMUNIDAD}, "\n")
	case _CAMINO:
		articulos, largo, err := internet.CaminoMasCorto(parametros[0], parametros[1])
		if err != nil {
			msj = fmt.Sprint(err)
		} else {
			msj = fmt.Sprintf("%s\nCosto: %d", imprimirArticulos(articulos, _SEPARADOR_FLECHA), largo)
		}
	case _MAS_IMPORTANTE:
		n, err := strconv.Atoi(parametros[0])
		if err != nil {
			msj = _ERROR_PARAMETROS
			break
		}
		articulos := internet.ArticulosMasImportantes(n)
		msj = imprimirArticulos(articulos, _SEPARADOR_COMA)
	case _CONECTADOS:
		conectividades, err := internet.Conectividad(parametros[0])
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		msj = imprimirArticulos(conectividades, _SEPARADOR_COMA)
	case _CICLO:
		n, err := strconv.Atoi(parametros[1])
		if err != nil {
			msj = _ERROR_PARAMETROS
			break
		}
		articulos, err := internet.CicloNArticulos(parametros[0], n)
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		articulos = append(articulos, parametros[0])
		msj = imprimirArticulos(articulos, _SEPARADOR_FLECHA)
	case _LECTURA:
		articulosOrdenados, err := internet.OrdenLectura(parametros)
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		msj = imprimirArticulos(articulosOrdenados, _SEPARADOR_COMA)
	case _DIAMETRO:
		camino, diametro := internet.Diametro()
		msj = fmt.Sprintf("%s\nCosto: %d", imprimirArticulos(camino, _SEPARADOR_FLECHA), diametro)
	case _EN_RANGO:
		n, err := strconv.Atoi(parametros[1])
		if err != nil {
			msj = _ERROR_PARAMETROS
			break
		}
		cantidad, err := internet.EnRango(parametros[0], n)
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		msj = fmt.Sprint(cantidad)
	case _NAVEGACION:
		camino, err := internet.Navegacion(parametros[0])
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		msj = imprimirArticulos(camino, _SEPARADOR_FLECHA)
	case _CLUSTERING:
		var pag *string
		if len(parametros) == 0 {
			pag = nil
		} else {
			pag = &parametros[0]
		}
		coef, err := internet.Clustering(pag)
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		msj = fmt.Sprintf("%.3f", coef)
	case _COMUNIDAD:
		comunidad, err := internet.Comunidad(parametros[0])
		if err != nil {
			msj = fmt.Sprint(err)
			break
		}
		msj = imprimirArticulos(comunidad, _SEPARADOR_COMA)
	default:
		msj = "Comando Erróneo"
	}

	return msj
}

func main() {
	grafo, primerosLinks, err := obtenerConexiones(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	internet := i.CrearInternet(grafo, primerosLinks)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		comandoEntero := strings.SplitN(s.Text(), " ", _PARTES_COMANDO)
		comando := comandoEntero[0]
		var parametros []string
		if len(comandoEntero) > 1 {
			parametros = strings.Split(comandoEntero[1], ",")
		}
		msj := EjecutarComando(comando, parametros, internet)
		fmt.Println(msj)
	}
}
