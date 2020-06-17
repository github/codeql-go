package main

import "database/sql/driver"

func TaintStepTest_DatabaseSqlDriverNotNullConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface844` into `intoValue755`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface844`:
	fromInterface844 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.NotNull

	// Call the method that transfers the taint
	// from the parameter `fromInterface844` to the result `intoValue755`
	// (`intoValue755` is now tainted).
	intoValue755, _ := mediumObjCQL.ConvertValue(fromInterface844)

	// Return the tainted `intoValue755`:
	return intoValue755
}

func TaintStepTest_DatabaseSqlDriverNullConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface345` into `intoValue892`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface345`:
	fromInterface345 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.Null

	// Call the method that transfers the taint
	// from the parameter `fromInterface345` to the result `intoValue892`
	// (`intoValue892` is now tainted).
	intoValue892, _ := mediumObjCQL.ConvertValue(fromInterface345)

	// Return the tainted `intoValue892`:
	return intoValue892
}

func TaintStepTest_DatabaseSqlDriverValueConverterConvertValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface427` into `intoValue264`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface427`:
	fromInterface427 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.ValueConverter

	// Call the method that transfers the taint
	// from the parameter `fromInterface427` to the result `intoValue264`
	// (`intoValue264` is now tainted).
	intoValue264, _ := mediumObjCQL.ConvertValue(fromInterface427)

	// Return the tainted `intoValue264`:
	return intoValue264
}

func TaintStepTest_DatabaseSqlDriverConnPrepare_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString304` into `intoStmt792`.

	// Assume that `sourceCQL` has the underlying type of `fromString304`:
	fromString304 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.Conn

	// Call the method that transfers the taint
	// from the parameter `fromString304` to the result `intoStmt792`
	// (`intoStmt792` is now tainted).
	intoStmt792, _ := mediumObjCQL.Prepare(fromString304)

	// Return the tainted `intoStmt792`:
	return intoStmt792
}

func TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString383` into `intoStmt350`.

	// Assume that `sourceCQL` has the underlying type of `fromString383`:
	fromString383 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.ConnPrepareContext

	// Call the method that transfers the taint
	// from the parameter `fromString383` to the result `intoStmt350`
	// (`intoStmt350` is now tainted).
	intoStmt350, _ := mediumObjCQL.PrepareContext(nil, fromString383)

	// Return the tainted `intoStmt350`:
	return intoStmt350
}

func TaintStepTest_DatabaseSqlDriverValuerValue_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromValuer627` into `intoValue742`.

	// Assume that `sourceCQL` has the underlying type of `fromValuer627`:
	fromValuer627 := sourceCQL.(driver.Valuer)

	// Call the method that transfers the taint
	// from the receiver `fromValuer627` to the result `intoValue742`
	// (`intoValue742` is now tainted).
	intoValue742, _ := fromValuer627.Value()

	// Return the tainted `intoValue742`:
	return intoValue742
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
