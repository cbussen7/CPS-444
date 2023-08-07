package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "math"

func main(){
	fmt.Println("Enter the radius and length of a cylinder: ")

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

	radius, err1 := strconv.ParseFloat(lines[0], 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	length, err2 := strconv.ParseFloat(lines[1], 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	area := radius * radius * math.Pi
	volume := area * length

	fmt.Println("The area is", area)
	fmt.Println("The volume is", volume)
}