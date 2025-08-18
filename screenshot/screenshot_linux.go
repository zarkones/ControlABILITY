//go:build linux

package screenshot

import (
	"bytes"
	"image/png"

	screen "github.com/kbinani/screenshot"
)

func Capture() ([]byte, error) {
	displayCount := screen.NumActiveDisplays()

	height, width := 0, 0
	for i := 0; i < displayCount; i++ {
		rect := screen.GetDisplayBounds(i)
		if rect.Dy() > height {
			height = rect.Dy()
		}
		width += rect.Dx()
	}

	image, err := screen.Capture(0, 0, width, height)
	if err != nil {
		return nil, err
	}

	var imageBuffer bytes.Buffer

	if err := png.Encode(&imageBuffer, image); err != nil {
		return nil, err
	}

	return imageBuffer.Bytes(), nil
}
