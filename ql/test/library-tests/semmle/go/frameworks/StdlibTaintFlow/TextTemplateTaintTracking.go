// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"text/template"
)

func TaintStepTest_TextTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte656` into `intoWriter414`.

	// Assume that `sourceCQL` has the underlying type of `fromByte656`:
	fromByte656 := sourceCQL.([]byte)

	// Declare `intoWriter414` variable:
	var intoWriter414 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte656` to parameter `intoWriter414`;
	// `intoWriter414` is now tainted.
	template.HTMLEscape(intoWriter414, fromByte656)

	// Return the tainted `intoWriter414`:
	return intoWriter414
}

func TaintStepTest_TextTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString518` into `intoString650`.

	// Assume that `sourceCQL` has the underlying type of `fromString518`:
	fromString518 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString518` to result `intoString650`
	// (`intoString650` is now tainted).
	intoString650 := template.HTMLEscapeString(fromString518)

	// Return the tainted `intoString650`:
	return intoString650
}

func TaintStepTest_TextTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface784` into `intoString957`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface784`:
	fromInterface784 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface784` to result `intoString957`
	// (`intoString957` is now tainted).
	intoString957 := template.HTMLEscaper(fromInterface784)

	// Return the tainted `intoString957`:
	return intoString957
}

func TaintStepTest_TextTemplateJSEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte520` into `intoWriter443`.

	// Assume that `sourceCQL` has the underlying type of `fromByte520`:
	fromByte520 := sourceCQL.([]byte)

	// Declare `intoWriter443` variable:
	var intoWriter443 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte520` to parameter `intoWriter443`;
	// `intoWriter443` is now tainted.
	template.JSEscape(intoWriter443, fromByte520)

	// Return the tainted `intoWriter443`:
	return intoWriter443
}

func TaintStepTest_TextTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString127` into `intoString483`.

	// Assume that `sourceCQL` has the underlying type of `fromString127`:
	fromString127 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString127` to result `intoString483`
	// (`intoString483` is now tainted).
	intoString483 := template.JSEscapeString(fromString127)

	// Return the tainted `intoString483`:
	return intoString483
}

func TaintStepTest_TextTemplateJSEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface989` into `intoString982`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface989`:
	fromInterface989 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface989` to result `intoString982`
	// (`intoString982` is now tainted).
	intoString982 := template.JSEscaper(fromInterface989)

	// Return the tainted `intoString982`:
	return intoString982
}

func TaintStepTest_TextTemplateNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString417` into `intoTemplate584`.

	// Assume that `sourceCQL` has the underlying type of `fromString417`:
	fromString417 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString417` to result `intoTemplate584`
	// (`intoTemplate584` is now tainted).
	intoTemplate584 := template.New(fromString417)

	// Return the tainted `intoTemplate584`:
	return intoTemplate584
}

func TaintStepTest_TextTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString991` into `intoTemplate881`.

	// Assume that `sourceCQL` has the underlying type of `fromString991`:
	fromString991 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString991` to result `intoTemplate881`
	// (`intoTemplate881` is now tainted).
	intoTemplate881, _ := template.ParseFiles(fromString991)

	// Return the tainted `intoTemplate881`:
	return intoTemplate881
}

func TaintStepTest_TextTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString186` into `intoTemplate284`.

	// Assume that `sourceCQL` has the underlying type of `fromString186`:
	fromString186 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString186` to result `intoTemplate284`
	// (`intoTemplate284` is now tainted).
	intoTemplate284, _ := template.ParseGlob(fromString186)

	// Return the tainted `intoTemplate284`:
	return intoTemplate284
}

func TaintStepTest_TextTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface908` into `intoString137`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface908`:
	fromInterface908 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface908` to result `intoString137`
	// (`intoString137` is now tainted).
	intoString137 := template.URLQueryEscaper(fromInterface908)

	// Return the tainted `intoString137`:
	return intoString137
}

func TaintStepTest_TextTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate494` into `intoWriter873`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate494`:
	fromTemplate494 := sourceCQL.(template.Template)

	// Declare `intoWriter873` variable:
	var intoWriter873 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate494` to the argument `intoWriter873`
	// (`intoWriter873` is now tainted).
	fromTemplate494.Execute(intoWriter873, nil)

	// Return the tainted `intoWriter873`:
	return intoWriter873
}

func TaintStepTest_TextTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface599` into `intoWriter409`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface599`:
	fromInterface599 := sourceCQL.(interface{})

	// Declare `intoWriter409` variable:
	var intoWriter409 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface599` to the parameter `intoWriter409`
	// (`intoWriter409` is now tainted).
	mediumObjCQL.Execute(intoWriter409, fromInterface599)

	// Return the tainted `intoWriter409`:
	return intoWriter409
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate246` into `intoWriter898`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate246`:
	fromTemplate246 := sourceCQL.(template.Template)

	// Declare `intoWriter898` variable:
	var intoWriter898 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate246` to the argument `intoWriter898`
	// (`intoWriter898` is now tainted).
	fromTemplate246.ExecuteTemplate(intoWriter898, "", nil)

	// Return the tainted `intoWriter898`:
	return intoWriter898
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface598` into `intoWriter631`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface598`:
	fromInterface598 := sourceCQL.(interface{})

	// Declare `intoWriter631` variable:
	var intoWriter631 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface598` to the parameter `intoWriter631`
	// (`intoWriter631` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter631, "", fromInterface598)

	// Return the tainted `intoWriter631`:
	return intoWriter631
}

func RunAllTaints_TextTemplate() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateHTMLEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateHTMLEscapeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateHTMLEscaper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateJSEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateJSEscapeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateJSEscaper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateNew_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateParseFiles_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateParseGlob_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateURLQueryEscaper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateTemplateExecute_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateTemplateExecute_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
