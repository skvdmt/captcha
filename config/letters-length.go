package config

import (
	"errors"
	"fmt"
)

// LettersLength setup letters length configuration
type LettersLength struct {
	Min int
	Max int
}

// Validate validation letters lengths
func (l *LettersLength) Validate() error {
	if l.Min > l.Max {
		return fmt.Errorf(
			"%w: min > max",
			errors.New("set letters range error"),
		)
	}
	if l.Min < 1 || l.Min > 30 {
		return fmt.Errorf(
			"%w: min must between 1 and 30",
			errors.New("set letters range error"),
		)
	}
	if l.Max < 1 || l.Max > 30 {
		return fmt.Errorf(
			"%w: max must between 1 and 30",
			errors.New("set letters range error"),
		)
	}
	return nil
}
