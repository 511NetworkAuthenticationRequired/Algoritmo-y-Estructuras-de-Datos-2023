### ESQUELETOS FRECUENTES

#### CORTE DE CONTROL
  ```ruby
ACCION CDC ES
    AMBIENTE
      Procedimiento TratarCorte Es
        Si (registro3.clave <> resguardo3) Entonces
            Corte3() #De mayor jerarquía
         Contrario
            Si (registro2.clave <> resguardo2) Entonces
              Corte2()
            Contrario
              Si (registro1.clave <> resguardo1) Entonces
                Corte1() #De menor importancia
              FinSi
            FinSi
        FinSi
      FinProcedimiento
      Procedimiento CorteN Es #Generalización
        CorteN-1() #LLamada al corte de menor nivel
        EmitirTotalesN() #Emición de resultados de nivel
        totalesn+1:= TotalesN+1 + TotalesN #Acumulación de totales al nivel superior
        totalesn:= 0 #Reinicio de totales
        resguardon:= claven #Resguardo de la clave nueva
      FinProcedimiento
    PROCESO
      Inicializar() #Subaccion que abre archivos e iguala los totalizarores y resguardos a 0
      Mientras NFDA(archivo) Hacer
        TratarCorte()
        TratarRegistro()
        LeerRegistro()
      FinMientras
      EmitirTotales()
      Corte3()
      EmitirTotales()
      Finalizar() #Subaccion que cierra los archivos
FINACCION
  ```
#### ACTUALIZACIÓN
  ```js

  ```
### SUBACCIONES ÚTILES
#### Lectura de archivos en actualización
#### 
```js
Procedimiento Leer[Maestro/Movimiento]() Es
	Leer(archivo, registro)
	Si FDA(archivo) Entonces
		registro.clave:= HV
	FinSi
FinProcedimiento
```
#### Convertir caracter a entero
```js
Funcion ConvertirCAE(variable: caracter):entero Es
	Segun (variable) Hacer
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
```
#### Convertir entero a caracter
```js
// AMBIENTE
Funcion ConvertirEAC(numero: entero):caracter Es
	Segun (variable) Hacer
  	=0: ConvertirCAE:= "0"
  	=1: ConvertirCAE:= "1"
  	=2: ConvertirCAE:= "2"
  	=3: ConvertirCAE:= "3"
  	=4: ConvertirCAE:= "4"
  	=5: ConvertirCAE:= "5"
  	=6: ConvertirCAE:= "6"
  	=7: ConvertirCAE:= "7"
  	=8: ConvertirCAE:= "8"
  	=9: ConvertirCAE:= "9"
 	FinSegun
FinFuncion

// PROCESO
x = valor
Para (i:= digitos hasta 1, -1) HACER
    var = x DIV (10**(i-1)) // x * 10ⁱ⁻¹
```
#### Avanzar manejado por contador
```js
i: entero
Procedimiento AvanzarN (x: entero) Es
	Para (i:=1 hasta x) Hacer
		Avanzar(secuencia, ventana)
	FinPara
FinProcedimiento
```
#### Contadores a 0 (Ejemplo)
```js
Procedimiento InciarContadores() Es
	[Nombre del contyador]:= 0
	[Nombre del acumulador]:= 0
	[x]:=; [y]:= 0; [z]:= 0
 	[i]:=; [j]:= 0; [k]:= 0
FinProcedimiento
```
