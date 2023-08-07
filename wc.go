package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var fileName string
    var option string
    if len(os.Args) > 1 {
        // if first arg starts w - store as option
        //then check if there's a second arg, if not exit
        if strings.HasPrefix(os.Args[1], "-") {
            option = os.Args[1]
            if len(os.Args) > 2 {
                fileName = os.Args[2]
            } else {
                fmt.Println("Please pass a file name")
                os.Exit(1)
            }
        } else {
            fileName = os.Args[1]
        }
        
    } else {
        fmt.Println("Please pass a file name")
        os.Exit(1)
    }
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Err ", err)
    }
    scanner := bufio.NewScanner(file)
    lines, words, characters := 0, 0, 0
    for scanner.Scan() {
        lines++

        line := scanner.Text()
        characters += len(line)

        splitLines := strings.Split(line, " ")
        words += len(splitLines)
    }
    if len(os.Args) > 1{
        fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
    } else if len(os.Args) > 2{
        if option == "-l"{
            fmt.Printf("%8d %s\n", lines, fileName)
        } else if option == "-w"{
            fmt.Printf("%8d %s\n", words, fileName)
        } else if option == "-m"{
            fmt.Printf("%8d %s\n", characters, fileName, option)
        } else if option == "-lw" || option == "-wl"{
            fmt.Printf("%8d%8d %s\n", lines, words, fileName, option)
        } else if option == "-lm" || option == "-ml"{
            fmt.Printf("%8d%8d %s\n", lines, characters, fileName, option)
        } else if option == "-wm" || option == "-mw"{
            fmt.Printf("%8d%8d %s\n", words, characters, fileName, option)
        } else if option == "-lwm" || option == "-lmw" || option == "-wlm" || option == "-wml" || option == "-mwl" || option == "-mlw"{
            fmt.Printf("%8d%8d %s\n", lines, words, characters, fileName, option)
        } else{
            fmt.Println("Invalid option given.")
            os.Exit(1)
        }
    }
}