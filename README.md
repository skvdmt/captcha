# Captcha
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
	fileFontOpenSansBoldItalic   = "./fonts/OpenSans-BoldItalic.ttf"
	fileFontOswaldBold           = "./fonts/Oswald-Bold.ttf"
	fileFontGoBold               = "./fonts/Go-Bold.ttf"
	fileFontRobotoMonoBoldItalic = "./fonts/RobotoMono-BoldItalic.ttf"

	captchaFile  = "./captcha.png"
)

func main() {
	capt, err := captcha.New(&captcha.Config{
		FontFiles: []string{
			fileFontGoBold,
			fileFontRobotoMonoBoldItalic,
			fileFontOswaldBold,
			fileFontOpenSansBoldItalic,
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
```