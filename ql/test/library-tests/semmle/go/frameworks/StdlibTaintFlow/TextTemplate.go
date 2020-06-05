package main

import (
	"io"
	"text/template"
)

func TaintStepTest_TextTemplateHTMLEscape(sourceCQL interface{}) {
	// The flow is from `b` into `w`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `b` to parameter `w`;
	// `w` is now tainted.
	template.HTMLEscape(w, b)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_TextTemplateHTMLEscapeString(sourceCQL interface{}) {
	// The flow is from `s` into `into281`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into281`
	// (`into281` is now tainted).
	into281 := template.HTMLEscapeString(s)

	// Sink the tainted `into281`:
	sink(into281)
}

func TaintStepTest_TextTemplateHTMLEscaper(sourceCQL interface{}) {
	// The flow is from `args` into `into178`.

	// Assume that `sourceCQL` has the underlying type of `args`:
	args := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `args` to result `into178`
	// (`into178` is now tainted).
	into178 := template.HTMLEscaper(args)

	// Sink the tainted `into178`:
	sink(into178)
}

func TaintStepTest_TextTemplateJSEscape(sourceCQL interface{}) {
	// The flow is from `b` into `w`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `b` to parameter `w`;
	// `w` is now tainted.
	template.JSEscape(w, b)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_TextTemplateJSEscapeString(sourceCQL interface{}) {
	// The flow is from `s` into `into124`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into124`
	// (`into124` is now tainted).
	into124 := template.JSEscapeString(s)

	// Sink the tainted `into124`:
	sink(into124)
}

func TaintStepTest_TextTemplateJSEscaper(sourceCQL interface{}) {
	// The flow is from `args` into `into186`.

	// Assume that `sourceCQL` has the underlying type of `args`:
	args := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `args` to result `into186`
	// (`into186` is now tainted).
	into186 := template.JSEscaper(args)

	// Sink the tainted `into186`:
	sink(into186)
}

func TaintStepTest_TextTemplateURLQueryEscaper(sourceCQL interface{}) {
	// The flow is from `args` into `into551`.

	// Assume that `sourceCQL` has the underlying type of `args`:
	args := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `args` to result `into551`
	// (`into551` is now tainted).
	into551 := template.URLQueryEscaper(args)

	// Sink the tainted `into551`:
	sink(into551)
}

func TaintStepTest_TextTemplateTemplateExecute(sourceCQL interface{}) {
	// The flow is from `data` into `wr`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.(interface{})

	// Declare `wr` variable:
	var wr io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `data` to the parameter `wr`
	// (`wr` is now tainted).
	mediumObjCQL.Execute(wr, data)

	// Sink the tainted `wr`:
	sink(wr)
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate(sourceCQL interface{}) {
	// The flow is from `data` into `wr`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.(interface{})

	// Declare `wr` variable:
	var wr io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `data` to the parameter `wr`
	// (`wr` is now tainted).
	mediumObjCQL.ExecuteTemplate(wr, "", data)

	// Sink the tainted `wr`:
	sink(wr)
}

func RunAllTaints_TextTemplate(v interface{}) {
	{
		source := newSource()
		TaintStepTest_TextTemplateHTMLEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateHTMLEscapeString(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateHTMLEscaper(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateJSEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateJSEscapeString(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateJSEscaper(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateURLQueryEscaper(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateTemplateExecute(source)
	}
	{
		source := newSource()
		TaintStepTest_TextTemplateTemplateExecuteTemplate(source)
	}
}
