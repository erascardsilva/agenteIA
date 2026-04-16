// Erasmo Cardoso - Software Engineer |Electronics Technician
package ai

import (
	"context"
)

type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
}

type AIProvider interface {
	Chat(ctx context.Context, messages []Message, systemPrompt string, tools []Tool) (string, []ToolCall, error)
	GetName() string
}
