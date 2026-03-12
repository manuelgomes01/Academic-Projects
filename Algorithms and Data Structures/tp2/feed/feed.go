package feed

type Feed[T any] interface {
	// AgregarPost añade el post pasado por parámetro al feed.
	AgregarPost(post T)

	// Proximo devuelve el siguiente elemento del feed. En caso de error, devuelve el mismo.
	Proximo() (T, error)
}
