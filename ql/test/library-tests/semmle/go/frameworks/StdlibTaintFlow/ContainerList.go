package main

import "container/list"

func TaintStepTest_ContainerListListBack(sourceCQL interface{}) {
	// The flow is from `from572` into `into360`.

	// Assume that `sourceCQL` has the underlying type of `from572`:
	from572 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `from572` to the result `into360`
	// (`into360` is now tainted).
	into360 := from572.Back()

	// Sink the tainted `into360`:
	sink(into360)
}

func TaintStepTest_ContainerListListFront(sourceCQL interface{}) {
	// The flow is from `from551` into `into614`.

	// Assume that `sourceCQL` has the underlying type of `from551`:
	from551 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `from551` to the result `into614`
	// (`into614` is now tainted).
	into614 := from551.Front()

	// Sink the tainted `into614`:
	sink(into614)
}

func TaintStepTest_ContainerListListInit(sourceCQL interface{}) {
	// The flow is from `from693` into `into888`.

	// Assume that `sourceCQL` has the underlying type of `from693`:
	from693 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `from693` to the result `into888`
	// (`into888` is now tainted).
	into888 := from693.Init()

	// Sink the tainted `into888`:
	sink(into888)
}

func TaintStepTest_ContainerListListInsertAfter(sourceCQL interface{}) {
	// The flow is from `v` into `into726`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into726` variable:
	var into726 list.List

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into726`
	// (`into726` is now tainted).
	into726.InsertAfter(v, nil)

	// Sink the tainted `into726`:
	sink(into726)
}

func TaintStepTest_ContainerListListInsertBefore(sourceCQL interface{}) {
	// The flow is from `v` into `into629`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into629` variable:
	var into629 list.List

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into629`
	// (`into629` is now tainted).
	into629.InsertBefore(v, nil)

	// Sink the tainted `into629`:
	sink(into629)
}

func TaintStepTest_ContainerListListMoveAfter(sourceCQL interface{}) {
	// The flow is from `e` into `into414`.

	// Assume that `sourceCQL` has the underlying type of `e`:
	e := sourceCQL.(*list.Element)

	// Declare `into414` variable:
	var into414 list.List

	// Call the method that transfers the taint
	// from the parameter `e` to the receiver `into414`
	// (`into414` is now tainted).
	into414.MoveAfter(e, nil)

	// Sink the tainted `into414`:
	sink(into414)
}

func TaintStepTest_ContainerListListMoveBefore(sourceCQL interface{}) {
	// The flow is from `e` into `into544`.

	// Assume that `sourceCQL` has the underlying type of `e`:
	e := sourceCQL.(*list.Element)

	// Declare `into544` variable:
	var into544 list.List

	// Call the method that transfers the taint
	// from the parameter `e` to the receiver `into544`
	// (`into544` is now tainted).
	into544.MoveBefore(e, nil)

	// Sink the tainted `into544`:
	sink(into544)
}

func TaintStepTest_ContainerListListMoveToBack(sourceCQL interface{}) {
	// The flow is from `e` into `into801`.

	// Assume that `sourceCQL` has the underlying type of `e`:
	e := sourceCQL.(*list.Element)

	// Declare `into801` variable:
	var into801 list.List

	// Call the method that transfers the taint
	// from the parameter `e` to the receiver `into801`
	// (`into801` is now tainted).
	into801.MoveToBack(e)

	// Sink the tainted `into801`:
	sink(into801)
}

func TaintStepTest_ContainerListListMoveToFront(sourceCQL interface{}) {
	// The flow is from `e` into `into233`.

	// Assume that `sourceCQL` has the underlying type of `e`:
	e := sourceCQL.(*list.Element)

	// Declare `into233` variable:
	var into233 list.List

	// Call the method that transfers the taint
	// from the parameter `e` to the receiver `into233`
	// (`into233` is now tainted).
	into233.MoveToFront(e)

	// Sink the tainted `into233`:
	sink(into233)
}

func TaintStepTest_ContainerListListPushBack(sourceCQL interface{}) {
	// The flow is from `v` into `into201`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into201` variable:
	var into201 list.List

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into201`
	// (`into201` is now tainted).
	into201.PushBack(v)

	// Sink the tainted `into201`:
	sink(into201)
}

func TaintStepTest_ContainerListListPushBackList(sourceCQL interface{}) {
	// The flow is from `other` into `into839`.

	// Assume that `sourceCQL` has the underlying type of `other`:
	other := sourceCQL.(*list.List)

	// Declare `into839` variable:
	var into839 list.List

	// Call the method that transfers the taint
	// from the parameter `other` to the receiver `into839`
	// (`into839` is now tainted).
	into839.PushBackList(other)

	// Sink the tainted `into839`:
	sink(into839)
}

func TaintStepTest_ContainerListListPushFront(sourceCQL interface{}) {
	// The flow is from `v` into `into292`.

	// Assume that `sourceCQL` has the underlying type of `v`:
	v := sourceCQL.(interface{})

	// Declare `into292` variable:
	var into292 list.List

	// Call the method that transfers the taint
	// from the parameter `v` to the receiver `into292`
	// (`into292` is now tainted).
	into292.PushFront(v)

	// Sink the tainted `into292`:
	sink(into292)
}

func TaintStepTest_ContainerListListPushFrontList(sourceCQL interface{}) {
	// The flow is from `other` into `into498`.

	// Assume that `sourceCQL` has the underlying type of `other`:
	other := sourceCQL.(*list.List)

	// Declare `into498` variable:
	var into498 list.List

	// Call the method that transfers the taint
	// from the parameter `other` to the receiver `into498`
	// (`into498` is now tainted).
	into498.PushFrontList(other)

	// Sink the tainted `into498`:
	sink(into498)
}

func TaintStepTest_ContainerListListRemove(sourceCQL interface{}) {
	// The flow is from `e` into `into899`.

	// Assume that `sourceCQL` has the underlying type of `e`:
	e := sourceCQL.(*list.Element)

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `e` to the result `into899`
	// (`into899` is now tainted).
	into899 := mediumObjCQL.Remove(e)

	// Sink the tainted `into899`:
	sink(into899)
}

func RunAllTaints_ContainerList(v interface{}) {
	{
		source := newSource()
		TaintStepTest_ContainerListListBack(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListFront(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInit(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInsertAfter(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInsertBefore(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveAfter(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveBefore(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveToBack(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveToFront(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushBack(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushBackList(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushFront(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushFrontList(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListRemove(source)
	}
}
