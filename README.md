<div align="center">

# AI Agent - Multimodal Desktop Assistant
### Fast and local interface for multimodal AI productivity

![Go 1.25+](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Wails v2](https://img.shields.io/badge/Wails-v2-E13238?style=for-the-badge&logo=wails&logoColor=white)
![Svelte](https://img.shields.io/badge/Svelte-%23FF3E00.svg?style=for-the-badge&logo=svelte&logoColor=white)
![Linux ready](https://img.shields.io/badge/Linux-ready-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![Windows ready](https://img.shields.io/badge/Windows-ready-0078D4?style=for-the-badge&logo=windows&logoColor=white)

<a href="https://snapcraft.io/agenteia">
  <img src="https://snapcraft.io/pt/dark/install.svg" alt="Available on the Snap Store" height="60">
</a>

</div>

AI Agent is a cross-platform desktop assistant compatible with all Linux distributions and Windows. Developed in Go and Wails, it integrates with various language model providers (LLMs) to provide a fast and local interface for productivity.

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
*   **Compatibility:**
    *   **Linux:** All distributions (Ubuntu, Fedora, Arch, Debian, etc.) via Snap or native build.
    *   **Windows:** Windows 10 and 11.

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

Binary files and installers (Windows .exe and Linux) are already available in the repository, located in the [build/bin](./build/bin) directory.

### Generate binary for your system
To generate the files manually, use the commands below. After execution, the files will be updated in the `build/bin` directory.

```bash
# Build for Linux
wails build -platform linux/amd64

# Build for Windows
wails build -platform windows/amd64 -nsis
```

### Linux (Snap)
The store version (Snap) has security restrictions (Sandbox) that may limit Autonomous Mode.
[![Available on the Snap Store](https://snapcraft.io/pt/dark/install.svg)](https://snapcraft.io/agenteia)

---

## 🚀 Full Version (Direct Download)

To utilize the full power of the **Autonomous Agent** (command execution, file and script management), we recommend using the full version available in the [build/bin](./build/bin) folder.

### 🪟 Windows
1.  Go to [build/bin](./build/bin).
2.  Download the `agente-ia-installer.exe` file.
3.  Run the installer and follow the steps.

### 🐧 Debian / Ubuntu / Mint (.deb)
1.  Download the `agenteia_amd64.deb` file from [build/bin](./build/bin).
2.  Install via terminal:
    ```bash
    sudo dpkg -i agenteia_amd64.deb
    ```
3.  In case of missing dependencies: `sudo apt install -f`.

### 🎩 Fedora / Red Hat / Arch (Binary)
1.  Download the `agenteIA` binary from [build/bin](./build/bin).
2.  Give execution permission:
    ```bash
    chmod +x agenteIA
    ```
3.  Run directly:
    ```bash
    ./agenteIA
    ```

---

**Erasmo Cardoso**<br>
**Software Engineer |Electronics Technician**

---

<div align="center">

### Support the Development
If this project helped you, consider making a donation to support continuous development.

[![Apoie com PayPal](https://img.shields.io/badge/Donate-PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/ncp/payment/8V6WQCGN6HDCQ)

</div>
