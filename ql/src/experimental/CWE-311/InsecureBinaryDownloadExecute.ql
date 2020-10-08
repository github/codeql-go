/**
 * @name Download over plaint-http and execution of a binary file.
 * @description Downloading binary files over unencrypted connections and then executing them is dangerous.
 * @kind path-problem
 * @problem.severity error
 * @id go/dl-exec-over-http
 * @tags security
 *       external/cwe/cwe-311
 */

import go
import DataFlow::PathGraph
import semmle.go.security.CommandInjection

/**
 * A flow of a string to the URL parameter of a HTTP client request.
 */
class FlowHttpUrlToExecutor extends TaintTracking::Configuration {
  FlowHttpUrlToExecutor() { this = "FlowHttpUrlToExecutor" }

  predicate isSource(DataFlow::Node source, StringLit val) {
    source.asExpr() = val and
    val.getStringValue().matches("http:%")
  }

  predicate isSink(DataFlow::Node sink, SystemCommandExecution exec) {
    sink = exec.getCommandName()
  }

  override predicate isAdditionalTaintStep(DataFlow::Node input, DataFlow::Node output) {
    isUrlToRequestResponseStep(input, output)
  }

  override predicate isSource(DataFlow::Node source) { this.isSource(source, _) }

  override predicate isSink(DataFlow::Node sink) { this.isSink(sink, _) }
}

predicate isUrlToRequestResponseStep(DataFlow::Node url, DataFlow::Node response) {
  exists(HTTP::ClientRequest request | url = request.getUrl() and response = request.getResponse())
}

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
  cfg.isSource(source.getNode(), httpUrl)
select sink.getNode(), source, sink, "Executes an insecure download from $@", httpUrl, "here"
