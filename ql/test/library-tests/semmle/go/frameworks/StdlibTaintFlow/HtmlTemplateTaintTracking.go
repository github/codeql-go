package main

import (
	"html/template"
	"io"
)

func TaintStepTest_HtmlTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte727` into `intoWriter261`.

	// Assume that `sourceCQL` has the underlying type of `fromByte727`:
	fromByte727 := sourceCQL.([]byte)

	// Declare `intoWriter261` variable:
	var intoWriter261 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte727` to parameter `intoWriter261`;
	// `intoWriter261` is now tainted.
	template.HTMLEscape(intoWriter261, fromByte727)

	// Return the tainted `intoWriter261`:
	return intoWriter261
}

func TaintStepTest_HtmlTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString213` into `intoString420`.

	// Assume that `sourceCQL` has the underlying type of `fromString213`:
	fromString213 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString213` to result `intoString420`
	// (`intoString420` is now tainted).
	intoString420 := template.HTMLEscapeString(fromString213)

	// Return the tainted `intoString420`:
	return intoString420
}

func TaintStepTest_HtmlTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface916` into `intoString204`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface916`:
	fromInterface916 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface916` to result `intoString204`
	// (`intoString204` is now tainted).
	intoString204 := template.HTMLEscaper(fromInterface916)

	// Return the tainted `intoString204`:
	return intoString204
}

func TaintStepTest_HtmlTemplateJSEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte671` into `intoWriter406`.

	// Assume that `sourceCQL` has the underlying type of `fromByte671`:
	fromByte671 := sourceCQL.([]byte)

	// Declare `intoWriter406` variable:
	var intoWriter406 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte671` to parameter `intoWriter406`;
	// `intoWriter406` is now tainted.
	template.JSEscape(intoWriter406, fromByte671)

	// Return the tainted `intoWriter406`:
	return intoWriter406
}

func TaintStepTest_HtmlTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString705` into `intoString880`.

	// Assume that `sourceCQL` has the underlying type of `fromString705`:
	fromString705 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString705` to result `intoString880`
	// (`intoString880` is now tainted).
	intoString880 := template.JSEscapeString(fromString705)

	// Return the tainted `intoString880`:
	return intoString880
}

func TaintStepTest_HtmlTemplateJSEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface513` into `intoString282`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface513`:
	fromInterface513 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface513` to result `intoString282`
	// (`intoString282` is now tainted).
	intoString282 := template.JSEscaper(fromInterface513)

	// Return the tainted `intoString282`:
	return intoString282
}

func TaintStepTest_HtmlTemplateMust_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate916` into `intoTemplate518`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate916`:
	fromTemplate916 := sourceCQL.(*template.Template)

	// Call the function that transfers the taint
	// from the parameter `fromTemplate916` to result `intoTemplate518`
	// (`intoTemplate518` is now tainted).
	intoTemplate518 := template.Must(fromTemplate916, nil)

	// Return the tainted `intoTemplate518`:
	return intoTemplate518
}

func TaintStepTest_HtmlTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString851` into `intoTemplate504`.

	// Assume that `sourceCQL` has the underlying type of `fromString851`:
	fromString851 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString851` to result `intoTemplate504`
	// (`intoTemplate504` is now tainted).
	intoTemplate504, _ := template.ParseFiles(fromString851)

	// Return the tainted `intoTemplate504`:
	return intoTemplate504
}

func TaintStepTest_HtmlTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString961` into `intoTemplate333`.

	// Assume that `sourceCQL` has the underlying type of `fromString961`:
	fromString961 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString961` to result `intoTemplate333`
	// (`intoTemplate333` is now tainted).
	intoTemplate333, _ := template.ParseGlob(fromString961)

	// Return the tainted `intoTemplate333`:
	return intoTemplate333
}

func TaintStepTest_HtmlTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface684` into `intoString272`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface684`:
	fromInterface684 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface684` to result `intoString272`
	// (`intoString272` is now tainted).
	intoString272 := template.URLQueryEscaper(fromInterface684)

	// Return the tainted `intoString272`:
	return intoString272
}

func TaintStepTest_HtmlTemplateTemplateClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate961` into `intoTemplate508`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate961`:
	fromTemplate961 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate961` to the result `intoTemplate508`
	// (`intoTemplate508` is now tainted).
	intoTemplate508, _ := fromTemplate961.Clone()

	// Return the tainted `intoTemplate508`:
	return intoTemplate508
}

func TaintStepTest_HtmlTemplateTemplateDefinedTemplates_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate147` into `intoString775`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate147`:
	fromTemplate147 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate147` to the result `intoString775`
	// (`intoString775` is now tainted).
	intoString775 := fromTemplate147.DefinedTemplates()

	// Return the tainted `intoString775`:
	return intoString775
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate412` into `intoTemplate404`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate412`:
	fromTemplate412 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate412` to the result `intoTemplate404`
	// (`intoTemplate404` is now tainted).
	intoTemplate404 := fromTemplate412.Delims("", "")

	// Return the tainted `intoTemplate404`:
	return intoTemplate404
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString322` into `intoTemplate574`.

	// Assume that `sourceCQL` has the underlying type of `fromString322`:
	fromString322 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString322` to the result `intoTemplate574`
	// (`intoTemplate574` is now tainted).
	intoTemplate574 := mediumObjCQL.Delims(fromString322, "")

	// Return the tainted `intoTemplate574`:
	return intoTemplate574
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString648` into `intoTemplate845`.

	// Assume that `sourceCQL` has the underlying type of `fromString648`:
	fromString648 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString648` to the result `intoTemplate845`
	// (`intoTemplate845` is now tainted).
	intoTemplate845 := mediumObjCQL.Delims("", fromString648)

	// Return the tainted `intoTemplate845`:
	return intoTemplate845
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate922` into `intoWriter486`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate922`:
	fromTemplate922 := sourceCQL.(template.Template)

	// Declare `intoWriter486` variable:
	var intoWriter486 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate922` to the argument `intoWriter486`
	// (`intoWriter486` is now tainted).
	fromTemplate922.Execute(intoWriter486, nil)

	// Return the tainted `intoWriter486`:
	return intoWriter486
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface150` into `intoWriter396`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface150`:
	fromInterface150 := sourceCQL.(interface{})

	// Declare `intoWriter396` variable:
	var intoWriter396 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface150` to the parameter `intoWriter396`
	// (`intoWriter396` is now tainted).
	mediumObjCQL.Execute(intoWriter396, fromInterface150)

	// Return the tainted `intoWriter396`:
	return intoWriter396
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate252` into `intoWriter289`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate252`:
	fromTemplate252 := sourceCQL.(template.Template)

	// Declare `intoWriter289` variable:
	var intoWriter289 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate252` to the argument `intoWriter289`
	// (`intoWriter289` is now tainted).
	fromTemplate252.ExecuteTemplate(intoWriter289, "", nil)

	// Return the tainted `intoWriter289`:
	return intoWriter289
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface970` into `intoWriter360`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface970`:
	fromInterface970 := sourceCQL.(interface{})

	// Declare `intoWriter360` variable:
	var intoWriter360 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface970` to the parameter `intoWriter360`
	// (`intoWriter360` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter360, "", fromInterface970)

	// Return the tainted `intoWriter360`:
	return intoWriter360
}

func TaintStepTest_HtmlTemplateTemplateLookup_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate277` into `intoTemplate962`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate277`:
	fromTemplate277 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate277` to the result `intoTemplate962`
	// (`intoTemplate962` is now tainted).
	intoTemplate962 := fromTemplate277.Lookup("")

	// Return the tainted `intoTemplate962`:
	return intoTemplate962
}

func TaintStepTest_HtmlTemplateTemplateNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate221` into `intoTemplate653`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate221`:
	fromTemplate221 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate221` to the result `intoTemplate653`
	// (`intoTemplate653` is now tainted).
	intoTemplate653 := fromTemplate221.New("")

	// Return the tainted `intoTemplate653`:
	return intoTemplate653
}

func TaintStepTest_HtmlTemplateTemplateOption_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate279` into `intoTemplate748`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate279`:
	fromTemplate279 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate279` to the result `intoTemplate748`
	// (`intoTemplate748` is now tainted).
	intoTemplate748 := fromTemplate279.Option("")

	// Return the tainted `intoTemplate748`:
	return intoTemplate748
}

func TaintStepTest_HtmlTemplateTemplateParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString540` into `intoTemplate775`.

	// Assume that `sourceCQL` has the underlying type of `fromString540`:
	fromString540 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString540` to the result `intoTemplate775`
	// (`intoTemplate775` is now tainted).
	intoTemplate775, _ := mediumObjCQL.Parse(fromString540)

	// Return the tainted `intoTemplate775`:
	return intoTemplate775
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate577` into `intoTemplate589`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate577`:
	fromTemplate577 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate577` to the result `intoTemplate589`
	// (`intoTemplate589` is now tainted).
	intoTemplate589, _ := fromTemplate577.ParseFiles("")

	// Return the tainted `intoTemplate589`:
	return intoTemplate589
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString378` into `intoTemplate394`.

	// Assume that `sourceCQL` has the underlying type of `fromString378`:
	fromString378 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString378` to the result `intoTemplate394`
	// (`intoTemplate394` is now tainted).
	intoTemplate394, _ := mediumObjCQL.ParseFiles(fromString378)

	// Return the tainted `intoTemplate394`:
	return intoTemplate394
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate624` into `intoTemplate211`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate624`:
	fromTemplate624 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate624` to the result `intoTemplate211`
	// (`intoTemplate211` is now tainted).
	intoTemplate211, _ := fromTemplate624.ParseGlob("")

	// Return the tainted `intoTemplate211`:
	return intoTemplate211
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString806` into `intoTemplate540`.

	// Assume that `sourceCQL` has the underlying type of `fromString806`:
	fromString806 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString806` to the result `intoTemplate540`
	// (`intoTemplate540` is now tainted).
	intoTemplate540, _ := mediumObjCQL.ParseGlob(fromString806)

	// Return the tainted `intoTemplate540`:
	return intoTemplate540
}

func TaintStepTest_HtmlTemplateTemplateTemplates_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate665` into `intoTemplate228`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate665`:
	fromTemplate665 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate665` to the result `intoTemplate228`
	// (`intoTemplate228` is now tainted).
	intoTemplate228 := fromTemplate665.Templates()

	// Return the tainted `intoTemplate228`:
	return intoTemplate228
}

func RunAllTaints_HtmlTemplate() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateHTMLEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateHTMLEscapeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateHTMLEscaper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateJSEscape_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateJSEscapeString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateJSEscaper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateMust_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateParseFiles_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateParseGlob_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateURLQueryEscaper_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateClone_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateDefinedTemplates_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateDelims_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateDelims_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateDelims_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateExecute_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateExecute_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateLookup_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateNew_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateOption_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateParse_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateParseFiles_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateParseFiles_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateParseGlob_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateParseGlob_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_HtmlTemplateTemplateTemplates_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
