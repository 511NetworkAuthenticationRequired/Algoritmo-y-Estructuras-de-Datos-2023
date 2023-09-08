# ALGORITMO Y ESTRUCTURAS DE DATOS 2023
Ejercicios en: [Guía de Trabajos Prácticos](https://aed-frre.github.io).

## SINTAXIS

### VARIABLES, CONSTANTES Y CONJUNTOS
```js
// VARIABLES
[Nombre de la variable]: entero
[Nombre de la variable]: alfanumerico
// CONSTANTES
[Nombre de la constante]:= [alfanumerico]
[Nombre de la constante]:= [entero]
// CONJUNTOS
[Nombre del conjunto]:= {[alfanumerico], [entero]}
```

### GENERAL

  #### CONDICIONALES
  ```js
  Si (condición) Entonces
    [proceso]
  Contrario/Sino
    [proceso]
  FinSi
  ```
  
  #### ITERACIONES
  ##### PRE-TEST
  ```js
  Mientras (condición) Hacer
   [proceso]
  FinMientras
  ```
  ##### POST-TEST
  ```js
  Repetir
    [proceso]
  Hasta que (condición)
  ```
  ##### MANEJADO POR CONTADOR
  ```js
  Para i:=x hasta y Hacer // "x" suele ser 1, "y" hace referencia al valor maximo de iteraciones
    [proceso]
  FinPara
  ```
  
  #### SUBACCIONES
  ##### FUNCION
  Retorna algo, ejemplo: cálculos matemáticos
  ```js
  Funcion Nombre ([Nombre del dato]: [Tipo de dato]): [Dato a devolver]
    [proceso]
  FinFuncion
  ```
  ##### PROCEDIMIENTO
  ```js
  Procedimiento Nombre ([parámetros]) Es
    [proceso]
  FinProcedimiento
  ```
  
### ESTRUCTURAS DE DATOS

  #### SECUENCIAS
  ```js
  [Nombre de la secuencia]: secuencia de [tipo de dato]
  [Nombre de la ventana]: [tipo de dato]

  Crear([Nombre de la secuencia])
  Abrir([Nombre de la secuencia])
  Cerrar([Nombre de la secuencia])
  Avanzar([Nombre de la secuencia], [Nombre de la ventana])
  Grabar([Nombre de la secuencia], [Nombre de la ventana])
  ```
  #### REGISTROS
  ```js
  [Nombre del registro] = registro
    [Nombre de la clave] = registro
      [claves]
    FinRegistro
    [campos]
  FinRegistro

  // Los campos y las claves van a compañadas de su tipo de dato (alfanumerico, numerico, booleano o conjunto) y de la cantidad, ejemplo: Clave: AN(2) --> "H1".
  ```
  #### ARCHIVOS 
  ```js
  [Nombre del archivo]: archivo de [Nombre del registro] ordenado por [Nombre de la clave]
  [Nombre del reg]: [Nombre del registro]

  Abrir([Nombre del archivo])
  Cerrar([Nombre del archivo])
  Leer([Nombre del archivo], [Nombre del reg])
  ```

### CORTE DE CONTROL
#### ESQUELETO
  ```js
  Inicializar()
  Mientras NFDA([archivo]) Hacer
    TratarCorte()
    TratarRegistro()
    LeerRegistro()
  FinMientras
  EmitirTotales()
  Corte3()
  EmitirTotales()
  Finalizar()
  ```
### MEZCLA

  #### CICLOS DE APAREO
  ##### INCLUYENTE
  ```JS
  Mientras ([clave1] <> HV) o ([clave2] <> HV) o ... ([claveN] <> HV) Hacer // HV hace referencia a "High Value".
    [proceso]
  FinMientras
  ```
  #### EXCLUYENTE
  ```JS
  Mientras NFDA ([archivo]) y NFDA ([archivo2]) Hacer
    [proceso común]
  FinMientras

  Mientras NFDA ([archivo]) Hacer
    [proceso archivo]       
  FinMientras
  Mientras NFDA ([archivo2]) Hacer
    [proceso archivo2]  
  FinMientras
  ```
  
  
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
#### Convertir caracter a entero
```js
Funcion ConvertirCAE([variable]: caracter) Es
			Segun [variable] Hacer
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
#### Avanzar manejado por contador
```js
x: entero
Procedimiento AvanzarN Es
			Para (i:=1 hasta x) Hacer
				Avanzar(secuencia, ventana)
			FinPara
		FinProcedimiento
```
