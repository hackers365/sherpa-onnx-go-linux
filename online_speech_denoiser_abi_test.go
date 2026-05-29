package sherpa_onnx

import (
	"os"
	"strings"
	"testing"
)

func TestSpeechDenoiserCAPIHeaderIncludesDpdfnetField(t *testing.T) {
	b, err := os.ReadFile("c-api.h")
	if err != nil {
		t.Fatal(err)
	}

	header := string(b)
	if !strings.Contains(header, "SherpaOnnxOfflineSpeechDenoiserDpdfNetModelConfig") {
		t.Fatal("c-api.h is missing SherpaOnnxOfflineSpeechDenoiserDpdfNetModelConfig")
	}
	if !strings.Contains(header, "SherpaOnnxOfflineSpeechDenoiserDpdfNetModelConfig dpdfnet;") {
		t.Fatal("c-api.h is missing SherpaOnnxOfflineSpeechDenoiserModelConfig.dpdfnet")
	}
}
