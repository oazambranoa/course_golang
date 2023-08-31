package main

import "fmt"

func main() {
	// Declaracion de variables

	hellMessage := "hello"
	worldMessage := "World"

	// Println
	fmt.Println(hellMessage, worldMessage)
	fmt.Println(hellMessage, worldMessage)

	//Printf
	nombre := "Omar"
	edad := 30
	fmt.Printf("%s tiene mas de %d 29 años\n", nombre, edad)
	fmt.Printf("%v tiene mas de %v 29 años\n", nombre, edad)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d 29 años", nombre, edad)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("hellMessage: %T", hellMessage)
	fmt.Printf("Años: %T", edad)
	

	// Declaracion de constantes

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos
	// -Rectangulo, trapecio y de un circulo

}
