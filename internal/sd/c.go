package sd

/*
#cgo LDFLAGS: -L../../stable-diffusion.cpp/build/bin -lstable-diffusion
#cgo darwin LDFLAGS: -Wl,-rpath,\@loader_path
#cgo linux LDFLAGS: -Wl,-rpath,$ORIGIN
#include "../../include/stable-diffusion.h"
*/
import "C"

func GetSystemInfo() string {
	return C.GoString(C.sd_get_system_info())
}
