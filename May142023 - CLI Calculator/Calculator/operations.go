package Calculator

type Calculator struct {
	Num1      int
	Num2      int
	Operation string
}

func Add(c Calculator) int {
	return c.Num1 + c.Num2
}

func Subtract(c Calculator) int {
	return c.Num1 - c.Num2
}

func Multiply(c Calculator) int {
	return c.Num1 * c.Num2
}

func Divide(c Calculator) int {
	if c.Num2 == 0 {
		return 0 // Handle division by zero
	}
	return c.Num1 / c.Num2
}
