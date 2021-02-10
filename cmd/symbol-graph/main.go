/*
   Execution:    go run cmd/symbol-graph/main.go filename.txt delimiter
   Data files:   https://algs4.cs.princeton.edu/41graph/routes.txt
                 https://algs4.cs.princeton.edu/41graph/movies.txt
                 https://algs4.cs.princeton.edu/41graph/moviestiny.txt
                 https://algs4.cs.princeton.edu/41graph/moviesG.txt
                 https://algs4.cs.princeton.edu/41graph/moviestopGrossing.txt

   %  go run cmd/symbol-graph/main.go routes.txt " "
   JFK
      MCO
      ATL
      ORD
   LAX
      PHX
      LAS

   % go run cmd/symbol-graph/main.go movies.txt "/"
   Tin Men (1987)
      Hershey, Barbara
      Geppi, Cindy
      Jones, Kathy (II)
      Herr, Marcia
      ...
      Blumenfeld, Alan
      DeBoy, David
   Bacon, Kevin
      Woodsman, The (2004)
      Wild Things (1998)
      Where the Truth Lies (2005)
      Tremors (1990)
      ...
      Apollo 13 (1995)
      Animal House (1978)


   Assumes that input file is encoded using UTF-8.
   % iconv -f ISO-8859-1 -t UTF-8 movies-iso8859.txt > movies.txt
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/huggin/algo/algs4"
	"github.com/huggin/algo/stdin"
)

func main() {
	sg := algs4.NewSymbolGraph(os.Args[1], os.Args[2])
	g := sg.G()
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		source := strings.Trim(stdin.ReadString(), " ")
		if sg.Contains(source) {
			s := sg.Index(source)
			for _, v := range g.Adj(s) {
				fmt.Println("  ", sg.Name(v))
			}
			fmt.Println()
		} else {
			fmt.Println("input not contains source: ", source)
		}
	}
}
