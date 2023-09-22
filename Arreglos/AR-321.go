ACCION 3.21 ES
	AMBIENTE
		// VARIABLES & CONSTANTES
		x; y: entero
		suma_fila; suma_columna; total: entero
		// ARREGLOS & MATRICES
		matriz: arreglo de [1...6, 1...6] de enteros
		// SUBACCIONES
		Procedimiento Inicializar() Es
			suma_fila:= 0; suma_columna:= 0; total:= 0
		FinProcedimiento
	PROCESO
		Inicializar()
		// SUMA COLUMNAS
		Para (x:=1 hasta 5) Hacer
			Para (y:= 1 hasta 5) Hacer
				suma_columna:= suma_columna + matriz[x, y]
			FinPara
			matriz[6, y]:= suma_columna
		FinPara
		// SUMA FILAS
		Para (y:=1 hasta 5) Hacer
			Para (x:=1 hasta 5) Hacer
				suma_fila:= suma_fila + matriz[x, y]
			FinPara
			matriz[x, 6]:= suma_fila
		FinPara
		total:= suma_fila + suma_columna
		matriz[6, 6]:= total
FINACCION