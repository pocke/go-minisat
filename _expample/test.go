package main

import (
	"fmt"

	"github.com/pocke/go-minisat"
)

func main() {
	s := minisat.NewSolver()
	fmt.Println(s)
	v := s.NewVar()
	fmt.Println(v)
	fmt.Println(v.Not())
	fmt.Println(s.Solve())
}
