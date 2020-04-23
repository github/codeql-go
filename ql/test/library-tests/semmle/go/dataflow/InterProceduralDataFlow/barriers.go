package main

func validate(x interface{}) bool {
	return something(x) && isSafe(x)
}

func funcBarrier() {
	x := source()

	if validate(x) {
		sink(x) // OK
	}
}

func validate1(x interface{}) bool {
	return something(x) || isSafe(x)
}

func funcBarrier1() {
	x := source()

	if validate1(x) {
		sink(x) // NOT OK
	}
}

func phi() {
	x := source()

	if something(x) && isSafe(x) {
		// this input to the phi node for 'x' should be sanitized
	} else {
		x = nil
	}
	sink(x) // OK
}

func phi2() {
	x := source()

	if something(x) || isSafe(x) {
		// this input to the phi node for 'x' is not fully sanitized
	} else {
		x = nil
	}
	sink(x) // NOT OK
}
