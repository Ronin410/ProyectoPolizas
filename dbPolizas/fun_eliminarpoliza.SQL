CREATE OR REPLACE FUNCTION public.fun_eliminarpoliza(integer)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Eliminar una poliza
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	id_polizas ALIAS for $1;
	resultado	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT idpoliza FROM ctl_polizas WHERE idpoliza=id_polizas) THEN
  		resultado = 1;
		DELETE FROM ctl_polizas WHERE idpoliza = id_polizas;
		RETURN resultado;
	ELSE
		resultado = 0;
		RETURN resultado;
	END IF;
END;
$function$