package main

import "database/sql/driver"

func TaintStepTest_DatabaseSqlDriverNotNullConvertValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface739` into `intoValue461`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface739`:
	fromInterface739 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.NotNull

	// Call the method that transfers the taint
	// from the parameter `fromInterface739` to the result `intoValue461`
	// (`intoValue461` is now tainted).
	intoValue461, _ := mediumObjCQL.ConvertValue(fromInterface739)

	// Sink the tainted `intoValue461`:
	sink(intoValue461)
}

func TaintStepTest_DatabaseSqlDriverNullConvertValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface849` into `intoValue310`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface849`:
	fromInterface849 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.Null

	// Call the method that transfers the taint
	// from the parameter `fromInterface849` to the result `intoValue310`
	// (`intoValue310` is now tainted).
	intoValue310, _ := mediumObjCQL.ConvertValue(fromInterface849)

	// Sink the tainted `intoValue310`:
	sink(intoValue310)
}

func TaintStepTest_DatabaseSqlDriverValueConverterConvertValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface234` into `intoValue778`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface234`:
	fromInterface234 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.ValueConverter

	// Call the method that transfers the taint
	// from the parameter `fromInterface234` to the result `intoValue778`
	// (`intoValue778` is now tainted).
	intoValue778, _ := mediumObjCQL.ConvertValue(fromInterface234)

	// Sink the tainted `intoValue778`:
	sink(intoValue778)
}

func TaintStepTest_DatabaseSqlDriverConnPrepare_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString479` into `intoStmt958`.

	// Assume that `sourceCQL` has the underlying type of `fromString479`:
	fromString479 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString479` to the result `intoStmt958`
	// (`intoStmt958` is now tainted).
	intoStmt958, _ := mediumObjCQL.Prepare(fromString479)

	// Sink the tainted `intoStmt958`:
	sink(intoStmt958)
}

func TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString368` into `intoStmt853`.

	// Assume that `sourceCQL` has the underlying type of `fromString368`:
	fromString368 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.ConnPrepareContext

	// Call the method that transfers the taint
	// from the parameter `fromString368` to the result `intoStmt853`
	// (`intoStmt853` is now tainted).
	intoStmt853, _ := mediumObjCQL.PrepareContext(nil, fromString368)

	// Sink the tainted `intoStmt853`:
	sink(intoStmt853)
}

func TaintStepTest_DatabaseSqlDriverValuerValue_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromValuer463` into `intoValue165`.

	// Assume that `sourceCQL` has the underlying type of `fromValuer463`:
	fromValuer463 := sourceCQL.(driver.Valuer)

	// Call the method that transfers the taint
	// from the receiver `fromValuer463` to the result `intoValue165`
	// (`intoValue165` is now tainted).
	intoValue165, _ := fromValuer463.Value()

	// Sink the tainted `intoValue165`:
	sink(intoValue165)
}

func RunAllTaints_DatabaseSqlDriver() {
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverNotNullConvertValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverNullConvertValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverValueConverterConvertValue_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverConnPrepare_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverValuerValue_B0I0O0(source)
	}
}
