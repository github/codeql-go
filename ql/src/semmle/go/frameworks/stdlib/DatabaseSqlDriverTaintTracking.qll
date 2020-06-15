/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `database/sql/driver` package. */
module DatabaseSqlDriverTaintTracking {
  private class NotNullConvertValue extends TaintTracking::FunctionModel, Method {
    // signature: func (NotNull).ConvertValue(v interface{}) (Value, error)
    NotNullConvertValue() {
      this.(Method).hasQualifiedName("database/sql/driver", "NotNull", "ConvertValue")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class NullConvertValue extends TaintTracking::FunctionModel, Method {
    // signature: func (Null).ConvertValue(v interface{}) (Value, error)
    NullConvertValue() {
      this.(Method).hasQualifiedName("database/sql/driver", "Null", "ConvertValue")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ValueConverterConvertValue extends TaintTracking::FunctionModel, Method {
    // signature: func (ValueConverter).ConvertValue(v interface{}) (Value, error)
    ValueConverterConvertValue() {
      this.implements("database/sql/driver", "ValueConverter", "ConvertValue")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ConnPrepare extends TaintTracking::FunctionModel, Method {
    // signature: func (Conn).Prepare(query string) (Stmt, error)
    ConnPrepare() { this.implements("database/sql/driver", "Conn", "Prepare") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ConnPrepareContextPrepareContext extends TaintTracking::FunctionModel, Method {
    // signature: func (ConnPrepareContext).PrepareContext(ctx context.Context, query string) (Stmt, error)
    ConnPrepareContextPrepareContext() {
      this.implements("database/sql/driver", "ConnPrepareContext", "PrepareContext")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult(0))
    }
  }

  private class ValuerValue extends TaintTracking::FunctionModel, Method {
    // signature: func (Valuer).Value() (Value, error)
    ValuerValue() { this.implements("database/sql/driver", "Valuer", "Value") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }
}
