package main

import "container/ring"

func TaintStepTest_ContainerRingRingLink(sourceCQL interface{}) {
	// The flow is from `s` into `into484`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(*ring.Ring)

	// Declare medium object/interface:
	var mediumObjCQL ring.Ring

	// Call the method that transfers the taint
	// from the parameter `s` to the result `into484`
	// (`into484` is now tainted).
	into484 := mediumObjCQL.Link(s)

	// Sink the tainted `into484`:
	sink(into484)
}

func TaintStepTest_ContainerRingRingMove(sourceCQL interface{}) {
	// The flow is from `from179` into `into813`.

	// Assume that `sourceCQL` has the underlying type of `from179`:
	from179 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `from179` to the result `into813`
	// (`into813` is now tainted).
	into813 := from179.Move(0)

	// Sink the tainted `into813`:
	sink(into813)
}

func TaintStepTest_ContainerRingRingNext(sourceCQL interface{}) {
	// The flow is from `from887` into `into660`.

	// Assume that `sourceCQL` has the underlying type of `from887`:
	from887 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `from887` to the result `into660`
	// (`into660` is now tainted).
	into660 := from887.Next()

	// Sink the tainted `into660`:
	sink(into660)
}

func TaintStepTest_ContainerRingRingPrev(sourceCQL interface{}) {
	// The flow is from `from410` into `into704`.

	// Assume that `sourceCQL` has the underlying type of `from410`:
	from410 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `from410` to the result `into704`
	// (`into704` is now tainted).
	into704 := from410.Prev()

	// Sink the tainted `into704`:
	sink(into704)
}

func TaintStepTest_ContainerRingRingUnlink(sourceCQL interface{}) {
	// The flow is from `from966` into `into505`.

	// Assume that `sourceCQL` has the underlying type of `from966`:
	from966 := sourceCQL.(ring.Ring)

	// Call the method that transfers the taint
	// from the receiver `from966` to the result `into505`
	// (`into505` is now tainted).
	into505 := from966.Unlink(0)

	// Sink the tainted `into505`:
	sink(into505)
}

func RunAllTaints_ContainerRing(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ContainerRingRingLink(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingMove(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingNext(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingPrev(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerRingRingUnlink(source)
	}
}
