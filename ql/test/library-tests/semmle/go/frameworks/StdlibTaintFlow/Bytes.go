package main

import (
	"bytes"
	"io"
)

func TaintStepTest_BytesFields(sourceCQL interface{}) {
	// The flow is from `s` into `into316`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into316`
	// (`into316` is now tainted).
	into316 := bytes.Fields(s)

	// Sink the tainted `into316`:
	sink(into316)
}

func TaintStepTest_BytesFieldsFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into656`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into656`
	// (`into656` is now tainted).
	into656 := bytes.FieldsFunc(s, nil)

	// Sink the tainted `into656`:
	sink(into656)
}

func TaintStepTest_BytesJoin(sourceCQL interface{}) {
	// The flow is from `s` into `into440`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([][]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into440`
	// (`into440` is now tainted).
	into440 := bytes.Join(s, nil)

	// Sink the tainted `into440`:
	sink(into440)
}

func TaintStepTest_BytesMap(sourceCQL interface{}) {
	// The flow is from `s` into `into960`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into960`
	// (`into960` is now tainted).
	into960 := bytes.Map(nil, s)

	// Sink the tainted `into960`:
	sink(into960)
}

func TaintStepTest_BytesNewBuffer(sourceCQL interface{}) {
	// The flow is from `buf` into `into511`.

	// Assume that `sourceCQL` has the underlying type of `buf`:
	buf := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `buf` to result `into511`
	// (`into511` is now tainted).
	into511 := bytes.NewBuffer(buf)

	// Sink the tainted `into511`:
	sink(into511)
}

func TaintStepTest_BytesNewBufferString(sourceCQL interface{}) {
	// The flow is from `s` into `into265`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into265`
	// (`into265` is now tainted).
	into265 := bytes.NewBufferString(s)

	// Sink the tainted `into265`:
	sink(into265)
}

func TaintStepTest_BytesNewReader(sourceCQL interface{}) {
	// The flow is from `b` into `into289`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `b` to result `into289`
	// (`into289` is now tainted).
	into289 := bytes.NewReader(b)

	// Sink the tainted `into289`:
	sink(into289)
}

func TaintStepTest_BytesRepeat(sourceCQL interface{}) {
	// The flow is from `b` into `into258`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `b` to result `into258`
	// (`into258` is now tainted).
	into258 := bytes.Repeat(b, 0)

	// Sink the tainted `into258`:
	sink(into258)
}

func TaintStepTest_BytesReplace(sourceCQL interface{}) {
	// The flow is from `s` into `into706`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into706`
	// (`into706` is now tainted).
	into706 := bytes.Replace(s, nil, nil, 0)

	// Sink the tainted `into706`:
	sink(into706)
}

func TaintStepTest_BytesReplaceAll(sourceCQL interface{}) {
	// The flow is from `s` into `into549`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into549`
	// (`into549` is now tainted).
	into549 := bytes.ReplaceAll(s, nil, nil)

	// Sink the tainted `into549`:
	sink(into549)
}

func TaintStepTest_BytesRunes(sourceCQL interface{}) {
	// The flow is from `s` into `into170`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into170`
	// (`into170` is now tainted).
	into170 := bytes.Runes(s)

	// Sink the tainted `into170`:
	sink(into170)
}

func TaintStepTest_BytesSplit(sourceCQL interface{}) {
	// The flow is from `s` into `into954`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into954`
	// (`into954` is now tainted).
	into954 := bytes.Split(s, nil)

	// Sink the tainted `into954`:
	sink(into954)
}

func TaintStepTest_BytesSplitAfter(sourceCQL interface{}) {
	// The flow is from `s` into `into771`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into771`
	// (`into771` is now tainted).
	into771 := bytes.SplitAfter(s, nil)

	// Sink the tainted `into771`:
	sink(into771)
}

func TaintStepTest_BytesSplitAfterN(sourceCQL interface{}) {
	// The flow is from `s` into `into239`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into239`
	// (`into239` is now tainted).
	into239 := bytes.SplitAfterN(s, nil, 0)

	// Sink the tainted `into239`:
	sink(into239)
}

func TaintStepTest_BytesSplitN(sourceCQL interface{}) {
	// The flow is from `s` into `into407`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into407`
	// (`into407` is now tainted).
	into407 := bytes.SplitN(s, nil, 0)

	// Sink the tainted `into407`:
	sink(into407)
}

func TaintStepTest_BytesTitle(sourceCQL interface{}) {
	// The flow is from `s` into `into308`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into308`
	// (`into308` is now tainted).
	into308 := bytes.Title(s)

	// Sink the tainted `into308`:
	sink(into308)
}

func TaintStepTest_BytesToLower(sourceCQL interface{}) {
	// The flow is from `s` into `into624`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into624`
	// (`into624` is now tainted).
	into624 := bytes.ToLower(s)

	// Sink the tainted `into624`:
	sink(into624)
}

func TaintStepTest_BytesToLowerSpecial(sourceCQL interface{}) {
	// The flow is from `s` into `into768`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into768`
	// (`into768` is now tainted).
	into768 := bytes.ToLowerSpecial(nil, s)

	// Sink the tainted `into768`:
	sink(into768)
}

func TaintStepTest_BytesToTitle(sourceCQL interface{}) {
	// The flow is from `s` into `into479`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into479`
	// (`into479` is now tainted).
	into479 := bytes.ToTitle(s)

	// Sink the tainted `into479`:
	sink(into479)
}

func TaintStepTest_BytesToTitleSpecial(sourceCQL interface{}) {
	// The flow is from `s` into `into264`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into264`
	// (`into264` is now tainted).
	into264 := bytes.ToTitleSpecial(nil, s)

	// Sink the tainted `into264`:
	sink(into264)
}

func TaintStepTest_BytesToUpper(sourceCQL interface{}) {
	// The flow is from `s` into `into360`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into360`
	// (`into360` is now tainted).
	into360 := bytes.ToUpper(s)

	// Sink the tainted `into360`:
	sink(into360)
}

func TaintStepTest_BytesToUpperSpecial(sourceCQL interface{}) {
	// The flow is from `s` into `into708`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into708`
	// (`into708` is now tainted).
	into708 := bytes.ToUpperSpecial(nil, s)

	// Sink the tainted `into708`:
	sink(into708)
}

func TaintStepTest_BytesToValidUTF8(sourceCQL interface{}) {
	// The flow is from `s` into `into844`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into844`
	// (`into844` is now tainted).
	into844 := bytes.ToValidUTF8(s, nil)

	// Sink the tainted `into844`:
	sink(into844)
}

func TaintStepTest_BytesTrim(sourceCQL interface{}) {
	// The flow is from `s` into `into325`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into325`
	// (`into325` is now tainted).
	into325 := bytes.Trim(s, "")

	// Sink the tainted `into325`:
	sink(into325)
}

func TaintStepTest_BytesTrimFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into720`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into720`
	// (`into720` is now tainted).
	into720 := bytes.TrimFunc(s, nil)

	// Sink the tainted `into720`:
	sink(into720)
}

func TaintStepTest_BytesTrimLeft(sourceCQL interface{}) {
	// The flow is from `s` into `into949`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into949`
	// (`into949` is now tainted).
	into949 := bytes.TrimLeft(s, "")

	// Sink the tainted `into949`:
	sink(into949)
}

func TaintStepTest_BytesTrimLeftFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into682`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into682`
	// (`into682` is now tainted).
	into682 := bytes.TrimLeftFunc(s, nil)

	// Sink the tainted `into682`:
	sink(into682)
}

func TaintStepTest_BytesTrimPrefix(sourceCQL interface{}) {
	// The flow is from `s` into `into604`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into604`
	// (`into604` is now tainted).
	into604 := bytes.TrimPrefix(s, nil)

	// Sink the tainted `into604`:
	sink(into604)
}

func TaintStepTest_BytesTrimRight(sourceCQL interface{}) {
	// The flow is from `s` into `into144`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into144`
	// (`into144` is now tainted).
	into144 := bytes.TrimRight(s, "")

	// Sink the tainted `into144`:
	sink(into144)
}

func TaintStepTest_BytesTrimRightFunc(sourceCQL interface{}) {
	// The flow is from `s` into `into654`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into654`
	// (`into654` is now tainted).
	into654 := bytes.TrimRightFunc(s, nil)

	// Sink the tainted `into654`:
	sink(into654)
}

func TaintStepTest_BytesTrimSpace(sourceCQL interface{}) {
	// The flow is from `s` into `into689`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into689`
	// (`into689` is now tainted).
	into689 := bytes.TrimSpace(s)

	// Sink the tainted `into689`:
	sink(into689)
}

func TaintStepTest_BytesTrimSuffix(sourceCQL interface{}) {
	// The flow is from `s` into `into224`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.([]byte)

	// Call the function that transfers the taint
	// from the parameter `s` to result `into224`
	// (`into224` is now tainted).
	into224 := bytes.TrimSuffix(s, nil)

	// Sink the tainted `into224`:
	sink(into224)
}

func TaintStepTest_BytesBufferBytes(sourceCQL interface{}) {
	// The flow is from `from504` into `into134`.

	// Assume that `sourceCQL` has the underlying type of `from504`:
	from504 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from504` to the result `into134`
	// (`into134` is now tainted).
	into134 := from504.Bytes()

	// Sink the tainted `into134`:
	sink(into134)
}

func TaintStepTest_BytesBufferNext(sourceCQL interface{}) {
	// The flow is from `from786` into `into834`.

	// Assume that `sourceCQL` has the underlying type of `from786`:
	from786 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from786` to the result `into834`
	// (`into834` is now tainted).
	into834 := from786.Next(0)

	// Sink the tainted `into834`:
	sink(into834)
}

func TaintStepTest_BytesBufferRead(sourceCQL interface{}) {
	// The flow is from `from387` into `p`.

	// Assume that `sourceCQL` has the underlying type of `from387`:
	from387 := sourceCQL.(bytes.Buffer)

	// Declare `p` variable:
	var p []byte

	// Call the method that transfers the taint
	// from the receiver `from387` to the argument `p`
	// (`p` is now tainted).
	from387.Read(p)

	// Sink the tainted `p`:
	sink(p)
}

func TaintStepTest_BytesBufferReadByte(sourceCQL interface{}) {
	// The flow is from `from354` into `into991`.

	// Assume that `sourceCQL` has the underlying type of `from354`:
	from354 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from354` to the result `into991`
	// (`into991` is now tainted).
	into991, _ := from354.ReadByte()

	// Sink the tainted `into991`:
	sink(into991)
}

func TaintStepTest_BytesBufferReadBytes(sourceCQL interface{}) {
	// The flow is from `from231` into `line`.

	// Assume that `sourceCQL` has the underlying type of `from231`:
	from231 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from231` to the result `line`
	// (`line` is now tainted).
	line, _ := from231.ReadBytes(0)

	// Sink the tainted `line`:
	sink(line)
}

func TaintStepTest_BytesBufferReadFrom(sourceCQL interface{}) {
	// The flow is from `r` into `into571`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(io.Reader)

	// Declare `into571` variable:
	var into571 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into571`
	// (`into571` is now tainted).
	into571.ReadFrom(r)

	// Sink the tainted `into571`:
	sink(into571)
}

func TaintStepTest_BytesBufferReadRune(sourceCQL interface{}) {
	// The flow is from `from776` into `r`.

	// Assume that `sourceCQL` has the underlying type of `from776`:
	from776 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from776` to the result `r`
	// (`r` is now tainted).
	r, _, _ := from776.ReadRune()

	// Sink the tainted `r`:
	sink(r)
}

func TaintStepTest_BytesBufferReadString(sourceCQL interface{}) {
	// The flow is from `from619` into `line`.

	// Assume that `sourceCQL` has the underlying type of `from619`:
	from619 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from619` to the result `line`
	// (`line` is now tainted).
	line, _ := from619.ReadString(0)

	// Sink the tainted `line`:
	sink(line)
}

func TaintStepTest_BytesBufferString(sourceCQL interface{}) {
	// The flow is from `from163` into `into179`.

	// Assume that `sourceCQL` has the underlying type of `from163`:
	from163 := sourceCQL.(bytes.Buffer)

	// Call the method that transfers the taint
	// from the receiver `from163` to the result `into179`
	// (`into179` is now tainted).
	into179 := from163.String()

	// Sink the tainted `into179`:
	sink(into179)
}

func TaintStepTest_BytesBufferWrite(sourceCQL interface{}) {
	// The flow is from `p` into `into853`.

	// Assume that `sourceCQL` has the underlying type of `p`:
	p := sourceCQL.([]byte)

	// Declare `into853` variable:
	var into853 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `p` to the receiver `into853`
	// (`into853` is now tainted).
	into853.Write(p)

	// Sink the tainted `into853`:
	sink(into853)
}

func TaintStepTest_BytesBufferWriteByte(sourceCQL interface{}) {
	// The flow is from `c` into `into356`.

	// Assume that `sourceCQL` has the underlying type of `c`:
	c := sourceCQL.(byte)

	// Declare `into356` variable:
	var into356 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `c` to the receiver `into356`
	// (`into356` is now tainted).
	into356.WriteByte(c)

	// Sink the tainted `into356`:
	sink(into356)
}

func TaintStepTest_BytesBufferWriteRune(sourceCQL interface{}) {
	// The flow is from `r` into `into411`.

	// Assume that `sourceCQL` has the underlying type of `r`:
	r := sourceCQL.(rune)

	// Declare `into411` variable:
	var into411 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `r` to the receiver `into411`
	// (`into411` is now tainted).
	into411.WriteRune(r)

	// Sink the tainted `into411`:
	sink(into411)
}

func TaintStepTest_BytesBufferWriteString(sourceCQL interface{}) {
	// The flow is from `s` into `into890`.

	// Assume that `sourceCQL` has the underlying type of `s`:
	s := sourceCQL.(string)

	// Declare `into890` variable:
	var into890 bytes.Buffer

	// Call the method that transfers the taint
	// from the parameter `s` to the receiver `into890`
	// (`into890` is now tainted).
	into890.WriteString(s)

	// Sink the tainted `into890`:
	sink(into890)
}

func TaintStepTest_BytesBufferWriteTo(sourceCQL interface{}) {
	// The flow is from `from941` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from941`:
	from941 := sourceCQL.(bytes.Buffer)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from941` to the argument `w`
	// (`w` is now tainted).
	from941.WriteTo(w)

	// Sink the tainted `w`:
	sink(w)
}

func TaintStepTest_BytesReaderRead(sourceCQL interface{}) {
	// The flow is from `from622` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from622`:
	from622 := sourceCQL.(bytes.Reader)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from622` to the argument `b`
	// (`b` is now tainted).
	from622.Read(b)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_BytesReaderReadAt(sourceCQL interface{}) {
	// The flow is from `from141` into `b`.

	// Assume that `sourceCQL` has the underlying type of `from141`:
	from141 := sourceCQL.(bytes.Reader)

	// Declare `b` variable:
	var b []byte

	// Call the method that transfers the taint
	// from the receiver `from141` to the argument `b`
	// (`b` is now tainted).
	from141.ReadAt(b, 0)

	// Sink the tainted `b`:
	sink(b)
}

func TaintStepTest_BytesReaderReadByte(sourceCQL interface{}) {
	// The flow is from `from455` into `into490`.

	// Assume that `sourceCQL` has the underlying type of `from455`:
	from455 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `from455` to the result `into490`
	// (`into490` is now tainted).
	into490, _ := from455.ReadByte()

	// Sink the tainted `into490`:
	sink(into490)
}

func TaintStepTest_BytesReaderReadRune(sourceCQL interface{}) {
	// The flow is from `from307` into `ch`.

	// Assume that `sourceCQL` has the underlying type of `from307`:
	from307 := sourceCQL.(bytes.Reader)

	// Call the method that transfers the taint
	// from the receiver `from307` to the result `ch`
	// (`ch` is now tainted).
	ch, _, _ := from307.ReadRune()

	// Sink the tainted `ch`:
	sink(ch)
}

func TaintStepTest_BytesReaderReset(sourceCQL interface{}) {
	// The flow is from `b` into `into763`.

	// Assume that `sourceCQL` has the underlying type of `b`:
	b := sourceCQL.([]byte)

	// Declare `into763` variable:
	var into763 bytes.Reader

	// Call the method that transfers the taint
	// from the parameter `b` to the receiver `into763`
	// (`into763` is now tainted).
	into763.Reset(b)

	// Sink the tainted `into763`:
	sink(into763)
}

func TaintStepTest_BytesReaderWriteTo(sourceCQL interface{}) {
	// The flow is from `from273` into `w`.

	// Assume that `sourceCQL` has the underlying type of `from273`:
	from273 := sourceCQL.(bytes.Reader)

	// Declare `w` variable:
	var w io.Writer

	// Call the method that transfers the taint
	// from the receiver `from273` to the argument `w`
	// (`w` is now tainted).
	from273.WriteTo(w)

	// Sink the tainted `w`:
	sink(w)
}

func RunAllTaints_Bytes(v interface{}) {
	{
		source := newSource()
		TaintStepTest_BytesFields(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesFieldsFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesJoin(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesMap(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesNewBuffer(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesNewBufferString(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesNewReader(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesRepeat(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReplace(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReplaceAll(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesRunes(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplit(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplitAfter(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplitAfterN(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesSplitN(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTitle(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToLower(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToLowerSpecial(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToTitle(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToTitleSpecial(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToUpper(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToUpperSpecial(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesToValidUTF8(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrim(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimLeft(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimLeftFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimPrefix(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimRight(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimRightFunc(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimSpace(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesTrimSuffix(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferNext(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferRead(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadByte(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadBytes(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadFrom(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadRune(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferReadString(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferString(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWrite(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteByte(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteRune(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteString(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesBufferWriteTo(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderRead(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReadAt(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReadByte(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReadRune(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderReset(source)
	}
	{
		source := newSource()
		TaintStepTest_BytesReaderWriteTo(source)
	}
}
