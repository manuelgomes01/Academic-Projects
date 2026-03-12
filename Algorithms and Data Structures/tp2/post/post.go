package post

type Post interface {
	// Autor devuelve el autor de la publicación.
	Autor() string

	// VerPost devuelve el texto de la publicación.
	VerPost() string

	// Likear añade el like al post del usuario pasado por parámetro.
	Likear(usuario string)

	// CantidadLikes devuelve la cantidad de likes que tiene el post.
	CantidadLikes() int

	// MostrarLikes devuelve la cantidad de likes que tiene el post y los usuarios que lo likearon.
	MostrarLikes() (int, []string)
}
