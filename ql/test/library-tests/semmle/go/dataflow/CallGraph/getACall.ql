import go

query predicate missingCall(DeclaredFunction f, DataFlow::CallNode call) {
  call.getACallee() = f.getFuncDecl() and
  not call = f.getACallIncludingExternals()
}

query predicate spuriousCall(DeclaredFunction f, DataFlow::CallNode call) {
  call = f.getACallIncludingExternals() and
  exists(FuncDecl fd | fd = f.getFuncDecl() | not call.getACallee() = fd)
}
