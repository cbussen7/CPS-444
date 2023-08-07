package main

import "fmt"
import "os"

func main(){
	name := os.Args[1]
	nameLength := len(name)
	//fmt.Println(name)
	//fmt.Println(nameLength)
	var edgeOfBox string //:= "-" * nameLength

	//use for loop to construct top and bottom of box
	for i:=0; i<nameLength+2; i++{
		edgeOfBox = edgeOfBox + "-"
	}
	
	fmt.Println("+" + edgeOfBox + "+")
	fmt.Println("| " + name + " |")
	fmt.Println("+" + edgeOfBox + "+")
}