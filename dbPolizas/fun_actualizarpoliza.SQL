CREATE OR REPLACE FUNCTION public.fun_actualizarpoliza(integer, integer, integer, integer, character varying)
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
	id_poliza ALIAS FOR $1;
	iempleadogenero ALIAS FOR $2;
    isku ALIAS FOR $3;
	icantidad ALIAS FOR $4;
	snombrecliente ALIAS FOR $5;
	resultado INTEGER DEFAULT 0;
	skuanterior INTEGER DEFAULT 0;
	cantidadanterior INTEGER DEFAULT 0;
	icantidadactualizar INTEGER DEFAULT 0;
	cSql TEXT;
	cantidadinv	SMALLINT DEFAULT 0;
	-----------------------------------
BEGIN
	IF EXISTS(SELECT idpoliza FROM ctl_polizas WHERE idpoliza = id_poliza) THEN
		SELECT sku, cantidad INTO skuanterior, cantidadanterior FROM ctl_polizas WHERE idpoliza = id_poliza;
		IF EXISTS(SELECT idempleado FROM ctl_empleado WHERE idempleado = iempleadogenero) THEN 
			IF EXISTS(SELECT sku FROM ctl_inventario WHERE sku = isku) THEN
				IF EXISTS(SELECT sku FROM ctl_polizas WHERE sku = isku and idpoliza = id_poliza) THEN
					icantidadactualizar = icantidad - cantidadanterior;
					IF EXISTS(SELECT sku FROM ctl_inventario WHERE cantidad >=icantidadactualizar and sku = isku) THEN
						UPDATE ctl_polizas SET empleadogenero = iempleadogenero, nombrecliente = snombrecliente, sku = isku, cantidad = icantidad, fechacambio = CURRENT_TIMESTAMP WHERE idpoliza = id_poliza;
						SELECT cantidad INTO cantidadinv FROM ctl_inventario WHERE sku = isku;
						UPDATE ctl_inventario SET cantidad =(cantidadinv - (icantidadactualizar)) WHERE sku = isku;
						resultado = 1;
						RETURN resultado;
					ELSE 
						resultado = 2;
						RETURN resultado;
					END IF;
				ELSE 
					IF EXISTS(SELECT sku FROM ctl_inventario WHERE cantidad >=icantidad and sku = isku) THEN
						UPDATE ctl_polizas SET empleadogenero = iempleadogenero, nombrecliente = snombrecliente, sku = isku, cantidad = icantidad, fechacambio = CURRENT_TIMESTAMP WHERE idpoliza = id_poliza;
						SELECT cantidad INTO cantidadinv FROM ctl_inventario WHERE sku = isku;
						UPDATE ctl_inventario SET cantidad =(cantidadinv - icantidad) WHERE sku = isku;
						SELECT cantidad INTO cantidadinv FROM ctl_inventario WHERE sku = skuanterior;
						UPDATE ctl_inventario SET cantidad =(cantidadinv + cantidadanterior) WHERE sku = skuanterior;			
						resultado = 1;
						RETURN resultado;
					ELSE
						resultado = 2;
						RETURN resultado;
					END IF;
				END IF;
			ELSE
				resultado = 3;
				RETURN resultado;
			END IF;
		ELSE
			resultado = 4;
			RETURN resultado;
		END IF;
	ELSE
		resultado = 5;
		RETURN resultado;
	END IF;
END;
$function$
