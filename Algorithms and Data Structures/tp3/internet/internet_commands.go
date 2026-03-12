package internet

import (
	"errors"
	biblioteca "netstats/biblioteca"
	d "tdas/diccionario"
	g "tdas/grafo"
)

const (
	_PAGINA_NO_EXISTENTE   string = "Página no existente"
	_RECORRIDO_INEXISTENTE string = "No se encontro recorrido"
	_ORDEN_INEXISTENTE     string = "No existe forma de leer las paginas en orden"
	_NAVEGACIONES          int    = 21
)

type estadosDiametro struct {
	diametro    int
	componentes []string
	calculado   bool
}

type internet struct {
	grafo         g.Grafo[string, int]
	primerLink    d.Diccionario[string, string]
	pagerank      d.Diccionario[string, float64]
	cfcs          d.Diccionario[int, []string]
	idsCfcs       d.Diccionario[string, int]
	diametro      estadosDiametro
	idIncremental int
	labels        d.Diccionario[string, int]
}

// CrearInternet devuelve un TDA Internet en estado válido con los elementos de conexiones como vértices.
func CrearInternet(conexiones g.Grafo[string, int], primerosLinks d.Diccionario[string, string]) Internet {
	return &internet{
		grafo:         conexiones,
		primerLink:    primerosLinks,
		pagerank:      nil,
		cfcs:          d.CrearHash[int, []string](func(a, b int) bool { return a == b }),
		idsCfcs:       d.CrearHash[string, int](func(a, b string) bool { return a == b }),
		idIncremental: 0,
	}
}

func (inter *internet) CaminoMasCorto(origen, destino string) ([]string, int, error) {
	if !inter.grafo.ExisteVertice(origen) || !inter.grafo.ExisteVertice(destino) {
		return nil, 0, errors.New(_PAGINA_NO_EXISTENTE)
	}
	padres, _, _, _ := biblioteca.CaminoMinimoBFS(inter.grafo, origen, &destino)
	camino, err := biblioteca.ReconstruirCamino(padres, origen, destino)
	if err != nil {
		return nil, 0, errors.New(_RECORRIDO_INEXISTENTE)
	}
	costo := len(camino) - 1
	return camino, costo, nil
}

func (inter *internet) ArticulosMasImportantes(n int) []string {
	if inter.pagerank == nil {
		inter.pagerank = biblioteca.PageRank(inter.grafo)
	}
	return biblioteca.TopK(n, inter.pagerank)
}

func (inter *internet) Conectividad(pag string) ([]string, error) {
	if !inter.grafo.ExisteVertice(pag) {
		return nil, errors.New(_PAGINA_NO_EXISTENTE)
	}

	if inter.idsCfcs.Pertenece(pag) {
		return inter.cfcs.Obtener(inter.idsCfcs.Obtener(pag)), nil
	}

	cfc := biblioteca.ComponentesFuertementeConexas(inter.grafo, pag)

	id := inter.idIncremental
	inter.idIncremental += 1
	for _, v := range cfc {
		inter.idsCfcs.Guardar(v, id)
	}
	inter.cfcs.Guardar(id, cfc)

	return cfc, nil
}

func (inter *internet) CicloNArticulos(pag string, n int) ([]string, error) {
	if !inter.grafo.ExisteVertice(pag) {
		return nil, errors.New(_PAGINA_NO_EXISTENTE)
	}
	ciclo := biblioteca.Ciclo(inter.grafo, pag, n)

	if ciclo == nil {
		return nil, errors.New(_RECORRIDO_INEXISTENTE)
	}

	return ciclo, nil
}

func (inter *internet) crearGrafoAux(paginas []string) (g.Grafo[string, int], error) {
	grafoAux := g.CrearGrafo[string, int](true, paginas)

	for _, v := range paginas {
		if !inter.grafo.ExisteVertice(v) {
			return nil, errors.New(_PAGINA_NO_EXISTENTE)
		}
		for _, w := range paginas {
			if inter.grafo.EstanUnidos(v, w) {
				grafoAux.AgregarArista(w, v, 1)
			}
		}
	}
	return grafoAux, nil
}

func (inter *internet) OrdenLectura(paginas []string) ([]string, error) {
	grafoAux, err := inter.crearGrafoAux(paginas)
	if err != nil {
		return nil, err
	}
	orden := biblioteca.OrdenTopologico(grafoAux)
	if orden == nil {
		return nil, errors.New(_ORDEN_INEXISTENTE)
	}
	return orden, nil
}

func (inter *internet) Diametro() ([]string, int) {
	if inter.diametro.calculado {
		return inter.diametro.componentes, inter.diametro.diametro
	}
	inter.diametro.componentes, inter.diametro.diametro = biblioteca.Diametro(inter.grafo)
	inter.diametro.calculado = true
	return inter.diametro.componentes, inter.diametro.diametro
}

func (inter *internet) EnRango(pagina string, n int) (int, error) {
	if !inter.grafo.ExisteVertice(pagina) {
		return 0, errors.New(_PAGINA_NO_EXISTENTE)
	}
	_, distancias, _, _ := biblioteca.CaminoMinimoBFS(inter.grafo, pagina, nil)
	contador := 0
	distancias.Iterar(func(_ string, dist int) bool {
		if dist == n {
			contador += 1
		}
		return true
	})
	return contador, nil
}

func (inter *internet) Navegacion(origen string) ([]string, error) {
	if !inter.grafo.ExisteVertice(origen) {
		return nil, errors.New(_PAGINA_NO_EXISTENTE)
	}
	camino := make([]string, 0)
	actual := origen
	for i := 0; i < _NAVEGACIONES; i++ {
		camino = append(camino, actual)
		if !inter.primerLink.Pertenece(actual) {
			break
		}
		actual = inter.primerLink.Obtener(actual)
	}
	return camino, nil
}

func (inter *internet) Clustering(pagina *string) (float64, error) {
	if pagina != nil && !inter.grafo.ExisteVertice(*pagina) {
		return 0.0, errors.New(_PAGINA_NO_EXISTENTE)
	}

	if pagina != nil {
		return biblioteca.ClusteringPagina(inter.grafo, *pagina), nil
	}
	return biblioteca.ClusteringPromedio(inter.grafo), nil
}

func (inter *internet) Comunidad(pagina string) ([]string, error) {
	if !inter.grafo.ExisteVertice(pagina) {
		return nil, errors.New(_PAGINA_NO_EXISTENTE)
	}

	if inter.labels == nil {
		inter.labels = biblioteca.LabelPropagation(inter.grafo)
	}
	labelPagina := inter.labels.Obtener(pagina)

	comunidad := make([]string, 0)
	inter.labels.Iterar(func(v string, l int) bool {
		if l == labelPagina {
			comunidad = append(comunidad, v)
		}
		return true
	})

	return comunidad, nil
}
