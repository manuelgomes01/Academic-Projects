package user

import (
	p "estructuras/post"
)

type Usuario[T any] interface {
	// Nombre devuelve el nombre del usuario.
	Nombre() string

	// Publicar devuelve el post publicado.
	Publicar(texto string) p.Post

	// RecibirPost añade al feed del usuario la publicación recibida por parámetros.
	RecibirPost(post T)

	// LikearPost likea el post pasado por parámetros.
	LikearPost(post p.Post)

	// VerSiguienteFeed devuelve la siguiente publicación del feed del usuario.
	// En caso de error, devuelve el mismo.
	VerSiguienteFeed() (T, error)
}
