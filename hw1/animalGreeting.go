package main

import "fmt"
import "os"

func main(){
	name := os.Args[1]
	nameLength := len(name)
	//declare variable for dashes at top of greeting and initialize using for loop
	var topOfGreeting string
	for i:=0; i<nameLength; i++{
		topOfGreeting = topOfGreeting + "-"
	}

	//variable spaces is used to add proper spacing between hello and \ depending on name
	var spacesTop string
	for i:=0; i<nameLength-5; i++{
		spacesTop = spacesTop + " "
	}

	var spacesBottom string
	for i:=0; i<nameLength; i++{
		spacesBottom = spacesBottom + " "
	}
	fmt.Println(
		" /\\___/\\     " + topOfGreeting + "  \n" +
		"(  ' '  )  / Hello" + spacesTop + "\\ \n" + 
		"(   -   ) <  " + name + " | \n" + 
		" |  |  |   \\" + spacesBottom +      " / \n" + 
		"(___|___)    " + topOfGreeting + "    ")
}