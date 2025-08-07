package sd

import (
	"github.com/binozo/gostablediffusion/internal/sd"
)

func (c *Context) GenerateImage(options *ImageGenerationOptions) sd.Image {
	return c.internal.GenerateImage((*sd.ImageGenerationOptions)(options))
}
