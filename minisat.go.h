#ifdef __cplusplus
extern "C" {
#endif

  typedef int WrapVar;
  typedef void* WrapSolver;

  WrapVar WrapSolverNewVar(WrapSolver slv);

  WrapSolver NewSolver();

#ifdef __cplusplus
}
#endif
