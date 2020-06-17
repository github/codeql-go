package main

import "io"

func TaintStepTest_IoCopy_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader557` into `intoWriter196`.

	// Assume that `sourceCQL` has the underlying type of `fromReader557`:
	fromReader557 := sourceCQL.(io.Reader)

	// Declare `intoWriter196` variable:
	var intoWriter196 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader557` to parameter `intoWriter196`;
	// `intoWriter196` is now tainted.
	io.Copy(intoWriter196, fromReader557)

	// Return the tainted `intoWriter196`:
	return intoWriter196
}

func TaintStepTest_IoCopyBuffer_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader327` into `intoWriter995`.

	// Assume that `sourceCQL` has the underlying type of `fromReader327`:
	fromReader327 := sourceCQL.(io.Reader)

	// Declare `intoWriter995` variable:
	var intoWriter995 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader327` to parameter `intoWriter995`;
	// `intoWriter995` is now tainted.
	io.CopyBuffer(intoWriter995, fromReader327, nil)

	// Return the tainted `intoWriter995`:
	return intoWriter995
}

func TaintStepTest_IoCopyN_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader285` into `intoWriter530`.

	// Assume that `sourceCQL` has the underlying type of `fromReader285`:
	fromReader285 := sourceCQL.(io.Reader)

	// Declare `intoWriter530` variable:
	var intoWriter530 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader285` to parameter `intoWriter530`;
	// `intoWriter530` is now tainted.
	io.CopyN(intoWriter530, fromReader285, 0)

	// Return the tainted `intoWriter530`:
	return intoWriter530
}

func TaintStepTest_IoLimitReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader866` into `intoReader673`.

	// Assume that `sourceCQL` has the underlying type of `fromReader866`:
	fromReader866 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader866` to result `intoReader673`
	// (`intoReader673` is now tainted).
	intoReader673 := io.LimitReader(fromReader866, 0)

	// Return the tainted `intoReader673`:
	return intoReader673
}

func TaintStepTest_IoMultiReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader920` into `intoReader439`.

	// Assume that `sourceCQL` has the underlying type of `fromReader920`:
	fromReader920 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader920` to result `intoReader439`
	// (`intoReader439` is now tainted).
	intoReader439 := io.MultiReader(fromReader920)

	// Return the tainted `intoReader439`:
	return intoReader439
}

func TaintStepTest_IoMultiWriter_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriter731` into `intoWriter508`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter731`:
	fromWriter731 := sourceCQL.(io.Writer)

	// Declare `intoWriter508` variable:
	var intoWriter508 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter508`:
	intermediateCQL := io.MultiWriter(intoWriter508)

	// Extra step (`fromWriter731` taints `intermediateCQL`, which taints `intoWriter508`:
	link(fromWriter731, intermediateCQL)

	// Return the tainted `intoWriter508`:
	return intoWriter508
}

func TaintStepTest_IoNewSectionReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt745` into `intoSectionReader743`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt745`:
	fromReaderAt745 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt745` to result `intoSectionReader743`
	// (`intoSectionReader743` is now tainted).
	intoSectionReader743 := io.NewSectionReader(fromReaderAt745, 0, 0)

	// Return the tainted `intoSectionReader743`:
	return intoSectionReader743
}

func TaintStepTest_IoPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPipeWriter935` into `intoPipeReader885`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeWriter935`:
	fromPipeWriter935 := sourceCQL.(*io.PipeWriter)

	// Call the function that transfers the taint
	// from the result `fromPipeWriter935` to result `intoPipeReader885`
	// (extra steps needed)
	intoPipeReader885, intermediateCQL := io.Pipe()

	// Extra step (`fromPipeWriter935` taints `intermediateCQL`, which taints `intoPipeReader885`:
	link(fromPipeWriter935, intermediateCQL)

	// Return the tainted `intoPipeReader885`:
	return intoPipeReader885
}

func TaintStepTest_IoReadAtLeast_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader808` into `intoByte794`.

	// Assume that `sourceCQL` has the underlying type of `fromReader808`:
	fromReader808 := sourceCQL.(io.Reader)

	// Declare `intoByte794` variable:
	var intoByte794 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader808` to parameter `intoByte794`;
	// `intoByte794` is now tainted.
	io.ReadAtLeast(fromReader808, intoByte794, 0)

	// Return the tainted `intoByte794`:
	return intoByte794
}

func TaintStepTest_IoReadFull_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader649` into `intoByte302`.

	// Assume that `sourceCQL` has the underlying type of `fromReader649`:
	fromReader649 := sourceCQL.(io.Reader)

	// Declare `intoByte302` variable:
	var intoByte302 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader649` to parameter `intoByte302`;
	// `intoByte302` is now tainted.
	io.ReadFull(fromReader649, intoByte302)

	// Return the tainted `intoByte302`:
	return intoByte302
}

func TaintStepTest_IoTeeReader_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader570` into `intoWriter471`.

	// Assume that `sourceCQL` has the underlying type of `fromReader570`:
	fromReader570 := sourceCQL.(io.Reader)

	// Declare `intoWriter471` variable:
	var intoWriter471 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader570` to parameter `intoWriter471`;
	// `intoWriter471` is now tainted.
	io.TeeReader(fromReader570, intoWriter471)

	// Return the tainted `intoWriter471`:
	return intoWriter471
}

func TaintStepTest_IoTeeReader_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader545` into `intoReader664`.

	// Assume that `sourceCQL` has the underlying type of `fromReader545`:
	fromReader545 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader545` to result `intoReader664`
	// (`intoReader664` is now tainted).
	intoReader664 := io.TeeReader(fromReader545, nil)

	// Return the tainted `intoReader664`:
	return intoReader664
}

func TaintStepTest_IoWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString348` into `intoWriter751`.

	// Assume that `sourceCQL` has the underlying type of `fromString348`:
	fromString348 := sourceCQL.(string)

	// Declare `intoWriter751` variable:
	var intoWriter751 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString348` to parameter `intoWriter751`;
	// `intoWriter751` is now tainted.
	io.WriteString(intoWriter751, fromString348)

	// Return the tainted `intoWriter751`:
	return intoWriter751
}

func TaintStepTest_IoLimitedReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromLimitedReader517` into `intoByte407`.

	// Assume that `sourceCQL` has the underlying type of `fromLimitedReader517`:
	fromLimitedReader517 := sourceCQL.(io.LimitedReader)

	// Declare `intoByte407` variable:
	var intoByte407 []byte

	// Call the method that transfers the taint
	// from the receiver `fromLimitedReader517` to the argument `intoByte407`
	// (`intoByte407` is now tainted).
	fromLimitedReader517.Read(intoByte407)

	// Return the tainted `intoByte407`:
	return intoByte407
}

func TaintStepTest_IoPipeReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPipeReader922` into `intoByte816`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeReader922`:
	fromPipeReader922 := sourceCQL.(io.PipeReader)

	// Declare `intoByte816` variable:
	var intoByte816 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPipeReader922` to the argument `intoByte816`
	// (`intoByte816` is now tainted).
	fromPipeReader922.Read(intoByte816)

	// Return the tainted `intoByte816`:
	return intoByte816
}

func TaintStepTest_IoPipeWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte274` into `intoPipeWriter437`.

	// Assume that `sourceCQL` has the underlying type of `fromByte274`:
	fromByte274 := sourceCQL.([]byte)

	// Declare `intoPipeWriter437` variable:
	var intoPipeWriter437 io.PipeWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte274` to the receiver `intoPipeWriter437`
	// (`intoPipeWriter437` is now tainted).
	intoPipeWriter437.Write(fromByte274)

	// Return the tainted `intoPipeWriter437`:
	return intoPipeWriter437
}

func TaintStepTest_IoSectionReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromSectionReader900` into `intoByte594`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader900`:
	fromSectionReader900 := sourceCQL.(io.SectionReader)

	// Declare `intoByte594` variable:
	var intoByte594 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader900` to the argument `intoByte594`
	// (`intoByte594` is now tainted).
	fromSectionReader900.Read(intoByte594)

	// Return the tainted `intoByte594`:
	return intoByte594
}

func TaintStepTest_IoSectionReaderReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromSectionReader736` into `intoByte730`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader736`:
	fromSectionReader736 := sourceCQL.(io.SectionReader)

	// Declare `intoByte730` variable:
	var intoByte730 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader736` to the argument `intoByte730`
	// (`intoByte730` is now tainted).
	fromSectionReader736.ReadAt(intoByte730, 0)

	// Return the tainted `intoByte730`:
	return intoByte730
}

func TaintStepTest_IoReaderRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader908` into `intoByte347`.

	// Assume that `sourceCQL` has the underlying type of `fromReader908`:
	fromReader908 := sourceCQL.(io.Reader)

	// Declare `intoByte347` variable:
	var intoByte347 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader908` to the argument `intoByte347`
	// (`intoByte347` is now tainted).
	fromReader908.Read(intoByte347)

	// Return the tainted `intoByte347`:
	return intoByte347
}

func TaintStepTest_IoReaderAtReadAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReaderAt600` into `intoByte406`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt600`:
	fromReaderAt600 := sourceCQL.(io.ReaderAt)

	// Declare `intoByte406` variable:
	var intoByte406 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReaderAt600` to the argument `intoByte406`
	// (`intoByte406` is now tainted).
	fromReaderAt600.ReadAt(intoByte406, 0)

	// Return the tainted `intoByte406`:
	return intoByte406
}

func TaintStepTest_IoByteReaderReadByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByteReader744` into `intoByte500`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader744`:
	fromByteReader744 := sourceCQL.(io.ByteReader)

	// Call the method that transfers the taint
	// from the receiver `fromByteReader744` to the result `intoByte500`
	// (`intoByte500` is now tainted).
	intoByte500, _ := fromByteReader744.ReadByte()

	// Return the tainted `intoByte500`:
	return intoByte500
}

func TaintStepTest_IoReaderFromReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader418` into `intoReaderFrom775`.

	// Assume that `sourceCQL` has the underlying type of `fromReader418`:
	fromReader418 := sourceCQL.(io.Reader)

	// Declare `intoReaderFrom775` variable:
	var intoReaderFrom775 io.ReaderFrom

	// Call the method that transfers the taint
	// from the parameter `fromReader418` to the receiver `intoReaderFrom775`
	// (`intoReaderFrom775` is now tainted).
	intoReaderFrom775.ReadFrom(fromReader418)

	// Return the tainted `intoReaderFrom775`:
	return intoReaderFrom775
}

func TaintStepTest_IoRuneReaderReadRune_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRuneReader915` into `intoRune847`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader915`:
	fromRuneReader915 := sourceCQL.(io.RuneReader)

	// Call the method that transfers the taint
	// from the receiver `fromRuneReader915` to the result `intoRune847`
	// (`intoRune847` is now tainted).
	intoRune847, _, _ := fromRuneReader915.ReadRune()

	// Return the tainted `intoRune847`:
	return intoRune847
}

func TaintStepTest_IoWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte114` into `intoWriter150`.

	// Assume that `sourceCQL` has the underlying type of `fromByte114`:
	fromByte114 := sourceCQL.([]byte)

	// Declare `intoWriter150` variable:
	var intoWriter150 io.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte114` to the receiver `intoWriter150`
	// (`intoWriter150` is now tainted).
	intoWriter150.Write(fromByte114)

	// Return the tainted `intoWriter150`:
	return intoWriter150
}

func TaintStepTest_IoWriterAtWriteAt_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte679` into `intoWriterAt944`.

	// Assume that `sourceCQL` has the underlying type of `fromByte679`:
	fromByte679 := sourceCQL.([]byte)

	// Declare `intoWriterAt944` variable:
	var intoWriterAt944 io.WriterAt

	// Call the method that transfers the taint
	// from the parameter `fromByte679` to the receiver `intoWriterAt944`
	// (`intoWriterAt944` is now tainted).
	intoWriterAt944.WriteAt(fromByte679, 0)

	// Return the tainted `intoWriterAt944`:
	return intoWriterAt944
}

func TaintStepTest_IoByteWriterWriteByte_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte736` into `intoByteWriter282`.

	// Assume that `sourceCQL` has the underlying type of `fromByte736`:
	fromByte736 := sourceCQL.(byte)

	// Declare `intoByteWriter282` variable:
	var intoByteWriter282 io.ByteWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte736` to the receiver `intoByteWriter282`
	// (`intoByteWriter282` is now tainted).
	intoByteWriter282.WriteByte(fromByte736)

	// Return the tainted `intoByteWriter282`:
	return intoByteWriter282
}

func TaintStepTest_IoStringWriterWriteString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString744` into `intoStringWriter359`.

	// Assume that `sourceCQL` has the underlying type of `fromString744`:
	fromString744 := sourceCQL.(string)

	// Declare `intoStringWriter359` variable:
	var intoStringWriter359 io.StringWriter

	// Call the method that transfers the taint
	// from the parameter `fromString744` to the receiver `intoStringWriter359`
	// (`intoStringWriter359` is now tainted).
	intoStringWriter359.WriteString(fromString744)

	// Return the tainted `intoStringWriter359`:
	return intoStringWriter359
}

func TaintStepTest_IoWriterToWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromWriterTo673` into `intoWriter741`.

	// Assume that `sourceCQL` has the underlying type of `fromWriterTo673`:
	fromWriterTo673 := sourceCQL.(io.WriterTo)

	// Declare `intoWriter741` variable:
	var intoWriter741 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriterTo673` to the argument `intoWriter741`
	// (`intoWriter741` is now tainted).
	fromWriterTo673.WriteTo(intoWriter741)

	// Return the tainted `intoWriter741`:
	return intoWriter741
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
