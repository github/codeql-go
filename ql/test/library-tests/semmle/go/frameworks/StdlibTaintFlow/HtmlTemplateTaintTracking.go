// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"html/template"
	"io"
)

func TaintStepTest_HtmlTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte747` into `intoWriter162`.

	// Assume that `sourceCQL` has the underlying type of `fromByte747`:
	fromByte747 := sourceCQL.([]byte)

	// Declare `intoWriter162` variable:
	var intoWriter162 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte747` to parameter `intoWriter162`;
	// `intoWriter162` is now tainted.
	template.HTMLEscape(intoWriter162, fromByte747)

	// Return the tainted `intoWriter162`:
	return intoWriter162
}

func TaintStepTest_HtmlTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString638` into `intoString421`.

	// Assume that `sourceCQL` has the underlying type of `fromString638`:
	fromString638 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString638` to result `intoString421`
	// (`intoString421` is now tainted).
	intoString421 := template.HTMLEscapeString(fromString638)

	// Return the tainted `intoString421`:
	return intoString421
}

func TaintStepTest_HtmlTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface143` into `intoString676`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface143`:
	fromInterface143 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface143` to result `intoString676`
	// (`intoString676` is now tainted).
	intoString676 := template.HTMLEscaper(fromInterface143)

	// Return the tainted `intoString676`:
	return intoString676
}

func TaintStepTest_HtmlTemplateJSEscape_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte387` into `intoWriter297`.

	// Assume that `sourceCQL` has the underlying type of `fromByte387`:
	fromByte387 := sourceCQL.([]byte)

	// Declare `intoWriter297` variable:
	var intoWriter297 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte387` to parameter `intoWriter297`;
	// `intoWriter297` is now tainted.
	template.JSEscape(intoWriter297, fromByte387)

	// Return the tainted `intoWriter297`:
	return intoWriter297
}

func TaintStepTest_HtmlTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString166` into `intoString845`.

	// Assume that `sourceCQL` has the underlying type of `fromString166`:
	fromString166 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString166` to result `intoString845`
	// (`intoString845` is now tainted).
	intoString845 := template.JSEscapeString(fromString166)

	// Return the tainted `intoString845`:
	return intoString845
}

func TaintStepTest_HtmlTemplateJSEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface419` into `intoString883`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface419`:
	fromInterface419 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface419` to result `intoString883`
	// (`intoString883` is now tainted).
	intoString883 := template.JSEscaper(fromInterface419)

	// Return the tainted `intoString883`:
	return intoString883
}

func TaintStepTest_HtmlTemplateMust_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate812` into `intoTemplate363`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate812`:
	fromTemplate812 := sourceCQL.(*template.Template)

	// Call the function that transfers the taint
	// from the parameter `fromTemplate812` to result `intoTemplate363`
	// (`intoTemplate363` is now tainted).
	intoTemplate363 := template.Must(fromTemplate812, nil)

	// Return the tainted `intoTemplate363`:
	return intoTemplate363
}

func TaintStepTest_HtmlTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString222` into `intoTemplate241`.

	// Assume that `sourceCQL` has the underlying type of `fromString222`:
	fromString222 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString222` to result `intoTemplate241`
	// (`intoTemplate241` is now tainted).
	intoTemplate241, _ := template.ParseFiles(fromString222)

	// Return the tainted `intoTemplate241`:
	return intoTemplate241
}

func TaintStepTest_HtmlTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString583` into `intoTemplate934`.

	// Assume that `sourceCQL` has the underlying type of `fromString583`:
	fromString583 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString583` to result `intoTemplate934`
	// (`intoTemplate934` is now tainted).
	intoTemplate934, _ := template.ParseGlob(fromString583)

	// Return the tainted `intoTemplate934`:
	return intoTemplate934
}

func TaintStepTest_HtmlTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface408` into `intoString290`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface408`:
	fromInterface408 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface408` to result `intoString290`
	// (`intoString290` is now tainted).
	intoString290 := template.URLQueryEscaper(fromInterface408)

	// Return the tainted `intoString290`:
	return intoString290
}

func TaintStepTest_HtmlTemplateTemplateClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate530` into `intoTemplate637`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate530`:
	fromTemplate530 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate530` to the result `intoTemplate637`
	// (`intoTemplate637` is now tainted).
	intoTemplate637, _ := fromTemplate530.Clone()

	// Return the tainted `intoTemplate637`:
	return intoTemplate637
}

func TaintStepTest_HtmlTemplateTemplateDefinedTemplates_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate735` into `intoString865`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate735`:
	fromTemplate735 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate735` to the result `intoString865`
	// (`intoString865` is now tainted).
	intoString865 := fromTemplate735.DefinedTemplates()

	// Return the tainted `intoString865`:
	return intoString865
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate148` into `intoTemplate429`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate148`:
	fromTemplate148 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate148` to the result `intoTemplate429`
	// (`intoTemplate429` is now tainted).
	intoTemplate429 := fromTemplate148.Delims("", "")

	// Return the tainted `intoTemplate429`:
	return intoTemplate429
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString501` into `intoTemplate412`.

	// Assume that `sourceCQL` has the underlying type of `fromString501`:
	fromString501 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString501` to the result `intoTemplate412`
	// (`intoTemplate412` is now tainted).
	intoTemplate412 := mediumObjCQL.Delims(fromString501, "")

	// Return the tainted `intoTemplate412`:
	return intoTemplate412
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString335` into `intoTemplate651`.

	// Assume that `sourceCQL` has the underlying type of `fromString335`:
	fromString335 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString335` to the result `intoTemplate651`
	// (`intoTemplate651` is now tainted).
	intoTemplate651 := mediumObjCQL.Delims("", fromString335)

	// Return the tainted `intoTemplate651`:
	return intoTemplate651
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate154` into `intoWriter124`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate154`:
	fromTemplate154 := sourceCQL.(template.Template)

	// Declare `intoWriter124` variable:
	var intoWriter124 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate154` to the argument `intoWriter124`
	// (`intoWriter124` is now tainted).
	fromTemplate154.Execute(intoWriter124, nil)

	// Return the tainted `intoWriter124`:
	return intoWriter124
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface342` into `intoWriter713`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface342`:
	fromInterface342 := sourceCQL.(interface{})

	// Declare `intoWriter713` variable:
	var intoWriter713 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface342` to the parameter `intoWriter713`
	// (`intoWriter713` is now tainted).
	mediumObjCQL.Execute(intoWriter713, fromInterface342)

	// Return the tainted `intoWriter713`:
	return intoWriter713
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate618` into `intoWriter378`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate618`:
	fromTemplate618 := sourceCQL.(template.Template)

	// Declare `intoWriter378` variable:
	var intoWriter378 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate618` to the argument `intoWriter378`
	// (`intoWriter378` is now tainted).
	fromTemplate618.ExecuteTemplate(intoWriter378, "", nil)

	// Return the tainted `intoWriter378`:
	return intoWriter378
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface345` into `intoWriter493`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface345`:
	fromInterface345 := sourceCQL.(interface{})

	// Declare `intoWriter493` variable:
	var intoWriter493 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface345` to the parameter `intoWriter493`
	// (`intoWriter493` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter493, "", fromInterface345)

	// Return the tainted `intoWriter493`:
	return intoWriter493
}

func TaintStepTest_HtmlTemplateTemplateLookup_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate853` into `intoTemplate375`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate853`:
	fromTemplate853 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate853` to the result `intoTemplate375`
	// (`intoTemplate375` is now tainted).
	intoTemplate375 := fromTemplate853.Lookup("")

	// Return the tainted `intoTemplate375`:
	return intoTemplate375
}

func TaintStepTest_HtmlTemplateTemplateNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate248` into `intoTemplate844`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate248`:
	fromTemplate248 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate248` to the result `intoTemplate844`
	// (`intoTemplate844` is now tainted).
	intoTemplate844 := fromTemplate248.New("")

	// Return the tainted `intoTemplate844`:
	return intoTemplate844
}

func TaintStepTest_HtmlTemplateTemplateOption_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate377` into `intoTemplate360`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate377`:
	fromTemplate377 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate377` to the result `intoTemplate360`
	// (`intoTemplate360` is now tainted).
	intoTemplate360 := fromTemplate377.Option("")

	// Return the tainted `intoTemplate360`:
	return intoTemplate360
}

func TaintStepTest_HtmlTemplateTemplateParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString837` into `intoTemplate215`.

	// Assume that `sourceCQL` has the underlying type of `fromString837`:
	fromString837 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString837` to the result `intoTemplate215`
	// (`intoTemplate215` is now tainted).
	intoTemplate215, _ := mediumObjCQL.Parse(fromString837)

	// Return the tainted `intoTemplate215`:
	return intoTemplate215
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate990` into `intoTemplate987`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate990`:
	fromTemplate990 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate990` to the result `intoTemplate987`
	// (`intoTemplate987` is now tainted).
	intoTemplate987, _ := fromTemplate990.ParseFiles("")

	// Return the tainted `intoTemplate987`:
	return intoTemplate987
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString816` into `intoTemplate769`.

	// Assume that `sourceCQL` has the underlying type of `fromString816`:
	fromString816 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString816` to the result `intoTemplate769`
	// (`intoTemplate769` is now tainted).
	intoTemplate769, _ := mediumObjCQL.ParseFiles(fromString816)

	// Return the tainted `intoTemplate769`:
	return intoTemplate769
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate139` into `intoTemplate872`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate139`:
	fromTemplate139 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate139` to the result `intoTemplate872`
	// (`intoTemplate872` is now tainted).
	intoTemplate872, _ := fromTemplate139.ParseGlob("")

	// Return the tainted `intoTemplate872`:
	return intoTemplate872
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString841` into `intoTemplate945`.

	// Assume that `sourceCQL` has the underlying type of `fromString841`:
	fromString841 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString841` to the result `intoTemplate945`
	// (`intoTemplate945` is now tainted).
	intoTemplate945, _ := mediumObjCQL.ParseGlob(fromString841)

	// Return the tainted `intoTemplate945`:
	return intoTemplate945
}

func TaintStepTest_HtmlTemplateTemplateTemplates_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate346` into `intoTemplate694`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate346`:
	fromTemplate346 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate346` to the result `intoTemplate694`
	// (`intoTemplate694` is now tainted).
	intoTemplate694 := fromTemplate346.Templates()

	// Return the tainted `intoTemplate694`:
	return intoTemplate694
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
