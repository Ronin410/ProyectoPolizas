CREATE OR REPLACE FUNCTION public.fun_consultarinventario()
 RETURNS TABLE(sku_articulo integer, nombre_articulo character varying, descripcion_articulo character varying, cantidad_articulo integer)
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Consultar sku del inventario.
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	-----------------------------------
BEGIN
		RETURN QUERY SELECT sku, nombrearticulo, descripcionarticulo, cantidad FROM ctl_inventario;
END;
$function$