/**
 * @name Bad redirect check
 * @description A redirect check that checks for a leading slash but not two
 *              leading slashes or a leading slash followed by a backslash is
 *              incomplete.
 * @kind path-problem
 * @problem.severity error
 * @id go/bad-redirect-check
 * @tags security
 *       external/cwe/cwe-601
 * @precision high
 */

import go
import semmle.go.security.OpenUrlRedirectCustomizations
import DataFlow::PathGraph

StringOps::HasPrefix checkForLeadingSlash(SsaWithFields v) {
  exists(DataFlow::Node substr |
    result.getBaseString() = v.getAUse() and result.getSubstring() = substr
  |
    substr.getStringValue() = "/"
  )
}

DataFlow::Node checkForSecondSlash(SsaWithFields v) {
  exists(StringOps::HasPrefix hp | result = hp and hp.getBaseString() = v.getAUse() |
    hp.getSubstring().getStringValue() = "//"
  )
  or
  exists(DataFlow::EqualityTestNode eq, DataFlow::Node slash, DataFlow::ElementReadNode er |
    result = eq
  |
    slash.getStringValue() = "/" and
    er.getBase() = v.getAUse() and
    er.getIndex().getIntValue() = 1 and
    eq.eq(_, er, slash)
  )
  or
  // a call to path.Clean will strip away multiple leading slashes
  exists(DataFlow::CallNode call | call.getTarget().hasQualifiedName("path", "Clean") |
    call.getArgument(0) = v.getAUse()
  )
}

DataFlow::Node checkForSecondBackslash(SsaWithFields v) {
  exists(StringOps::HasPrefix hp | result = hp and hp.getBaseString() = v.getAUse() |
    hp.getSubstring().getStringValue() = "/\\"
  )
  or
  exists(DataFlow::EqualityTestNode eq, DataFlow::Node slash, DataFlow::ElementReadNode er |
    result = eq
  |
    slash.getStringValue() = "\\" and
    er.getBase() = v.getAUse() and
    er.getIndex().getIntValue() = 1 and
    eq.eq(_, er, slash)
  )
  or
  // if this variable comes from or is a net/url.URL.Path, backslashes are most likely sanitized,
  // as the parse functions turn them into "%5C"
  exists(DataFlow::FieldReadNode frn |
    frn.getField().getName() = "Path" and frn.getBase().getType().hasQualifiedName("net/url", "URL")
  |
    frn.getASuccessor*() = v.getAUse()
  )
}

class Configuration extends TaintTracking::Configuration {
  Configuration() { this = "BadRedirectCheck" }

  override predicate isSource(DataFlow::Node source) { this.isSource(source, _) }

  predicate isSource(DataFlow::Node source, DataFlow::Node check) {
    exists(SsaWithFields v |
      DataFlow::localFlow(source, v.getAUse()) and
      not exists(source.getAPredecessor()) and
      isBadRedirectCheckOrWrapper(check, v)
    )
  }

  override predicate isAdditionalTaintStep(DataFlow::Node pred, DataFlow::Node succ) {
    // this is very over-approximate, because most filtering is done by the isSource predicate
    exists(Write w | w.writesField(succ, _, pred))
  }

  override predicate isSanitizerOut(DataFlow::Node node) {
    // assume this value is safe if something is prepended to it.
    exists(StringOps::Concatenation conc, int i, int j | i < j |
      node = conc.getOperand(j) and
      exists(conc.getOperand(i))
    )
    or
    exists(DataFlow::CallNode call, int i | call.getTarget().hasQualifiedName("path", "Join") |
      i > 0 and node = call.getArgument(i)
    )
  }

  override predicate isSink(DataFlow::Node sink) { sink instanceof OpenUrlRedirect::Sink }
}

/**
 * Holds there is a check `check` that is a bad redirect check, and `v` is either
 * checked directly by `check` or checked by a function that contains `check`.
 */
predicate isBadRedirectCheckOrWrapper(DataFlow::Node check, SsaWithFields v) {
  isBadRedirectCheck(check, v)
  or
  exists(DataFlow::CallNode call, FuncDef f, FunctionInput input |
    call = f.getACall() and
    input.getEntryNode(call) = v.getAUse() and
    isBadRedirectCheckWrapper(check, f, input)
  )
}

/**
 * Holds if `check` checks that `v` has a leading slash, but not whether it has another slash or a
 * backslash in its second position.
 */
predicate isBadRedirectCheck(DataFlow::Node check, SsaWithFields v) {
  // a check for a leading slash
  check = checkForLeadingSlash(v) and
  // where there does not exist a check for both a second slash and a second backslash
  not (exists(checkForSecondSlash(v)) and exists(checkForSecondBackslash(v)))
}

/**
 * Holds if `f` contains a bad redirect check `check`, that checks the parameter `input`.
 */
predicate isBadRedirectCheckWrapper(DataFlow::Node check, FuncDef f, FunctionInput input) {
  exists(SsaVariable ssa |
    input.getExitNode(f) = DataFlow::ssaNode(ssa) and
    isBadRedirectCheck(check, SsaWithFields::fromSsa(ssa))
  )
}

from Configuration cfg, DataFlow::PathNode source, DataFlow::PathNode sink, DataFlow::Node check
where
  cfg.isSource(source.getNode(), check) and
  cfg.hasFlowPath(source, sink)
select check, source, sink,
  "This is a check that $@, which flows into a $@, has a leading slash, but not that does not have '/' or '\\' in its second position.",
  source.getNode(), "this value", sink.getNode(), "redirect"
