package main

import "io"

func TaintStepTest_IoCopy_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader863` into `intoWriter501`.

	// Assume that `sourceCQL` has the underlying type of `fromReader863`:
	fromReader863 := sourceCQL.(io.Reader)

	// Declare `intoWriter501` variable:
	var intoWriter501 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader863` to parameter `intoWriter501`;
	// `intoWriter501` is now tainted.
	io.Copy(intoWriter501, fromReader863)

	// Sink the tainted `intoWriter501`:
	sink(intoWriter501)
}

func TaintStepTest_IoCopyBuffer_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader362` into `intoWriter996`.

	// Assume that `sourceCQL` has the underlying type of `fromReader362`:
	fromReader362 := sourceCQL.(io.Reader)

	// Declare `intoWriter996` variable:
	var intoWriter996 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader362` to parameter `intoWriter996`;
	// `intoWriter996` is now tainted.
	io.CopyBuffer(intoWriter996, fromReader362, nil)

	// Sink the tainted `intoWriter996`:
	sink(intoWriter996)
}

func TaintStepTest_IoCopyN_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader656` into `intoWriter134`.

	// Assume that `sourceCQL` has the underlying type of `fromReader656`:
	fromReader656 := sourceCQL.(io.Reader)

	// Declare `intoWriter134` variable:
	var intoWriter134 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader656` to parameter `intoWriter134`;
	// `intoWriter134` is now tainted.
	io.CopyN(intoWriter134, fromReader656, 0)

	// Sink the tainted `intoWriter134`:
	sink(intoWriter134)
}

func TaintStepTest_IoLimitReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader827` into `intoReader644`.

	// Assume that `sourceCQL` has the underlying type of `fromReader827`:
	fromReader827 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader827` to result `intoReader644`
	// (`intoReader644` is now tainted).
	intoReader644 := io.LimitReader(fromReader827, 0)

	// Sink the tainted `intoReader644`:
	sink(intoReader644)
}

func TaintStepTest_IoMultiReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader980` into `intoReader322`.

	// Assume that `sourceCQL` has the underlying type of `fromReader980`:
	fromReader980 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader980` to result `intoReader322`
	// (`intoReader322` is now tainted).
	intoReader322 := io.MultiReader(fromReader980)

	// Sink the tainted `intoReader322`:
	sink(intoReader322)
}

func TaintStepTest_IoMultiWriter_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriter132` into `intoWriter694`.

	// Assume that `sourceCQL` has the underlying type of `fromWriter132`:
	fromWriter132 := sourceCQL.(io.Writer)

	// Declare `intoWriter694` variable:
	var intoWriter694 io.Writer

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoWriter694`:
	intermediateCQL := io.MultiWriter(intoWriter694)

	// Extra step (`fromWriter132` taints `intermediateCQL`, which taints `intoWriter694`:
	link(fromWriter132, intermediateCQL)

	// Sink the tainted `intoWriter694`:
	sink(intoWriter694)
}

func TaintStepTest_IoNewSectionReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReaderAt909` into `intoSectionReader188`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt909`:
	fromReaderAt909 := sourceCQL.(io.ReaderAt)

	// Call the function that transfers the taint
	// from the parameter `fromReaderAt909` to result `intoSectionReader188`
	// (`intoSectionReader188` is now tainted).
	intoSectionReader188 := io.NewSectionReader(fromReaderAt909, 0, 0)

	// Sink the tainted `intoSectionReader188`:
	sink(intoSectionReader188)
}

func TaintStepTest_IoPipe_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPipeWriter521` into `intoPipeReader673`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeWriter521`:
	fromPipeWriter521 := sourceCQL.(*io.PipeWriter)

	// Call the function that transfers the taint
	// from the result `fromPipeWriter521` to result `intoPipeReader673`
	// (extra steps needed)
	intoPipeReader673, intermediateCQL := io.Pipe()

	// Extra step (`fromPipeWriter521` taints `intermediateCQL`, which taints `intoPipeReader673`:
	link(fromPipeWriter521, intermediateCQL)

	// Sink the tainted `intoPipeReader673`:
	sink(intoPipeReader673)
}

func TaintStepTest_IoReadAtLeast_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader678` into `intoByte206`.

	// Assume that `sourceCQL` has the underlying type of `fromReader678`:
	fromReader678 := sourceCQL.(io.Reader)

	// Declare `intoByte206` variable:
	var intoByte206 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader678` to parameter `intoByte206`;
	// `intoByte206` is now tainted.
	io.ReadAtLeast(fromReader678, intoByte206, 0)

	// Sink the tainted `intoByte206`:
	sink(intoByte206)
}

func TaintStepTest_IoReadFull_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader945` into `intoByte478`.

	// Assume that `sourceCQL` has the underlying type of `fromReader945`:
	fromReader945 := sourceCQL.(io.Reader)

	// Declare `intoByte478` variable:
	var intoByte478 []byte

	// Call the function that transfers the taint
	// from the parameter `fromReader945` to parameter `intoByte478`;
	// `intoByte478` is now tainted.
	io.ReadFull(fromReader945, intoByte478)

	// Sink the tainted `intoByte478`:
	sink(intoByte478)
}

func TaintStepTest_IoTeeReader_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader234` into `intoWriter475`.

	// Assume that `sourceCQL` has the underlying type of `fromReader234`:
	fromReader234 := sourceCQL.(io.Reader)

	// Declare `intoWriter475` variable:
	var intoWriter475 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromReader234` to parameter `intoWriter475`;
	// `intoWriter475` is now tainted.
	io.TeeReader(fromReader234, intoWriter475)

	// Sink the tainted `intoWriter475`:
	sink(intoWriter475)
}

func TaintStepTest_IoTeeReader_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromReader673` into `intoReader671`.

	// Assume that `sourceCQL` has the underlying type of `fromReader673`:
	fromReader673 := sourceCQL.(io.Reader)

	// Call the function that transfers the taint
	// from the parameter `fromReader673` to result `intoReader671`
	// (`intoReader671` is now tainted).
	intoReader671 := io.TeeReader(fromReader673, nil)

	// Sink the tainted `intoReader671`:
	sink(intoReader671)
}

func TaintStepTest_IoWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString389` into `intoWriter739`.

	// Assume that `sourceCQL` has the underlying type of `fromString389`:
	fromString389 := sourceCQL.(string)

	// Declare `intoWriter739` variable:
	var intoWriter739 io.Writer

	// Call the function that transfers the taint
	// from the parameter `fromString389` to parameter `intoWriter739`;
	// `intoWriter739` is now tainted.
	io.WriteString(intoWriter739, fromString389)

	// Sink the tainted `intoWriter739`:
	sink(intoWriter739)
}

func TaintStepTest_IoLimitedReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromLimitedReader586` into `intoByte143`.

	// Assume that `sourceCQL` has the underlying type of `fromLimitedReader586`:
	fromLimitedReader586 := sourceCQL.(io.LimitedReader)

	// Declare `intoByte143` variable:
	var intoByte143 []byte

	// Call the method that transfers the taint
	// from the receiver `fromLimitedReader586` to the argument `intoByte143`
	// (`intoByte143` is now tainted).
	fromLimitedReader586.Read(intoByte143)

	// Sink the tainted `intoByte143`:
	sink(intoByte143)
}

func TaintStepTest_IoPipeReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPipeReader334` into `intoByte858`.

	// Assume that `sourceCQL` has the underlying type of `fromPipeReader334`:
	fromPipeReader334 := sourceCQL.(io.PipeReader)

	// Declare `intoByte858` variable:
	var intoByte858 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPipeReader334` to the argument `intoByte858`
	// (`intoByte858` is now tainted).
	fromPipeReader334.Read(intoByte858)

	// Sink the tainted `intoByte858`:
	sink(intoByte858)
}

func TaintStepTest_IoPipeWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte924` into `intoPipeWriter429`.

	// Assume that `sourceCQL` has the underlying type of `fromByte924`:
	fromByte924 := sourceCQL.([]byte)

	// Declare `intoPipeWriter429` variable:
	var intoPipeWriter429 io.PipeWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte924` to the receiver `intoPipeWriter429`
	// (`intoPipeWriter429` is now tainted).
	intoPipeWriter429.Write(fromByte924)

	// Sink the tainted `intoPipeWriter429`:
	sink(intoPipeWriter429)
}

func TaintStepTest_IoSectionReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromSectionReader257` into `intoByte117`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader257`:
	fromSectionReader257 := sourceCQL.(io.SectionReader)

	// Declare `intoByte117` variable:
	var intoByte117 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader257` to the argument `intoByte117`
	// (`intoByte117` is now tainted).
	fromSectionReader257.Read(intoByte117)

	// Sink the tainted `intoByte117`:
	sink(intoByte117)
}

func TaintStepTest_IoSectionReaderReadAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromSectionReader513` into `intoByte614`.

	// Assume that `sourceCQL` has the underlying type of `fromSectionReader513`:
	fromSectionReader513 := sourceCQL.(io.SectionReader)

	// Declare `intoByte614` variable:
	var intoByte614 []byte

	// Call the method that transfers the taint
	// from the receiver `fromSectionReader513` to the argument `intoByte614`
	// (`intoByte614` is now tainted).
	fromSectionReader513.ReadAt(intoByte614, 0)

	// Sink the tainted `intoByte614`:
	sink(intoByte614)
}

func TaintStepTest_IoReaderRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader533` into `intoByte477`.

	// Assume that `sourceCQL` has the underlying type of `fromReader533`:
	fromReader533 := sourceCQL.(io.Reader)

	// Declare `intoByte477` variable:
	var intoByte477 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReader533` to the argument `intoByte477`
	// (`intoByte477` is now tainted).
	fromReader533.Read(intoByte477)

	// Sink the tainted `intoByte477`:
	sink(intoByte477)
}

func TaintStepTest_IoReaderAtReadAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReaderAt683` into `intoByte434`.

	// Assume that `sourceCQL` has the underlying type of `fromReaderAt683`:
	fromReaderAt683 := sourceCQL.(io.ReaderAt)

	// Declare `intoByte434` variable:
	var intoByte434 []byte

	// Call the method that transfers the taint
	// from the receiver `fromReaderAt683` to the argument `intoByte434`
	// (`intoByte434` is now tainted).
	fromReaderAt683.ReadAt(intoByte434, 0)

	// Sink the tainted `intoByte434`:
	sink(intoByte434)
}

func TaintStepTest_IoByteReaderReadByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByteReader605` into `intoByte658`.

	// Assume that `sourceCQL` has the underlying type of `fromByteReader605`:
	fromByteReader605 := sourceCQL.(io.ByteReader)

	// Call the method that transfers the taint
	// from the receiver `fromByteReader605` to the result `intoByte658`
	// (`intoByte658` is now tainted).
	intoByte658, _ := fromByteReader605.ReadByte()

	// Sink the tainted `intoByte658`:
	sink(intoByte658)
}

func TaintStepTest_IoReaderFromReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader987` into `intoReaderFrom575`.

	// Assume that `sourceCQL` has the underlying type of `fromReader987`:
	fromReader987 := sourceCQL.(io.Reader)

	// Declare `intoReaderFrom575` variable:
	var intoReaderFrom575 io.ReaderFrom

	// Call the method that transfers the taint
	// from the parameter `fromReader987` to the receiver `intoReaderFrom575`
	// (`intoReaderFrom575` is now tainted).
	intoReaderFrom575.ReadFrom(fromReader987)

	// Sink the tainted `intoReaderFrom575`:
	sink(intoReaderFrom575)
}

func TaintStepTest_IoRuneReaderReadRune_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromRuneReader852` into `intoRune987`.

	// Assume that `sourceCQL` has the underlying type of `fromRuneReader852`:
	fromRuneReader852 := sourceCQL.(io.RuneReader)

	// Call the method that transfers the taint
	// from the receiver `fromRuneReader852` to the result `intoRune987`
	// (`intoRune987` is now tainted).
	intoRune987, _, _ := fromRuneReader852.ReadRune()

	// Sink the tainted `intoRune987`:
	sink(intoRune987)
}

func TaintStepTest_IoWriterWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte257` into `intoWriter668`.

	// Assume that `sourceCQL` has the underlying type of `fromByte257`:
	fromByte257 := sourceCQL.([]byte)

	// Declare `intoWriter668` variable:
	var intoWriter668 io.Writer

	// Call the method that transfers the taint
	// from the parameter `fromByte257` to the receiver `intoWriter668`
	// (`intoWriter668` is now tainted).
	intoWriter668.Write(fromByte257)

	// Sink the tainted `intoWriter668`:
	sink(intoWriter668)
}

func TaintStepTest_IoWriterAtWriteAt_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte595` into `intoWriterAt847`.

	// Assume that `sourceCQL` has the underlying type of `fromByte595`:
	fromByte595 := sourceCQL.([]byte)

	// Declare `intoWriterAt847` variable:
	var intoWriterAt847 io.WriterAt

	// Call the method that transfers the taint
	// from the parameter `fromByte595` to the receiver `intoWriterAt847`
	// (`intoWriterAt847` is now tainted).
	intoWriterAt847.WriteAt(fromByte595, 0)

	// Sink the tainted `intoWriterAt847`:
	sink(intoWriterAt847)
}

func TaintStepTest_IoByteWriterWriteByte_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte259` into `intoByteWriter415`.

	// Assume that `sourceCQL` has the underlying type of `fromByte259`:
	fromByte259 := sourceCQL.(byte)

	// Declare `intoByteWriter415` variable:
	var intoByteWriter415 io.ByteWriter

	// Call the method that transfers the taint
	// from the parameter `fromByte259` to the receiver `intoByteWriter415`
	// (`intoByteWriter415` is now tainted).
	intoByteWriter415.WriteByte(fromByte259)

	// Sink the tainted `intoByteWriter415`:
	sink(intoByteWriter415)
}

func TaintStepTest_IoStringWriterWriteString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString355` into `intoStringWriter588`.

	// Assume that `sourceCQL` has the underlying type of `fromString355`:
	fromString355 := sourceCQL.(string)

	// Declare `intoStringWriter588` variable:
	var intoStringWriter588 io.StringWriter

	// Call the method that transfers the taint
	// from the parameter `fromString355` to the receiver `intoStringWriter588`
	// (`intoStringWriter588` is now tainted).
	intoStringWriter588.WriteString(fromString355)

	// Sink the tainted `intoStringWriter588`:
	sink(intoStringWriter588)
}

func TaintStepTest_IoWriterToWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromWriterTo875` into `intoWriter811`.

	// Assume that `sourceCQL` has the underlying type of `fromWriterTo875`:
	fromWriterTo875 := sourceCQL.(io.WriterTo)

	// Declare `intoWriter811` variable:
	var intoWriter811 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromWriterTo875` to the argument `intoWriter811`
	// (`intoWriter811` is now tainted).
	fromWriterTo875.WriteTo(intoWriter811)

	// Sink the tainted `intoWriter811`:
	sink(intoWriter811)
}

func RunAllTaints_Io() {
	{
		source := newSource()
		TaintStepTest_IoCopy_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoCopyBuffer_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoCopyN_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoLimitReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoMultiReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoMultiWriter_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoNewSectionReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoPipe_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReadAtLeast_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReadFull_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoTeeReader_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoTeeReader_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriteString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoLimitedReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoPipeReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoPipeWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoSectionReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoSectionReaderReadAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReaderRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReaderAtReadAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoByteReaderReadByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoReaderFromReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoRuneReaderReadRune_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterAtWriteAt_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoByteWriterWriteByte_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoStringWriterWriteString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_IoWriterToWriteTo_B0I0O0(source)
	}
}
