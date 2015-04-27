package minisat

/*
#cgo LDFLAGS: minisat/build/release/lib/libminisat.a
#include "minisat.go.h"
*/
import "C"

type Solver struct {
	CSolver *C.WrapSolver
}

func NewSolver() *Solver {
	s := C.NewSolver()
	return &Solver{
		CSolver: &s,
	}
}
