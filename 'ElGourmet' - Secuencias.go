ACCION ElGourmet Es
	AMBIENTE
		sec_recetas, sec_salida: secuencia de caracteres
		ven_recetas: caracter
		minutos, cant_rd, porcentaje_rd, total: entero
		Procedimiento Inicializar Es
			Arrancar(sec_recetas)
			Crear(sec_salida)
			Avanzar(sec_recetas)
			minutos:= 0; cant_rd:= 0; porcentaje_rd:= 0; total:= 0
			Escribir("Programa Inicializado.")
		FinProcedimiento
		Procedimiento Finalizar Es
			Cerrar(sec_recetas)
			Cerrar(sec_salida)
			Escribir("Programa finalizado.")
		FinProcedimiento
		Procedimiento AvanzarRecetas(i, x: entero) Es
			Para (i:=1 hasta x)
				Avanzar(sec_recetas, ven_recetas)
			FinPara
		FinProcedimiento
		Funcion ConvertirCAE(ven_recetas: caracter) Es
			Segun ven_recetas Hacer
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
		Funcion AnalizarDificultad Es
			Si (ven_recetas = "D") Entonces
				cant_rd:= cant_rd + 1
			Contrario
				total:= total + 1
			FinSi
	PROCESO
		Mientras (NFDS(sec_recetas)) Hacer
			Avanzar(sec_recetas, ven_recetas)
			Si (ven_recetas = "D") Entonces
				Avanzar(sec_recetas, ven_recetas)
				minutos:= ConvertirCAE(ven_recetas)*100
				Avanzar(sec_recetas, ven_recetas)
				minutos:= minutos + ConvertirCAE(ven_recetas)*10
				Avanzar(sec_recetas, ven_recetas)
				minutos:= minutos + ConvertirCAE(ven_recetas)
				Si (minutos < 30) Entonces
					Mientras (ven_recetas <> "#") Hacer
						Avanzar(sec_recetas, ven_recetas)
						Grabar(sec_salida, ven_recetas)
					FinMientras
					Avanzar(sec_recetas, ven_recetas)
					AnalizarDificultad()
					Grabar(sec_salida, ven_recetas)
				Contrario
					Mientras (ven_recetas <> "#") Hacer
						Avanzar(sec_recetas, ven_recetas)
					FinMientras
					Avanzar(sec_recetas, ven_recetas)
					AnalizarDificultad()
				FinSi
			Contrario
				Mientras (ven_recetas <> "#") Hacer
					Avanzar(sec_recetas, ven_recetas)
				FinMientras
				Avanzar(sec_recetas, ven_recetas)
				AnalizarDificultad()
			FinSi
			porcentaje_rd:= (cant_rd/total)*100
			Escribir("Total general: ", total, " | Cantidad de recetas difíciles: ", cant_rd, " | Porcentaje de recetas difíciles: ", porcentaje_rd)
FINACCION