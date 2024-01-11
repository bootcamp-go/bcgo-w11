# Estructura de Proyecto en Go

La estructura de un proyecto en Go puede variar según las necesidades y preferencias del equipo de desarrollo. Una forma común de organizar un proyecto es dividirlo en varias carpetas principales, cada una con un propósito específico. Aquí hay una guía de cómo podría estructurarse un proyecto:

## CMD

Esta carpeta contiene todos los puntos de entrada de la aplicación. Cada punto de entrada es un paquete dentro de la carpeta `/cmd` y tiene su propio `main.go`.


## Internal

La carpeta `internal` contiene todos los paquetes que representan los componentes internos de nuestra aplicación. Dentro de `internal`, se pueden encontrar varias subcarpetas que representan diferentes aspectos de la aplicación.

### Application

Dentro de `internal`, una carpeta `Application` representa la implementación de una app con un método `Run`, donde se realiza la inicialización e inyección de dependencias.


### Controller, Service, y Repository

En `internal`, se sigue el patrón controlador - servicio - repositorio, donde cada capa desarrolla tareas específicas:

- **Controller**: Maneja la petición del cliente (`request`) y envía la respuesta adecuada (`response`). Se enfoca en el paso de datos entre el cliente y la capa de servicio.
- **Service**: Maneja toda la lógica de negocio, validaciones y comunicación con servicios externos (como APIs de terceros, bases de datos).
- **Repository**: Maneja todas las operaciones relacionadas con la persistencia de datos (en memoria, archivos, SGBD, etc).

### Coordinacion del sistema

Cada capa de la arquitectura puede ser representada mediante `interfaces`, que definen un conjunto de funcionalidades esenciales. Estas interfaces sirven para abstraer los detalles de implementación, permitiendo que el sistema sea flexible y escalable. Por ejemplo, una interfaz de servicio puede especificar los métodos para realizar operaciones de negocio, pero no cómo estas operaciones son efectivamente realizadas.

En este sentido tendríamos 3 interfaces principales:
- **Controller**
- **Service**
- **Repository**

En la práctica para `coordinar` al sistema y las `interfaces`, desarrollaremos `implementaciones concretas` de estas interfaces para cada capa. En este esquema, una implementación del controlador (controller) interactuará con una interfaz de servicio (service), delegando responsabilidades específicas de la lógica de negocio. A su vez, la implementación del servicio dependerá de una o más interfaces de repositorio (repository), las cuales manejan la interacción con la capa de persistencia de datos.

Este enfoque de diseño permite que el flujo de trabajo dentro del sistema sea claro y bien definido. El controlador gestiona las interacciones con el usuario o el cliente, pasando los requisitos al servicio. Luego, el servicio procesa la lógica de negocio y se comunica con la capa de repositorio para acceder o modificar los datos. Así, la estructura general del sistema sigue un flujo lineal: controller -> service -> repository.


#### Notas
Es importante destacar que, aunque las interfaces definen cómo se comunican las distintas capas, no dictan la lógica específica de cada implementación.
La verdadera coordinación y flujo del sistema se establece a través de cómo estas interfaces son implementadas y cómo se manejan las dependencias entre ellas, permitiendo construir sistemas robustos a partir de componentes más pequeños y bien definidos.

## Platform

La carpeta `Platform` contiene componentes de utilidad general para el desarrollo del código, como funciones de uso común y utilidades para operaciones específicas.

## Ejemplo de estructura de proyecto

```bash
├── cmd
│   ├── app
│   │   main.go
│   ├── api
│   │   main.go
│   └── worker
│        main.go
├── internal
│   |   item.go
│   |   item_service.go
│   |   item_repository.go
│   |   user.go
│   |   user_service.go
│   |   user_repository.go
|   ├── application
│   │      default.go
|   ├── handler
│   │      default.go
|   ├── service
│   │      item_default.go
│   │      user_default.go
|   └── repository
│          item_default.go
│          user_default.go
|
└── platform
    └── web
        ├── request
        │   request.go
        └── response
            response.go
```