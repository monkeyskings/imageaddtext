package comment

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
)

func (c *Comment) DecodeImage(fd io.Reader) (image.Image, error) {
	switch c.filetype {
	case "png":
		log.Println("decode png image")
		return png.Decode(fd)
	case "jpg":
		log.Println("decode jpg image")
		return jpeg.Decode(fd)
	case "gif":
		log.Println("decode gif image")
		return gif.Decode(fd)
	default:
		return nil, errors.New("UnkownType Error")
	}

}

func (c *Comment) EncodeImage(fd io.Writer, img image.Image) error {
	switch c.filetype {
	case "png":
		log.Println("encode png image")
		return png.Encode(fd, img)
	case "jpg":
		log.Println("encode jpg image")
		return jpeg.Encode(fd, img, &jpeg.Options{Quality: 100})
	case "gif":
		log.Println("encode gif image")
		//return errors.New("GifType Error")
		return gif.Encode(fd, img, &gif.Options{NumColors: 256})
	default:
		return errors.New("UnkownType Error")
	}
}
