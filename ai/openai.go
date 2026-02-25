// Erasmo Cardoso - Dev
package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAIProvider struct {
	ApiKey string
	Model  string
	Url    string
}

func (p *OpenAIProvider) Chat(ctx context.Context, messages []Message, systemPrompt string, tools []Tool) (string, []ToolCall, error) {
	var msgs []map[string]interface{}
	msgs = append(msgs, map[string]interface{}{"role": "system", "content": systemPrompt})

	for _, m := range messages {
		var content interface{} = m.Content
		// Groq/OpenAI: Se houver tool_calls e content for "", alguns modelos preferem "" em vez de null
		if m.Role == "assistant" && len(m.ToolCalls) > 0 && m.Content == "" {
			content = ""
		}

		msg := map[string]interface{}{
			"role":    m.Role,
			"content": content,
		}
		if m.ToolCallID != "" {
			msg["tool_call_id"] = m.ToolCallID
		}
		if len(m.ToolCalls) > 0 {
			msg["tool_calls"] = m.ToolCalls
		}
		msgs = append(msgs, msg)
	}

	payload := map[string]interface{}{
		"model":       p.Model,
		"messages":    msgs,
		"tool_choice": "auto",
	}

	if len(tools) > 0 {
		payload["tools"] = tools
	}

	// Adicionar tool_call_id e tool_calls às mensagens se presentes
	// (O OpenAI/Groq exige isso para mensagens de ferramenta e respostas que acionam ferramentas)

	body, err := json.Marshal(payload)
	if err != nil {
		return "", nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.Url, bytes.NewBuffer(body))
	if err != nil {
		return "", nil, err
	}

	req.Header.Set("Authorization", "Bearer "+p.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errData map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errData)
		return "", nil, fmt.Errorf("API error: %v", errData)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content   string     `json:"content"`
				ToolCalls []ToolCall `json:"tool_calls"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", nil, err
	}

	if len(result.Choices) == 0 {
		return "", nil, fmt.Errorf("no response from AI")
	}

	msg := result.Choices[0].Message
	return msg.Content, msg.ToolCalls, nil
}

func (p *OpenAIProvider) GetName() string {
	return "OpenAI/Groq"
}
