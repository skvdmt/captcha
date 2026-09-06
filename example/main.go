package main

import (
	"fmt"
	"log"

	"github.com/skvdmt/captcha"
)

const (
	FONT_FILE_OPEN_SANS_BOLD_ITALIC   = "./fonts/OpenSans-BoldItalic.ttf"
	FONT_FILE_OSWALD_BOLD             = "./fonts/Oswald-Bold.ttf"
	FONT_FILE_GO_BOLD                 = "./fonts/Go-Bold.ttf"
	FONT_FILE_ROBOTO_MONO_BOLD_ITALIC = "./fonts/RobotoMono-BoldItalic.ttf"

	CAPTCHA_FILE  = "./captcha.png"
	LETTERS_COUNT = 4
	MIN_FONT_SIZE = 80
	MAX_FONT_SIZE = 100
	MIN_ROTATE    = -70
	MAX_ROTATE    = 70
)

func main() {
	c, err := captcha.New(
		captcha.WithFontFiles([]string{
			FONT_FILE_OPEN_SANS_BOLD_ITALIC,
			FONT_FILE_OSWALD_BOLD,
			FONT_FILE_GO_BOLD,
			FONT_FILE_ROBOTO_MONO_BOLD_ITALIC,
		}),
		captcha.WithLettersLength(LETTERS_COUNT, LETTERS_COUNT),
		captcha.WithFontSizes(MIN_FONT_SIZE, MAX_FONT_SIZE),
		captcha.WithRotate(MIN_ROTATE, MAX_ROTATE),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = c.SaveToFile(CAPTCHA_FILE); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value: %s\n", c.Value)
}
