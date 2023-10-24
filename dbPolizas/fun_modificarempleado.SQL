CREATE OR REPLACE FUNCTION public.fun_modificarempleado(integer, character varying, character varying, character varying)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Modificar datos de un empleado.
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
  		 UPDATE ctl_empleado set nombreempleado = snombre,apellidoempleado = sapellido, puesto = spuesto  , fechacambio = current_timestamp WHERE idempleado = inidempleado;
		resultado = 1;
		RETURN resultado;
	ELSE
		resultado = 2;
		RETURN resultado;
	END IF;
END;
$function$
