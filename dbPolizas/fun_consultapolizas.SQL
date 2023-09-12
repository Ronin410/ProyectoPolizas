CREATE OR REPLACE FUNCTION public.fun_consultapolizas()
 RETURNS TABLE(idpoliza integer, empleado integer, skus integer, cantidades integer, fecha timestamp without time zone, nombreclientes character varying)
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Agregar nuevo empleado.
-- Fecha     : 28/05/2023	
---------------------------------------------------------------------------------
BEGIN
	RETURN QUERY SELECT idpolizas, empleadogenero, sku, cantidad, fechamovto, nombrecliente  FROM ctl_polizas;
END;
$function$