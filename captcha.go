// Package captcha creating unique images with text for user verification
package captcha

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
)

// Captcha image generator
type Captcha struct {
	config  *Config
	fonts   []*Font
	Value   string
	letters []*Letter
	Image   *image.RGBA
}

// New constructor of image generator
func New(config *Config) (*Captcha, error) {
	c := &Captcha{
		config: config,
	}
	var err error
	if err = c.config.Setup(); err != nil {
		return nil, err
	}

	if err = c.setFonts(); err != nil {
		return nil, err
	}
	c.Value = c.generateValue()
	if err = c.setLetters(); err != nil {
		return nil, err
	}

	c.joinLetterImages()

	//fmt.Printf("Captcha: %+v\n", c)
	return c, nil
}

// setFonts setting fonts from config
func (c *Captcha) setFonts() error {
	for _, ff := range c.config.FontFiles {
		fnt, err := NewFont(ff)
		if err != nil {
			return err
		}
		c.fonts = append(c.fonts, fnt)
	}
	return nil
}

// generateValue generate captcha value of letters
func (c *Captcha) generateValue() string {
	vl := rand.Intn(
		c.config.LettersLength.Max+1-c.config.LettersLength.Min) + c.config.LettersLength.Min
	l := []rune(c.config.Letters)
	vs := make([]rune, vl)
	for i := range vs {
		vs[i] = l[rand.Intn(len(l))]
	}
	return string(vs)
}

// setLetters making letters
func (c *Captcha) setLetters() error {
	for _, l := range c.Value {
		ltr, err := NewLetter(l, &ConfigLetter{
			Font: c.fonts[rand.Intn(len(c.fonts))],
			Size: float64(rand.Intn(
				c.config.FontSizes.Max+1-c.config.FontSizes.Min) + c.config.FontSizes.Min),
			Color: c.config.TextColors[rand.Intn(len(c.config.TextColors))],
			Rotate: float64(
				c.config.Rotate.Min + rand.Int()%(c.config.Rotate.Max-c.config.Rotate.Min+1)),
			BGColor: color.White,
		})
		if err != nil {
			return err
		}
		c.letters = append(c.letters, ltr)
	}
	return nil
}

// joinLetterImages completing captcha image
func (c *Captcha) joinLetterImages() {
	// calculate width and height
	var width int
	var height int

	for i, l := range c.letters {
		if i == len(c.letters)-1 {
			width += l.CompleteImage.Bounds().Max.X
		} else {
			width += int(float64(l.CompleteImage.Bounds().Max.X) * 0.7)
		}
		if l.CompleteImage.Bounds().Max.Y > height {
			height = l.CompleteImage.Bounds().Max.Y
		}
	}

	// create image
	c.Image = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(c.Image, c.Image.Bounds(), image.NewUniform(color.White), image.ZP, draw.Src)

	// add letter images
	var x int
	for _, l := range c.letters {
		rct := image.Rectangle{
			Min: image.Point{
				X: x,
				Y: 0,
			},
			Max: image.Point{
				X: x + l.CompleteImage.Bounds().Max.X,
				Y: l.CompleteImage.Bounds().Max.Y,
			},
		}
		x += int(float64(l.CompleteImage.Bounds().Max.X) * 0.7)
		draw.Draw(c.Image, rct, l.CompleteImage, image.Point{}, draw.Over)
	}
}

// SaveToFile save image to file.
func (c *Captcha) SaveToFile(filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf(
			"error captcha file opening: %w",
			err,
		)
	}
	b := bufio.NewWriter(outFile)
	if err = png.Encode(b, c.Image); err != nil {
		return fmt.Errorf(
			"error captcha writing: %w",
			err,
		)
	}
	if err = b.Flush(); err != nil {
		return fmt.Errorf(
			"error captcha flush: %w",
			err,
		)
	}
	if err = outFile.Close(); err != nil {
		return fmt.Errorf(
			"error captcha file closing: %w",
			err,
		)
	}
	return nil
}
