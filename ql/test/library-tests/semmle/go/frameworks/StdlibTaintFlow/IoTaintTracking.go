// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "io"

func TaintStepTest_IoCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader516` into `intoWriter638`.

	// Assume that `sourceCQL` has the underlying type of `fromReader516`:
	fromReader516 := sourceCQL.(io.Reader)

	// Declare `intoWriter638` variable:
	var intoWriter638 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader516` to parameter `intoWriter638`;
	// `intoWriter638` is now tainted.
	io.Copy(intoWriter638, fromReader516)

	// Return the tainted `intoWriter638`:
	return intoWriter638
}

func TaintStepTest_IoCopyBuffer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader207` into `intoWriter585`.

	// Assume that `sourceCQL` has the underlying type of `fromReader207`:
	fromReader207 := sourceCQL.(io.Reader)

	// Declare `intoWriter585` variable:
	var intoWriter585 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader207` to parameter `intoWriter585`;
	// `intoWriter585` is now tainted.
	io.CopyBuffer(intoWriter585, fromReader207, nil)

	// Return the tainted `intoWriter585`:
	return intoWriter585
}

func TaintStepTest_IoCopyN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader924` into `intoWriter566`.

	// Assume that `sourceCQL` has the underlying type of `fromReader924`:
	fromReader924 := sourceCQL.(io.Reader)

	// Declare `intoWriter566` variable:
	var intoWriter566 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader924` to parameter `intoWriter566`;
	// `intoWriter566` is now tainted.
	io.CopyN(intoWriter566, fromReader924, 0)

	// Return the tainted `intoWriter566`:
	return intoWriter566
}

func TaintStepTest_IoLimitReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader387` into `intoReader834`.

	// Assume that `sourceCQL` has the underlying type of `fromReader387`:
	fromReader387 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader387` to result `intoReader834`
	// (`intoReader834` is now tainted).
	intoReader834 := io.LimitReader(fromReader387, 0)

	// Return the tainted `intoReader834`:
	return intoReader834
}

func TaintStepTest_IoMultiReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader509` into `intoReader734`.

	// Assume that `sourceCQL` has the underlying type of `fromReader509`:
	fromReader509 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader509` to result `intoReader734`
	// (`intoReader734` is now tainted).
	intoReader734 := io.MultiReader(fromReader509)

	// Return the tainted `intoReader734`:
	return intoReader734
}

func TaintStepTest_IoMultiWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter751` into `intoWriter428`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter751`:
	fromWriter751 := sourceCQL.(io.Writer)

	// Declare `intoWriter428` variable:
	var intoWriter428 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter428`:
	intermediateCQL := io.MultiWriter(intoWriter428)

	// Extra step (`fromWriter751` taints `intermediateCQL`, which taints `intoWriter428`:
	link(fromWriter751, intermediateCQL)

	// Return the tainted `intoWriter428`:
	return intoWriter428
}

func TaintStepTest_IoNewSectionReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt788` into `intoSectionReader483`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt788`:
	fromReaderAt788 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt788` to result `intoSectionReader483`
	// (`intoSectionReader483` is now tainted).
	intoSectionReader483 := io.NewSectionReader(fromReaderAt788, 0, 0)

	// Return the tainted `intoSectionReader483`:
	return intoSectionReader483
}

func TaintStepTest_IoPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPipeWriter215` into `intoPipeReader441`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeWriter215`:
	fromPipeWriter215 := sourceCQL.(*io.PipeWriter)

	// Call the function that transfers the taint
	// from the result `fromPipeWriter215` to result `intoPipeReader441`
	// (extra steps needed)
	intoPipeReader441, intermediateCQL := io.Pipe()

	// Extra step (`fromPipeWriter215` taints `intermediateCQL`, which taints `intoPipeReader441`:
	link(fromPipeWriter215, intermediateCQL)

	// Return the tainted `intoPipeReader441`:
	return intoPipeReader441
}

func TaintStepTest_IoReadAtLeast_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader662` into `intoByte945`.

	// Assume that `sourceCQL` has the underlying type of `fromReader662`:
	fromReader662 := sourceCQL.(io.Reader)

	// Declare `intoByte945` variable:
	var intoByte945 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader662` to parameter `intoByte945`;
	// `intoByte945` is now tainted.
	io.ReadAtLeast(fromReader662, intoByte945, 0)

	// Return the tainted `intoByte945`:
	return intoByte945
}

func TaintStepTest_IoReadFull_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader583` into `intoByte987`.

	// Assume that `sourceCQL` has the underlying type of `fromReader583`:
	fromReader583 := sourceCQL.(io.Reader)

	// Declare `intoByte987` variable:
	var intoByte987 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader583` to parameter `intoByte987`;
	// `intoByte987` is now tainted.
	io.ReadFull(fromReader583, intoByte987)

	// Return the tainted `intoByte987`:
	return intoByte987
}

func TaintStepTest_IoTeeReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader359` into `intoWriter557`.

	// Assume that `sourceCQL` has the underlying type of `fromReader359`:
	fromReader359 := sourceCQL.(io.Reader)

	// Declare `intoWriter557` variable:
	var intoWriter557 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader359` to parameter `intoWriter557`;
	// `intoWriter557` is now tainted.
	io.TeeReader(fromReader359, intoWriter557)

	// Return the tainted `intoWriter557`:
	return intoWriter557
}

func TaintStepTest_IoTeeReader_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader868` into `intoReader869`.

	// Assume that `sourceCQL` has the underlying type of `fromReader868`:
	fromReader868 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader868` to result `intoReader869`
	// (`intoReader869` is now tainted).
	intoReader869 := io.TeeReader(fromReader868, nil)

	// Return the tainted `intoReader869`:
	return intoReader869
}

func TaintStepTest_IoWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString252` into `intoWriter933`.

	// Assume that `sourceCQL` has the underlying type of `fromString252`:
	fromString252 := sourceCQL.(string)

	// Declare `intoWriter933` variable:
	var intoWriter933 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString252` to parameter `intoWriter933`;
	// `intoWriter933` is now tainted.
	io.WriteString(intoWriter933, fromString252)

	// Return the tainted `intoWriter933`:
	return intoWriter933
}

func TaintStepTest_IoLimitedReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLimitedReader976` into `intoByte791`.

	// Assume that `sourceCQL` has the underlying type of `fromLimitedReader976`:
	fromLimitedReader976 := sourceCQL.(io.LimitedReader)

	// Declare `intoByte791` variable:
	var intoByte791 []byte

	// Call the method that transfers the taint
	// from the receiver `fromLimitedReader976` to the argument `intoByte791`
	// (`intoByte791` is now tainted).
	fromLimitedReader976.Read(intoByte791)

	// Return the tainted `intoByte791`:
	return intoByte791
}

func TaintStepTest_IoPipeReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPipeReader943` into `intoByte330`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeReader943`:
	fromPipeReader943 := sourceCQL.(io.PipeReader)

	// Declare `intoByte330` variable:
	var intoByte330 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPipeReader943` to the argument `intoByte330`
	// (`intoByte330` is now tainted).
	fromPipeReader943.Read(intoByte330)

	// Return the tainted `intoByte330`:
	return intoByte330
}

func TaintStepTest_IoPipeWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte450` into `intoPipeWriter717`.

	// Assume that `sourceCQL` has the underlying type of `fromByte450`:
	fromByte450 := sourceCQL.([]byte)

	// Declare `intoPipeWriter717` variable:
	var intoPipeWriter717 io.PipeWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte450` to the receiver `intoPipeWriter717`
	// (`intoPipeWriter717` is now tainted).
	intoPipeWriter717.Write(fromByte450)

	// Return the tainted `intoPipeWriter717`:
	return intoPipeWriter717
}

func TaintStepTest_IoSectionReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromSectionReader498` into `intoByte474`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader498`:
	fromSectionReader498 := sourceCQL.(io.SectionReader)

	// Declare `intoByte474` variable:
	var intoByte474 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader498` to the argument `intoByte474`
	// (`intoByte474` is now tainted).
	fromSectionReader498.Read(intoByte474)

	// Return the tainted `intoByte474`:
	return intoByte474
}

func TaintStepTest_IoSectionReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromSectionReader700` into `intoByte336`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader700`:
	fromSectionReader700 := sourceCQL.(io.SectionReader)

	// Declare `intoByte336` variable:
	var intoByte336 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader700` to the argument `intoByte336`
	// (`intoByte336` is now tainted).
	fromSectionReader700.ReadAt(intoByte336, 0)

	// Return the tainted `intoByte336`:
	return intoByte336
}

func TaintStepTest_IoReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader416` into `intoByte236`.

	// Assume that `sourceCQL` has the underlying type of `fromReader416`:
	fromReader416 := sourceCQL.(io.Reader)

	// Declare `intoByte236` variable:
	var intoByte236 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader416` to the argument `intoByte236`
	// (`intoByte236` is now tainted).
	fromReader416.Read(intoByte236)

	// Return the tainted `intoByte236`:
	return intoByte236
}

func TaintStepTest_IoReaderAtReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt746` into `intoByte115`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt746`:
	fromReaderAt746 := sourceCQL.(io.ReaderAt)

	// Declare `intoByte115` variable:
	var intoByte115 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReaderAt746` to the argument `intoByte115`
	// (`intoByte115` is now tainted).
	fromReaderAt746.ReadAt(intoByte115, 0)

	// Return the tainted `intoByte115`:
	return intoByte115
}

func TaintStepTest_IoByteReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteReader657` into `intoByte587`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader657`:
	fromByteReader657 := sourceCQL.(io.ByteReader)

	// Call the method that transfers the taint
	// from the receiver `fromByteReader657` to the result `intoByte587`
	// (`intoByte587` is now tainted).
	intoByte587, _ := fromByteReader657.ReadByte()

	// Return the tainted `intoByte587`:
	return intoByte587
}

func TaintStepTest_IoReaderFromReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader736` into `intoReaderFrom844`.

	// Assume that `sourceCQL` has the underlying type of `fromReader736`:
	fromReader736 := sourceCQL.(io.Reader)

	// Declare `intoReaderFrom844` variable:
	var intoReaderFrom844 io.ReaderFrom

	// Call the method that transfers the taint
	// from the parameter `fromReader736` to the receiver `intoReaderFrom844`
	// (`intoReaderFrom844` is now tainted).
	intoReaderFrom844.ReadFrom(fromReader736)

	// Return the tainted `intoReaderFrom844`:
	return intoReaderFrom844
}

func TaintStepTest_IoRuneReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader125` into `intoRune487`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader125`:
	fromRuneReader125 := sourceCQL.(io.RuneReader)

	// Call the method that transfers the taint
	// from the receiver `fromRuneReader125` to the result `intoRune487`
	// (`intoRune487` is now tainted).
	intoRune487, _, _ := fromRuneReader125.ReadRune()

	// Return the tainted `intoRune487`:
	return intoRune487
}

func TaintStepTest_IoWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte722` into `intoWriter140`.

	// Assume that `sourceCQL` has the underlying type of `fromByte722`:
	fromByte722 := sourceCQL.([]byte)

	// Declare `intoWriter140` variable:
	var intoWriter140 io.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte722` to the receiver `intoWriter140`
	// (`intoWriter140` is now tainted).
	intoWriter140.Write(fromByte722)

	// Return the tainted `intoWriter140`:
	return intoWriter140
}

func TaintStepTest_IoWriterAtWriteAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte697` into `intoWriterAt830`.

	// Assume that `sourceCQL` has the underlying type of `fromByte697`:
	fromByte697 := sourceCQL.([]byte)

	// Declare `intoWriterAt830` variable:
	var intoWriterAt830 io.WriterAt

	// Call the method that transfers the taint
	// from the parameter `fromByte697` to the receiver `intoWriterAt830`
	// (`intoWriterAt830` is now tainted).
	intoWriterAt830.WriteAt(fromByte697, 0)

	// Return the tainted `intoWriterAt830`:
	return intoWriterAt830
}

func TaintStepTest_IoByteWriterWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte280` into `intoByteWriter203`.

	// Assume that `sourceCQL` has the underlying type of `fromByte280`:
	fromByte280 := sourceCQL.(byte)

	// Declare `intoByteWriter203` variable:
	var intoByteWriter203 io.ByteWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte280` to the receiver `intoByteWriter203`
	// (`intoByteWriter203` is now tainted).
	intoByteWriter203.WriteByte(fromByte280)

	// Return the tainted `intoByteWriter203`:
	return intoByteWriter203
}

func TaintStepTest_IoStringWriterWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString850` into `intoStringWriter985`.

	// Assume that `sourceCQL` has the underlying type of `fromString850`:
	fromString850 := sourceCQL.(string)

	// Declare `intoStringWriter985` variable:
	var intoStringWriter985 io.StringWriter

	// Call the method that transfers the taint
	// from the parameter `fromString850` to the receiver `intoStringWriter985`
	// (`intoStringWriter985` is now tainted).
	intoStringWriter985.WriteString(fromString850)

	// Return the tainted `intoStringWriter985`:
	return intoStringWriter985
}

func TaintStepTest_IoWriterToWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriterTo392` into `intoWriter596`.

	// Assume that `sourceCQL` has the underlying type of `fromWriterTo392`:
	fromWriterTo392 := sourceCQL.(io.WriterTo)

	// Declare `intoWriter596` variable:
	var intoWriter596 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriterTo392` to the argument `intoWriter596`
	// (`intoWriter596` is now tainted).
	fromWriterTo392.WriteTo(intoWriter596)

	// Return the tainted `intoWriter596`:
	return intoWriter596
}

func RunAllTaints_Io() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoCopy_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoCopyBuffer_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoCopyN_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoLimitReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoMultiReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoMultiWriter_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoNewSectionReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoPipe_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoReadAtLeast_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoReadFull_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoTeeReader_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoTeeReader_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoLimitedReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoPipeReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoPipeWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoSectionReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoSectionReaderReadAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoReaderRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoReaderAtReadAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoByteReaderReadByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoReaderFromReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoRuneReaderReadRune_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoWriterWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoWriterAtWriteAt_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoByteWriterWriteByte_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoStringWriterWriteString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_IoWriterToWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
