package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type Guidance struct {
	// TxtCfg is just normal Cfg
	TxtCfg            float32
	ImgCfg            float32
	MinCfg            float32
	DistilledGuidance float32
	SlgParams         SlgParams

	internalGuidance *C.sd_guidance_params_t
}

func (g *Guidance) internal() *C.sd_guidance_params_t {
	if g.internalGuidance != nil {
		return g.internalGuidance
	}

	var guidance *C.sd_guidance_params_t = (*C.sd_guidance_params_t)(C.malloc(C.sizeof_sd_guidance_params_t))

	guidance.txt_cfg = C.float(g.TxtCfg)
	guidance.img_cfg = C.float(g.ImgCfg)
	guidance.min_cfg = C.float(g.MinCfg)
	guidance.distilled_guidance = C.float(g.DistilledGuidance)
	guidance.slg = *g.SlgParams.internal()

	g.internalGuidance = guidance
	return guidance
}

func (g *Guidance) Free() {
	if g.internalGuidance == nil {
		return
	}

	g.SlgParams.Free()
	C.free(unsafe.Pointer(g.internalGuidance))
	g.internalGuidance = nil
}
