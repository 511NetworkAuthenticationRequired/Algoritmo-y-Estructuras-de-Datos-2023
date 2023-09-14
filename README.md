# ALGORITMO Y ESTRUCTURAS DE DATOS 2023
Ejercicios en: [Guía de Trabajos Prácticos](https://aed-frre.github.io).

## SINTAXIS

### OPERADORES
```yaml
|----------|------------------------------------|-------------|
| Operador | Operación                          | Tipo        |
|----------|------------------------------------|-------------|
| +        | Suma                               | Aritmético  |
| -        | Resta                              | Aritmético  |
| *        | Multiplicación                     | Aritmético  |
| /        | División Real                      | Aritmético  |
| DIV      | División Entera                    | Aritmético  |
| MOD      | Módulo o Resto de la División      | Aritmético  |
| ^ o **   | Exponenciación                     | Aritmético  |
| ABSO     | Valor Absoluto                     | Aritmético  |
| TRUNC    | Truncado (Parte Entera)            | Aritmético  |
| REDOND   | Redondeo                           | Aritmético  |
| ==       | Igual a                            | Relacional  |
| <>       | No Igual a                         | Relacional  |
| >        | Mayor que                          | Relacional  |
| >=       | Mayor o Igual que                  | Relacional  |
| <        | Menor que                          | Relacional  |
| <=       | Menor o Igual que                  | Relacional  |
| -        | Negación (No)                      | Lógico      |
| ^        | Conjunción (Y)                     | Lógico      |
| ˅        | Disyunción (O)                     | Lógico      |
|----------|------------------------------------|-------------|
```
### ESTRUCTURAS DE DATOS
 #### DATOS SIMPLES
 Es importante conocer los siguientes tipos de datos simples con sus respectivos ejemplos:
 
 - **Caracter**: ```"h", "1", " ", "&"```
 - **Alfanumerico/AN**: ```"Hola", "  ", "122", "!#"```
 - **Entero/N**: ```1, 7, 22, 127```
 - **Booleano/Lógico**:``` V/F, S/N, 0/1```
 - **Rango/Rankeado**: ```([Valor númerico]...[Valor númerico])``` --> ```v:= (1..10) // Toma valores del 1 al 10```
   
 **IMPORTANTE**: 1 <> "1", el primero es un entero, el segundo un caracter/alfanumerico.
 
 #### VARIABLES
 ```js
 [Nombre de la variable]: [Tipo de dato]
  ```
 #### CONSTANTES
 ```js
 [Nombre de la constante]:= [Valor del tipo de dato]
 ```
 #### CONJUNTOS
 ```js
 [Nombre del conjunto]:= {[Valor de dato], ..., [Valor de dato]}
 ```
  #### SECUENCIAS
  ```js
  [Nombre de la secuencia]: secuencia de [Tipo de dato]
  [Nombre de la ventana]: [Tipo de dato]

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
      [Claves]
    FinRegistro
    [Campos]
  FinRegistro

  // Los campos y las claves van a compañadas de su tipo de dato (alfanumerico, numerico, etc) y de la cantidad, ejemplo: Clave: AN(10) --> "Hello7world".
  ```
  #### ARCHIVOS 
  ```js
  [Nombre del archivo]: archivo de [Nombre del registro] ordenado por [Nombre de la clave]
  [Nombre del registro]: [Nombre del registro]

  Abrir([Nombre del archivo])
  Cerrar([Nombre del archivo])
  Leer([Nombre del archivo], [Nombre del registro])
  ```
  #### ARREGLOS
  ```js
  [Nombre del vector]:= arreglo de [Valor de Dato....Valor de Dato] de [Tipo de dato]
  // Acesso al elemento
  [Nombre del vector][Número del índice] // Ejemplo: vector[1]
  // Asignación de valores
	[Nombre del vector]:= [Valor del dato] // Ejemplo: vector[1]:= 7
  
  ```

### GENERAL

  #### CONDICIONALES
  ##### SI & CONTRARIO
  ```js
  Si (condición) Entonces
    [proceso]
  Contrario/Sino
    [proceso]
  FinSi
  ```
  ##### SEGUN
  ```js
  Segun [variable] Entonces
    [Proceso]
  FinSegun
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
  Para i:=x hasta y Hacer // "x" suele ser 1, "y" hace referencia al valor maximo de iteraciones.
    [proceso]
  FinPara
  ```
  
  #### SUBACCIONES
  ##### FUNCION
  Retorna algo, ejemplo: cálculos matemáticos.
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
 **IMPORTANTE:** Nunca usar ninguna de las palabras reservadas.
 
### CORTE DE CONTROL
#### ESQUELETO
  ```js
  Inicializar()
  Mientras NFDA ([archivo]) Hacer // NFDA significa "No fin de archivo".
    TratarCorte()
    TratarRegistro()
    LeerRegistro()
  FinMientras
  EmitirTotales()
  Corte3()
  EmitirTotales()
  Finalizar()
  ```
**IMPORTANTE:** Revisar también este [ejemplo](https://github.com/511NetworkAuthenticationRequired/Algoritmo-y-Estructuras-de-Datos-2023/edit/main/README.md#constantes).
### MEZCLA/APAREO

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
Funcion ConvertirCAE(variable: caracter) Es
	Segun variable Hacer
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
// Algunos la llaman Convertir_A_Caracter(), y según el enunciado inclusive no es preciso crearla y se la puede usar directamente en el proceso.
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
#### Contadores a 0
Subaccion que pone todos los contadores a 0. No se requiere crear la subaccion y se la puede usar directamente en el proceso.
```js
Contador_A_Cero()
```
