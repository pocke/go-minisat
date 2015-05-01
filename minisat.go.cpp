#include <minisat/core/Solver.h>
#include <minisat/core/SolverTypes.h>
#include "minisat.go.h"

using namespace Minisat;


extern "C" WrapSolver NewSolver() {
  return (WrapSolver) new Solver();
}

extern "C" WrapVar WrapSolverNewVar(WrapSolver slv) {
    return ((Solver*) slv)->newVar();
}

extern "C" int WrapSolverSolve(WrapSolver slv) {
  return ((Solver*) slv)->solve() ? 1 : 0;
}

extern "C" WrapLit WrapMkLit(WrapVar v, int sign) {
  return mkLit(v, bool(sign)).x;
}

extern "C" int WrapSolverAddClause(WrapSolver slv, WrapLit* lits, int len) {
 vec<Lit> vec_lits;
 
  for (int i = 0; i < len; i++) {
    Lit l;
    l.x = lits[i];
    vec_lits.push(l);
  }
  return ((Solver*) slv)->addClause_(vec_lits) ? 1 : 0;
}
