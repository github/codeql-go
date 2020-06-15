package main

import "container/list"

func TaintStepTest_ContainerListListBack_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromList759` into `intoElement193`.

	// Assume that `sourceCQL` has the underlying type of `fromList759`:
	fromList759 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList759` to the result `intoElement193`
	// (`intoElement193` is now tainted).
	intoElement193 := fromList759.Back()

	// Sink the tainted `intoElement193`:
	sink(intoElement193)
}

func TaintStepTest_ContainerListListFront_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromList503` into `intoElement726`.

	// Assume that `sourceCQL` has the underlying type of `fromList503`:
	fromList503 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList503` to the result `intoElement726`
	// (`intoElement726` is now tainted).
	intoElement726 := fromList503.Front()

	// Sink the tainted `intoElement726`:
	sink(intoElement726)
}

func TaintStepTest_ContainerListListInit_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromList564` into `intoList984`.

	// Assume that `sourceCQL` has the underlying type of `fromList564`:
	fromList564 := sourceCQL.(list.List)

	// Call the method that transfers the taint
	// from the receiver `fromList564` to the result `intoList984`
	// (`intoList984` is now tainted).
	intoList984 := fromList564.Init()

	// Sink the tainted `intoList984`:
	sink(intoList984)
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface245` into `intoList493`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface245`:
	fromInterface245 := sourceCQL.(interface{})

	// Declare `intoList493` variable:
	var intoList493 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface245` to the receiver `intoList493`
	// (`intoList493` is now tainted).
	intoList493.InsertAfter(fromInterface245, nil)

	// Sink the tainted `intoList493`:
	sink(intoList493)
}

func TaintStepTest_ContainerListListInsertAfter_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromInterface965` into `intoElement866`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface965`:
	fromInterface965 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface965` to the result `intoElement866`
	// (`intoElement866` is now tainted).
	intoElement866 := mediumObjCQL.InsertAfter(fromInterface965, nil)

	// Sink the tainted `intoElement866`:
	sink(intoElement866)
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface746` into `intoList127`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface746`:
	fromInterface746 := sourceCQL.(interface{})

	// Declare `intoList127` variable:
	var intoList127 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface746` to the receiver `intoList127`
	// (`intoList127` is now tainted).
	intoList127.InsertBefore(fromInterface746, nil)

	// Sink the tainted `intoList127`:
	sink(intoList127)
}

func TaintStepTest_ContainerListListInsertBefore_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromInterface151` into `intoElement437`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface151`:
	fromInterface151 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface151` to the result `intoElement437`
	// (`intoElement437` is now tainted).
	intoElement437 := mediumObjCQL.InsertBefore(fromInterface151, nil)

	// Sink the tainted `intoElement437`:
	sink(intoElement437)
}

func TaintStepTest_ContainerListListMoveAfter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromElement399` into `intoList576`.

	// Assume that `sourceCQL` has the underlying type of `fromElement399`:
	fromElement399 := sourceCQL.(*list.Element)

	// Declare `intoList576` variable:
	var intoList576 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement399` to the receiver `intoList576`
	// (`intoList576` is now tainted).
	intoList576.MoveAfter(fromElement399, nil)

	// Sink the tainted `intoList576`:
	sink(intoList576)
}

func TaintStepTest_ContainerListListMoveBefore_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromElement127` into `intoList233`.

	// Assume that `sourceCQL` has the underlying type of `fromElement127`:
	fromElement127 := sourceCQL.(*list.Element)

	// Declare `intoList233` variable:
	var intoList233 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement127` to the receiver `intoList233`
	// (`intoList233` is now tainted).
	intoList233.MoveBefore(fromElement127, nil)

	// Sink the tainted `intoList233`:
	sink(intoList233)
}

func TaintStepTest_ContainerListListMoveToBack_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromElement357` into `intoList934`.

	// Assume that `sourceCQL` has the underlying type of `fromElement357`:
	fromElement357 := sourceCQL.(*list.Element)

	// Declare `intoList934` variable:
	var intoList934 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement357` to the receiver `intoList934`
	// (`intoList934` is now tainted).
	intoList934.MoveToBack(fromElement357)

	// Sink the tainted `intoList934`:
	sink(intoList934)
}

func TaintStepTest_ContainerListListMoveToFront_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromElement620` into `intoList322`.

	// Assume that `sourceCQL` has the underlying type of `fromElement620`:
	fromElement620 := sourceCQL.(*list.Element)

	// Declare `intoList322` variable:
	var intoList322 list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement620` to the receiver `intoList322`
	// (`intoList322` is now tainted).
	intoList322.MoveToFront(fromElement620)

	// Sink the tainted `intoList322`:
	sink(intoList322)
}

func TaintStepTest_ContainerListListPushBack_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface759` into `intoList851`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface759`:
	fromInterface759 := sourceCQL.(interface{})

	// Declare `intoList851` variable:
	var intoList851 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface759` to the receiver `intoList851`
	// (`intoList851` is now tainted).
	intoList851.PushBack(fromInterface759)

	// Sink the tainted `intoList851`:
	sink(intoList851)
}

func TaintStepTest_ContainerListListPushBack_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromInterface649` into `intoElement809`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface649`:
	fromInterface649 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface649` to the result `intoElement809`
	// (`intoElement809` is now tainted).
	intoElement809 := mediumObjCQL.PushBack(fromInterface649)

	// Sink the tainted `intoElement809`:
	sink(intoElement809)
}

func TaintStepTest_ContainerListListPushBackList_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromList431` into `intoList715`.

	// Assume that `sourceCQL` has the underlying type of `fromList431`:
	fromList431 := sourceCQL.(*list.List)

	// Declare `intoList715` variable:
	var intoList715 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList431` to the receiver `intoList715`
	// (`intoList715` is now tainted).
	intoList715.PushBackList(fromList431)

	// Sink the tainted `intoList715`:
	sink(intoList715)
}

func TaintStepTest_ContainerListListPushFront_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface581` into `intoList814`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface581`:
	fromInterface581 := sourceCQL.(interface{})

	// Declare `intoList814` variable:
	var intoList814 list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface581` to the receiver `intoList814`
	// (`intoList814` is now tainted).
	intoList814.PushFront(fromInterface581)

	// Sink the tainted `intoList814`:
	sink(intoList814)
}

func TaintStepTest_ContainerListListPushFront_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromInterface903` into `intoElement114`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface903`:
	fromInterface903 := sourceCQL.(interface{})

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromInterface903` to the result `intoElement114`
	// (`intoElement114` is now tainted).
	intoElement114 := mediumObjCQL.PushFront(fromInterface903)

	// Sink the tainted `intoElement114`:
	sink(intoElement114)
}

func TaintStepTest_ContainerListListPushFrontList_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromList774` into `intoList350`.

	// Assume that `sourceCQL` has the underlying type of `fromList774`:
	fromList774 := sourceCQL.(*list.List)

	// Declare `intoList350` variable:
	var intoList350 list.List

	// Call the method that transfers the taint
	// from the parameter `fromList774` to the receiver `intoList350`
	// (`intoList350` is now tainted).
	intoList350.PushFrontList(fromList774)

	// Sink the tainted `intoList350`:
	sink(intoList350)
}

func TaintStepTest_ContainerListListRemove_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromElement439` into `intoInterface753`.

	// Assume that `sourceCQL` has the underlying type of `fromElement439`:
	fromElement439 := sourceCQL.(*list.Element)

	// Declare medium object/interface:
	var mediumObjCQL list.List

	// Call the method that transfers the taint
	// from the parameter `fromElement439` to the result `intoInterface753`
	// (`intoInterface753` is now tainted).
	intoInterface753 := mediumObjCQL.Remove(fromElement439)

	// Sink the tainted `intoInterface753`:
	sink(intoInterface753)
}

func RunAllTaints_ContainerList() {
	{
		source := newSource()
		TaintStepTest_ContainerListListBack_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListFront_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInit_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInsertAfter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInsertAfter_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInsertBefore_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListInsertBefore_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveAfter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveBefore_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveToBack_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListMoveToFront_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushBack_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushBack_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushBackList_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushFront_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushFront_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListPushFrontList_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_ContainerListListRemove_B0I0O0(source)
	}
}
