/**
 * Provides classes for working with concepts relating to the Macaron web framework
 */

import go

private module Macaron {
  private class Context extends HTTP::ResponseWriter::Range {
    Context() {
      this
          .getType()
          .getMethod("Redirect")
          .hasQualifiedName("gopkg.in/macaron.v1", "Context", "Redirect")
    }
  }

  private class RedirectCall extends HTTP::Redirect::Range, DataFlow::MethodCallNode {
    RedirectCall() {
      this.getTarget().hasQualifiedName("gopkg.in/macaron.v1", "Context", "Redirect")
    }

    override DataFlow::Node getUrl() { result = this.getArgument(0) }

    override HTTP::ResponseWriter getResponseWriter() { result.getARead() = this.getReceiver() }
  }
}
