package sd

/*
#cgo LDFLAGS: -L../../stable-diffusion.cpp/build/bin -lstable-diffusion
#include "../../include/stable-diffusion.h"
*/
import "C"

func GetSystemInfo() string {
	return C.GoString(C.sd_get_system_info())
}
