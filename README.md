# Algoritmo-y-Estructuras-de-Datos-2023
Ejercicios resueltos de AED (2023)

# SINTAXIS

## VARIABLES, CONSTANTES Y CONJUNTOS
```js
// VARIABLES
var: entero
var: alfanumerico
// CONSTANTES
cons:= [alfanumerico]
cons:= [entero]
// CONJUNTOS
con:= {[alfanumerico], [entero]}
```

## GENERAL

  ### CONDICIONALES
  ```js
  Si (condición) Entonces
    [proceso]
  Contrario/Sino
    [proceso]
  FinSi
  ```
  
  ### ITERACIONES
  #### PRE-TEST
  ```js
  Mientras (condición) Hacer
   [proceso]
  FinMientras
  ```
  #### POST-TEST
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
  
  ### SUBACCIONES
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
  
## ESTRUCTURAS DE DATOS

  ### SECUENCIAS
  ```js
  sec: secuencia de [tipo de dato]
  ven: [tipo de dato]

  Crear(sec)
  Abrir(sec)
  Cerrar(sec)
  Avanzar(sec, ven)
  Grabar(sec, ven)
  ```
  ### REGISTROS
  ```js
  reg = registro
    clave = registro
      [claves]
    FinRegistro
    [campos]
  FinRegistro

  // Los campos y las claves van a compañadas de su tipo de dato (alfanumerico, numerico, booleano o conjunto) y de la cantidad, ejemplo: Clave: AN(2) --> "H1".
  ```
  ### ARCHIVOS 
  ```js
  arch: archivo de [Nombre del registro] ordenado por [Nombre de la clave]
  reg: [Nombre del registro]
  ```
