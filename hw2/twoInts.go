package main

import "fmt"
import "os"
import "math"
import "strconv"

func main(){
	//declare input variables as ints
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	num1, err1 := strconv.Atoi(arg1)
	if err1 != nil{
		fmt.Println("Error")
		return
	}
	num2, err2 := strconv.Atoi(arg2)
	if err2 != nil{
		fmt.Println("Error")
		return
	}


	sum := num1 + num2
	fmt.Println("Sum is", sum)

	difference := num1 - num2
	fmt.Println("Difference is", difference)

	product := num1 * num2
	fmt.Println("Product is", product)

	var avg float64 = (float64(sum) / 2)
	fmt.Println("Average is", avg)

	distance := math.Abs(float64(difference))
	fmt.Println("Distance is", distance)

	max := math.Max(float64(num1), float64(num2))
	fmt.Println("Max is", max)

	min := math.Min(float64(num1), float64(num2))
	fmt.Println("Min is", min)
}
