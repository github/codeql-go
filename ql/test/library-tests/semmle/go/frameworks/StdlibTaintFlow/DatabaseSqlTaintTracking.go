// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "database/sql"

func TaintStepTest_DatabaseSqlNamed_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString656` into `intoNamedArg414`.

	// Assume that `sourceCQL` has the underlying type of `fromString656`:
	fromString656 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString656` to result `intoNamedArg414`
	// (`intoNamedArg414` is now tainted).
	intoNamedArg414 := sql.Named(fromString656, nil)

	// Return the tainted `intoNamedArg414`:
	return intoNamedArg414
}

func TaintStepTest_DatabaseSqlNamed_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface518` into `intoNamedArg650`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface518`:
	fromInterface518 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface518` to result `intoNamedArg650`
	// (`intoNamedArg650` is now tainted).
	intoNamedArg650 := sql.Named("", fromInterface518)

	// Return the tainted `intoNamedArg650`:
	return intoNamedArg650
}

func TaintStepTest_DatabaseSqlConnPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString784` into `intoStmt957`.

	// Assume that `sourceCQL` has the underlying type of `fromString784`:
	fromString784 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString784` to the result `intoStmt957`
	// (`intoStmt957` is now tainted).
	intoStmt957, _ := mediumObjCQL.PrepareContext(nil, fromString784)

	// Return the tainted `intoStmt957`:
	return intoStmt957
}

func TaintStepTest_DatabaseSqlDBPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoStmt443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString520` to the result `intoStmt443`
	// (`intoStmt443` is now tainted).
	intoStmt443, _ := mediumObjCQL.Prepare(fromString520)

	// Return the tainted `intoStmt443`:
	return intoStmt443
}

func TaintStepTest_DatabaseSqlDBPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoStmt483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString127` to the result `intoStmt483`
	// (`intoStmt483` is now tainted).
	intoStmt483, _ := mediumObjCQL.PrepareContext(nil, fromString127)

	// Return the tainted `intoStmt483`:
	return intoStmt483
}

func TaintStepTest_DatabaseSqlTxPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString989` into `intoStmt982`.

	// Assume that `sourceCQL` has the underlying type of `fromString989`:
	fromString989 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString989` to the result `intoStmt982`
	// (`intoStmt982` is now tainted).
	intoStmt982, _ := mediumObjCQL.Prepare(fromString989)

	// Return the tainted `intoStmt982`:
	return intoStmt982
}

func TaintStepTest_DatabaseSqlTxPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoStmt584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString417` to the result `intoStmt584`
	// (`intoStmt584` is now tainted).
	intoStmt584, _ := mediumObjCQL.PrepareContext(nil, fromString417)

	// Return the tainted `intoStmt584`:
	return intoStmt584
}

func TaintStepTest_DatabaseSqlTxStmt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStmt991` into `intoStmt881`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt991`:
	fromStmt991 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt991` to the result `intoStmt881`
	// (`intoStmt881` is now tainted).
	intoStmt881 := mediumObjCQL.Stmt(fromStmt991)

	// Return the tainted `intoStmt881`:
	return intoStmt881
}

func TaintStepTest_DatabaseSqlTxStmtContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromStmt186` into `intoStmt284`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt186`:
	fromStmt186 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt186` to the result `intoStmt284`
	// (`intoStmt284` is now tainted).
	intoStmt284 := mediumObjCQL.StmtContext(nil, fromStmt186)

	// Return the tainted `intoStmt284`:
	return intoStmt284
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
