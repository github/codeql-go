// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "container/list"

func TaintStepTest_ContainerListListBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList554` into `intoElement692`.

	// Assume that `sourceCQL` has the underlying type of `fromList554`:
	fromList554 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList554` to the result `intoElement692`
	// (`intoElement692` is now tainted).
	intoElement692 := fromList554.Back()

	// Return the tainted `intoElement692`:
	return intoElement692
}

func TaintStepTest_ContainerListListFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList174` into `intoElement853`.

	// Assume that `sourceCQL` has the underlying type of `fromList174`:
	fromList174 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList174` to the result `intoElement853`
	// (`intoElement853` is now tainted).
	intoElement853 := fromList174.Front()

	// Return the tainted `intoElement853`:
	return intoElement853
}

func TaintStepTest_ContainerListListInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList299` into `intoList852`.

	// Assume that `sourceCQL` has the underlying type of `fromList299`:
	fromList299 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList299` to the result `intoList852`
	// (`intoList852` is now tainted).
	intoList852 := fromList299.Init()

	// Return the tainted `intoList852`:
	return intoList852
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface174` into `intoList489`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface174`:
	fromInterface174 := sourceCQL.(interface{})

	// Declare `intoList489` variable:
	var intoList489 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface174` to the receiver `intoList489`
	// (`intoList489` is now tainted).
	intoList489.InsertAfter(fromInterface174, nil)

	// Return the tainted `intoList489`:
	return intoList489
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface560` into `intoElement689`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface560`:
	fromInterface560 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface560` to the result `intoElement689`
	// (`intoElement689` is now tainted).
	intoElement689 := mediumObjCQL.InsertAfter(fromInterface560, nil)

	// Return the tainted `intoElement689`:
	return intoElement689
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface648` into `intoList519`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface648`:
	fromInterface648 := sourceCQL.(interface{})

	// Declare `intoList519` variable:
	var intoList519 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface648` to the receiver `intoList519`
	// (`intoList519` is now tainted).
	intoList519.InsertBefore(fromInterface648, nil)

	// Return the tainted `intoList519`:
	return intoList519
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface465` into `intoElement723`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface465`:
	fromInterface465 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface465` to the result `intoElement723`
	// (`intoElement723` is now tainted).
	intoElement723 := mediumObjCQL.InsertBefore(fromInterface465, nil)

	// Return the tainted `intoElement723`:
	return intoElement723
}

func TaintStepTest_ContainerListListMoveAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement596` into `intoList315`.

	// Assume that `sourceCQL` has the underlying type of `fromElement596`:
	fromElement596 := sourceCQL.(*list.Element)

	// Declare `intoList315` variable:
	var intoList315 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement596` to the receiver `intoList315`
	// (`intoList315` is now tainted).
	intoList315.MoveAfter(fromElement596, nil)

	// Return the tainted `intoList315`:
	return intoList315
}

func TaintStepTest_ContainerListListMoveBefore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement616` into `intoList613`.

	// Assume that `sourceCQL` has the underlying type of `fromElement616`:
	fromElement616 := sourceCQL.(*list.Element)

	// Declare `intoList613` variable:
	var intoList613 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement616` to the receiver `intoList613`
	// (`intoList613` is now tainted).
	intoList613.MoveBefore(fromElement616, nil)

	// Return the tainted `intoList613`:
	return intoList613
}

func TaintStepTest_ContainerListListMoveToBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement915` into `intoList650`.

	// Assume that `sourceCQL` has the underlying type of `fromElement915`:
	fromElement915 := sourceCQL.(*list.Element)

	// Declare `intoList650` variable:
	var intoList650 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement915` to the receiver `intoList650`
	// (`intoList650` is now tainted).
	intoList650.MoveToBack(fromElement915)

	// Return the tainted `intoList650`:
	return intoList650
}

func TaintStepTest_ContainerListListMoveToFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement542` into `intoList277`.

	// Assume that `sourceCQL` has the underlying type of `fromElement542`:
	fromElement542 := sourceCQL.(*list.Element)

	// Declare `intoList277` variable:
	var intoList277 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement542` to the receiver `intoList277`
	// (`intoList277` is now tainted).
	intoList277.MoveToFront(fromElement542)

	// Return the tainted `intoList277`:
	return intoList277
}

func TaintStepTest_ContainerListListPushBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface416` into `intoList594`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface416`:
	fromInterface416 := sourceCQL.(interface{})

	// Declare `intoList594` variable:
	var intoList594 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface416` to the receiver `intoList594`
	// (`intoList594` is now tainted).
	intoList594.PushBack(fromInterface416)

	// Return the tainted `intoList594`:
	return intoList594
}

func TaintStepTest_ContainerListListPushBack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface666` into `intoElement615`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface666`:
	fromInterface666 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface666` to the result `intoElement615`
	// (`intoElement615` is now tainted).
	intoElement615 := mediumObjCQL.PushBack(fromInterface666)

	// Return the tainted `intoElement615`:
	return intoElement615
}

func TaintStepTest_ContainerListListPushBackList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList684` into `intoList741`.

	// Assume that `sourceCQL` has the underlying type of `fromList684`:
	fromList684 := sourceCQL.(*list.List)

	// Declare `intoList741` variable:
	var intoList741 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList684` to the receiver `intoList741`
	// (`intoList741` is now tainted).
	intoList741.PushBackList(fromList684)

	// Return the tainted `intoList741`:
	return intoList741
}

func TaintStepTest_ContainerListListPushFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface129` into `intoList218`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface129`:
	fromInterface129 := sourceCQL.(interface{})

	// Declare `intoList218` variable:
	var intoList218 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface129` to the receiver `intoList218`
	// (`intoList218` is now tainted).
	intoList218.PushFront(fromInterface129)

	// Return the tainted `intoList218`:
	return intoList218
}

func TaintStepTest_ContainerListListPushFront_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface391` into `intoElement401`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface391`:
	fromInterface391 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface391` to the result `intoElement401`
	// (`intoElement401` is now tainted).
	intoElement401 := mediumObjCQL.PushFront(fromInterface391)

	// Return the tainted `intoElement401`:
	return intoElement401
}

func TaintStepTest_ContainerListListPushFrontList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList179` into `intoList313`.

	// Assume that `sourceCQL` has the underlying type of `fromList179`:
	fromList179 := sourceCQL.(*list.List)

	// Declare `intoList313` variable:
	var intoList313 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList179` to the receiver `intoList313`
	// (`intoList313` is now tainted).
	intoList313.PushFrontList(fromList179)

	// Return the tainted `intoList313`:
	return intoList313
}

func TaintStepTest_ContainerListListRemove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement932` into `intoInterface431`.

	// Assume that `sourceCQL` has the underlying type of `fromElement932`:
	fromElement932 := sourceCQL.(*list.Element)

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement932` to the result `intoInterface431`
	// (`intoInterface431` is now tainted).
	intoInterface431 := mediumObjCQL.Remove(fromElement932)

	// Return the tainted `intoInterface431`:
	return intoInterface431
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
