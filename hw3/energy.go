package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	fmt.Println("Enter the amount of water in kilograms:")
	waterScanner := bufio.NewScanner(os.Stdin)
	waterScanner.Scan()
	text1 := waterScanner.Text()

	fmt.Println("Enter the initial temperature:")
	temp1Scanner := bufio.NewScanner(os.Stdin)
	temp1Scanner.Scan()
	text2 := temp1Scanner.Text()

	fmt.Println("Enter the final temperature:")
	temp2Scanner := bufio.NewScanner(os.Stdin)
	temp2Scanner.Scan()
	text3 := temp2Scanner.Text()

	waterKG, err1 := strconv.ParseFloat(text1, 64)
	if err1 != nil{
		fmt.Println("Error")
		return
	}

	initTemp, err2 := strconv.ParseFloat(text2, 64)
	if err2 != nil{
		fmt.Println("Error")
		return
	}

	finTemp, err3 := strconv.ParseFloat(text3, 64)
	if err3 != nil{
		fmt.Println("Error")
		return
	}

	energy := waterKG * (finTemp - initTemp) * 4184

	fmt.Printf("The energy needed is %f\n", energy)
}