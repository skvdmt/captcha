package config

import (
	"errors"
	"fmt"
)

type Rotate struct {
	Min int
	Max int
}

// Validate validation rotate angle letters
func (r *Rotate) Validate() error {
	if r.Min > r.Max {
		return fmt.Errorf(
			"%w: min > max",
			errors.New("set rotate error"),
		)
	}
	if r.Min < -70 || r.Min > 70 {
		return fmt.Errorf(
			"%w: min must between -70 and 70",
			errors.New("set rotate error"),
		)
	}
	if r.Max < -70 || r.Max > 70 {
		return fmt.Errorf(
			"%w: max must between -70 and 70",
			errors.New("set rotate error"),
		)
	}
	return nil
}
