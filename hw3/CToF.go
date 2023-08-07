package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	fmt.Println("Enter a degree in Celsius: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	degreesC, err := strconv.ParseFloat(text, 64)
	if err != nil{
		fmt.Println("Error")
		return
	}

	fahrenheit := (1.8 * degreesC) + 32

	fmt.Println(degreesC, "Celsius is", fahrenheit, "Fahrenheit")

}