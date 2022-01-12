import go

from RegexpMatchFunction match, DataFlow::CallNode call
where call = match.getACallIncludingExternals()
select match, match.getRegexp(call), match.getValue().getNode(call), match.getResult().getNode(call)
