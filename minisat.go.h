#ifdef __cplusplus
extern "C" {
#endif

  typedef int WrapVar;
  typedef void* WrapSolver;
  typedef int WrapLit;


  WrapSolver NewSolver(double seed);

  WrapVar WrapSolverNewVar(WrapSolver slv);
  int WrapSolverSolve(WrapSolver slv);
  int WrapSolverAddClause(WrapSolver slv, WrapLit* lits, int len);
  int WrapSolverModelValue(WrapSolver slv, WrapVar v);
  void WrapSolverFree(WrapSolver slv);
  WrapLit WrapMkLit(WrapVar v, int sign);

#ifdef __cplusplus
}
#endif
