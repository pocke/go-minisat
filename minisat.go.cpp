#include "minisat/minisat/core/Solver.h"
#include "minisat.go.h"

using namespace Minisat;

// variable

// solver
WrapSolver NewSolver() {
  return (WrapSolver) new Solver();
}
