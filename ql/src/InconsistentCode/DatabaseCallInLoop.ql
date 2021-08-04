/**
 * @name Database call in loop
 * @description Detects database operations within loops.
 *              Doing operations in series can be slow and lead to N+1 situations.
 * @kind path-problem
 * @problem.severity warning
 * @precision high
 * @id go/database-call-in-loop
 */

import go

class DatabaseAccess extends CallExpr {
  DatabaseAccess() {
    this.getTarget().(Method).getReceiverBaseType().getQualifiedName() = "github.com/jinzhu/gorm.DB" and
    (
      // List all terminating Gorm methods
      this.getTarget().getName() = "Find" or
      this.getTarget().getName() = "Take" or
      this.getTarget().getName() = "Last" or
      this.getTarget().getName() = "Scan" or
      this.getTarget().getName() = "Row" or
      this.getTarget().getName() = "Rows" or
      this.getTarget().getName() = "ScanRows" or
      this.getTarget().getName() = "Pluck" or
      this.getTarget().getName() = "Count" or
      this.getTarget().getName() = "FirstOrInit" or
      this.getTarget().getName() = "FindOrCreate" or
      this.getTarget().getName() = "Update" or
      this.getTarget().getName() = "Updates" or
      this.getTarget().getName() = "UpdateColumn" or
      this.getTarget().getName() = "UpdateColumns" or
      this.getTarget().getName() = "Save" or
      this.getTarget().getName() = "Create" or
      this.getTarget().getName() = "Delete" or
      this.getTarget().getName() = "Exec"
    )
  }
}

// The entrypoint for a twirp method
class TwirpApiMethod extends MethodDecl {
  TwirpApiMethod() {
    this.getFile().getAbsolutePath().matches("%/ts/twirp/results_resolver.go") and
    this.getReceiverBaseType().getName() = "ResultsResolver"
  }
}

class CallGraphNode extends Locatable {
  CallGraphNode() {
    this instanceof LoopStmt
    or
    this instanceof CallExpr
    or
    this instanceof FuncDef
  }
}

/**
 * Holds if `pred` calls `succ`, i.e. is an edge in the call graph,
 * This includes explicit edges from call -> callee, to produce better paths.
 */
predicate callGraphEdge(CallGraphNode pred, CallGraphNode succ) {
  // Go from a loop to an enclosed expression.
  pred.(LoopStmt).getBody().getAChild*() = succ.(CallExpr)
  or
  // Go from a call to the called function.
  pred.(CallExpr) = succ.(FuncDef).getACall().asExpr()
  or
  // Go from a function to an enclosed loop.
  pred.(FuncDef) = succ.(LoopStmt).getEnclosingFunction()
  or
  // Go from a function to an enclosed call.
  pred.(FuncDef) = succ.(CallExpr).getEnclosingFunction()
}

query predicate edges(CallGraphNode pred, CallGraphNode succ) {
  callGraphEdge(pred, succ) and
  // Limit the range of edges to only those that are relevant.
  // This helps to speed up the query by reducing the size of the outputted path information.
  (
    // Is between an API method and a loop
    exists(TwirpApiMethod apiMethod, LoopStmt loop |
      callGraphEdge*(apiMethod, pred) and
      callGraphEdge*(succ, loop)
    )
    or
    // Or is between a loop and a db access
    exists(LoopStmt loop, DatabaseAccess dbAccess |
      callGraphEdge*(loop, pred) and
      callGraphEdge*(succ, dbAccess)
    )
  )
}

from TwirpApiMethod apiMethod, LoopStmt loop, DatabaseAccess dbAccess
where edges*(apiMethod, loop) and edges*(loop, dbAccess)
select dbAccess, loop, dbAccess, "$@ is called in $@ and can be called from the $@ twirp endpoint", dbAccess,
  dbAccess.toString(), loop, "a loop", apiMethod, apiMethod.getName()
