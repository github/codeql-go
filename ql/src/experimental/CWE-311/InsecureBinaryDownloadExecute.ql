/**
 * @name Download over plain HTTP and execution of a binary file.
 * @description Downloading binary files over unencrypted connections and then executing them is dangerous.
 * @kind path-problem
 * @problem.severity error
 * @id go/dl-exec-over-http
 * @tags security
 *       external/cwe/cwe-311
 */

import go
import semmle.go.security.CommandInjection
import DataFlow::PathGraph

private predicate isHttpUrlSource(DataFlow::Node source, StringLit val) {
  source.asExpr() = val and
  val.getStringValue().matches("http:%")
}

/**
 * A flow of a string beginning with `http:` to the URL parameter of a HTTP client request.
 */
class FlowHttpUrlToClientRequest extends TaintTracking::Configuration {
  FlowHttpUrlToClientRequest() { this = "FlowHttpUrlToClientRequest" }

  override predicate isSource(DataFlow::Node source) {
    exists(StringLit val | isHttpUrlSource(source, val))
  }

  override predicate isSink(DataFlow::Node sink) {
    exists(HTTP::ClientRequest clientRequest | sink = clientRequest.getUrl())
  }
}

/**
 * Holds if `source` is an HTTP URL that is used to make an HTTP client request.
 */
predicate urlFlowsToClientRequest(DataFlow::Node source) {
  exists(FlowHttpUrlToClientRequest config | config.hasFlow(source, _))
}

/**
 * A flow of an HTTP URL to an executor.
 */
class FlowHttpUrlToExecutor extends TaintTracking::Configuration {
  FlowHttpUrlToExecutor() { this = "FlowHttpUrlToExecutor" }

  predicate isSource(DataFlow::Node source, StringLit val) { isHttpUrlSource(source, val) }

  override predicate isAdditionalTaintStep(DataFlow::Node url, DataFlow::Node response) {
    exists(HTTP::ClientRequest request |
      url = request.getUrl() and response = request.getResponse()
    )
  }

  override predicate isSource(DataFlow::Node source) { this.isSource(source, _) }

  override predicate isSink(DataFlow::Node sink) {
    exists(SystemCommandExecution exec | sink = exec.getCommandName())
  }
}

/**
 * Model of `os.Create`, propagating taint from the name of a created file to the created `os.File` instance.
 */
class OsCreateModel extends DataFlow::FunctionModel {
  OsCreateModel() { hasQualifiedName("os", "Create") }

  override predicate hasDataFlow(FunctionInput input, FunctionOutput output) {
    input.isResult(0) and output.isParameter(0)
  }
}

from
  FlowHttpUrlToExecutor cfg, DataFlow::PathNode source, DataFlow::PathNode sink, StringLit httpUrl
where
  cfg.hasFlowPath(source, sink) and
  cfg.isSource(source.getNode(), httpUrl) and
  // Cross-check: make sure an HTTP Client Request is involved in the path somewhere:
  urlFlowsToClientRequest(source.getNode())
select sink.getNode(), source, sink, "Executes an insecure download from $@", httpUrl, "here"
