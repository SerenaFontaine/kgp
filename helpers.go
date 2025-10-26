package kgp

import (
	"bytes"
	"compress/zlib"
	"image"
	"image/png"
)

// ImageToRGBA converts an image.Image to raw RGBA bytes.
func ImageToRGBA(img image.Image) []byte {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	rgba := make([]byte, width*height*4)
	idx := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgba[idx+0] = byte(r >> 8)
			rgba[idx+1] = byte(g >> 8)
			rgba[idx+2] = byte(b >> 8)
			rgba[idx+3] = byte(a >> 8)
			idx += 4
		}
	}

	return rgba
}

// ImageToRGB converts an image.Image to raw RGB bytes.
func ImageToRGB(img image.Image) []byte {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	rgb := make([]byte, width*height*3)
	idx := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rgb[idx+0] = byte(r >> 8)
			rgb[idx+1] = byte(g >> 8)
			rgb[idx+2] = byte(b >> 8)
			idx += 3
		}
	}

	return rgb
}

// ImageToPNG converts an image.Image to PNG bytes.
func ImageToPNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CompressZlib compresses data using ZLIB (RFC 1950).
func CompressZlib(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)

	_, err := w.Write(data)
	if err != nil {
		w.Close()
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TransmitImage is a convenience function to transmit and display an image.Image
// using PNG format.
func TransmitImage(img image.Image) (*Command, error) {
	pngData, err := ImageToPNG(img)
	if err != nil {
		return nil, err
	}

	cmd := NewTransmitDisplay().
		Format(FormatPNG).
		TransmitDirect(pngData).
		Build()

	return cmd, nil
}

// TransmitImageWithID is a convenience function to transmit and display an image.Image
// with a specific image ID using PNG format.
func TransmitImageWithID(img image.Image, imageID uint32) (*Command, error) {
	pngData, err := ImageToPNG(img)
	if err != nil {
		return nil, err
	}

	cmd := NewTransmitDisplay().
		ImageID(imageID).
		Format(FormatPNG).
		TransmitDirect(pngData).
		Build()

	return cmd, nil
}

// TransmitImageRGBA is a convenience function to transmit and display an image.Image
// using raw RGBA format with optional compression.
func TransmitImageRGBA(img image.Image, compress bool) (*Command, error) {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	rgba := ImageToRGBA(img)

	builder := NewTransmitDisplay().
		Format(FormatRGBA).
		Dimensions(width, height)

	if compress {
		compressed, err := CompressZlib(rgba)
		if err != nil {
			return nil, err
		}
		builder.Compress().TransmitDirect(compressed)
	} else {
		builder.TransmitDirect(rgba)
	}

	return builder.Build(), nil
}

// CreateRGBAColor creates a 32-bit RGBA color value.
func CreateRGBAColor(r, g, b, a uint8) uint32 {
	return uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a)
}

// SolidColorImage creates raw RGBA bytes for a solid color image.
func SolidColorImage(width, height int, r, g, b, a uint8) []byte {
	size := width * height * 4
	rgba := make([]byte, size)

	for i := 0; i < size; i += 4 {
		rgba[i+0] = r
		rgba[i+1] = g
		rgba[i+2] = b
		rgba[i+3] = a
	}

	return rgba
}
