// CONSIGNA
// Elgourmet.com, un reconocido sitio web de cocina, lo ha contratado para agregar una funcionalidad de emisión de informes. 
// Para ello se cuenta con una secuencia de caracteres que contiene la siguiente información sobre las recetas disponibles en el sitio: Categoria(1 carácter: (C)ames —(E)ntrantes-(M)ariscos-(P)astas-(D)ulces), Tiempo de preparación en minutos (3 caracteres), Nombre de la receta (indeterminado, termina con un -), Ingrediente principal (indeterminado, termina con un #) y Dificultad (1 carácter: (F)ácil-(N)ormal-(D)ificil). Ejemplo: C045escalopes al jerez-lomo#F+DO3Otiramisu de la nona-bizcochos yainilla#N,........+FDS Los datos de cada receta finalizan en un signo ".". Se solicita que realice un algoritmo en pseudocódigo que permita: 
// 1. Generar una secuencia de salida con todas las recetas "Dulces" cuyo tiempo de preparación sea menor a 30 minutos, incluyendo solamente el nombre de la receta, el ingrediente principal y la dificultad. 
// 2. Informar la cantidad y porcentaje que representan sobre el total, las recetas difíciles (todas). 

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
