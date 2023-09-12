CREATE OR REPLACE FUNCTION public.fun_actualizarnombreclientepoliza(integer, character varying)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Actualizar una poliza
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	id_poliza ALIAS FOR $1;
	snombrecliente ALIAS FOR $2;
	resultado SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT idpoliza FROM ctl_polizas WHERE idpoliza=id_poliza) THEN
		UPDATE ctl_polizas SET nombrecliente = snombrecliente, fechacambio = CURRENT_TIMESTAMP WHERE idpoliza = id_poliza;
		resultado = 1;
		RETURN resultado;
	ELSE
		resultado = 0;
		RETURN resultado;
	END IF;
END;
$function$
