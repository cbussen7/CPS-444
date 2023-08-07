package main

import "fmt"
import "os"
import "strconv"
import "math"

func main(){
	arg1 := os.Args[1]

	radius, err := strconv.ParseFloat(arg1, 64)
	if err != nil{
		fmt.Println("Error")
		return
	}

	area := (math.Pi) * math.Pow(radius, 2)
	fmt.Println("Area of circle is", area)

	circumference := 2 * (math.Pi) * radius
	fmt.Println("Circumference of circle is", circumference)

	volume := (4.0/3.0) * (math.Pi) * math.Pow(radius, 3)
	fmt.Println("Volume of sphere is", volume)

	surfaceArea := 4 * (math.Pi) * math.Pow(radius, 2)
	fmt.Println("Surface area of sphere is", surfaceArea)
}