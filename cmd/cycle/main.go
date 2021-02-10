/******************************************************************************
 *  Execution:    go run cmd/cycle/main.go filename.txt
 *  Data files:   https://algs4.cs.princeton.edu/41graph/tinyG.txt
 *                https://algs4.cs.princeton.edu/41graph/mediumG.txt
 *                https://algs4.cs.princeton.edu/41graph/largeG.txt
 *
 *  Identifies a cycle.
 *  Runs in O(E + V) time.
 *
 *  % go run cmd/cycle/main.go tinyG.txt
 *  3 4 5 3
 *
 *  % go run cmd/cycle/main.go mediumG.txt
 *  15 0 225 15
 *
 *  % go run cmd/cycle/main.go largeG.txt
 *  996673 762 840164 4619 785187 194717 996673
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/huggin/algo/algs4"
	"github.com/huggin/algo/stdin"
)

func main() {
	graph := algs4.NewGraph(stdin.NewIn(os.Args[1]))
	finder := algs4.NewCycle(graph)
	if finder.HasCycle() {
		for _, v := range finder.Cycle().Slice() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Graph is acyclic")
	}

}
