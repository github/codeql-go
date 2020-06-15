package main

import "database/sql"

func TaintStepTest_DatabaseSqlNamed_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString695` into `intoNamedArg843`.

	// Assume that `sourceCQL` has the underlying type of `fromString695`:
	fromString695 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString695` to result `intoNamedArg843`
	// (`intoNamedArg843` is now tainted).
	intoNamedArg843 := sql.Named(fromString695, nil)

	// Sink the tainted `intoNamedArg843`:
	sink(intoNamedArg843)
}

func TaintStepTest_DatabaseSqlNamed_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface304` into `intoNamedArg991`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface304`:
	fromInterface304 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface304` to result `intoNamedArg991`
	// (`intoNamedArg991` is now tainted).
	intoNamedArg991 := sql.Named("", fromInterface304)

	// Sink the tainted `intoNamedArg991`:
	sink(intoNamedArg991)
}

func TaintStepTest_DatabaseSqlConnPrepareContext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString949` into `intoStmt367`.

	// Assume that `sourceCQL` has the underlying type of `fromString949`:
	fromString949 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString949` to the result `intoStmt367`
	// (`intoStmt367` is now tainted).
	intoStmt367, _ := mediumObjCQL.PrepareContext(nil, fromString949)

	// Sink the tainted `intoStmt367`:
	sink(intoStmt367)
}

func TaintStepTest_DatabaseSqlDBPrepare_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString352` into `intoStmt789`.

	// Assume that `sourceCQL` has the underlying type of `fromString352`:
	fromString352 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString352` to the result `intoStmt789`
	// (`intoStmt789` is now tainted).
	intoStmt789, _ := mediumObjCQL.Prepare(fromString352)

	// Sink the tainted `intoStmt789`:
	sink(intoStmt789)
}

func TaintStepTest_DatabaseSqlDBPrepareContext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString383` into `intoStmt820`.

	// Assume that `sourceCQL` has the underlying type of `fromString383`:
	fromString383 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.DB

	// Call the method that transfers the taint
	// from the parameter `fromString383` to the result `intoStmt820`
	// (`intoStmt820` is now tainted).
	intoStmt820, _ := mediumObjCQL.PrepareContext(nil, fromString383)

	// Sink the tainted `intoStmt820`:
	sink(intoStmt820)
}

func TaintStepTest_DatabaseSqlTxPrepare_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString783` into `intoStmt592`.

	// Assume that `sourceCQL` has the underlying type of `fromString783`:
	fromString783 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString783` to the result `intoStmt592`
	// (`intoStmt592` is now tainted).
	intoStmt592, _ := mediumObjCQL.Prepare(fromString783)

	// Sink the tainted `intoStmt592`:
	sink(intoStmt592)
}

func TaintStepTest_DatabaseSqlTxPrepareContext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString498` into `intoStmt364`.

	// Assume that `sourceCQL` has the underlying type of `fromString498`:
	fromString498 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromString498` to the result `intoStmt364`
	// (`intoStmt364` is now tainted).
	intoStmt364, _ := mediumObjCQL.PrepareContext(nil, fromString498)

	// Sink the tainted `intoStmt364`:
	sink(intoStmt364)
}

func TaintStepTest_DatabaseSqlTxStmt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStmt441` into `intoStmt114`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt441`:
	fromStmt441 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt441` to the result `intoStmt114`
	// (`intoStmt114` is now tainted).
	intoStmt114 := mediumObjCQL.Stmt(fromStmt441)

	// Sink the tainted `intoStmt114`:
	sink(intoStmt114)
}

func TaintStepTest_DatabaseSqlTxStmtContext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromStmt795` into `intoStmt461`.

	// Assume that `sourceCQL` has the underlying type of `fromStmt795`:
	fromStmt795 := sourceCQL.(*sql.Stmt)

	// Declare medium object/interface:
	var mediumObjCQL sql.Tx

	// Call the method that transfers the taint
	// from the parameter `fromStmt795` to the result `intoStmt461`
	// (`intoStmt461` is now tainted).
	intoStmt461 := mediumObjCQL.StmtContext(nil, fromStmt795)

	// Sink the tainted `intoStmt461`:
	sink(intoStmt461)
}

func RunAllTaints_DatabaseSql() {
	{
		source := newSource()
		TaintStepTest_DatabaseSqlNamed_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlNamed_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlConnPrepareContext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDBPrepare_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDBPrepareContext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxPrepare_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxPrepareContext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxStmt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlTxStmtContext_B0I0O0(source)
	}
}
