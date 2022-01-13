---
category: deprecated
---
* The member predicate `getACall` on the class `Function` has been deprecated. The recommended replacement is `getACallIncludingExternals`, which gives more results in some cases. See its QLDoc for more details. If these extra results are not desired then instead use `getACallExcludingExternals`, which gives exactly the same results as `getACall`.
