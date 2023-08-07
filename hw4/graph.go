// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 99.

// Graph shows how to use a map of maps to represent a directed graph.
package main

import "fmt"

//!+
var graph [26][26]bool

func addEdge(from, to int) {
	graph[from][to] = true
}

func hasEdge(from, to int) bool {
	return graph[from][to]
}

//!-

func main() {
	addEdge(1, 2)
	addEdge(3, 4)
	addEdge(1, 4)
	addEdge(4, 1)
	addEdge(24, 25)
	fmt.Println(hasEdge(1, 2))
	fmt.Println(hasEdge(3, 4))
	fmt.Println(hasEdge(1, 4))
	fmt.Println(hasEdge(4, 1))
	fmt.Println(hasEdge(24, 2))
	fmt.Println(hasEdge(3, 4))
	fmt.Println(hasEdge(24, 4))
	fmt.Println(hasEdge(4, 24))
	fmt.Println(hasEdge(24, 25))

}