/**
 * @name Cookie Attributes Checking.
 * @description This rule helps you to found out the Cookie set HttpOnly as True or not, and whether used HttpOnly, which decrease the damage of xss.
 * @kind problem
 * @problem.severity warning
 */

import go

from StructLit s, KeyValueExpr k, string infor
where
  s.getType().hasQualifiedName("net/http", "Cookie") and
  (
    k = s.getAnElement() and
    k.getKey().(Ident).getName() = "HttpOnly" and
    k.getValue().(Ident).getName() = "false" and
    infor = "You set Cookie attribute HttpOnly, but did not set it as True"
    or
    not exists(KeyValueExpr k2 |
      k2 = s.getAnElement() and
      k2.getKey().(Ident).getName() = "HttpOnly"
    ) and
    infor = "You did not use Cookie attribute HttpOnly"
  )
select s, infor
