/******************************************************************************
 *  Execution:    go run cmd/frequency-counter/main.go L < input.txt
*  Data files:   https://algs4.cs.princeton.edu/31elementary/tinyTale.txt
*                https://algs4.cs.princeton.edu/31elementary/tale.txt
*                https://algs4.cs.princeton.edu/31elementary/leipzig100K.txt
*                https://algs4.cs.princeton.edu/31elementary/leipzig300K.txt
*                https://algs4.cs.princeton.edu/31elementary/leipzig1M.txt
*
*  Read in a list of words from standard input and print out
*  the most frequently occurring word that has length greater than
*  a given threshold.
*
*  % go run cmd/frequency-counter/main.go 1 < tinyTale.txt
*  it 10
*
*  % go run cmd/frequency-counter/main.go 8 < tale.txt
*  business 122
*
*  % go run cmd/frequency-counter/main.go 10 < leipzig1M.txt
*  government 24763
*
*
******************************************************************************/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/huggin/algo/algs4"
	"github.com/huggin/algo/stdin"
)

func main() {
	minLen, _ := strconv.Atoi(os.Args[1])
	//st := algs4.NewSequentialSearchST()
	//st := algs4.NewBinarySearchST()
	//st := algs4.NewBST()
	// st := algs4.NewRedBlackBST()
	//st := algs4.NewSeparateChainHashST(0)
	st := algs4.NewLinearProbingHashST(0)
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		word := stdin.ReadString()
		if len(word) < minLen {
			continue
		}
		wordKey := algs4.StringHashKey(word)
		if !st.Contains(wordKey) {
			st.Put(wordKey, 1)
		} else {
			st.Put(wordKey, st.GetInt(wordKey)+1)
		}
	}

	max := algs4.StringHashKey("")
	st.Put(max, 0)
	for _, word := range st.Keys() {
		if st.GetInt(word.(algs4.StringHashKey)) > st.GetInt(max) {
			max = word.(algs4.StringHashKey)
		}
	}
	fmt.Printf("%s %d", max, st.GetInt(max))
}
