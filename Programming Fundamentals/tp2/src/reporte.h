#ifndef __REPORTE_H__
#define __REPORTE_H__

#include <stdbool.h>

typedef struct incidente {
    char *descripcion; // Un puntero a un string creado en el heap
    int minutos_desde_medianoche; // Minutos pasadas la medianoche en el momento que ocurrió el incidente. Por ejemplo, si el incidente ocurrió las 6:23, significa que pasaron 383 minutos después de las 00:00. Este es el valor que se guardaría en ese caso.
    char dia;
    int prioridad;
} incidente_t;

struct reporte;
typedef struct reporte reporte_t;


// Pre: -
// Post: Crea un TDA reporte en el heap y devuelve un puntero al mismo. 
// Devuelve NULL si no lo pudo crear.
reporte_t *reporte_crear();

// Pre: - El reporte fue previamente creado con `reporte_crear`.
//      - El incidente ya está completamente inicializado, incluido 
//        el string dinámico.
// Post: Agrega un incidente **al final** del reporte. 
void agregar_incidente(reporte_t *reporte, incidente_t incidente);

// Pre: - El reporte fue previamente creado con `reporte_crear`.
// Post: Elimina el incidente **al principio** del reporte. 
//       Devuelve true si había incidentes para eliminar o false
//       si el reporte estaba vacio. Además devuelve el incidente
//       eliminado mediante la referencia `incidente_eliminado`.
bool eliminar_incidente(reporte_t *reporte, incidente_t *incidente_eliminado);

// Pre: El reporte recibido fue previamente creado con `reporte_crear`.
// Post: Libera la memoria que se reservó en el heap para el reporte.
void reporte_destruir(reporte_t *reporte);




#endif /* __REPORTE_H__ */