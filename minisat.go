package minisat

/*
#cgo LDFLAGS: minisat/build/release/lib/libminisat.a
#include "minisat.go.h"
*/
import "C"

type Solver struct {
	CSolver *C.WrapSolver
}

type Var struct {
	CVar *C.WrapVar
}

func NewSolver() *Solver {
	s := C.NewSolver()
	return &Solver{
		CSolver: &s,
	}
}

func (s *Solver) NewVar() *Var {
	v := C.WrapSolverNewVar(*s.CSolver)
	return &Var{
		CVar: v,
	}
}
