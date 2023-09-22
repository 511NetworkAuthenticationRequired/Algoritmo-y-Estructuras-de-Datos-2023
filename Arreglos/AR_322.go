ACCION 3.22 ES
	AMBIENTE
		// VARIABLES & CONSTANTES
		x; y: entero
		suma_fila; suma_columna; total: entero
		// ARREGLOS & MATRICES
		entrada: arreglo de [1...5, 1...5] de enteros
		salida: arreglo de [1...2, 1...5] de enteros
		// SUBACCIONES
		Procedimiento Inicializar() Es
			suma_fila:= 0; suma_columna:= 0
		FinProcedimiento
	PROCESO
		Inicializar()
		// SUMA COLUMNAS
		Para (x:=1 hasta 5) Hacer
			Para (y:= 1 hasta 5) Hacer
				suma_columna:= suma_columna + entrada[x, y]
				y:= y
			FinPara
			salida[2, y]:= suma_columna
			y:= 0
		FinPara
		// SUMA FILAS
		Para (y:=1 hasta 5) Hacer
			Para (x:=1 hasta 5) Hacer
				suma_fila:= suma_fila + entrada[x, y]
				x:= x
			FinPara
			salida[1, x]:= suma_fila
			x:= 0
		FinPara
FINACCION