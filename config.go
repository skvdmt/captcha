package captcha

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"unicode/utf8"
)

const (
	DEFAULT_LETTERS       = "abcdefghkmnpqrstuvwxyzABCDEFGHKMNPQRSTUVWXYZ23456789"
	DEFAULT_LETTERS_MIN   = 4
	DEFAULT_LETTERS_MAX   = 6
	DEFAULT_FONT_SIZE_MIN = 48
	DEFAULT_FONT_SIZE_MAX = 64
	DEFAULT_ROTATE_MIN    = -70
	DEFAULT_ROTATE_MAX    = 70
)

var (
	DEFAULT_COLOR_BLACK  = color.RGBA{R: 0, G: 0, B: 0, A: uint8(rand.Intn(100)) + 100}
	DEFAULT_COLOR_RED    = color.RGBA{R: 100, G: 0, B: 0, A: uint8(rand.Intn(100)) + 100}
	DEFAULT_COLOR_GREEN  = color.RGBA{R: 0, G: 100, B: 0, A: uint8(rand.Intn(100)) + 100}
	DEFAULT_COLOR_BLUE   = color.RGBA{R: 0, G: 0, B: 100, A: uint8(rand.Intn(100)) + 100}
	DEFAULT_COLOR_CYAN   = color.RGBA{R: 0, G: 100, B: 100, A: uint8(rand.Intn(100)) + 100}
	DEFAULT_COLOR_PURPLE = color.RGBA{R: 100, G: 0, B: 100, A: uint8(rand.Intn(100)) + 100}
	DEFAULT_COLOR_YELLOW = color.RGBA{R: 100, G: 100, B: 0, A: uint8(rand.Intn(100)) + 155}
)

// Config configuration image generator
type Config struct {
	FontFiles     []string
	Letters       string
	LettersLength *LettersLength
	FontSizes     *FontSizes
	Rotate        *Rotate
	TextColors    []color.Color
}

// NewConfig Constructor.
func NewConfig() *Config {
	return &Config{
		Letters: DEFAULT_LETTERS,
		LettersLength: &LettersLength{
			Min: DEFAULT_LETTERS_MIN,
			Max: DEFAULT_LETTERS_MAX,
		},
		FontSizes: &FontSizes{
			Min: DEFAULT_FONT_SIZE_MIN,
			Max: DEFAULT_FONT_SIZE_MAX,
		},
		Rotate: &Rotate{
			Min: DEFAULT_ROTATE_MIN,
			Max: DEFAULT_ROTATE_MAX,
		},
		TextColors: []color.Color{
			DEFAULT_COLOR_BLACK,
			DEFAULT_COLOR_RED,
			DEFAULT_COLOR_GREEN,
			DEFAULT_COLOR_BLUE,
			DEFAULT_COLOR_CYAN,
			DEFAULT_COLOR_PURPLE,
			DEFAULT_COLOR_YELLOW,
		},
	}
}

// Validate Config validation.
func (c *Config) Validate() error {
	if err := c.ValidateFontFiles(); err != nil {
		return err
	}
	if err := c.ValidateLetters(); err != nil {
		return err
	}
	if err := c.ValidateTextColors(); err != nil {
		return err
	}
	if err := c.LettersLength.Validate(); err != nil {
		return err
	}
	if err := c.FontSizes.Validate(); err != nil {
		return err
	}
	if err := c.Rotate.Validate(); err != nil {
		return err
	}
	return nil
}

// ValidateFontFiles validate font files
func (c *Config) ValidateFontFiles() error {
	if len(c.FontFiles) == 0 {
		return fmt.Errorf(
			"%w: no font files",
			errors.New("font files error"),
		)
	}
	return nil
}

// ValidateLetters validate letters
func (c *Config) ValidateLetters() error {
	if utf8.RuneCountInString(c.Letters) == 0 {
		return fmt.Errorf(
			"%w: no letters",
			errors.New("letters error"),
		)
	}
	return nil
}

// ValidateTextColors validate text colors
func (c *Config) ValidateTextColors() error {
	if len(c.TextColors) == 0 {
		return fmt.Errorf(
			"%w: no text colors",
			errors.New("text colors error"),
		)
	}
	return nil
}
