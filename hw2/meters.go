package main

import "fmt"
import "os"
import "strconv"

func main(){
	arg1 := os.Args[1]

	metersNum, err := strconv.ParseFloat(arg1, 64)
	if err != nil{
		fmt.Println("Error")
		return
	}

	//create variables for the conversion factor of meters to miles, feet, and inches
	metersToMiles := 0.000621371
	metersToFeet := 3.28084
	metersToInches := 39.3701

	milesNum := metersNum * metersToMiles
	fmt.Println("Miles:", milesNum)

	feetNum := metersNum * metersToFeet
	fmt.Println("Feet:", feetNum)

	inchesNum := metersNum * metersToInches
	fmt.Println("Inches:", inchesNum)

}