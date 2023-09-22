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
|----------|------------------------------------|-------------|
| =        | Igual a                            | Relacional  |
| <>       | Distinto de                        | Relacional  |
| >        | Mayor que                          | Relacional  |
| >=       | Mayor o Igual que                  | Relacional  |
| <        | Menor que                          | Relacional  |
| <=       | Menor o Igual que                  | Relacional  |
|----------|------------------------------------|-------------|
| -        | Negación (No)                      | Lógico      |
| ^        | Conjunción (Y)                     | Lógico      |
| ˅        | Disyunción (O)                     | Lógico      |
|----------|------------------------------------|-------------|
```
### ESTRUCTURAS DE DATOS
 #### DATOS SIMPLES
 - **Caracter**: ```"h", "1", " ", "&"```
 - **Alfanumerico/AN**: ```"Hola", "  ", "122", "!#"```
 - **Entero/N**: ```1, 7, 22, 127```
 - **Booleano/Lógico**:``` V/F, S/N, 0/1```
 - **Rango/Rankeado**: ```([Valor númerico]...[Valor númerico])``` --> ```v:= (1..10) // Toma valores del 1 al 10```
   
 **IMPORTANTE**: 1 <> "1", el primero es un entero, el segundo un caracter/alfanumerico.
 
 #### VARIABLES
 Se le puede asignar nuevos valores.
 ```js
 [Nombre de la variable]: [Tipo de dato]
 // Asignación
 [Nombre de la variable]:= [Valor]
  ```
 #### CONSTANTES
 Su valor no cambia en el proceso.
 ```js
 [Nombre de la constante]:= [Valor del tipo de dato]
 ```
 #### CONJUNTOS
 ```js
 [Nombre del conjunto]:= {[Valor del dato], ..., [Valor del dato]}
 ```
 #### PUNTEROS & NODOS/CELDAS
 Los punteros hacen referencia a una dirección de memoria de una variable dinámica.
 ```js
 [Nombre del puntero]: puntero a [Nombre del nodo/celda]
 ```
 Los nodos o celdas son campos a los que solo puede acceder mediante un puntero.
 ```js
 [Nombre del nodo/celda] = registro
  [Dato]: [es la información que se va almacenar]
  [Prox]: puntero a [Nombre del nodo/celda]
 FinRegistro
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

  // Asignación
  [Nombre del registro].[Nombre del campo]:= [Valor]
  [Nombre del registro].[Nombre de la clave].[Nombre del campo]:= [Valor]

  // Los campos y las claves van a compañadas de su tipo de dato (alfanumerico, numerico, etc) y de la cantidad, ejemplo: Clave: AN(10) --> "Hello7world".
  ```
  #### ARCHIVOS 
  ```js
  [Nombre del archivo]: archivo de [Nombre del registro] ordenado/indexado por [Nombre de la clave]
  [Nombre del registro]: [Nombre del registro]
  ```
   ##### SECUENCIALES
   ```js
    AbrirE/S([Nombre del archivo])
    Cerrar([Nombre del archivo])
    Grabar([Nombre del archivo], [Nombre del registro]) // Escribir()
    Leer([Nombre del archivo], [Nombre del registro])
  ```
  ##### INDEXADOS
  A los procedimientos de secuenciales se le agregan:
  ```js
  ReGrabar([Nombre del archivo], [Nombre del registro) // ReEscribir()
  Borrar([Nombre del archivo], [Nombre del registro)
  ```
  #### ARREGLOS
  ```js
  [Nombre del vector]: arreglo de [Valor del dato....Valor del dato] de [Tipo de dato]
  // Acesso al elemento
  [Nombre del vector][[Número del índice]] // Ejemplo: vector[1]
  // Asignación de valores
  [Nombre del vector][[Número del índice]]:= [Valor del dato] // Ejemplo: vector[1]:= 7
  ```
  #### MATRICES
  Son arreglos bidimensionales; es decir, poseen dos subíndices (fila y columna).
  ```js
  [Nombre de la matriz]:= arreglo de [Valor del dato....Valor del dato, Valor del dato....Valor del dato] de [Tipo de dato]
  ```
  LISTAS
  SIMPLEMENTE ENLAZADAS
  ```js
  
  ```
  DOBLEMENTE ENLAZADAS
  ```js
  
  ```
  CIRCULARES
  ```js
  
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
  **NOTA:** Revisar también este [ejemplo](https://github.com/511NetworkAuthenticationRequired/Algoritmo-y-Estructuras-de-Datos-2023/blob/main/README.md#convertir-caracter-a-entero).
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
  #### METODOS DE BUSQUEDA
  ##### LINEAL
  ```js
 
  ```
  ##### LINEAL CON CENTINELA
  ```js
 
  ```
  ##### BINARIA/DICOTOMICA
  ```js
 
  ```
  #### METODOS DE ORDENAMIENTO
  ##### INSERCION DIRECTA
  ```js
 
  ```
  ##### SELECCION DIRECTA
  ```js
 
  ```
 ##### INTERCAMBIO DIRECTA/BURBUJA
  ```js
 
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
**NOTA:** Revisar también este [ejemplo](https://github.com/511NetworkAuthenticationRequired/Algoritmo-y-Estructuras-de-Datos-2023/tree/main#corte-de-control-1).

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
// Algunos la llaman Convertir_A_Caracter(), y según el enunciado inclusive no es preciso crearla y se la puede usar directamente en el proceso.
```
#### Avanzar manejado por contador
```js
Procedimiento AvanzarN (x: entero) Es
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
