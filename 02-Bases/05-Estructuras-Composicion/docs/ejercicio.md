# Bootcamp Go

## Composición

### Objetivo
El objetivo de esta guía práctica es que podamos afianzar y profundizar los conceptos sobre estructuras, métodos y composición. Para esto, vamos a plantear una serie de ejercicios simples que nos permitirán repasar los temas que estudiamos.

### ¿Are you ready?

#### Ejercicio 1
Crear un programa que cumpla los siguiente puntos:

- Tener una estructura llamada `Product` con los campos `ID`, `Name`, `Price`, `Description` y `Category`.
- Tener un slice global de `Product` llamado `Products` instanciado con valores.
- 2 métodos asociados a la estructura `Product`: `Save()`, `GetAll()`. El método `Save()` deberá tomar el slice de `Products` y añadir el producto desde el cual se llama al método. El método `GetAll()` deberá imprimir todos los productos guardados en el slice `Products`.
- Una función `getById()` al cual se le deberá pasar un `INT` como parámetro y retorna el producto correspondiente al parámetro pasado.
- Ejecutar al menos una vez cada método y función definido desde `main()`.

#### Ejercicio 2 - Employee
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:

- Crear una estructura `Person` con los campos `ID`, `Name`, `DateOfBirth`.
- Crear una estructura `Employee` con los campos: `ID`, `Position` y una composición con la estructura `Person`.
- Realizar el método a la estructura `Employe` que se llame `PrintEmployee()`, lo que hará es realizar la impresión de los campos de un empleado.
- Instanciar en la función `main()` tanto una `Person` como un `Employee` cargando sus respectivos campos y por último ejecutar el método `PrintEmployee()`.
- Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
