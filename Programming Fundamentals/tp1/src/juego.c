#include "plutonio.h"
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

const char ARRIBA = 'W';
const char ABAJO = 'S';
const char DERECHA = 'D';
const char IZQUIERDA = 'A';
const char LINTERNA = 'L';

const int JUEGO_GANADO = 1;
const int JUEGO_PERDIDIO = -1;
const int JUEGO_EN_CURSO = 0;

/*
 * Pre: -
 * Post: Devuelve true siempre que el usuario ingrese uno de los 5 caracteres válidos. De no cumplirlo, devuelve false. 
 */
bool es_accion_valida(char accion) {
    return (accion == ARRIBA) || (accion == ABAJO) || (accion == DERECHA) || (accion == IZQUIERDA) || (accion == LINTERNA); 
}

/*
 * Pre: -
 * Post: Verifica que la acción ingresada por el usuario sea válida (W,S,A,D,L). De no cumplirlo, repite la pregunta hasta que sea válido. 
 * Si ya hay una linterna activa, no deja que el usuario vuelva a ingresar una L. 
 * Una vez válida, guarda el caracter ingresado en accion. 
 */
void preguntar_movimiento(char* accion) {
    printf("Ingrese un movimiento [W, S, A, D] o active una linterna [L] \n");
    system("stty raw");
    *accion = (char)getchar();
    system("stty cooked");  
    while(!es_accion_valida(*accion)) {
        printf("Accion Invalida. Vuelva a ingresar. \n");
        system("stty raw");
        *accion = (char)getchar();
        system("stty cooked");
    }
}

/*
 * Pre: El juego debe estar en estado perdido
 * Post: Imprime por pantalla el mensaje cuando se pierde el juego 
 */
void mensaje_perdio_juego() {
    system("clear");
        printf("\033[1;31m *'no estoy cansado, nada cansado...'\033[0m\n");
        printf("\033[1;31m * PERDISTE!!\033[0m \033[31mHomero se quedo sin energía.\033[0m\n");
}

/*
 * Pre: El juego debe estar en estado ganado
 * Post: Imprime por pantalla el mensaje cuando se gana el juego 
 */
void mensaje_gano_juego() {
    system("clear");
    printf("\033[1;32m * 'SONO, SONO, SONO!! ME LLAMAN DEL BAR DE MOE!!EN ESE BUEN LUGAR!! ME GUSTA BEBER ALCOHOL!!' \033[0m \033[32m\n");
    printf("\033[1;32m * GANASTEE!! \033[0m \033[32mHOMERO LOGRO JUNTAR TODAS LAS BARRAS!!\033[0m \n");
}

/*
 * Pre: -
 * Post: Imprime por pantalla el detalle del juego y pide que se inicie 
 */
void mensaje_introduccion() {
    printf("\033[1m* DESASTRE NUCLEAR *\033[0m\n");
    printf(" > Después de un largo día en la planta nuclear, Homero sale apurado rumbo al bar de Moe.\n");
    printf(" > Pero distraído y con la mente puesta en su cerveza, tropieza con una caja repleta de barras de plutonio, desparramándolas por todos lados.\n");
    printf(" > Ayudalo a juntarlas todas antes de que se quede sin energía y decida que lo mejor es rendirse e irse a tomar una cerveza.\n");
    printf("\n");
    printf("\033[1m* COMANDOS:\033[0m\n");
    printf("\t W: ARRIBA \n");
    printf("\t S: ABAJO \n");
    printf("\t A: IZQUIERDA \n");
    printf("\t D: DERECHA \n");
    printf("\t L: ACTIVAR LA LINTERNA \n");
    printf("\n");
    printf("\033[1mPARA INICIAR PRESIONE UNA TECLA CUALQUIERA\033[0m\n");
    printf("Cual es 'Cualquiera'? Veo la 'ESC', 'Ctrl' y 'PgUp'. NO TIENE TECLA 'CUALQUIERA' \n");

    system("stty raw");
    getchar();
    system("stty cooked");
}

int main () {
    srand ((unsigned)time(NULL));

    mensaje_introduccion() ;
    system("clear");

    juego_t juego;
    inicializar_juego(&juego);
    mostrar_juego(juego);

    char accion = ' ';

    while(estado_juego(juego) == JUEGO_EN_CURSO) { 
        preguntar_movimiento(&accion);
        realizar_jugada(&juego, accion);
        system("clear");
        mostrar_juego(juego);
    }

    if (estado_juego(juego) == JUEGO_GANADO) {
        mensaje_gano_juego();
    } else if (estado_juego(juego) == JUEGO_PERDIDIO) {
        mensaje_perdio_juego();
    }

    return 0;
}