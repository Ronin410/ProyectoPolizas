CREATE OR REPLACE FUNCTION public.fun_consultaempleados()
 RETURNS TABLE(id_emp integer, nombre_emp character varying, apellido_emp character varying, puesto_emp character varying)
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Agregar nuevo empleado.
-- Fecha     : 28/05/2023	
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------

	-----------------------------------
BEGIN
	RETURN QUERY SELECT idempleado, nombreempleado, apellidoempleado,puesto FROM ctl_empleado;
END;
$function$