private import go
private import DataFlowUtil
private import DataFlowImplCommon

private newtype TReturnKind =
  MkReturnKind(int i) { exists(SignatureType st | exists(st.getResultType(i))) }

/**
 * A return kind. A return kind describes how a value can be returned
 * from a callable. For Go, this is either a return of a single value
 * or of one of multiple values.
 */
class ReturnKind extends TReturnKind {
  /** Gets a textual representation of this return kind. */
  string toString() { exists(int i | this = MkReturnKind(i) | result = "return[" + i + "]") }
}

/** A data flow node that represents returning a value from a function. */
class ReturnNode extends ResultNode {
  ReturnKind kind;

  ReturnNode() { kind = MkReturnKind(i) }

  /** Gets the kind of this returned value. */
  ReturnKind getKind() { result = kind }
}

/** A data flow node that represents the output of a call. */
class OutNode extends DataFlow::Node {
  DataFlow::CallNode call;
  int i;

  OutNode() { this = call.getResult(i) }

  /** Gets the underlying call. */
  DataFlowCall getCall() { result = call.asExpr() }
}

/**
 * Gets a node that can read the value returned from `call` with return kind
 * `kind`.
 */
OutNode getAnOutNode(DataFlowCall call, ReturnKind kind) {
  exists(DataFlow::CallNode c, int i | c.asExpr() = call and kind = MkReturnKind(i) |
    result = c.getResult(i)
  )
}

/**
 * Holds if data can flow from `node1` to `node2` in a way that loses the
 * calling context. For example, this would happen with flow through a
 * global or static variable.
 */
predicate jumpStep(Node n1, Node n2) {
  exists(ValueEntity v, Write w |
    not v instanceof SsaSourceVariable and
    not v instanceof Field and
    w.writes(v, n1) and
    n2 = v.getARead()
  )
}

private newtype TContent =
  TFieldContent(Field f) or
  TCollectionContent() or
  TArrayContent() or
  TPointerContent(PointerType p)

/**
 * A reference contained in an object. Examples include instance fields, the
 * contents of a collection object, the contents of an array or pointer.
 */
class Content extends TContent {
  /** Gets a textual representation of this element. */
  abstract string toString();

  predicate hasLocationInfo(string path, int sl, int sc, int el, int ec) {
    path = "" and sl = 0 and sc = 0 and el = 0 and ec = 0
  }
}

private class FieldContent extends Content, TFieldContent {
  Field f;

  FieldContent() { this = TFieldContent(f) }

  override string toString() { result = f.toString() }

  override predicate hasLocationInfo(string path, int sl, int sc, int el, int ec) {
    f.getDeclaration().hasLocationInfo(path, sl, sc, el, ec)
  }
}

private class CollectionContent extends Content, TCollectionContent {
  override string toString() { result = "collection" }
}

private class ArrayContent extends Content, TArrayContent {
  override string toString() { result = "array" }
}

private class PointerContent extends Content, TPointerContent {
  override string toString() { result = "pointer" }
}

/**
 * Holds if data can flow from `node1` to `node2` via an assignment to `c`.
 * Thus, `node2` references an object with a field `f` that contains the
 * value of `node1`.
 */
predicate storeStep(Node node1, Content c, PostUpdateNode node2) {
  // a write `(*p).f = rhs` is modelled as two store steps: `rhs` is flows into field `f` of `(*p)`,
  // which in turn flows into the pointer content of `p`
  exists(Write w, Field f, DataFlow::Node base, DataFlow::Node rhs | w.writesField(base, f, rhs) |
    node1 = rhs and
    node2.getPreUpdateNode() = base and
    c = TFieldContent(f)
    or
    node1 = base and
    node2.getPreUpdateNode() = node1.(PointerDereferenceNode).getOperand() and
    c = TPointerContent(node2.getType())
  )
  or
  node1 = node2.(AddressOperationNode).getOperand() and
  c = TPointerContent(node2.getType())
}

/**
 * Holds if data can flow from `node1` to `node2` via a read of `f`.
 * Thus, `node1` references an object with a field `f` whose value ends up in
 * `node2`.
 */
predicate readStep(Node node1, Content f, Node node2) {
  node1 = node2.(PointerDereferenceNode).getOperand() and
  f = TPointerContent(node1.getType())
  or
  exists(FieldReadNode read |
    node2 = read and
    node1 = read.getBase() and
    f = TFieldContent(read.getField())
  )
}

/**
 * Holds if values stored inside content `c` are cleared at node `n`.
 */
predicate clearsContent(Node n, Content c) {
  none() // stub implementation
}

/** Gets the type of `n` used for type pruning. */
DataFlowType getNodeType(Node n) { result = n.getType() }

/** Gets a string representation of a type returned by `getNodeType()`. */
string ppReprType(Type t) { result = t.toString() }

/**
 * Holds if `t1` and `t2` are compatible, that is, whether data can flow from
 * a node of type `t1` to a node of type `t2`.
 */
pragma[inline]
predicate compatibleTypes(Type t1, Type t2) {
  any() // stub implementation
}

//////////////////////////////////////////////////////////////////////////////
// Java QL library compatibility wrappers
//////////////////////////////////////////////////////////////////////////////
/** A node that performs a type cast. */
class CastNode extends ExprNode {
  override ConversionExpr expr;
}

class DataFlowCallable = FuncDef;

class DataFlowExpr = Expr;

class DataFlowType = Type;

class DataFlowLocation = Location;

/** A function call relevant for data flow. */
class DataFlowCall extends Expr {
  DataFlow::CallNode call;

  DataFlowCall() { this = call.asExpr() }

  /**
   * Gets the nth argument for this call.
   */
  Node getArgument(int n) { result = call.getArgument(n) }

  /** Gets the data flow node corresponding to this call. */
  ExprNode getNode() { result = call }

  /** Gets the enclosing callable of this call. */
  DataFlowCallable getEnclosingCallable() { result = this.getEnclosingFunction() }
}

/** Holds if `e` is an expression that always has the same Boolean value `val`. */
private predicate constantBooleanExpr(Expr e, boolean val) {
  e.getBoolValue() = val
  or
  exists(SsaExplicitDefinition v, Expr src |
    IR::evalExprInstruction(e) = v.getVariable().getAUse() and
    IR::evalExprInstruction(src) = v.getRhs() and
    constantBooleanExpr(src, val)
  )
}

/** An argument that always has the same Boolean value. */
private class ConstantBooleanArgumentNode extends ArgumentNode, ExprNode {
  ConstantBooleanArgumentNode() { constantBooleanExpr(this.getExpr(), _) }

  /** Gets the Boolean value of this expression. */
  boolean getBooleanValue() { constantBooleanExpr(this.getExpr(), result) }
}

/**
 * Returns a guard that will certainly not hold in calling context `call`.
 *
 * In particular it does not hold because it checks that `param` has value `b`, but
 * in context `call` it is known to have value `!b`. Note this is `noinline`d in order
 * to avoid a bad join order in `isUnreachableInCall`.
 */
pragma[noinline]
private ControlFlow::ConditionGuardNode getAFalsifiedGuard(DataFlowCall call) {
  exists(ParameterNode param, ConstantBooleanArgumentNode arg |
    // get constant bool argument and parameter for this call
    viableParamArg(call, param, arg) and
    // which is used in a guard controlling `n` with the opposite value of `arg`
    result.ensures(param.getAUse(), arg.getBooleanValue().booleanNot())
  )
}

/**
 * Holds if the node `n` is unreachable when the call context is `call`.
 */
predicate isUnreachableInCall(Node n, DataFlowCall call) {
  getAFalsifiedGuard(call).dominates(n.getBasicBlock())
}

int accessPathLimit() { result = 5 }

/** The unit type. */
private newtype TUnit = TMkUnit()

/** The trivial type with a single element. */
class Unit extends TUnit {
  /** Gets a textual representation of this element. */
  string toString() { result = "unit" }
}

/**
 * Gets the `i`th argument of call `c`, where the receiver of a method call
 * counts as argument -1.
 */
Node getArgument(CallNode c, int i) {
  result = c.getArgument(i)
  or
  result = c.(MethodCallNode).getReceiver() and
  i = -1
}

/** Holds if `n` should be hidden from path explanations. */
predicate nodeIsHidden(Node n) { none() }

class LambdaCallKind = Unit;

/** Holds if `creation` is an expression that creates a lambda of kind `kind` for `c`. */
predicate lambdaCreation(Node creation, LambdaCallKind kind, DataFlowCallable c) { none() }

/** Holds if `call` is a lambda call of kind `kind` where `receiver` is the lambda expression. */
predicate lambdaCall(DataFlowCall call, LambdaCallKind kind, Node receiver) { none() }

/** Extra data-flow steps needed for lambda flow analysis. */
predicate additionalLambdaFlowStep(Node nodeFrom, Node nodeTo, boolean preservesValue) { none() }
