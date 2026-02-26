// Erasmo Cardoso - Dev
package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiProvider struct {
	ApiKey       string
	Model        string
	UnlockModels bool
}

func (p *GeminiProvider) Chat(ctx context.Context, messages []Message, systemPrompt string, tools []Tool) (string, []ToolCall, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(p.ApiKey))
	if err != nil {
		return "", nil, err
	}
	defer client.Close()

	// Configurar modelo. O Gemini exige nomes com o prefixo "models/"
	modelName := p.Model
	if !strings.HasPrefix(modelName, "models/") && !strings.Contains(modelName, "/") {
		modelName = "models/" + modelName
	}
	model := client.GenerativeModel(modelName)

	if p.UnlockModels {
		model.SafetySettings = []*genai.SafetySetting{
			{
				Category:  genai.HarmCategoryHarassment,
				Threshold: genai.HarmBlockNone,
			},
			{
				Category:  genai.HarmCategoryHateSpeech,
				Threshold: genai.HarmBlockNone,
			},
			{
				Category:  genai.HarmCategorySexuallyExplicit,
				Threshold: genai.HarmBlockNone,
			},
			{
				Category:  genai.HarmCategoryDangerousContent,
				Threshold: genai.HarmBlockNone,
			},
		}
	}
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemPrompt)},
	}

	// Converter ferramentas para Gemini
	var geminiTools []*genai.Tool
	availableTools := GetAvailableTools()
	if len(availableTools) > 0 {
		var toolDefs []*genai.FunctionDeclaration
		for _, t := range availableTools {
			decl := &genai.FunctionDeclaration{
				Name:        t.Function.Name,
				Description: t.Function.Description,
				Parameters: &genai.Schema{
					Type:       genai.TypeObject,
					Properties: map[string]*genai.Schema{},
				},
			}
			// Mapeamento simplificado de parâmetros para o Gemini
			params, ok := t.Function.Parameters.(map[string]interface{})
			if ok {
				props, ok := params["properties"].(map[string]interface{})
				if ok {
					for k, v := range props {
						propMap := v.(map[string]interface{})
						decl.Parameters.Properties[k] = &genai.Schema{
							Type:        genai.TypeString, // Simplificado para string
							Description: propMap["description"].(string),
						}
					}
				}
			}
			toolDefs = append(toolDefs, decl)
		}
		geminiTools = append(geminiTools, &genai.Tool{FunctionDeclarations: toolDefs})
		model.Tools = geminiTools
	}

	// Histórico para o chat do Gemini
	// O Gemini exige uma sequência estrita: user -> model (com function call) -> function (com response) -> model
	var geminiHistory []*genai.Content

	for i := 0; i < len(messages)-1; i++ {
		msg := messages[i]
		content := &genai.Content{}

		if msg.Role == "assistant" {
			content.Role = "model"
			if len(msg.ToolCalls) > 0 {
				for _, tc := range msg.ToolCalls {
					var args map[string]interface{}
					json.Unmarshal([]byte(tc.Function.Arguments), &args)
					content.Parts = append(content.Parts, genai.FunctionCall{
						Name: tc.Function.Name,
						Args: args,
					})
				}
			}
			if msg.Content != "" {
				content.Parts = append(content.Parts, genai.Text(msg.Content))
			} else if len(msg.ToolCalls) == 0 {
				continue
			}
		} else if msg.Role == "tool" {
			if msg.Content == "" {
				continue
			}
			content.Role = "function"
			content.Parts = append(content.Parts, genai.FunctionResponse{
				Name:     msg.ToolCallID,
				Response: map[string]interface{}{"result": msg.Content},
			})
		} else if msg.Role == "user" {
			content.Role = "user"
			content.Parts = append(content.Parts, genai.Text(msg.Content))
		} else {
			continue // Ignorar system ou outros roles não suportados no histórico
		}

		geminiHistory = append(geminiHistory, content)
	}

	// Poda de segurança para Gemini: O histórico deve começar com 'user'
	// e não pode terminar com 'function' se a próxima mensagem enviada não for a resposta.
	for len(geminiHistory) > 0 && geminiHistory[0].Role != "user" {
		geminiHistory = geminiHistory[1:]
	}

	chat := model.StartChat()
	chat.History = geminiHistory

	// Enviar última mensagem
	lastMsg := messages[len(messages)-1].Content
	resp, err := chat.SendMessage(ctx, genai.Text(lastMsg))
	if err != nil {
		return "", nil, fmt.Errorf("Gemini Error: %v (Verifique se sua chave tem permissão para o modelo %s)", err, p.Model)
	}

	if len(resp.Candidates) == 0 {
		return "", nil, fmt.Errorf("no response from Gemini")
	}

	candidate := resp.Candidates[0]
	fmt.Printf("Gemini Finish Reason: %v\n", candidate.FinishReason)

	var responseText string
	var toolCalls []ToolCall

	for _, part := range candidate.Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			responseText += string(txt)
		} else if fn, ok := part.(genai.FunctionCall); ok {
			args, _ := json.Marshal(fn.Args)
			toolCalls = append(toolCalls, ToolCall{
				ID:   fn.Name,
				Type: "function",
				Function: struct {
					Name      string `json:"name"`
					Arguments string `json:"arguments"`
				}{
					Name:      fn.Name,
					Arguments: string(args),
				},
			})
		}
	}

	if responseText == "" && len(toolCalls) == 0 {
		return "", nil, fmt.Errorf("Gemini retornou conteúdo vazio (Motivo: %v). Verifique se o modelo %s está ativo ou se houve bloqueio de segurança.", candidate.FinishReason, p.Model)
	}

	return responseText, toolCalls, nil
}

func (p *GeminiProvider) GetName() string {
	return "Gemini"
}

func ListGeminiModels(apiKey string) ([]string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	var models []string
	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if err != nil {
			break
		}
		// Filtrar apenas modelos que suportam geração de conteúdo
		supportsGenerate := false
		for _, op := range m.SupportedGenerationMethods {
			if op == "generateContent" {
				supportsGenerate = true
				break
			}
		}

		if supportsGenerate {
			name := strings.TrimPrefix(m.Name, "models/")

			// Filtrar apenas modelos modernos que começam com "gemini-"
			// e evitar modelos de treinamento ou versões "vision" isoladas que não são para chat direto.
			isGemini := strings.HasPrefix(strings.ToLower(name), "gemini-")
			isTuning := strings.Contains(strings.ToLower(name), "tuned")

			if isGemini && !isTuning {
				models = append(models, name)
			}
		}
	}

	return models, nil
}
