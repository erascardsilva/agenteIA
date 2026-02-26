// Erasmo Cardoso - Dev
package system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func ListFiles(path string) (string, error) {
	if path == "" {
		path = "."
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var result string
	for _, f := range files {
		typeStr := "file"
		if f.IsDir() {
			typeStr = "dir"
		}
		result += fmt.Sprintf("[%s] %s\n", typeStr, f.Name())
	}
	return result, nil
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(path, content string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func AppendToFile(path, content string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func ExecuteCommand(command string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		// Detectar se está rodando dentro de um Flatpak
		if _, err := os.Stat("/.flatpak-info"); err == nil {
			// Se estiver no Flatpak, usa flatpak-spawn --host para rodar binários do sistema
			// Usamos sh -c de forma direta para garantir que o comando chegue completo ao host
			cmd = exec.Command("flatpak-spawn", "--host", "sh", "-c", command)
		} else {
			cmd = exec.Command("sh", "-c", command)
		}
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("error executing command: %v", err)
	}
	return string(out), nil
}

func OpenEmail(to, subject, body string) error {
	url := fmt.Sprintf("mailto:%s?subject=%s&body=%s", to, subject, body)
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("xdg-open", url)
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", url)
	} else {
		cmd = exec.Command("cmd", "/C", "start", url)
	}
	return cmd.Run()
}
