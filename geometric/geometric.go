// Package geometric provides functions to work
// with [geometric progression]
//
// [geometric progression]: https://en.wikipedia.org/wiki/Geometric_progression
package geometric

import (
	"fmt"
	"github.com/punzrr/progression/internal"
	// "errors"
)

// GetVal function returns value ofa
// progression depending on progression info
func GetVal(q, first float32, index int) (float32, error) {
	// returns value depending on progression
	if index == 0 {
		return 0, fmt.Errorf("in GetVal: %w", internal.ErrZeroIndex)
	}
	if index < 0 {
		return 0, fmt.Errorf("in GetVal: %w", internal.ErrNegIndex)
	}
	progress, err := Degree(q, index-1)
	if err != nil {
		return 0, fmt.Errorf("in GetVal: %w", err)
	}
	return first * progress, nil

}

func Degree(val float32, degree int) (float32, error) {
	// Provides degree operation
	if degree == 0 {
		return 1, nil
	}
	if degree < 0 {
		return 0, fmt.Errorf("in Degree: %w", internal.ErrNegIndex)
	}
	var result float32 = 1
	for degree != 0 {
		result *= val
		degree -= 1
	}
	return result, nil
}
