CREATE OR REPLACE FUNCTION public.fun_agregarinventario(integer, character varying, character varying, integer)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Agregar nuevo sku al inventario.
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	isku ALIAS for $1;
	snombre ALIAS for $2;
	sdescripcion ALIAS for $3;
	icantidad ALIAS for $4;
	resultado	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT sku FROM ctl_inventario WHERE sku=isku) THEN
  		resultado = 2;
		RETURN resultado;
	ELSE
		INSERT INTO ctl_inventario(sku,nombrearticulo,descripcionarticulo,cantidad,fechamovto,fechacambio) VALUES (isku, snombre, sdescripcion, icantidad, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
		resultado = 1;
		RETURN resultado;
	END IF;
END;
$function$
