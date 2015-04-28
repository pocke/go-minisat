#ifdef __cplusplus
extern "C" {
#endif

  typedef int WrapVar;
  typedef void* WrapSolver;


  WrapSolver NewSolver();

  WrapVar WrapSolverNewVar(WrapSolver slv);
  int WrapSolverSolve(WrapSolver slv);

#ifdef __cplusplus
}
#endif
