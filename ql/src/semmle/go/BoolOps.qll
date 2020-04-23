/**
 * Provides predicates and classes for working with boolean operations.
 */

import go

/** Provides predicates and classes for working with boolean operations. */
module BoolOps {
  /** BoolExpr is an expression that has boolean type. */
  class BoolExpr extends Expr {
    BoolExpr() { this.getType() instanceof BoolType }

    /**
     * Holds if the value of this node can influence `node`.
     *
     * Note that this does not hold if this is a parent of `node`.
     */
    predicate influences(BoolExpr node) {
      this = node
      or
      exists(BoolExpr mid | this.influences(mid) | node.(LogicalExpr).getAnOperand() = mid)
      or
      this.influences(node.(ParenExpr).getExpr())
    }

    /** Holds if, given that this node has value `value`, `node` must have value `ensured`. */
    predicate ensuresIf(Boolean value, BoolExpr node, boolean ensured) {
      this = node and value = ensured
      or
      exists(BoolExpr mid, boolean midVal | this.ensuresIf(value, mid, midVal) |
        // child -> parent, e.g. expr ensures (expr)
        node.(ParenExpr).getExpr() = mid and ensured = midVal
        or
        node.(NotExpr).getAnOperand() = mid and ensured = midVal.booleanNot()
        or
        node.(LandExpr).getAnOperand() = mid and midVal = false and ensured = false
        or
        node.(LorExpr).getAnOperand() = mid and midVal = true and ensured = true
        or
        // parent -> child, e.g. (expr) ensures expr
        node = mid.(ParenExpr).getExpr() and ensured = midVal
        or
        node = mid.(NotExpr).getAnOperand() and ensured = midVal.booleanNot()
        or
        node = mid.(LandExpr).getAnOperand() and midVal = true and ensured = true
        or
        node = mid.(LorExpr).getAnOperand() and midVal = false and ensured = false
      )
    }
  }
}
