-- DDL
DROP DATABASE IF EXISTS `empresa_db`;
CREATE DATABASE `empresa_db`;
USE `empresa_db`;

-- Table: departamentos
CREATE TABLE IF NOT EXISTS departamentos (
	`depto_nro` varchar(7) NOT NULL,
	`nombre_depto` varchar(50) NOT NULL,
    `localidad` varchar(150) NULL,
    PRIMARY KEY (`depto_nro`)
);

-- Table: empleados
CREATE TABLE IF NOT EXISTS empleados (
	`cod_emp` varchar(6) NOT NULL,
    `nombre` varchar(50) NOT NULL,
    `apellido` varchar(50) NOT NULL,
    `puesto` varchar(50) NOT NULL,
    `fecha_alta` date NULL,
    `salario` int NULL,
    `comision` int DEFAULT 0,
    `depto_nro` varchar(7) NULL,
    PRIMARY KEY (`cod_emp`),
    KEY `idx_employee_puesto` (`puesto`),
    KEY `idx_employee_depto_nro` (`depto_nro`),
    CONSTRAINT `fk_employee_depto_nro` FOREIGN KEY (`depto_nro`) REFERENCES `departamentos` (`depto_nro`)
);