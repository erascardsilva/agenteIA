<!-- Erasmo Cardoso - Dev -->
<script>
    import { onMount, tick } from "svelte";
    import { marked } from "marked";
    import {
        SendMessage,
        GetChatHistory,
        ClearChatHistory,
        TextToSpeech,
        SpeakLinux,
        StopLinux,
    } from "../../wailsjs/go/main/App";
    import * as runtime from "../../wailsjs/runtime/runtime";
    import { configStore } from "./store";

    marked.setOptions({
        breaks: true,
        gfm: true,
    });

    let messages = [];
    let newMessage = "";
    let chatContainer;
    let loading = false;
    let thinkingMessage = "Pensando...";

    async function clear() {
        if (!confirm("Tem certeza que deseja apagar todo o histórico de chat?"))
            return;
        try {
            await ClearChatHistory();
            messages = [];
        } catch (e) {
            alert("Erro ao limpar histórico: " + e);
        }
    }

    async function send() {
        if (!newMessage.trim() || loading) return;

        const userMsg = newMessage;
        newMessage = "";
        messages = [...messages, { role: "user", content: userMsg }];

        loading = true;
        await scrollToBottom();

        try {
            const response = await SendMessage(userMsg);
            console.log("Frontend received response:", response);
            messages = [...messages, { role: "assistant", content: response }];

            // Auto-play se habilitado
            if ($configStore.voiceSettings.enabled) {
                speak(response);
            }
        } catch (e) {
            messages = [
                ...messages,
                { role: "assistant", content: "Erro: " + e },
            ];
        } finally {
            loading = false;
            await scrollToBottom();
        }
    }

    let currentPlayingMsg = null;
    let nativeUtterance = null;
    let audioObject = null;

    async function stopAudio() {
        if ($configStore.voiceSettings.engine === "native") {
            window.speechSynthesis.cancel();
        } else if ($configStore.voiceSettings.engine === "linux_local") {
            StopLinux();
        } else if (audioObject) {
            audioObject.pause();
            audioObject.currentTime = 0;
        }
        currentPlayingMsg = null;
    }

    async function speak(text, msgId) {
        try {
            if (!text) return;

            // Se já está tocando ESTA mensagem, para tudo
            if (currentPlayingMsg === msgId) {
                stopAudio();
                return;
            }

            // Se está tocando outra coisa, para antes de começar
            if (currentPlayingMsg !== null) {
                stopAudio();
            }

            currentPlayingMsg = msgId;

            // Se for motor nativo, usamos a Web Speech API
            if ($configStore.voiceSettings.engine === "native") {
                const utterance = new SpeechSynthesisUtterance(text);
                utterance.lang = "pt-BR";
                utterance.rate = 1.0;

                utterance.onend = () => {
                    if (currentPlayingMsg === msgId) currentPlayingMsg = null;
                };

                const speakNow = () => {
                    const voices = window.speechSynthesis.getVoices();
                    const ptVoice =
                        voices.find((v) => v.lang.includes("pt-BR")) ||
                        voices.find((v) => v.lang.includes("pt"));
                    if (ptVoice) utterance.voice = ptVoice;
                    window.speechSynthesis.speak(utterance);
                };

                if (window.speechSynthesis.getVoices().length === 0) {
                    window.speechSynthesis.onvoiceschanged = speakNow;
                } else {
                    speakNow();
                }
                return;
            }

            // Se for motor Linux (Local), usamos o backend Go -> spd-say
            if ($configStore.voiceSettings.engine === "linux_local") {
                const cleanText = text.replace(/[*#`_]/g, "");
                await SpeakLinux(
                    cleanText,
                    $configStore.voiceSettings.voiceId || "",
                );
                // Como spd-say é assíncrono no backend, resetamos o ícone depois de um tempo estimado ou imediatamente
                // Por simplicidade, resetamos após 3 segundos ou deixamos o usuário parar manual se implementarmos o stop
                setTimeout(() => {
                    if (currentPlayingMsg === msgId) currentPlayingMsg = null;
                }, 3000); // Estimativa básica
                return;
            }

            const cleanText = text.replace(/[*#`_]/g, "");

            const base64Audio = await TextToSpeech(cleanText);
            if (!base64Audio) {
                alert("O sistema de áudio não retornou dados.");
                currentPlayingMsg = null;
                return;
            }
            audioObject = new Audio("data:audio/mp3;base64," + base64Audio);
            audioObject.onended = () => {
                if (currentPlayingMsg === msgId) currentPlayingMsg = null;
            };
            audioObject.oncanplaythrough = () => audioObject.play();
            audioObject.onerror = (e) => {
                alert("Erro no player de áudio: " + e);
                currentPlayingMsg = null;
            };
        } catch (e) {
            console.error("Erro ao reproduzir voz:", e);
            alert("Erro na Voz do Jarvis: " + e);
            currentPlayingMsg = null;
        }
    }

    async function scrollToBottom() {
        await tick();
        if (chatContainer) {
            chatContainer.scrollTop = chatContainer.scrollHeight;
        }
    }

    onMount(async () => {
        // Escutar status de pensamento do backend
        runtime.EventsOn("thinking_status", (status) => {
            thinkingMessage = status;
        });

        const hist = await GetChatHistory(50);
        if (hist) {
            // Inverter para ordem cronológica (DB retorna DESC)
            messages = hist.reverse();
        }
        scrollToBottom();
    });
</script>

<div class="chat-wrapper">
    <div class="chat-header">
        <button
            class="clear-btn glass"
            on:click={clear}
            title="Apagar Histórico"
        >
            <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor">
                <path
                    d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"
                ></path>
            </svg>
            Apagar Chat
        </button>
    </div>
    <div class="messages" bind:this={chatContainer}>
        {#each messages.filter((m) => m.role !== "tool") as msg, i}
            <div class="message-row {msg.role}">
                <div class="message-bubble glass fade-in">
                    {@html marked.parse(
                        msg.content ||
                            "_Jarvis não retornou texto para esta mensagem._",
                    )}
                    {#if msg.role === "assistant"}
                        <button
                            class="speak-btn {currentPlayingMsg === i
                                ? 'playing'
                                : ''}"
                            on:click={() => speak(msg.content, i)}
                            title={currentPlayingMsg === i ? "Parar" : "Ouvir"}
                        >
                            {#if currentPlayingMsg === i}
                                <!-- Ícone Stop (Quadrado) -->
                                <svg
                                    viewBox="0 0 24 24"
                                    width="14"
                                    height="14"
                                    fill="currentColor"
                                >
                                    <rect x="6" y="6" width="12" height="12"
                                    ></rect>
                                </svg>
                            {:else}
                                <!-- Ícone Play (Triângulo) -->
                                <svg
                                    viewBox="0 0 24 24"
                                    width="14"
                                    height="14"
                                    fill="currentColor"
                                >
                                    <path d="M8 5v14l11-7z"></path>
                                </svg>
                            {/if}
                        </button>
                    {/if}
                </div>
            </div>
        {/each}
        {#if loading}
            <div class="message-row assistant">
                <div class="message-bubble glass loading">
                    <span class="thinking-text">{thinkingMessage}</span>
                    <span class="cursor">_</span>
                </div>
            </div>
        {/if}
    </div>

    <div class="input-area glass">
        <textarea
            placeholder="Digite sua mensagem..."
            bind:value={newMessage}
            on:keydown={(e) => {
                if (e.key === "Enter" && !e.shiftKey) {
                    e.preventDefault();
                    send();
                }
            }}
        ></textarea>
        <button
            class="send-btn"
            on:click={() => send()}
            disabled={loading || !newMessage.trim()}
        >
            <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
                <path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z"></path>
            </svg>
        </button>
    </div>
</div>

<style>
    .chat-wrapper {
        display: flex;
        flex-direction: column;
        height: 100%;
        padding: 20px;
        gap: 20px;
        flex: 1;
    }

    .messages {
        flex: 1;
        overflow-y: auto;
        display: flex;
        flex-direction: column;
        gap: 16px;
        padding-right: 10px;
    }

    .messages::-webkit-scrollbar {
        width: 6px;
    }

    .messages::-webkit-scrollbar-thumb {
        background: var(--border-color);
        border-radius: 10px;
    }

    .message-row {
        display: flex;
        width: 100%;
    }

    .message-row.user {
        justify-content: flex-end;
    }
    .message-row.assistant {
        justify-content: flex-start;
    }

    .message-bubble {
        max-width: 80%;
        padding: 12px 18px;
        border-radius: 18px;
        font-size: 15px;
        line-height: 1.5;
    }

    .user .message-bubble {
        background: var(--primary);
        color: white;
        border-bottom-right-radius: 4px;
        border: none;
    }

    .assistant .message-bubble {
        border-bottom-left-radius: 4px;
        background: var(--sidebar-bg);
        position: relative;
        padding-bottom: 30px; /* Espaço para o botão */
        color: white; /* Garante visibilidade */
    }

    .speak-btn {
        position: absolute;
        right: 8px;
        bottom: 6px;
        background: rgba(255, 255, 255, 0.05);
        border: none;
        border-radius: 50%;
        width: 24px;
        height: 24px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--text-muted);
        cursor: pointer;
        transition: all 0.2s;
    }

    .speak-btn:hover {
        color: var(--primary);
        background: rgba(255, 255, 255, 0.1);
        transform: scale(1.1);
    }

    .speak-btn.playing {
        color: #ff4444;
        background: rgba(255, 68, 68, 0.1);
        animation: pulse-red 1.5s infinite;
    }

    @keyframes pulse-red {
        0% {
            transform: scale(1);
        }
        50% {
            transform: scale(1.2);
        }
        100% {
            transform: scale(1);
        }
    }

    /* Estilos para Markdown dentro das mensagens */
    .message-bubble :global(p) {
        margin: 0 0 12px 0;
    }
    .message-bubble :global(p:last-child) {
        margin-bottom: 0;
    }
    .message-bubble :global(ul),
    .message-bubble :global(ol) {
        margin: 8px 0;
        padding-left: 20px;
    }
    .message-bubble :global(li) {
        margin-bottom: 4px;
    }

    .input-area {
        display: flex;
        padding: 12px;
        gap: 12px;
        border-radius: 16px;
        align-items: flex-end;
    }

    textarea {
        flex: 1;
        background: none;
        border: none;
        resize: none;
        min-height: 48px;
        max-height: 200px;
        padding: 8px;
        color: white;
        font-size: 15px;
    }

    .send-btn {
        background: var(--primary);
        color: white;
        width: 48px;
        height: 48px;
        border-radius: 12px;
        flex-shrink: 0;
    }

    .send-btn:disabled {
        background: var(--text-muted);
        opacity: 0.5;
    }

    .chat-header {
        display: flex;
        justify-content: flex-end;
        padding-bottom: 10px;
    }

    .clear-btn {
        display: flex;
        align-items: center;
        gap: 8px;
        padding: 8px 16px;
        font-size: 13px;
        color: var(--text-muted);
        border-radius: 10px;
        transition: all 0.2s;
    }

    .clear-btn:hover {
        background: rgba(255, 0, 0, 0.1);
        color: #ff4444;
    }

    .loading {
        display: flex;
        align-items: center;
        gap: 4px;
        min-width: 100px;
    }

    .thinking-text {
        font-family: "Courier New", Courier, monospace;
        font-size: 14px;
        color: var(--primary);
    }

    .cursor {
        animation: blink 1s infinite;
        font-weight: bold;
        color: var(--primary);
    }

    @keyframes blink {
        0%,
        100% {
            opacity: 1;
        }
        50% {
            opacity: 0;
        }
    }
</style>
