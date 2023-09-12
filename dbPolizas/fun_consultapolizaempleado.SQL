CREATE OR REPLACE FUNCTION public.fun_consultapolizaempleadofun_consultapolizaempleado(integer)
 RETURNS TABLE(id_poliza integer, empleado_genero integer, sku_pol integer, cantidad_pol integer, fecha_pol timestamp without time zone, nombre_cliente character varying)
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Consultar polizas de un empleado.
-- Fecha     : 28/05/2023	
---------------------------------------------------------------------------------
DECLARE
	idempleado ALIAS for $1;
BEGIN
	RETURN QUERY SELECT idpoliza, empleadogenero, sku, cantidad, fechamovto, nombrecliente  FROM ctl_polizas WHERE empleadogenero = idempleado;
END;
$function$