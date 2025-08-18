//go:build windows

package screenshot

import (
	"bytes"
	"image"
	"image/png"

	screen "github.com/kbinani/screenshot"
)

func Capture() ([]byte, error) {
	displayCount := screen.NumActiveDisplays()

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < displayCount; i++ {
		rect := screen.GetDisplayBounds(i)
		all = rect.Union(all)
	}

	img, err := screen.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		return nil, err
	}

	var imageBuffer bytes.Buffer

	if err := png.Encode(&imageBuffer, img); err != nil {
		return nil, err
	}

	return imageBuffer.Bytes(), nil
}
