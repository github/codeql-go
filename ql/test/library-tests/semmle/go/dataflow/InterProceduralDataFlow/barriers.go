package main

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
