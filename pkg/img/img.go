package img

import (
	"image"
	"image/jpeg"
	"io"
	"math"
	"strings"

	"golang.org/x/image/draw"
)

type Img struct {
	Date    string
	Caption string
	Path    string
}

func (i *Img) PreviewPath() string {
	return strings.ReplaceAll(i.Path, ".jpg", "-resized.jpg")
}

func CreatePreview(r io.Reader, w io.Writer) error {
	var src image.Image
	var err error
	width := 800

	src, err = jpeg.Decode(r)

	if err != nil {
		return err
	}

	ratio := (float64)(src.Bounds().Max.Y) / (float64)(src.Bounds().Max.X)
	height := int(math.Round(float64(width) * ratio))

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	err = jpeg.Encode(w, dst, nil)
	if err != nil {
		return err
	}

	return nil
}
