CREATE OR REPLACE FUNCTION public.fun_agregarempleado(integer, character varying, character varying, character varying)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Agregar nuevo empleado.
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	inidempleado ALIAS for $1;
	snombre ALIAS for $2;
	sapellido ALIAS for $3;
	spuesto ALIAS for $4;
	resultado	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT idempleado FROM ctl_empleado WHERE idempleado=inidempleado) THEN
  		resultado = 2;
		RETURN resultado;
	ELSE
		INSERT INTO ctl_empleado(idempleado,nombreempleado,apellidoempleado,puesto,fechamovto,fechacambio) VALUES (inidempleado, snombre, sapellido, spuesto, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
		resultado = 1;
		RETURN resultado;
	END IF;
END;
$function$