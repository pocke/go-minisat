package minisat

/*
#cgo LDFLAGS: -lminisat
#include "minisat.go.h"
*/
import "C"
import "fmt"

type lBool int

const (
	l_True lBool = iota
	l_False
	l_Undef
)

type solverState int

const (
	sat solverState = iota
	unsat
	notSolved
)

type Solver struct {
	CSolver *C.WrapSolver
	state   solverState
}

type Var struct {
	CVar *C.WrapVar
	CLit *C.WrapLit
}

func NewSolver() *Solver {
	s := C.NewSolver()
	return &Solver{
		CSolver: &s,
		state:   notSolved,
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
	switch res {
	case 0:
		s.state = unsat
	case 1:
		s.state = sat
	}
	return res != 0
}

func (s *Solver) AddClause(vars ...*Var) bool {
	lits := make([]C.WrapLit, 0, len(vars))
	for _, v := range vars {
		lits = append(lits, *v.CLit)
	}
	return C.WrapSolverAddClause(*s.CSolver, (*C.WrapLit)(&lits[0]), (C.int)(len(lits))) != 0
}

// TODO: error message
func (s *Solver) ModelValue(v *Var) (bool, error) {
	switch s.state {
	case unsat:
		return false, fmt.Errorf("Unsat")
	case notSolved:
		return false, fmt.Errorf("not solved yet")
	}
	lb := (lBool)(C.WrapSolverModelValue(*s.CSolver, *v.CVar))
	if lb == l_Undef {
		return false, fmt.Errorf("undefined")
	}
	return lb == l_True, nil
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

func (s *Solver) Free() {
	C.WrapSolverFree(*s.CSolver)
}
