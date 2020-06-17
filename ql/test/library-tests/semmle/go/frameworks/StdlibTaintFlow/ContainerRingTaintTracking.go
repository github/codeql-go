package main

import "container/ring"

func TaintStepTest_ContainerRingRingLink_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing922` into `intoRing273`.

	// Assume that `sourceCQL` has the underlying type of `fromRing922`:
	fromRing922 := sourceCQL.(*ring.Ring)

	// Declare medium object/interface:
	var mediumObjCQL ring.Ring

	// Call the method that transfers the taint
	// from the parameter `fromRing922` to the result `intoRing273`
	// (`intoRing273` is now tainted).
	intoRing273 := mediumObjCQL.Link(fromRing922)

	// Return the tainted `intoRing273`:
	return intoRing273
}

func TaintStepTest_ContainerRingRingMove_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing241` into `intoRing234`.

	// Assume that `sourceCQL` has the underlying type of `fromRing241`:
	fromRing241 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing241` to the result `intoRing234`
	// (`intoRing234` is now tainted).
	intoRing234 := fromRing241.Move(0)

	// Return the tainted `intoRing234`:
	return intoRing234
}

func TaintStepTest_ContainerRingRingNext_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing882` into `intoRing858`.

	// Assume that `sourceCQL` has the underlying type of `fromRing882`:
	fromRing882 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing882` to the result `intoRing858`
	// (`intoRing858` is now tainted).
	intoRing858 := fromRing882.Next()

	// Return the tainted `intoRing858`:
	return intoRing858
}

func TaintStepTest_ContainerRingRingPrev_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing793` into `intoRing548`.

	// Assume that `sourceCQL` has the underlying type of `fromRing793`:
	fromRing793 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing793` to the result `intoRing548`
	// (`intoRing548` is now tainted).
	intoRing548 := fromRing793.Prev()

	// Return the tainted `intoRing548`:
	return intoRing548
}

func TaintStepTest_ContainerRingRingUnlink_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRing187` into `intoRing554`.

	// Assume that `sourceCQL` has the underlying type of `fromRing187`:
	fromRing187 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing187` to the result `intoRing554`
	// (`intoRing554` is now tainted).
	intoRing554 := fromRing187.Unlink(0)

	// Return the tainted `intoRing554`:
	return intoRing554
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
