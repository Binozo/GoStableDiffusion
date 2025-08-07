package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
// #include <stdlib.h>
import "C"

type Ctx struct {
	internal *C.sd_ctx_t
	params   *CtxParams
}

func NewSdContext(params *CtxParams) *Ctx {
	var ctx *C.sd_ctx_t = C.new_sd_ctx(params.internal)

	return &Ctx{
		internal: ctx,
		params:   params,
	}
}

func (c *Ctx) Free() {
	if c.params != nil {
		c.params.Free()
		c.params = nil
	}

	if c.internal != nil {
		C.free_sd_ctx(c.internal)
		c.internal = nil
	}
}
