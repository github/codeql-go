package main

import "database/sql"

func TaintStepTest_DatabaseSqlNamed(sourceCQL interface{}) {
	// The flow is from `value` into `into398`.

	// Assume that `sourceCQL` has the underlying type of `value`:
	value := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `value` to result `into398`
	// (`into398` is now tainted).
	into398 := sql.Named("", value)

	// Sink the tainted `into398`:
	sink(into398)
}

func TaintStepTest_DatabaseSqlConnPrepareContext(sourceCQL interface{}) {
	// The flow is from `query` into `into420`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Conn

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into420`
	// (`into420` is now tainted).
	into420, _ := mediumObjCQL.PrepareContext(nil, query)

	// Sink the tainted `into420`:
	sink(into420)
}

func TaintStepTest_DatabaseSqlDBPrepare(sourceCQL interface{}) {
	// The flow is from `query` into `into810`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into810`
	// (`into810` is now tainted).
	into810, _ := mediumObjCQL.Prepare(query)

	// Sink the tainted `into810`:
	sink(into810)
}

func TaintStepTest_DatabaseSqlDBPrepareContext(sourceCQL interface{}) {
	// The flow is from `query` into `into460`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into460`
	// (`into460` is now tainted).
	into460, _ := mediumObjCQL.PrepareContext(nil, query)

	// Sink the tainted `into460`:
	sink(into460)
}

func TaintStepTest_DatabaseSqlTxPrepare(sourceCQL interface{}) {
	// The flow is from `query` into `into226`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into226`
	// (`into226` is now tainted).
	into226, _ := mediumObjCQL.Prepare(query)

	// Sink the tainted `into226`:
	sink(into226)
}

func TaintStepTest_DatabaseSqlTxPrepareContext(sourceCQL interface{}) {
	// The flow is from `query` into `into877`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into877`
	// (`into877` is now tainted).
	into877, _ := mediumObjCQL.PrepareContext(nil, query)

	// Sink the tainted `into877`:
	sink(into877)
}

func TaintStepTest_DatabaseSqlTxStmt(sourceCQL interface{}) {
	// The flow is from `stmt` into `into193`.

	// Assume that `sourceCQL` has the underlying type of `stmt`:
	stmt := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `stmt` to the result `into193`
	// (`into193` is now tainted).
	into193 := mediumObjCQL.Stmt(stmt)

	// Sink the tainted `into193`:
	sink(into193)
}

func TaintStepTest_DatabaseSqlTxStmtContext(sourceCQL interface{}) {
	// The flow is from `stmt` into `into988`.

	// Assume that `sourceCQL` has the underlying type of `stmt`:
	stmt := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `stmt` to the result `into988`
	// (`into988` is now tainted).
	into988 := mediumObjCQL.StmtContext(nil, stmt)

	// Sink the tainted `into988`:
	sink(into988)
}

func RunAllTaints_DatabaseSql(v interface{}) {
	{
		source := newSource()
		TaintStepTest_DatabaseSqlNamed(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlConnPrepareContext(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDBPrepare(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDBPrepareContext(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxPrepare(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxPrepareContext(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxStmt(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxStmtContext(source)
	}
}
