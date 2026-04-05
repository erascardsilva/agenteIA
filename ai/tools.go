// Erasmo Cardoso - Software Engineer | Electronics Specialist
package ai

type Tool struct {
	Type     string       `json:"type"`
	Function ToolFunction `json:"function"`
}

type ToolFunction struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  interface{} `json:"parameters"`
}

func GetAvailableTools() []Tool {
	return []Tool{
		{
			Type: "function",
			Function: ToolFunction{
				Name:        "list_files",
				Description: "Lista os arquivos em um diretório específico.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "Caminho do diretório (vazio para o diretório atual)",
						},
					},
					"required": []string{"path"},
				},
			},
		},
		{
			Type: "function",
			Function: ToolFunction{
				Name:        "read_file",
				Description: "Lê o conteúdo de um arquivo.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "Caminho do arquivo",
						},
					},
					"required": []string{"path"},
				},
			},
		},
		{
			Type: "function",
			Function: ToolFunction{
				Name:        "write_file",
				Description: "Escreve conteúdo em um arquivo (sobrescreve se já existir).",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "Caminho do arquivo",
						},
						"content": map[string]interface{}{
							"type":        "string",
							"description": "Conteúdo a ser escrito",
						},
					},
					"required": []string{"path", "content"},
				},
			},
		},
		{
			Type: "function",
			Function: ToolFunction{
				Name:        "append_to_file",
				Description: "Concatena conteúdo ao final de um arquivo já existente (ou cria um novo). Útil para escrever arquivos grandes em blocos.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "Caminho do arquivo",
						},
						"content": map[string]interface{}{
							"type":        "string",
							"description": "Conteúdo a ser concatenado",
						},
					},
					"required": []string{"path", "content"},
				},
			},
		},
		{
			Type: "function",
			Function: ToolFunction{
				Name:        "execute_command",
				Description: "Executa um comando de terminal (Bash/Shell). Use com cautela.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"command": map[string]interface{}{
							"type":        "string",
							"description": "O comando a ser executado",
						},
					},
					"required": []string{"command"},
				},
			},
		},
		{
			Type: "function",
			Function: ToolFunction{
				Name:        "open_email",
				Description: "Abre o cliente de e-mail padrão para enviar uma mensagem.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"to": map[string]interface{}{
							"type":        "string",
							"description": "Destinatário do e-mail",
						},
						"subject": map[string]interface{}{
							"type":        "string",
							"description": "Assunto do e-mail",
						},
						"body": map[string]interface{}{
							"type":        "string",
							"description": "Corpo do e-mail",
						},
					},
					"required": []string{"to", "subject", "body"},
				},
			},
		},
	}
}

type ToolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}
