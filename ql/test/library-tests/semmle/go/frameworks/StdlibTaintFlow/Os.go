package main

import "os"

func TaintStepTest_OsExpand(sourceCQL interface{}) {
	// The flow is from `s` into `into594`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into594`
	// (`into594` is now tainted).
	into594 := os.Expand(s, nil)

	// Sink the tainted `into594`:
	sink(into594)
}

func TaintStepTest_OsExpandEnv(sourceCQL interface{}) {
	// The flow is from `s` into `into566`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into566`
	// (`into566` is now tainted).
	into566 := os.ExpandEnv(s)

	// Sink the tainted `into566`:
	sink(into566)
}

func TaintStepTest_OsNewFile(sourceCQL interface{}) {
	// The flow is from `fd` into `into816`.

	// Assume that `sourceCQL` has the underlying type of `fd`:
	fd := sourceCQL.(uintptr)

	// Call the function that transfers the taint
	// from the parameter `fd` to result `into816`
	// (`into816` is now tainted).
	into816 := os.NewFile(fd, "")

	// Sink the tainted `into816`:
	sink(into816)
}

func TaintStepTest_OsPipe(sourceCQL interface{}) {
	// The flow is from `w` into `r`.

	// Assume that `sourceCQL` has the underlying type of `w`:
	w := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the result `w` to result `r`
	// (extra steps needed)
	r, intermediateCQL, _ := os.Pipe()

	// Extra step (`w` taints `intermediateCQL`, which taints `r`:
	link(w, intermediateCQL)

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_OsFileFd(sourceCQL interface{}) {
	// The flow is from `from656` into `into490`.

	// Assume that `sourceCQL` has the underlying type of `from656`:
	from656 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `from656` to the result `into490`
	// (`into490` is now tainted).
	into490 := from656.Fd()

	// Sink the tainted `into490`:
	sink(into490)
}

func TaintStepTest_OsFileRead(sourceCQL interface{}) {
	// The flow is from `from691` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from691`:
	from691 := sourceCQL.(os.File)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from691` to the argument `b`
	// (`b` is now tainted).
	from691.Read(b)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_OsFileReadAt(sourceCQL interface{}) {
	// The flow is from `from477` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from477`:
	from477 := sourceCQL.(os.File)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from477` to the argument `b`
	// (`b` is now tainted).
	from477.ReadAt(b, 0)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_OsFileSyscallConn(sourceCQL interface{}) {
	// The flow is from `from646` into `into314`.

	// Assume that `sourceCQL` has the underlying type of `from646`:
	from646 := sourceCQL.(os.File)

	// Call the method that transfers the taint
	// from the receiver `from646` to the result `into314`
	// (`into314` is now tainted).
	into314, _ := from646.SyscallConn()

	// Sink the tainted `into314`:
	sink(into314)
}

func TaintStepTest_OsFileWrite(sourceCQL interface{}) {
	// The flow is from `b` into `into594`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `into594` variable:
	var into594 os.File

	// Call the method that transfers the taint
	// from the parameter `b` to the receiver `into594`
	// (`into594` is now tainted).
	into594.Write(b)

	// Sink the tainted `into594`:
	sink(into594)
}

func TaintStepTest_OsFileWriteAt(sourceCQL interface{}) {
	// The flow is from `b` into `into492`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `into492` variable:
	var into492 os.File

	// Call the method that transfers the taint
	// from the parameter `b` to the receiver `into492`
	// (`into492` is now tainted).
	into492.WriteAt(b, 0)

	// Sink the tainted `into492`:
	sink(into492)
}

func TaintStepTest_OsFileWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `into322`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `into322` variable:
	var into322 os.File

	// Call the method that transfers the taint
	// from the parameter `s` to the receiver `into322`
	// (`into322` is now tainted).
	into322.WriteString(s)

	// Sink the tainted `into322`:
	sink(into322)
}

func RunAllTaints_Os(v interface{}) {
	{
		source := newSource()
		TaintStepTest_OsExpand(source)
	}
	{
		source := newSource()
		TaintStepTest_OsExpandEnv(source)
	}
	{
		source := newSource()
		TaintStepTest_OsNewFile(source)
	}
	{
		source := newSource()
		TaintStepTest_OsPipe(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileFd(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileRead(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileReadAt(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileSyscallConn(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileWriteAt(source)
	}
	{
		source := newSource()
		TaintStepTest_OsFileWriteString(source)
	}
}
