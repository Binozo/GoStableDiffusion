package sd

// #include "stable-diffusion.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type ImageGenerationOptions struct {
	Prompt          string
	NegativePrompt  string
	ClipSkip        int
	Guidance        Guidance
	initImage       Image
	maskImage       Image
	Width           int
	Height          int
	SampleMethod    SampleMethod
	SampleSteps     int
	Eta             float32
	Strength        float32
	Seed            int64
	BatchCount      int
	ControlStrength float32
	StyleStrength   float32
	NormalizeInput  bool

	internalParams *C.sd_img_gen_params_t
}

func (i *ImageGenerationOptions) internal() *C.sd_img_gen_params_t {
	if i.internalParams != nil {
		return i.internalParams
	}

	var params *C.sd_img_gen_params_t = (*C.sd_img_gen_params_t)(C.malloc(C.sizeof_sd_img_gen_params_t))

	params.prompt = C.CString(i.Prompt)
	params.negative_prompt = C.CString(i.NegativePrompt)
	params.clip_skip = C.int(i.ClipSkip)
	params.guidance = *i.Guidance.internal()

	initImage := Image{
		width:   i.Width,
		height:  i.Height,
		channel: 3,
		data:    nil,
	}
	i.initImage = initImage
	params.init_image = *initImage.internal()

	maskImageData := make([]byte, i.Width*i.Height)

	for index := range maskImageData {
		maskImageData[index] = byte(255)
	}

	maskImage := Image{
		width:   i.Width,
		height:  i.Height,
		channel: 1,
		data:    maskImageData,
	}
	i.maskImage = maskImage

	params.mask_image = *maskImage.internal()
	params.width = C.int(i.Width)
	params.height = C.int(i.Height)
	params.sample_method = i.SampleMethod.internal()
	params.sample_steps = C.int(i.SampleSteps)
	params.eta = C.float(i.Eta)
	params.strength = C.float(i.Strength)
	params.seed = C.int64_t(i.Seed)
	params.batch_count = C.int(i.BatchCount)
	params.control_strength = C.float(i.ControlStrength)
	params.style_strength = C.float(i.StyleStrength)
	params.normalize_input = C.bool(i.NormalizeInput)

	i.internalParams = params
	return params
}

func (i *ImageGenerationOptions) Free() {
	if i.internalParams == nil {
		return
	}

	C.free(unsafe.Pointer(i.internalParams.prompt))
	C.free(unsafe.Pointer(i.internalParams.negative_prompt))

	i.Guidance.Free()

	i.initImage.Free()
	i.maskImage.Free()

	C.free(unsafe.Pointer(i.internalParams))
	i.internalParams = nil
}
