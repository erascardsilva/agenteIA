// Erasmo Cardoso - Software Engineer | Electronics Specialist
package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	appDir := filepath.Join(configDir, "agente-ia")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		_ = os.MkdirAll(appDir, 0755)
	}

	dbPath := filepath.Join(appDir, "data.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	DB = db

	// Otimizações para evitar "database is locked"
	_, _ = DB.Exec("PRAGMA journal_mode=WAL;")
	_, _ = DB.Exec("PRAGMA synchronous=NORMAL;")
	_, _ = DB.Exec("PRAGMA busy_timeout=5000;")

	err = createTables()
	return err
}

func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS configs (
			key TEXT PRIMARY KEY,
			value TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS chat_history (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			role TEXT,
			content TEXT,
			tool_call_id TEXT,
			tool_calls TEXT,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			return err
		}
	}

	// Migração: Adicionar colunas se não existirem
	DB.Exec("ALTER TABLE chat_history ADD COLUMN tool_call_id TEXT")
	DB.Exec("ALTER TABLE chat_history ADD COLUMN tool_calls TEXT")

	return nil
}

func SetConfig(key, value string) error {
	_, err := DB.Exec("INSERT OR REPLACE INTO configs (key, value) VALUES (?, ?)", key, value)
	return err
}

func GetConfig(key string) (string, error) {
	var value string
	err := DB.QueryRow("SELECT value FROM configs WHERE key = ?", key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return value, err
}

func SaveMessage(role, content, toolCallID, toolCalls string) error {
	if content == "" && toolCalls == "" {
		return nil // Não salva mensagens totalmente vazias
	}
	_, err := DB.Exec("INSERT INTO chat_history (role, content, tool_call_id, tool_calls) VALUES (?, ?, ?, ?)", role, content, toolCallID, toolCalls)
	return err
}

func GetHistory(limit int) ([]map[string]string, error) {
	rows, err := DB.Query("SELECT role, content, COALESCE(tool_call_id, ''), COALESCE(tool_calls, '') FROM chat_history ORDER BY timestamp DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []map[string]string
	for rows.Next() {
		var role, content, tcid, tcs string
		if err := rows.Scan(&role, &content, &tcid, &tcs); err != nil {
			return nil, err
		}
		history = append(history, map[string]string{
			"role":         role,
			"content":      content,
			"tool_call_id": tcid,
			"tool_calls":   tcs,
		})
	}
	return history, nil
}
func ClearHistory() error {
	_, err := DB.Exec("DELETE FROM chat_history")
	return err
}
