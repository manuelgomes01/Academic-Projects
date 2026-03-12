#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "reporte.h"

#define MAX_DESCRIPCION 500
#define MAX_COMANDO 100
#define MAX_DIAS 7
#define MAX_CANTIDAD_COMANDOS 6
#define MAX_TERMINAL 10

#define ARCHIVO_VALIDO "csv"
#define FORMATO_ARCHIVO "%[^;];%d:%d;%c;%i\n"
#define FORMATO_INCIDENTE_AGREGADO " %[^;];%d:%d;%c;%i"
#define FORMATO_ASIGNAR_INCIDENTE "%s;%02d:%02d;%c;%i\n"
#define MODO_LECTURA "r"
#define MODO_ESCRITURA "w"

const char DIAS_VALIDOS[MAX_DIAS] = {'L', 'M', 'X', 'J', 'V', 'S', 'D'};
const char DIGITO_MIN = '0';
const char DIGITO_MAX = '9';

const char* COMANDOS_VALIDOS[MAX_CANTIDAD_COMANDOS] = {"asignar_incidente", "agregar_incidente", "imprimir_incidentes", "agregar_multiples_incidentes", "dia_mas_accidentado", "salir"};
const char* NOMBRES_DIAS_VALIDOS[MAX_DIAS] = {"LUNES", "MARTES", "MIERCOLES", "JUEVES", "VIERNES", "SABADO", "DOMINGO"};

const int ERROR = -1;
const int ERROR_MEMORIA = -2;
const int ARGUMENTOS_MINIMOS = 2;
const int ARGUMENTOS_ARCHIVO_SALIDA = 3;
const int POSICION_ARCHIVO_ENTRADA = 1;
const int POSICION_ARCHIVO_SALIDA = 2;
const int LINEA_VALIDA = 5;
const int HORAS_MINUTOS = 60;
const int HORA_MAX = 24;
const int HORA_MIN = 0;
const int MINUTOS_MIN = 0;
const int MINUTOS_MAX = 60;
const int MINUTOS_TOTALES = 1440;
const int PRIORIDAD_MIN = 100;
const int PRIORIDAD_MAX = 1;
const int ENTERO_VALIDO = 1;
const int VACIO = 0;
const int POSICION_LUNES = 0;
const int POSICION_MARTES = 1;
const int POSICION_MIERCOLES = 2;
const int POSICION_JUEVES = 3;
const int POSICION_VIERNES = 4;
const int POSICION_SABADO = 5;
const int POSICION_DOMINGO = 6;
const int POS_COMANDO_ASIGNAR_INCIDENTE = 0;
const int POS_COMANDO_AGREGAR_INCIDENTE = 1;
const int POS_COMANDO_IMPRIMIR_INCIDENTES = 2; 
const int POS_COMANDO_AGREGAR_MULTIPLES_INCIDENTES = 3;
const int POS_COMANDO_DIA_MAS_ACCIODENTADO = 4;
const int POS_COMANDO_SALIR = 5;

typedef struct incidente_local {
    char descripcion[MAX_DESCRIPCION];
    int horas;
    int minutos;
    char dia;
    int prioridad;
} incidente_local_t;

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
 * Post: * Elimina los incidentes del reporte
         * Librea la memoria de la descripcion los incidentes del reporte que devuelve 'eliminar_incidente'
 */
void liberar_descripciones(reporte_t* reporte) {
    incidente_t incidente;
    while(eliminar_incidente(reporte, &incidente)) {
        free(incidente.descripcion);
    }
}

/*
 * Pre: - 
 * Post: Devuelve true cuando la prioridad del incidente esta entre PRIORIDAD_MIN y PRIORIDAD_MAX, sino devuelve false
 */
bool es_prioridad_valida(int prioridad) {
    return prioridad >= PRIORIDAD_MAX && prioridad <= PRIORIDAD_MIN;
}

/*
 * Pre: - 
 * Post: Devuelve true cuando el dia del incidente esta dentro del vector DIAS_VALIDOS, sino devuelve false
 */
bool es_dia_valido(char dia) {
    bool es_dia_valido = false;
    int i = 0;
    while(!es_dia_valido && i < MAX_DIAS) {
        if(dia == DIAS_VALIDOS[i]) {
            es_dia_valido = true;
        }
        i++;
    }
    return es_dia_valido;
}

/*
 * Pre: - 
 * Post: Devuelve true cuando la hora del incidente esta entre HORA_MIN y HORA_MAX y cuando los minutos del incidente estan entre
 *       MINUTOS_MIN y MINUTOS_MAX. Si no se cumples las dos en simultaneo, devuelve false
 */
bool son_minutos_validos(int hora, int minutos) {
    return ((hora >= HORA_MIN && hora < HORA_MAX) && (minutos >= MINUTOS_MIN && minutos < MINUTOS_MAX));
}

/*
 * Pre: - 
 * Post: * Devuelve true cuando el incidente cumple las condiciones de 'son_minutos_validos', 'es_dia_valido' y 'es_prioridad_valida'
 *       * En caso de que el incidente no cumpla con una de las tres, devuelve false
 */
bool incidente_valido(incidente_local_t incidente) {
    return son_minutos_validos(incidente.horas, incidente.minutos) && es_dia_valido(incidente.dia) && es_prioridad_valida(incidente.prioridad);
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * archivo_salida se debe abrir con MODO_ESCRITURA
 * Post: * Elimina los incidentes del reporte y los guarda en archivo_salida
         * Librea la memoria de la descripción los incidentes del reporte que devuelve 'eliminar_incidente'
         * Cuando 'eliminar_incidente' devuelve false, cierra incidentes_f
         * En caso de que no se pueda abrir incidentes_f, devuelve ERROR. Devuelve 0 si todo sale bien
 */
int guardar_reporte_actualizado(reporte_t* reporte, char* archivo_salida) {
    FILE* incidentes_f = fopen(archivo_salida, MODO_ESCRITURA);
    if(!incidentes_f) {
        printf("Hubo un error a la hora de abrir el archivo \n");
        return ERROR;
    }
    incidente_t incidente;
    while(eliminar_incidente(reporte, &incidente)){
        int horas = incidente.minutos_desde_medianoche / HORAS_MINUTOS;
        int minutos = incidente.minutos_desde_medianoche % HORAS_MINUTOS;
        fprintf(incidentes_f, "%s;%02d:%02d;%c;%i\n", incidente.descripcion, horas, minutos, incidente.dia, incidente.prioridad);
        free(incidente.descripcion);
    }
    fclose(incidentes_f);
    
    return 0;
}

/*
 * Pre: * El reporte_origen debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * El reporte_destino debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
 * Post: * Mueve los incidentes del reporte_origen a reporte_destino
         * Librea la memoria de la descripcion los incidentes del reporte_origen que devuelve eliminar_incidente
 */
void mover_incidentes(reporte_t* reporte_origen, reporte_t* reporte_destino) {
    incidente_t incidente;
    char* memoria_reservada;
    while(eliminar_incidente(reporte_origen, &incidente)) {
        memoria_reservada = incidente.descripcion;
        agregar_incidente(reporte_destino, incidente);
        free(memoria_reservada);
    }
}

/*
 * Pre: * La posicion del dia debe estar definida como constande y debe coincidir 
        con la posicion que el dia tiene dentro del vector NOMBRE_DIAS_VALIDOS
 * Post: * Imprime por pantalla el nombre del dia en que ocurrieron mas incidentes
 */
void mostrar_dia_mas_accidentado(int posicion_dia) {
    printf("\nEl día más peligroso es el *%s*.\n", NOMBRES_DIAS_VALIDOS[posicion_dia]);
}

/*
 * Pre: * El vector accidentes_dias debe estar correctamente inicializado
 * Post: * Devuelve la posicion del dia que mas incidentes tiene. En caso de que dos dias tengan la misma cantidad, 
 *         devuelve el primero con mas cantidad, empezando desde el lunes
 */
int buscar_dia_mas_accidentado(int accidentes_dias[MAX_DIAS]) {
    int incidentes_max = POSICION_LUNES;
    for(int i = 1; i < MAX_DIAS; i++) {
        if(accidentes_dias[i] > accidentes_dias[incidentes_max]) {
            incidentes_max = i;
        }
    }
    return incidentes_max;
}

/*
 * Pre: * Las posiciones hasta MAX_DIAS del vector accidentes_dias deben estar previamente inicializadas en VACIO
        * Debe recibir un incidente con un día válido
 * Post: * Suma 1 a la posicion del vector accidentes_dias correspondiente al día del incidente 
 */
void dia_del_incidente(incidente_t incidente, int accidentes_dias[MAX_DIAS]) {
    bool encontre_el_dia = false;
    int i = 0;
    while(!encontre_el_dia && i < MAX_DIAS) {
        if(incidente.dia == DIAS_VALIDOS[i]) {
            accidentes_dias[i]++;
            encontre_el_dia = true;
        }
        i++;
    }
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
 * Post: * Calcula cual es el NOMBRE_DIA_VALIDO en el que ocurrieron más incidentes y lo imprime por pantalla
         * Devuelve 0 si todo sale bien, o ERROR_MEMORIA si no se pudo crear el reporte_aux
 */
int calculo_accidentes_por_dia(reporte_t* reporte) {
    reporte_t* reporte_aux = reporte_crear();
    if(!reporte_aux) {
        printf("Hubo un error al crear el reporte auxilar para dia_mas_accidentado");
        return ERROR;
    }
    int accidentes_dias[MAX_DIAS];
    for(int i = 0; i < MAX_DIAS; i++){
        accidentes_dias[i] = VACIO;
    }
    incidente_t incidente;
    bool tiene_incidentes = false;
    char* memoria_reservada;
    while(eliminar_incidente(reporte, &incidente)) {
        tiene_incidentes = true;
        memoria_reservada = incidente.descripcion;
        dia_del_incidente(incidente, accidentes_dias);
        agregar_incidente(reporte_aux, incidente);
        free(memoria_reservada);
    }
    if(!tiene_incidentes) {
        printf("No hay incidentes en el reporte. No se puede determinar el día más accidentado.\n");
        reporte_destruir(reporte_aux);
        return 0;
    }
    mover_incidentes(reporte_aux, reporte);
    int posicion_dia_mas_accidentado = buscar_dia_mas_accidentado(accidentes_dias);
    mostrar_dia_mas_accidentado(posicion_dia_mas_accidentado);
    reporte_destruir(reporte_aux);

    return 0;
}

/*
 * Pre: - 
 * Post: Imprime por pantalla los campos del incidente
 */
void mostrar_incidente(incidente_t incidente, int hora, int minutos) {
    printf(" \033[1m*\033[0m DESCRIPCION: %s\n", incidente.descripcion);
    printf(" \033[1m*\033[0m PRIORIDAD: %i\n", incidente.prioridad);
    printf(" \033[1m*\033[0m HORARIO: %02d:%02d\n", hora, minutos);
    printf(" \033[1m*\033[0m DIA: %c\n", incidente.dia);
    printf("==============================\n");
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * Los incidentes dentro del reporte deben estar ordenados por prioridad (de PRIORIDAD_MAX a PRIORIDAD_MIN)
 * Post: * Imprime todos los incidentes cargados dentro del reporte por orden de prioridad
         * Devuelve 0 si todo sale bien o ERROR_MEMORIA si no se pudo crear el reporte_aux
 */
int imprimir_incidentes(reporte_t* reporte) {
    reporte_t* reporte_aux = reporte_crear();
    if(!reporte_aux) {
        printf("Hubo un error al crear el reporte auxiliar para imprimir_incidentes.\n");
        return ERROR_MEMORIA;
    }
    incidente_t incidente;
    bool tiene_incidentes = false;
    char* memoria_reservada;
    int i = 0;
    while(eliminar_incidente(reporte, &incidente)) {
        tiene_incidentes = true;
        int horas = incidente.minutos_desde_medianoche / HORAS_MINUTOS;
        int minutos = incidente.minutos_desde_medianoche % HORAS_MINUTOS;
        printf("\n\033[1mINCIDENTE NUMERO %i\033[0m\n", ++i);
        mostrar_incidente(incidente, horas, minutos);
        memoria_reservada = incidente.descripcion;
        agregar_incidente(reporte_aux, incidente);
        free(memoria_reservada);
    }
    if(!tiene_incidentes) {
        printf("No hay incidentes en el reporte.\n");
        reporte_destruir(reporte_aux);
        return 0;
    }
    mover_incidentes(reporte_aux, reporte);
    reporte_destruir(reporte_aux);

    return 0;
}

/*
 * Pre: - 
 * Post: * Devuelve true cuando prioridad_anterior es más chica que prioridad_actual. 
 *       * En caso de que las prioridades sean iguales, devuelve true cuando minutos_anteriores son más chicos que minutos_actuales
 *       * Si ambos son iguales (prioridades y minutos), tabien devuelve true
 *       * Si no se cumple una, devuelve false
 */
bool esta_ordenado(int prioridad_anterior, int prioridad_actual, int minutos_anteriores, int minutos_actuales) {
    return prioridad_anterior < prioridad_actual || (prioridad_anterior == prioridad_actual && minutos_anteriores <= minutos_actuales);
}

/*
 * Pre: - 
 * Post: * Devuelve true cuando cargue_el_incidente es false y esta_ordenado devuelve true
         * En cualquier otro caso, devuelve false
 */
bool es_posicion_ordenada(bool cargue_el_incidente, incidente_t incidente, incidente_t incidente_cargado) {
    return !cargue_el_incidente && esta_ordenado(incidente.prioridad, incidente_cargado.prioridad, incidente.minutos_desde_medianoche, incidente_cargado.minutos_desde_medianoche);
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * El incidente_t debe estar previamnete inicializado con todos sus campos validos
 * Post: * Agrega el incidente ingresado de forma ordenada (respoetando el orden) dentro de reporte_aux
         * Los incidentes ordenados dentro de reporte_aux los mueve al reporte
         * Destruye el reporte_aux con reporte_destruir
         * En caso reporte_crear devuelva NULL al crear reporte_aux, libera la memoria del incidente ingresado y devuelve ERROR
           Si se pudieron ordenar bien los incidentes, devuelve 0
 */
int ordenar_incidentes(reporte_t* reporte, incidente_t incidente) {
    reporte_t* reporte_aux = reporte_crear();
    if(!reporte_aux) {
        printf("Hubo un error al crear el reporte auxiliar para ordenar_incidentes.\n");
        free(incidente.descripcion);
        return ERROR_MEMORIA;
    }
    incidente_t incidente_cargado;
    bool cargue_el_incidente = false;
    char* memoria_reservada;
    while(eliminar_incidente(reporte, &incidente_cargado)) {
        if(es_posicion_ordenada(cargue_el_incidente, incidente, incidente_cargado)) {
            memoria_reservada = incidente.descripcion;
            agregar_incidente(reporte_aux, incidente);
            free(memoria_reservada);
            cargue_el_incidente = true;
        }
        memoria_reservada = incidente_cargado.descripcion;
        agregar_incidente(reporte_aux, incidente_cargado);
        free(memoria_reservada);
    }
    
    if(!cargue_el_incidente) {
        memoria_reservada = incidente.descripcion;
        agregar_incidente(reporte_aux, incidente);
        free(memoria_reservada);
    }

    mover_incidentes(reporte_aux, reporte);
    reporte_destruir(reporte_aux);

    return 0;
}

/*
 * Pre: - 
 * Post: Devuelve la cantidad de minutos totales que tiene la hora en la que ocurrio en incidente
 */
int minutos_totales(int hora, int minutos){
    return (hora * HORAS_MINUTOS) + minutos;
}

/*
 * Pre: * El incidente_t debe estar previamente inicializado
        * El incidente_local_t debe estar previamnete inicializado con todos sus campos validos
 * Post: * Copia los campos del incidente_local_t en incidente_t
         * Si no se pudo reservar memoria dinamica para la descripcion, devuelve ERROR. Devuelve 0 si se pudo inicializar bien el incidente
 */
int inicializar_incidente(incidente_t* incidente, incidente_local_t incidente_local) {
    incidente->descripcion = malloc(sizeof(char) * MAX_DESCRIPCION + 1);
    if(!incidente->descripcion) {
        printf("Hubo un error al reservar memoria para la descripción.\n");
        return ERROR_MEMORIA;
    }
    strcpy(incidente->descripcion, incidente_local.descripcion);
    incidente->minutos_desde_medianoche = minutos_totales(incidente_local.horas, incidente_local.minutos);
    incidente->dia = incidente_local.dia;
    incidente->prioridad = incidente_local.prioridad;

    return 0;
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * incidente_local_t fue previamente inicializado
 * Post: * Verifica que se haya ingresado un incidente valido y, en caso de serlo, lo carga dentro del reporte
           manteniedo el orden de prioridad
         * Devuelve 0 si es un incidente válido y se pudo cargar bien dentro del reporte, ERROR_MEMORIA si no se pudo 
           ordenar el reporte o no se pudo inicializar el incidente, o ERROR si es un incidente invalido
 */
int agregar_incidente_ingresado(reporte_t* reporte, incidente_local_t incidente_local) {
    if(incidente_valido(incidente_local)){
        incidente_t incidente_aux;
        int inicializado = inicializar_incidente(&incidente_aux, incidente_local);
        if(inicializado == ERROR_MEMORIA) {
            return ERROR_MEMORIA;
        }
        if (ordenar_incidentes(reporte, incidente_aux) == ERROR_MEMORIA) {
            return ERROR_MEMORIA;
        }
        printf("\033[1;32mEl incidente se agrego exitosamente.\033[0m\n");
        return 0;
    }
    printf("\033[1;31mEl incidente ingresado es inválido.\033[0m\n");
    return ERROR;
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * cantidad_incidentes debe ser un número válido
 * Post: * Valida el incidente que se ingresa y lo carga dentro del reporte. De no ser válido, vuelve a pedir que se ingrese
         * Repite esto cantidad_incidentes veces o hasta que se_cargo_incidente sea ERROR_MEMORIA 
 */
void agregar_multiples_incidentes(reporte_t* reporte, int cantidad_incidentes) {
    incidente_local_t incidente;
    int se_cargo_incidente = 0;
    int i = 0;
    while(i < cantidad_incidentes && se_cargo_incidente != ERROR_MEMORIA) {
        printf("Ingrese: \033[1m<\033[0mdescripcion\033[1m>\033[0m\033[1m;\033[0m<hora>\033[1m;\033[0m<dia>\033[1m;\033[0m<prioridad>\033[1m>\033[0m:\n");
        scanf(FORMATO_INCIDENTE_AGREGADO, incidente.descripcion, &incidente.horas, &incidente.minutos, &incidente.dia, &incidente.prioridad);
        se_cargo_incidente = agregar_incidente_ingresado(reporte, incidente);
        while(se_cargo_incidente == ERROR) {
            printf("Vuelva a ingresar: ");
            scanf(FORMATO_INCIDENTE_AGREGADO, incidente.descripcion, &incidente.horas, &incidente.minutos, &incidente.dia, &incidente.prioridad);
            se_cargo_incidente = agregar_incidente_ingresado(reporte, incidente);
        }
        i++;
    }
    
    if(se_cargo_incidente == ERROR_MEMORIA) {
        printf("Hubo un error al reservar memoria.\n");
    } 
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * Los incidentes dentro del reporte deben estar ordenados por prioridad (de PRIORIDAD_MAX a PRIORIDAD_MIN)
 * Post: * Muestra por pantalla el incidente con mayor prioridad y lo elimina del reporte
 */
void asignar_incidente(reporte_t* reporte){
    incidente_t incidente;
    if(eliminar_incidente(reporte, &incidente)) {
        int horas = incidente.minutos_desde_medianoche / HORAS_MINUTOS;
        int minutos = incidente.minutos_desde_medianoche % HORAS_MINUTOS;
        printf("\n\033[1mINCIDENTE ASIGNADO:\033[0m\n");
        mostrar_incidente(incidente, horas, minutos); 
        free(incidente.descripcion);
    } else {
        printf("\033[1mNo hay incidentes en el reporte.\033[0m\n");
    }
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * comando_ingresado debe estar dentro del vector COMANDOS_VALIDOS y ser distinto a COMANDOS_VALIDOS[POS_COMANDO_SALIR]
 * Post: * Ejecuta la función correspondiente al comando ingresado
         * Para COMANDOS_VALIDOS[POS_COMANDO_AGREGAR_MULTIPLES_INCIDENTES] verifica que se haya ingresado un número válido
 */
void ejecutar_comando(reporte_t* reporte, char comando_ingresado[MAX_COMANDO]) {
    incidente_local_t incidente;
    int cantidad_incidentes = 0;

    if(strcmp(comando_ingresado, COMANDOS_VALIDOS[POS_COMANDO_ASIGNAR_INCIDENTE]) == 0) {
        asignar_incidente(reporte);
    } else if(strcmp(comando_ingresado, COMANDOS_VALIDOS[POS_COMANDO_AGREGAR_INCIDENTE]) == 0) {
        scanf(FORMATO_INCIDENTE_AGREGADO, incidente.descripcion, &incidente.horas, &incidente.minutos, &incidente.dia, &incidente.prioridad);
        agregar_incidente_ingresado(reporte, incidente);
    } else if(strcmp(comando_ingresado, COMANDOS_VALIDOS[POS_COMANDO_IMPRIMIR_INCIDENTES]) == 0) {
        imprimir_incidentes(reporte);
    } else if(strcmp(comando_ingresado, COMANDOS_VALIDOS[POS_COMANDO_AGREGAR_MULTIPLES_INCIDENTES]) == 0) {
        int entero_valido = scanf("%i", &cantidad_incidentes);
        if(entero_valido != ENTERO_VALIDO || cantidad_incidentes < 0) {
            printf("Cantidad ingresada inválida\n");
        } else {
            agregar_multiples_incidentes(reporte, cantidad_incidentes);
        }
    } else if(strcmp(comando_ingresado, COMANDOS_VALIDOS[POS_COMANDO_DIA_MAS_ACCIODENTADO]) == 0) {
        calculo_accidentes_por_dia(reporte);
    }
}

/*
 * Pre: - 
 * Post: Devuelve true si comando ingresado está dentro del vector COMANDOS_VALIDOS, sino devuelve false
 */
bool es_comando_valido(char comando[MAX_COMANDO]) {
    bool estado_comando = false; 
    int i = 0;
    while(!estado_comando && i < MAX_CANTIDAD_COMANDOS){
        if(strcmp(comando, COMANDOS_VALIDOS[i]) == 0) {
            estado_comando = true;
        }
        i++;
    }
    return estado_comando;
}

/*
 * Pre: - 
 * Post: Valida el comando que se ingresa y lo guarda en comando_ingresado
 */
void pedir_comando(char comando_ingresado[MAX_COMANDO]){
    printf("Ingrese un comando: ");
    scanf("%s", comando_ingresado);
    while(!es_comando_valido(comando_ingresado)) {
        printf("Comando inválido. Vuelva a ingresar: ");
        scanf("%s", comando_ingresado);
    }
}

/*
 * Pre: - 
 * Post: Imprime el mensaje de error correspondiente a porqué no se pudo cargar el archivo ingresado al reporte
 */
void mostrar_error_procesar_archivo(bool archivo_ordenado, bool es_incidente_valido, bool reservo_memoria) {
    if(!archivo_ordenado) {
        printf("El archivo ingresado no está ordenado.\n");
    } else if(!es_incidente_valido){
        printf("Hay un incidente inválido dentro del archivo ingresado.\n");
    } else if(!reservo_memoria) {
        printf("Hubo un problema al reservar memoria heap.\n");
    }
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * El incidente_local debe ser un incidente con todos sus campos válidos
        * es_incidente_valido debe estar previamente inicializado
 * Post: * Agrega el incidente al reporte y libera la memoria dinámica de la descripción que se reservó en 'inicializar_incidente'
         * Si 'inicializar_incidente' devuelve ERROR_MEMORIA, cambia el estado de es_incidente_valido a false 
 */
void cargar_incidente(reporte_t* reporte, incidente_local_t incidente_local, bool* reservo_memoria) {
    incidente_t incidente;
    int inicializado = inicializar_incidente(&incidente, incidente_local);
    if(inicializado == ERROR_MEMORIA) {
        printf("Hubo un error al reservar memoria para la descripción.\n");
        *reservo_memoria = false;
    }
    char* memoria_reservada = incidente.descripcion;
    agregar_incidente(reporte, incidente);
    free(memoria_reservada);
}

/*
 * Pre: * El reporte debe estar previamente creado con 'reporte_crear' y ser distinto de NULL
        * El archivo_ingresado se debe abrir en MODO_LECTURA
 * Post: * Carga en el reporte los incidentes dentro del archivo_ingresado
         * Si no se pudo abrir el archivo_ingresado, si el archivo_ingresado no esta ordenado, el archivo_ingresado tiene un incidente invalido, 
           o si no se pudo reservar memoria para uno de los incidentes, cierra el archivo_ingresado, 
           libera la memoria de las descripciones (si se cargó algún incidente) y devuelve ERROR
 */
int procesar_archivo(reporte_t* reporte, char* archivo_ingresado) {
    FILE* incidentes_f = fopen(archivo_ingresado, MODO_LECTURA);
    if(!incidentes_f) {
        printf("Hubo un error a la hora de abrir el archivo.\n");
        return ERROR;
    }
    incidente_local_t incidente_local;
    
    int incidente_leido = fscanf(incidentes_f, FORMATO_ARCHIVO, incidente_local.descripcion, &incidente_local.horas, &incidente_local.minutos, &incidente_local.dia, &incidente_local.prioridad);

    int prioridad_anterior = VACIO;
    int minutos_totales_anterior = VACIO;
    int minutos_totales_actuales = minutos_totales(incidente_local.horas, incidente_local.minutos);
    bool archivo_ordenado = true;
    bool es_incidente_valido = true;
    bool reservo_memoria = true; 

    while (incidente_leido != EOF && incidente_leido == LINEA_VALIDA && archivo_ordenado && es_incidente_valido && reservo_memoria) {
        es_incidente_valido = incidente_valido(incidente_local);
        cargar_incidente(reporte, incidente_local, &reservo_memoria);
            
        prioridad_anterior = incidente_local.prioridad;
        minutos_totales_anterior = minutos_totales_actuales;
        
        incidente_leido = fscanf(incidentes_f, FORMATO_ARCHIVO, incidente_local.descripcion, &incidente_local.horas, &incidente_local.minutos, &incidente_local.dia, &incidente_local.prioridad);
        
        archivo_ordenado = esta_ordenado(prioridad_anterior, incidente_local.prioridad, minutos_totales_anterior, minutos_totales_actuales);
        minutos_totales_actuales = minutos_totales(incidente_local.horas, incidente_local.minutos);
    }
    if(!archivo_ordenado || !es_incidente_valido || !reservo_memoria) {
        mostrar_error_procesar_archivo(archivo_ordenado, es_incidente_valido, reservo_memoria);
        fclose(incidentes_f);
        liberar_descripciones(reporte);
        return ERROR;
    }
        
    fclose(incidentes_f);
    return 0;
}

/*
 * Pre: * La cantidad de argumentos ingresados por línea de comando deben ser mayor a ARGUMENTOS_MINIMOS
 * Post: * Devuelve cual va a ser el archivo_salida, dependiendo de si se ingresaron ARGUMENTOS_MINIMOS o más
 */
char* definir_archivo_salida(int argc, char* argv[]) {
    if(argc > ARGUMENTOS_MINIMOS) {
        return argv[POSICION_ARCHIVO_SALIDA];
    } 

    return argv[POSICION_ARCHIVO_ENTRADA];
}

/*
 * Pre: - 
 * Post: Devuelve true cuando se ingresan por linea de comando, como mínimo ARGUMENTOS_MINIMOS
 */
bool es_cantidad_argumentos_validos(int argumentos) {
    return argumentos >= ARGUMENTOS_MINIMOS;
}

/*
 * Pre: - 
 * Post: * Verifica que los argumentos ingresados por linea de comando sean válidos
         * Si los argumentos no cumples con 'es_cantidad_argumentos_validos', devuelve ERROR
 */
int validar_linea_comando(int argc, char* argv[]) {
    if (!es_cantidad_argumentos_validos(argc)) {
        printf("Cantidad de argumentos inválida.\n");
        return ERROR;
    }

    return 0;
}

int main(int argc, char* argv[]) {
    int estado_linea_comando = validar_linea_comando(argc, argv);
    if(estado_linea_comando == ERROR) {
        return ERROR;
    }

    char* archivo_ingresado = argv[POSICION_ARCHIVO_ENTRADA];
    char* archivo_salida = definir_archivo_salida(argc, argv);
    
    reporte_t* reporte = reporte_crear();
    if(!reporte) {
        printf("Hubo un error al crear el reporte.\n");
        return ERROR;
    }
    
    if(procesar_archivo(reporte, archivo_ingresado) == ERROR) {
        liberar_descripciones(reporte);
        reporte_destruir(reporte);
        return ERROR;
    }

    char comando_ingresado[MAX_COMANDO] = "";
    while(strcmp(comando_ingresado, COMANDOS_VALIDOS[POS_COMANDO_SALIR]) != 0) {
        pedir_comando(comando_ingresado);
        ejecutar_comando(reporte, comando_ingresado);
    }

    int reporte_actualizado = guardar_reporte_actualizado(reporte, archivo_salida);
    if(reporte_actualizado == ERROR) {
        liberar_descripciones(reporte);
        reporte_destruir(reporte);
        printf("Hubo un error al actualizar el archivo de incidentes.\n");
        return ERROR;
    }

    printf("\033[1;32mEl reporte se guardó exitosamente.\033[0m\n");
    reporte_destruir(reporte);
    
    return 0;
}
