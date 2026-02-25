# 🤖 AI Agent - Multimodal Desktop Assistant

AI Agent is a powerful and flexible artificial intelligence ecosystem designed to interact directly with your operating system. It combines the power of modern LLMs with a rich, fast desktop interface.

---

## 🛠️ Technologies Used

This project was built using the following cutting-edge technologies:

*   **Backend:** [Go (Golang)](https://golang.org/) + [Wails](https://wails.io/) (Native interface)
*   **Frontend:** [Svelte](https://svelte.dev/) + [Vite](https://vitejs.dev/)
*   **Artificial Intelligence (LLMs):**
    *   **Groq:** Ultra-fast models (Llama 3.3, Mixtral).
    *   **Google Gemini:** Advanced vision and reasoning (Flash/Pro).
    *   **OpenAI:** Market standard (GPT-4o).
    *   **DeepSeek:** Reasoning (V3/R1).
    *   **OpenRouter:** Universal access to hundreds of models.
    *   **Ollama:** Support for locally running models.
*   **🔓 AI Unlock (Unlock Models):** Option to disable restrictive LLM filters, allowing for more direct, creative, and uncensored responses.
*   **Audio & Voice:**
    *   **OpenAI TTS:** High-quality neural voices.
    *   **Web Speech API:** Native system voice synthesis.
*   **Database:** SQLite (Contexts, History, and Settings).

---

## 🏗️ Engineering and Operation

The app uses a hybrid architecture where the Frontend (Svelte) communicates asynchronously with the Backend (Go) via the Wails Bridge.

```mermaid
graph TD
    User((User)) <--> Frontend[Frontend - Svelte/Chat]
    Frontend <--> Wails[Wails Bridge - IPC]
    Wails <--> Go[Go Backend - App Logic]
    Go <--> FS[File System]
    Go <--> AI[Cloud/Local LLM APIs]
    Go <--> Audio[OpenAI/Native Audio]
    Go <--> DB[(SQLite - Configs/Context)]
```

---

## 🔑 How to configure your API Keys

To unlock the full power of AI, you need to add your keys:

1.  Open AI Agent.
2.  Click the **⚙️ Settings** icon in the sidebar.
3.  Enter the keys for the providers you wish to use:
    *   **Groq:** Get it at [console.groq.com](https://console.groq.com/keys).
    *   **Gemini:** Get it at [Google AI Studio](https://aistudio.google.com/app/apikey).
    *   **OpenAI:** Get it at the [OpenAI dashboard](https://platform.openai.com/api-keys).
4.  Click **Save Settings**.

---

## 🔊 Audio Differences by System

AI Agent supports two audio modes: **OpenAI Premium (Cloud)** and **Native (Local)**.

*   **Windows:** Native audio uses Microsoft SAPI/Edge voices, which are very natural and fast.
*   **macOS:** Uses premium Apple voices (like Siri and Alex), providing an excellent experience with no delay.
*   **Linux:** Native audio depends on `speech-dispatcher`. It may sound more robotic depending on the installed distribution. **Tip:** On Linux, use the **OpenAI Premium** engine for perfect human voices.

---

## 🚀 How to Install or Build

### Requirements
*   Go 1.21+
*   Node.js & NPM
*   Wails CLI

### Generate Build for your system
```bash
# Install Wails if you don't have it
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Build for Linux
wails build -platform linux/amd64

# Build for Windows (requires Mingw on Linux for cross-compile)
wails build -platform windows/amd64
```

---

## 📦 Executables

The compiled executables are available in the `build/bin/` folder.
*   `agenteIA.exe` (Windows)
*   `agenteIA` (Linux)

---

**Erasmo Cardoso - Dev**
