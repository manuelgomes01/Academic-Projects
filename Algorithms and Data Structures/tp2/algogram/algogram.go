package algogram

type Algogram interface {
	// Login devuelve un mensaje de bienvenida al usuario loggeado.
	// En caso de error, devuelve el mismo.
	Login(usuario string) (string, error)

	// Logout devuelve un mensaje de despedida al usuario.
	// En caso de error, devuelve el mismo.
	Logout() (string, error)

	// Publicar sube la publicación del usuario en la red para los demás.
	// En caso de error, devuelve el mismo.
	Publicar(texto string) (string, error)

	// VerSiguienteFeed devuelve un mensaje con los datos de la siguiente publicación para ver.
	// En caso de error, devuelve el mismo.
	VerSiguienteFeed() (string, error)

	// LikearPost likea el post con el id pasado por parámetro.
	// En caso de error, devuelve el mismo.
	LikearPost(idPost int) (string, error)

	// MostrarLikes devuelve un mensaje con la cantidad de likes y los usuarios que likearon el post con el id pasado por parámetro.
	// En caso de error, devuelve el mismo.
	MostrarLikes(idPost int) (string, error)
}
