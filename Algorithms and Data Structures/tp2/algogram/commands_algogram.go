package algogram

import (
	"errors"
	f "estructuras/feed"
	p "estructuras/post"
	u "estructuras/usuario"
	"fmt"
	dicc "tdas/diccionario"
)

type feedUsuario struct {
	post     p.Post
	idPost   int
	afinidad int
}

type usuarioAlgogram struct {
	pos     int
	usuario u.Usuario[feedUsuario]
}

type algogram struct {
	usuarios dicc.Diccionario[string, usuarioAlgogram]
	posts    dicc.Diccionario[int, p.Post]
	loggeado u.Usuario[feedUsuario]
}

func afinidadUsuarios(a, b feedUsuario) int {
	if a.afinidad == b.afinidad {
		return b.idPost - a.idPost
	}
	return b.afinidad - a.afinidad
}

func calcularAfinidad(a, b int) int {
	afinidad := a - b
	if afinidad < 0 {
		afinidad *= -1
	}
	return afinidad
}

func CrearAlgogramConUsuarios(usuarios []string) (Algogram, error) {
	users := dicc.CrearHash[string, usuarioAlgogram](func(a, b string) bool { return a == b })
	for i := 0; i < len(usuarios); i++ {
		feed := f.CrearFeed(afinidadUsuarios)
		user := usuarioAlgogram{pos: i, usuario: u.CrearUsuario(usuarios[i], feed)}
		users.Guardar(usuarios[i], user)
	}

	return &algogram{
		usuarios: users,
		posts:    dicc.CrearHash[int, p.Post](func(a, b int) bool { return a == b }),
		loggeado: nil,
	}, nil
}

func (app *algogram) Login(nombre string) (string, error) {
	if app.loggeado != nil {
		return "", errors.New("Error: Ya habia un usuario loggeado")
	}
	if !app.usuarios.Pertenece(nombre) {
		return "", errors.New("Error: usuario no existente")
	}
	app.loggeado = app.usuarios.Obtener(nombre).usuario
	return fmt.Sprintf("Hola %s\n", nombre), nil
}

func (app *algogram) Logout() (string, error) {
	if app.loggeado == nil {
		return "", errors.New("Error: no habia usuario loggeado")
	}
	app.loggeado = nil
	return fmt.Sprintln("Adios"), nil
}

func (app *algogram) Publicar(texto string) (string, error) {
	if app.loggeado == nil {
		return "", errors.New("Error: no habia usuario loggeado")
	}
	publicado := app.loggeado.Publicar(texto)
	app.posts.Guardar(app.posts.Cantidad(), publicado)

	posPublicador := app.usuarios.Obtener(app.loggeado.Nombre()).pos
	app.usuarios.Iterar(func(nombre string, us usuarioAlgogram) bool {
		if nombre != app.loggeado.Nombre() {
			af := calcularAfinidad(posPublicador, us.pos)
			postFeed := feedUsuario{post: publicado, idPost: app.posts.Cantidad() - 1, afinidad: af}
			us.usuario.RecibirPost(postFeed)
		}
		return true
	})
	return fmt.Sprintln("Post publicado"), nil
}

func (app *algogram) VerSiguienteFeed() (string, error) {
	if app.loggeado == nil {
		return "", errors.New("Usuario no loggeado o no hay mas posts para ver")
	}
	feedUsuario, err := app.loggeado.VerSiguienteFeed()
	if err != nil {
		return "", err
	}
	post := feedUsuario.post
	id := feedUsuario.idPost
	mensaje := fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d\n", id, post.Autor(), post.VerPost(), post.CantidadLikes())
	return mensaje, nil
}

func (app *algogram) LikearPost(id int) (string, error) {
	if app.loggeado == nil || !app.posts.Pertenece(id) {
		return "", errors.New("Error: Usuario no loggeado o Post inexistente")
	}
	post := app.posts.Obtener(id)
	app.loggeado.LikearPost(post)
	return fmt.Sprintln("Post likeado"), nil
}

func (app *algogram) MostrarLikes(id int) (string, error) {
	if !app.posts.Pertenece(id) {
		return "", errors.New("Error: Post inexistente o sin likes")
	}
	post := app.posts.Obtener(id)
	cant, usuarios := post.MostrarLikes()
	if cant == 0 {
		return "", errors.New("Error: Post inexistente o sin likes")
	}
	mensaje := fmt.Sprintf("El post tiene %d likes:\n", post.CantidadLikes())
	for _, u := range usuarios {
		mensaje += fmt.Sprintf("\t%s\n", u)
	}
	return mensaje, nil
}
