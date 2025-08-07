package sd

// #include "stable-diffusion.h"
import "C"

type Scheduler C.enum_schedule_t

func (s Scheduler) internal() C.enum_schedule_t {
	return C.enum_schedule_t(s)
}

const (
	Default     Scheduler = C.DEFAULT
	Discrete    Scheduler = C.DISCRETE
	Karras      Scheduler = C.KARRAS
	Exponential Scheduler = C.EXPONENTIAL
	Ays         Scheduler = C.AYS
	Gits        Scheduler = C.GITS
)
