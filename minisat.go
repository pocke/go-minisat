package minisat

/*
#cgo CFLAGS: -D __STDC_LIMIT_MACROS -D __STDC_FORMAT_MACROS -D __STDC_CONSTANT_MACROS -U STDINT_H
#cgo LDFLAGS: -lminisat
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
	l := C.WrapMkLit(v, 0)
	return &Var{
		CVar: &v,
		CLit: &l,
	}
}

func (s *Solver) Solve() bool {
	res := C.WrapSolverSolve(*s.CSolver)
	return res != 0
}

func (s *Solver) AddClause(vars ...*Var) bool {
	lits := make([]C.WrapLit, 0, len(vars))
	for _, v := range vars {
		lits = append(lits, *v.CLit)
	}
	return C.WrapSolverAddClause(*s.CSolver, (*C.WrapLit)(&lits[0]), (C.int)(len(lits))) != 0
}

func (v *Var) Not() *Var {
	res := &Var{
		CVar: v.CVar,
		CLit: v.CLit,
	}

	lit := C.WrapMkLit(*v.CVar, 1)
	res.CLit = &lit

	return res
}
