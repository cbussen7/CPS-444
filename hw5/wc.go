package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//if slice contains more than one element check if it contains each letter - 
//if only one element check that element

func contains(s []string, str string) bool{
    for _, i:= range s{
        if i == str{
            return true
        }
    }
    return false
}

func main() {
    var fileName string
    var option []string
    var lines, words, characters int
    //if both option and file are given
    if len(os.Args) > 2 {
        for i := range os.Args{
            option = append(option, os.Args[i])
        } 
        
        //if last arg is not an option, make it filename and truncate options slice
        if !strings.HasPrefix(option[len(option)-1], "-"){
            fileName = os.Args[len(os.Args)-1]
            option = option[1:len(option)-1]

            //read from file
            file, err := os.Open(fileName)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error: %s\n", err)
                os.Exit(1)
            }
            scanner := bufio.NewScanner(file)
            lines, words, characters = 0, 0, 0
            for scanner.Scan() {
                lines++

                line := scanner.Text()
                characters += len(line)

                splitLines := strings.Split(line, " ")
                words += len(splitLines)
            }
        }else{
            //last arg is option, read from stdin
            //read from stdin
            scanner := bufio.NewScanner(os.Stdin)
            lines, words, characters = 0, 0, 0
            for scanner.Scan() {
                lines++

                line := scanner.Text()
                characters += len(line)

                splitLines := strings.Split(line, " ")
                words += len(splitLines)
            }
        }
    

        if len(option) == 1{
            //check which option was passed - fmt.Printf writes to stdout
            if option[0] == "-l"{
                fmt.Printf("%8d %s\n", lines, fileName)
            } else if option[0] == "-w"{
                fmt.Printf("%8d %s\n", words, fileName)
            } else if option[0] == "-m"{
                fmt.Printf("%8d %s\n", characters, fileName)
            } else if option[0] == "-lw" || option[0] == "-wl"{
                fmt.Printf("%8d%8d %s\n", lines, words, fileName)
            } else if option[0] == "-lm" || option[0] == "-ml"{
                fmt.Printf("%8d%8d %s\n", lines, characters, fileName)
            } else if option[0] == "-wm" || option[0] == "-mw"{
                fmt.Printf("%8d%8d %s\n", words, characters, fileName)
            } else if option[0] == "-lwm" || option[0] == "-lmw" || option[0] == "-wlm" || option[0] == "-wml" || option[0] == "-mwl" || option[0] == "-mlw"{
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            } else{
                fmt.Fprintf(os.Stderr, "Invalid option given\n")
                os.Exit(1)
            }
        } else if len(option) > 1{
            if contains(option, "-l") && contains(option, "-w") && contains(option, "-m"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-lw") && contains(option, "-m"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-l") && contains(option, "-wm"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-lm") && contains(option, "-w"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-wl") && contains(option, "-m"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-l") && contains(option, "-mw"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-ml") && contains(option, "-w"){
                fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
            }else if contains(option, "-l") && contains(option, "-w"){
                fmt.Printf("%8d%8d %s\n", lines, words, fileName)
            }else if contains(option, "-l") && contains(option, "-m"){
                fmt.Printf("%8d%8d %s\n", lines, characters, fileName)
            }else if contains(option, "-w") && contains(option, "-m"){
                fmt.Printf("%8d%8d %s\n", words, characters, fileName)
            }else if contains(option, "-l"){
                fmt.Printf("%8d %s\n", lines, fileName)
            }else if contains(option, "-w"){
                fmt.Printf("%8d %s\n", words, fileName)
            }else if contains(option, "-m"){
                fmt.Printf("%8d %s\n", characters, fileName)
            }
        }
        

        
      // if one of option or file are given  
    } else if len(os.Args) == 2{
        //option only was passed
        if strings.HasPrefix(os.Args[1], "-") {
            option = append(option, os.Args[1])

            //read from stdin
            scanner := bufio.NewScanner(os.Stdin)
            lines, words, characters := 0, 0, 0
            for scanner.Scan() {
                lines++

                line := scanner.Text()
                characters += len(line)

                splitLines := strings.Split(line, " ")
                words += len(splitLines)
            }

            //check which option was passed - fmt.Printf writes to stdout
            if option[0] == "-l"{
                fmt.Printf("%8d\n", lines)
            } else if option[0] == "-w"{
                fmt.Printf("%8d\n", words)
            } else if option[0] == "-m"{
                fmt.Printf("%8d\n", characters)
            } else if option[0] == "-lw" || option[0] == "-wl"{
                fmt.Printf("%8d%8d\n", lines, words)
            } else if option[0] == "-lm" || option[0] == "-ml"{
                fmt.Printf("%8d%8d\n", lines, characters)
            } else if option[0] == "-wm" || option[0] == "-mw"{
                fmt.Printf("%8d%8d\n", words, characters)
            } else if option[0] == "-lwm" || option[0] == "-lmw" || option[0] == "-wlm" || option[0] == "-wml" || option[0] == "-mwl" || option[0] == "-mlw"{
                fmt.Printf("%8d%8d%8d\n", lines, words, characters)
            } else{
                fmt.Fprintf(os.Stderr, "Invalid option given\n")
                os.Exit(1)
            }

        //file only was passed
        } else {
            fileName = os.Args[1]

            //read file 
            file, err := os.Open(fileName)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error:%s\n", err)
                os.Exit(1)
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

            //fmt.Printf writes to stdout
            fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)
        }
    
    //neither option nor file was passed
    } else {
        //read from stdin
        scanner := bufio.NewScanner(os.Stdin)
        lines, words, characters := 0, 0, 0
        for scanner.Scan() {
            lines++

            line := scanner.Text()
            characters += len(line)

            splitLines := strings.Split(line, " ")
            words += len(splitLines)
        }

        //fmt.Printf writes to stdout
        fmt.Printf("%8d%8d%8d\n", lines, words, characters)

    }
}