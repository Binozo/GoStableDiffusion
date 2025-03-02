package sd

import (
	"GoStableDiffusion/internal/sd"
	"errors"
	"os"
)

type ContextParams struct {
	ModelPath             string
	ClipLPath             string
	ClipGPath             string
	T5xxlPath             string
	DiffusionModelPath    string
	VaePath               string
	TaeSdPath             string
	ControlNetPath        string
	LoraModelDir          string
	EmbedDir              string
	StackedIdEmbedDir     string
	VaeDecodeOnly         bool
	VaeTiling             bool
	FreeParamsImmediately bool
	NThreads              int
	WType                 sd.Type
	RngType               sd.Rng
	Schedule              sd.Scheduler
	KeepClipOnCpu         bool
	KeepControlNetCpu     bool
	KeepVaeOnCpu          bool
	DiffusionFlashAttn    bool
}

func (cp *ContextParams) validate() error {
	if cp.ModelPath == "" {
		return errors.New("model path required")
	} else if !fileExists(cp.ModelPath) {
		return errors.New("model path does not exist")
	}

	if cp.ClipLPath != "" && !fileExists(cp.ClipLPath) {
		return errors.New("clip l path does not exist")
	}

	if cp.ClipGPath != "" && !fileExists(cp.ClipGPath) {
		return errors.New("clip g path does not exist")
	}

	if cp.T5xxlPath != "" && !fileExists(cp.T5xxlPath) {
		return errors.New("t5xxl path does not exist")
	}

	if cp.DiffusionModelPath != "" && !fileExists(cp.DiffusionModelPath) {
		return errors.New("diffusion model path does not exist")
	}

	if cp.VaePath != "" && !fileExists(cp.VaePath) {
		return errors.New("va path does not exist")
	}

	if cp.TaeSdPath != "" && !fileExists(cp.TaeSdPath) {
		return errors.New("tae-sd path does not exist")
	}

	if cp.ControlNetPath != "" && !fileExists(cp.ControlNetPath) {
		return errors.New("control net path does not exist")
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
