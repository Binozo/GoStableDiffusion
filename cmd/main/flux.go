package main

import (
	"fmt"
	"image/png"
	"os"
	"time"
	"unsafe"

	"github.com/binozo/gostablediffusion/pkg/sd"
)

func main() {
	fmt.Println("Warming up")

	sd.SetLogCallback(func(level sd.LogLevel, text string, data unsafe.Pointer) {
		switch level {
		case sd.LogDebug:
			fmt.Print("DEBUG:", text)
		case sd.LogInfo:
			fmt.Print("INFO:", text)
		case sd.LogWarn:
			fmt.Print("WARN:", text)
		case sd.LogError:
			fmt.Print("ERROR:", text)
		default:
			fmt.Print("???:", text)
		}
	})

	sd.SetProgressCallback(func(step, steps int, time time.Duration, data unsafe.Pointer) {
		fmt.Printf("PROGRESS: Completed step %d of %d in %0.2fs\n", step, steps, time.Seconds())
	})

	ctx, err := sd.New().
		SetDiffusionModel("models/flux1-dev-q8_0.gguf").
		SetClipL("models/clip_l.safetensors").
		SetVaePath("models/ae.safetensors").
		SetT5xxlPath("models/t5xxl_fp16.safetensors").
		SetLoRaDir("models/lora").
		UseFlashAttn().
		Load()

	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	params := sd.NewImageGenerationParams()
	params.Guidance.TxtCfg = 1
	params.SampleMethod = sd.Euler
	params.Height = 512
	params.Width = 512
	params.Prompt = "a lovely cat holding a sign says 'flux.cpp'<lora:realism_lora_comfy_converted:1>"

	fmt.Println("Running inference")
	result := ctx.GenerateImage(params)

	fmt.Println("Writing result to output.png")
	targetFile, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer targetFile.Close()
	if err := png.Encode(targetFile, result.Image()); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
