package captcha

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

type ConfigLetter struct {
	Font    *Font
	Size    float64
	BGColor color.Color
	Color   color.Color
	Rotate  float64
}

// Letter символ
type Letter struct {
	Value         rune
	config        *ConfigLetter
	ctx           *freetype.Context
	image         *image.RGBA
	CompleteImage *image.RGBA
	letterWidth   int
	letterHeight  int
}

// NewLetter конструктор символа
func NewLetter(value rune, config *ConfigLetter) (*Letter, error) {
	l := &Letter{
		Value:  value,
		config: config,
	}
	l.initContext()
	if err := l.setImageLetterSize(); err != nil {
		return nil, err
	}
	l.createImage()
	if err := l.drawLetterToImage(); err != nil {
		return nil, err
	}
	l.rotateLetterImage()
	l.chopImageLetter()
	l.changeWhitePixelToAlpha()
	return l, nil
}

// initContext init context
func (l *Letter) initContext() {
	l.ctx = freetype.NewContext()
	l.ctx.SetDPI(72)
	l.ctx.SetFont(l.config.Font.Font)
	l.ctx.SetFontSize(l.config.Size)
	l.ctx.SetSrc(image.NewUniform(l.config.Color))
	l.ctx.SetHinting(font.HintingNone)
}

// setImageLetterSize calculate image size by letter
func (l *Letter) setImageLetterSize() error {
	pnt, err := l.ctx.DrawString(string(l.Value), freetype.Pt(0, int(l.config.Size)))
	if err != nil {
		return fmt.Errorf("error get image letter size: %w", err)
	}
	l.letterWidth = pnt.X.Round()
	l.letterHeight = pnt.Y.Round()
	return nil
}

// createImage create new rectangle image
func (l *Letter) createImage() {
	iw := l.letterWidth * 3
	ih := l.letterHeight * 3
	l.image = image.NewRGBA(image.Rect(0, 0, iw, ih))
	draw.Draw(l.image, l.image.Bounds(), image.NewUniform(l.config.BGColor), image.Point{}, draw.Over)
	l.ctx.SetClip(l.image.Bounds())
	l.ctx.SetDst(l.image)
}

// drawLetterToImage draw letter to image
func (l *Letter) drawLetterToImage() error {
	pt := freetype.Pt(l.letterWidth, l.letterHeight*2)
	_, err := l.ctx.DrawString(string(l.Value), pt)
	if err != nil {
		return fmt.Errorf("error draw letter: %w", err)
	}
	return nil
}

// rotateLetterImage rotate letter image to rotate angle
func (l *Letter) rotateLetterImage() {
	l.image = (*image.RGBA)(imaging.Rotate(
		l.image, l.config.Rotate, image.NewUniform(l.config.BGColor)))
}

// changeWhitePixelToAlpha change white pixel to alpha
func (l *Letter) changeWhitePixelToAlpha() {
	var col uint32 = 50000
	for x := 0; x <= l.CompleteImage.Bounds().Dx(); x++ {
		for y := 0; y <= l.CompleteImage.Bounds().Dy(); y++ {
			r, g, b, a := l.CompleteImage.At(x, y).RGBA()
			if r > col && g > col && b > col && a > col {
				l.CompleteImage.Set(x, y, image.NewUniform(color.Alpha{}))
			}
		}
	}
}

// chopImageLetter chop image letter around space
func (l *Letter) chopImageLetter() {
	minX, minY := l.image.Bounds().Dx(), l.image.Bounds().Dy()
	maxX, maxY := 0, 0
	var col uint32 = 65535
	for y := 0; y < l.image.Bounds().Dy(); y++ {
		for x := 0; x < l.image.Bounds().Dx(); x++ {
			r, g, b, a := l.image.At(x, y).RGBA()
			if r < col || g < col || b < col || a < col {
				if minX > x {
					minX = x
				}
				if maxX < x {
					maxX = x
				}
				if minY > y {
					minY = y
				}
				if maxY < y {
					maxY = y
				}
			}
		}
	}
	l.CompleteImage = image.NewRGBA(image.Rect(0, 0, maxX-minX, maxY-minY))
	draw.Draw(
		l.CompleteImage, l.CompleteImage.Bounds(),
		image.NewUniform(color.White), image.Point{}, draw.Src)

	draw.Draw(
		l.CompleteImage,
		image.Rect(0, 0, maxX-minX, maxY-minY),
		l.image,
		image.Point{X: minX, Y: minY},
		draw.Over,
	)
}
