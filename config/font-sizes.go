package config

import (
	"errors"
	"fmt"
)

type FontSizes struct {
	Min int
	Max int
}

// Validate validation font size range
func (f *FontSizes) Validate() error {
	if f.Min > f.Max {
		return fmt.Errorf(
			"%w: min > max",
			errors.New("set font size range error"),
		)
	}
	if f.Min < 24 || f.Min > 200 {
		return fmt.Errorf(
			"%w: min must between 24 and 200",
			errors.New("set font size range error"),
		)
	}
	if f.Max < 24 || f.Max > 200 {
		return fmt.Errorf(
			"%w: max must between 24 and 200",
			errors.New("set font size range error"),
		)
	}
	return nil
}
