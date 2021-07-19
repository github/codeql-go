/**
 * Provides classes and predicates used by the Cross Site Request Forgery
 * (CSRF) query.
 */

import go

/** Provides classes and predicates for the Cross Site Request Forgery (CSRF) query. */
module Csrf {
  /** A data flow source for Cross Site Request Forgery (CSRF) vulnerabilities. */
  abstract class Source extends DataFlow::Node { }

  /** A data flow sink for Cross Site Request Forgery (CSRF) vulnerabilities. */
  abstract class Sink extends DataFlow::Node { }

  /** A sanitizer for Cross Site Request Forgery (CSRF) vulnerabilities. */
  abstract class Sanitizer extends DataFlow::Node { }

  /** An outgoing sanitizer edge for Cross Site Request Forgery (CSRF) vulnerabilities. */
  abstract class SanitizerEdge extends DataFlow::Node { }

  /**
   * A sanitizer guard for Cross Site Request Forgery (CSRF) vulnerabilities.
   */
  abstract class SanitizerGuard extends DataFlow::BarrierGuard { }

  /**
   * A gorilla mux router creation expression, considered as a flow source for cross site request forgery (CSRF).
   */
  private class MuxRouterCreation extends Source {
    MuxRouterCreation() {
      exists(Method m |
        m.hasQualifiedName("gorilla/mux", "NewRouter") or
        m.hasQualifiedName("gorilla/mux", "Route", "SubRouter")
      |
        m.getACall() = this
      )
    }
  }

  /**
   * A handler creation expression, considered as a flow source for cross site request forgery (CSRF).
   */
  private class HandlerCreation extends Source {
    HandlerCreation() {
      exists(Method m | m.hasQualifiedName("net/http", ["Handle", "HandleFunc"]) |
        m.getACall() = this
      )
    }
  }

  /**
   * The handler of an HTTP request, viewed as a sink for Cross Site Request Forgery (CSRF) vulnerabilities.
   */
  private class HttpHandlerSink extends Sink {
    HttpHandlerSink() {
      exists(Method m, int i |
        m.hasQualifiedName("net/http", _) and
        m.getParameterType(i).hasQualifiedName("net/http", "Handler") and
        m.getParameter(i) = this.asParameter()
      )
    }
  }
}
