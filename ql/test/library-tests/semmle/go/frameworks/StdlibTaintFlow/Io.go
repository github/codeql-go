package main

import "io"

func TaintStepTest_IoCopy(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(io.Reader)

	// Declare `dst` variable:
	var dst io.Writer

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	io.Copy(dst, src)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_IoCopyBuffer(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(io.Reader)

	// Declare `dst` variable:
	var dst io.Writer

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	io.CopyBuffer(dst, src, nil)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_IoCopyN(sourceCQL interface{}) {
	// The flow is from `src` into `dst`.

	// Assume that `sourceCQL` has the underlying type of `src`:
	src := sourceCQL.(io.Reader)

	// Declare `dst` variable:
	var dst io.Writer

	// Call the function that transfers the taint
	// from the parameter `src` to parameter `dst`;
	// `dst` is now tainted.
	io.CopyN(dst, src, 0)

	// Sink the tainted `dst`:
	sink(dst)
}

func TaintStepTest_IoLimitReader(sourceCQL interface{}) {
	// The flow is from `r` into `into690`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into690`
	// (`into690` is now tainted).
	into690 := io.LimitReader(r, 0)

	// Sink the tainted `into690`:
	sink(into690)
}

func TaintStepTest_IoMultiReader(sourceCQL interface{}) {
	// The flow is from `readers` into `into565`.

	// Assume that `sourceCQL` has the underlying type of `readers`:
	readers := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `readers` to result `into565`
	// (`into565` is now tainted).
	into565 := io.MultiReader(readers)

	// Sink the tainted `into565`:
	sink(into565)
}

func TaintStepTest_IoMultiWriter(sourceCQL interface{}) {
	// The flow is from `from862` into `writers`.

	// Assume that `sourceCQL` has the underlying type of `from862`:
	from862 := sourceCQL.(io.Writer)

	// Declare `writers` variable:
	var writers io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `writers`:
	intermediateCQL := io.MultiWriter(writers)

	// Extra step (`from862` taints `intermediateCQL`, which taints `writers`:
	link(from862, intermediateCQL)

	// Sink the tainted `writers`:
	sink(writers)
}

func TaintStepTest_IoNewSectionReader(sourceCQL interface{}) {
	// The flow is from `r` into `into602`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `r` to result `into602`
	// (`into602` is now tainted).
	into602 := io.NewSectionReader(r, 0, 0)

	// Sink the tainted `into602`:
	sink(into602)
}

func TaintStepTest_IoPipe(sourceCQL interface{}) {
	// The flow is from `from164` into `into574`.

	// Assume that `sourceCQL` has the underlying type of `from164`:
	from164 := sourceCQL.(*io.PipeWriter)

	// Call the function that transfers the taint
	// from the result `from164` to result `into574`
	// (extra steps needed)
	into574, intermediateCQL := io.Pipe()

	// Extra step (`from164` taints `intermediateCQL`, which taints `into574`:
	link(from164, intermediateCQL)

	// Sink the tainted `into574`:
	sink(into574)
}

func TaintStepTest_IoReadAtLeast(sourceCQL interface{}) {
	// The flow is from `r` into `buf`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `buf` variable:
	var buf []byte

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `buf`;
	// `buf` is now tainted.
	io.ReadAtLeast(r, buf, 0)

	// Sink the tainted `buf`:
	sink(buf)
}

func TaintStepTest_IoReadFull(sourceCQL interface{}) {
	// The flow is from `r` into `buf`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `buf` variable:
	var buf []byte

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `buf`;
	// `buf` is now tainted.
	io.ReadFull(r, buf)

	// Sink the tainted `buf`:
	sink(buf)
}

func TaintStepTest_IoTeeReader(sourceCQL interface{}) {
	// The flow is from `r` into `w`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `r` to parameter `w`;
	// `w` is now tainted.
	io.TeeReader(r, w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_IoWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `w`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `w` variable:
	var w io.Writer

	// Call the function that transfers the taint
	// from the parameter `s` to parameter `w`;
	// `w` is now tainted.
	io.WriteString(w, s)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_IoLimitedReaderRead(sourceCQL interface{}) {
	// The flow is from `from974` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from974`:
	from974 := sourceCQL.(io.LimitedReader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from974` to the argument `p`
	// (`p` is now tainted).
	from974.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_IoPipeReaderRead(sourceCQL interface{}) {
	// The flow is from `from490` into `data`.

	// Assume that `sourceCQL` has the underlying type of `from490`:
	from490 := sourceCQL.(io.PipeReader)

	// Declare `data` variable:
	var data []byte

	// Call the method that transfers the taint
	// from the receiver `from490` to the argument `data`
	// (`data` is now tainted).
	from490.Read(data)

	// Sink the tainted `data`:
	sink(data)
}

func TaintStepTest_IoPipeWriterWrite(sourceCQL interface{}) {
	// The flow is from `data` into `into689`.

	// Assume that `sourceCQL` has the underlying type of `data`:
	data := sourceCQL.([]byte)

	// Declare `into689` variable:
	var into689 io.PipeWriter

	// Call the method that transfers the taint
	// from the parameter `data` to the receiver `into689`
	// (`into689` is now tainted).
	into689.Write(data)

	// Sink the tainted `into689`:
	sink(into689)
}

func TaintStepTest_IoSectionReaderRead(sourceCQL interface{}) {
	// The flow is from `from188` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from188`:
	from188 := sourceCQL.(io.SectionReader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from188` to the argument `p`
	// (`p` is now tainted).
	from188.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_IoSectionReaderReadAt(sourceCQL interface{}) {
	// The flow is from `from701` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from701`:
	from701 := sourceCQL.(io.SectionReader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from701` to the argument `p`
	// (`p` is now tainted).
	from701.ReadAt(p, 0)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_IoByteReaderReadByte(sourceCQL interface{}) {
	// The flow is from `from677` into `into462`.

	// Assume that `sourceCQL` has the underlying type of `from677`:
	from677 := sourceCQL.(io.ByteReader)

	// Call the method that transfers the taint
	// from the receiver `from677` to the result `into462`
	// (`into462` is now tainted).
	into462, _ := from677.ReadByte()

	// Sink the tainted `into462`:
	sink(into462)
}

func TaintStepTest_IoByteWriterWriteByte(sourceCQL interface{}) {
	// The flow is from `c` into `into735`.

	// Assume that `sourceCQL` has the underlying type of `c`:
	c := sourceCQL.(byte)

	// Declare `into735` variable:
	var into735 io.ByteWriter

	// Call the method that transfers the taint
	// from the parameter `c` to the receiver `into735`
	// (`into735` is now tainted).
	into735.WriteByte(c)

	// Sink the tainted `into735`:
	sink(into735)
}

func TaintStepTest_IoReaderRead(sourceCQL interface{}) {
	// The flow is from `from831` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from831`:
	from831 := sourceCQL.(io.Reader)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from831` to the argument `p`
	// (`p` is now tainted).
	from831.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_IoReaderAtReadAt(sourceCQL interface{}) {
	// The flow is from `from796` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from796`:
	from796 := sourceCQL.(io.ReaderAt)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from796` to the argument `p`
	// (`p` is now tainted).
	from796.ReadAt(p, 0)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_IoReaderFromReadFrom(sourceCQL interface{}) {
	// The flow is from `r` into `into630`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into630` variable:
	var into630 io.ReaderFrom

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into630`
	// (`into630` is now tainted).
	into630.ReadFrom(r)

	// Sink the tainted `into630`:
	sink(into630)
}

func TaintStepTest_IoRuneReaderReadRune(sourceCQL interface{}) {
	// The flow is from `from886` into `r`.

	// Assume that `sourceCQL` has the underlying type of `from886`:
	from886 := sourceCQL.(io.RuneReader)

	// Call the method that transfers the taint
	// from the receiver `from886` to the result `r`
	// (`r` is now tainted).
	r, _, _ := from886.ReadRune()

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_IoStringWriterWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `into455`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `into455` variable:
	var into455 io.StringWriter

	// Call the method that transfers the taint
	// from the parameter `s` to the receiver `into455`
	// (`into455` is now tainted).
	into455.WriteString(s)

	// Sink the tainted `into455`:
	sink(into455)
}

func TaintStepTest_IoWriterWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into430`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into430` variable:
	var into430 io.Writer

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into430`
	// (`into430` is now tainted).
	into430.Write(p)

	// Sink the tainted `into430`:
	sink(into430)
}

func TaintStepTest_IoWriterAtWriteAt(sourceCQL interface{}) {
	// The flow is from `p` into `into388`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into388` variable:
	var into388 io.WriterAt

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into388`
	// (`into388` is now tainted).
	into388.WriteAt(p, 0)

	// Sink the tainted `into388`:
	sink(into388)
}

func TaintStepTest_IoWriterToWriteTo(sourceCQL interface{}) {
	// The flow is from `from741` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from741`:
	from741 := sourceCQL.(io.WriterTo)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from741` to the argument `w`
	// (`w` is now tainted).
	from741.WriteTo(w)

	// Sink the tainted `w`:
	sink(w)
}

func RunAllTaints_Io(v interface{}) {
	{
		source := newSource()
		TaintStepTest_IoCopy(source)
	}
	{
		source := newSource()
		TaintStepTest_IoCopyBuffer(source)
	}
	{
		source := newSource()
		TaintStepTest_IoCopyN(source)
	}
	{
		source := newSource()
		TaintStepTest_IoLimitReader(source)
	}
	{
		source := newSource()
		TaintStepTest_IoMultiReader(source)
	}
	{
		source := newSource()
		TaintStepTest_IoMultiWriter(source)
	}
	{
		source := newSource()
		TaintStepTest_IoNewSectionReader(source)
	}
	{
		source := newSource()
		TaintStepTest_IoPipe(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReadAtLeast(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReadFull(source)
	}
	{
		source := newSource()
		TaintStepTest_IoTeeReader(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriteString(source)
	}
	{
		source := newSource()
		TaintStepTest_IoLimitedReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_IoPipeReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_IoPipeWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_IoSectionReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_IoSectionReaderReadAt(source)
	}
	{
		source := newSource()
		TaintStepTest_IoByteReaderReadByte(source)
	}
	{
		source := newSource()
		TaintStepTest_IoByteWriterWriteByte(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReaderAtReadAt(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReaderFromReadFrom(source)
	}
	{
		source := newSource()
		TaintStepTest_IoRuneReaderReadRune(source)
	}
	{
		source := newSource()
		TaintStepTest_IoStringWriterWriteString(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterAtWriteAt(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterToWriteTo(source)
	}
}
