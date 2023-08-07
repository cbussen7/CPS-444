package main

import "fmt"
import "os"
import "strconv"

func main(){
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]
	
	gallonsInTank, err1 := strconv.ParseFloat(arg1, 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	fuelEfficiency, err2 := strconv.ParseFloat(arg2, 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	pricePerGallon, err3 := strconv.ParseFloat(arg3, 64)
	if err3 != nil{
		fmt.Println("Error")
		return
	}

	costPer100 := (100 / fuelEfficiency) * pricePerGallon

	distance := gallonsInTank * fuelEfficiency

	fmt.Println("The cost per 100 miles is $", costPer100)
	fmt.Println("Miles with remaining gas in tank:", distance)
	
}