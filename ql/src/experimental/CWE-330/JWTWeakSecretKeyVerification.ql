import go
import DataFlow::PathGraph

class JWTWeakSecretKeyVerificationSink extends DataFlow::Node {
  JWTWeakSecretKeyVerificationSink() {
    exists(DataFlow::MethodCallNode call |
      call.getTarget().hasQualifiedName("github.com/dgrijalva/jwt-go", "Token", "SignedString")
    |
      this = call.getArgument(0)
    )
  }
}

class JWTWeakSecretKeyVerificationSource extends DataFlow::Node {
  JWTWeakSecretKeyVerificationSource() { this.getStringValue().length() = 0 }
}

class JWTWeakSecretKeyVerificationConfiguration extends TaintTracking::Configuration {
  JWTWeakSecretKeyVerificationConfiguration() { this = "JWTWeakSecretKeyVerificationConfiguration" }

  override predicate isSource(DataFlow::Node source) {
    source instanceof JWTWeakSecretKeyVerificationSource
  }

  override predicate isSink(DataFlow::Node sink) {
    sink instanceof JWTWeakSecretKeyVerificationSink
  }
}

from
  JWTWeakSecretKeyVerificationConfiguration config, DataFlow::PathNode source,
  DataFlow::PathNode sink
where config.hasFlowPath(source, sink)
select source
