/**
 * Provides classes for working with concepts in the Mux HTTP middleware library.
 */

import go

/**
 * Provides classes for working with concepts in the Mux HTTP middleware library.
 */
module Mux {
  /** An access to a Mux middleware variable. */
  class RequestVars extends DataFlow::UntrustedFlowSource::Range, DataFlow::CallNode {
    RequestVars() { this.getTarget().hasQualifiedName("github.com/gorilla/mux", "Vars") }
  }

  /** 
   * Gets a data flow call node that corresponds to an expression that create a new
   * Gorilla Mux router.
   */
  class RouterDefinition extends DataFlow::CallNode {
    RouterDefinition() {
        this.getTarget().hasQualifiedName("github.com/gorilla/mux", "NewRouter")
    }
  }


  /** A call to a Gorilla Mux router method that sets up a route */
  class RouteSetup extends DataFlow::MethodCallNode {
      RouteSetup() {
          getTarget().hasQualifiedName("github.com/gorilla/mux", "Router", "HandleFunc")
          or
          getTarget().hasQualifiedName("github.com/gorilla/mux", "Router", "Handle")
          or
          getTarget().hasQualifiedName("github.com/gorilla/mux", "Route", "HandleFunc")
          or
          getTarget().hasQualifiedName("github.com/gorilla/mux", "Route", "Handle")
      }

      /** Gets the path associated with a route. */
      DataFlow::Node getPath() { 
          result = getArgument(0)
      }

      /** Gets the route handler registered by the setup */
      DataFlow::FunctionNode getRouteHandler() { 
          exists(DataFlow::CallNode call |
              call.asExpr().getAChild() = getAnArgument().asExpr().getAChild*() 
              and 
              (
                  result.getASuccessor*() = call.getAnArgument()
                  or
                  result.getFunction() = call.getTarget()
              )
          )
      }
  }
}
