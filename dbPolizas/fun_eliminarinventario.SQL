CREATE OR REPLACE FUNCTION public.fun_eliminarinventario(integer)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Eliminar sku del inventario.
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	isku ALIAS for $1;
	resultado	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT sku FROM ctl_inventario WHERE sku=isku) THEN
  		DELETE FROM ctl_inventario WHERE sku = isku;
		resultado = 1;
		RETURN resultado;
	ELSE
		resultado = 2;
		RETURN resultado;
	END IF;
END;
$function$
