// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "database/sql"

func TaintStepTest_DatabaseSqlNamed_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString953` into `intoNamedArg403`.

	// Assume that `sourceCQL` has the underlying type of `fromString953`:
	fromString953 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString953` to result `intoNamedArg403`
	// (`intoNamedArg403` is now tainted).
	intoNamedArg403 := sql.Named(fromString953, nil)

	// Return the tainted `intoNamedArg403`:
	return intoNamedArg403
}

func TaintStepTest_DatabaseSqlNamed_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface414` into `intoNamedArg712`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface414`:
	fromInterface414 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface414` to result `intoNamedArg712`
	// (`intoNamedArg712` is now tainted).
	intoNamedArg712 := sql.Named("", fromInterface414)

	// Return the tainted `intoNamedArg712`:
	return intoNamedArg712
}

func TaintStepTest_DatabaseSqlConnPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString825` into `intoStmt646`.

	// Assume that `sourceCQL` has the underlying type of `fromString825`:
	fromString825 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString825` to the result `intoStmt646`
	// (`intoStmt646` is now tainted).
	intoStmt646, _ := mediumObjCQL.PrepareContext(nil, fromString825)

	// Return the tainted `intoStmt646`:
	return intoStmt646
}

func TaintStepTest_DatabaseSqlDBPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString632` into `intoStmt416`.

	// Assume that `sourceCQL` has the underlying type of `fromString632`:
	fromString632 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString632` to the result `intoStmt416`
	// (`intoStmt416` is now tainted).
	intoStmt416, _ := mediumObjCQL.Prepare(fromString632)

	// Return the tainted `intoStmt416`:
	return intoStmt416
}

func TaintStepTest_DatabaseSqlDBPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString469` into `intoStmt512`.

	// Assume that `sourceCQL` has the underlying type of `fromString469`:
	fromString469 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString469` to the result `intoStmt512`
	// (`intoStmt512` is now tainted).
	intoStmt512, _ := mediumObjCQL.PrepareContext(nil, fromString469)

	// Return the tainted `intoStmt512`:
	return intoStmt512
}

func TaintStepTest_DatabaseSqlTxPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString528` into `intoStmt430`.

	// Assume that `sourceCQL` has the underlying type of `fromString528`:
	fromString528 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString528` to the result `intoStmt430`
	// (`intoStmt430` is now tainted).
	intoStmt430, _ := mediumObjCQL.Prepare(fromString528)

	// Return the tainted `intoStmt430`:
	return intoStmt430
}

func TaintStepTest_DatabaseSqlTxPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString182` into `intoStmt241`.

	// Assume that `sourceCQL` has the underlying type of `fromString182`:
	fromString182 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString182` to the result `intoStmt241`
	// (`intoStmt241` is now tainted).
	intoStmt241, _ := mediumObjCQL.PrepareContext(nil, fromString182)

	// Return the tainted `intoStmt241`:
	return intoStmt241
}

func TaintStepTest_DatabaseSqlTxStmt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStmt535` into `intoStmt123`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt535`:
	fromStmt535 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt535` to the result `intoStmt123`
	// (`intoStmt123` is now tainted).
	intoStmt123 := mediumObjCQL.Stmt(fromStmt535)

	// Return the tainted `intoStmt123`:
	return intoStmt123
}

func TaintStepTest_DatabaseSqlTxStmtContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStmt977` into `intoStmt648`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt977`:
	fromStmt977 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt977` to the result `intoStmt648`
	// (`intoStmt648` is now tainted).
	intoStmt648 := mediumObjCQL.StmtContext(nil, fromStmt977)

	// Return the tainted `intoStmt648`:
	return intoStmt648
}

func RunAllTaints_DatabaseSql() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlNamed_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlNamed_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlConnPrepareContext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDBPrepare_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDBPrepareContext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlTxPrepare_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlTxPrepareContext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlTxStmt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlTxStmtContext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
