// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"html/template"
	"io"
)

func TaintStepTest_HtmlTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateJSEscape_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateJSEscaper_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateMust_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate417` into `intoTemplate584`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate417`:
	fromTemplate417 := sourceCQL.(*template.Template)

	// Call the function that transfers the taint
	// from the parameter `fromTemplate417` to result `intoTemplate584`
	// (`intoTemplate584` is now tainted).
	intoTemplate584 := template.Must(fromTemplate417, nil)

	// Return the tainted `intoTemplate584`:
	return intoTemplate584
}

func TaintStepTest_HtmlTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) interface{} {
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

func TaintStepTest_HtmlTemplateTemplateClone_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate494` into `intoTemplate873`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate494`:
	fromTemplate494 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate494` to the result `intoTemplate873`
	// (`intoTemplate873` is now tainted).
	intoTemplate873, _ := fromTemplate494.Clone()

	// Return the tainted `intoTemplate873`:
	return intoTemplate873
}

func TaintStepTest_HtmlTemplateTemplateDefinedTemplates_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate599` into `intoString409`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate599`:
	fromTemplate599 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate599` to the result `intoString409`
	// (`intoString409` is now tainted).
	intoString409 := fromTemplate599.DefinedTemplates()

	// Return the tainted `intoString409`:
	return intoString409
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate246` into `intoTemplate898`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate246`:
	fromTemplate246 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate246` to the result `intoTemplate898`
	// (`intoTemplate898` is now tainted).
	intoTemplate898 := fromTemplate246.Delims("", "")

	// Return the tainted `intoTemplate898`:
	return intoTemplate898
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString598` into `intoTemplate631`.

	// Assume that `sourceCQL` has the underlying type of `fromString598`:
	fromString598 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString598` to the result `intoTemplate631`
	// (`intoTemplate631` is now tainted).
	intoTemplate631 := mediumObjCQL.Delims(fromString598, "")

	// Return the tainted `intoTemplate631`:
	return intoTemplate631
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoTemplate150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString165` to the result `intoTemplate150`
	// (`intoTemplate150` is now tainted).
	intoTemplate150 := mediumObjCQL.Delims("", fromString165)

	// Return the tainted `intoTemplate150`:
	return intoTemplate150
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate340` into `intoWriter471`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate340`:
	fromTemplate340 := sourceCQL.(template.Template)

	// Declare `intoWriter471` variable:
	var intoWriter471 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate340` to the argument `intoWriter471`
	// (`intoWriter471` is now tainted).
	fromTemplate340.Execute(intoWriter471, nil)

	// Return the tainted `intoWriter471`:
	return intoWriter471
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface290` into `intoWriter758`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface290`:
	fromInterface290 := sourceCQL.(interface{})

	// Declare `intoWriter758` variable:
	var intoWriter758 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface290` to the parameter `intoWriter758`
	// (`intoWriter758` is now tainted).
	mediumObjCQL.Execute(intoWriter758, fromInterface290)

	// Return the tainted `intoWriter758`:
	return intoWriter758
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate396` into `intoWriter707`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate396`:
	fromTemplate396 := sourceCQL.(template.Template)

	// Declare `intoWriter707` variable:
	var intoWriter707 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate396` to the argument `intoWriter707`
	// (`intoWriter707` is now tainted).
	fromTemplate396.ExecuteTemplate(intoWriter707, "", nil)

	// Return the tainted `intoWriter707`:
	return intoWriter707
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromInterface912` into `intoWriter718`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface912`:
	fromInterface912 := sourceCQL.(interface{})

	// Declare `intoWriter718` variable:
	var intoWriter718 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface912` to the parameter `intoWriter718`
	// (`intoWriter718` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter718, "", fromInterface912)

	// Return the tainted `intoWriter718`:
	return intoWriter718
}

func TaintStepTest_HtmlTemplateTemplateLookup_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate972` into `intoTemplate633`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate972`:
	fromTemplate972 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate972` to the result `intoTemplate633`
	// (`intoTemplate633` is now tainted).
	intoTemplate633 := fromTemplate972.Lookup("")

	// Return the tainted `intoTemplate633`:
	return intoTemplate633
}

func TaintStepTest_HtmlTemplateTemplateNew_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate316` into `intoTemplate145`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate316`:
	fromTemplate316 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate316` to the result `intoTemplate145`
	// (`intoTemplate145` is now tainted).
	intoTemplate145 := fromTemplate316.New("")

	// Return the tainted `intoTemplate145`:
	return intoTemplate145
}

func TaintStepTest_HtmlTemplateTemplateOption_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate817` into `intoTemplate474`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate817`:
	fromTemplate817 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate817` to the result `intoTemplate474`
	// (`intoTemplate474` is now tainted).
	intoTemplate474 := fromTemplate817.Option("")

	// Return the tainted `intoTemplate474`:
	return intoTemplate474
}

func TaintStepTest_HtmlTemplateTemplateParse_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString832` into `intoTemplate378`.

	// Assume that `sourceCQL` has the underlying type of `fromString832`:
	fromString832 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString832` to the result `intoTemplate378`
	// (`intoTemplate378` is now tainted).
	intoTemplate378, _ := mediumObjCQL.Parse(fromString832)

	// Return the tainted `intoTemplate378`:
	return intoTemplate378
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate541` into `intoTemplate139`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate541`:
	fromTemplate541 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate541` to the result `intoTemplate139`
	// (`intoTemplate139` is now tainted).
	intoTemplate139, _ := fromTemplate541.ParseFiles("")

	// Return the tainted `intoTemplate139`:
	return intoTemplate139
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString814` into `intoTemplate768`.

	// Assume that `sourceCQL` has the underlying type of `fromString814`:
	fromString814 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString814` to the result `intoTemplate768`
	// (`intoTemplate768` is now tainted).
	intoTemplate768, _ := mediumObjCQL.ParseFiles(fromString814)

	// Return the tainted `intoTemplate768`:
	return intoTemplate768
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate468` into `intoTemplate736`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate468`:
	fromTemplate468 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate468` to the result `intoTemplate736`
	// (`intoTemplate736` is now tainted).
	intoTemplate736, _ := fromTemplate468.ParseGlob("")

	// Return the tainted `intoTemplate736`:
	return intoTemplate736
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoTemplate246`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString516` to the result `intoTemplate246`
	// (`intoTemplate246` is now tainted).
	intoTemplate246, _ := mediumObjCQL.ParseGlob(fromString516)

	// Return the tainted `intoTemplate246`:
	return intoTemplate246
}

func TaintStepTest_HtmlTemplateTemplateTemplates_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTemplate679` into `intoTemplate736`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate679`:
	fromTemplate679 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate679` to the result `intoTemplate736`
	// (`intoTemplate736` is now tainted).
	intoTemplate736 := fromTemplate679.Templates()

	// Return the tainted `intoTemplate736`:
	return intoTemplate736
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
