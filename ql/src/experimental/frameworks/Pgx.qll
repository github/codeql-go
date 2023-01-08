/**
 * Provides predicates and class for working with sql injection sinks
 * from the `github.com/jackc/pgx` and `github.com/jackc/pgconn` packages.
 */

import go

module Pgx {
  /** Gets the package name for Jackc. */
  string pgx3PackagePath() { result = "github.com/jackc/pgx" }

  string pgxStdlibPackagePath() { result = package("github.com/jackc/pgx", "stdlib") }

  string pgxPackagePath() {
    result = package("github.com/jackc/pgx", "") and result != pgx3PackagePath()
  }

  string pgxAnyPackagePath() { result = package("github.com/jackc/pgx", "") }

  string pgxPoolPackagePath() { result = package("github.com/jackc/pgx", "pgxpool") }

  string pgconnPackagePath() { result = package("github.com/jackc/pgconn", "") }

  predicate batchQueries(string pkg, string tp, string m, int arg) {
    pkg = pgxAnyPackagePath() and
    tp = "Batch" and
    (m = "Queue" and arg = 0)
  }

  predicate stdlibQueries(string pkg, string tp, string m, int arg) {
    pkg = pgxStdlibPackagePath() and
    tp = "Conn" and
    (
      m = "Prepare" and arg = 0
      or
      m = "PrepareContext" and arg = 1
      or
      m = "Exec" and arg = 0
      or
      m = "ExecContext" and arg = 1
      or
      m = "Query" and arg = 0
      or
      m = "QueryContext" and arg = 1
    )
  }

  predicate v3Queries(string pkg, string tp, string m, int arg) {
    pkg = pgx3PackagePath() and
    tp = ["Tx", "Conn", "ConnPool", "ReplicationConn"] and
    (
      m = "Exec" and arg = 0
      or
      m = "ExecEx" and arg = 1
      or
      m = "Prepare" and arg = 1
      or
      m = "PrepareEx" and arg = 2
      or
      m = "Query" and arg = 0
      or
      m = "QueryEx" and arg = 1
      or
      m = "QueryRow" and arg = 0
      or
      m = "QueryRowEx" and arg = 1
    )
  }

  predicate v4Queries(string pkg, string tp, string m, int arg) {
    pkg = pgxPackagePath() and
    tp = ["Conn", "Tx"] and
    (
      m = "Prepare" and arg = 2
      or
      m = "Exec" and arg = 1
      or
      m = "Query" and arg = 1
      or
      m = "QueryRow" and arg = 1
      or
      m = "QueryFunc" and arg = 1
    )
  }

  predicate pgxpoolConnPoolQueries(string pkg, string tp, string m, int arg) {
    pkg = pgxPoolPackagePath() and
    tp = ["Conn", "Pool"] and
    (
      m = "Exec" and arg = 1
      or
      m = "Query" and arg = 1
      or
      m = "QueryRow" and arg = 1
      or
      m = "QueryFunc" and arg = 1
    )
  }

  predicate pgxpoolTxQueries(string pkg, string tp, string m, int arg) {
    pkg = pgxPoolPackagePath() and
    tp = ["Tx"] and
    (
      m = "Exec" and arg = 1
      or
      m = "Query" and arg = 1
      or
      m = "QueryRow" and arg = 1
      or
      m = "QueryFunc" and arg = 1
      or
      m = "Prepare" and arg = 2
    )
  }

  predicate pgconnQueries(string pkg, string tp, string m, int arg) {
    pkg = pgconnPackagePath() and
    tp = "PgConn" and
    (
      m = "Exec" and arg = 1
      or
      m = "ExecParams" and arg = 1
      or
      m = "CopyTo" and arg = 2
      or
      m = "CopyFrom" and arg = 2
      or
      m = "Prepare" and arg = 2
    )
  }

  predicate pgxQueries(string pkg, string tp, string m, int arg) {
    pgconnQueries(pkg, tp, m, arg)
    or
    pgxpoolTxQueries(pkg, tp, m, arg)
    or
    pgxpoolConnPoolQueries(pkg, tp, m, arg)
    or
    v3Queries(pkg, tp, m, arg)
    or
    v4Queries(pkg, tp, m, arg)
    or
    stdlibQueries(pkg, tp, m, arg)
    or
    batchQueries(pkg, tp, m, arg)
  }

  /** A model for sinks of github.com/jackc/pgx. */
  class QueryString extends DataFlow::Node {
    // class QueryString extends SQL::QueryString {
    QueryString() {
      exists(Method meth, string pkg, string tp, string m, int arg |
        pgxQueries(pkg, tp, m, arg) and
        meth.hasQualifiedName(pkg, tp, m) and
        this = meth.getACall().getArgument(arg)
      )
    }
  }
}
