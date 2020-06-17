package main

import "container/list"

func TaintStepTest_ContainerListListBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList827` into `intoElement134`.

	// Assume that `sourceCQL` has the underlying type of `fromList827`:
	fromList827 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList827` to the result `intoElement134`
	// (`intoElement134` is now tainted).
	intoElement134 := fromList827.Back()

	// Return the tainted `intoElement134`:
	return intoElement134
}

func TaintStepTest_ContainerListListFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList801` into `intoElement795`.

	// Assume that `sourceCQL` has the underlying type of `fromList801`:
	fromList801 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList801` to the result `intoElement795`
	// (`intoElement795` is now tainted).
	intoElement795 := fromList801.Front()

	// Return the tainted `intoElement795`:
	return intoElement795
}

func TaintStepTest_ContainerListListInit_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList179` into `intoList744`.

	// Assume that `sourceCQL` has the underlying type of `fromList179`:
	fromList179 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList179` to the result `intoList744`
	// (`intoList744` is now tainted).
	intoList744 := fromList179.Init()

	// Return the tainted `intoList744`:
	return intoList744
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface353` into `intoList333`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface353`:
	fromInterface353 := sourceCQL.(interface{})

	// Declare `intoList333` variable:
	var intoList333 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface353` to the receiver `intoList333`
	// (`intoList333` is now tainted).
	intoList333.InsertAfter(fromInterface353, nil)

	// Return the tainted `intoList333`:
	return intoList333
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface231` into `intoElement539`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface231`:
	fromInterface231 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface231` to the result `intoElement539`
	// (`intoElement539` is now tainted).
	intoElement539 := mediumObjCQL.InsertAfter(fromInterface231, nil)

	// Return the tainted `intoElement539`:
	return intoElement539
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface319` into `intoList649`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface319`:
	fromInterface319 := sourceCQL.(interface{})

	// Declare `intoList649` variable:
	var intoList649 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface319` to the receiver `intoList649`
	// (`intoList649` is now tainted).
	intoList649.InsertBefore(fromInterface319, nil)

	// Return the tainted `intoList649`:
	return intoList649
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface974` into `intoElement168`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface974`:
	fromInterface974 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface974` to the result `intoElement168`
	// (`intoElement168` is now tainted).
	intoElement168 := mediumObjCQL.InsertBefore(fromInterface974, nil)

	// Return the tainted `intoElement168`:
	return intoElement168
}

func TaintStepTest_ContainerListListMoveAfter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement234` into `intoList729`.

	// Assume that `sourceCQL` has the underlying type of `fromElement234`:
	fromElement234 := sourceCQL.(*list.Element)

	// Declare `intoList729` variable:
	var intoList729 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement234` to the receiver `intoList729`
	// (`intoList729` is now tainted).
	intoList729.MoveAfter(fromElement234, nil)

	// Return the tainted `intoList729`:
	return intoList729
}

func TaintStepTest_ContainerListListMoveBefore_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement447` into `intoList203`.

	// Assume that `sourceCQL` has the underlying type of `fromElement447`:
	fromElement447 := sourceCQL.(*list.Element)

	// Declare `intoList203` variable:
	var intoList203 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement447` to the receiver `intoList203`
	// (`intoList203` is now tainted).
	intoList203.MoveBefore(fromElement447, nil)

	// Return the tainted `intoList203`:
	return intoList203
}

func TaintStepTest_ContainerListListMoveToBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement803` into `intoList858`.

	// Assume that `sourceCQL` has the underlying type of `fromElement803`:
	fromElement803 := sourceCQL.(*list.Element)

	// Declare `intoList858` variable:
	var intoList858 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement803` to the receiver `intoList858`
	// (`intoList858` is now tainted).
	intoList858.MoveToBack(fromElement803)

	// Return the tainted `intoList858`:
	return intoList858
}

func TaintStepTest_ContainerListListMoveToFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement908` into `intoList483`.

	// Assume that `sourceCQL` has the underlying type of `fromElement908`:
	fromElement908 := sourceCQL.(*list.Element)

	// Declare `intoList483` variable:
	var intoList483 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement908` to the receiver `intoList483`
	// (`intoList483` is now tainted).
	intoList483.MoveToFront(fromElement908)

	// Return the tainted `intoList483`:
	return intoList483
}

func TaintStepTest_ContainerListListPushBack_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface714` into `intoList520`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface714`:
	fromInterface714 := sourceCQL.(interface{})

	// Declare `intoList520` variable:
	var intoList520 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface714` to the receiver `intoList520`
	// (`intoList520` is now tainted).
	intoList520.PushBack(fromInterface714)

	// Return the tainted `intoList520`:
	return intoList520
}

func TaintStepTest_ContainerListListPushBack_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface984` into `intoElement885`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface984`:
	fromInterface984 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface984` to the result `intoElement885`
	// (`intoElement885` is now tainted).
	intoElement885 := mediumObjCQL.PushBack(fromInterface984)

	// Return the tainted `intoElement885`:
	return intoElement885
}

func TaintStepTest_ContainerListListPushBackList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList160` into `intoList597`.

	// Assume that `sourceCQL` has the underlying type of `fromList160`:
	fromList160 := sourceCQL.(*list.List)

	// Declare `intoList597` variable:
	var intoList597 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList160` to the receiver `intoList597`
	// (`intoList597` is now tainted).
	intoList597.PushBackList(fromList160)

	// Return the tainted `intoList597`:
	return intoList597
}

func TaintStepTest_ContainerListListPushFront_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface190` into `intoList311`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface190`:
	fromInterface190 := sourceCQL.(interface{})

	// Declare `intoList311` variable:
	var intoList311 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface190` to the receiver `intoList311`
	// (`intoList311` is now tainted).
	intoList311.PushFront(fromInterface190)

	// Return the tainted `intoList311`:
	return intoList311
}

func TaintStepTest_ContainerListListPushFront_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface504` into `intoElement993`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface504`:
	fromInterface504 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface504` to the result `intoElement993`
	// (`intoElement993` is now tainted).
	intoElement993 := mediumObjCQL.PushFront(fromInterface504)

	// Return the tainted `intoElement993`:
	return intoElement993
}

func TaintStepTest_ContainerListListPushFrontList_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromList368` into `intoList586`.

	// Assume that `sourceCQL` has the underlying type of `fromList368`:
	fromList368 := sourceCQL.(*list.List)

	// Declare `intoList586` variable:
	var intoList586 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList368` to the receiver `intoList586`
	// (`intoList586` is now tainted).
	intoList586.PushFrontList(fromList368)

	// Return the tainted `intoList586`:
	return intoList586
}

func TaintStepTest_ContainerListListRemove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromElement859` into `intoInterface679`.

	// Assume that `sourceCQL` has the underlying type of `fromElement859`:
	fromElement859 := sourceCQL.(*list.Element)

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement859` to the result `intoInterface679`
	// (`intoInterface679` is now tainted).
	intoInterface679 := mediumObjCQL.Remove(fromElement859)

	// Return the tainted `intoInterface679`:
	return intoInterface679
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
