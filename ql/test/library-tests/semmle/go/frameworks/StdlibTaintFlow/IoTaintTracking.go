// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import "io"

func TaintStepTest_IoCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader656` into `intoWriter414`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Declare `intoWriter414` variable:
	var intoWriter414 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to parameter `intoWriter414`;
	// `intoWriter414` is now tainted.
	io.Copy(intoWriter414, fromReader656)

	// Return the tainted `intoWriter414`:
	return intoWriter414
}

func TaintStepTest_IoCopyBuffer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader518` into `intoWriter650`.

	// Assume that `sourceCQL` has the underlying type of `fromReader518`:
	fromReader518 := sourceCQL.(io.Reader)

	// Declare `intoWriter650` variable:
	var intoWriter650 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader518` to parameter `intoWriter650`;
	// `intoWriter650` is now tainted.
	io.CopyBuffer(intoWriter650, fromReader518, nil)

	// Return the tainted `intoWriter650`:
	return intoWriter650
}

func TaintStepTest_IoCopyN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader784` into `intoWriter957`.

	// Assume that `sourceCQL` has the underlying type of `fromReader784`:
	fromReader784 := sourceCQL.(io.Reader)

	// Declare `intoWriter957` variable:
	var intoWriter957 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader784` to parameter `intoWriter957`;
	// `intoWriter957` is now tainted.
	io.CopyN(intoWriter957, fromReader784, 0)

	// Return the tainted `intoWriter957`:
	return intoWriter957
}

func TaintStepTest_IoLimitReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader520` into `intoReader443`.

	// Assume that `sourceCQL` has the underlying type of `fromReader520`:
	fromReader520 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader520` to result `intoReader443`
	// (`intoReader443` is now tainted).
	intoReader443 := io.LimitReader(fromReader520, 0)

	// Return the tainted `intoReader443`:
	return intoReader443
}

func TaintStepTest_IoMultiReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader127` into `intoReader483`.

	// Assume that `sourceCQL` has the underlying type of `fromReader127`:
	fromReader127 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader127` to result `intoReader483`
	// (`intoReader483` is now tainted).
	intoReader483 := io.MultiReader(fromReader127)

	// Return the tainted `intoReader483`:
	return intoReader483
}

func TaintStepTest_IoMultiWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter989` into `intoWriter982`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter989`:
	fromWriter989 := sourceCQL.(io.Writer)

	// Declare `intoWriter982` variable:
	var intoWriter982 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter982`:
	intermediateCQL := io.MultiWriter(intoWriter982)

	// Extra step (`fromWriter989` taints `intermediateCQL`, which taints `intoWriter982`:
	link(fromWriter989, intermediateCQL)

	// Return the tainted `intoWriter982`:
	return intoWriter982
}

func TaintStepTest_IoNewSectionReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt417` into `intoSectionReader584`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt417`:
	fromReaderAt417 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt417` to result `intoSectionReader584`
	// (`intoSectionReader584` is now tainted).
	intoSectionReader584 := io.NewSectionReader(fromReaderAt417, 0, 0)

	// Return the tainted `intoSectionReader584`:
	return intoSectionReader584
}

func TaintStepTest_IoPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPipeWriter991` into `intoPipeReader881`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeWriter991`:
	fromPipeWriter991 := sourceCQL.(*io.PipeWriter)

	// Call the function that transfers the taint
	// from the result `fromPipeWriter991` to result `intoPipeReader881`
	// (extra steps needed)
	intoPipeReader881, intermediateCQL := io.Pipe()

	// Extra step (`fromPipeWriter991` taints `intermediateCQL`, which taints `intoPipeReader881`:
	link(fromPipeWriter991, intermediateCQL)

	// Return the tainted `intoPipeReader881`:
	return intoPipeReader881
}

func TaintStepTest_IoReadAtLeast_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader186` into `intoByte284`.

	// Assume that `sourceCQL` has the underlying type of `fromReader186`:
	fromReader186 := sourceCQL.(io.Reader)

	// Declare `intoByte284` variable:
	var intoByte284 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader186` to parameter `intoByte284`;
	// `intoByte284` is now tainted.
	io.ReadAtLeast(fromReader186, intoByte284, 0)

	// Return the tainted `intoByte284`:
	return intoByte284
}

func TaintStepTest_IoReadFull_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader908` into `intoByte137`.

	// Assume that `sourceCQL` has the underlying type of `fromReader908`:
	fromReader908 := sourceCQL.(io.Reader)

	// Declare `intoByte137` variable:
	var intoByte137 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader908` to parameter `intoByte137`;
	// `intoByte137` is now tainted.
	io.ReadFull(fromReader908, intoByte137)

	// Return the tainted `intoByte137`:
	return intoByte137
}

func TaintStepTest_IoTeeReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader494` into `intoWriter873`.

	// Assume that `sourceCQL` has the underlying type of `fromReader494`:
	fromReader494 := sourceCQL.(io.Reader)

	// Declare `intoWriter873` variable:
	var intoWriter873 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader494` to parameter `intoWriter873`;
	// `intoWriter873` is now tainted.
	io.TeeReader(fromReader494, intoWriter873)

	// Return the tainted `intoWriter873`:
	return intoWriter873
}

func TaintStepTest_IoTeeReader_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader599` into `intoReader409`.

	// Assume that `sourceCQL` has the underlying type of `fromReader599`:
	fromReader599 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader599` to result `intoReader409`
	// (`intoReader409` is now tainted).
	intoReader409 := io.TeeReader(fromReader599, nil)

	// Return the tainted `intoReader409`:
	return intoReader409
}

func TaintStepTest_IoWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString246` into `intoWriter898`.

	// Assume that `sourceCQL` has the underlying type of `fromString246`:
	fromString246 := sourceCQL.(string)

	// Declare `intoWriter898` variable:
	var intoWriter898 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString246` to parameter `intoWriter898`;
	// `intoWriter898` is now tainted.
	io.WriteString(intoWriter898, fromString246)

	// Return the tainted `intoWriter898`:
	return intoWriter898
}

func TaintStepTest_IoLimitedReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLimitedReader598` into `intoByte631`.

	// Assume that `sourceCQL` has the underlying type of `fromLimitedReader598`:
	fromLimitedReader598 := sourceCQL.(io.LimitedReader)

	// Declare `intoByte631` variable:
	var intoByte631 []byte

	// Call the method that transfers the taint
	// from the receiver `fromLimitedReader598` to the argument `intoByte631`
	// (`intoByte631` is now tainted).
	fromLimitedReader598.Read(intoByte631)

	// Return the tainted `intoByte631`:
	return intoByte631
}

func TaintStepTest_IoPipeReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPipeReader165` into `intoByte150`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeReader165`:
	fromPipeReader165 := sourceCQL.(io.PipeReader)

	// Declare `intoByte150` variable:
	var intoByte150 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPipeReader165` to the argument `intoByte150`
	// (`intoByte150` is now tainted).
	fromPipeReader165.Read(intoByte150)

	// Return the tainted `intoByte150`:
	return intoByte150
}

func TaintStepTest_IoPipeWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte340` into `intoPipeWriter471`.

	// Assume that `sourceCQL` has the underlying type of `fromByte340`:
	fromByte340 := sourceCQL.([]byte)

	// Declare `intoPipeWriter471` variable:
	var intoPipeWriter471 io.PipeWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte340` to the receiver `intoPipeWriter471`
	// (`intoPipeWriter471` is now tainted).
	intoPipeWriter471.Write(fromByte340)

	// Return the tainted `intoPipeWriter471`:
	return intoPipeWriter471
}

func TaintStepTest_IoSectionReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromSectionReader290` into `intoByte758`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader290`:
	fromSectionReader290 := sourceCQL.(io.SectionReader)

	// Declare `intoByte758` variable:
	var intoByte758 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader290` to the argument `intoByte758`
	// (`intoByte758` is now tainted).
	fromSectionReader290.Read(intoByte758)

	// Return the tainted `intoByte758`:
	return intoByte758
}

func TaintStepTest_IoSectionReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromSectionReader396` into `intoByte707`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader396`:
	fromSectionReader396 := sourceCQL.(io.SectionReader)

	// Declare `intoByte707` variable:
	var intoByte707 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader396` to the argument `intoByte707`
	// (`intoByte707` is now tainted).
	fromSectionReader396.ReadAt(intoByte707, 0)

	// Return the tainted `intoByte707`:
	return intoByte707
}

func TaintStepTest_IoReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader912` into `intoByte718`.

	// Assume that `sourceCQL` has the underlying type of `fromReader912`:
	fromReader912 := sourceCQL.(io.Reader)

	// Declare `intoByte718` variable:
	var intoByte718 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader912` to the argument `intoByte718`
	// (`intoByte718` is now tainted).
	fromReader912.Read(intoByte718)

	// Return the tainted `intoByte718`:
	return intoByte718
}

func TaintStepTest_IoReaderAtReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt972` into `intoByte633`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt972`:
	fromReaderAt972 := sourceCQL.(io.ReaderAt)

	// Declare `intoByte633` variable:
	var intoByte633 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReaderAt972` to the argument `intoByte633`
	// (`intoByte633` is now tainted).
	fromReaderAt972.ReadAt(intoByte633, 0)

	// Return the tainted `intoByte633`:
	return intoByte633
}

func TaintStepTest_IoByteReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteReader316` into `intoByte145`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader316`:
	fromByteReader316 := sourceCQL.(io.ByteReader)

	// Call the method that transfers the taint
	// from the receiver `fromByteReader316` to the result `intoByte145`
	// (`intoByte145` is now tainted).
	intoByte145, _ := fromByteReader316.ReadByte()

	// Return the tainted `intoByte145`:
	return intoByte145
}

func TaintStepTest_IoReaderFromReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader817` into `intoReaderFrom474`.

	// Assume that `sourceCQL` has the underlying type of `fromReader817`:
	fromReader817 := sourceCQL.(io.Reader)

	// Declare `intoReaderFrom474` variable:
	var intoReaderFrom474 io.ReaderFrom

	// Call the method that transfers the taint
	// from the parameter `fromReader817` to the receiver `intoReaderFrom474`
	// (`intoReaderFrom474` is now tainted).
	intoReaderFrom474.ReadFrom(fromReader817)

	// Return the tainted `intoReaderFrom474`:
	return intoReaderFrom474
}

func TaintStepTest_IoRuneReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader832` into `intoRune378`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader832`:
	fromRuneReader832 := sourceCQL.(io.RuneReader)

	// Call the method that transfers the taint
	// from the receiver `fromRuneReader832` to the result `intoRune378`
	// (`intoRune378` is now tainted).
	intoRune378, _, _ := fromRuneReader832.ReadRune()

	// Return the tainted `intoRune378`:
	return intoRune378
}

func TaintStepTest_IoWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte541` into `intoWriter139`.

	// Assume that `sourceCQL` has the underlying type of `fromByte541`:
	fromByte541 := sourceCQL.([]byte)

	// Declare `intoWriter139` variable:
	var intoWriter139 io.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte541` to the receiver `intoWriter139`
	// (`intoWriter139` is now tainted).
	intoWriter139.Write(fromByte541)

	// Return the tainted `intoWriter139`:
	return intoWriter139
}

func TaintStepTest_IoWriterAtWriteAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte814` into `intoWriterAt768`.

	// Assume that `sourceCQL` has the underlying type of `fromByte814`:
	fromByte814 := sourceCQL.([]byte)

	// Declare `intoWriterAt768` variable:
	var intoWriterAt768 io.WriterAt

	// Call the method that transfers the taint
	// from the parameter `fromByte814` to the receiver `intoWriterAt768`
	// (`intoWriterAt768` is now tainted).
	intoWriterAt768.WriteAt(fromByte814, 0)

	// Return the tainted `intoWriterAt768`:
	return intoWriterAt768
}

func TaintStepTest_IoByteWriterWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte468` into `intoByteWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromByte468`:
	fromByte468 := sourceCQL.(byte)

	// Declare `intoByteWriter736` variable:
	var intoByteWriter736 io.ByteWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte468` to the receiver `intoByteWriter736`
	// (`intoByteWriter736` is now tainted).
	intoByteWriter736.WriteByte(fromByte468)

	// Return the tainted `intoByteWriter736`:
	return intoByteWriter736
}

func TaintStepTest_IoStringWriterWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString516` into `intoStringWriter246`.

	// Assume that `sourceCQL` has the underlying type of `fromString516`:
	fromString516 := sourceCQL.(string)

	// Declare `intoStringWriter246` variable:
	var intoStringWriter246 io.StringWriter

	// Call the method that transfers the taint
	// from the parameter `fromString516` to the receiver `intoStringWriter246`
	// (`intoStringWriter246` is now tainted).
	intoStringWriter246.WriteString(fromString516)

	// Return the tainted `intoStringWriter246`:
	return intoStringWriter246
}

func TaintStepTest_IoWriterToWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriterTo679` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromWriterTo679`:
	fromWriterTo679 := sourceCQL.(io.WriterTo)

	// Declare `intoWriter736` variable:
	var intoWriter736 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriterTo679` to the argument `intoWriter736`
	// (`intoWriter736` is now tainted).
	fromWriterTo679.WriteTo(intoWriter736)

	// Return the tainted `intoWriter736`:
	return intoWriter736
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
