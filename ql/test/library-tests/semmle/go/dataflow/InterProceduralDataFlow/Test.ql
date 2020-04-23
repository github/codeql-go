import go

class IsSafeBarrier extends DataFlow::BarrierGuard, DataFlow::CallNode {
  IsSafeBarrier() { this.getCalleeName() = "isSafe" }

  override predicate checks(Expr e, boolean outcome) {
    e = this.getArgument(0).asExpr() and outcome = true
  }
}

class MyConfiguration extends DataFlow::Configuration {
  MyConfiguration() { this = "MyConfiguration" }

  override predicate isSource(DataFlow::Node nd) {
    exists(ValueEntity v, Write w |
      v.getName().matches("source%") and
      w.writes(v, nd)
    )
    or
    nd.(DataFlow::CallNode).getCalleeName() = "source"
  }

  override predicate isSink(DataFlow::Node nd) {
    exists(ValueEntity v, Write w |
      v.getName().matches("sink%") and
      w.writes(v, nd)
    )
    or
    exists(DataFlow::CallNode call | call.getCalleeName() = "sink" and nd = call.getAnArgument())
  }

  override predicate isBarrierGuard(DataFlow::BarrierGuard guard) { guard instanceof IsSafeBarrier }
}

from MyConfiguration cfg, DataFlow::Node source, DataFlow::Node sink
where cfg.hasFlow(source, sink)
select source, sink
