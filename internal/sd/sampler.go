package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
import "C"

type SampleMethod C.enum_sample_method_t

func (s SampleMethod) internal() C.enum_sample_method_t {
	return C.enum_sample_method_t(s)
}

const (
	EulerA       SampleMethod = C.EULER_A
	Euler        SampleMethod = C.EULER
	Heun         SampleMethod = C.HEUN
	Dpm2         SampleMethod = C.DPM2
	DpmPP2SA     SampleMethod = C.DPMPP2S_A
	DpmPP2M      SampleMethod = C.DPMPP2M
	DpmPP2Mv2    SampleMethod = C.DPMPP2Mv2
	IpnDm        SampleMethod = C.IPNDM
	IpnDmV       SampleMethod = C.IPNDM_V
	Lcm          SampleMethod = C.LCM
	DdimTrailing SampleMethod = C.DDIM_TRAILING
	Tcd          SampleMethod = C.TCD
)
