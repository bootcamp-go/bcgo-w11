# Tercera FORMA NORMAL (2NF)
La tercera forma normal (3NF) es otro nivel en el diseño de bases de datos relacionales que busca reducir la redundancia y mejorar la integridad de los datos. Una tabla está en 3NF si ya está en 2NF y, además, todos sus atributos no clave dependen directamente de la clave primaria y no de otros atributos no clave. Esto significa que no debe haber dependencias transitivas, donde un atributo no clave depende de otro atributo no clave.

## Ejemplo

Consideremos las siguiente tabla que ya están en 2NF:

Tabla `Estudiantes`:
| ID_Estudiante | Nombre_Estudiante | Dirección       | Fecha_Nacimiento |
|---------------|-------------------|-----------------|------------------|
| 1             | Ana               | Calle A, 123    | 2001-04-15       |
| 2             | Juan              | Avenida B, 456  | 2002-08-22       |
| 3             | María             | Calle C, 789    | 2003-01-30       |

Tabla `Clases` (combinación de estudiante, asignatura y profesor):
| Id            | Asignatura  | Profesor   |
|---------------|-------------|------------|
| 1             | Matemáticas | Sr. García |
| 2             | Física      | Sra. Pérez |
| 3             | Historia    | Sr. López  |
| 4             | Matemáticas | Sr. García |
| 5             | Biología    | Sra. Ortiz |
| 6             | Física      | Sra. Pérez |

---

### Dependencia transitiva
Aquí, la tabla tiene una dependencia transitoria:
- `profesor` depende de Asignatura, no de ID_Estudiante.

Es un atributo no clave que no depende directamente de la clave primaria, sino de otro atributo no clave. Para llevar esta tabla a 3NF, eliminamos la dependencia transitoria creando una nueva tabla para los profesores y sus asignaturas:

---

### Normalizacion

Tabla `Estudiantes` (sin cambios):
| ID_Estudiante | Nombre_Estudiante | Dirección       | Fecha_Nacimiento |
|---------------|-------------------|-----------------|------------------|
| 1             | Ana               | Calle A, 123    | 2001-04-15       |
| 2             | Juan              | Avenida B, 456  | 2002-08-22       |
| 3             | María             | Calle C, 789    | 2003-01-30       |

Tabla `Asignaturas`:
| ID_Asignatura | Nombre_Asignatura |
|---------------|-------------------|
| 101           | Matemáticas       |
| 102           | Física            |
| 103           | Historia          |
| 104           | Biología          |

Tabla `Profesores`:
| ID_Profesor | Nombre_Profesor |
|-------------|-----------------|
| P1          | Sr. García      |
| P2          | Sra. Pérez      |
| P3          | Sr. López       |
| P4          | Sra. Ortiz      |

Tabla `Clases` (actualizada con IDs de asignaturas y profesores):
| ID_Estudiante | ID_Asignatura | ID_Profesor |
|---------------|---------------|-------------|
| 1             | 101           | P1          |
| 1             | 102           | P2          |
| 2             | 103           | P3          |
| 2             | 101           | P1          |
| 3             | 104           | P4          |
| 3             | 102           | P2          |

Con esta nueva estructura, la tabla `Clases` está en 3NF. La tabla `Estudiantes` y `Asignaturas` eliminan la redundancia de nombres, y la tabla `Clases` representa las relaciones entre estudiantes, asignaturas y profesores, dependiendo completamente de la clave compuesta (ID_Estudiante, ID_Asignatura, ID_Profesor).