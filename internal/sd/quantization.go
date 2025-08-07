package sd

// #include "../../stable-diffusion.cpp/stable-diffusion.h"
import "C"

type Type C.enum_sd_type_t

func (t Type) internal() C.enum_sd_type_t {
	return C.enum_sd_type_t(t)
}

const (
	F32         Type = C.SD_TYPE_F32
	F16         Type = C.SD_TYPE_F16
	Q4_0        Type = C.SD_TYPE_Q4_0
	Q4_1        Type = C.SD_TYPE_Q4_1
	Q5_0        Type = C.SD_TYPE_Q5_0
	Q5_1        Type = C.SD_TYPE_Q5_1
	Q8_0        Type = C.SD_TYPE_Q8_0
	Q8_1        Type = C.SD_TYPE_Q8_1
	Q2_K        Type = C.SD_TYPE_Q2_K
	Q3_K        Type = C.SD_TYPE_Q3_K
	Q4_K        Type = C.SD_TYPE_Q4_K
	Q5_K        Type = C.SD_TYPE_Q5_K
	Q6_K        Type = C.SD_TYPE_Q6_K
	Q8_K        Type = C.SD_TYPE_Q8_K
	IQ2_XSS     Type = C.SD_TYPE_IQ2_XXS
	IQ2_XS      Type = C.SD_TYPE_IQ2_XS
	IQ3_XXS     Type = C.SD_TYPE_IQ3_XXS
	IQ1_S       Type = C.SD_TYPE_IQ1_S
	IQ4_NL      Type = C.SD_TYPE_IQ4_NL
	IQ3_S       Type = C.SD_TYPE_IQ3_S
	IQ2_S       Type = C.SD_TYPE_IQ2_S
	IQ4_XS      Type = C.SD_TYPE_IQ4_XS
	I8          Type = C.SD_TYPE_I8
	I16         Type = C.SD_TYPE_I16
	I32         Type = C.SD_TYPE_I32
	I64         Type = C.SD_TYPE_I64
	F64         Type = C.SD_TYPE_F64
	IQ1_M       Type = C.SD_TYPE_IQ1_M
	BF16        Type = C.SD_TYPE_BF16
	TQ1_0       Type = C.SD_TYPE_TQ1_0
	TQ2_0       Type = C.SD_TYPE_TQ2_0
	Unspecified Type = C.SD_TYPE_COUNT
)

func (t Type) Name() string {
	return C.GoString(C.sd_type_name(C.enum_sd_type_t(t)))
}
