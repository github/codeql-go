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
