package user

import (
	f "src/feed"
	p "src/post"
)

type usuario[T any] struct {
	nombre string
	feed   f.Feed[T]
}

func CrearUsuario[T any](nombre string, feed f.Feed[T]) Usuario[T] {
	return &usuario[T]{nombre: nombre, feed: feed}
}

func (u *usuario[T]) Nombre() string {
	return u.nombre
}

func (u *usuario[T]) Publicar(texto string) p.Post {
	return p.CrearPost(u.nombre, texto)
}

func (u *usuario[T]) RecibirPost(post T) {
	u.feed.AgregarPost(post)
}

func (u *usuario[T]) LikearPost(post p.Post) {
	post.Likear(u.nombre)
}

func (u *usuario[T]) VerSiguienteFeed() (T, error) {
	post, err := u.feed.Proximo()
	if err != nil {
		var zero T
		return zero, err
	}
	return post, nil
}
