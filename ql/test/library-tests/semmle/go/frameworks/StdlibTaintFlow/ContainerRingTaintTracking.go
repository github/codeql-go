// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "container/ring"

func TaintStepTest_ContainerRingRingLink_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing750` into `intoRing817`.

	// Assume that `sourceCQL` has the underlying type of `fromRing750`:
	fromRing750 := sourceCQL.(*ring.Ring)

	// Declare medium object/interface:
	var mediumObjCQL ring.Ring

	// Call the method that transfers the taint
	// from the parameter `fromRing750` to the result `intoRing817`
	// (`intoRing817` is now tainted).
	intoRing817 := mediumObjCQL.Link(fromRing750)

	// Return the tainted `intoRing817`:
	return intoRing817
}

func TaintStepTest_ContainerRingRingMove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing659` into `intoRing232`.

	// Assume that `sourceCQL` has the underlying type of `fromRing659`:
	fromRing659 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing659` to the result `intoRing232`
	// (`intoRing232` is now tainted).
	intoRing232 := fromRing659.Move(0)

	// Return the tainted `intoRing232`:
	return intoRing232
}

func TaintStepTest_ContainerRingRingNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing810` into `intoRing501`.

	// Assume that `sourceCQL` has the underlying type of `fromRing810`:
	fromRing810 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing810` to the result `intoRing501`
	// (`intoRing501` is now tainted).
	intoRing501 := fromRing810.Next()

	// Return the tainted `intoRing501`:
	return intoRing501
}

func TaintStepTest_ContainerRingRingPrev_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing438` into `intoRing858`.

	// Assume that `sourceCQL` has the underlying type of `fromRing438`:
	fromRing438 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing438` to the result `intoRing858`
	// (`intoRing858` is now tainted).
	intoRing858 := fromRing438.Prev()

	// Return the tainted `intoRing858`:
	return intoRing858
}

func TaintStepTest_ContainerRingRingUnlink_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing546` into `intoRing364`.

	// Assume that `sourceCQL` has the underlying type of `fromRing546`:
	fromRing546 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing546` to the result `intoRing364`
	// (`intoRing364` is now tainted).
	intoRing364 := fromRing546.Unlink(0)

	// Return the tainted `intoRing364`:
	return intoRing364
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
