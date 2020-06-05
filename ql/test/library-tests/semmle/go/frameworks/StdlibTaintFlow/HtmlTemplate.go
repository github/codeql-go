package main

import (
	"html/template"
	"io"
)

func TaintStepTest_HtmlTemplateHTMLEscape(sourceCQL interface{}) {
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

func TaintStepTest_HtmlTemplateHTMLEscapeString(sourceCQL interface{}) {
	// The flow is from `s` into `into286`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into286`
	// (`into286` is now tainted).
	into286 := template.HTMLEscapeString(s)

	// Sink the tainted `into286`:
	sink(into286)
}

func TaintStepTest_HtmlTemplateHTMLEscaper(sourceCQL interface{}) {
	// The flow is from `args` into `into924`.

	// Assume that `sourceCQL` has the underlying type of `args`:
	args := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `args` to result `into924`
	// (`into924` is now tainted).
	into924 := template.HTMLEscaper(args)

	// Sink the tainted `into924`:
	sink(into924)
}

func TaintStepTest_HtmlTemplateJSEscape(sourceCQL interface{}) {
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

func TaintStepTest_HtmlTemplateJSEscapeString(sourceCQL interface{}) {
	// The flow is from `s` into `into627`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into627`
	// (`into627` is now tainted).
	into627 := template.JSEscapeString(s)

	// Sink the tainted `into627`:
	sink(into627)
}

func TaintStepTest_HtmlTemplateJSEscaper(sourceCQL interface{}) {
	// The flow is from `args` into `into628`.

	// Assume that `sourceCQL` has the underlying type of `args`:
	args := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `args` to result `into628`
	// (`into628` is now tainted).
	into628 := template.JSEscaper(args)

	// Sink the tainted `into628`:
	sink(into628)
}

func TaintStepTest_HtmlTemplateMust(sourceCQL interface{}) {
	// The flow is from `t` into `into135`.

	// Assume that `sourceCQL` has the underlying type of `t`:
	t := sourceCQL.(*template.Template)

	// Call the function that transfers the taint
	// from the parameter `t` to result `into135`
	// (`into135` is now tainted).
	into135 := template.Must(t, nil)

	// Sink the tainted `into135`:
	sink(into135)
}

func TaintStepTest_HtmlTemplateParseFiles(sourceCQL interface{}) {
	// The flow is from `filenames` into `into678`.

	// Assume that `sourceCQL` has the underlying type of `filenames`:
	filenames := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `filenames` to result `into678`
	// (`into678` is now tainted).
	into678, _ := template.ParseFiles(filenames)

	// Sink the tainted `into678`:
	sink(into678)
}

func TaintStepTest_HtmlTemplateParseGlob(sourceCQL interface{}) {
	// The flow is from `pattern` into `into834`.

	// Assume that `sourceCQL` has the underlying type of `pattern`:
	pattern := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `pattern` to result `into834`
	// (`into834` is now tainted).
	into834, _ := template.ParseGlob(pattern)

	// Sink the tainted `into834`:
	sink(into834)
}

func TaintStepTest_HtmlTemplateURLQueryEscaper(sourceCQL interface{}) {
	// The flow is from `args` into `into621`.

	// Assume that `sourceCQL` has the underlying type of `args`:
	args := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `args` to result `into621`
	// (`into621` is now tainted).
	into621 := template.URLQueryEscaper(args)

	// Sink the tainted `into621`:
	sink(into621)
}

func TaintStepTest_HtmlTemplateTemplateClone(sourceCQL interface{}) {
	// The flow is from `from811` into `into688`.

	// Assume that `sourceCQL` has the underlying type of `from811`:
	from811 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from811` to the result `into688`
	// (`into688` is now tainted).
	into688, _ := from811.Clone()

	// Sink the tainted `into688`:
	sink(into688)
}

func TaintStepTest_HtmlTemplateTemplateDefinedTemplates(sourceCQL interface{}) {
	// The flow is from `from539` into `into975`.

	// Assume that `sourceCQL` has the underlying type of `from539`:
	from539 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from539` to the result `into975`
	// (`into975` is now tainted).
	into975 := from539.DefinedTemplates()

	// Sink the tainted `into975`:
	sink(into975)
}

func TaintStepTest_HtmlTemplateTemplateDelims(sourceCQL interface{}) {
	// The flow is from `from307` into `into914`.

	// Assume that `sourceCQL` has the underlying type of `from307`:
	from307 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from307` to the result `into914`
	// (`into914` is now tainted).
	into914 := from307.Delims("", "")

	// Sink the tainted `into914`:
	sink(into914)
}

func TaintStepTest_HtmlTemplateTemplateExecute(sourceCQL interface{}) {
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

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate(sourceCQL interface{}) {
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

func TaintStepTest_HtmlTemplateTemplateLookup(sourceCQL interface{}) {
	// The flow is from `from168` into `into612`.

	// Assume that `sourceCQL` has the underlying type of `from168`:
	from168 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from168` to the result `into612`
	// (`into612` is now tainted).
	into612 := from168.Lookup("")

	// Sink the tainted `into612`:
	sink(into612)
}

func TaintStepTest_HtmlTemplateTemplateNew(sourceCQL interface{}) {
	// The flow is from `from336` into `into505`.

	// Assume that `sourceCQL` has the underlying type of `from336`:
	from336 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from336` to the result `into505`
	// (`into505` is now tainted).
	into505 := from336.New("")

	// Sink the tainted `into505`:
	sink(into505)
}

func TaintStepTest_HtmlTemplateTemplateOption(sourceCQL interface{}) {
	// The flow is from `from669` into `into286`.

	// Assume that `sourceCQL` has the underlying type of `from669`:
	from669 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from669` to the result `into286`
	// (`into286` is now tainted).
	into286 := from669.Option("")

	// Sink the tainted `into286`:
	sink(into286)
}

func TaintStepTest_HtmlTemplateTemplateParse(sourceCQL interface{}) {
	// The flow is from `text` into `into261`.

	// Assume that `sourceCQL` has the underlying type of `text`:
	text := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `text` to the result `into261`
	// (`into261` is now tainted).
	into261, _ := mediumObjCQL.Parse(text)

	// Sink the tainted `into261`:
	sink(into261)
}

func TaintStepTest_HtmlTemplateTemplateParseFiles(sourceCQL interface{}) {
	// The flow is from `from811` into `into696`.

	// Assume that `sourceCQL` has the underlying type of `from811`:
	from811 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from811` to the result `into696`
	// (`into696` is now tainted).
	into696, _ := from811.ParseFiles("")

	// Sink the tainted `into696`:
	sink(into696)
}

func TaintStepTest_HtmlTemplateTemplateParseGlob(sourceCQL interface{}) {
	// The flow is from `from206` into `into171`.

	// Assume that `sourceCQL` has the underlying type of `from206`:
	from206 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from206` to the result `into171`
	// (`into171` is now tainted).
	into171, _ := from206.ParseGlob("")

	// Sink the tainted `into171`:
	sink(into171)
}

func TaintStepTest_HtmlTemplateTemplateTemplates(sourceCQL interface{}) {
	// The flow is from `from374` into `into678`.

	// Assume that `sourceCQL` has the underlying type of `from374`:
	from374 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `from374` to the result `into678`
	// (`into678` is now tainted).
	into678 := from374.Templates()

	// Sink the tainted `into678`:
	sink(into678)
}

func RunAllTaints_HtmlTemplate(v interface{}) {
	{
		source := newSource()
		TaintStepTest_HtmlTemplateHTMLEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateHTMLEscapeString(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateHTMLEscaper(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateJSEscape(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateJSEscapeString(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateJSEscaper(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateMust(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateParseFiles(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateParseGlob(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateURLQueryEscaper(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateClone(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateDefinedTemplates(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateDelims(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateExecute(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateExecuteTemplate(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateLookup(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateNew(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateOption(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParse(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParseFiles(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParseGlob(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateTemplates(source)
	}
}
