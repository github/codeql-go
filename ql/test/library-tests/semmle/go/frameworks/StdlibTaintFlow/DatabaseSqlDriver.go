package main

import "database/sql/driver"

func TaintStepTest_DatabaseSqlDriverNotNullConvertValue(sourceCQL interface{}) {
	// The flow is from `v` into `into871`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.NotNull

	// Call the method that transfers the taint
	// from the parameter `v` to the result `into871`
	// (`into871` is now tainted).
	into871, _ := mediumObjCQL.ConvertValue(v)

	// Sink the tainted `into871`:
	sink(into871)
}

func TaintStepTest_DatabaseSqlDriverNullConvertValue(sourceCQL interface{}) {
	// The flow is from `v` into `into899`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.Null

	// Call the method that transfers the taint
	// from the parameter `v` to the result `into899`
	// (`into899` is now tainted).
	into899, _ := mediumObjCQL.ConvertValue(v)

	// Sink the tainted `into899`:
	sink(into899)
}

func TaintStepTest_DatabaseSqlDriverValueConverterConvertValue(sourceCQL interface{}) {
	// The flow is from `v` into `into548`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL driver.ValueConverter

	// Call the method that transfers the taint
	// from the parameter `v` to the result `into548`
	// (`into548` is now tainted).
	into548, _ := mediumObjCQL.ConvertValue(v)

	// Sink the tainted `into548`:
	sink(into548)
}

func TaintStepTest_DatabaseSqlDriverConnPrepare(sourceCQL interface{}) {
	// The flow is from `query` into `into295`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.Conn

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into295`
	// (`into295` is now tainted).
	into295, _ := mediumObjCQL.Prepare(query)

	// Sink the tainted `into295`:
	sink(into295)
}

func TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext(sourceCQL interface{}) {
	// The flow is from `query` into `into562`.

	// Assume that `sourceCQL` has the underlying type of `query`:
	query := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL driver.ConnPrepareContext

	// Call the method that transfers the taint
	// from the parameter `query` to the result `into562`
	// (`into562` is now tainted).
	into562, _ := mediumObjCQL.PrepareContext(nil, query)

	// Sink the tainted `into562`:
	sink(into562)
}

func TaintStepTest_DatabaseSqlDriverValuerValue(sourceCQL interface{}) {
	// The flow is from `from294` into `into320`.

	// Assume that `sourceCQL` has the underlying type of `from294`:
	from294 := sourceCQL.(driver.Valuer)

	// Call the method that transfers the taint
	// from the receiver `from294` to the result `into320`
	// (`into320` is now tainted).
	into320, _ := from294.Value()

	// Sink the tainted `into320`:
	sink(into320)
}

func RunAllTaints_DatabaseSqlDriver(v interface{}) {
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverNotNullConvertValue(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverNullConvertValue(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverValueConverterConvertValue(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverConnPrepare(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverConnPrepareContextPrepareContext(source)
	}
	{
		source := newSource()
		TaintStepTest_DatabaseSqlDriverValuerValue(source)
	}
}
