/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `database/sql` package. */
module DatabaseSqlTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Named(name string, value interface{}) NamedArg
      hasQualifiedName("database/sql", "Named") and
      (inp.isParameter(_) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }

  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*Conn).PrepareContext(ctx context.Context, query string) (*Stmt, error)
      this.(Method).hasQualifiedName("database/sql", "Conn", "PrepareContext") and
      (inp.isParameter(1) and outp.isResult(0))
      or
      // signature: func (*DB).Prepare(query string) (*Stmt, error)
      this.(Method).hasQualifiedName("database/sql", "DB", "Prepare") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (*DB).PrepareContext(ctx context.Context, query string) (*Stmt, error)
      this.(Method).hasQualifiedName("database/sql", "DB", "PrepareContext") and
      (inp.isParameter(1) and outp.isResult(0))
      or
      // signature: func (*Tx).Prepare(query string) (*Stmt, error)
      this.(Method).hasQualifiedName("database/sql", "Tx", "Prepare") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (*Tx).PrepareContext(ctx context.Context, query string) (*Stmt, error)
      this.(Method).hasQualifiedName("database/sql", "Tx", "PrepareContext") and
      (inp.isParameter(1) and outp.isResult(0))
      or
      // signature: func (*Tx).Stmt(stmt *Stmt) *Stmt
      this.(Method).hasQualifiedName("database/sql", "Tx", "Stmt") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Tx).StmtContext(ctx context.Context, stmt *Stmt) *Stmt
      this.(Method).hasQualifiedName("database/sql", "Tx", "StmtContext") and
      (inp.isParameter(1) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
