CREATE OR REPLACE FUNCTION public.fun_agregarpoliza(integer, integer, integer, character varying)
 RETURNS smallint
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
---------------------------------------------------------------------------------
-- Realizo   : Alejandro Bueno Mendoza
-- Actividad : Agregar nueva poliza
-- Fecha     : 28/05/2023
---------------------------------------------------------------------------------
DECLARE
	-- Variables ----------------------
	iempleadogenero ALIAS for $1;
	isku ALIAS for $2;
	icantidad ALIAS for $3;
	snombrecliente ALIAS for $4;
	resultado	SMALLINT DEFAULT 0;
	cantidadinv	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT sku FROM ctl_inventario WHERE cantidad >=icantidad) THEN
		IF EXISTS(SELECT sku FROM ctl_inventario WHERE sku=isku) THEN
			IF EXISTS(SELECT idempleado FROM ctl_empleado WHERE idempleado = iempleadogenero) THEN 
				INSERT INTO ctl_polizas(empleadogenero,sku,cantidad,fechamovto,fechacambio,nombrecliente) 
				VALUES (iempleadogenero, isku, icantidad, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP,snombrecliente);

				SELECT cantidad INTO cantidadinv FROM ctl_inventario WHERE sku = isku;
				UPDATE ctl_inventario SET cantidad =(cantidadinv - icantidad) WHERE sku = isku;
				resultado = 1;
				RETURN resultado;
			ELSE
				resultado = 2;
				RETURN resultado;
			END IF;
		ELSE
			resultado = 3;
			RETURN resultado;
		END IF;
	ELSE
		resultado = 4;
		RETURN resultado;
	END IF;
END;
$function$
