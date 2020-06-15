package main

import (
	"html/template"
	"io"
)

func TaintStepTest_HtmlTemplateHTMLEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte183` into `intoWriter358`.

	// Assume that `sourceCQL` has the underlying type of `fromByte183`:
	fromByte183 := sourceCQL.([]byte)

	// Declare `intoWriter358` variable:
	var intoWriter358 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte183` to parameter `intoWriter358`;
	// `intoWriter358` is now tainted.
	template.HTMLEscape(intoWriter358, fromByte183)

	// Sink the tainted `intoWriter358`:
	sink(intoWriter358)
}

func TaintStepTest_HtmlTemplateHTMLEscapeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString467` into `intoString282`.

	// Assume that `sourceCQL` has the underlying type of `fromString467`:
	fromString467 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString467` to result `intoString282`
	// (`intoString282` is now tainted).
	intoString282 := template.HTMLEscapeString(fromString467)

	// Sink the tainted `intoString282`:
	sink(intoString282)
}

func TaintStepTest_HtmlTemplateHTMLEscaper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface316` into `intoString357`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface316`:
	fromInterface316 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface316` to result `intoString357`
	// (`intoString357` is now tainted).
	intoString357 := template.HTMLEscaper(fromInterface316)

	// Sink the tainted `intoString357`:
	sink(intoString357)
}

func TaintStepTest_HtmlTemplateJSEscape_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte939` into `intoWriter653`.

	// Assume that `sourceCQL` has the underlying type of `fromByte939`:
	fromByte939 := sourceCQL.([]byte)

	// Declare `intoWriter653` variable:
	var intoWriter653 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromByte939` to parameter `intoWriter653`;
	// `intoWriter653` is now tainted.
	template.JSEscape(intoWriter653, fromByte939)

	// Sink the tainted `intoWriter653`:
	sink(intoWriter653)
}

func TaintStepTest_HtmlTemplateJSEscapeString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString181` into `intoString570`.

	// Assume that `sourceCQL` has the underlying type of `fromString181`:
	fromString181 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString181` to result `intoString570`
	// (`intoString570` is now tainted).
	intoString570 := template.JSEscapeString(fromString181)

	// Sink the tainted `intoString570`:
	sink(intoString570)
}

func TaintStepTest_HtmlTemplateJSEscaper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface982` into `intoString493`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface982`:
	fromInterface982 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface982` to result `intoString493`
	// (`intoString493` is now tainted).
	intoString493 := template.JSEscaper(fromInterface982)

	// Sink the tainted `intoString493`:
	sink(intoString493)
}

func TaintStepTest_HtmlTemplateMust_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate969` into `intoTemplate120`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate969`:
	fromTemplate969 := sourceCQL.(*template.Template)

	// Call the function that transfers the taint
	// from the parameter `fromTemplate969` to result `intoTemplate120`
	// (`intoTemplate120` is now tainted).
	intoTemplate120 := template.Must(fromTemplate969, nil)

	// Sink the tainted `intoTemplate120`:
	sink(intoTemplate120)
}

func TaintStepTest_HtmlTemplateParseFiles_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString403` into `intoTemplate890`.

	// Assume that `sourceCQL` has the underlying type of `fromString403`:
	fromString403 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString403` to result `intoTemplate890`
	// (`intoTemplate890` is now tainted).
	intoTemplate890, _ := template.ParseFiles(fromString403)

	// Sink the tainted `intoTemplate890`:
	sink(intoTemplate890)
}

func TaintStepTest_HtmlTemplateParseGlob_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString352` into `intoTemplate415`.

	// Assume that `sourceCQL` has the underlying type of `fromString352`:
	fromString352 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString352` to result `intoTemplate415`
	// (`intoTemplate415` is now tainted).
	intoTemplate415, _ := template.ParseGlob(fromString352)

	// Sink the tainted `intoTemplate415`:
	sink(intoTemplate415)
}

func TaintStepTest_HtmlTemplateURLQueryEscaper_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromInterface750` into `intoString959`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface750`:
	fromInterface750 := sourceCQL.(interface{})

	// Call the function that transfers the taint
	// from the parameter `fromInterface750` to result `intoString959`
	// (`intoString959` is now tainted).
	intoString959 := template.URLQueryEscaper(fromInterface750)

	// Sink the tainted `intoString959`:
	sink(intoString959)
}

func TaintStepTest_HtmlTemplateTemplateClone_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate652` into `intoTemplate714`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate652`:
	fromTemplate652 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate652` to the result `intoTemplate714`
	// (`intoTemplate714` is now tainted).
	intoTemplate714, _ := fromTemplate652.Clone()

	// Sink the tainted `intoTemplate714`:
	sink(intoTemplate714)
}

func TaintStepTest_HtmlTemplateTemplateDefinedTemplates_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate953` into `intoString479`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate953`:
	fromTemplate953 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate953` to the result `intoString479`
	// (`intoString479` is now tainted).
	intoString479 := fromTemplate953.DefinedTemplates()

	// Sink the tainted `intoString479`:
	sink(intoString479)
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate675` into `intoTemplate971`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate675`:
	fromTemplate675 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate675` to the result `intoTemplate971`
	// (`intoTemplate971` is now tainted).
	intoTemplate971 := fromTemplate675.Delims("", "")

	// Sink the tainted `intoTemplate971`:
	sink(intoTemplate971)
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString710` into `intoTemplate750`.

	// Assume that `sourceCQL` has the underlying type of `fromString710`:
	fromString710 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString710` to the result `intoTemplate750`
	// (`intoTemplate750` is now tainted).
	intoTemplate750 := mediumObjCQL.Delims(fromString710, "")

	// Sink the tainted `intoTemplate750`:
	sink(intoTemplate750)
}

func TaintStepTest_HtmlTemplateTemplateDelims_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromString267` into `intoTemplate560`.

	// Assume that `sourceCQL` has the underlying type of `fromString267`:
	fromString267 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString267` to the result `intoTemplate560`
	// (`intoTemplate560` is now tainted).
	intoTemplate560 := mediumObjCQL.Delims("", fromString267)

	// Sink the tainted `intoTemplate560`:
	sink(intoTemplate560)
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate417` into `intoWriter150`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate417`:
	fromTemplate417 := sourceCQL.(template.Template)

	// Declare `intoWriter150` variable:
	var intoWriter150 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate417` to the argument `intoWriter150`
	// (`intoWriter150` is now tainted).
	fromTemplate417.Execute(intoWriter150, nil)

	// Sink the tainted `intoWriter150`:
	sink(intoWriter150)
}

func TaintStepTest_HtmlTemplateTemplateExecute_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface954` into `intoWriter244`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface954`:
	fromInterface954 := sourceCQL.(interface{})

	// Declare `intoWriter244` variable:
	var intoWriter244 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface954` to the parameter `intoWriter244`
	// (`intoWriter244` is now tainted).
	mediumObjCQL.Execute(intoWriter244, fromInterface954)

	// Sink the tainted `intoWriter244`:
	sink(intoWriter244)
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate556` into `intoWriter855`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate556`:
	fromTemplate556 := sourceCQL.(template.Template)

	// Declare `intoWriter855` variable:
	var intoWriter855 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromTemplate556` to the argument `intoWriter855`
	// (`intoWriter855` is now tainted).
	fromTemplate556.ExecuteTemplate(intoWriter855, "", nil)

	// Sink the tainted `intoWriter855`:
	sink(intoWriter855)
}

func TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromInterface737` into `intoWriter598`.

	// Assume that `sourceCQL` has the underlying type of `fromInterface737`:
	fromInterface737 := sourceCQL.(interface{})

	// Declare `intoWriter598` variable:
	var intoWriter598 io.Writer

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromInterface737` to the parameter `intoWriter598`
	// (`intoWriter598` is now tainted).
	mediumObjCQL.ExecuteTemplate(intoWriter598, "", fromInterface737)

	// Sink the tainted `intoWriter598`:
	sink(intoWriter598)
}

func TaintStepTest_HtmlTemplateTemplateLookup_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate623` into `intoTemplate464`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate623`:
	fromTemplate623 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate623` to the result `intoTemplate464`
	// (`intoTemplate464` is now tainted).
	intoTemplate464 := fromTemplate623.Lookup("")

	// Sink the tainted `intoTemplate464`:
	sink(intoTemplate464)
}

func TaintStepTest_HtmlTemplateTemplateNew_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate713` into `intoTemplate386`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate713`:
	fromTemplate713 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate713` to the result `intoTemplate386`
	// (`intoTemplate386` is now tainted).
	intoTemplate386 := fromTemplate713.New("")

	// Sink the tainted `intoTemplate386`:
	sink(intoTemplate386)
}

func TaintStepTest_HtmlTemplateTemplateOption_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate832` into `intoTemplate158`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate832`:
	fromTemplate832 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate832` to the result `intoTemplate158`
	// (`intoTemplate158` is now tainted).
	intoTemplate158 := fromTemplate832.Option("")

	// Sink the tainted `intoTemplate158`:
	sink(intoTemplate158)
}

func TaintStepTest_HtmlTemplateTemplateParse_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString741` into `intoTemplate444`.

	// Assume that `sourceCQL` has the underlying type of `fromString741`:
	fromString741 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString741` to the result `intoTemplate444`
	// (`intoTemplate444` is now tainted).
	intoTemplate444, _ := mediumObjCQL.Parse(fromString741)

	// Sink the tainted `intoTemplate444`:
	sink(intoTemplate444)
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate393` into `intoTemplate595`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate393`:
	fromTemplate393 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate393` to the result `intoTemplate595`
	// (`intoTemplate595` is now tainted).
	intoTemplate595, _ := fromTemplate393.ParseFiles("")

	// Sink the tainted `intoTemplate595`:
	sink(intoTemplate595)
}

func TaintStepTest_HtmlTemplateTemplateParseFiles_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString294` into `intoTemplate779`.

	// Assume that `sourceCQL` has the underlying type of `fromString294`:
	fromString294 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString294` to the result `intoTemplate779`
	// (`intoTemplate779` is now tainted).
	intoTemplate779, _ := mediumObjCQL.ParseFiles(fromString294)

	// Sink the tainted `intoTemplate779`:
	sink(intoTemplate779)
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate836` into `intoTemplate466`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate836`:
	fromTemplate836 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate836` to the result `intoTemplate466`
	// (`intoTemplate466` is now tainted).
	intoTemplate466, _ := fromTemplate836.ParseGlob("")

	// Sink the tainted `intoTemplate466`:
	sink(intoTemplate466)
}

func TaintStepTest_HtmlTemplateTemplateParseGlob_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString689` into `intoTemplate744`.

	// Assume that `sourceCQL` has the underlying type of `fromString689`:
	fromString689 := sourceCQL.(string)

	// Declare medium object/interface:
	var mediumObjCQL template.Template

	// Call the method that transfers the taint
	// from the parameter `fromString689` to the result `intoTemplate744`
	// (`intoTemplate744` is now tainted).
	intoTemplate744, _ := mediumObjCQL.ParseGlob(fromString689)

	// Sink the tainted `intoTemplate744`:
	sink(intoTemplate744)
}

func TaintStepTest_HtmlTemplateTemplateTemplates_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTemplate399` into `intoTemplate963`.

	// Assume that `sourceCQL` has the underlying type of `fromTemplate399`:
	fromTemplate399 := sourceCQL.(template.Template)

	// Call the method that transfers the taint
	// from the receiver `fromTemplate399` to the result `intoTemplate963`
	// (`intoTemplate963` is now tainted).
	intoTemplate963 := fromTemplate399.Templates()

	// Sink the tainted `intoTemplate963`:
	sink(intoTemplate963)
}

func RunAllTaints_HtmlTemplate() {
	{
		source := newSource()
		TaintStepTest_HtmlTemplateHTMLEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateHTMLEscapeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateHTMLEscaper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateJSEscape_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateJSEscapeString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateJSEscaper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateMust_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateParseFiles_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateParseGlob_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateURLQueryEscaper_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateClone_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateDefinedTemplates_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateDelims_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateDelims_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateDelims_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateExecute_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateExecute_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateExecuteTemplate_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateLookup_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateNew_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateOption_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParse_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParseFiles_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParseFiles_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParseGlob_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateParseGlob_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_HtmlTemplateTemplateTemplates_B0I0O0(source)
	}
}
