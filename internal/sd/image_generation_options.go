package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type ImageGenerationOptions struct {
	Prompt            string
	NegativePrompt    string
	ClipSkip          int
	Guidance          Guidance
	initImage         *Image
	refImages         []Image
	maskImage         *Image
	Width             int
	Height            int
	SampleMethod      SampleMethod
	SampleSteps       int
	Eta               float32
	Strength          float32
	Seed              int64
	BatchCount        int
	ControlCond       *Image
	ControlStrength   float32
	StyleStrength     float32
	NormalizeInput    bool
	InputIdImagesPath string

	internalParams   *C.sd_img_gen_params_t
	cPrompt          *C.char
	cNegativePrompt  *C.char
	cInputIdPath     *C.char
	cRefImages       *C.sd_image_t
	defaultMaskImage *Image
}

func (i *ImageGenerationOptions) internal() *C.sd_img_gen_params_t {
	if i.internalParams != nil {
		return i.internalParams
	}

	var params *C.sd_img_gen_params_t = (*C.sd_img_gen_params_t)(C.malloc(C.sizeof_sd_img_gen_params_t))
	C.sd_img_gen_params_init(params)

	i.cPrompt = C.CString(i.Prompt)
	params.prompt = i.cPrompt

	i.cNegativePrompt = C.CString(i.NegativePrompt)
	params.negative_prompt = i.cNegativePrompt

	params.clip_skip = C.int(i.ClipSkip)
	params.guidance = *i.Guidance.internal()

	if i.initImage != nil {
		params.init_image = *i.initImage.internal()
	}

	if len(i.refImages) > 0 {
		size := C.size_t(len(i.refImages)) * C.sizeof_sd_image_t
		i.cRefImages = (*C.sd_image_t)(C.malloc(size))

		refImagesSlice := unsafe.Slice(i.cRefImages, len(i.refImages))

		for idx, refImg := range i.refImages {
			refImagesSlice[idx] = *refImg.internal()
		}

		params.ref_images = i.cRefImages
		params.ref_images_count = C.int(len(i.refImages))

	} else {
		params.ref_images = nil
		params.ref_images_count = 0
	}

	if i.maskImage != nil {
		params.mask_image = *i.maskImage.internal()
	} else if i.initImage != nil {
		// Default white mask
		maskImageData := make([]byte, i.Width*i.Height)
		for index := range maskImageData {
			maskImageData[index] = byte(255)
		}

		i.defaultMaskImage = &Image{
			width:   i.Width,
			height:  i.Height,
			channel: 1,
			data:    maskImageData,
		}
		params.mask_image = *i.defaultMaskImage.internal()
	}

	params.width = C.int(i.Width)
	params.height = C.int(i.Height)
	params.sample_method = i.SampleMethod.internal()
	params.sample_steps = C.int(i.SampleSteps)
	params.eta = C.float(i.Eta)
	params.strength = C.float(i.Strength)
	params.seed = C.int64_t(i.Seed)
	params.batch_count = C.int(i.BatchCount)

	if i.ControlCond != nil {
		params.control_cond = i.ControlCond.internal()
	} else {
		params.control_cond = nil
	}

	params.control_strength = C.float(i.ControlStrength)
	params.style_strength = C.float(i.StyleStrength)
	params.normalize_input = C.bool(i.NormalizeInput)

	i.cInputIdPath = C.CString(i.InputIdImagesPath)
	params.input_id_images_path = i.cInputIdPath

	i.internalParams = params
	return params
}

func (i *ImageGenerationOptions) Free() {
	if i.internalParams == nil {
		return
	}

	if i.cPrompt != nil {
		C.free(unsafe.Pointer(i.cPrompt))
		i.cPrompt = nil
	}
	if i.cNegativePrompt != nil {
		C.free(unsafe.Pointer(i.cNegativePrompt))
		i.cNegativePrompt = nil
	}
	if i.cInputIdPath != nil {
		C.free(unsafe.Pointer(i.cInputIdPath))
		i.cInputIdPath = nil
	}

	if i.cRefImages != nil {
		C.free(unsafe.Pointer(i.cRefImages))
		i.cRefImages = nil
	}

	i.Guidance.Free()

	if i.initImage != nil {
		i.initImage.Free()
	}
	if i.maskImage != nil {
		i.maskImage.Free()
	}

	for _, refImg := range i.refImages {
		refImg.Free()
	}

	if i.ControlCond != nil {
		i.ControlCond.Free()
	}

	if i.defaultMaskImage != nil {
		i.defaultMaskImage.Free()
	}

	C.free(unsafe.Pointer(i.internalParams))
	i.internalParams = nil
}
