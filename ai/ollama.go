// Erasmo Cardoso - Software Engineer |Electronics Technician
package ai

import (
	"context"
)

type OllamaProvider struct {
	Model string
	Url   string
}

func (p *OllamaProvider) Chat(ctx context.Context, messages []Message, systemPrompt string, tools []Tool) (string, []ToolCall, error) {
	// Ollama suporta a API da OpenAI no endpoint /v1/chat/completions
	// Então podemos reutilizar a lógica da OpenAI
	internal := &OpenAIProvider{
		ApiKey: "ollama", // Valor dummy, Ollama local não costuma exigir chave
		Model:  p.Model,
		Url:    p.Url + "/v1/chat/completions",
	}
	return internal.Chat(ctx, messages, systemPrompt, tools)
}

func (p *OllamaProvider) GetName() string {
	return "Ollama (Local)"
}
