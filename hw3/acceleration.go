package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	fmt.Println("Enter v0, v1, and t: ")

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

	v0, err1 := strconv.ParseFloat(lines[0], 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	v1, err2 := strconv.ParseFloat(lines[1], 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	t, err3 := strconv.ParseFloat(lines[2], 64)
	if err3 != nil{
		fmt.Println("Error")
		return
	}

	acceleration := (v1 - v0) / t

	fmt.Println("The average acceleration is", acceleration)
}