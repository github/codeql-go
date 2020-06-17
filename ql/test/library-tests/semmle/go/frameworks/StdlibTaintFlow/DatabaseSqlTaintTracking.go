package main

import "database/sql"

func TaintStepTest_DatabaseSqlNamed_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString873` into `intoNamedArg963`.

	// Assume that `sourceCQL` has the underlying type of `fromString873`:
	fromString873 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString873` to result `intoNamedArg963`
	// (`intoNamedArg963` is now tainted).
	intoNamedArg963 := sql.Named(fromString873, nil)

	// Return the tainted `intoNamedArg963`:
	return intoNamedArg963
}

func TaintStepTest_DatabaseSqlNamed_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface355` into `intoNamedArg273`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface355`:
	fromInterface355 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface355` to result `intoNamedArg273`
	// (`intoNamedArg273` is now tainted).
	intoNamedArg273 := sql.Named("", fromInterface355)

	// Return the tainted `intoNamedArg273`:
	return intoNamedArg273
}

func TaintStepTest_DatabaseSqlConnPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString256` into `intoStmt508`.

	// Assume that `sourceCQL` has the underlying type of `fromString256`:
	fromString256 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString256` to the result `intoStmt508`
	// (`intoStmt508` is now tainted).
	intoStmt508, _ := mediumObjCQL.PrepareContext(nil, fromString256)

	// Return the tainted `intoStmt508`:
	return intoStmt508
}

func TaintStepTest_DatabaseSqlDBPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString934` into `intoStmt270`.

	// Assume that `sourceCQL` has the underlying type of `fromString934`:
	fromString934 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString934` to the result `intoStmt270`
	// (`intoStmt270` is now tainted).
	intoStmt270, _ := mediumObjCQL.Prepare(fromString934)

	// Return the tainted `intoStmt270`:
	return intoStmt270
}

func TaintStepTest_DatabaseSqlDBPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString348` into `intoStmt800`.

	// Assume that `sourceCQL` has the underlying type of `fromString348`:
	fromString348 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString348` to the result `intoStmt800`
	// (`intoStmt800` is now tainted).
	intoStmt800, _ := mediumObjCQL.PrepareContext(nil, fromString348)

	// Return the tainted `intoStmt800`:
	return intoStmt800
}

func TaintStepTest_DatabaseSqlTxPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString849` into `intoStmt202`.

	// Assume that `sourceCQL` has the underlying type of `fromString849`:
	fromString849 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString849` to the result `intoStmt202`
	// (`intoStmt202` is now tainted).
	intoStmt202, _ := mediumObjCQL.Prepare(fromString849)

	// Return the tainted `intoStmt202`:
	return intoStmt202
}

func TaintStepTest_DatabaseSqlTxPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString589` into `intoStmt743`.

	// Assume that `sourceCQL` has the underlying type of `fromString589`:
	fromString589 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString589` to the result `intoStmt743`
	// (`intoStmt743` is now tainted).
	intoStmt743, _ := mediumObjCQL.PrepareContext(nil, fromString589)

	// Return the tainted `intoStmt743`:
	return intoStmt743
}

func TaintStepTest_DatabaseSqlTxStmt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStmt561` into `intoStmt503`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt561`:
	fromStmt561 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt561` to the result `intoStmt503`
	// (`intoStmt503` is now tainted).
	intoStmt503 := mediumObjCQL.Stmt(fromStmt561)

	// Return the tainted `intoStmt503`:
	return intoStmt503
}

func TaintStepTest_DatabaseSqlTxStmtContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStmt451` into `intoStmt779`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt451`:
	fromStmt451 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt451` to the result `intoStmt779`
	// (`intoStmt779` is now tainted).
	intoStmt779 := mediumObjCQL.StmtContext(nil, fromStmt451)

	// Return the tainted `intoStmt779`:
	return intoStmt779
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
