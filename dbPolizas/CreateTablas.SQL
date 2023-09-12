--Tablas
--=========================================================================
--Elaboro  : Alejandro Bueno
--Descripcion: Creacion de tabla 'ctl_inventario'
--=========================================================================
DO
$$
DECLARE
iExiste integer;
BEGIN

SELECT count(*) INTO iExiste FROM pg_tables WHERE tablename= 'ctl_inventario';

IF (iExiste = 0) THEN

	CREATE TABLE IF NOT EXISTS public.ctl_inventario
	(
		sku integer NOT NULL,
		nombrearticulo character varying(50) COLLATE pg_catalog."default",
		descripcionarticulo character varying(70) COLLATE pg_catalog."default",
		cantidad integer,
		fechamovto timestamp without time zone,
		fechacambio timestamp without time zone,
		CONSTRAINT inventario_pkey PRIMARY KEY (sku)
	
	);	
	GRANT ALL ON TABLE ctl_inventario TO postgres;

	COMMENT ON TABLE ctl_inventario IS 'Se crea tabla para el catalogo de inventario';

	COMMENT ON COLUMN ctl_inventario.sku IS 'Guarda el número del sku';

	COMMENT ON COLUMN ctl_inventario.nombrearticulo IS 'Guarda el nombre del sku';
	
	COMMENT ON COLUMN ctl_inventario.descripcionarticulo IS 'Guarda la descripcion del articulo';
	
	COMMENT ON COLUMN ctl_inventario.cantidad IS 'Guarda la cantidad de unidades del articulo';

	COMMENT ON COLUMN ctl_inventario.fechamovto IS 'Guarda la fecha en la que se agrego el articulo al inventario';

	COMMENT ON COLUMN ctl_inventario.fechacambio IS 'Guarda la fecha en la que se hizo un cambio del articulo en el inventario';
		  
END IF;

END;

$$

--=========================================================================
-- Elaboro  : Alejandro Bueno Mendoza
--Descripcion: Creacion de tabla 'ctl_empleado'
--=========================================================================
DO
$$
DECLARE
iExiste integer;
BEGIN

SELECT count(*) INTO iExiste FROM pg_tables WHERE tablename= 'ctl_empleado';

IF (iExiste = 0) THEN

	CREATE TABLE IF NOT EXISTS public.ctl_empleado
	(
		idempleado integer NOT NULL,
		nombreempleado character varying(30) COLLATE pg_catalog."default",
		apellidoempleado character varying(50) COLLATE pg_catalog."default",
		puesto character varying(30) COLLATE pg_catalog."default",
		fechamovto timestamp without time zone,
		fechacambio timestamp without time zone,
		CONSTRAINT empleado_pkey PRIMARY KEY (idempleado)
	
	);
	
	GRANT ALL ON TABLE ctl_empleado TO postgres;

	COMMENT ON TABLE ctl_empleado IS 'Se crea tabla para el catalogo de empleados';

	COMMENT ON COLUMN ctl_empleado.idempleado IS 'Guarda el número del empleado';

	COMMENT ON COLUMN ctl_empleado.nombreempleado IS 'Guarda el nombre del empleado';
	
	COMMENT ON COLUMN ctl_empleado.apellidoempleado IS 'Guarda el apellido del empleado';
	
	COMMENT ON COLUMN ctl_empleado.puesto IS 'Guarda el puesto del empleado';
		  
	COMMENT ON COLUMN ctl_empleado.fechamovto IS 'Guarda la fecha en la que se agrego el articulo al inventario';

	COMMENT ON COLUMN ctl_empleado.fechacambio IS 'Guarda la fecha en la que se hizo un cambio del articulo en el inventario';
		  
END IF;

END;

$$

 --=========================================================================
-- Elaboro  : Alejandro Bueno Mendoza
--Descripcion: Creacion de tabla 'ctl_polizas'
---------------------------------------------------------------------------
DO
$$
DECLARE
iExiste integer;
BEGIN

SELECT count(*) INTO iExiste FROM pg_tables WHERE tablename= 'ctl_polizas';

IF (iExiste = 0) THEN

   CREATE SEQUENCE IF NOT EXISTS ctl_polizas_idpolizas_seq
   MINVALUE 1000
   MAXVALUE 999999
   INCREMENT BY 1;

   CREATE TABLE IF NOT EXISTS public.ctl_polizas
   (
		idpoliza integer NOT NULL DEFAULT nextval('ctl_polizas_idpolizas_seq'::regclass),
		empleadogenero integer,
		sku integer,
		cantidad integer,
		fechamovto timestamp without time zone,
		fechacambio timestamp without time zone,
		nombrecliente character varying(50) COLLATE pg_catalog."default",
		CONSTRAINT ctl_polizas_pkey PRIMARY KEY (idpoliza)
	);
	
	GRANT ALL ON TABLE ctl_polizas TO postgres;

	COMMENT ON TABLE ctl_polizas IS 'Se crea catalogo de empleado';

	COMMENT ON COLUMN ctl_polizas.idpoliza IS 'Guarda el número de la poliza';

	COMMENT ON COLUMN ctl_polizas.empleadogenero IS 'Guarda el nombre del empleado que genero la poliza';
	
	COMMENT ON COLUMN ctl_polizas.sku IS 'Guarda el numero del sku del articulo';
	
	COMMENT ON COLUMN ctl_polizas.cantidad IS 'Guarda la cantidad de poliza';
	
	COMMENT ON COLUMN ctl_polizas.nombrecliente IS 'Guarda el nombre del cliente que genero la poliza';

	COMMENT ON COLUMN ctl_polizas.fechamovto IS 'Guarda la fecha de la creacion de la poliza';
	
	COMMENT ON COLUMN ctl_polizas.fechamovto IS 'Guarda la fecha en la que se hizo una modificacion a la poliza';	
		  
END IF;

END;

$$

------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
--INSERT CTL_EMPLEADO
INSERT INTO ctl_empleado VALUES 
(90222124,'ALEJANDRO','BUENO','DESARROLLADOR',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(90221112,'RYAN','BUENO ORDOÑEZ','VENDEDOR',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(92212111,'YENY','ZULUAGA','GERENTE',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(97848591,'PEDRO','PEREZ','VENDEDOR',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);

--INSERT ctl_inventario
INSERT INTO ctl_inventario VALUES
(232323,'PATINETA','PATINETA PARA SKATERS',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(554334,'BICICLETA','BICICLETA DE BARBIE',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(644838,'BARBIE PLAYERA','BARBIE EDICION PLAYA CON BIKINI',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(547557,'PELOTA','PELOTA DE PLAYA DE COLORES',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);
(432345,'CELULAR','MOTOROLA G50',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(999999,'IPAD','IPAD AIR 5',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(222222,'SILLON','SILLON RECLINABLE',10,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);

--INSERT ctl_polizas
INSERT INTO ctl_polizas (empleadogenero,sku,cantidad,nombrecliente,fechamovto,fechacambio) VALUES
(90222124,232323,1,'Alejandro Bueno',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(90222124,232323,2,'Ryan Ordoñez',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(90222124,554334,3,'Yeny Zuluaga',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(92212111,644838,1,'Oscar Gutierrez',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(90221112,547557,1,'Fernando Valenzuela',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP),
(97848591,547557,1,'Jose Ruiz',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);

------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------



 
