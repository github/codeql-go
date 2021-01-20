import go
import DataFlow::PathGraph

class JWTWeakSecretKeyVerificationSink extends DataFlow::Node {
  JWTWeakSecretKeyVerificationSink() {
    exists(DataFlow::CallNode call |
      call.getTarget().getPackage().getPath() = "github.com/dgrijalva/jwt-go" and
      this.(DataFlow::MethodCallNode).getCalleeName() = "SignedString"
    )
  }
}

class JWTWeakSecretKeyVerificationSource extends DataFlow::Node {
  JWTWeakSecretKeyVerificationSource() {
    this.(DataFlow::MethodCallNode).getArgument(0).getAPredecessor*().getStringValue().length() = 0
  }
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
