/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `database/sql/driver` package. */
module DatabaseSqlDriverTaintTracking {
  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (NotNull).ConvertValue(v interface{}) (Value, error)
      this.(Method).hasQualifiedName("database/sql/driver", "NotNull", "ConvertValue") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (Null).ConvertValue(v interface{}) (Value, error)
      this.(Method).hasQualifiedName("database/sql/driver", "Null", "ConvertValue") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // Interfaces:
      // signature: func (ValueConverter).ConvertValue(v interface{}) (Value, error)
      this.implements("database/sql/driver", "ValueConverter", "ConvertValue") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (Conn).Prepare(query string) (Stmt, error)
      this.implements("database/sql/driver", "Conn", "Prepare") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (ConnPrepareContext).PrepareContext(ctx context.Context, query string) (Stmt, error)
      this.implements("database/sql/driver", "ConnPrepareContext", "PrepareContext") and
      (inp.isParameter(1) and outp.isResult(0))
      or
      // signature: func (Valuer).Value() (Value, error)
      this.implements("database/sql/driver", "Valuer", "Value") and
      (inp.isReceiver() and outp.isResult(0))
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
