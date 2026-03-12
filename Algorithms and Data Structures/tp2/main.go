package tp2

import (
	"bufio"
	"errors"
	a "estructuras/algogram"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	_LOGIN          string = "login"
	_LOGOUT         string = "logout"
	_PUBLICAR       string = "publicar"
	_VER_FEED       string = "ver_siguiente_feed"
	_LIKEAR         string = "likear_post"
	_MOSTRAR_LIKES  string = "mostrar_likes"
	_MAX_PARAMETROS int    = 2
)

func procesarArchivoUsuarios(rutaArchivo string) ([]string, error) {
	arch, err := os.Open(rutaArchivo)
	if err != nil {
		return nil, errors.New("Error al abrir el archivo")
	}
	s := bufio.NewScanner(arch)
	var usuarios []string
	for s.Scan() {
		nombre := s.Text()
		usuarios = append(usuarios, nombre)
	}
	return usuarios, nil
}

func EjecutarComando(comando []string, algogram a.Algogram) string {
	var err error
	var msj string
	switch comando[0] {
	case _LOGIN:
		msj, err = algogram.Login(comando[1])
	case _LOGOUT:
		msj, err = algogram.Logout()
	case _PUBLICAR:
		msj, err = algogram.Publicar(comando[1])
	case _VER_FEED:
		msj, err = algogram.VerSiguienteFeed()
	case _LIKEAR:
		id, _ := strconv.Atoi(comando[1])
		msj, err = algogram.LikearPost(id)
	case _MOSTRAR_LIKES:
		id, _ := strconv.Atoi(comando[1])
		msj, err = algogram.MostrarLikes(id)
	}

	if err != nil {
		return fmt.Sprintln(err)
	}
	return msj
}

func main() {
	usuarios, err := procesarArchivoUsuarios(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	algogram, err := a.CrearAlgogramConUsuarios(usuarios)
	if err != nil {
		panic("Error al abrir el archivo")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		comando := strings.SplitN(s.Text(), " ", _MAX_PARAMETROS)
		msj := EjecutarComando(comando, algogram)
		fmt.Print(msj)
	}
}
