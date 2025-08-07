package sd

// #include "stable-diffusion.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type CtxParams struct {
	internal *C.sd_ctx_params_t
}

func NewContextParams(
	modelPath, clipLPath, clipGPath, t5xxlPath, diffusionModelPath, vaePath, taeSdPath, controlNetPath, loraModelDir, embedDir, stackedIdEmbedDir string,
	vaeDecodingOnly, vaeTiling, freeParamsImmediately bool,
	nThreads int,
	wType Type,
	rngType Rng,
	schedule Scheduler,
	keepClipOnCpu, keepControlNetCpu, keepVaeOnCpu, diffusionFlashAttn bool,
) *CtxParams {
	var cParams *C.sd_ctx_params_t = (*C.sd_ctx_params_t)(C.malloc(C.sizeof_sd_ctx_params_t))
	C.sd_ctx_params_init(cParams)

	cParams.model_path = C.CString(modelPath)
	cParams.clip_l_path = C.CString(clipLPath)
	cParams.clip_g_path = C.CString(clipGPath)
	cParams.t5xxl_path = C.CString(t5xxlPath)
	cParams.diffusion_model_path = C.CString(diffusionModelPath)
	cParams.vae_path = C.CString(vaePath)
	cParams.taesd_path = C.CString(taeSdPath)
	cParams.control_net_path = C.CString(controlNetPath)
	cParams.lora_model_dir = C.CString(loraModelDir)
	cParams.embedding_dir = C.CString(embedDir)
	cParams.stacked_id_embed_dir = C.CString(stackedIdEmbedDir)
	cParams.vae_decode_only = C.bool(vaeDecodingOnly)
	cParams.vae_tiling = C.bool(vaeTiling)
	cParams.free_params_immediately = C.bool(freeParamsImmediately)
	cParams.n_threads = C.int(nThreads)
	cParams.wtype = wType.internal()
	cParams.rng_type = rngType.internal()
	cParams.schedule = schedule.internal()
	cParams.keep_clip_on_cpu = C.bool(keepClipOnCpu)
	cParams.keep_control_net_on_cpu = C.bool(keepControlNetCpu)
	cParams.keep_vae_on_cpu = C.bool(keepVaeOnCpu)
	cParams.diffusion_flash_attn = C.bool(diffusionFlashAttn)

	// Unsupported
	cParams.chroma_use_dit_mask = C.bool(false)
	cParams.chroma_use_t5_mask = C.bool(false)
	cParams.chroma_t5_mask_pad = C.int(0)

	return &CtxParams{
		internal: cParams,
	}
}

func (c *CtxParams) Free() {
	if c.internal == nil {
		return
	}

	C.free(unsafe.Pointer(c.internal.model_path))
	C.free(unsafe.Pointer(c.internal.clip_l_path))
	C.free(unsafe.Pointer(c.internal.clip_g_path))
	C.free(unsafe.Pointer(c.internal.t5xxl_path))
	C.free(unsafe.Pointer(c.internal.diffusion_model_path))
	C.free(unsafe.Pointer(c.internal.vae_path))
	C.free(unsafe.Pointer(c.internal.taesd_path))
	C.free(unsafe.Pointer(c.internal.control_net_path))
	C.free(unsafe.Pointer(c.internal.lora_model_dir))
	C.free(unsafe.Pointer(c.internal.embedding_dir))
	C.free(unsafe.Pointer(c.internal.stacked_id_embed_dir))

	C.free(unsafe.Pointer(c.internal))
	c.internal = nil
}
