#include <stdio.h>
#include <stdbool.h>

#define MESES_ANUALES 12
#define PUNTAJE_SI_TRABAJO_MATERIALES_RADIOACTIVOS 10
#define PUNTAJE_NO_TRABAJO_MATERIALES_RADIOACTIVOS 0
#define PUNTAJE_ACTIVAR_PROTOCOLO_SEGURIDAD 15
#define PUNTAJE_CORRER 10
#define PUNTAJE_SALTAR_PISO_NUEVE 5
#define PUNTAJE_SEGUIR_DURMIENDO 0
#define PUNTOS_MANTENIMIENTO 15
#define PUNTOS_OPERATIVO 25
#define PUNTOS_TECNICO 35

const int EXPERIENCIA_MINIMA = 0;
const int EXPERIENCIA_MAXIMA = 300;
const char ACTIVAR_PROTOCOLO_SEGURIDAD = 'P';
const char CORRER = 'C'; 
const char SALTAR_PISO_NUEVE = 'S';
const char SEGUIR_DURMIENDO = 'D';
const char SI_TRABAJO_MATERIALES_RADIOACTIVOS = 'S';
const char NO_TRABAJO_MATERIALES_RADIOACTIVOS = 'N';
const float CANTIDAD_MINIMA_DONAS = 0;
const float CANTIDAD_MAXIMA_DONAS = 12;

//Pre: Que el usuario haya ingresado correctamente las 4 respuestas. 
//Pos: Dependiendo del valor de la sumatoria del puntaje de cada respuesta, devuelve que puesto esta capacitado para ocupar. 
void puesto_trabajo_correspondiente(int resultado_puntaje_total) {
    if (resultado_puntaje_total <= PUNTOS_MANTENIMIENTO){
        printf("Con las respuestas brindadas, tu posición es: -MANTENIMIENTO-\n");
    } else if ((PUNTOS_MANTENIMIENTO < resultado_puntaje_total) && (resultado_puntaje_total <= PUNTOS_OPERATIVO)){
        printf("Con las respuestas brindadas, tu posición es: -OPERATIVO-\n");
    } else if ((PUNTOS_OPERATIVO < resultado_puntaje_total) && (resultado_puntaje_total <= PUNTOS_TECNICO)) {
        printf("Con las respuestas brindadas, tu posición es: -TECNICO-\n");
    } else {
        printf("Con las respuestas brindadas, tu posición es: -INGENIERO-\n");
    }
}

//Pre: Que cada valor ingresado por el usuario cumpla con los requisitos de cada pregunta. 
//Pos: Transforma en un int el calculo de los valores correspondientes a las respuestas ingresadas. 
int calcular_puntaje_total(int experiencia_final, int puntaje_final_pregunta_experiencia, int puntaje_final_trabajo_materiales_radioactivos, float cuantas_donas_come) {
    float calculo_puntaje_total = ((float) experiencia_final) + ((float)puntaje_final_pregunta_experiencia) + ((float)puntaje_final_trabajo_materiales_radioactivos) - cuantas_donas_come;
    return (int)calculo_puntaje_total;
}

//Pre: -- 
//Pos: Devuelve true siempre que el float ingresado sea mayor o igual a CANTIDAD_MINIMA_DONAS y menor o igual a CANTIDAD_MAXIMA_DONAS. 
bool cantidad_donas_validas(float cantidad_donas_que_come) {
    return (cantidad_donas_que_come >= CANTIDAD_MINIMA_DONAS) && (cantidad_donas_que_come <= CANTIDAD_MAXIMA_DONAS);
}

//Pre: Verifica que el float ingresado cumpla los requisitos del bool cantidad_donas_validas.
//Pos: Cuando lo cumple, guarda el valor ingresado en la variable. 
void preguntar_cantidad_donas_jornada (float* cantidad_donas_que_come) {
    printf("¿Cuántas donas podrías comer en una jornada laboral?\n");
    scanf(" %f", cantidad_donas_que_come);
    while(!cantidad_donas_validas(*cantidad_donas_que_come)) {
        printf("La cantidad tiene que ser entre 0 y 12 donas.\n");
        scanf(" %f", cantidad_donas_que_come);
    }
}

//Pre: Dependiendo del caracter ingresado, le asigna un 10 si es igual a SI_TRABAJO_MATERIALES_RADIOACTIVOS o un 0 si es distinto. 
//Pos: Guarda el puntaje correspondiente a SI_TRABAJO_MATERIALES_RADIOACTIVOS o a NO_TRABAJO_MATERIALES_RADIOACTIVOS en la variable. 

int puntaje_trabajo_materiales_radioactivos(char trabajo_materiales_radioactivos, int puntaje_fianl_trabajo_materiales_radioactivos) {
    if (trabajo_materiales_radioactivos == true){
        puntaje_fianl_trabajo_materiales_radioactivos = PUNTAJE_SI_TRABAJO_MATERIALES_RADIOACTIVOS;

    } else {
        puntaje_fianl_trabajo_materiales_radioactivos = PUNTAJE_NO_TRABAJO_MATERIALES_RADIOACTIVOS;
    }
    return puntaje_fianl_trabajo_materiales_radioactivos;
}

//Pre: El caracter ingresado tiene que ser S o N.
//Pos: Le asigna al caracter un valor true si es igual a S o false si es disntinto de S (o sea cuando es N).

bool valido_trabajo_materiales_radioactivos(char trabajo_materiales_radioactivos){
    if (trabajo_materiales_radioactivos == SI_TRABAJO_MATERIALES_RADIOACTIVOS){
        return true;
    } else {
        return false;
    }
}

//Pre: -- 
//Pos: Devuelve true siempre que el caracter ingresado sea distinto a las constantes declaradas. 

bool es_respuesta_valida(char trabajaste_materiales_radioactivos) {
    return (trabajaste_materiales_radioactivos != SI_TRABAJO_MATERIALES_RADIOACTIVOS) && (trabajaste_materiales_radioactivos != NO_TRABAJO_MATERIALES_RADIOACTIVOS);
}

//Pre: Verifica que el caracter ingresado sea S o N. De no cumplirlo, repite el pedido hasta que coincida con uno de los dos.
//Pos: Una vez que cumpla con lo necesario, guarda ese caracter en la variable. 

void preguntar_trabajo_materiales_radioactivos (char* trabajaste_materiales_radioactivos) {
    printf("¿Ha trabajado previamente con materiales radiactivos? [S/N] \n");
    scanf(" %c", trabajaste_materiales_radioactivos);
    while (es_respuesta_valida(*trabajaste_materiales_radioactivos)) {
        printf("Tu respuesta tiene que ser [S/N] \n");
        scanf(" %c", trabajaste_materiales_radioactivos);
    }
}

//Pre: Según al caracter correspondiente a la reaccion ingresada, busca que valor asignarle.
//Pos: Devuelve el valor correspondiente al caracter ingresado. 
int puntaje_pregunta_emergencia(char como_reacciona_en_emergencias, int puntaje_final_pregunta_emergencias) {
    if (como_reacciona_en_emergencias == ACTIVAR_PROTOCOLO_SEGURIDAD){
        puntaje_final_pregunta_emergencias = PUNTAJE_ACTIVAR_PROTOCOLO_SEGURIDAD;
    } else if (como_reacciona_en_emergencias == CORRER) {
        puntaje_final_pregunta_emergencias = PUNTAJE_CORRER;
    } else if (como_reacciona_en_emergencias == SALTAR_PISO_NUEVE) {
        puntaje_final_pregunta_emergencias = PUNTAJE_SALTAR_PISO_NUEVE;
    } else {
        puntaje_final_pregunta_emergencias = PUNTAJE_SEGUIR_DURMIENDO;
    }
    return puntaje_final_pregunta_emergencias;
}

//Pre: --
//Pos: Devuelve true si el caracter ingresado es distinto al que se le asignó a las constantes ACTIVAR_PROTOCOLO_SEGURIDAD, CORRER, SALTAR_PISO_NUEVE y SEGUIR_DURMIENDO. 
bool es_reaccion_valida(char reaccion_ingresada) {
    return (reaccion_ingresada != ACTIVAR_PROTOCOLO_SEGURIDAD) && (reaccion_ingresada != CORRER) && (reaccion_ingresada != SALTAR_PISO_NUEVE) && (reaccion_ingresada != SEGUIR_DURMIENDO);
}

//Pre: Verifica que el caracter ingresado este dentro de las 4 posibilidades.
//Pos: Si cumple con lo pedido, lo guarda en la variable reaccion_ingresada. 
void preguntar_reaccion_emergencia (char* reaccion_ingresada) {
    printf("En caso de emergencia, ¿cómo actuaría primero? Ingrese P, C, S o D.\n");
    scanf(" %c", reaccion_ingresada);
    while (es_reaccion_valida(*reaccion_ingresada)) {
        printf("Tiene que ser P, C, S o D.\n");
        scanf(" %c", reaccion_ingresada);
    }
}

//Pre: --
//Pos: Devuelve la cantidad de años truncada. 
int calcular_experiencia(int experiencia_laboral) {
    return experiencia_laboral / MESES_ANUALES;
}

//Pre: --
//Pos: Devuelve true si la experiencia tiene un valor entre 0 y 300 meses. De no cumplirlo, devuelve false.
bool es_experiencia_valida(int experiencia_ingresada) {
    return (experiencia_ingresada >= EXPERIENCIA_MINIMA) && (experiencia_ingresada <= EXPERIENCIA_MAXIMA);
}

//Pre: Verificar que el número ingresado cumpla los requisitos del bool es_experiencia_valida. Si no lo hace, repite el pedido hasta que lo cumpla.
//Pos: Una vez que cumple los requisitos, guarda el valor en la variable.
void preguntar_experiencia(int* experiencia_ingresada) {
    printf("¿Cuántos meses tenes de experiencia en el sector nuclear? \n");
    scanf(" %i", experiencia_ingresada);
    while (!es_experiencia_valida(*experiencia_ingresada)) {
        printf("Tiene que ser entre 0 y 300 meses.\n");
        scanf(" %i", experiencia_ingresada);
    }
}

int main() {
    //PREGUNTA 1
    int experiencia_laboral = 0;
    preguntar_experiencia(&experiencia_laboral); 
    calcular_experiencia(experiencia_laboral);
    int experiencia_final = calcular_experiencia(experiencia_laboral);
    //PREGUNTA 2
    char como_reacciona_en_emergencias;
    int puntaje_final_pregunta_emergencias = 0; 
    preguntar_reaccion_emergencia(&como_reacciona_en_emergencias);
    puntaje_pregunta_emergencia(como_reacciona_en_emergencias, puntaje_final_pregunta_emergencias); 
    puntaje_final_pregunta_emergencias = puntaje_pregunta_emergencia(como_reacciona_en_emergencias, puntaje_final_pregunta_emergencias); 
    //PREGUNTA 3
    char trabajo_materiales_radioactivos = 0;
    int puntaje_final_trabajo_materiales_radioactivos = 0; 
    preguntar_trabajo_materiales_radioactivos(&trabajo_materiales_radioactivos);
    puntaje_trabajo_materiales_radioactivos(trabajo_materiales_radioactivos, puntaje_final_trabajo_materiales_radioactivos);
    valido_trabajo_materiales_radioactivos(trabajo_materiales_radioactivos);
    bool trabajaste_materiales_radioactivos = valido_trabajo_materiales_radioactivos(trabajo_materiales_radioactivos);
    puntaje_final_trabajo_materiales_radioactivos = puntaje_trabajo_materiales_radioactivos(trabajaste_materiales_radioactivos, puntaje_final_trabajo_materiales_radioactivos);
    //PREGUNTA 4
    float cuantas_donas_come = 0;
    preguntar_cantidad_donas_jornada(&cuantas_donas_come);
    //ASIGNACION DE PUESTO
    calcular_puntaje_total(experiencia_final, puntaje_final_pregunta_emergencias, puntaje_final_trabajo_materiales_radioactivos, cuantas_donas_come);
    int resultado_puntaje_total =(calcular_puntaje_total(experiencia_final, puntaje_final_pregunta_emergencias, puntaje_final_trabajo_materiales_radioactivos, cuantas_donas_come));
    //lo que devuelve calcular_puntaje_total lo estoy guardando en calculo_puntaje_total y se lo asigno a resultado_puntaje_total
    puesto_trabajo_correspondiente(resultado_puntaje_total);
    
    return 0;
}