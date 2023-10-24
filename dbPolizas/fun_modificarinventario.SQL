CREATE OR REPLACE FUNCTION public.fun_modificarinventario(integer, integer)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Modificar sku del inventario.
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	isku ALIAS for $1;
	icantidad ALIAS for $2;
	resultado	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT sku FROM ctl_inventario WHERE sku=isku) THEN
  		 UPDATE ctl_inventario set cantidad = icantidad, fechacambio = current_timestamp where sku = isku;
		resultado = 1;
		RETURN resultado;
	ELSE
		resultado = 2;
		RETURN resultado;
	END IF;
END;
$function$
