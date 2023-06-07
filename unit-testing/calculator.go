package unit_testing

type InterfaceCalculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
	Multiply(a, b int) int
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}
