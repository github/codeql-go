// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "container/ring"

func TaintStepTest_ContainerRingRingLink_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing656` into `intoRing414`.

	// Assume that `sourceCQL` has the underlying type of `fromRing656`:
	fromRing656 := sourceCQL.(*ring.Ring)

	// Declare medium object/interface:
	var mediumObjCQL ring.Ring

	// Call the method that transfers the taint
	// from the parameter `fromRing656` to the result `intoRing414`
	// (`intoRing414` is now tainted).
	intoRing414 := mediumObjCQL.Link(fromRing656)

	// Return the tainted `intoRing414`:
	return intoRing414
}

func TaintStepTest_ContainerRingRingMove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing518` into `intoRing650`.

	// Assume that `sourceCQL` has the underlying type of `fromRing518`:
	fromRing518 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing518` to the result `intoRing650`
	// (`intoRing650` is now tainted).
	intoRing650 := fromRing518.Move(0)

	// Return the tainted `intoRing650`:
	return intoRing650
}

func TaintStepTest_ContainerRingRingNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing784` into `intoRing957`.

	// Assume that `sourceCQL` has the underlying type of `fromRing784`:
	fromRing784 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing784` to the result `intoRing957`
	// (`intoRing957` is now tainted).
	intoRing957 := fromRing784.Next()

	// Return the tainted `intoRing957`:
	return intoRing957
}

func TaintStepTest_ContainerRingRingPrev_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing520` into `intoRing443`.

	// Assume that `sourceCQL` has the underlying type of `fromRing520`:
	fromRing520 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing520` to the result `intoRing443`
	// (`intoRing443` is now tainted).
	intoRing443 := fromRing520.Prev()

	// Return the tainted `intoRing443`:
	return intoRing443
}

func TaintStepTest_ContainerRingRingUnlink_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing127` into `intoRing483`.

	// Assume that `sourceCQL` has the underlying type of `fromRing127`:
	fromRing127 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing127` to the result `intoRing483`
	// (`intoRing483` is now tainted).
	intoRing483 := fromRing127.Unlink(0)

	// Return the tainted `intoRing483`:
	return intoRing483
}

func RunAllTaints_ContainerRing() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerRingRingLink_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerRingRingMove_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerRingRingNext_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerRingRingPrev_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_ContainerRingRingUnlink_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
