/**
 * Provides a taint-tracking configuration for reasoning about
 * unvalidated URL redirection problems on the server side.
 *
 * Note, for performance reasons: only import this file if
 * `OpenUrlRedirect::Configuration` is needed, otherwise
 * `OpenUrlRedirectCustomizations` should be imported instead.
 */

import go
import UrlConcatenation

module OpenUrlRedirect {
  import OpenUrlRedirectCustomizations::OpenUrlRedirect

  /**
   * A data-flow configuration for reasoning about unvalidated URL redirections.
   */
  class Configuration extends DataFlow::Configuration {
    Configuration() { this = "OpenUrlRedirect" }

    override predicate isSource(DataFlow::Node source) { source instanceof Source }

    override predicate isSink(DataFlow::Node sink) { sink instanceof Sink }

    override predicate isBarrier(DataFlow::Node node) { node instanceof Barrier }

    override predicate isAdditionalFlowStep(DataFlow::Node pred, DataFlow::Node succ) {
      // A write to URL.Host
      exists(Write write, Field f, DataFlow::SsaNode var |
        write.writesField(var.getAUse(), f, pred) and
        succ = var.getAUse() and
        write.getASuccessor+() = succ.asInstruction() and
        f.getName() = "Host" and
        var.getType().hasQualifiedName("net/url", "URL")
      )
      or
      // taint steps that do not include flow through fields
      TaintTracking::localTaintStep(pred, succ) and not TaintTracking::fieldReadStep(pred, succ)
    }

    override predicate isBarrierOut(DataFlow::Node node) { hostnameSanitizingPrefixEdge(node, _) }

    override predicate isBarrierGuard(DataFlow::BarrierGuard guard) {
      guard instanceof BarrierGuard
    }
  }

  /**
   * A data-flow configuration for reasoning about safe URLs for unvalidated URL redirections.
   */
  class SafeUrlConfiguration extends DataFlow::Configuration {
    SafeUrlConfiguration() { this = "SafeUrlFlow" }

    override predicate isSource(DataFlow::Node source) {
      source.(DataFlow::FieldReadNode).getField().hasQualifiedName("net/http", "Request", "URL")
      or
      source.(DataFlow::FieldReadNode).getField().hasQualifiedName("net/http", "Request", "Host")
    }

    override predicate isSink(DataFlow::Node sink) { sink instanceof Sink }

    override predicate isAdditionalFlowStep(DataFlow::Node pred, DataFlow::Node succ) {
      TaintTracking::functionModelStep(any(SafeUrlMethod m), pred, succ)
      or
      TaintTracking::referenceStep(pred, succ)
      or
      TaintTracking::stringConcatStep(pred, succ)
      or
      TaintTracking::tupleStep(pred, succ)
      or
      exists(DataFlow::FieldReadNode frn | succ = frn |
        frn.getBase() = pred and frn.getFieldName() = "Host"
      )
    }

    override predicate isBarrierOut(DataFlow::Node node) {
      exists(Write w, Field f | f.hasQualifiedName("net/url", "URL", "Path") |
        w.writesField(node.getASuccessor*(), f, _)
      )
    }
  }
}
