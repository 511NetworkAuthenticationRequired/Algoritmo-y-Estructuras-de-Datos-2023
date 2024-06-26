# ALGORITMO Y ESTRUCTURAS DE DATOS (2023)
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

### DATOS SIMPLES
 - **Caracter**: ```"h", "1", " ", "&"```
 - **Alfanumerico/AN**: ```"hola", "  ", "122", "!#", "Es todo un tema."``` 
 - **Entero/N**: ```1, 7, 22, 127```
 - **Booleano/Lógico**:``` V/F, S/N, 0/1```
 - **Rango/Rankeado**: ```([Valor númerico]...[Valor númerico])``` --> ```v:= (1..10) // Toma valores del 1 al 10```
   
 **IMPORTANTE:** 1 <> "1", el primero es un entero, el segundo un caracter/alfanumerico.

#### CONSTANTES
 Su valor no cambia en el proceso.
 ```js
	nombre_constante = valor

  // Ejemplo: pi = 3,14159265359
 ```
 #### VARIABLES
 Se le puede asignar nuevos valores.
 ```js
  nombre_variable: dato
 ```
 La asignación se especifica colocando ":=".
 ```js
  nombre_variable:= valor

	// Ejemplo:
		nombre_variable: entero; nombre_variable:= 1
		nombre_variable₂: alfanumerico; nombre_variable₂:= "Hola mundo"
  ```
 #### CONJUNTOS
 ```js
  nombre_conjunto: {valor₁, valor₂, valor₃, ..., valorᵢ} 
 ```
 ```js
	nombre_conjunto = {valor₁, valor₂, valor₃, ..., valorᵢ}
 ```
 #### PUNTEROS & NODOS/CELDAS
 Los punteros hacen referencia a una dirección de memoria de una variable dinámica.
 ```js
 [Nombre del puntero]: puntero a [Nombre del nodo/celda]
 ```
 Los nodos o celdas son campos a los que solo puede acceder mediante un puntero.
 ```js
 nombre_nodo = registro
   Dato: [Información que se va almacenar]
   Prox: puntero a nombre_nodo
 FinRegistro
 ```
### ESTRUCTURAS DE DATOS
  #### SECUENCIAS
  ```js
  nombre_secuencia: secuencia de [Tipo de dato]
  nombre_ventana: [Tipo de dato]

  Crear(nombre_secuencia)
  Abrir(nombre_secuencia)
  Cerrar(nombre_secuencia)
  Avanzar(nombre_secuencia, nombre_ventana)
  Grabar(nombre_secuencia, nombre_ventana)
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
  [Nombre del archivo]: archivo de [Nombre del registro] ordenado/indexado por [Nombre de la clave] // Si esta 	 
  ordenado o indexado por alguna clave
  [Nombre del archivo]: archivo de [Nombre del registro]  // Sino posee campo clave
  [Nombre del registro]: [Nombre del registro]
  ```
   ##### SECUENCIALES
   ```js
    AbrirE/S([Nombre del archivo])
    AbrirE/([Nombre del archivo])
    AbrirS/([Nombre del archivo])
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
Los índices generalmente reciben los nombres de "i", "j" y "k" o bien "x", "y" y "z", aunque al ser variables pueden tomar el nombre que se le quiera dar.
  ```js
  [Nombre del vector]: arreglo de [Valor del dato....Valor del dato] de [Tipo de dato]
  // Acesso al elemento
  [Nombre del vector][[Número del índice]] // Ejemplo: vector[1]
  // Asignación de valores
  [Nombre del vector][[Número del índice]]:= [Valor del dato] // Ejemplo: vector[1]:= 7
  ```
  #### MATRICES
  Son arreglos bidimensionales o de mayor dimensión; es decir, poseen dos subíndices o más (fila y columna, etc).
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
  CorteMayorNivel()
  EmitirTotales()
  Finalizar()
  ```
**NOTA:** Revisar también este [ejemplo](https://github.com/511NetworkAuthenticationRequired/Algoritmo-y-Estructuras-de-Datos-2023/tree/main#corte-de-control-1).

### MEZCLA/APAREO
  
  #### CICLOS DE APAREO
  ##### INCLUYENTE
  ```JS
  Mientras ([archivo1.clave] <> HV) o ([archivo2.clave] <> HV) o ... ([claveN] <> HV) Hacer // HV hace referencia a "High Value".
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
