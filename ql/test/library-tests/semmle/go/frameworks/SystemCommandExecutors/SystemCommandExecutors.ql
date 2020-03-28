/**
 * @kind path-problem
 */

import go
import semmle.go.security.CommandInjection
import DataFlow::PathGraph

from CommandInjection::Configuration cfg, DataFlow::PathNode source, DataFlow::PathNode sink
where cfg.hasFlowPath(source, sink)
select sink.getNode(), source, sink, "This command depends on $@.", source.getNode(),
  "a user-provided value"
