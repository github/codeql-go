import go

from BoolOps::BoolExpr a, boolean av, boolean wv, Write w
where
  a = any(Read r).asExpr() and
  a.ensuresIf(av, w.getRhs().asExpr(), wv)
select "If " + a + " is " + av + " then " + w.getLhs() + " is " + wv
