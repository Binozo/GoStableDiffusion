package sd

/*
#include "stable-diffusion.h"
#include <stdlib.h>

extern void goLogCallback(enum sd_log_level_t level, char* text, void* data);

static inline void cLogCallback(enum sd_log_level_t level, const char* text, void* data) {
    goLogCallback(level, (char*)text, data);
}
*/
import "C"
import "unsafe"

type LogLevel C.enum_sd_log_level_t

const (
	LogDebug LogLevel = C.SD_LOG_DEBUG
	LogInfo  LogLevel = C.SD_LOG_INFO
	LogWarn  LogLevel = C.SD_LOG_WARN
	LogError LogLevel = C.SD_LOG_ERROR
)

type LogCallback func(level LogLevel, text string, data unsafe.Pointer)

var logCallback LogCallback

//export goLogCallback
func goLogCallback(level C.enum_sd_log_level_t, text *C.char, data unsafe.Pointer) {
	if logCallback != nil {
		logCallback(LogLevel(level), C.GoString(text), data)
	}
}

func SetLogCallback(callback LogCallback) {
	logCallback = callback
	C.sd_set_log_callback((C.sd_log_cb_t)(C.goLogCallback), nil)
}
