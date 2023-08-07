package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	fmt.Println("Enter speed and acceleration: ")

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

	speed, err1 := strconv.ParseFloat(lines[0], 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	acceleration, err2 := strconv.ParseFloat(lines[1], 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	runwayLength := (speed * speed) / (2 * acceleration)

	fmt.Println("The minimum runway length for this airplane is", runwayLength)
}