package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	fmt.Println("Enter weight in pounds:")
	weightScanner := bufio.NewScanner(os.Stdin)
	weightScanner.Scan()
	text1 := weightScanner.Text()

	fmt.Println("Enter height in inches:")
	heightScanner := bufio.NewScanner(os.Stdin)
	heightScanner.Scan()
	text2 := heightScanner.Text()

	weight, err1 := strconv.ParseFloat(text1, 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	height, err2 := strconv.ParseFloat(text2, 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	weightInKG := weight * 0.45359237
	heightInMeters := height * 0.0254

	bmi := weightInKG / (heightInMeters * heightInMeters)

	fmt.Println("BMI is", bmi)
}