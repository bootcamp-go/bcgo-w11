# Segunda FORMA NORMAL (2NF)
La segunda forma normal (2NF) es un paso más en el diseño de bases de datos relacionales que se enfoca en la eliminación de redundancias y dependencias parciales. Una tabla está en 2NF si, además de cumplir con la primera forma normal (1NF), cada atributo no clave depende completamente de la clave primaria. Esto significa que no debe haber dependencias parciales, es decir, ningún atributo no clave debe depender únicamente de una parte de la clave primaria en una tabla con una clave compuesta.

## Ejemplo

Supongamos que tenemos una tabla de calificaciones de estudiantes que ya está en 1NF:

| ID_Estudiante | ID_Asignatura | Nombre_Estudiante | Nombre_Asignatura | Profesor   |
|---------------|---------------|-------------------|-------------------|------------|
| 1             | 101           | Ana               | Matemáticas       | Sr. García |
| 1             | 102           | Ana               | Física            | Sra. Pérez |
| 2             | 101           | Juan              | Matemáticas       | Sr. García |
| 2             | 103           | Juan              | Historia          | Sra. Ortiz |

---

### Dependencias parciales
Aquí, la clave primaria es la combinación de ID_Estudiante e ID_Asignatura. Sin embargo, observamos dependencias parciales:
- Nombre_Estudiante depende solo de ID_Estudiante.
- Nombre_Asignatura depende solo de ID_Asignatura.

Estas dependencias parciales hacen que la tabla no esté en 2NF. Para convertirla en 2NF, dividimos la tabla en varias tablas, eliminando redundancias y dependencias parciales:

---

### Normalizacion

Tabla `Estudiantes`:
| ID_Estudiante | Nombre_Estudiante |
|---------------|-------------------|
| 1             | Ana               |
| 2             | Juan              |

Tabla `Asignaturas`:
| ID_Asignatura | Nombre_Asignatura | Profesor   |
|---------------|-------------------|------------|
| 101           | Matemáticas       | Sr. García |
| 102           | Física            | Sra. Pérez |
| 103           | Historia          | Sra. Ortiz |

Con esta estructura, cada tabla está en 2NF. La tabla `Estudiantes` y `Asignaturas` eliminan la redundancia de nombres, y la tabla `Calificaciones` representa las relaciones entre estudiantes y asignaturas con sus calificaciones, dependiendo completamente de la clave compuesta (ID_Estudiante, ID_Asignatura).