package sd

// #include "stable-diffusion.h"
import "C"

type Rng C.enum_rng_type_t

func (r Rng) internal() C.enum_rng_type_t {
	return C.enum_rng_type_t(r)
}

const (
	StdDefaultRng Rng = C.STD_DEFAULT_RNG
	CudaRng       Rng = C.CUDA_RNG
)
