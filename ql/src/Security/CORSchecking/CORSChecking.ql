/**
 * @name CORS checking
 * @description This rule is used to check your code whether config correctly about CORS(Cross-Origin Resource Sharing)
 * @kind cors
 */
import go
from Function f, DataFlow::CallNode c, string s
where f.hasQualifiedName("net/http.Header", "Set")
and 
    (
        (
            c = f.getACall()
            and c.getArgument(0).toString().matches("\"Access-Control-Allow-Origin\"")
            and c.getArgument(1).toString().matches("\"*\"")
            and s = "Critical! You set all origin can access,  it could be attakced by CORS"
        )
        or
        (
            // f.getACall().getArgument(0).toString().matches("\"Access-Control-Allow-Credentials\"")
            // and f.getACall().getArgument(1).toString().matches("\"true\"")
            // and c = f.getACall()
            c = f.getACall()
            and c.getArgument(0).toString().matches("\"Access-Control-Allow-Credentials\"")
            and c.getArgument(1).toString().matches("\"true\"")
            and s = "Warning! You set send credentials as ture. If your config Access-Control-Allow-Origin unproperly, it could be attakced by CORS"
            //c.getArgument(0).toString().matches("\"Access-Control-Allow-Origin\"")

        )
        or
        (
            c = f.getACall()
            and c.getArgument(0).toString().matches("\"Access-Control-Allow-Methods\"")
            and c.getArgument(1).toString().regexpMatch("(?i).*?(put|head|delete|options).*?")
            and s = "Info! Unexpected http method used, please check it"
        )

    )
select "Usage here: ", c, c.asExpr().getLocation(), s