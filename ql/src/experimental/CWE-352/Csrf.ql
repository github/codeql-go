/**
 * @name Cross site Request Forgery
 * @description When a web server is designed to receive a request from a client without any
 *  mechanism for verifying that it was intentionally sent, then it might be possible for an
 * attacker to trick a client into making an unintentional request to the web server which
 * will be treated as an authentic request. This can be done via a URL, image load,
 * XMLHttpRequest, etc. and can result in exposure of data or unintended code execution.
 * @kind path-problem
 * @problem.severity error
 * @precision high
 * @id go/request-forgery
 * @tags security
 *       external/cwe/cwe-918
 */

import go
import Csrf::Csrf
import DataFlow::PathGraph

from Configuration cfg, DataFlow::PathNode source, DataFlow::PathNode sink
where cfg.hasFlowPath(source, sink)
select sink, source, sink, "This router $@, doesn't add a check for CSRF $@.", source.getNode(),
  "router ", sink.getNode(), "sink"
