package internet

type Internet interface {
	// devuelve una lista con las páginas con los cuales navegamos de la página origen a la página destino, navegando lo menos posible.
	CaminoMasCorto(origen, destino string) ([]string, int, error) //O(P+L), P cant pags y L cant de links (aristas)

	// devuelve las n páginas más centrales/importantes del mundo según el algoritmo de pagerank, ordenadas de mayor importancia a menor importancia.
	// O(K(P+L)+Plog(n)), siendo K la cantidad de iteraciones a realizar para llegar a la convergencia (puede simplificarse a O(P logn + L).
	// (puede simplificarse a O(P logn + L) (El término O(P log⁡ n) proviene de obtener los Top-n luego de haber aplicado el algoritmo)
	ArticulosMasImportantes(n int) []string

	// devuelve todas las páginas a los que podemos llegar desde la página pasado por parámetro y que, a su vez, puedan también volver a dicha página.
	// Este comando debe ejecutar en O(P+L). Considerar que a todas las páginas a las que lleguemos también se conectan entre sí, y con el tamaño del set de datos puede convenir guardar los resultados.
	// IMPORTANTE: En ambos casos la CFC está compuesta por los mismos artículos. Es importante notar que la segunda consulta debería obtener un resultado en tiempo constante.
	Conectividad(art string) ([]string, error)

	// permite obtener un ciclo de largo n que comience en la página indicada.
	// en O(P^n). Es importante considerar realizar todas las optimizaciones posibles para raducir (empíricamente) el tiempo de ejecución del algoritmo.
	CicloNArticulos(art string, n int) ([]string, error)

	//  Permite obtener un orden en el que es válido leer las páginas indicados. Para que un orden sea válido, si página_i tiene un link a página_j, entonces es necesario primero leer página_j
	// en O(n+L_n), siendo n la cantidad de páginas indicadas, y L_n la cantidad de links entre estas.
	OrdenLectura(articulos []string) ([]string, error)

	// permite obtener el diámetro de toda la red. Esto es, obtener el camino mínimo más grande de toda la red.
	Diametro() ([]string, int) // en O(P(P+L))

	// permite obtener la cantidad de páginas que se encuenten a exactamente n links/saltos desde la página pasada por parámetro
	EnRango(art string, n int) (int, error) // O(P + L)

	// Navegacion devuelve el camino navegando únicamente por el primer link encontrado.
	Navegacion(origen string) ([]string, error) // O(n)

	// Clustering devuelve el coeficiente de Clustering de la página pasada por parámetros.
	// Si pagina es nil, entonces devuelve el coeficiente promedio de la red.
	// En caso de pagina no existente, devuelve un error.
	Clustering(pagina *string) (float64, error)

	// permite obtener la comunidad dentro de la red a la que pertenezca la página pasada por parámetro.
	// Devuelve un error en caso de que la página no exista.
	Comunidad(pagina string) ([]string, error)
}
