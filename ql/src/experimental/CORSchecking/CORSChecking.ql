/**
 * @name CORS checking
 * @description This rule is used to check your code whether config correctly about CORS(Cross-Origin Resource Sharing)
 * @kind problem
 * @problem.severity warning
 */
import go
from Function f, DataFlow::CallNode c, string s
where f.hasQualifiedName("net/http.Header", "Set")
and 
    (
        (
            c = f.getACall()
            and c.getArgument(0).getStringValue() = "Access-Control-Allow-Origin"
            and c.getArgument(1).getStringValue() = "*"
            and s = "Allowing access from arbitrary origins may facilitate CORS attacks."
        )
        or
        (
            c = f.getACall()
            and c.getArgument(0).getStringValue() = "Access-Control-Allow-Credentials"
            and c.getArgument(1).getStringValue() = "true"
            and s = "Allowing Access-Control-Allow-Credentials as true would expand you attack surface if CORS attacks appears。"

        )
        or
        (
            c = f.getACall()
            and c.getArgument(0).getStringValue() = "Access-Control-Allow-Methods"
            and c.getArgument(1).getStringValue().regexpMatch("(?i).*?(put|head|delete|options).*?")
            and s = "Allowing HTTP methods such as PUT and DELETE may be dangerous."
        )

    )
select c, s