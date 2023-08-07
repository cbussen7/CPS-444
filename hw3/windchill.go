package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "math"

func main(){
	fmt.Println("Enter the temperature in Fahrenheit between -58 and 41:")
	tempScanner := bufio.NewScanner(os.Stdin)
	tempScanner.Scan()
	text1 := tempScanner.Text()

	fmt.Println("Enter the wind speed in miles per hour:")
	windScanner := bufio.NewScanner(os.Stdin)
	windScanner.Scan()
	text2 := windScanner.Text()

	temp, err1 := strconv.ParseFloat(text1, 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	wind, err2 := strconv.ParseFloat(text2, 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	if (temp < -58 || temp > 41 || wind < 2){
		fmt.Println("Error: temp needs to be between -58 and 41 and wind speed needs to be 2 or higher")
		return
	}

	windchill := 35.74 + (0.6215 * temp) - (35.75 * math.Pow(wind, 0.16)) + (0.4275 * temp * math.Pow(wind, 0.16))

	fmt.Println("The wind chill index is", windchill)
}