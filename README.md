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
  ##### FUNCION // Retorna algo, ejemplo: cálculos matemáticos
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
### MEZCLA

  #### CICLOS DE APAREO
  Siendo HV el "High Value".
  ##### INCLUYENTE
  ```JS
  Mientras ([clave1] <> HV) o ([clave2] <> HV) o ... ([claveN] <> HV) Hacer  
    [proceso]
  FinMientras
  ```
  #### EXCLUYENTE
  ```JS
  Mientras NFDA ([arch1]) y NFDA ([arch2]) Hacer
    [proceso común]
  FinMientras

  Mientras NFDA ([arch1]) Hacer
    [proceso arch1]       
  FinMientras
  Mientras NFDA ([arch2]) Hacer
    [proceso arch2]  
  FinMientras
  ```
  
  
### ESQUELETOS FRECUENTES

#### CORTE DE CONTROL
  ```js
    
  ```
#### ACTUALIZACIÓN
  ```js
  ```
