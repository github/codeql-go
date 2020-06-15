package main

import "container/ring"

func TaintStepTest_ContainerRingRingLink_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRing965` into `intoRing415`.

	// Assume that `sourceCQL` has the underlying type of `fromRing965`:
	fromRing965 := sourceCQL.(*ring.Ring)

	// Declare medium object/interface:
	var mediumObjCQL ring.Ring

	// Call the method that transfers the taint
	// from the parameter `fromRing965` to the result `intoRing415`
	// (`intoRing415` is now tainted).
	intoRing415 := mediumObjCQL.Link(fromRing965)

	// Sink the tainted `intoRing415`:
	sink(intoRing415)
}

func TaintStepTest_ContainerRingRingMove_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRing989` into `intoRing701`.

	// Assume that `sourceCQL` has the underlying type of `fromRing989`:
	fromRing989 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing989` to the result `intoRing701`
	// (`intoRing701` is now tainted).
	intoRing701 := fromRing989.Move(0)

	// Sink the tainted `intoRing701`:
	sink(intoRing701)
}

func TaintStepTest_ContainerRingRingNext_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRing303` into `intoRing732`.

	// Assume that `sourceCQL` has the underlying type of `fromRing303`:
	fromRing303 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing303` to the result `intoRing732`
	// (`intoRing732` is now tainted).
	intoRing732 := fromRing303.Next()

	// Sink the tainted `intoRing732`:
	sink(intoRing732)
}

func TaintStepTest_ContainerRingRingPrev_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRing757` into `intoRing633`.

	// Assume that `sourceCQL` has the underlying type of `fromRing757`:
	fromRing757 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing757` to the result `intoRing633`
	// (`intoRing633` is now tainted).
	intoRing633 := fromRing757.Prev()

	// Sink the tainted `intoRing633`:
	sink(intoRing633)
}

func TaintStepTest_ContainerRingRingUnlink_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRing547` into `intoRing582`.

	// Assume that `sourceCQL` has the underlying type of `fromRing547`:
	fromRing547 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `fromRing547` to the result `intoRing582`
	// (`intoRing582` is now tainted).
	intoRing582 := fromRing547.Unlink(0)

	// Sink the tainted `intoRing582`:
	sink(intoRing582)
}

func RunAllTaints_ContainerRing() {
	{
		source := newSource()
		TaintStepTest_ContainerRingRingLink_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingMove_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingNext_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingPrev_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingUnlink_B0I0O0(source)
	}
}
