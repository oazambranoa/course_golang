package main

import (
	pk "course_golang/src/mypackage"
	"fmt"
)

/*
type car struct {
	brand string
	year int
} */

/*
func normalFunction(message string) {
	fmt.Println(message)
}

func tirpeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleRetunr(a int) (c, d int) {
	return a, a * 2
} */

func main() {
	fmt.Println("Hola")
	/*
	normalFunction("Hola mundo")
	tirpeArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	//	value1, value2 := doubleRetunr(2)
	//	fmt.Println("Value 1 y 2:", value1, value2)

	value1, _ := doubleRetunr(2)
	fmt.Println("Value 1 :", value1)

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

	//FOr condicional

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	
		// For forever
		counterForever := 0

		for {
			fmt.Println(counterForever)
			counterForever++
		} 

	// Llave valor
	m := make(map[string]int)

	m["Omar"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// ENcontrar un valor

	value3, ok := m["Josep"]
	fmt.Println(value3, ok) 

	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)*/

	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage()


}
