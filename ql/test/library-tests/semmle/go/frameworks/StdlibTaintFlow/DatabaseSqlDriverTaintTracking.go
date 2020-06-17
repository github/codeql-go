// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "database/sql/driver"

func TaintStepTest_DatabaseSqlDriverNotNullConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface656` into `intoValue414`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface656`:
	fromInterface656 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.NotNull

	// Call the method that transfers the taint
	// from the parameter `fromInterface656` to the result `intoValue414`
	// (`intoValue414` is now tainted).
	intoValue414, _ := mediumObjCQL.ConvertValue(fromInterface656)

	// Return the tainted `intoValue414`:
	return intoValue414
}

func TaintStepTest_DatabaseSqlDriverNullConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface518` into `intoValue650`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface518`:
	fromInterface518 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.Null

	// Call the method that transfers the taint
	// from the parameter `fromInterface518` to the result `intoValue650`
	// (`intoValue650` is now tainted).
	intoValue650, _ := mediumObjCQL.ConvertValue(fromInterface518)

	// Return the tainted `intoValue650`:
	return intoValue650
}

func TaintStepTest_DatabaseSqlDriverValueConverterConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface784` into `intoValue957`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface784`:
	fromInterface784 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.ValueConverter

	// Call the method that transfers the taint
	// from the parameter `fromInterface784` to the result `intoValue957`
	// (`intoValue957` is now tainted).
	intoValue957, _ := mediumObjCQL.ConvertValue(fromInterface784)

	// Return the tainted `intoValue957`:
	return intoValue957
}

func TaintStepTest_DatabaseSqlDriverConnPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString520` into `intoStmt443`.

	// Assume that `sourceCQL` has the underlying type of `fromString520`:
	fromString520 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString520` to the result `intoStmt443`
	// (`intoStmt443` is now tainted).
	intoStmt443, _ := mediumObjCQL.Prepare(fromString520)

	// Return the tainted `intoStmt443`:
	return intoStmt443
}

func TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoStmt483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.ConnPrepareContext

	// Call the method that transfers the taint
	// from the parameter `fromString127` to the result `intoStmt483`
	// (`intoStmt483` is now tainted).
	intoStmt483, _ := mediumObjCQL.PrepareContext(nil, fromString127)

	// Return the tainted `intoStmt483`:
	return intoStmt483
}

func TaintStepTest_DatabaseSqlDriverValuerValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValuer989` into `intoValue982`.

	// Assume that `sourceCQL` has the underlying type of `fromValuer989`:
	fromValuer989 := sourceCQL.(driver.Valuer)

	// Call the method that transfers the taint
	// from the receiver `fromValuer989` to the result `intoValue982`
	// (`intoValue982` is now tainted).
	intoValue982, _ := fromValuer989.Value()

	// Return the tainted `intoValue982`:
	return intoValue982
}

func RunAllTaints_DatabaseSqlDriver() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDriverNotNullConvertValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDriverNullConvertValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDriverValueConverterConvertValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDriverConnPrepare_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_DatabaseSqlDriverValuerValue_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
