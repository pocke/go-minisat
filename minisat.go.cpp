#include <minisat/core/Solver.h>
#include <minisat/core/SolverTypes.h>
#include "minisat.go.h"


extern "C" WrapSolver NewSolver() {
  return (WrapSolver) new Minisat::Solver();
}

extern "C" WrapVar WrapSolverNewVar(WrapSolver slv) {
    return ((Minisat::Solver*) slv)->newVar();
}

extern "C" int WrapSolverSolve(WrapSolver slv) {
  return ((Minisat::Solver*) slv)->solve() ? 1 : 0;
}

extern "C" WrapLit WrapMkLit(WrapVar v, int sign) {
  return Minisat::mkLit(v, bool(sign)).x;
}

extern "C" int WrapSolverAddClause(WrapSolver slv, WrapLit* lits, int len) {
  Minisat::vec<Minisat::Lit> vec_lits;
 
  for (int i = 0; i < len; i++) {
    Minisat::Lit l;
    l.x = lits[i];
    vec_lits.push(l);
  }
  return ((Minisat::Solver*) slv)->addClause_(vec_lits) ? 1 : 0;
}

extern "C" int WrapSolverModelValue(WrapSolver slv, WrapVar v) {
  Minisat::lbool b = ((Minisat::Solver*) slv)->model[v];
  return toInt(b);
}
