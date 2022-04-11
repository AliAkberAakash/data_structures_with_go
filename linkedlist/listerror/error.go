package listerror

import "strconv"

type OutOfBoundsError struct {
	Position int
}

func (err *OutOfBoundsError) Error() string {
	return "Invalid position" + strconv.Itoa(err.Position) + ", Out of bounds"
}
