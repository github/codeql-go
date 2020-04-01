/**
 * Provides classes for working with SQL-related concepts such as queries.
 */

import go

module SQL {
  /**
   * A data-flow node whose string value is interpreted as (part of) a SQL query.
   *
   * Extends this class to refine existing API models. If you want to model new APIs,
   * extend `SQL::QueryString::Range` instead.
   */
  class QueryString extends DataFlow::Node {
    QueryString::Range self;

    QueryString() { this = self }
  }

  module QueryString {
    /**
     * A data-flow node whose string value is interpreted as (part of) a SQL query.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `SQL::QueryString` instead.
     */
    abstract class Range extends DataFlow::Node { }

    /** A query string used in an API function of the standard `database/sql` package. */
    private class StandardQueryString extends Range {
      StandardQueryString() {
        exists(Method meth, string base, string m, int n |
          (
            meth.hasQualifiedName("database/sql", "DB", m) or
            meth.hasQualifiedName("database/sql", "Tx", m)
          ) and
          this = meth.getACall().getArgument(n)
        |
          (base = "Exec" or base = "Prepare" or base = "Query" or base = "QueryRow") and
          (
            m = base and n = 0
            or
            m = base + "Context" and n = 1
          )
        )
      }
    }

    /**
     * An argument to an API of the squirrel library that is directly interpreted as SQL without
     * taking syntactic structure into account.
     */
    private class SquirrelQueryString extends Range {
      SquirrelQueryString() {
        exists(Function fn |
          exists(string sq |
            sq = "github.com/Masterminds/squirrel" or
            sq = "github.com/lann/squirrel"
          |
            // first argument to `squirrel.Expr`
            fn.hasQualifiedName(sq, "Expr")
            or
            // first argument to the `Prefix` or `Suffix` method of one of the `*Builder` classes
            exists(string builder | builder.matches("%Builder") |
              fn.(Method).hasQualifiedName(sq, builder, "Prefix") or
              fn.(Method).hasQualifiedName(sq, builder, "Suffix")
            )
          ) and
          this = fn.getACall().getArgument(0) and
          this.getType().getUnderlyingType() instanceof StringType
        )
      }
    }

    /** A string that might identify package `go-pg/pg` or a specific version of it. */
    bindingset[result]
    private string gopg() {
      result.regexpMatch("github.com/go-pg/pg(/v[^/]+)?")
    }

    /** A string that might identify package `go-pg/pg/orm` or a specific version of it. */
    bindingset[result]
    private string gopgorm() {
      result.regexpMatch("github.com/go-pg/pg(/v[^/]+)?/orm")
    }

    /**
     * A string argument to an API of `go-pg/pg` that is directly interpreted as SQL without
     * taking syntactic structure into account.
     */
    private class PgQueryString extends Range {
      PgQueryString() {
        exists(Function f, int arg |
          f.hasQualifiedName(gopg(), "Q") and
          arg = 0
          or
          exists(string tp, string m | f.(Method).hasQualifiedName(gopg(), tp, m) |
            (tp = "Conn" or tp = "DB" or tp = "Tx") and
            (
              m = "FormatQuery" and arg = 1
              or
              m = "Prepare" and arg = 0
            )
          )
        |
          this = f.getACall().getArgument(arg)
        )
      }
    }

    /**
     * A string argument to an API of `go-pg/pg/orm` that is directly interpreted as SQL without
     * taking syntactic structure into account.
     */
    private class PgOrmQueryString extends Range {
      PgOrmQueryString() {
        exists(Function f, int arg |
          f.hasQualifiedName(gopgorm(), "Q") and
          arg = 0
          or
          exists(string tp, string m |
            f.(Method).hasQualifiedName(gopgorm(), tp, m)
          |
            tp = "Query" and
            (
              m = "ColumnExpr" or
              m = "For" or
              m = "Having" or
              m = "Where" or
              m = "WhereIn" or
              m = "WhereInMulti" or
              m = "WhereOr"
            ) and
            arg = 0
            or
            tp = "Query" and
            m = "FormatQuery" and
            arg = 1
          )
        |
          this = f.getACall().getArgument(arg)
        )
      }
    }

    /** A taint model for various methods on the struct `Formatter` of `go-pg/pg/orm`. */
    private class PgOrmFormatterFunction extends TaintTracking::FunctionModel, Method {
      FunctionInput i;
      FunctionOutput o;

      PgOrmFormatterFunction() {
        exists(string m | this.hasQualifiedName(gopgorm(), "Formatter", m) |
          // func (f Formatter) Append(dst []byte, src string, params ...interface{}) []byte
          // func (f Formatter) AppendBytes(dst, src []byte, params ...interface{}) []byte
          // func (f Formatter) FormatQuery(dst []byte, query string, params ...interface{}) []byte
          (m = "Append" or m = "AppendBytes" or m = "FormatQuery") and
          i.isParameter(1) and
          (o.isParameter(0) or o.isResult())
        )
      }

      override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
        inp = i and outp = o
      }
    }
  }
}
