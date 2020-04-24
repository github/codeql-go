import go

from BoolOps::BoolExpr a, BoolOps::BoolExpr b
where a.influences(b)
select a, b
