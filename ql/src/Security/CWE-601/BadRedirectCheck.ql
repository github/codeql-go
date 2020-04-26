/**
 * @name Bad redirect check
 * @description A redirect check that checks for a leading slash but not two
 *              leading slashes or a leading slash followed by a backslash is
 *              incomplete.
 * @kind path-problem
 * @problem.severity warning
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
  }

  override predicate isSink(DataFlow::Node sink) { sink instanceof OpenUrlRedirect::Sink }
}

predicate isBadRedirectCheckOrWrapper(DataFlow::Node check, SsaWithFields v) {
  isBadRedirectCheck(check, v)
  or
  exists(FuncDef f, FunctionInput input |
    check = f.getACall() and
    input.getEntryNode(check) = v.getAUse() and
    isBadRedirectCheckWrapper(f, input)
  )
}

predicate isBadRedirectCheck(DataFlow::Node check, SsaWithFields v) {
  // a check for a leading slash
  check = checkForLeadingSlash(v) and
  // where there does not exist a check for both a second slash and a second backslash
  not (exists(checkForSecondSlash(v)) and exists(checkForSecondBackslash(v)))
}

predicate isBadRedirectCheckWrapper(FuncDef f, FunctionInput input) {
  exists(DataFlow::Node check, SsaVariable ssa |
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
