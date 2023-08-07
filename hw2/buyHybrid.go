package main

import "fmt"
import "os"
import "strconv"

func main(){
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]
	arg4 := os.Args[4]
	arg5 := os.Args[5]
	
	costOfNew, err1 := strconv.ParseFloat(arg1, 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	milesPerYear, err2 := strconv.ParseFloat(arg2, 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	gasPrice, err3 := strconv.ParseFloat(arg3, 64)
	if err3 != nil{
		fmt.Println("Error")
		return
	}

	milesPerGallon, err4 := strconv.ParseFloat(arg4, 64)
	if err4 != nil{
		fmt.Println("Error")
		return
	}

	resale, err5 := strconv.ParseFloat(arg5, 64)
	if err5 != nil{
		fmt.Println("Error")
		return
	}

	yearsBeforeResale := 5.
	gasInYear := milesPerYear / milesPerGallon

	// cost to own = cost of new car plus 
	costToOwn := costOfNew + (yearsBeforeResale * gasInYear * gasPrice)
	fmt.Println("Cost to own: $", costToOwn)

	costAfterResale := costToOwn - resale
	fmt.Println("Cost after resale: $", costAfterResale)
}