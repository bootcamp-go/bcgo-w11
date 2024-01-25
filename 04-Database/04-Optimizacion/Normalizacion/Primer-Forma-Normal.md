# Primer FORMA NORMAL (1NF)
La primera forma normal (1NF) en bases de datos es un concepto fundamental en el diseño de bases de datos relacionales. Un diseño está en 1NF si cumple con las siguientes condiciones:

- Cada columna de la tabla contiene valores atómicos, es decir, indivisibles.
- Todos los valores en una columna son del mismo tipo de datos.
- Cada columna tiene un nombre único.
- El orden en que se almacenan los datos no afecta su integridad.
- Cada fila es única y se identifica por una clave primaria.

El objetivo de la 1NF es asegurar que la tabla represente adecuadamente la relación entre sus datos, eliminando grupos repetitivos y facilitando las operaciones básicas de una base de datos como son la inserción, modificación y eliminación de datos.

## Ejemplo

Supongamos que tenemos una tabla de estudiantes que inicialmente no está en 1NF:

| ID_Estudiante | Nombre     | Asignaturas            |
|---------------|------------|------------------------|
| 1             | Ana        | Matemáticas, Física    |
| 2             | Juan       | Historia, Matemáticas  |
| 3             | María      | Biología               |

---

### Grupos repetitivos
Esta tabla no está en 1NF porque la columna "Asignaturas" contiene valores no atómicos (listas de asignaturas).
Para convertirla en una tabla en 1NF, reestructuramos la tabla

---

### Normalizacion

| ID_Estudiante | Nombre | Asignatura   |
|---------------|--------|--------------|
| 1             | Ana    | Matemáticas  |
| 1             | Ana    | Física       |
| 2             | Juan   | Historia     |
| 2             | Juan   | Matemáticas  |
| 3             | María  | Biología     |

En esta nueva estructura, cada fila representa un registro único de un estudiante con una asignatura específica. Ahora, la tabla cumple con todas las condiciones de la 1NF: cada columna tiene valores atómicos, del mismo tipo de dato, con nombres únicos, y cada fila es única y se identifica por la combinación de ID_Estudiante y Asignatura.