package sd

import (
	"math"

	"github.com/binozo/gostablediffusion/internal/sd"
)

var INFINITY float32 = float32(math.Inf(1))

type ImageGenerationOptions sd.ImageGenerationOptions

func NewImageGenerationParams() *ImageGenerationOptions {
	return &ImageGenerationOptions{
		ClipSkip: -1,
		Guidance: sd.Guidance{
			TxtCfg:            7,
			MinCfg:            1,
			ImgCfg:            INFINITY,
			DistilledGuidance: 3.5,
			SlgParams: sd.SlgParams{
				Layers:     []int{7, 8, 9},
				LayerStart: 0.01,
				LayerEnd:   0.2,
				Scale:      0,
			},
		},
		Width:         512,
		Height:        512,
		SampleMethod:  sd.EulerA,
		SampleSteps:   20,
		Strength:      0.75,
		Seed:          -1,
		BatchCount:    1,
		StyleStrength: 20,
	}
}
