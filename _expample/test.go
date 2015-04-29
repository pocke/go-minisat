package main

import (
	"fmt"

	"github.com/pocke/go-minisat"
)
import "C"

func main() {
	s := minisat.NewSolver()
	v := s.NewVar()
	fmt.Println((C.int)(*v.CVar))
	v2 := s.NewVar()
	fmt.Println((C.int)(*v2.CVar))
}
