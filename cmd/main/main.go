package main

import (
	"GoStableDiffusion/pkg/sd"
	"fmt"
	"image/png"
	"os"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("Warming up")
	fmt.Println("Running with Cuda:", sd.HasCuda())

	ctx, err := sd.NewDefault("models/sd-v1-4.ckpt")
	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	sd.SetLogCallback(func(level sd.LogLevel, text string, data unsafe.Pointer) {
		switch level {
		case sd.LogDebug:
			fmt.Println("DEBUG:", text)
		case sd.LogInfo:
			fmt.Println("INFO:", text)
		case sd.LogError:
			fmt.Println("ERROR:", text)
		default:
			fmt.Println("???:", text)
		}
	})

	sd.SetProgressCallback(func(step, steps int, time time.Duration, data unsafe.Pointer) {
		fmt.Printf("PROGRESS: Completed step %d of %d in %0.2fs", step, steps, time.Seconds())
	})

	params := sd.NewDefaultParams()
	params.Prompt = "A beautiful and sunny landscape with lots of dogs"

	fmt.Println("Running inference")
	result := ctx.Text2Img(params)

	fmt.Println("Writing result to output.png")
	targetFile, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer targetFile.Close()
	if err := png.Encode(targetFile, result.Image()); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
