package captcha

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"os"
)

// Font шрифт
type Font struct {
	file string
	Font *truetype.Font
}

// NewFont конструктор шрифта
func NewFont(file string) (*Font, error) {
	f := &Font{
		file: file,
	}
	if err := f.loadFontFromFile(); err != nil {
		return nil, err
	}
	return f, nil
}

// loadFontFromFile загрузка шрифта из файла
func (f *Font) loadFontFromFile() error {
	fb, err := os.ReadFile(f.file)
	if err != nil {
		return fmt.Errorf("load font file error: %w", err)
	}
	ft, err := freetype.ParseFont(fb)
	if err != nil {
		return fmt.Errorf("parse font error: %w", err)
	}
	f.Font = ft
	return nil
}
