# Captcha
![Logo](./captcha.png "Captcha Example")

Creating unique images with text for user verification.

## Installation
```
go get -u github.com/skvdmt/captcha
```

## Example

```go
package main

import (
	"fmt"
	"github.com/skvdmt/captcha"
	"log"
)

const (
	FONT_FILE_OPEN_SANS_BOLD_ITALIC   = "./fonts/OpenSans-BoldItalic.ttf"
	FONT_FILE_OSWALD_BOLD             = "./fonts/Oswald-Bold.ttf"
	FONT_FILE_GO_BOLD                 = "./fonts/Go-Bold.ttf"
	FONT_FILE_ROBOTO_MONO_BOLD_ITALIC = "./fonts/RobotoMono-BoldItalic.ttf"

	CAPTCHA_FILE  = "./captcha.png"
)

func main() {
	c, err := captcha.New(
		captcha.WithFontFiles([]string{
			FONT_FILE_OPEN_SANS_BOLD_ITALIC,
			FONT_FILE_OSWALD_BOLD,
			FONT_FILE_GO_BOLD,
			FONT_FILE_ROBOTO_MONO_BOLD_ITALIC,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = c.SaveToFile(CAPTCHA_FILE); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value: %s\n", c.Value)
}
```