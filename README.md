# AI Agent - Multimodal Desktop Assistant

[![Disponível na Snap Store](https://snapcraft.io/pt/dark/install.svg)](https://snapcraft.io/agenteia)

AI Agent is a desktop assistant developed in Go and Wails for integration with various language model providers (LLMs). The goal is to provide a fast and local interface for productivity.

---

## Technologies

The project utilizes the following components:

*   **Backend:** [Go](https://golang.org/) + [Wails v2](https://wails.io/)
*   **Frontend:** [Svelte](https://svelte.dev/) + [Vite](https://vitejs.dev/)
*   **Supported Providers:**
    *   **Groq:** Llama and Mixtral models (focus on speed).
    *   **Google Gemini:** Flash and Pro models from the 1.5/2.0 series.
    *   **OpenAI:** Support for GPT-4o and legacy models.
    *   **DeepSeek:** OpenAI-compatible API.
    *   **OpenRouter:** Unified access to multiple models.
    *   **Ollama:** Support for local models via Ollama server.
*   **Audio & Voice:**
    *   Support for **OpenAI TTS** for cloud synthesis.
    *   Native integration via **Web Speech API**.
    *   Support for **spd-say** on Linux for local system voices.
*   **Persistence:** SQLite for chat history and local settings.

---

## Architecture

The application follows a hybrid IPC (Inter-Process Communication) architecture.

```mermaid
graph TD
    User((User)) <--> Frontend[Frontend - Svelte/Chat]
    Frontend <--> Wails[Wails Bridge - IPC]
    Wails <--> Go[Go Backend - Business Logic]
    Go <--> FS[File System]
    Go <--> AI[Cloud/Local LLM APIs]
    Go <--> Audio[Native/Cloud Audio]
    Go <--> DB[(SQLite - Local Data)]
```

---

## Configuration

To use the assistant, you need to configure your API keys:

1.  Open AI Agent.
2.  Go to the **Settings** menu (gear icon).
3.  Enter the keys for the desired providers (e.g., Gemini, Groq, OpenAI).
4.  Click **Save Settings**.

---

## Build and Compilation

### Requirements
*   Go 1.21+
*   Node.js (LTS recommended)
*   Wails CLI

### Generate binary for your system
```bash
# Build for Linux
wails build -platform linux/amd64

# Build for Windows
wails build -platform windows/amd64 -nsis
```

### Linux (Flatpak)
The project includes a script for automated Flatpak builds:
```bash
./build-flatpak.sh
```

### Linux (Snap)
O projeto está disponível na Snap Store:
[![Disponível na Snap Store](https://snapcraft.io/pt/dark/install.svg)](https://snapcraft.io/agenteia)

Ou instale via terminal:
```bash
sudo snap install agenteia
```

---

**Erasmo Cardoso - Dev**
