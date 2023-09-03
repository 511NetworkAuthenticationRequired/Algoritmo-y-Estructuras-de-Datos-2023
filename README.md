# Algoritmo-y-Estructuras-de-Datos-2023
Ejercicios resueltos de AED (2023)

# SINTAXIS

  ### CONDICIONALES
  ```
  Si (condición) Entonces
    [proceso]
  Contrario/Sino
    [proceso]
  FinSi
  ```
  
  ### ITERACIONES
    #### PRE-TEST
    ```
    Mientras (condición) Hacer
      [proceso]
    FinMientras
    ```
    #### POST-TEST
    ```
    Repetir
      [proceso]
    Hasta que (condición)
    ```
    #### MANEJADO POR CONTADOR
    ```
    Para i:=x hasta y Hacer // "x" suele ser 1, "y" hace referencia al valor maximo de iteraciones
      [proceso]
    FinPara
    ```
  
  ### SUBACCIONES
    #### FUNCION // Retorna algo, ejemplo: cálculos matemáticos
    ```
    Funcion Nombre ([Nombre del dato]: [Tipo de dato]): [Dato a devolver]
      [proceso]
    FinFuncion
    ```
    #### PROCEDIMIENTO
    ```
    Procedimiento Nombre ([parámetros]) Es
      [proceso]
    FinProcedimiento
    ```
  
## ESTRUCTURAS DE DATOS

  ### SECUENCIAS
  ```
  secuencia: secuencia de [tipo de dato]
  ventana: [tipo de dato]

  Crear(secuencia)
  Abrir(secuencia)
  Cerrar(secuencia)
  Avanzar(secuencia, ventana)
  Grabar(secuencia, ventana)
  ```

  ### REGISTROS
  ```
  reg = registro
    clave = registro
      [claves]
    FinRegistro
    [campos]
  FinRegistro

  // Los campos y las claves van a compañadas de su tipo de dato (alfanumerico, numerico, booleano o conjunto) y de la cantidad, ejemplo: Clave: AN(2) --> "H1".
  ```
