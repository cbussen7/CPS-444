package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	fmt.Println("Enter the subtotal and gratuity rate: ")

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
        scanner.Scan()
        line := scanner.Text()
        if len(line) == 0 {
            break
        }
        lines = append(lines, line)
    }

	subtotal, err1 := strconv.ParseFloat(lines[0], 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	gratuityRate, err2 := strconv.ParseFloat(lines[1], 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	gratuity := subtotal * (gratuityRate / 100.0)
	total := subtotal + gratuity

	fmt.Println("The gratuity is $", gratuity, "and total is $", total)
}