/**
 * @name 'HttpOnly' attribute is not set to true
 * @description Omitting the 'HttpOnly' attribute for security sensitive data allows
 *              malicious JavaScript to steal it in case of XSS vulnerability. Always set
 *              'HttpOnly' to 'true' to authentication related cookie to make it
 *              not accessible by JavaScript.
 * @kind problem
 * @problem.severity warning
 * @precision high
 * @id go/cookie-httponly-not-set
 * @tags security
 *       external/cwe/cwe-1004
 */

import go
import AuthCookie

predicate isNetHttpCookieFlow(Expr expr) {
  exists(
    HttpOnlyCookieTrackingConfiguration httpOnlyCfg, AuthCookieNameConfiguration cookieNameCfg,
    SetCookieSink sink, DataFlow::Node source
  |
    httpOnlyCfg.hasFlow(source, sink) and
    cookieNameCfg.hasFlow(source, sink) and
    sink.asExpr() = expr
  )
}

predicate isGinContextCookieFlow(Expr expr) {
  exists(CallExpr c |
    c.getTarget().getQualifiedName() = "github.com/gin-gonic/gin.Context.SetCookie" and
    c.getArgument(6) = expr and
    exists(DataFlow::Node valSrc, DataFlow::Node httpOnlyArg |
      DataFlow::localFlow(valSrc, httpOnlyArg) and
      httpOnlyArg.asExpr() = c.getArgument(6) and
      valSrc.asExpr().getBoolValue() = false
    ) and
    exists(DataFlow::Node nameSrc, DataFlow::Node nameArg |
      DataFlow::localFlow(nameSrc, nameArg) and
      nameArg.asExpr() = c.getArgument(0) and
      isAuthVariable(nameSrc.asExpr())
    )
  )
}

predicate isGorillaSessionsCookieFlow(Expr expr) {
  exists(DataFlow::Node sessionSave |
    sessionSave.asExpr() = expr and
    exists(CookieStoreSaveTrackingConfiguration cfg | cfg.hasFlow(_, sessionSave)) and
    (
      not exists(SessionOptionsTrackingConfiguration cfg | cfg.hasFlow(_, sessionSave))
      or
      exists(SessionOptionsTrackingConfiguration cfg, DataFlow::Node options |
        cfg.hasFlow(options, sessionSave) and
        (
          not exists(DataFlow::Node rhs | rhs = getValueForFieldWrite(options.asExpr(), "HttpOnly"))
          or
          exists(DataFlow::Node rhs |
            rhs = getValueForFieldWrite(options.asExpr(), "HttpOnly") and
            exists(DataFlow::Node valSrc |
              DataFlow::localFlow(valSrc, rhs) and
              valSrc.asExpr().getBoolValue() = false
            )
          )
        )
      )
    )
  )
}

from Expr expr
where
  isNetHttpCookieFlow(expr) or
  isGinContextCookieFlow(expr) or
  isGorillaSessionsCookieFlow(expr)
select expr, "Cookie attribute 'HttpOnly' is not set to true."
