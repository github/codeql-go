package main

import (
	"io"
	"text/template"
)

func TaintStepTest_TextTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte774` into `intoWriter307`.

	// Assume that `sourceCQL` has the underlying type of `fromByte774`:
	fromByte774 := sourceCQL.([]byte)

	// Declare `intoWriter307` variable:
	var intoWriter307 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte774` to parameter `intoWriter307`;
	// `intoWriter307` is now tainted.
	template.HTMLEscape(intoWriter307, fromByte774)

	// Sink the tainted `intoWriter307`:
	sink(intoWriter307)
}

func TaintStepTest_TextTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString186` into `intoString152`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoString152`
	// (`intoString152` is now tainted).
	intoString152 := template.HTMLEscapeString(fromString186)

	// Sink the tainted `intoString152`:
	sink(intoString152)
}

func TaintStepTest_TextTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface600` into `intoString752`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface600`:
	fromInterface600 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface600` to result `intoString752`
	// (`intoString752` is now tainted).
	intoString752 := template.HTMLEscaper(fromInterface600)

	// Sink the tainted `intoString752`:
	sink(intoString752)
}

func TaintStepTest_TextTemplateJSEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte202` into `intoWriter557`.

	// Assume that `sourceCQL` has the underlying type of `fromByte202`:
	fromByte202 := sourceCQL.([]byte)

	// Declare `intoWriter557` variable:
	var intoWriter557 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte202` to parameter `intoWriter557`;
	// `intoWriter557` is now tainted.
	template.JSEscape(intoWriter557, fromByte202)

	// Sink the tainted `intoWriter557`:
	sink(intoWriter557)
}

func TaintStepTest_TextTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString217` into `intoString665`.

	// Assume that `sourceCQL` has the underlying type of `fromString217`:
	fromString217 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString217` to result `intoString665`
	// (`intoString665` is now tainted).
	intoString665 := template.JSEscapeString(fromString217)

	// Sink the tainted `intoString665`:
	sink(intoString665)
}

func TaintStepTest_TextTemplateJSEscaper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface600` into `intoString492`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface600`:
	fromInterface600 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface600` to result `intoString492`
	// (`intoString492` is now tainted).
	intoString492 := template.JSEscaper(fromInterface600)

	// Sink the tainted `intoString492`:
	sink(intoString492)
}

func TaintStepTest_TextTemplateNew_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString173` into `intoTemplate888`.

	// Assume that `sourceCQL` has the underlying type of `fromString173`:
	fromString173 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString173` to result `intoTemplate888`
	// (`intoTemplate888` is now tainted).
	intoTemplate888 := template.New(fromString173)

	// Sink the tainted `intoTemplate888`:
	sink(intoTemplate888)
}

func TaintStepTest_TextTemplateParseFiles_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString596` into `intoTemplate221`.

	// Assume that `sourceCQL` has the underlying type of `fromString596`:
	fromString596 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString596` to result `intoTemplate221`
	// (`intoTemplate221` is now tainted).
	intoTemplate221, _ := template.ParseFiles(fromString596)

	// Sink the tainted `intoTemplate221`:
	sink(intoTemplate221)
}

func TaintStepTest_TextTemplateParseGlob_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString764` into `intoTemplate880`.

	// Assume that `sourceCQL` has the underlying type of `fromString764`:
	fromString764 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString764` to result `intoTemplate880`
	// (`intoTemplate880` is now tainted).
	intoTemplate880, _ := template.ParseGlob(fromString764)

	// Sink the tainted `intoTemplate880`:
	sink(intoTemplate880)
}

func TaintStepTest_TextTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface144` into `intoString922`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface144`:
	fromInterface144 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface144` to result `intoString922`
	// (`intoString922` is now tainted).
	intoString922 := template.URLQueryEscaper(fromInterface144)

	// Sink the tainted `intoString922`:
	sink(intoString922)
}

func TaintStepTest_TextTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate720` into `intoWriter867`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate720`:
	fromTemplate720 := sourceCQL.(template.Template)

	// Declare `intoWriter867` variable:
	var intoWriter867 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate720` to the argument `intoWriter867`
	// (`intoWriter867` is now tainted).
	fromTemplate720.Execute(intoWriter867, nil)

	// Sink the tainted `intoWriter867`:
	sink(intoWriter867)
}

func TaintStepTest_TextTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface449` into `intoWriter867`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface449`:
	fromInterface449 := sourceCQL.(interface{})

	// Declare `intoWriter867` variable:
	var intoWriter867 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface449` to the parameter `intoWriter867`
	// (`intoWriter867` is now tainted).
	mediumObjCQL.Execute(intoWriter867, fromInterface449)

	// Sink the tainted `intoWriter867`:
	sink(intoWriter867)
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate988` into `intoWriter488`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate988`:
	fromTemplate988 := sourceCQL.(template.Template)

	// Declare `intoWriter488` variable:
	var intoWriter488 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate988` to the argument `intoWriter488`
	// (`intoWriter488` is now tainted).
	fromTemplate988.ExecuteTemplate(intoWriter488, "", nil)

	// Sink the tainted `intoWriter488`:
	sink(intoWriter488)
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface732` into `intoWriter370`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface732`:
	fromInterface732 := sourceCQL.(interface{})

	// Declare `intoWriter370` variable:
	var intoWriter370 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface732` to the parameter `intoWriter370`
	// (`intoWriter370` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter370, "", fromInterface732)

	// Sink the tainted `intoWriter370`:
	sink(intoWriter370)
}

func RunAllTaints_TextTemplate() {
	{
		source := newSource()
		TaintStepTest_TextTemplateHTMLEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateHTMLEscapeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateHTMLEscaper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateJSEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateJSEscapeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateJSEscaper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateNew_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateParseFiles_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateParseGlob_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateURLQueryEscaper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateTemplateExecute_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateTemplateExecute_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I1O0(source)
	}
}
