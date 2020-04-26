/**
 * Provides classes for working with concepts relating to the Macaron web framework
 */

import go

private module Macaron {
  private predicate containsContextType(Type t) {
    t.hasQualifiedName("gopkg.in/macaron.v1", "Context")
    or
    t.(StructType).hasOwnField(_, _, any(Type t1 | containsContextType(t1)), true)
    or
    containsContextType(t.getUnderlyingType())
    or
    containsContextType(t.(PointerType).getBaseType())
  }

  private class Context extends HTTP::ResponseWriter::Range {
    Context() { containsContextType(this.getType()) }
  }

  private class RedirectCall extends HTTP::Redirect::Range, DataFlow::MethodCallNode {
    RedirectCall() {
      this.getTarget().hasQualifiedName("gopkg.in/macaron.v1", "Context", "Redirect")
    }

    override DataFlow::Node getUrl() { result = this.getArgument(0) }

    override HTTP::ResponseWriter getResponseWriter() { result.getARead() = this.getReceiver() }
  }
}
