package sd

// #include "stable-diffusion.h"
// #include <stdlib.h>
import "C"
import (
	"image"
	"unsafe"
)

type Image struct {
	width, height int
	channel       int
	data          []uint8
	internalImage *C.sd_image_t
}

func (i *Image) internal() *C.sd_image_t {
	if i.internalImage != nil {
		return i.internalImage
	}

	var internalImage *C.sd_image_t = (*C.sd_image_t)(C.malloc(C.sizeof_sd_image_t))
	i.internalImage = internalImage

	internalImage.width = C.uint32_t(i.width)
	internalImage.height = C.uint32_t(i.height)
	internalImage.channel = C.uint32_t(i.channel)

	if i.data != nil && len(i.data) > 0 {
		size := len(i.data)
		internalImage.data = (*C.uint8_t)(C.malloc(C.size_t(size)))

		C.memcpy(unsafe.Pointer(internalImage.data), unsafe.Pointer(&i.data[0]), C.size_t(size))
	} else {
		internalImage.data = nil
	}

	return internalImage
}

func (i *Image) Free() {
	if i.internalImage != nil {
		if i.internalImage.data != nil {
			C.free(unsafe.Pointer(i.internalImage.data))
			i.internalImage.data = nil
		}
		C.free(unsafe.Pointer(i.internalImage))
		i.internalImage = nil

	}
}

func (c *Ctx) GenerateImage(options *ImageGenerationOptions) Image {
	defer options.Free()
	var imageOptions *C.sd_img_gen_params_t = options.internal()
	var generatedImage *C.sd_image_t = C.generate_image(c.internal, imageOptions)

	generatedImageSize := uint32(generatedImage.width) * uint32(generatedImage.height) * uint32(generatedImage.channel)

	return Image{
		internalImage: generatedImage,
		width:         options.Width,
		height:        options.Height,
		channel:       int(generatedImage.channel),
		data:          unsafe.Slice((*uint8)(unsafe.Pointer(generatedImage.data)), generatedImageSize),
	}
}

func (i *Image) Image() *image.RGBA {
	return parseImage(i)
}

func parseImage(img *Image) *image.RGBA {
	if img.channel != 3 {
		return nil
	}

	rect := image.Rect(0, 0, img.width, img.height)
	return &image.RGBA{
		Pix:    convertRGBtoRGBA(img.data, img.width, img.height),
		Rect:   rect,
		Stride: img.width * 4, // By default 3 channels are getting generated but for RGBA we need 4
	}
}

func convertRGBtoRGBA(data []uint8, width, height int) []uint8 {
	rgbaData := make([]uint8, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			offsetRGB := (y*width + x) * 3
			offsetRGBA := (y*width + x) * 4
			rgbaData[offsetRGBA] = data[offsetRGB]     // R
			rgbaData[offsetRGBA+1] = data[offsetRGB+1] // G
			rgbaData[offsetRGBA+2] = data[offsetRGB+2] // B
			rgbaData[offsetRGBA+3] = 255               // Alpha
		}
	}
	return rgbaData
}
