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
		SetModel("models/sd-v1-4.ckpt").
		Load()

	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	params := sd.NewImageGenerationParams()
	params.Guidance.TxtCfg = 3
	params.SampleSteps = 30
	params.SampleMethod = sd.Euler
	params.Height = 768
	params.Width = 768
	params.Seed = 42
	params.Prompt = "fantasy medieval village world inside a glass sphere , high detail, fantasy, realistic, light effect, hyper detail, volumetric lighting, cinematic, macro, depth of field, blur, red light and clouds from the back, highly detailed epic cinematic concept art cg render made in maya, blender and photoshop, octane render, excellent composition, dynamic dramatic cinematic lighting, aesthetic, very inspirational, world inside a glass sphere by james gurney by artgerm with james jean, joe fenton and tristan eaton by ross tran, fine details, 4k resolution"

	fmt.Println("Running inference")
	result := ctx.GenerateImage(params)
	defer result.Free()

	fmt.Println("Writing result to output.png")
	targetFile, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer targetFile.Close()
	if err = png.Encode(targetFile, result.Image()); err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
