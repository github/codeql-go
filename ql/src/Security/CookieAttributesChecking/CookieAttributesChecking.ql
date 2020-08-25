/**
 * @name Cookie Attributes Checking.
 * @description This rule helps you to found out the Cookie set HttpOnly as True or not, and whether used HttpOnly, which decrease the damage of xss.
 * @kind xss
 */

import go
from Function f, StructLit s, int x, string infor, string tmp
where
    f.getName() = "SetCookie"
    and s.getType().toString() = "Cookie"
    and x in [1 .. 20]
    and 
        (
            s.getKey(x).toString() = "HttpOnly"
            and s.getAnElement().getChild(1).toString() = "false"
            and infor = "set Cookie attribute HttpOnly, but did not set it as True"
        )
        and tmp = ""
    or
        (
            x = 1
            and tmp = concat(int i| i in [1 .. s.getNumElement()]| s.getElement(i).getChild(0).toString(), "|")
            and not tmp.matches("%HttpOnly%")
            and infor = "did not set Cookie attribute HttpOnly"
        )
select "Cookie located here: ", s.getLocation(), " found this situation: ", infor