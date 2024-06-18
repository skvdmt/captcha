package main

import (
	"fmt"
	"github.com/skvdmt/captcha"
	"github.com/skvdmt/captcha/config"
	"log"
)

const (
	fileFontOpenSansBoldItalic   = "./fonts/OpenSans-BoldItalic.ttf"
	fileFontOswaldBold           = "./fonts/Oswald-Bold.ttf"
	fileFontGoBold               = "./fonts/Go-Bold.ttf"
	fileFontRobotoMonoBoldItalic = "./fonts/RobotoMono-BoldItalic.ttf"

	captchaFile  = "./captcha.png"
	lettersCount = 4
	minFontSize  = 100
	maxFontSize  = 200
	minRotate    = -20
	maxRotate    = 20
)

func main() {
	capt, err := captcha.New(&captcha.Config{
		FontFiles: []string{
			fileFontGoBold,
			fileFontRobotoMonoBoldItalic,
			fileFontOswaldBold,
			fileFontOpenSansBoldItalic,
		},
		LettersLength: &config.LettersLength{
			Min: lettersCount,
			Max: lettersCount,
		},
		FontSizes: &config.FontSizes{
			Min: minFontSize,
			Max: maxFontSize,
		},
		Rotate: &config.Rotate{
			Min: minRotate,
			Max: maxRotate,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if err = capt.SaveToFile(captchaFile); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value: %s\n", capt.Value)
}
