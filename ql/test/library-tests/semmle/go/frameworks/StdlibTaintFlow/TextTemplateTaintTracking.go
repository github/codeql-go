// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"text/template"
)

func TaintStepTest_TextTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte785` into `intoWriter156`.

	// Assume that `sourceCQL` has the underlying type of `fromByte785`:
	fromByte785 := sourceCQL.([]byte)

	// Declare `intoWriter156` variable:
	var intoWriter156 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte785` to parameter `intoWriter156`;
	// `intoWriter156` is now tainted.
	template.HTMLEscape(intoWriter156, fromByte785)

	// Return the tainted `intoWriter156`:
	return intoWriter156
}

func TaintStepTest_TextTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString995` into `intoString636`.

	// Assume that `sourceCQL` has the underlying type of `fromString995`:
	fromString995 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString995` to result `intoString636`
	// (`intoString636` is now tainted).
	intoString636 := template.HTMLEscapeString(fromString995)

	// Return the tainted `intoString636`:
	return intoString636
}

func TaintStepTest_TextTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface235` into `intoString876`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface235`:
	fromInterface235 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface235` to result `intoString876`
	// (`intoString876` is now tainted).
	intoString876 := template.HTMLEscaper(fromInterface235)

	// Return the tainted `intoString876`:
	return intoString876
}

func TaintStepTest_TextTemplateJSEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte473` into `intoWriter287`.

	// Assume that `sourceCQL` has the underlying type of `fromByte473`:
	fromByte473 := sourceCQL.([]byte)

	// Declare `intoWriter287` variable:
	var intoWriter287 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte473` to parameter `intoWriter287`;
	// `intoWriter287` is now tainted.
	template.JSEscape(intoWriter287, fromByte473)

	// Return the tainted `intoWriter287`:
	return intoWriter287
}

func TaintStepTest_TextTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString618` into `intoString838`.

	// Assume that `sourceCQL` has the underlying type of `fromString618`:
	fromString618 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString618` to result `intoString838`
	// (`intoString838` is now tainted).
	intoString838 := template.JSEscapeString(fromString618)

	// Return the tainted `intoString838`:
	return intoString838
}

func TaintStepTest_TextTemplateJSEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface318` into `intoString298`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface318`:
	fromInterface318 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface318` to result `intoString298`
	// (`intoString298` is now tainted).
	intoString298 := template.JSEscaper(fromInterface318)

	// Return the tainted `intoString298`:
	return intoString298
}

func TaintStepTest_TextTemplateNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString758` into `intoTemplate459`.

	// Assume that `sourceCQL` has the underlying type of `fromString758`:
	fromString758 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString758` to result `intoTemplate459`
	// (`intoTemplate459` is now tainted).
	intoTemplate459 := template.New(fromString758)

	// Return the tainted `intoTemplate459`:
	return intoTemplate459
}

func TaintStepTest_TextTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString571` into `intoTemplate472`.

	// Assume that `sourceCQL` has the underlying type of `fromString571`:
	fromString571 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString571` to result `intoTemplate472`
	// (`intoTemplate472` is now tainted).
	intoTemplate472, _ := template.ParseFiles(fromString571)

	// Return the tainted `intoTemplate472`:
	return intoTemplate472
}

func TaintStepTest_TextTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString191` into `intoTemplate123`.

	// Assume that `sourceCQL` has the underlying type of `fromString191`:
	fromString191 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString191` to result `intoTemplate123`
	// (`intoTemplate123` is now tainted).
	intoTemplate123, _ := template.ParseGlob(fromString191)

	// Return the tainted `intoTemplate123`:
	return intoTemplate123
}

func TaintStepTest_TextTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface451` into `intoString118`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface451`:
	fromInterface451 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface451` to result `intoString118`
	// (`intoString118` is now tainted).
	intoString118 := template.URLQueryEscaper(fromInterface451)

	// Return the tainted `intoString118`:
	return intoString118
}

func TaintStepTest_TextTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate216` into `intoWriter748`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate216`:
	fromTemplate216 := sourceCQL.(template.Template)

	// Declare `intoWriter748` variable:
	var intoWriter748 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate216` to the argument `intoWriter748`
	// (`intoWriter748` is now tainted).
	fromTemplate216.Execute(intoWriter748, nil)

	// Return the tainted `intoWriter748`:
	return intoWriter748
}

func TaintStepTest_TextTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface568` into `intoWriter543`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface568`:
	fromInterface568 := sourceCQL.(interface{})

	// Declare `intoWriter543` variable:
	var intoWriter543 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface568` to the parameter `intoWriter543`
	// (`intoWriter543` is now tainted).
	mediumObjCQL.Execute(intoWriter543, fromInterface568)

	// Return the tainted `intoWriter543`:
	return intoWriter543
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate391` into `intoWriter672`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate391`:
	fromTemplate391 := sourceCQL.(template.Template)

	// Declare `intoWriter672` variable:
	var intoWriter672 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate391` to the argument `intoWriter672`
	// (`intoWriter672` is now tainted).
	fromTemplate391.ExecuteTemplate(intoWriter672, "", nil)

	// Return the tainted `intoWriter672`:
	return intoWriter672
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface175` into `intoWriter140`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface175`:
	fromInterface175 := sourceCQL.(interface{})

	// Declare `intoWriter140` variable:
	var intoWriter140 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface175` to the parameter `intoWriter140`
	// (`intoWriter140` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter140, "", fromInterface175)

	// Return the tainted `intoWriter140`:
	return intoWriter140
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
