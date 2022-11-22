package calculator

import (
	"errors"
	"math"
)

type Calculator struct {
	operations map[string]func(float64, float64) float64
	result     float64
}

func NewCalculator() *Calculator {
	return &Calculator{
		operations: map[string]func(float64, float64) float64{
			"+": add,
			"-": sub,
			"/": div,
			"*": mul,
		},
		result: 0.0,
	}
}

func (c *Calculator) Evaluate(operation string, num float64) error {
	if f, ok := c.operations[operation]; ok {
		c.result = f(c.result, num)

		return nil
	}

	return errors.New("no such operation")
}

func (c *Calculator) GetResult() float64 {
	return c.result
}

func (c *Calculator) Restart() {
	c.result = 0
}

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func sub(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func div(num1 float64, num2 float64) float64 {
	if num2 == 0 {
		switch {
		case num1 > 0:
			return math.Inf(1)
		case num1 < 0:
			return math.Inf(-1)
		default:
			return math.NaN()
		}
	}

	return num1 / num2
}

func mul(num1 float64, num2 float64) float64 {
	return num1 * num2
}
