package sherpa_onnx

import "testing"

func TestOnlineSpeechDenoiserEngineAPICompile(t *testing.T) {
	_ = OnlineSpeechDenoiserConfig{
		Model: OnlineSpeechDenoiserModelConfig{
			Gtcrn: OnlineSpeechDenoiserGtcrnModelConfig{
				Model: "model.onnx",
			},
			NumThreads: 1,
			Debug:      0,
			Provider:   "cpu",
		},
		PoolSize: 1,
	}

	var _ *OnlineSpeechDenoiserEngine
	var _ *OnlineSpeechDenoiserStream

	var _ func(*OnlineSpeechDenoiserConfig) *OnlineSpeechDenoiserEngine = NewOnlineSpeechDenoiserEngine
	var _ func(*OnlineSpeechDenoiserEngine) = DeleteOnlineSpeechDenoiserEngine
	var _ func(*OnlineSpeechDenoiserEngine) int = (*OnlineSpeechDenoiserEngine).SampleRate
	var _ func(*OnlineSpeechDenoiserEngine) int = (*OnlineSpeechDenoiserEngine).FrameShiftInSamples
	var _ func(*OnlineSpeechDenoiserEngine) *OnlineSpeechDenoiserStream = (*OnlineSpeechDenoiserEngine).CreateStream
	var _ func(*OnlineSpeechDenoiserStream, []float32, int) *DenoisedAudio = (*OnlineSpeechDenoiserStream).Run
	var _ func(*OnlineSpeechDenoiserStream) *DenoisedAudio = (*OnlineSpeechDenoiserStream).Flush
	var _ func(*OnlineSpeechDenoiserStream) = (*OnlineSpeechDenoiserStream).Reset
	var _ func(*OnlineSpeechDenoiserStream) = DeleteOnlineSpeechDenoiserStream
}
