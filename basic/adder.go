package adder

import "errors"

var (
	ErrInvalidSummand = errors.New("Invalid summand")
	ErrZeroSummand = errors.New("Zero summand")
)

func Add(x, y int) (int, error) {
	// when values are zero
	if x == 0 || y == 0 {
		return 0, ErrZeroSummand
	}

	// when values are negative
	if x < 0 || y < 0 {
		return 0, ErrInvalidSummand
	}
	
	return x + y, nil
}