package post

import (
	dicc "tdas/diccionario"
)

type post struct {
	autor string
	texto string
	likes dicc.DiccionarioOrdenado[string, bool]
}

func cmp(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}

func CrearPost(autor string, texto string) Post {
	return &post{autor: autor, texto: texto, likes: dicc.CrearABB[string, bool](cmp)}
}

func (p *post) Autor() string {
	return p.autor
}

func (p *post) VerPost() string {
	return p.texto
}

func (p *post) Likear(usuario string) {
	if p.likes.Pertenece(usuario) {
		return
	}
	p.likes.Guardar(usuario, true)
}

func (p *post) CantidadLikes() int {
	return p.likes.Cantidad()
}

func (p *post) MostrarLikes() (int, []string) {
	var usuarios []string
	p.likes.Iterar(func(u string, _ bool) bool {
		usuarios = append(usuarios, u)
		return true
	})
	return p.CantidadLikes(), usuarios
}
