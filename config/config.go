// Erasmo Cardoso - Software Engineer |Electronics Technician
package config

import (
	"agente-ia/db"
	"encoding/json"
)

type Config struct {
	AssistantName     string            `json:"assistantName"`
	ApiKeys           map[string]string `json:"apiKeys"`
	PreferredProvider string            `json:"preferredProvider"` // "groq", "openai", "gemini"
	PreferredModel    string            `json:"preferredModel"`
	VoiceSettings     VoiceSettings     `json:"voiceSettings"`
	Context           ContextSettings   `json:"context"`
	Language          string            `json:"language"` // "en", "pt-BR"
}

type VoiceSettings struct {
	Engine  string `json:"engine"`
	VoiceID string `json:"voiceId"`
	Enabled bool   `json:"enabled"`
}

type ContextSettings struct {
	Humor         string `json:"humor"`
	UnlockModels  bool   `json:"unlockModels"`
	SystemProcess bool   `json:"systemProcess"`
	UserNome      string `json:"userNome"`
	UserProfissao string `json:"userProfissao"`
	UserIdade     string `json:"userIdade"`
	Funcoes       string `json:"funcoes"`
	Regras        string `json:"regras"`
	AutonomousMode bool   `json:"autonomousMode"`
}

func LoadConfig() (*Config, error) {
	data, err := db.GetConfig("main_config")
	if err != nil {
		return nil, err
	}

	if data == "" {
		return &Config{
			AssistantName:     "Jarvis",
			ApiKeys:           make(map[string]string),
			PreferredProvider: "groq",
			PreferredModel:    "llama-3.3-70b-versatile",
			VoiceSettings:     VoiceSettings{Enabled: false, Engine: "kokoro", VoiceID: "af_heart"},
			Context:           ContextSettings{Humor: "Prestativo", UnlockModels: false},
			Language:          "en",
		}, nil
	}

	var cfg Config
	err = json.Unmarshal([]byte(data), &cfg)
	return &cfg, err
}

func SaveConfig(cfg *Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return db.SetConfig("main_config", string(data))
}
