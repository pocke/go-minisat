#include "minisat/minisat/core/Solver.h"
#include "minisat.go.h"

using namespace Minisat;

extern "C" WrapVar WrapSolverNewVar(WrapSolver slv) {
    return ((Solver*) slv)->newVar();
}

extern "C" WrapSolver NewSolver() {
  return (WrapSolver) new Solver();
}
