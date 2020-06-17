// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "database/sql/driver"

func TaintStepTest_DatabaseSqlDriverNotNullConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface887` into `intoValue805`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface887`:
	fromInterface887 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.NotNull

	// Call the method that transfers the taint
	// from the parameter `fromInterface887` to the result `intoValue805`
	// (`intoValue805` is now tainted).
	intoValue805, _ := mediumObjCQL.ConvertValue(fromInterface887)

	// Return the tainted `intoValue805`:
	return intoValue805
}

func TaintStepTest_DatabaseSqlDriverNullConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface523` into `intoValue237`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface523`:
	fromInterface523 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.Null

	// Call the method that transfers the taint
	// from the parameter `fromInterface523` to the result `intoValue237`
	// (`intoValue237` is now tainted).
	intoValue237, _ := mediumObjCQL.ConvertValue(fromInterface523)

	// Return the tainted `intoValue237`:
	return intoValue237
}

func TaintStepTest_DatabaseSqlDriverValueConverterConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface304` into `intoValue989`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface304`:
	fromInterface304 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.ValueConverter

	// Call the method that transfers the taint
	// from the parameter `fromInterface304` to the result `intoValue989`
	// (`intoValue989` is now tainted).
	intoValue989, _ := mediumObjCQL.ConvertValue(fromInterface304)

	// Return the tainted `intoValue989`:
	return intoValue989
}

func TaintStepTest_DatabaseSqlDriverConnPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString370` into `intoStmt823`.

	// Assume that `sourceCQL` has the underlying type of `fromString370`:
	fromString370 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString370` to the result `intoStmt823`
	// (`intoStmt823` is now tainted).
	intoStmt823, _ := mediumObjCQL.Prepare(fromString370)

	// Return the tainted `intoStmt823`:
	return intoStmt823
}

func TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString945` into `intoStmt848`.

	// Assume that `sourceCQL` has the underlying type of `fromString945`:
	fromString945 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.ConnPrepareContext

	// Call the method that transfers the taint
	// from the parameter `fromString945` to the result `intoStmt848`
	// (`intoStmt848` is now tainted).
	intoStmt848, _ := mediumObjCQL.PrepareContext(nil, fromString945)

	// Return the tainted `intoStmt848`:
	return intoStmt848
}

func TaintStepTest_DatabaseSqlDriverValuerValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValuer919` into `intoValue365`.

	// Assume that `sourceCQL` has the underlying type of `fromValuer919`:
	fromValuer919 := sourceCQL.(driver.Valuer)

	// Call the method that transfers the taint
	// from the receiver `fromValuer919` to the result `intoValue365`
	// (`intoValue365` is now tainted).
	intoValue365, _ := fromValuer919.Value()

	// Return the tainted `intoValue365`:
	return intoValue365
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
