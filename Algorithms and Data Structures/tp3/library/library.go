package library

import (
	"errors"
	rand "math/rand"
	q "tdas/cola"
	d "tdas/diccionario"
	g "tdas/grafo"
	p "tdas/pila"
)

const (
	_COEFICIENTE_AMORTIGUACION_PAGERANK float64 = 0.85
	_ITERACIONES_PAGERANK               int     = 50
	_ITERACIONES_LABEL_PROPAGATION      int     = 30
	_CICLO_MIN                          int     = 2
	_MIN_ADY_CLUTERING                  int     = 2
)

type parametrosCfc[V comparable] struct {
	visitados  d.Diccionario[V, bool]
	pila       p.Pila[V]
	apilados   d.Diccionario[V, bool]
	orden      d.Diccionario[V, int]
	masBajo    d.Diccionario[V, int]
	cfc        []V
	origen     V
	contador   int
	encontrada bool
}

// ObtenerEntrantesYGrados obtiene los vértices de entrada y el grado de salida para cada vértice del grafo.
func ObtenerEntrantesYGrados[V comparable, W any](g g.Grafo[V, W], vertices []V) (d.Diccionario[V, []V], d.Diccionario[V, float64]) {
	entrantes := d.CrearHash[V, []V](func(a, b V) bool { return a == b })
	gradosSalida := d.CrearHash[V, float64](func(a, b V) bool { return a == b })
	for _, v := range vertices {
		entrantes.Guardar(v, []V{})
		gradosSalida.Guardar(v, float64(len(g.Adyacentes(v))))
	}

	for _, pj := range vertices {
		for _, pi := range g.Adyacentes(pj) {
			arr := entrantes.Obtener(pi)
			arr = append(arr, pj)
			entrantes.Guardar(pi, arr)
		}
	}
	return entrantes, gradosSalida
}

// PageRank devuelve un diccionario con el valor asociado al pageRank de cada vertice.
func PageRank[V comparable, W any](g g.Grafo[V, W]) d.Diccionario[V, float64] {
	vertices := g.ObtenerVertices()
	N := float64(len(vertices))

	prActual := d.CrearHash[V, float64](func(a, b V) bool { return a == b })
	prNuevo := d.CrearHash[V, float64](func(a, b V) bool { return a == b })
	for _, v := range vertices {
		prActual.Guardar(v, 1/N)
	}

	entrantes, gradosSalida := ObtenerEntrantesYGrados(g, vertices)

	for i := 0; i < _ITERACIONES_PAGERANK; i++ {
		for _, pi := range vertices {
			nuevoValor := (1 - _COEFICIENTE_AMORTIGUACION_PAGERANK) / N
			for _, pj := range entrantes.Obtener(pi) {
				prj := prActual.Obtener(pj)
				salidas := gradosSalida.Obtener(pj)
				if salidas == 0 {
					nuevoValor += _COEFICIENTE_AMORTIGUACION_PAGERANK * (prj / N)
				} else {
					nuevoValor += _COEFICIENTE_AMORTIGUACION_PAGERANK * (prj / salidas)
				}
			}
			prNuevo.Guardar(pi, nuevoValor)
		}
		prActual, prNuevo = prNuevo, prActual
	}
	return prActual
}

// ReconstruirCamino reconstuye el camino desde destino utilizando el diccionario de padres. Devueelve un error si no existe tal camino.
func ReconstruirCamino[V comparable](padres d.Diccionario[V, V], origen, destino V) ([]V, error) {
	if !padres.Pertenece(destino) {
		return nil, errors.New("No existe recorrido")
	}
	res := make([]V, 0)
	actual := destino
	for {
		res = append(res, actual)
		if actual == origen {
			break
		}
		padre := padres.Obtener(actual)
		if padre == actual {
			return nil, errors.New("No existe recorrido")
		}
		actual = padres.Obtener(actual)
	}
	return Invertir(res), nil
}

// CaminoMinimo devuelve el camino más corto entre origen y destino usando BFS.
// Devuelve nil si no existe camino.
func CaminoMinimoBFS[V comparable, W any](g g.Grafo[V, W], origen V, destino *V) (d.Diccionario[V, V], d.Diccionario[V, int], V, int) {
	padres := d.CrearHash[V, V](func(a, b V) bool { return a == b })
	dist := d.CrearHash[V, int](func(a, b V) bool { return a == b })
	visitados := d.CrearHash[V, bool](func(a, b V) bool { return a == b })
	cola := q.CrearColaEnlazada[V]()
	cola.Encolar(origen)
	padres.Guardar(origen, origen)
	dist.Guardar(origen, 0)
	visitados.Guardar(origen, true)
	distMax := 0
	ultimo := origen

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		distV := dist.Obtener(v)
		if destino != nil && v == *destino {
			return padres, dist, v, distV
		}
		for _, w := range g.Adyacentes(v) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				padres.Guardar(w, v)
				cola.Encolar(w)

				distActual := distV + 1
				dist.Guardar(w, distActual)
				if distActual > distMax {
					distMax = distActual
					ultimo = w
				}
			}
		}
	}
	return padres, dist, ultimo, distMax
}

// ObtenerPrimeraComponente devuelve la primer componente iterada del diccionario cfcs.
func ObtenerPrimeraComponente[V comparable](cfcs d.Diccionario[int, []V]) []V {
	res := make([]V, 0)
	cfcs.Iterar(func(_ int, cfc []V) bool {
		res = cfc
		return false
	})
	return res
}

// ComponentesFuertementeConexas devuelve las CFC del grafo.
func ComponentesFuertementeConexas[V comparable, W any](g g.Grafo[V, W], origen V) []V {
	params := parametrosCfc[V]{
		visitados:  d.CrearHash[V, bool](func(a, b V) bool { return a == b }),
		pila:       p.CrearPilaDinamica[V](),
		apilados:   d.CrearHash[V, bool](func(a, b V) bool { return a == b }),
		orden:      d.CrearHash[V, int](func(a, b V) bool { return a == b }),
		masBajo:    d.CrearHash[V, int](func(a, b V) bool { return a == b }),
		cfc:        make([]V, 0),
		origen:     origen,
		contador:   0,
		encontrada: false,
	}

	dfsCFC(g, origen, &params)
	return params.cfc
}

// dfsCFC obtiene las componentes fuertemente conexas del grafo recorrido con un dfs.
func dfsCFC[V comparable, W any](grafo g.Grafo[V, W], v V, params *parametrosCfc[V]) {
	if params.encontrada {
		return
	}

	params.orden.Guardar(v, params.contador)
	params.masBajo.Guardar(v, params.contador)
	params.contador++
	params.pila.Apilar(v)
	params.visitados.Guardar(v, true)
	params.apilados.Guardar(v, true)
	for _, w := range grafo.Adyacentes(v) {
		if !params.visitados.Pertenece(w) {
			dfsCFC(grafo, w, params)
		}

		if params.apilados.Pertenece(w) {
			params.masBajo.Guardar(v, min(params.masBajo.Obtener(w), params.masBajo.Obtener(v)))
		}
	}

	if params.masBajo.Obtener(v) == params.orden.Obtener(v) {
		nuevaCfc := make([]V, 0)
		for {
			w := params.pila.Desapilar()
			params.apilados.Borrar(w)
			nuevaCfc = append(nuevaCfc, w)
			if w == v {
				break
			}
		}

		if v == params.origen {
			params.cfc = nuevaCfc
			params.encontrada = true
		}
	}
}

// Ciclo obtiene un ciclo de longitud n partiendo del origen. Devuelve nil si no existe dicho ciclo.
func Ciclo[V comparable, W any](g g.Grafo[V, W], origen V, n int) []V {
	if n < _CICLO_MIN {
		return nil
	}
	camino := make([]V, 0, n)
	visitados := d.CrearHash[V, bool](func(a, b V) bool { return a == b })
	if cicloDfs(g, origen, origen, visitados, &camino, n) {
		return camino
	}
	return nil
}

func cicloDfs[V comparable, W any](g g.Grafo[V, W], v, origen V, visitados d.Diccionario[V, bool], camino *[]V, n int) bool {
	visitados.Guardar(v, true)
	*camino = append(*camino, v)
	if len(*camino) == n {
		for _, w := range g.Adyacentes(v) {
			if w == origen {
				return true
			}
		}
		visitados.Borrar(v)
		*camino = (*camino)[:len(*camino)-1]
		return false
	}
	for _, w := range g.Adyacentes(v) {
		if !visitados.Pertenece(w) {
			if cicloDfs(g, w, origen, visitados, camino, n) {
				return true
			}
		}
	}

	visitados.Borrar(v)
	*camino = (*camino)[:len(*camino)-1]
	return false
}

// GradosEntrada devuelve los grados de entrada de cada vertice en un diccionario.
func GradosEntrada[V comparable, W any](g g.Grafo[V, W]) d.Diccionario[V, int] {
	entradas := d.CrearHash[V, int](func(a, b V) bool { return a == b })
	for _, v := range g.ObtenerVertices() {
		entradas.Guardar(v, 0)
	}
	for _, v := range g.ObtenerVertices() {
		for _, w := range g.Adyacentes(v) {
			entradas.Guardar(w, entradas.Obtener(w)+1)
		}
	}
	return entradas
}

// OrdenTopologico devuelve un arreglo ordenado topologicamente. Devuelve nil si no existe dicho orden.
func OrdenTopologico[V comparable, W any](grafo g.Grafo[V, W]) []V {
	entradas := GradosEntrada(grafo)
	q := q.CrearColaEnlazada[V]()
	entradas.Iterar(func(a V, b int) bool {
		if b == 0 {
			q.Encolar(a)
		}
		return true
	})
	res := make([]V, 0)

	for !q.EstaVacia() {
		v := q.Desencolar()
		res = append(res, v)
		for _, w := range grafo.Adyacentes(v) {
			entradas.Guardar(w, entradas.Obtener(w)-1)
			if entradas.Obtener(w) == 0 {
				q.Encolar(w)
			}
		}
	}
	if grafo.Cantidad() != len(res) {
		return nil
	}
	return res
}

// Diametro devuelve el camino y el costo del diámetro del grafo.
func Diametro[V comparable, W any](grafo g.Grafo[V, W]) ([]V, int) {
	diametro := 0
	padresDiametro := d.CrearHash[V, V](func(a, b V) bool { return a == b })
	var inicioDiametro V
	var ultimoDiametro V
	for _, v := range grafo.ObtenerVertices() {
		padres, _, ultimo, distActual := CaminoMinimoBFS(grafo, v, nil)
		if distActual > diametro {
			diametro = distActual
			padresDiametro = padres
			inicioDiametro = v
			ultimoDiametro = ultimo
		}
	}
	caminoMax, _ := ReconstruirCamino(padresDiametro, inicioDiametro, ultimoDiametro)
	return caminoMax, diametro
}

// ClusteringPromedio devuelve el coeficiente de Clustering promedio del grafo g.
func ClusteringPromedio[V comparable, W any](g g.Grafo[V, W]) float64 {
	suma := 0.0
	for _, v := range g.ObtenerVertices() {
		suma += ClusteringPagina(g, v)
	}
	return suma / float64(g.Cantidad())
}

// CLusteringPagina devuelve el coeficiente de Clustering del vértice v.
func ClusteringPagina[V comparable, W any](g g.Grafo[V, W], v V) float64 {
	adyConLoops := g.Adyacentes(v)
	ady := make([]V, 0)
	for _, w := range adyConLoops {
		if w != v {
			ady = append(ady, w)
		}
	}

	k := len(ady)
	if k < _MIN_ADY_CLUTERING {
		return 0.0
	}

	contador := 0.0
	for _, x := range ady {
		for _, w := range ady {
			if x != w && g.EstanUnidos(x, w) {
				contador += 1.0
			}
		}
	}
	return contador / float64(k*(k-1))
}

// maxFreq devuelve la label con mayor frecuencia.
func maxFreq[V comparable](entrantes []V, label d.Diccionario[V, int]) int {
	frequencies := d.CrearHash[int, int](func(a, b int) bool { return a == b })
	for _, v := range entrantes {
		i := label.Obtener(v)
		if !frequencies.Pertenece(i) {
			frequencies.Guardar(i, 1)
		} else {
			frequencies.Guardar(i, frequencies.Obtener(i)+1)
		}
	}

	maxLabel := 0
	maxFreq := -1
	frequencies.Iterar(func(l, freq int) bool {
		if freq > maxFreq {
			maxFreq = freq
			maxLabel = l
		}
		return true
	})

	return maxLabel
}

// LabelPropagation devuelve la vértices comunidad utilizando el algoritmo LabelPropagation.
func LabelPropagation[V comparable, W any](g g.Grafo[V, W]) d.Diccionario[V, int] {
	vertices := g.ObtenerVertices()
	label := d.CrearHash[V, int](func(a, b V) bool { return a == b })
	for i, v := range vertices {
		label.Guardar(v, i)
	}

	entradas, _ := ObtenerEntrantesYGrados(g, vertices)

	for range _ITERACIONES_LABEL_PROPAGATION {
		rand.Shuffle(len(vertices), func(i, j int) {
			vertices[i], vertices[j] = vertices[j], vertices[i]
		})
		for _, v := range vertices {
			labelEntrantes := entradas.Obtener(v)
			if len(labelEntrantes) == 0 {
				continue
			}

			moreFreqLabel := maxFreq(labelEntrantes, label)
			label.Guardar(v, moreFreqLabel)
		}
	}

	return label
}
