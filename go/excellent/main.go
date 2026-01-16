package main

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	// Example usage!
	result := EvenOrOdd(5)
	println(result) // Output: odd
}
