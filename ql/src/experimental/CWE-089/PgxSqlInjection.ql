/**
 * @name Database pgx query built from user-controlled sources
 * @description Building a database pgx query from user-controlled sources is vulnerable to insertion of
 *              malicious code by the user.
 * @kind path-problem
 * @problem.severity error
 * @security-severity 8.8
 * @precision high
 * @id go/pgx-sql-injection
 * @tags security
 *       external/cwe/cwe-089
 */

import go
import experimental.frameworks.Pgx
import DataFlow::PathGraph

class PgxSqlInjectionConfiguration extends TaintTracking::Configuration {
  PgxSqlInjectionConfiguration() { this = "PgxSqlInjectionConfiguration" }

  override predicate isSource(DataFlow::Node source) { source instanceof UntrustedFlowSource }

  override predicate isSink(DataFlow::Node sink) { sink instanceof Pgx::QueryString }
}

from PgxSqlInjectionConfiguration cfg, DataFlow::PathNode source, DataFlow::PathNode sink
where cfg.hasFlowPath(source, sink)
select sink.getNode(), source, sink, "This query depends on $@.", source.getNode(),
  "a user-provided value"
