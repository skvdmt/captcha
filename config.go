package captcha

import (
	"errors"
	"fmt"
	"github.com/skvdmt/captcha/config"
	"image/color"
	"math/rand"
)

const (
	defaultLetters     = "abcdefghkmnpqrstuvwxyzABCDEFGHKMNPQRSTUVWXYZ23456789"
	defaultMinLetters  = 4
	defaultMaxLetters  = 6
	defaultMinFontSize = 48
	defaultMaxFontSize = 64
	defaultMinRotate   = -70
	defaultMaxRotate   = 70
)

var (
	defaultColorBlack  = color.RGBA{R: 0, G: 0, B: 0, A: uint8(rand.Intn(100)) + 100}
	defaultColorRed    = color.RGBA{R: 100, G: 0, B: 0, A: uint8(rand.Intn(100)) + 100}
	defaultColorGreen  = color.RGBA{R: 0, G: 100, B: 0, A: uint8(rand.Intn(100)) + 100}
	defaultColorBlue   = color.RGBA{R: 0, G: 0, B: 100, A: uint8(rand.Intn(100)) + 100}
	defaultColorCyan   = color.RGBA{R: 0, G: 100, B: 100, A: uint8(rand.Intn(100)) + 100}
	defaultColorPurple = color.RGBA{R: 100, G: 0, B: 100, A: uint8(rand.Intn(100)) + 100}
	defaultColorYellow = color.RGBA{R: 100, G: 100, B: 0, A: uint8(rand.Intn(100)) + 155}
)

// Config configuration image generator
type Config struct {
	FontFiles     []string
	Letters       string
	LettersLength *config.LettersLength
	FontSizes     *config.FontSizes
	Rotate        *config.Rotate
	TextColors    []color.Color
}

// Setup settings and validation config
func (c *Config) Setup() error {
	var err error
	c.setConfigLetters()
	if err = c.validateFontFilesList(); err != nil {
		return err
	}
	if err = c.setLettersRange(); err != nil {
		return err
	}
	if err = c.setFontSizeRange(); err != nil {
		return err
	}
	if err = c.setRotate(); err != nil {
		return err
	}
	c.setTextColors()
	return nil
}

// setConfigLetters setting letters from config or default const letters
func (c *Config) setConfigLetters() {
	if len(c.Letters) == 0 {
		c.Letters = defaultLetters
	}
}

// validateFontFilesList validate font files list
func (c *Config) validateFontFilesList() error {
	if len(c.FontFiles) == 0 {
		return fmt.Errorf(
			"%w: no font files",
			errors.New("font files error"),
		)
	}
	return nil
}

// setLettersRange setting range letters
func (c *Config) setLettersRange() error {
	if c.LettersLength == nil {
		c.LettersLength = &config.LettersLength{
			Min: defaultMinLetters,
			Max: defaultMaxLetters,
		}
	}
	return c.LettersLength.Validate()
}

// setFontSizeRange setting font size range
func (c *Config) setFontSizeRange() error {
	if c.FontSizes == nil {
		c.FontSizes = &config.FontSizes{
			Min: defaultMinFontSize,
			Max: defaultMaxFontSize,
		}
	}
	return c.FontSizes.Validate()
}

// setTextColors setting text colors
func (c *Config) setTextColors() {
	if len(c.TextColors) == 0 {
		c.TextColors = []color.Color{
			defaultColorBlack,
			defaultColorRed,
			defaultColorGreen,
			defaultColorBlue,
			defaultColorCyan,
			defaultColorPurple,
			defaultColorYellow,
		}
	}
}

// setRotate setting rotate angle letters
func (c *Config) setRotate() error {
	if c.Rotate == nil {
		c.Rotate = &config.Rotate{
			Min: defaultMinRotate,
			Max: defaultMaxRotate,
		}
	}
	return c.Rotate.Validate()
}
