// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "container/list"

func TaintStepTest_ContainerListListBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList656` into `intoElement414`.

	// Assume that `sourceCQL` has the underlying type of `fromList656`:
	fromList656 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList656` to the result `intoElement414`
	// (`intoElement414` is now tainted).
	intoElement414 := fromList656.Back()

	// Return the tainted `intoElement414`:
	return intoElement414
}

func TaintStepTest_ContainerListListFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList518` into `intoElement650`.

	// Assume that `sourceCQL` has the underlying type of `fromList518`:
	fromList518 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList518` to the result `intoElement650`
	// (`intoElement650` is now tainted).
	intoElement650 := fromList518.Front()

	// Return the tainted `intoElement650`:
	return intoElement650
}

func TaintStepTest_ContainerListListInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList784` into `intoList957`.

	// Assume that `sourceCQL` has the underlying type of `fromList784`:
	fromList784 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList784` to the result `intoList957`
	// (`intoList957` is now tainted).
	intoList957 := fromList784.Init()

	// Return the tainted `intoList957`:
	return intoList957
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface520` into `intoList443`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface520`:
	fromInterface520 := sourceCQL.(interface{})

	// Declare `intoList443` variable:
	var intoList443 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface520` to the receiver `intoList443`
	// (`intoList443` is now tainted).
	intoList443.InsertAfter(fromInterface520, nil)

	// Return the tainted `intoList443`:
	return intoList443
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface127` into `intoElement483`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface127`:
	fromInterface127 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface127` to the result `intoElement483`
	// (`intoElement483` is now tainted).
	intoElement483 := mediumObjCQL.InsertAfter(fromInterface127, nil)

	// Return the tainted `intoElement483`:
	return intoElement483
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoList982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Declare `intoList982` variable:
	var intoList982 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface989` to the receiver `intoList982`
	// (`intoList982` is now tainted).
	intoList982.InsertBefore(fromInterface989, nil)

	// Return the tainted `intoList982`:
	return intoList982
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface417` into `intoElement584`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface417`:
	fromInterface417 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface417` to the result `intoElement584`
	// (`intoElement584` is now tainted).
	intoElement584 := mediumObjCQL.InsertBefore(fromInterface417, nil)

	// Return the tainted `intoElement584`:
	return intoElement584
}

func TaintStepTest_ContainerListListMoveAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement991` into `intoList881`.

	// Assume that `sourceCQL` has the underlying type of `fromElement991`:
	fromElement991 := sourceCQL.(*list.Element)

	// Declare `intoList881` variable:
	var intoList881 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement991` to the receiver `intoList881`
	// (`intoList881` is now tainted).
	intoList881.MoveAfter(fromElement991, nil)

	// Return the tainted `intoList881`:
	return intoList881
}

func TaintStepTest_ContainerListListMoveBefore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement186` into `intoList284`.

	// Assume that `sourceCQL` has the underlying type of `fromElement186`:
	fromElement186 := sourceCQL.(*list.Element)

	// Declare `intoList284` variable:
	var intoList284 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement186` to the receiver `intoList284`
	// (`intoList284` is now tainted).
	intoList284.MoveBefore(fromElement186, nil)

	// Return the tainted `intoList284`:
	return intoList284
}

func TaintStepTest_ContainerListListMoveToBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement908` into `intoList137`.

	// Assume that `sourceCQL` has the underlying type of `fromElement908`:
	fromElement908 := sourceCQL.(*list.Element)

	// Declare `intoList137` variable:
	var intoList137 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement908` to the receiver `intoList137`
	// (`intoList137` is now tainted).
	intoList137.MoveToBack(fromElement908)

	// Return the tainted `intoList137`:
	return intoList137
}

func TaintStepTest_ContainerListListMoveToFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement494` into `intoList873`.

	// Assume that `sourceCQL` has the underlying type of `fromElement494`:
	fromElement494 := sourceCQL.(*list.Element)

	// Declare `intoList873` variable:
	var intoList873 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement494` to the receiver `intoList873`
	// (`intoList873` is now tainted).
	intoList873.MoveToFront(fromElement494)

	// Return the tainted `intoList873`:
	return intoList873
}

func TaintStepTest_ContainerListListPushBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface599` into `intoList409`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface599`:
	fromInterface599 := sourceCQL.(interface{})

	// Declare `intoList409` variable:
	var intoList409 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface599` to the receiver `intoList409`
	// (`intoList409` is now tainted).
	intoList409.PushBack(fromInterface599)

	// Return the tainted `intoList409`:
	return intoList409
}

func TaintStepTest_ContainerListListPushBack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface246` into `intoElement898`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface246`:
	fromInterface246 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface246` to the result `intoElement898`
	// (`intoElement898` is now tainted).
	intoElement898 := mediumObjCQL.PushBack(fromInterface246)

	// Return the tainted `intoElement898`:
	return intoElement898
}

func TaintStepTest_ContainerListListPushBackList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList598` into `intoList631`.

	// Assume that `sourceCQL` has the underlying type of `fromList598`:
	fromList598 := sourceCQL.(*list.List)

	// Declare `intoList631` variable:
	var intoList631 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList598` to the receiver `intoList631`
	// (`intoList631` is now tainted).
	intoList631.PushBackList(fromList598)

	// Return the tainted `intoList631`:
	return intoList631
}

func TaintStepTest_ContainerListListPushFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface165` into `intoList150`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface165`:
	fromInterface165 := sourceCQL.(interface{})

	// Declare `intoList150` variable:
	var intoList150 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface165` to the receiver `intoList150`
	// (`intoList150` is now tainted).
	intoList150.PushFront(fromInterface165)

	// Return the tainted `intoList150`:
	return intoList150
}

func TaintStepTest_ContainerListListPushFront_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface340` into `intoElement471`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface340`:
	fromInterface340 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface340` to the result `intoElement471`
	// (`intoElement471` is now tainted).
	intoElement471 := mediumObjCQL.PushFront(fromInterface340)

	// Return the tainted `intoElement471`:
	return intoElement471
}

func TaintStepTest_ContainerListListPushFrontList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList290` into `intoList758`.

	// Assume that `sourceCQL` has the underlying type of `fromList290`:
	fromList290 := sourceCQL.(*list.List)

	// Declare `intoList758` variable:
	var intoList758 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList290` to the receiver `intoList758`
	// (`intoList758` is now tainted).
	intoList758.PushFrontList(fromList290)

	// Return the tainted `intoList758`:
	return intoList758
}

func TaintStepTest_ContainerListListRemove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement396` into `intoInterface707`.

	// Assume that `sourceCQL` has the underlying type of `fromElement396`:
	fromElement396 := sourceCQL.(*list.Element)

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement396` to the result `intoInterface707`
	// (`intoInterface707` is now tainted).
	intoInterface707 := mediumObjCQL.Remove(fromElement396)

	// Return the tainted `intoInterface707`:
	return intoInterface707
}

func RunAllTaints_ContainerList() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListBack_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListFront_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListInit_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListInsertAfter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListInsertAfter_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListInsertBefore_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListInsertBefore_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListMoveAfter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListMoveBefore_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListMoveToBack_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListMoveToFront_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListPushBack_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListPushBack_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListPushBackList_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListPushFront_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListPushFront_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListPushFrontList_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerListListRemove_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
