package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type SlgParams struct {
	Layers     []int
	LayerStart float32
	LayerEnd   float32
	Scale      float32

	internalParams *C.sd_slg_params_t
}

func (p *SlgParams) internal() *C.sd_slg_params_t {
	if p.internalParams != nil {
		return p.internalParams
	}

	params := (*C.sd_slg_params_t)(C.malloc(C.sizeof_sd_slg_params_t))

	if p.Layers != nil && len(p.Layers) > 0 {
		layersC := (*C.int)(C.malloc(C.size_t(len(p.Layers)) * C.sizeof_int))
		layersSlice := unsafe.Slice(layersC, len(p.Layers))
		for i, layer := range p.Layers {
			layersSlice[i] = C.int(layer)
		}
		params.layers = layersC
		params.layer_count = C.size_t(len(p.Layers))
	} else {
		params.layers = nil
		params.layer_count = 0
	}

	params.layer_start = C.float(p.LayerStart)
	params.layer_end = C.float(p.LayerEnd)
	params.scale = C.float(p.Scale)

	p.internalParams = params
	return p.internalParams
}

func (p *SlgParams) Free() {
	if p.internalParams == nil {
		return
	}

	if p.internalParams.layers != nil {
		C.free(unsafe.Pointer(p.internalParams.layers))
		p.internalParams.layers = nil
	}

	C.free(unsafe.Pointer(p.internalParams))
	p.internalParams = nil
}
