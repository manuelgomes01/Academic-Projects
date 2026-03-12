#include "plutonio.h"
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <time.h>

#define MAX_HERRAMIENTAS 30
#define MAX_OBSTACULOS 30
#define MAX_BARRAS 10
#define MAX_PROPULSORES 10

#define MAX_FIL 20
#define MAX_COL 20

const int CANTIDAD_BARRAS = 10;
const int CANTIDAD_DONAS = 5;
const int CANTIDAD_INTERRUPTORES = 4;
const int CANTIDAD_RATAS = 5;
const int CANTIDAD_BARRILES = 15;
const int CANTIDAD_ESCALERAS = 3;
const int CANTIDAD_LINTERNAS_INICIAL = 5;
const int CANTIDAD_CHARCOS = 3;
const int CANTIDAD_BARRAS_INICIAL = 0;

const int ENERGIA_INICIAL = 400;
const int MOVIMIENTO = 1;
const int MOV_LINTERNA_INICIAL = 0;
const int MAX_MOV_LIINTERNAS = 5;
const int MOV_PROPULSORES = 3;
const int DIST_MANHATTAN = 2;

const int PUNTOS_RATAS = 1;
const int PUNTOS_BARRIL = 15;
const int PUNTOS_DONAS = 10;

const int GANO_JUEGO = 1;
const int PERDIO_JUEGO = -1;
const int SIGUE_JUGANDO = 0;

const int BARRAS_TOTALES = 10;
const int SIN_ENERGIA = 0;

const char ACCION_ARRIBA = 'W';
const char ACCION_ABAJO = 'S';
const char ACCION_DERECHA = 'D';
const char ACCION_IZQUIERDA = 'A';
const char ACCION_LINTERNA = 'L';

const char HOMERO = 'H';
const char BARRAS = 'B';
const char DONAS = 'D';
const char INTERRUPTORES = 'I';
const char RATAS = 'R';
const char BARRILES = 'A';
const char ESCALERA = 'E';
const char CHARCOS = 'C';
const char VACIO = '-';

int estado_juego(juego_t juego) {
    if(juego.homero.cantidad_barras == BARRAS_TOTALES) {
        return GANO_JUEGO; 
    } else if(juego.homero.energia <= SIN_ENERGIA) {
        return PERDIO_JUEGO;
    }
    return SIGUE_JUGANDO; 
}

/*
 * Pre: * El juego debe estar correctamente inicializado
        * La posicion de Homero debe ser válida
        * Los objetos deben estar corectamente inicializados y llenos hasta su tope
 * Post: Devuelve el tipo de objeto que hay en esa posicion O VACIO cunado no hay ninguno
 */
char objeto_pisado(juego_t juego) {
    char objeto = VACIO;
    for(int barra = 0; barra < juego.cantidad_barras; barra++) {
        if((juego.barras[barra].posicion.fil == juego.homero.posicion.fil) && (juego.barras[barra].posicion.col == juego.homero.posicion.col)) {
            objeto = juego.barras[barra].tipo; 
        }
    }
    for(int herramienta = 0; herramienta < juego.cantidad_herramientas; herramienta++) {
        if((juego.herramientas[herramienta].posicion.fil == juego.homero.posicion.fil) && (juego.herramientas[herramienta].posicion.col == juego.homero.posicion.col)) {
            objeto = juego.herramientas[herramienta].tipo; 
        }
    }
    for(int obstaculo = 0; obstaculo < juego.cantidad_obstaculos; obstaculo++) {
        if((juego.obstaculos[obstaculo].posicion.fil == juego.homero.posicion.fil) && (juego.obstaculos[obstaculo].posicion.col == juego.homero.posicion.col)) {
            objeto = juego.obstaculos[obstaculo].tipo; 
        }
    }
    for(int propulsor = 0; propulsor < juego.cantidad_propulsores; propulsor++) {
        if((juego.propulsores[propulsor].posicion.fil == juego.homero.posicion.fil) && (juego.propulsores[propulsor].posicion.col == juego.homero.posicion.col)) {
            objeto = juego.propulsores[propulsor].tipo; 
        }
    }
    return objeto;
}

/*
 * Pre: * El juego debe estar correctamente inicializado 
        * La posicion de Homero debe ser válida
        * Los objetos deben estar corectamente inicializados y llenos hasta su tope
 * Post: Mustra un mensaje cuando la posicion de Homero coincida con la de un interruptor o con uno de los obstaculos 
 */
void mostrar_mensaje_objeto(juego_t juego) {
    char objeto = objeto_pisado(juego); 
    if (objeto == INTERRUPTORES) {
        printf("Que no panda el cúnico!! Activaste un interruptor \n"); 
    } else if(objeto == RATAS) {
        printf("De donde salio esa RATA? \n");
    } else if(objeto == BARRILES) {
        printf("D'oh!! Chocaste con un barril \n"); 
    } 
}

/*
 * Pre: * El juego debe estar correctamente inicializado 
        * Los datos de Homero deben estar correctamente inicializados y bien actualizados a lo largo del juego 
 * Post: Imprime por pantalla los datos de Homero
 */
void mostrar_datos(juego_t juego) {
    if (juego.homero.energia > 100) {
        printf("\033[1;32m * ENERGIA:\033[0m \033[32m%i\033[0m \n", juego.homero.energia);
    } else {
        printf("\033[1;31m * ENERGIA:\033[0m \033[31m%i\033[0m\n", juego.homero.energia);
    }
    printf("\033[1m * BARRAS ENCONTRADAS:\033[0m %i \n", juego.homero.cantidad_barras); 
    if (juego.homero.cantidad_linternas != 0) {
        printf("\033[1m * LINTERNAS:\033[0m %i \n", juego.homero.cantidad_linternas);
    } else {
        printf("\033[1m * TE QUEDASTE SIN LINTERNAS\033[0m \n");
    }
    if(juego.homero.linterna_activada) {
        printf("Hiciste %i movimientos con la linterna \n", juego.homero.mov_linterna);
    }
}

/*
 * Pre: * El juego debe estar correctamente inicializado
        * El tablero tiene que estar previamente inicializado 
        * Los objetos deben estar bien inicializados, llenos hasta su tope y correctamente asignados en el tablero
        * La posicion de Homero debe ser válida
 * Post:* Imprime el objeto (si es visible) que hay en esa posicion del tablero
        * Si no es visible, imprime un espacio   
 */
void mostrar_objeto(objeto_t *objeto, coordenada_t coordenada, int cantidad) {
    int posicion = 0;
    bool encontre_objeto = false;
    while ((posicion < cantidad) && (!encontre_objeto)) {
        if (objeto[posicion].visible && objeto[posicion].posicion.fil == coordenada.fil && objeto[posicion].posicion.col == coordenada.col) {
            if(objeto[posicion].tipo == BARRAS) {
                printf("\033[92m%c\033[0m", objeto[posicion].tipo);
            } else if(objeto[posicion].tipo == DONAS) {
                printf("\033[35m%c\033[0m", objeto[posicion].tipo);
            } else if(objeto[posicion].tipo == INTERRUPTORES) {
                printf("\033[1m%c\033[0m", objeto[posicion].tipo);
            } else if(objeto[posicion].tipo == RATAS) {
                printf("\033[31m%c\033[0m", objeto[posicion].tipo);
            } else if(objeto[posicion].tipo == BARRILES) {
                printf("\033[31m%c\033[0m", objeto[posicion].tipo);
            } else if(objeto[posicion].tipo == ESCALERA) {
                printf("\033[36m%c\033[0m", objeto[posicion].tipo);
            } else {
                printf("\033[36m%c\033[0m", objeto[posicion].tipo);
            }
            encontre_objeto = true;
        }
        posicion++;
    } if (!encontre_objeto) {
        printf(" ");
    } 
}

/*
 * Pre: * El juego El juego debe estar correctamente inicializado
        * El tablero tiene que estar previamente inicializado y lleno hasta su tope
        * Homero debe estar bien inicializado 
        * Los objetos esten bien inicializados y llenos hasta su tope
 * Post: Asigna todos los objetos y a Homero al tablero en sus correspondientes posiciones
 */
void asignar_objetos_tablero(char tablero[MAX_FIL][MAX_COL], juego_t juego) {
    for(int barra = 0; barra < (juego.cantidad_barras); barra++) {
        tablero[juego.barras[barra].posicion.fil][juego.barras[barra].posicion.col] = juego.barras[barra].tipo;
    } 
    for(int herramienta = 0; herramienta < (juego.cantidad_herramientas); herramienta++) {
        tablero[juego.herramientas[herramienta].posicion.fil][juego.herramientas[herramienta].posicion.col] = juego.herramientas[herramienta].tipo;
    } 
    for(int obstaculo = 0; obstaculo < (juego.cantidad_obstaculos); obstaculo++) {
        tablero[juego.obstaculos[obstaculo].posicion.fil][juego.obstaculos[obstaculo].posicion.col] = juego.obstaculos[obstaculo].tipo; 
    }
    for(int propulsor = 0; propulsor < (juego.cantidad_propulsores); propulsor++) {
        tablero[juego.propulsores[propulsor].posicion.fil][juego.propulsores[propulsor].posicion.col] = juego.propulsores[propulsor].tipo; 
    } 
    tablero[juego.homero.posicion.fil][juego.homero.posicion.col] = HOMERO;
}

/*
 * Pre: -
 * Post: Inicializa la matriz en VACIO 
 */
void inicializar_tablero(char tablero[MAX_FIL][MAX_COL]) {
    for(int fila =0; fila < MAX_FIL; fila++){
        for(int columna = 0; columna < MAX_COL; columna++) {
            tablero[fila][columna] = VACIO;
        }
    }     
}

void mostrar_juego(juego_t juego) {
    char tablero[MAX_FIL][MAX_COL];
    inicializar_tablero(tablero);
    asignar_objetos_tablero(tablero, juego);
    mostrar_mensaje_objeto(juego);
    for(int fila = 0; fila < MAX_FIL; fila++) {
        for(int columna = 0; columna < MAX_COL; columna++) {
            coordenada_t coordenada_tablero;
            coordenada_tablero.fil = fila;
            coordenada_tablero.col = columna;
            printf("|");
            if(tablero[fila][columna] == HOMERO) {
                printf("\033[1;93m%c\033[0m", tablero[fila][columna]);
            } else if(tablero[fila][columna] == BARRAS) {
                mostrar_objeto(juego.barras, coordenada_tablero, juego.cantidad_barras);
            } else if((tablero[fila][columna] == DONAS) || (tablero[fila][columna] == INTERRUPTORES)) {
                mostrar_objeto(juego.herramientas, coordenada_tablero, juego.cantidad_herramientas);
            } else if((tablero[fila][columna] == RATAS) || (tablero[fila][columna] == BARRILES)) {
                mostrar_objeto(juego.obstaculos, coordenada_tablero, juego.cantidad_obstaculos);
            } else if((tablero[fila][columna] == ESCALERA) || (tablero[fila][columna] == CHARCOS)) {
                mostrar_objeto(juego.propulsores, coordenada_tablero, juego.cantidad_propulsores);
            } else {
                printf(" ");
            }   
        }
        printf("|\n");
    }
    mostrar_datos(juego);
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La posicion de Homero debe ser válida
        * Los objetos deben estar correctamente inicializados y llenos hasta su tope
        * El objeto que se tiene que borrar tiene que estar dentro del vector
 * Post: Borra esa dona o barra de su vector correspondiente 
 */
void borrar_objeto(objeto_t *objeto, int* cantidad_objeto, coordenada_t coordenada_homero, char tipo) {
    int posicion = 0;
    bool objeto_borrado = false; 
    while(!objeto_borrado) {
        if(objeto[posicion].posicion.fil == coordenada_homero.fil && objeto[posicion].posicion.col == coordenada_homero.col && objeto[posicion].tipo == tipo) {
            objeto[posicion] = objeto[*cantidad_objeto - 1];
            (*cantidad_objeto)--;
            objeto_borrado = true; 
        } else {
            posicion++;
        }  
    }
}

/*
 * Pre: * El juego debe estar correctamente inicializado
        * Homero debe estar correctamente inicializado 
        * La coordenada que recibe debe estar dentro de los límites del juego
 * Post: Devuelve true cuando una posicion esta vacia (no tiene nungun objeto ni a Homero asignado) o false cuando ya este ocupada
 */
bool es_posicion_vacia(juego_t juego, coordenada_t coordenada) {
    bool posicion_vacia = true; 
    if((juego.homero.posicion.fil == coordenada.fil) && (juego.homero.posicion.col == coordenada.col)) {
        posicion_vacia = false; 
    }
    for(int barra = 0; barra < juego.cantidad_barras; barra++) {
        if((juego.barras[barra].posicion.fil == coordenada.fil) && (juego.barras[barra].posicion.col == coordenada.col)) {
            posicion_vacia = false; 
        }
    }
    for(int herramienta = 0; herramienta < juego.cantidad_herramientas; herramienta++) {
        if((juego.herramientas[herramienta].posicion.fil == coordenada.fil) && (juego.herramientas[herramienta].posicion.col == coordenada.col)) {
            posicion_vacia = false; 
        }
    }
    for(int obstaculo = 0; obstaculo < juego.cantidad_obstaculos; obstaculo++) {
        if((juego.obstaculos[obstaculo].posicion.fil == coordenada.fil) && (juego.obstaculos[obstaculo].posicion.col == coordenada.col)) {
            posicion_vacia = false; 
        }
    }
    for(int propulsor = 0; propulsor < juego.cantidad_propulsores; propulsor++) {
        if((juego.propulsores[propulsor].posicion.fil == coordenada.fil) && (juego.propulsores[propulsor].posicion.col == coordenada.col)) {
            posicion_vacia = false; 
        }
    }
    return posicion_vacia; 
}

/*
 * Pre: Que los topes de los objetos sean válidos y esten bien inicializados
 * Post: Devuelve una coordenada válida y vacía
 */
coordenada_t generar_coordenada(juego_t juego) {
    coordenada_t coordenada_aux;
    coordenada_aux.fil = rand() % MAX_FIL;
    coordenada_aux.col = rand() % MAX_COL;
    while(!es_posicion_vacia(juego, coordenada_aux)) {
        coordenada_aux.fil = rand() % MAX_FIL;
        coordenada_aux.col = rand() % MAX_COL;
    }
    return coordenada_aux;
}

/*
 * Pre: * El juego debe estar correctamente inicializado
        * La ratas deben estar bien inicializadas 
 * Post: Le va a asignar una nueva posicion random a las ratas cuando el usuario se pare en un interruptor o active una linterna
 */
void mover_ratas(juego_t *juego) {
    for(int rata = 0; rata < CANTIDAD_RATAS; rata++) {
        juego->obstaculos[rata].posicion = generar_coordenada(*juego);
    }
}

/*
 * Pre: * La posicion de Homero debe ser válida
        * Los objetos deben estar correctamente inicializados y llenos hasta su tope
 * Post: Hace visibles a los objetos que estan a distancia manhattan menor o igual a 2 de la posicion de Homero
 */
void mostrar_distancia_homero(objeto_t *objeto, int cantidad, coordenada_t coordenada_homero) {
    for(int posicion = 0; posicion < cantidad; posicion++) {
        int distancia_fil = abs(objeto[posicion].posicion.fil - coordenada_homero.fil);
        int distancia_col = abs(objeto[posicion].posicion.col - coordenada_homero.col);
        if((distancia_fil + distancia_col) <= DIST_MANHATTAN) {
            objeto[posicion].visible = true;
        }
    }
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La posicion de Homero debe ser válida
        * Los objetos deben estar correctamente inicializados y llenos hasta su tope
 * Post: Llama a mostrar_distancia_homero para ver los objetos a distancia manhattan menor o igual a 2 de Homero
 */
void activar_linterna(juego_t *juego) {
    mostrar_distancia_homero(juego->barras, juego->cantidad_barras, juego->homero.posicion);
    mostrar_distancia_homero(juego->herramientas, juego->cantidad_herramientas, juego->homero.posicion);
    mostrar_distancia_homero(juego->obstaculos, juego->cantidad_obstaculos, juego->homero.posicion);
    mostrar_distancia_homero(juego->propulsores, juego->cantidad_propulsores, juego->homero.posicion);
}

/*
 * Pre: -
 * Post: Devuelve true y mueve a las ratas de lugar cuando la posicion de homero coincide con la de un interruptor. Sino devuelve false
 */
bool es_interruptor(juego_t *juego, char tipo) {
    if(tipo == INTERRUPTORES) {
        mover_ratas(juego); 
        return true; 
    } 
    return false; 
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La posicion de Homero debe ser válida
        * Los objetos deben estar correctamente inicializados y llenos hasta su tope
 * Post: Hace la interaccion de Homero con ese objeto
 */
void interactuar_con_objeto(juego_t *juego, char objeto) {
    bool activo_interruptor = es_interruptor(juego, objeto);
    for(int barra = 0; barra < (juego->cantidad_barras); barra++) {
        juego->barras[barra].visible = activo_interruptor;
    }
    for(int herramienta = 0; herramienta < (juego->cantidad_herramientas); herramienta++) {
        if(juego->herramientas[herramienta].tipo == DONAS) {
            juego->herramientas[herramienta].visible = activo_interruptor;
        }
    }
    for(int obstaculo = 0; obstaculo < (juego->cantidad_obstaculos); obstaculo++) {
        juego->obstaculos[obstaculo].visible = activo_interruptor;
    }
    for(int propulsor = 0; propulsor < (juego->cantidad_propulsores); propulsor++) {
        juego->propulsores[propulsor].visible = activo_interruptor;
    }
    if(objeto == BARRAS) {
        juego->homero.cantidad_barras++;
        borrar_objeto(juego->barras, &juego->cantidad_barras, juego->homero.posicion, BARRAS);
    } else if(objeto == DONAS) {
        (juego->homero).energia += PUNTOS_DONAS;
        borrar_objeto(juego->herramientas, &juego->cantidad_herramientas, juego->homero.posicion, DONAS);
    } else if(objeto == RATAS) {
        if(juego->homero.cantidad_linternas != 0) {
            (juego->homero).cantidad_linternas -= PUNTOS_RATAS;
        }
    } else if(objeto == BARRILES) {
        (juego->homero).energia -= PUNTOS_BARRIL;
    }
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La posicion de Homero debe ser válida
        * Los objetos deben estar correctamente inicializados y llenos hasta su tope
 * Post: Hace que homero interactue con los objetos cuando lo mueve un propulsor y cambia los objetos que puede ver con la linterna
 */
void chequear_interaccion_nueva_posicion(juego_t *juego) {
    char objeto_nuevo = objeto_pisado(*juego);
    interactuar_con_objeto(juego, objeto_nuevo);
    if(juego->homero.linterna_activada) {
        activar_linterna(juego);
    }
}

/*
 * Pre: -
 * Post: Devuelve true siempre que la posicion que reciba este dentro de los limites del tablero. Sino devuelve false
 */
bool dentro_tablero(coordenada_t coordenada) {
    return (coordenada.fil >= 0 ) && (coordenada.fil < MAX_FIL) && (coordenada.col< MAX_COL) && (coordenada.col >=0);
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La posicion de Homero debe ser válida 
        * La acción ingresada debe ser válida
 * Post:* Mueve a Homero 3 veces en direccion a la accion que ingreso el usuario y en sentido de las escaleras
        * Se chequean las interacciones con objetos en cada nueva posicion valida 
 */
void activar_escalera(juego_t *juego, char accion) {
    coordenada_t nueva_coordenada_homero = juego->homero.posicion;
    int cantidad_diagonales = 0;
    if(accion == ACCION_DERECHA) {
        nueva_coordenada_homero.fil = nueva_coordenada_homero.fil - MOVIMIENTO;
        nueva_coordenada_homero.col = nueva_coordenada_homero.col + MOVIMIENTO;
        while((dentro_tablero(nueva_coordenada_homero)) && (cantidad_diagonales < MOV_PROPULSORES)) {
            juego->homero.posicion = nueva_coordenada_homero;
            nueva_coordenada_homero.fil--;
            nueva_coordenada_homero.col++;
            chequear_interaccion_nueva_posicion(juego);
            cantidad_diagonales++;
        }
    } else if(accion == ACCION_IZQUIERDA) {
        nueva_coordenada_homero.fil = nueva_coordenada_homero.fil - MOVIMIENTO;
        nueva_coordenada_homero.col = nueva_coordenada_homero.col - MOVIMIENTO;
        while((dentro_tablero(nueva_coordenada_homero)) && (cantidad_diagonales < MOV_PROPULSORES)) {
            juego->homero.posicion = nueva_coordenada_homero;
            nueva_coordenada_homero.fil--;
            nueva_coordenada_homero.col--;
            chequear_interaccion_nueva_posicion(juego);
            cantidad_diagonales++;
        }
    } else {
        nueva_coordenada_homero.fil = nueva_coordenada_homero.fil - MOVIMIENTO;
        while((dentro_tablero(nueva_coordenada_homero)) && (cantidad_diagonales < MOV_PROPULSORES)) {
            juego->homero.posicion = nueva_coordenada_homero;
            nueva_coordenada_homero.fil--;
            chequear_interaccion_nueva_posicion(juego);
            cantidad_diagonales++;
        }
    }
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La posicion de Homero debe ser válida 
        * La acción ingresada debe ser válida
 * Post:* Mueve a Homero 3 veces en direccion a la accion que ingreso el usuario y en sentido de los charcos
        * Se chequean las interacciones con objetos en cada nueva posicion valida   
 */
void activar_charco(juego_t *juego, char accion) {
    coordenada_t nueva_coordenada_homero = juego->homero.posicion;
    int cantidad_diagonales = 0;
    if(accion == ACCION_DERECHA) {
        nueva_coordenada_homero.fil = nueva_coordenada_homero.fil + MOVIMIENTO;
        nueva_coordenada_homero.col = nueva_coordenada_homero.col + MOVIMIENTO;
        while((dentro_tablero(nueva_coordenada_homero)) && (cantidad_diagonales < MOV_PROPULSORES)) {
            juego->homero.posicion = nueva_coordenada_homero;
            nueva_coordenada_homero.fil++;
            nueva_coordenada_homero.col++;
            chequear_interaccion_nueva_posicion(juego);
            cantidad_diagonales++;
        }
    } else if(accion == ACCION_IZQUIERDA) {
        nueva_coordenada_homero.fil = nueva_coordenada_homero.fil + MOVIMIENTO;
        nueva_coordenada_homero.col = nueva_coordenada_homero.col - MOVIMIENTO;
        while((dentro_tablero(nueva_coordenada_homero)) && (cantidad_diagonales < MOV_PROPULSORES)) {
            juego->homero.posicion = nueva_coordenada_homero;
            nueva_coordenada_homero.fil++;
            nueva_coordenada_homero.col--;
            chequear_interaccion_nueva_posicion(juego);
            cantidad_diagonales++;
        }
    } else {
        nueva_coordenada_homero.fil = nueva_coordenada_homero.fil + MOVIMIENTO;
        while((dentro_tablero(nueva_coordenada_homero)) && (cantidad_diagonales < MOV_PROPULSORES)) {
            juego->homero.posicion = nueva_coordenada_homero;
            nueva_coordenada_homero.fil++;
            chequear_interaccion_nueva_posicion(juego);
            cantidad_diagonales++;
        }
    }
}

/*
 * Pre: * El juego debe estar inicializado correctamente
 * Post: Llama a la funcion que corresponde segun corresponda a la accion ingresada
 */
void activar_propulsor(juego_t *juego, char objeto, char accion) {
    if (objeto == ESCALERA) {
        activar_escalera(juego, accion);
    } else if(objeto == CHARCOS) {
        activar_charco(juego, accion);
    }
}

/*
 * Pre: * El juego debe estar inicializado correctamente
        * La acción que se recibe debe ser válida
 * Post:* Actualiza la posicion de Homero en caso de que este dentro del tablero
        * Sino deja a Homero en la misma posicion hasta que si se pueda mover
 */
void realizar_accion(juego_t *juego, coordenada_t coordenada_homero, char accion) {
    if(dentro_tablero(coordenada_homero)) {
        juego->homero.posicion = coordenada_homero;
        char objeto = objeto_pisado(*juego);
        if(accion != ACCION_LINTERNA) {
            (juego->homero).energia -= MOVIMIENTO;
            interactuar_con_objeto(juego, objeto);
            activar_propulsor(juego, objeto, accion);
            if(juego->homero.linterna_activada) {
                juego->homero.mov_linterna++;
                if(!(juego->homero.mov_linterna > MAX_MOV_LIINTERNAS)) {
                    activar_linterna(juego);
                } else {
                    juego->homero.linterna_activada = false;
                    juego->homero.mov_linterna = MOV_LINTERNA_INICIAL;
                } 
            }
        } 
    }
    coordenada_homero = juego->homero.posicion;  
}

void realizar_jugada(juego_t *juego, char accion) {
    coordenada_t coordenada_homero = juego->homero.posicion;
    if(accion == ACCION_DERECHA) {
        coordenada_homero.col += MOVIMIENTO;
    } else if(accion == ACCION_IZQUIERDA) {
        coordenada_homero.col -= MOVIMIENTO;
    } else if(accion == ACCION_ARRIBA) {
        coordenada_homero.fil -= MOVIMIENTO;
    } else if(accion == ACCION_ABAJO) {
        coordenada_homero.fil += MOVIMIENTO;
    } else {
        if((juego->homero.cantidad_linternas != 0) && (!juego->homero.linterna_activada)) { 
            juego->homero.cantidad_linternas -=MOVIMIENTO; 
            juego->homero.linterna_activada = true; 
            juego->homero.mov_linterna += 1; 
            mover_ratas(juego);
            activar_linterna(juego);
        }
    }
    realizar_accion(juego, coordenada_homero, accion); 
}

/*
 * Pre: * Que Homero ya este inicializado correctamente
        * Los vectores de objetos deben estar definidos
 * Post: Inicializa los vectores de los objetos hasta sus respectivas cantidades (tope)
 */
void inicializar_objetos(juego_t *juego, objeto_t *objeto, int *indice, char tipo, int cantidad) {
    while(*indice < cantidad) {
        bool posicion_ocupada = false;
        while(!posicion_ocupada) {
            objeto[*indice].posicion = generar_coordenada(*juego);
            objeto[*indice].tipo = tipo; 
            if(tipo == INTERRUPTORES) {
                objeto[*indice].visible = true;
            } else {
                objeto[*indice].visible = false;
            }
            posicion_ocupada = true; 
        
        }
        (*indice)++;
    }
}

/*
 * Pre: -
 * Post: Inicializa personaje_t homero
 */
void inicializar_homero(juego_t *juego) {
    juego->homero.posicion = generar_coordenada(*juego);
    juego->homero.cantidad_linternas = CANTIDAD_LINTERNAS_INICIAL;
    juego->homero.cantidad_barras = CANTIDAD_BARRAS_INICIAL;
    juego->homero.energia = ENERGIA_INICIAL;
    juego->homero.linterna_activada = false;
    juego->homero.mov_linterna = MOV_LINTERNA_INICIAL;
}

void inicializar_juego(juego_t *juego) {
    juego->cantidad_herramientas = 0;
    juego->cantidad_obstaculos = 0;
    juego->cantidad_barras = 0;
    juego->cantidad_propulsores = 0;
 
    inicializar_homero(juego);

    inicializar_objetos(juego, juego->barras, &juego->cantidad_barras, BARRAS, CANTIDAD_BARRAS);

    inicializar_objetos(juego, juego->herramientas, &juego->cantidad_herramientas, DONAS, CANTIDAD_DONAS);
    inicializar_objetos(juego, juego->herramientas, &juego->cantidad_herramientas, INTERRUPTORES, (CANTIDAD_INTERRUPTORES + CANTIDAD_DONAS));

    inicializar_objetos(juego, juego->obstaculos, &juego->cantidad_obstaculos, RATAS, CANTIDAD_RATAS);
    inicializar_objetos(juego, juego->obstaculos, &juego->cantidad_obstaculos, BARRILES, (CANTIDAD_BARRILES + CANTIDAD_RATAS));

    inicializar_objetos(juego, juego->propulsores, &juego->cantidad_propulsores, ESCALERA, CANTIDAD_ESCALERAS);
    inicializar_objetos(juego, juego->propulsores, &juego->cantidad_propulsores, CHARCOS, (CANTIDAD_CHARCOS + CANTIDAD_ESCALERAS));
}