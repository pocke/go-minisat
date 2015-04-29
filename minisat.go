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
	CLit *C.WrapLit
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
		CVar: &v,
	}
}

func (s *Solver) Solve() bool {
	res := C.WrapSolverSolve(*s.CSolver)
	return res != 0
}

func (s *Solver) AddClause(lits ...Var) {

}

func (v *Var) Not() *Var {
	res := &Var{
		CVar: v.CVar,
		CLit: v.CLit,
	}
	if res.CLit == nil {
		lit := C.WrapMkLit(*v.CVar, 1)
		res.CLit = &lit
	} else {

	}

	return res
}
