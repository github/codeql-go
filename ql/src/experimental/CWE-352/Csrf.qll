/**
 * Provides a taint-tracking configuration for reasoning about Cross Site Request Forgery
 * (CSRF) vulnerabilities.
 *
 * Note, for performance reasons: only import this file if
 * `Csrf::Configuration` is needed, otherwise
 * `CsrfCustomizations` should be imported instead.
 */

import go

/**
 * Provides a taint-tracking configuration for reasoning about Cross Site Request Forgery
 * (CSRF) vulnerabilities.
 */
module Csrf {
  import CsrfCustomizations::Csrf

  /**
   * A taint-tracking configuration for reasoning about Cross Site Request Forgery
   * (CSRF) vulnerabilities.
   */
  class Configuration extends TaintTracking::Configuration {
    Configuration() { this = "Cross Site Request Forgery" }

    override predicate isSource(DataFlow::Node source) { source instanceof Source }

    override predicate isSink(DataFlow::Node sink) { sink instanceof Sink }

    override predicate isSanitizer(DataFlow::Node node) {
      super.isSanitizer(node) or
      node instanceof Sanitizer
    }

    override predicate isSanitizerOut(DataFlow::Node node) {
      super.isSanitizerOut(node) or
      node instanceof SanitizerEdge
    }

    override predicate isSanitizerGuard(DataFlow::BarrierGuard guard) {
      super.isSanitizerGuard(guard) or guard instanceof SanitizerGuard
    }
  }
}
