ACCION PremiumBeat Es
	AMBIENTE
		sec_pistas, sec_salida: secuencia de caracteres
		ven_pistas: caracter
		respuesta, genero: caracter
		i, tot_pistas, cant_c, cant_j, cant_l, cant_i, cant_r, cant_pistascc, cant_cc_c, cant_cc_i, cant_cc_j, cant_cc_l, cant_cc_r: entero
		tot_s, s1, s2, s3: entero
		Procedimiento Inicializar Es
			Arrancar(sec_pistas)
			Crear(sec_salida)
			Avanzar(sec_pistas)
			segundos:= 0; tot_pistas:= 0
			tot_s:= =; s1:= 0; s2:= 0; s3:= 0
			cant_c:= 0; cant_j:= =; cant_l:= 0; cant_i:= 0; cant_r:= 0; cant_cc_c:= 0; cant_cc_j:= 0; cant_cc_l:= 0; cant_cc_i:= 0; cant_cc_r:= 0
			Escribir("Programa Inicializado.")
		FinProcedimiento
		Procedimiento Finalizar Es
			Cerrar(sec_pistas)
			Cerrar(sec_salida)
			Escribir("Programa finalizado.")
		FinProcedimiento
		Procedimiento AvanzarPistas(x: entero) Es
			Para (i:=1 hasta x)
				Avanzar(sec_pistas, ven_pistas)
			FinPara
		FinProcedimiento
		Funcion ConvertirCAE(ven_pistas: caracter) Es
			Segun ven_pistas Hacer
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
		Funcion ContadorGeneros Es
			Segun (genero) Hacer
				="C": cant_c:= cant_c + 1
				="J": cant_j:= cant_j + 1
				="L": cant_l:= cant_l + 1
				="I": cant_i:= cant_i + 1
				="R": cant_r:= cant_r + 1
			FinSegun
		FinFuncion
	PROCESO
		Inicializar()
		Escribir("Ingrese un tipo de sentimiento (A/C/R/X):")
		Leer(respuesta)

		Mientras (NFDS(sec_pistas)) Hacer
			Avanzar(sec_pistas, ven_pistas)
			genero:= ven_pistas
			Avanzar(sec_pistas, ven_pistas)
			Si (ven_pistas = respuesta) Entonces
				Avanzar(sec_pistas, ven_pistas)
				s1:= ConvertirCAE(ven_pistas)*100
				Avanzar(sec_pistas, ven_pistas)
				s2:= ConvertirCAE(ven_pistas)*10
				Avanzar(sec_pistas, ven_pistas)
				s3:= ConvertirCAE(ven_pistas)
				tot_s:= s1 + s2 + s3
				Si (tot_s > 150) Entonces
					Grabar(sec_salida, s1)
					Grabar(sec_salida, s2)
					Grabar(sec_salida, s3)
					Mientras (ven_pistas <> "#") Entonces
						Avanzar(sec_pistas, ven_pistas)
						Grabar(sec_salida, ven_pistas)
					FinMientras
					Segun (genero) Hacer
						="C": cant_c:= cant_c + 1; cant_cc_c:= cant_cc_c + 1
						="J": cant_j:= cant_j + 1; cant_cc_j:= cant_cc_j + 1
						="L": cant_l:= cant_l + 1; cant_cc_l:= cant_cc_l + 1
						="I": cant_i:= cant_i + 1; cant_cc_i:= cant_cc_i + 1
						="R": cant_r:= cant_r + 1; cant_cc_r:= cant_cc_r + 1
					FinSegun
				Sino
					Mientras (ven_pistas <> "#") Entonces
						Avanzar(sec_pistas, ven_pistas)
					FinMientras
					ContadorGeneros()
				FinSi
			Sino
				Mientras (ven_pistas <> "#") Entonces
					Avanzar(sec_pistas, ven_pistas)
				FinMientras
				ContadorGeneros()
			FinSi
		FinMientras
		Escribir("CLÁSICA: Total: ", cant_c, " | Cumple con la condición:", cant_cc_c)
		Escribir("JAZZ: Total: ", cant_j, " | Cumple con la condición:", cant_cc_j)
		Escribir("LATINA: Total: ", cant_l, " | Cumple con la condición:", cant_cc_l)
		Escribir("INSTRUMENTAL: Total: ", cant_i, " | Cumple con la condición:", cant_cc_i)
		Escribir("ROCK: Total: ", cant_r, " | Cumple con la condición:", cant_cc_r)
FINACCION		