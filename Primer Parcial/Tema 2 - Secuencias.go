ACCION PrimaveraSound2023-2 ES
	AMBIENTE
		// SECUENCIAS
		sec_fila, sec_compras, sec_salida: secuencia de caracter
		ven_fila, ven_compras: caracter
		// CONSTANTES Y VARIABLES
		x: entero
		hora, minutos: entero
		cant_usuariosnc: entero
		// SUBACCIONES
		Procedimiento Inicializar Es
			Arrancar(sec_fila)
			Arrancar(sec_compras)
			Crear(sec_salida)
			Avanzar(sec_fila, ven_fila)
			Avanzar(sec_compras, ven_compras)
			cant_usuariosnc:= 0
			Escribir("Programa inicializado.")
		FinProcedimiento
		Procedimiento Finalizar Es
			Cerrar(sec_fila)
			Cerrar(sec_compras)
			Cerrar(sec_salida)
			Escribir("Programa finalizado.")
		FinProcedimiento
		Procedimiento AvanzarFila Es
			Para (i:=1 hasta x) Hacer
				Avanzar(sec_fila, ven_fila)
			FinPara
		FinProcedimiento
		Procedimiento AvanzarCompras Es
			Para (i:=1 hasta x) Hacer
				Avanzar(sec_compras, ven_compras)
			FinPara
		FinProcedimiento
		Funcion ConvertirCAE(ven_fila: caracter) Es
			Segun ven_fila Hacer
				="0": ConvertirCAE:= 0
				="1": ConvertirCAE:= 1
				="2": ConvertirCAE:= 2
				="3": ConvertirCAE:= 3
				="4": ConvertirCAE:= 4
				="5": ConvertirCAE:= 5
				="6": ConvertirCAE:= 6
				="7": ConvertirCAE:= 7
				="8": ConvertirCAE:= 8
				="9": ConvertirCAE:= 9
			FinSegun
		FinFuncion
	PROCESO
		Inicializar()
		Mientras NFDS(sec_fila) Hacer
			Avanzar(sec_fila, ven_fila)
			hora:= ConvertirCAE(ven_fila)*10
			Avanzar(sec_fila, ven_fila)
			hora:= hora + ConvertirCAE(ven_fila)
			Avanzar(sec_fila, ven_fila)
			minutos:= ConvertirCAE(ven_fila)*10
			Avanzar(sec_fila, ven_fila)
			minutos:= minutos + ConvertirCAE(ven_fila)
			Si (hora < 10) y (minutos < 60) Entonces
				Avanzar(sec_compras, ven_compras)
				Si (ven_compras = "#") Entonces
					cant_usuariosnc:= cant_usuariosnc + 1
				Contrario
					Para (i:=1 hasta 6) Hacer
						Avanzar(sec_compras, ven_compras)
						Grabar(sec_salida, ven_compras)
					FinPara
					Grabar(sec_salida, "+")
					Mientras (ven_compras <> "+")
						Avanzar(sec_compras, ven_compras)
					FinMientras
					Para (i:=1 hasta 8) Hacer
						Avanzar(sec_compras, ven_compras)
						Grabar(sec_salida, ven_compras)
					FinPara
					Grabar(sec_salida, "#")
				FinSi
			FinSi
		FinMientras
		Escribir("La cantidad de usuarios que estaban en fila, pero no compraron entradas es de: ", cant_usuariosnc)
		Finalizar()
FINACCION
