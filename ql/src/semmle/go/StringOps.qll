/**
 * Provides predicates and classes for working with string operations.
 */

import go

module StringOps {
  /**
   * An expression that is equivalent to `strings.HasPrefix(A, B)` or `!strings.HasPrefix(A, B)`.
   *
   * Extends this class to refine existing API models. If you want to model new APIs,
   * extend `StringOps::HasPrefix::Range` instead.
   */
  class HasPrefix extends DataFlow::Node {
    HasPrefix::Range range;

    HasPrefix() { range = this }

    /**
     * Gets the `A` in `strings.HasPrefix(A, B)`.
     */
    DataFlow::Node getBaseString() { result = range.getBaseString() }

    /**
     * Gets the `B` in `strings.HasPrefix(A, B)`.
     */
    DataFlow::Node getSubstring() { result = range.getSubstring() }

    /**
     * Gets the polarity of the check.
     *
     * If the polarity is `false` the check returns `true` if the string does not start
     * with the given substring.
     */
    boolean getPolarity() { result = range.getPolarity() }
  }

  class StartsWith = HasPrefix;

  module HasPrefix {
    /**
     * An expression that is equivalent to `strings.HasPrefix(A, B)` or `!strings.HasPrefix(A, B)`.
     *
     * Extends this class to model new APIs. If you want to refine existing API models, extend
     * `StringOps::HasPrefix` instead.
     */
    abstract class Range extends DataFlow::Node {
      /**
       * Gets the `A` in `strings.HasPrefix(A, B)`.
       */
      abstract DataFlow::Node getBaseString();

      /**
       * Gets the `B` in `strings.HasPrefix(A, B)`.
       */
      abstract DataFlow::Node getSubstring();

      /**
       * Gets the polarity of the check.
       *
       * If the polarity is `false` the check returns `true` if the string does not start
       * with the given substring.
       */
      boolean getPolarity() { result = true }
    }

    /**
     * An expression of form `strings.HasPrefix(A, B)`.
     */
    private class StringsHasPrefix extends Range, DataFlow::CallNode {
      StringsHasPrefix() { getTarget().hasQualifiedName("strings", "HasPrefix") }

      override DataFlow::Node getBaseString() { result = getArgument(0) }

      override DataFlow::Node getSubstring() { result = getArgument(1) }
    }

    /**
     * An expression of form `strings.Index(A, B) === 0`.
     */
    private class HasPrefix_IndexOfEquals extends Range, DataFlow::EqualityTestNode {
      DataFlow::CallNode indexOf;

      HasPrefix_IndexOfEquals() {
        indexOf.getTarget().hasQualifiedName("strings", "Index") and
        getAnOperand() = globalValueNumber(indexOf).getANode() and
        getAnOperand().getIntValue() = 0
      }

      override DataFlow::Node getBaseString() { result = indexOf.getArgument(0) }

      override DataFlow::Node getSubstring() { result = indexOf.getArgument(1) }

      override boolean getPolarity() { result = expr.getPolarity() }
    }

    /**
     * A comparison of form `x[0] === 'k'` for some rune literal `k`.
     */
    private class HasPrefix_FirstCharacter extends Range, DataFlow::EqualityTestNode {
      DataFlow::ElementReadNode read;
      DataFlow::Node runeLiteral;

      HasPrefix_FirstCharacter() {
        read.getBase().getType().getUnderlyingType() instanceof StringType and
        read.getIndex().getIntValue() = 0 and
        eq(_, globalValueNumber(read).getANode(), runeLiteral)
      }

      override DataFlow::Node getBaseString() { result = read.getBase() }

      override DataFlow::Node getSubstring() { result = runeLiteral }

      override boolean getPolarity() { result = expr.getPolarity() }
    }

    /**
     * A comparison of form `x[:len(y)] === y`.
     */
    private class HasPrefix_Substring extends Range, DataFlow::EqualityTestNode {
      DataFlow::SliceNode slice;
      DataFlow::Node substring;

      HasPrefix_Substring() {
        eq(_, slice, substring) and
        slice.getLow().getIntValue() = 0 and
        (
          exists(DataFlow::CallNode len |
            len = Builtin::len().getACall() and
            len.getArgument(0) = globalValueNumber(substring).getANode() and
            slice.getHigh() = globalValueNumber(len).getANode()
          )
          or
          substring.getStringValue().length() = slice.getHigh().getIntValue()
        )
      }

      override DataFlow::Node getBaseString() { result = slice.getBase() }

      override DataFlow::Node getSubstring() { result = substring }

      override boolean getPolarity() { result = expr.getPolarity() }
    }
  }

  /**
   * A data-flow node that performs string concatenation.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `StringOps::Concatenation::Range` instead.
   */
  class Concatenation extends DataFlow::Node {
    Concatenation::Range self;

    Concatenation() { this = self }

    /**
     * Gets the `n`th operand of this string concatenation, if there is a data-flow node for it.
     */
    DataFlow::Node getOperand(int n) { result = self.getOperand(n) }

    /**
     * Gets the string value of the `n`th operand of this string concatenation, if it is a constant.
     */
    string getOperandStringValue(int n) { result = self.getOperandStringValue(n) }

    /**
     * Gets the number of operands of this string concatenation.
     */
    int getNumOperand() { result = self.getNumOperand() }
  }

  module Concatenation {
    /**
     * A data-flow node that performs string concatenation.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `StringOps::Concatenation` instead.
     */
    abstract class Range extends DataFlow::Node {
      /**
       * Gets the `n`th operand of this string concatenation, if there is a data-flow node for it.
       */
      abstract DataFlow::Node getOperand(int n);

      /**
       * Gets the string value of the `n`th operand of this string concatenation, if it is
       * a constant.
       */
      string getOperandStringValue(int n) { result = getOperand(n).getStringValue() }

      /**
       * Gets the number of operands of this string concatenation.
       */
      int getNumOperand() { result = count(getOperand(_)) }
    }

    /** A string concatenation using the `+` or `+=` operator. */
    private class PlusConcat extends Range, DataFlow::BinaryOperationNode {
      PlusConcat() {
        getType() instanceof StringType and
        getOperator() = "+"
      }

      override DataFlow::Node getOperand(int n) {
        n = 0 and result = getLeftOperand()
        or
        n = 1 and result = getRightOperand()
      }
    }

    /**
     * Gets a regular expression for matching simple format-string components, including flags,
     * width and precision specifiers, but not including `*` specifiers or explicit argument
     * indices.
     */
    pragma[noinline]
    private string getFormatComponentRegex() {
      exists(string literal, string opt_flag, string width, string prec, string opt_width_and_prec, string operator, string verb |
        literal = "([^%]|%%)+" and
        opt_flag = "[-+ #0]?" and
        width = "\\d+|\\*" and
        prec = "\\.(\\d+|\\*)" and
        // either a width followed by an optional prec, or just a prec, or nothing
        opt_width_and_prec = "((" + width + ")(" + prec + ")?|(" + prec + "))?" and
        operator = "[bcdeEfFgGoOpqstTxXUv]" and
        verb = "(%" + opt_flag + opt_width_and_prec + operator + ")"
      |
        result = "(" + literal + "|" + verb + ")"
      )
    }

    /**
     * A call to `fmt.Sprintf`, considered as a string concatenation.
     *
     * Only calls with simple format strings (no `*` specifiers, no explicit argument indices)
     * are supported. Such format strings can be viewed as sequences of alternating literal and
     * non-literal components. A literal component contains no `%` characters except `%%` pairs,
     * while a non-literal component consists of `%`, a verb, and possibly flags and specifiers.
     * Each non-literal component consumes exactly one argument.
     *
     * Literal components give rise to concatenation operands that have a string value but no
     * data-flow node; non-literal `%s` or `%v` components give rise to concatenation operands
     * that do have an associated data-flow node but possibly no string value; any other non-literal
     * components give rise to concatenation operands that have neither an associated data-flow
     * node nor a string value. This is because verbs like `%q` perform additional string
     * transformations that we cannot easily represent.
     */
    private class SprintfConcat extends Range, DataFlow::CallNode {
      string fmt;

      SprintfConcat() {
        exists(Function sprintf | sprintf.hasQualifiedName("fmt", "Sprintf") |
          this = sprintf.getACall() and
          fmt = getArgument(0).getStringValue() and
          fmt.regexpMatch(getFormatComponentRegex() + "*")
        )
      }

      /**
       * Gets the `n`th component of this format string.
       */
      private string getComponent(int n) {
        result = fmt.regexpFind(getFormatComponentRegex(), n, _)
      }

      override DataFlow::Node getOperand(int n) {
        exists(int i, string part | part = "%s" or part = "%v" |
          part = getComponent(n) and
          i = n / 2 and
          result = getArgument(i + 1)
        )
      }

      override string getOperandStringValue(int n) {
        result = Range.super.getOperandStringValue(n)
        or
        exists(string cmp | cmp = getComponent(n) |
          (cmp.charAt(0) != "%" or cmp.charAt(1) = "%") and
          result = cmp.replaceAll("%%", "%")
        )
      }

      override int getNumOperand() { result = max(int i | exists(getComponent(i))) + 1 }
    }

    /**
     * Holds if `src` flows to `dst` through the `n`th operand of the given concatenation operator.
     */
    predicate taintStep(DataFlow::Node src, DataFlow::Node dst, Concatenation cat, int n) {
      src = cat.getOperand(n) and
      dst = cat
    }

    /**
     * Holds if there is a taint step from `src` to `dst` through string concatenation.
     */
    predicate taintStep(DataFlow::Node src, DataFlow::Node dst) { taintStep(src, dst, _, _) }
  }

  newtype TConcatenationElement =
    /** A root concatenation element that is not itself an operand of a string concatenation. */
    MkConcatenationRoot(Concatenation cat) { not cat = any(Concatenation parent).getOperand(_) } or
    /** A concatenation element that is an operand of a string concatenation. */
    MkConcatenationOperand(Concatenation parent, int i) { i in [0 .. parent.getNumOperand() - 1] }

  /**
   * An element of a string concatenation, which either itself performs a string concatenation or
   * occurs as an operand in a string concatenation.
   *
   * For example, the expression `x + y + z` contains the following concatenation
   * elements:
   *
   * - The leaf elements `x`, `y`, and `z`
   * - The intermediate element `x + y`, which is both a concatenation and an operand
   * - The root element `x + y + z`
   */
  class ConcatenationElement extends TConcatenationElement {
    /**
     * Gets the data-flow node corresponding to this concatenation element, if any.
     */
    DataFlow::Node asNode() {
      this = MkConcatenationRoot(result)
      or
      exists(Concatenation parent, int i | this = MkConcatenationOperand(parent, i) |
        result = parent.getOperand(i)
      )
    }

    /**
     * Gets the string value of this concatenation element if it is a constant.
     */
    string getStringValue() {
      result = asNode().getStringValue()
      or
      exists(Concatenation parent, int i | this = MkConcatenationOperand(parent, i) |
        result = parent.getOperandStringValue(i)
      )
    }

    /**
     * Gets the `n`th operand of this string concatenation.
     */
    ConcatenationOperand getOperand(int n) { result = MkConcatenationOperand(asNode(), n) }

    /**
     * Gets an operand of this string concatenation.
     */
    ConcatenationOperand getAnOperand() { result = this.getOperand(_) }

    /**
     * Gets the number of operands of this string concatenation.
     */
    int getNumOperand() { result = count(this.getAnOperand()) }

    /**
     * Gets the first operand of this string concatenation.
     *
     * For example, the first operand of `(x + y) + z` is `(x + y)`.
     */
    ConcatenationOperand getFirstOperand() { result = getOperand(0) }

    /**
     * Gets the last operand of this string concatenation.
     *
     * For example, the last operand of `x + (y + z)` is `(y + z)`.
     */
    ConcatenationOperand getLastOperand() { result = getOperand(getNumOperand() - 1) }

    /**
     * Gets the root of the concatenation tree to which this element belongs.
     */
    ConcatenationRoot getConcatenationRoot() { this = result.getAnOperand*() }

    /**
     * Gets a leaf in the concatenation tree that this element is the root of.
     */
    ConcatenationLeaf getALeaf() { result = this.getAnOperand*() }

    /**
     * Gets the first leaf in this concatenation tree.
     *
     * For example, the first leaf of `(x + y) + z` is `x`.
     */
    ConcatenationLeaf getFirstLeaf() { result = getFirstOperand*() }

    /**
     * Gets the last leaf in this concatenation tree.
     *
     * For example, the last leaf of `x + (y + z)` is `z`.
     */
    ConcatenationLeaf getLastLeaf() { result = getLastOperand*() }

    /** Gets a textual representation of this concatenation element. */
    string toString() {
      if exists(asNode())
      then result = asNode().toString()
      else
        if exists(getStringValue())
        then result = getStringValue()
        else result = "concatenation element"
    }

    /**
     * Holds if this element is at the specified location.
     * The location spans column `startcolumn` of line `startline` to
     * column `endcolumn` of line `endline` in file `filepath`.
     * For more information, see
     * [Locations](https://help.semmle.com/QL/learn-ql/ql/locations.html).
     */
    predicate hasLocationInfo(
      string filepath, int startline, int startcolumn, int endline, int endcolumn
    ) {
      asNode().hasLocationInfo(filepath, startline, startcolumn, endline, endcolumn)
      or
      // use dummy location for elements that don't have a corresponding node
      not exists(asNode()) and
      filepath = "" and
      startline = 0 and
      startcolumn = 0 and
      endline = 0 and
      endcolumn = 0
    }
  }

  /**
   * One of the operands in a string concatenation.
   *
   * See `ConcatenationElement` for more information.
   */
  class ConcatenationOperand extends ConcatenationElement, MkConcatenationOperand { }

  /**
   * A data-flow node that performs a string concatenation, and is not an
   * immediate operand in a larger string concatenation.
   *
   * See `ConcatenationElement` for more information.
   */
  class ConcatenationRoot extends ConcatenationElement, MkConcatenationRoot { }

  /**
   * An operand to a concatenation that is not itself a concatenation.
   *
   * See `ConcatenationElement` for more information.
   */
  class ConcatenationLeaf extends ConcatenationOperand {
    ConcatenationLeaf() { not exists(getAnOperand()) }

    /**
     * Gets the operand immediately preceding this one in its parent concatenation.
     *
     * For example, in `(x + y) + z`, the previous leaf for `z` is `y`.
     */
    ConcatenationLeaf getPreviousLeaf() {
      exists(ConcatenationElement parent, int i |
        result = parent.getOperand(i - 1).getLastLeaf() and
        this = parent.getOperand(i).getFirstLeaf()
      )
    }

    /**
     * Gets the operand immediately succeeding this one in its parent concatenation.
     *
     * For example, in `(x + y) + z`, the previous leaf for `y` is `z`.
     */
    ConcatenationLeaf getNextLeaf() { this = result.getPreviousLeaf() }
  }
}
