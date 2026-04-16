// Erasmo Cardoso - Software Engineer |Electronics Technician
package audio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AudioModule struct {
	ApiKey string
}

func NewAudioModule(basePath string) *AudioModule {
	return &AudioModule{}
}

func (m *AudioModule) TextToSpeech(ctx context.Context, text string, voice string) ([]byte, error) {
	if m.ApiKey == "" {
		return nil, fmt.Errorf("API Key não configurada para Áudio")
	}

	if voice == "" {
		voice = "alloy"
	}

	url := "https://api.openai.com/v1/audio/speech"
	payload := map[string]string{
		"model": "tts-1",
		"input": text,
		"voice": voice,
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+m.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro no TTS: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// Whisper STT via Groq (Grátis) ou OpenAI
