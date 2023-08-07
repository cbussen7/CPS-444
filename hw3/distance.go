package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "math"

func main(){
	fmt.Println("Enter x1 and y1: ")

	var lines []string
	firstScanner := bufio.NewScanner(os.Stdin)
	
	for {
        firstScanner.Scan()
        line := firstScanner.Text()
        if len(line) == 0 {
            break
        }
        lines = append(lines, line)
    }

	x1, err1 := strconv.ParseFloat(lines[0], 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	y1, err2 := strconv.ParseFloat(lines[1], 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	fmt.Println("Enter x2 and y2: ")

	var lines2 []string
	secondScanner := bufio.NewScanner(os.Stdin)
	
	for {
        secondScanner.Scan()
        line2 := secondScanner.Text()
        if len(line2) == 0 {
            break
        }
        lines2 = append(lines2, line2)
    }

	x2, err3 := strconv.ParseFloat(lines2[0], 64)
	if err3 != nil{
		fmt.Println("Error")
		return
	}

	y2, err4 := strconv.ParseFloat(lines2[1], 64)
	if err4 != nil{
		fmt.Println("Error")
		return
	}

	xDif := x2 - x1
	xDifSquared := math.Pow(xDif, 2)

	yDif := y2 - y1
	yDifSquared := math.Pow(yDif, 2)

	sumDif := xDifSquared + yDifSquared
	
	distance := math.Sqrt(sumDif)

	fmt.Println("The distance between the two points is", distance)
}