import go

query predicate test_RouteSetup_getHandler(Mux::RouteSetup rs, DataFlow::FunctionNode res) {
    res = rs.getRouteHandler()
}