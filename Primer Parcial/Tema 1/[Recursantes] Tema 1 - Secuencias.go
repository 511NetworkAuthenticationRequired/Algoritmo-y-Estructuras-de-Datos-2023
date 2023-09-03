ACCION supermercado ES
	AMBIENTE
		// SECUENCIAS
		sec_stock, sec_ventas, sec_salida: secuencia de caracteres
		ven_stock, ven_ventas, ven_salida: caracter
		// SUBACCIONES
		Subaccion Inicializar Es
			Arrancar(sec_stock)
			Arrancar(sec_ventas)
			Crear(sec_salida)

			stock, stock_vendidoS, stock_vendidoD, mes:=0

			Escribir("Programa inicializado.")
		FinSubaccion
		Subaccion Finalizar Es
			Cerrar(sec_stock)
			Cerrar(sec_ventas)
			Cerrar(sec_salida)

			Escribir("Programa finalizado.")
		FinSubaccion
		Subaccion ConvertirCAE Es
			Segun ven_stock Hacer
				="0": ven_stock:= 0
				="1": ven_stock:= 1
				="2": ven_stock:= 2
				="3": ven_stock:= 3
				="4": ven_stock:= 4
				="5": ven_stock:= 5
				="6": ven_stock:= 6
				="7": ven_stock:= 7
				="8": ven_stock:= 8
				="9": ven_stock:= 9
			FinSegun
		FinSubaccion
		x: entero
		Subaccion AvanzarStock Es
			Para i:=1 hasta x Hacer
				Avanzar(sec_stock, ven_stock)
			FinPara
		FinSubaccion
		Subaccion AvanzarVentas Es
			Para i:=1 hasta x Hacer
				Avanzar(sec_ventas, ven_ventas)
			FinPara
		FinSubaccion
		Subaccion ObtenerUnidades Es
			Avanzar(sec_ventas, ven_ventas)
			unidades:= ConvertirCAE(ven_ventas)*10
			Avanzar(sec_ventas, ven_ventas)
			unidades:= unidades + ConvertirCAE(ven_ventas)
		FinSubaccion
		// CONSTANTES Y VARIABLES
		// i: entero
		nombre: alfanumerico
		stock, stock_vendidoS, stock_vendidoD, mes, bandera: entero
	PROCESO
		Inicializar()

		Escribir("Ingrese un mes:")
		Leer(bandera)

		Mientras NFDS(sec_stock) Hacer
			AvanzarStock(7)
			stock:= ConvertirCAE(ven_stock)*100
			Avanzar(sec_stock, ven_stock)
			stock:= stock + ConvertirCAE(ven_stock)*10
			Avanzar(sec_stock, ven_stock)
			stock:= stock + ConvertirCAE(ven_stock)

			AvanzarVentas(3)
			mes:= ConvertirCAE(ven_ventas)*10
			Avanzar(sec_ventas, ven_ventas)
			mes:= mes + ConvertirCAE(ven_ventas)
			Avanzar(sec_ventas, ven_ventas)

			Si (mes = bandera) Entonces
				Avanzar(sec_ventas, ven_ventas)
				Si (ven_ventas = "S") Entonces
					stock_vendidoS:= ConvertirCAE(ven_ventas)
				Contrario
					stock_vendidoD:= ConvertirCAE(ven_ventas)
				FinSi

				ObtenerUnidades()
				stock:= stock - (stock_vendidoS*unidades)

				Escribir("Nombre del Artículo: ")
				Si (stock = 0) Entonces
					Mientras (ven_stock <> "&") Hacer
						Avanzar(sec_stock, ven_stock)
						Grabar(sec_salida, ven_stock)
						Escribir(ven_stock)
					FinMientras
				Contrario
					Mientras (ven_stock <> "&") Hacer
						Avanzar(sec_stock, ven_stock)
						Escribir(ven_stock)
					FinMientras
				FinSi
				Escribir(" | Cantidad de unidades entregadas en sucursal: ", stock_vendidoS, " | Cantidad de unidades entregadas a domicilio: ", stock_vendidoD)
			Contrario
				Avanzar(sec_ventas, ven_ventas)
				Si (ven_ventas = "S") Entonces
					stock_vendidoS:= ConvertirCAE(ven_ventas)
				FinSi

				ObtenerUnidades()
				stock:= stock - (stock_vendidoS*unidades)

				Si (stock = 0) Entonces
					Mientras (ven_stock <> "&") Hacer
						Avanzar(sec_stock, ven_stock)
						Escribir(ven_stock)
					FinMientras
				FinSi
			FinSi
		FinMientras
		Finalizar()
FINACCION
