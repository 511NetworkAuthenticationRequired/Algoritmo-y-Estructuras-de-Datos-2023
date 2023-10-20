ACCION ejercicio 3.23 ES
	AMBIENTE
		// VARIABLES & CONSTANTES
		x, y, z: entero
		suma_temperatura; promedio: entero
		max_temperatura; min_temperatura; max_dia; min_dia; max_lectura, min_lectura: entero
		dia: caracter // Lunes: L, Martes: M, Miercoles: X, Jueves: J, Viernes: V, Sabado: S, Domingo: D
		// ARREGLOS/MATRICES
		matriz_informacion: arreglo de [1...50, 1...7, 1...4] de enteros // x, y, z
		matriz_promedio: arreglo de [1...50, 1..7] de enteros
		// SUBACCIONES
		Procedimiento Inicializar() Es
			x:= 0; y:= 0; z:= 0
			suma_temperatura:= 0
			max_temperatura:= 0; min_temperatura:= 0; max_dia:= 0; min_dia:= 0; max_lectura:= 0; min_lectura:= 0
			cont_temperatura:= 0
		FinProcedimiento
		Funcion ObtenerDia(n: entero):alfanumerico Es
			Segun (n) Entonces
				=1: dia:= "L"
				=2: dia:= "M"
				=3: dia:= "X"
				=4: dia:= "J"
				=5: dia:= "V"
				=6: dia:= "S"
				=7: dia:= "D"
			FinSegun
		FinFuncion
	FinPara
FinPara
FinPara
	PROCESO
		Inicializar()
		Para (x:=1 hasta 50) Hacer
			max_temperatura:= matriz_informacion[HV, 1, 1] // [x, y, z]
			min_temperatura:= matriz_informacion[LV, 1, 1] // [x, y, z]
			Para (y:=1 hasta 7) Hacer
				suma_temperatura:= 0
				Para (z:=1 hasta 4) Hacer
					Si (matriz_informacion[x, y, z] > max_temperatura) Entonces
						max_temperatura:= matriz_informacion[x, y, z]
						max_dia:= y
						max_lectura:= x
					Contrario
						Si (matriz_informacion[x, y, z] < min_temperatura) Entonces
							min_temperatura:= matriz_informacion[x, y, z]
							min_dia:= y
							min_lectura:= x
						FinSi
					FinSi
					suma_temperatura: suma_temperatura + matriz_informacion[x, y, z]
				FinPara
			FinPara
			promedio:= suma_temperatura DIV 4
			matriz_promedio[z, y]:= suma_temperatura
		FinPara
		Escribir("Temperatura máxima: ", max_temperatura, "el dia: ", ObtenerDia(max_dia), "en la lectura: ", max_lectura)
		Escribir("Temperatura mínima: ", min_temperatura, "el dia: ", ObtenerDia(min_dia), "en la lectura: ", min_lectura)
FINACCION
