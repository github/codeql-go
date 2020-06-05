package main

import (
	"io"
	"net/mail"
)

func TaintStepTest_NetMailReadMessage(sourceCQL interface{}) {
	// The flow is from `r` into `msg`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `msg`
	// (`msg` is now tainted).
	msg, _ := mail.ReadMessage(r)

	// Sink the tainted `msg`:
	sink(msg)
}

func TaintStepTest_NetMailHeaderGet(sourceCQL interface{}) {
	// The flow is from `from159` into `into122`.

	// Assume that `sourceCQL` has the underlying type of `from159`:
	from159 := sourceCQL.(mail.Header)

	// Call the method that transfers the taint
	// from the receiver `from159` to the result `into122`
	// (`into122` is now tainted).
	into122 := from159.Get("")

	// Sink the tainted `into122`:
	sink(into122)
}

func RunAllTaints_NetMail(v interface{}) {
	{
		source := newSource()
		TaintStepTest_NetMailReadMessage(source)
	}
	{
		source := newSource()
		TaintStepTest_NetMailHeaderGet(source)
	}
}
