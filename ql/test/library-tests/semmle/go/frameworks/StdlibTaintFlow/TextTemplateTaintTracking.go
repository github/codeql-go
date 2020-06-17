package main

import (
	"io"
	"text/template"
)

func TaintStepTest_TextTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte914` into `intoWriter960`.

	// Assume that `sourceCQL` has the underlying type of `fromByte914`:
	fromByte914 := sourceCQL.([]byte)

	// Declare `intoWriter960` variable:
	var intoWriter960 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte914` to parameter `intoWriter960`;
	// `intoWriter960` is now tainted.
	template.HTMLEscape(intoWriter960, fromByte914)

	// Return the tainted `intoWriter960`:
	return intoWriter960
}

func TaintStepTest_TextTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString931` into `intoString288`.

	// Assume that `sourceCQL` has the underlying type of `fromString931`:
	fromString931 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString931` to result `intoString288`
	// (`intoString288` is now tainted).
	intoString288 := template.HTMLEscapeString(fromString931)

	// Return the tainted `intoString288`:
	return intoString288
}

func TaintStepTest_TextTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface135` into `intoString196`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface135`:
	fromInterface135 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface135` to result `intoString196`
	// (`intoString196` is now tainted).
	intoString196 := template.HTMLEscaper(fromInterface135)

	// Return the tainted `intoString196`:
	return intoString196
}

func TaintStepTest_TextTemplateJSEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte619` into `intoWriter285`.

	// Assume that `sourceCQL` has the underlying type of `fromByte619`:
	fromByte619 := sourceCQL.([]byte)

	// Declare `intoWriter285` variable:
	var intoWriter285 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte619` to parameter `intoWriter285`;
	// `intoWriter285` is now tainted.
	template.JSEscape(intoWriter285, fromByte619)

	// Return the tainted `intoWriter285`:
	return intoWriter285
}

func TaintStepTest_TextTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString470` into `intoString384`.

	// Assume that `sourceCQL` has the underlying type of `fromString470`:
	fromString470 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString470` to result `intoString384`
	// (`intoString384` is now tainted).
	intoString384 := template.JSEscapeString(fromString470)

	// Return the tainted `intoString384`:
	return intoString384
}

func TaintStepTest_TextTemplateJSEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface211` into `intoString312`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface211`:
	fromInterface211 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface211` to result `intoString312`
	// (`intoString312` is now tainted).
	intoString312 := template.JSEscaper(fromInterface211)

	// Return the tainted `intoString312`:
	return intoString312
}

func TaintStepTest_TextTemplateNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString897` into `intoTemplate510`.

	// Assume that `sourceCQL` has the underlying type of `fromString897`:
	fromString897 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString897` to result `intoTemplate510`
	// (`intoTemplate510` is now tainted).
	intoTemplate510 := template.New(fromString897)

	// Return the tainted `intoTemplate510`:
	return intoTemplate510
}

func TaintStepTest_TextTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString502` into `intoTemplate280`.

	// Assume that `sourceCQL` has the underlying type of `fromString502`:
	fromString502 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString502` to result `intoTemplate280`
	// (`intoTemplate280` is now tainted).
	intoTemplate280, _ := template.ParseFiles(fromString502)

	// Return the tainted `intoTemplate280`:
	return intoTemplate280
}

func TaintStepTest_TextTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString731` into `intoTemplate195`.

	// Assume that `sourceCQL` has the underlying type of `fromString731`:
	fromString731 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString731` to result `intoTemplate195`
	// (`intoTemplate195` is now tainted).
	intoTemplate195, _ := template.ParseGlob(fromString731)

	// Return the tainted `intoTemplate195`:
	return intoTemplate195
}

func TaintStepTest_TextTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface703` into `intoString643`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface703`:
	fromInterface703 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface703` to result `intoString643`
	// (`intoString643` is now tainted).
	intoString643 := template.URLQueryEscaper(fromInterface703)

	// Return the tainted `intoString643`:
	return intoString643
}

func TaintStepTest_TextTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate152` into `intoWriter577`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate152`:
	fromTemplate152 := sourceCQL.(template.Template)

	// Declare `intoWriter577` variable:
	var intoWriter577 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate152` to the argument `intoWriter577`
	// (`intoWriter577` is now tainted).
	fromTemplate152.Execute(intoWriter577, nil)

	// Return the tainted `intoWriter577`:
	return intoWriter577
}

func TaintStepTest_TextTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface411` into `intoWriter412`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface411`:
	fromInterface411 := sourceCQL.(interface{})

	// Declare `intoWriter412` variable:
	var intoWriter412 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface411` to the parameter `intoWriter412`
	// (`intoWriter412` is now tainted).
	mediumObjCQL.Execute(intoWriter412, fromInterface411)

	// Return the tainted `intoWriter412`:
	return intoWriter412
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate380` into `intoWriter162`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate380`:
	fromTemplate380 := sourceCQL.(template.Template)

	// Declare `intoWriter162` variable:
	var intoWriter162 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate380` to the argument `intoWriter162`
	// (`intoWriter162` is now tainted).
	fromTemplate380.ExecuteTemplate(intoWriter162, "", nil)

	// Return the tainted `intoWriter162`:
	return intoWriter162
}

func TaintStepTest_TextTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface446` into `intoWriter588`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface446`:
	fromInterface446 := sourceCQL.(interface{})

	// Declare `intoWriter588` variable:
	var intoWriter588 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface446` to the parameter `intoWriter588`
	// (`intoWriter588` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter588, "", fromInterface446)

	// Return the tainted `intoWriter588`:
	return intoWriter588
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
