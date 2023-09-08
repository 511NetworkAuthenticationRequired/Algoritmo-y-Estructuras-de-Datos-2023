ACCION CentroBioquimicoChaco ES
	AMBIENTE
		// SECUENCIAS
		sec_analisis, sec_salida: secuencia de caracteres
		ven_analisis, sec_salida: caracter
		// CONSTANTES Y VARIABLES
		x: entero
		porc_A, tot_A, tot_E, tot_I, total: entero
		cant_estudios: entero
		// SUBACCIONES
		Procedimiento Inicializar Es
			Arrancar(sec_analisis)
			Crear(sec_salida)
			Avanzar(sec_analisis, ven_analisis)
			porc_A, tot_A, tot_E, tot_I, total:= 0
			cant_estudios:= 0
			Escribir("Programa inicializado.")
		FinProcedimiento
		Procedimiento Finalizar Es
			Cerrar(sec_analisis)
			Cerrar(sec_salida)
			Escribir("Programa inicializado.")
		FinProcedimiento
		Procedimiento AvanzarAnalisis Es
			Para (i:=1 hasta x) Hacer
				Avanzar(sec_analisis, ven_analisis)
			FinPara
		FinProcedimiento
		Funcion ConvertirCAE(ven_analisis: caracter) Es
			Segun ven_analisis Hacer
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
		Mientras NFDS(sec_analisis) Hacer
			Avanzar(sec_analisis, ven_analisis)
			Si (ven_analisis = "A") Entonces
				Mientras (ven_analisis <> ",") Hacer
					Avanzar(sec_analisis, ven_analisis)
				FinMientras
				Avanzar(sec_analisis, ven_analisis)
				cant_estudios:= ConvertirCae(ven_analisis)*10
				Avanzar(sec_analisis, ven_analisis)
				cant_estudios:= cant_estudios + ConvertirCae(ven_analisis)
				Para (i:=1 hasta cant_estudios) Hacer
					Avanzar(sec_analisis, ven_analisis)
					Si (ven_analisis = "E") Entonces
						tot_E:= tot_E + 1
						Para (i:=1 hasta 3) Hacer
							Avanzar(sec_analisis, ven_analisis)
							Grabar(sec_salida, ven_analisis)
						FinPara
					Contrario
						Si (ven_analisis = "A") Entonces
							tot_A:= tot_A + 1
							Para (i:=1 hasta 3) Hacer
								Avanzar(sec_analisis, ven_analisis)
							FinPara
						Contrario
							tot_I:= tot_I + 1
							Para (i:=1 hasta 3) Hacer
								Avanzar(sec_analisis, ven_analisis)
							FinPara
						FinSi
					FinSi
				FinPara
			FinSi
		FinMientras
		Escribir("Total de cada estudio: Tipo A: ", tot_A, " | Tipo E: ", tot_E, " | Tipo I: ", tot_I, ".")
		total:= tot_A + tot_E + tot_I
		Escribir("El porcentaje de estudios tipo “A” sobre el total de estudios es: ", (tot_A/total)*100)
		Finalizar()
FINACCION
