package minisat_test

import (
	"testing"

	"github.com/pocke/go-minisat"
)

func TestSolveSat(t *testing.T) {
	s := minisat.NewSolver()
	v1 := s.NewVar()
	v2 := s.NewVar()
	v3 := s.NewVar()

	nv1 := v1.Not()
	nv2 := v2.Not()
	nv3 := v3.Not()

	s.AddClause(v1, v2, v3)
	s.AddClause(nv1, nv2)
	s.AddClause(nv1, nv3)
	s.AddClause(nv2, nv3)

	if !s.Solve() {
		t.Error("Expected: true, but got false")
	}
}

func TestSolveUnsat(t *testing.T) {
	s := minisat.NewSolver()
	v := s.NewVar()
	s.AddClause(v, v.Not())
	if s.Solve() {
		t.Error("Expected: false, but got true")
	}
}
