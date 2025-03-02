package sd

/*
#include "../../stable-diffusion.cpp/stable-diffusion.h"

// Deklaration der Go-Callback-Funktion
extern void goProgressCallback(int step, int steps, float time, void* data);

// C-Trampoline-Funktion, die den Go-Callback aufruft
static inline void cProgressCallback(int step, int steps, float time, void* data) {
    goProgressCallback(step, steps, time, data);
}
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

type ProgressCallback func(step, steps int, time time.Duration, data unsafe.Pointer)

var progressCallback ProgressCallback

//export goProgressCallback
func goProgressCallback(step C.int, steps C.int, seconds C.float, data unsafe.Pointer) {
	fmt.Printf("Progress: Step %d of %d, Time: %.2f\n", int(step), int(steps), float32(seconds))
	if progressCallback != nil {
		parsedDuration, err := time.ParseDuration(fmt.Sprintf("%fs", float32(seconds)))
		if err != nil {
			parsedDuration = time.Duration(0)
		}
		progressCallback(int(steps), int(steps), parsedDuration, data)
	}
}

func SetProgressCallback(callback ProgressCallback) {
	progressCallback = callback
	C.sd_set_progress_callback((C.sd_progress_cb_t)(C.goProgressCallback), nil)
}
