package arithmetic

import (
	"fmt"
	"github.com/punzrr/progression/internal"
)

func GetVal(d, first float32, index int) (float32, error) {
	if index == 0 {
		return 0, fmt.Errorf("in GetVal: %w", internal.ErrZeroIndex)
	}
	if index == 0 {
		return 0, fmt.Errorf("in GetVal: %w", internal.ErrNegIndex)
	}
	return first + float32(d*float32(index-1)), nil
}
