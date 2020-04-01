# Improvements to Go analysis

## General improvements

* Alert suppression can now be done with single-line block comments (`/* ... */`) as well as line comments (`// ...`).
* Analysis of flow through fields has been improved.
* More sources of untrusted input are modelled, which may lead to more results from the security queries.

## New queries

The CodeQL library for Go now contains a folder of simple "cookbook" queries that show how to access basic Go elements using the predicates defined by the standard library. They're intended to give you a starting point for your own experiments and to help you work out the best way to frame your questions using CodeQL. You can find them in the `examples/snippets` folder in the [CodeQL for Go repository](https://github.com/github/codeql-go/tree/master/ql/examples/snippets).

| **Query**                                                                 | **Tags**                                                                   | **Purpose**                                                                                                                                            |
|---------------------------------------------------------------------------|----------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| Bad check of redirect URL (`go/bad-redirect-check`) | correctness, security, external/cwe/cwe-601 | Highlights checks that ensure redirect URLs start with `/` but don't check for `//` or `/\`. Results are shown on LGTM by default. |
| Constant length comparison (`go/constant-length-comparison`)     | correctness | Highlights code that checks the length of an array or slice against a constant before indexing it using a variable, suggesting a logic error. Results are shown on LGTM by default. |
| Impossible interface nil check (`go/impossible-interface-nil-check`) | correctness | Highlights code that compares an interface value that cannot be `nil` to `nil`, suggesting a logic error. Results are shown on LGTM by default. |
| Incomplete URL scheme check (`go/incomplete-url-scheme-check`) | correctness, security, external/cwe/cwe-020 | Highlights checks for `javascript` URLs that do not take `data` or `vbscript` URLs into account. Results are shown on LGTM by default. |
| Potentially unsafe quoting (`go/unsafe-quoting`) | correctness, security, external/cwe/cwe-078, external/cwe/cwe-089, external/cwe/cwe-094 | Highlights code that constructs a quoted string literal containing data that may itself contain quotes. Results are shown on LGTM by default. |

## Changes to existing queries

| **Query**                                           | **Expected impact**          | **Change**                                                |
|-----------------------------------------------------|------------------------------|-----------------------------------------------------------|
| Database query built from user-controlled sources (`go/sql-injection`) | More results | The library models used by the query have been improved, allowing it to flag more potentially problematic cases. |
| Identical operands (`go/redundant-operation`) | Fewer false positives | The query no longer flags cases where the operands have the same value but are syntactically distinct, since this is usually intentional. |
| Incomplete regular expression for hostnames (`go/incomplete-hostname-regexp`) | More results | The query now flags unescaped dots before the TLD in a hostname regex. |
| Reflected cross-site scripting (`go/reflected-xss`) | Fewer results | Untrusted input flowing into an HTTP header definition or into an `fmt.Fprintf` call with a constant prefix is no longer flagged, since it is in both cases often harmless. |
| Useless assignment to field (`go/useless-assignment-to-field`) | Fewer false positives | The query now conservatively handles fields promoted through embedded pointer types. |
