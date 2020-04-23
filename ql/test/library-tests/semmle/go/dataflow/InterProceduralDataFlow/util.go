package main

func source() interface{} {
	return struct{}{}
}

func sink(xs ...interface{}) {}

func something(x interface{}) bool {
	return false
}

func isSafe(x interface{}) bool {
	return true
}
