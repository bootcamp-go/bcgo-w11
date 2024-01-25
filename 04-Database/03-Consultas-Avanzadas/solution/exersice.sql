USE empresa_db;

-- Seleccionar el nombre, el puesto del empleado y la localidad de los departamentos donde trabajan los vendedores.
SELECT e.nombre, e.puesto, d.localidad FROM empleados e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
HAVING e.puesto = "Vendedor";

-- Visualizar los departamentos con más de 1 empleado.
SELECT d.* FROM departamentos d
INNER JOIN empleados e ON d.depto_nro = e.depto_nro
GROUP BY d.depto_nro HAVING COUNT(e.cod_emp) > 1;

-- Mostrar el nombre, salario del empleado y nombre del departamento que tengan el mismo puesto
-- que ‘Mito Barchuk’.
SELECT e.nombre, e.salario, d.nombre_depto FROM empleados e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (
	SELECT e2.puesto FROM empleados e2 WHERE e2.nombre = "Jonathan" AND e2.apellido = "Aguilera"
);

-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad,
-- ordenados por nombre.
SELECT e.* FROM empleados e
WHERE e.depto_nro = (
	SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Contabilidad"
)
ORDER BY e.nombre;

-- Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT e.nombre FROM empleados e
WHERE e.salario = (
	SELECT MIN(e2.salario) FROM empleados e2
);

-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.* FROM empleados e
WHERE e.salario = (
	SELECT MAX(e2.salario) FROM empleados e2 WHERE e2.depto_nro = (
		SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Ventas"
    )
) AND e.depto_nro = (
	SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Ventas"
);

SELECT e.* FROM empleados e INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
AND d.nombre_depto = "Ventas" AND e.salario = (
	SELECT MAX(e2.salario) FROM empleados e2 WHERE e2.depto_nro = (
		SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Ventas"
    )
);

SELECT e.* FROM empleados e INNER JOIN departamentos d ON e.depto_nro = d.depto_nro AND d.nombre_depto = "Ventas"
ORDER BY e.salario DESC
LIMIT 1;

select e.* from empleados e join departamentos d on e.id = d.depto_nro and d.nombre_depto = "Ventas"
where e.salario = MAX(e.salario);
