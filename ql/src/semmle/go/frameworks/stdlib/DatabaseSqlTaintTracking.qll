/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `database/sql` package. */
module DatabaseSqlTaintTracking {
  private class Named extends TaintTracking::FunctionModel {
    // signature: func Named(name string, value interface{}) NamedArg
    Named() { hasQualifiedName("database/sql", "Named") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class ConnPrepareContext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Conn).PrepareContext(ctx context.Context, query string) (*Stmt, error)
    ConnPrepareContext() {
      this.(Method).hasQualifiedName("database/sql", "Conn", "PrepareContext")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult(0))
    }
  }

  private class DBPrepare extends TaintTracking::FunctionModel, Method {
    // signature: func (*DB).Prepare(query string) (*Stmt, error)
    DBPrepare() { this.(Method).hasQualifiedName("database/sql", "DB", "Prepare") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class DBPrepareContext extends TaintTracking::FunctionModel, Method {
    // signature: func (*DB).PrepareContext(ctx context.Context, query string) (*Stmt, error)
    DBPrepareContext() { this.(Method).hasQualifiedName("database/sql", "DB", "PrepareContext") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult(0))
    }
  }

  private class TxPrepare extends TaintTracking::FunctionModel, Method {
    // signature: func (*Tx).Prepare(query string) (*Stmt, error)
    TxPrepare() { this.(Method).hasQualifiedName("database/sql", "Tx", "Prepare") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class TxPrepareContext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Tx).PrepareContext(ctx context.Context, query string) (*Stmt, error)
    TxPrepareContext() { this.(Method).hasQualifiedName("database/sql", "Tx", "PrepareContext") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult(0))
    }
  }

  private class TxStmt extends TaintTracking::FunctionModel, Method {
    // signature: func (*Tx).Stmt(stmt *Stmt) *Stmt
    TxStmt() { this.(Method).hasQualifiedName("database/sql", "Tx", "Stmt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TxStmtContext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Tx).StmtContext(ctx context.Context, stmt *Stmt) *Stmt
    TxStmtContext() { this.(Method).hasQualifiedName("database/sql", "Tx", "StmtContext") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }
}
