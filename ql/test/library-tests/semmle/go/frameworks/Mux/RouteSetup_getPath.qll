import go

query predicate test_RouteSetup_getPath(Mux::RouteSetup rs, DataFlow::Node res) {
    res = rs.getPath()
}