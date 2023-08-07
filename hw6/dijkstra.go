// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 99.

// Graph shows how to use a map of maps to represent a directed graph.
package main

import "fmt"
import "math"
import "os"
import "strconv"
import "bufio"
import "strings"

//!+
var graph [8][8]int

func addEdge(from int, to int, weight int) {
	graph[from][to] = weight
}

// func hasEdge(from, to int) bool {
// 	return graph[from][to]
// }

//!-

func minVertex(dist [8]int, spt [8]bool) int{
	min := math.MaxInt
	minIndex := -1

	for i:=0; i < 8; i++{
		if spt[i] == false && dist[i] <= min{
			min = dist[i]
			minIndex = i
		}
	}

	return minIndex
}

func dijkstra(graph[8][8]int, start int){
	// verts := 8
	var dist [8]int

	var spt [8]bool

	for i:=0; i < 8; i++{
		dist[i] = math.MaxInt
		spt[i] = false
	}

	dist[start] = 0

	for count:=1; count < 8-1; count++{
		u := minVertex(dist, spt)

		spt[u] = true

		for v:=0; v < 8; v++{
			if !spt[v] && graph[u][v] != 0 && dist[u] != math.MaxInt && dist[u] + graph[u][v] < dist[v]{
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	display(dist, start)
}

func display(dist[8] int, start int){
	for i :=1; i < 8; i++{
		fmt.Println("Shortest distance from", start, "to", i, ":", dist[i])
	}
	
}

func displayAllNodes(graph[8][8]int){
	for i :=1; i < 8; i++{
		fmt.Print(i, ":",)
		for j :=1; j < 8; j++{
			if graph[i][j] > 0{
				fmt.Print(j, "[", graph[i][j], "] ")
			}
		}

		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: please provide a file to read the digraph from.\n")
        os.Exit(1)
	}
	fileName := os.Args[1]

	//read from file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// scan from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// split each line and then take 1st element is start node, 2nd as edge, 3rd as weight
		splitLines := strings.Split(line, " ")

		nodeStr := splitLines[0]
		edgeStr := splitLines[1]
		weightStr := splitLines[2]

		// convert each value for node, edge, and weight to ints
		node, err1 := strconv.Atoi(nodeStr)
		if err1 != nil{
			fmt.Println("Error")
			return
		}
		edge, err2 := strconv.Atoi(edgeStr)
		if err2 != nil{
			fmt.Println("Error")
			return
		}
		weight, err3 := strconv.Atoi(weightStr)
		if err3 != nil{
			fmt.Println("Error")
			return
		}

		// add each node to graph
		addEdge(node, edge, weight)
	}

	displayAllNodes(graph)

	fmt.Println("Start node? ")
	nodeScanner := bufio.NewScanner(os.Stdin)
	nodeScanner.Scan()
	input := nodeScanner.Text()

	startNode, err4 := strconv.Atoi(input)
	if err4 != nil{
		fmt.Println("Error")
		return
	}

	dijkstra(graph, startNode)

}