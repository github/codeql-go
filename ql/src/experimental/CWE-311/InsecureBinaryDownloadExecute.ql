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

class FlowStringToHTTPCall extends TaintTracking::Configuration {
  FlowStringToHTTPCall() { this = "FlowStringToHTTPCall" }

  predicate isSource(DataFlow::Node source, StringLit val) { source.asExpr() = val }

  predicate isSink(DataFlow::Node sink, HTTP::ClientRequest httpClientCall) {
    sink = httpClientCall.(DataFlow::CallNode).getArgument(0)
  }

  override predicate isSource(DataFlow::Node source) { this.isSource(source, _) }

  override predicate isSink(DataFlow::Node sink) { this.isSink(sink, _) }
}

class FlowFromResponseBodyToAFile extends TaintTracking::Configuration {
  FlowFromResponseBodyToAFile() { this = "FlowFromResponseBodyToAFile" }

  predicate isSource(DataFlow::Node source, HTTP::ClientRequest httpClientCall) {
    exists(Read read, DataFlow::Node base, Field fld |
      fld.hasQualifiedName("net/http", "Response", "Body") and
      read = fld.getARead() and
      read.readsField(base, fld) and
      base.getAPredecessor*() = httpClientCall.(DataFlow::CallNode).getResult(0).getASuccessor*()
    |
      source = read
    )
  }

  predicate isSink(DataFlow::Node sink, ValueEntity val) {
    val.getType().(PointerType).getBaseType().hasQualifiedName("os", "File") and
    sink = val.getARead()
  }

  override predicate isSource(DataFlow::Node source) { isSource(source, _) }

  override predicate isSink(DataFlow::Node sink) { isSink(sink, _) }
}

predicate responseBodyFlowsToFile(HTTP::ClientRequest httpClientCall, ValueEntity file) {
  exists(FlowFromResponseBodyToAFile cfg, DataFlow::PathNode source, DataFlow::PathNode sink |
    cfg.hasFlowPath(source, sink) and
    cfg.isSource(source.getNode(), httpClientCall) and
    cfg.isSink(sink.getNode(), file)
  )
}

predicate getFileCreationCall(DataFlow::ValueEntity file, DataFlow::CallNode fileCreationCall) {
  file.getAWrite().getRhs() = fileCreationCall.getResult(0)
}

class FlowFromStringToFileCreatorCall extends TaintTracking::Configuration {
  FlowFromStringToFileCreatorCall() { this = "FlowFromStringToFileCreatorCall" }

  predicate isSource(DataFlow::Node source, ValueEntity val) { source = val.getARead() }

  predicate isSink(DataFlow::Node sink, DataFlow::CallNode fileCreationCall) {
    // TODO: the argument index might be different for each creation function.
    sink = fileCreationCall.getArgument(0)
  }

  override predicate isSource(DataFlow::Node source) { isSource(source, _) }

  override predicate isSink(DataFlow::Node sink) { isSink(sink, _) }
}

predicate getFilename(DataFlow::CallNode fileCreationCall, ValueEntity filename) {
  exists(FlowFromStringToFileCreatorCall cfg, DataFlow::PathNode source, DataFlow::PathNode sink |
    cfg.hasFlowPath(source, sink) and
    cfg.isSink(sink.getNode(), fileCreationCall)
  |
    cfg.isSource(source.getNode(), filename)
  )
}

class FlowToExecutor extends TaintTracking::Configuration {
  FlowToExecutor() { this = "FlowToExecutor" }

  predicate isSource(DataFlow::Node source, ValueEntity val) { source = val.getARead() }

  predicate isSink(DataFlow::Node sink, SystemCommandExecution exec) {
    sink = exec.getCommandName()
  }

  override predicate isSource(DataFlow::Node source) { isSource(source, _) }

  override predicate isSink(DataFlow::Node sink) { isSink(sink, _) }
}

predicate flowsToExecutor(ValueEntity filename) {
  exists(FlowToExecutor cfg, DataFlow::PathNode source, DataFlow::PathNode sink |
    cfg.hasFlowPath(source, sink) and
    cfg.isSource(source.getNode(), filename)
  )
}

from
  FlowStringToHTTPCall cfg, DataFlow::PathNode source, DataFlow::PathNode sink,
  HTTP::ClientRequest httpClientCall, ValueEntity file, DataFlow::CallNode fileCreationCall,
  ValueEntity filename, string url
where
  // A string flows to an HTTP client call:
  cfg.hasFlowPath(source, sink) and
  // Populate httpClientCall:
  cfg.isSink(sink.getNode(), httpClientCall) and
  // Check that the HTTP client call response body flows to a file:
  responseBodyFlowsToFile(httpClientCall, file) and
  //
  getFileCreationCall(file, fileCreationCall) and
  getFilename(fileCreationCall, filename) and
  flowsToExecutor(filename) and
  url = source.getNode().getStringValue() and
  url.matches("http:%")
select sink.getNode(), source, sink, "Downloading and executing " + url
