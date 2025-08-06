package sd

// #include "stable-diffusion.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type SlgParams struct {
	Layers     int
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

	layers := (*C.int)(C.malloc(C.sizeof_int))
	*layers = C.int(p.Layers)

	params.layers = layers
	params.layer_count = 0
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
