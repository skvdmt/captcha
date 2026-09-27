package captcha

import (
	"image/color"
)

// Option Functional option.
type Option func(c *Config)

// WithFontFiles Set font files.
func WithFontFiles(fontFiles []string) Option {
	return func(c *Config) {
		c.FontFiles = fontFiles
	}
}

// WithLetters Set symbols.
func WithLetters(letters string) Option {
	return func(c *Config) {
		c.Letters = letters
	}
}

// WithLettersLength Set captcha size.
func WithLettersLength(Min, Max int) Option {
	return func(c *Config) {
		c.LettersLength = &LettersLength{Min: Min, Max: Max}
	}
}

// WithFontSizes Set font sizes.
func WithFontSizes(Min, Max int) Option {
	return func(c *Config) {
		c.FontSizes = &FontSizes{Min: Min, Max: Max}
	}
}

// WithRotate Set rotates.
func WithRotate(Min, Max int) Option {
	return func(c *Config) {
		c.Rotate = &Rotate{Min: Min, Max: Max}
	}
}

// WithTextColors Set text colors.
func WithTextColors(textColors []color.Color) Option {
	return func(c *Config) {
		c.TextColors = textColors
	}
}
