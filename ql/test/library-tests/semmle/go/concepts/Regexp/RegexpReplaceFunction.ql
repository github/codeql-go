import go

from RegexpReplaceFunction rrfc, DataFlow::CallNode call
where call = rrfc.getACallIncludingExternals()
select rrfc, rrfc.getRegexp(call), rrfc.getSource().getNode(call), rrfc.getResult().getNode(call)
