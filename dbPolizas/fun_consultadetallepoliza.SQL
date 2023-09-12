CREATE OR REPLACE FUNCTION public.fun_consultadetallepoliza(integer)
 RETURNS TABLE(idpoliza integer, empleado integer, num_sku integer, num_cantidad integer, fecha timestamp without time zone, nombre_cliente character varying, nombre_empleado character varying, apellido_empleado character varying, nombre_articulo character varying)
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Consultar el detalle de una poliza.
-- Fecha     : 28/05/2023	
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	inidpoliza ALIAS for $1;
	-----------------------------------
BEGIN
	RETURN QUERY SELECT a.idpolizas,a.empleadogenero,a.sku,a.cantidad,a.fechamovto,a.nombrecliente,b.nombreempleado, b.apellidoempleado,c.nombrearticulo 
				 FROM ctl_polizas AS a INNER JOIN ctl_empleado AS b ON b.idempleado = a.empleadogenero
				 INNER JOIN ctl_inventario AS c ON c.sku = a.sku WHERE a.idpolizas = inidpoliza;	

END;
$function$