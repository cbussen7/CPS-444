package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "math"

func main(){
	fmt.Println("Enter the length of the side: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	length, err := strconv.ParseFloat(text, 64)
	if err != nil{
		fmt.Println("Error")
		return
	}

	area := (3 * math.Sqrt(3) / 2.0) * math.Pow(length, 2)

	fmt.Println("The area of the hexagon is", area)
}