CREATE OR REPLACE FUNCTION public.fun_eliminarempleado(integer)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Eliminar empleado.
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	inidempleado ALIAS for $1;
	resultado	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT idempleado FROM ctl_empleado WHERE idempleado=inidempleado) THEN
  		DELETE FROM ctl_empleado WHERE idempleado = inidempleado;
		resultado = 1;
		RETURN resultado;
	ELSE
		resultado = 2;
		RETURN resultado;
	END IF;
END;
$function$
