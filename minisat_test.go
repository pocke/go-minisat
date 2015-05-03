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
	s.AddClause(v)
	s.AddClause(v.Not())
	if s.Solve() {
		t.Error("Expected: false, but got true")
	}
}

func TestNewVar(t *testing.T) {
	s := minisat.NewSolver()
	for i := 0; i < 3; i++ {
		v := s.NewVar()
		cvar := (int)(*v.CVar)
		clit := (int)(*v.CLit)
		if cvar*2 != clit {
			t.Errorf("clit expected: %d, but got %d", cvar*2, clit)
		}
	}
}

func TestVarNot(t *testing.T) {
	s := minisat.NewSolver()
	for i := 0; i < 3; i++ {
		v := s.NewVar().Not()
		cvar := (int)(*v.CVar)
		clit := (int)(*v.CLit)
		if cvar*2+1 != clit {
			t.Errorf("clit expected: %d, but got %d", cvar*2+1, clit)
		}
	}
}

func TestSolverModelValueWhenNotSolvedYet(t *testing.T) {
	s := minisat.NewSolver()
	v := s.NewVar()
	s.AddClause(v)
	s.AddClause(v.Not())
	_, err := s.ModelValue(v)
	if err == nil {
		t.Error("Expected: error, but got nil")
	}
}

func TestSolverModelValueWhenUnsat(t *testing.T) {
	s := minisat.NewSolver()
	v := s.NewVar()
	s.AddClause(v)
	s.AddClause(v.Not())
	s.Solve()
	_, err := s.ModelValue(v)
	if err == nil {
		t.Error("Expected: error, but got nil")
	}
}

func TestSolverModelValueWhenSat(t *testing.T) {
	s := minisat.NewSolver()
	v1 := s.NewVar()
	v2 := s.NewVar()

	s.AddClause(v1, v2)
	s.AddClause(v1, v2.Not())
	s.AddClause(v1.Not(), v2.Not())
	s.Solve()

	a1, err := s.ModelValue(v1)
	if err != nil {
		t.Errorf("Expected: nil, but got %s", err)
	}
	a2, err := s.ModelValue(v2)
	if err != nil {
		t.Errorf("Expected: nil, but got %s", err)
	}

	if !(a1 && !a2) {
		t.Errorf("Expected: a1 == true and a2 == false, but got a1 == %v, a2 == %v", a1, a2)
	}
}
