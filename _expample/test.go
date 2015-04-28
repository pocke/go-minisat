package main

import (
	"fmt"

	"github.com/pocke/go-minisat"
)

func main() {
	s := minisat.NewSolver()
	fmt.Println(s)
	s.NewVar()
	fmt.Println(s.Solve())
}
