#include "minisat/minisat/core/Solver.h"
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
