import go
import TestUtilities.InlineExpectationsTest
import experimental.frameworks.FastHTTP

class Configuration extends TaintTracking::Configuration {
  Configuration() { this = "test-configuration" }

  override predicate isSource(DataFlow::Node source) {
    exists(Function fn | fn.hasQualifiedName(_, "source") | source = fn.getACall().getResult())
  }

  override predicate isSink(DataFlow::Node sink) {
    exists(Function fn | fn.hasQualifiedName(_, "sink") | sink = fn.getACall().getAnArgument())
  }
}

class Link extends TaintTracking::FunctionModel {
  Link() { hasQualifiedName(_, "link") }

  override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
    inp.isParameter(0) and outp.isParameter(1)
  }
}

class TaintTrackingTest extends InlineExpectationsTest {
  TaintTrackingTest() { this = "TaintTrackingTest" }

  override string getARelevantTag() { result = "taintSink" }

  override predicate hasActualResult(string file, int line, string element, string tag, string value) {
    file.matches("%/TaintTracking.go") and
    tag = "taintSink" and
    exists(DataFlow::Node sink | any(Configuration c).hasFlow(_, sink) |
      element = sink.toString() and
      value = "" and
      sink.hasLocationInfo(file, line, _, _, _)
    )
  }
}
