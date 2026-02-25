<!-- Erasmo Cardoso - Dev -->
<script>
    import { configStore } from "./store";
    import { onMount } from "svelte";

    export let visible = false;
    let currentTab = "config"; // 'config' ou 'about'

    // Erasmo Cardoso - Dev
    // Lógica para dropdowns customizados
    let openDropdown = null; // ID do dropdown aberto atualmente

    const labelMaps = {
        "voiceSettings.engine": {
            kokoro: "Kokoro-82M (Local/Grátis)",
            openai: "OpenAI Premium (API Key)",
            native: "Nativo (Navegador)",
        },
        "voiceSettings.voiceId": {
            pf_dora: "Feminina PT [Offline]",
            pm_alex: "Masculina PT [Offline]",
            pm_santa: "Masculina PT (Santa) [Offline]",
        },
        preferredProvider: {
            groq: "Groq (Llama 3 / Mixtral)",
            gemini: "Google Gemini (Flash/Pro)",
            openai: "OpenAI (GPT-4o)",
            deepseek: "DeepSeek (V3/R1)",
            openrouter: "OpenRouter (Universal)",
            ollama: "Ollama (Local)",
        },
    };

    function toggleDropdown(id, e) {
        if (e) e.stopPropagation();
        openDropdown = openDropdown === id ? null : id;
    }

    function selectOption(path, value) {
        configStore.updateField(path, value);
        openDropdown = null;
    }

    onMount(() => {
        const handleOutsideClick = () => (openDropdown = null);
        window.addEventListener("click", handleOutsideClick);
        return () => window.removeEventListener("click", handleOutsideClick);
    });
    async function save() {
        await configStore.save($configStore);
        visible = false;
    }

    function handleKeydown(e) {
        if (e.key === "Escape") visible = false;
    }
</script>

{#if visible}
    <div
        class="modal-overlay"
        on:click|self={() => (visible = false)}
        on:keydown={handleKeydown}
        role="button"
        tabindex="0"
    >
        <div class="settings-modal glass fade-in">
            <header>
                <div class="tabs">
                    <button
                        class="tab-link {currentTab === 'config'
                            ? 'active'
                            : ''}"
                        on:click={() => (currentTab = "config")}
                        >Configurações</button
                    >
                    <button
                        class="tab-link {currentTab === 'about'
                            ? 'active'
                            : ''}"
                        on:click={() => (currentTab = "about")}>Sobre</button
                    >
                </div>
                <button class="close-btn" on:click={() => (visible = false)}
                    >&times;</button
                >
            </header>

            {#if currentTab === "config"}
                <section class="settings-content">
                    <div class="field">
                        <label for="assistantName">Nome do Assistente</label>
                        <input
                            id="assistantName"
                            type="text"
                            bind:value={$configStore.assistantName}
                            placeholder="Ex: Jarvis, Sexta-Feira..."
                        />
                    </div>

                    <div class="field">
                        <label for="apiKeyGroq">API Key (Groq - Grátis)</label>
                        <input
                            id="apiKeyGroq"
                            type="password"
                            bind:value={$configStore.apiKeys["groq"]}
                            placeholder="gsk_..."
                        />
                    </div>

                    <div class="field">
                        <label for="apiKeyGemini"
                            >API Key (Gemini - Grátis)</label
                        >
                        <input
                            id="apiKeyGemini"
                            type="password"
                            bind:value={$configStore.apiKeys["gemini"]}
                            placeholder="AIza..."
                        />
                    </div>

                    <div class="field">
                        <label for="apiKeyOpenAI">API Key (OpenAI)</label>
                        <input
                            id="apiKeyOpenAI"
                            type="password"
                            bind:value={$configStore.apiKeys["openai"]}
                            placeholder="sk-..."
                        />
                    </div>

                    <div class="field">
                        <label for="apiKeyDeepSeek">API Key (DeepSeek)</label>
                        <input
                            id="apiKeyDeepSeek"
                            type="password"
                            bind:value={$configStore.apiKeys["deepseek"]}
                            placeholder="sk-..."
                        />
                    </div>

                    <div class="field">
                        <label for="apiKeyOpenRouter"
                            >API Key (OpenRouter)</label
                        >
                        <input
                            id="apiKeyOpenRouter"
                            type="password"
                            bind:value={$configStore.apiKeys["openrouter"]}
                            placeholder="sk-or-v1-..."
                        />
                    </div>

                    <div class="field">
                        <label for="humor">Humor</label>
                        <div class="custom-select" id="humor">
                            <button
                                class="select-trigger"
                                on:click={(e) => toggleDropdown("humor", e)}
                            >
                                <span>{$configStore.context.humor}</span>
                                <i
                                    class="chevron"
                                    class:open={openDropdown === "humor"}
                                ></i>
                            </button>
                            {#if openDropdown === "humor"}
                                <div class="options-container">
                                    {#each ["Prestativo", "Sarcástico", "Profissional", "Amigável", "Direto", "Zombeteiro"] as hum}
                                        <button
                                            class="option"
                                            class:selected={$configStore.context
                                                .humor === hum}
                                            on:click={() =>
                                                selectOption(
                                                    "context.humor",
                                                    hum,
                                                )}
                                        >
                                            <span class="dot"></span>
                                            {hum}
                                        </button>
                                    {/each}
                                </div>
                            {/if}
                        </div>
                    </div>

                    <div class="field">
                        <label for="provider">Provedor Preferencial</label>
                        <div class="custom-select" id="provider">
                            <button
                                class="select-trigger"
                                on:click={(e) => toggleDropdown("provider", e)}
                            >
                                <span
                                    >{labelMaps["preferredProvider"][
                                        $configStore.preferredProvider
                                    ]}</span
                                >
                                <i
                                    class="chevron"
                                    class:open={openDropdown === "provider"}
                                ></i>
                            </button>
                            {#if openDropdown === "provider"}
                                <div class="options-container">
                                    {#each Object.entries(labelMaps["preferredProvider"]) as [val, label]}
                                        <button
                                            class="option"
                                            class:selected={$configStore.preferredProvider ===
                                                val}
                                            on:click={() => {
                                                selectOption(
                                                    "preferredProvider",
                                                    val,
                                                );
                                                $configStore.preferredModel =
                                                    "";
                                            }}
                                        >
                                            <span class="dot"></span>
                                            {label}
                                        </button>
                                    {/each}
                                </div>
                            {/if}
                        </div>
                    </div>

                    <div class="field">
                        <label for="model">Modelo</label>
                        {#if $configStore.preferredProvider === "groq"}
                            <select
                                id="model"
                                bind:value={$configStore.preferredModel}
                            >
                                <option value="llama-3.3-70b-versatile"
                                    >Llama 3.3 70B</option
                                >
                                <option value="llama-3.1-8b-instant"
                                    >Llama 3.1 8B (Instant)</option
                                >
                                <option value="mixtral-8x7b-32768"
                                    >Mixtral 8x7B</option
                                >
                            </select>
                        {:else if $configStore.preferredProvider === "gemini"}
                            <select
                                id="model"
                                bind:value={$configStore.preferredModel}
                            >
                                <option value="gemini-3.1-pro-preview"
                                    >Gemini 3.1 Pro (Estado da Arte)</option
                                >
                                <option value="gemini-3-flash-preview"
                                    >Gemini 3 Flash (Alta Velocidade)</option
                                >
                                <option value="gemini-3-pro-preview"
                                    >Gemini 3 Pro Preview</option
                                >
                                <option value="gemini-2.5-pro-preview"
                                    >Gemini 2.5 Pro (Estável/Complexo)</option
                                >
                                <option value="gemini-2.5-flash-preview"
                                    >Gemini 2.5 Flash (Estável/Rápido)</option
                                >
                            </select>
                        {:else if $configStore.preferredProvider === "openai"}
                            <select
                                id="model"
                                bind:value={$configStore.preferredModel}
                            >
                                <option value="gpt-4o">GPT-4o (Padrão)</option>
                                <option value="gpt-4o-mini">GPT-4o Mini</option>
                                <option value="gpt-3.5-turbo"
                                    >GPT-3.5 Turbo</option
                                >
                            </select>
                        {:else if $configStore.preferredProvider === "deepseek"}
                            <select
                                id="model"
                                bind:value={$configStore.preferredModel}
                            >
                                <option value="deepseek-chat"
                                    >DeepSeek V3 (Chat)</option
                                >
                                <option value="deepseek-reasoner"
                                    >DeepSeek R1 (Raciocínio)</option
                                >
                            </select>
                        {:else if $configStore.preferredProvider === "openrouter"}
                            <select
                                id="model"
                                bind:value={$configStore.preferredModel}
                            >
                                <option value="google/gemini-2.0-flash-exp:free"
                                    >Gemini 2.0 Flash (Free)</option
                                >
                                <option
                                    value="mistralai/mistral-7b-instruct:free"
                                    >Mistral 7B (Free)</option
                                >
                                <option
                                    value="huggingfaceh4/zephyr-7b-beta:free"
                                    >Zephyr 7B (Free)</option
                                >
                                <option value="openrouter/auto"
                                    >Auto (Melhor grátis)</option
                                >
                            </select>
                        {:else if $configStore.preferredProvider === "ollama"}
                            <input
                                id="model"
                                type="text"
                                bind:value={$configStore.preferredModel}
                                placeholder="ex: llama3, mistral, phi3..."
                            />
                        {:else}
                            <input
                                id="model"
                                type="text"
                                bind:value={$configStore.preferredModel}
                                placeholder="Nome do modelo..."
                            />
                        {/if}
                    </div>

                    <div class="field">
                        <label for="voiceEngine">Motor de Voz</label>
                        <div class="custom-select" id="voiceEngine">
                            <button
                                class="select-trigger"
                                on:click={(e) => toggleDropdown("engine", e)}
                            >
                                <span
                                    >{labelMaps["voiceSettings.engine"][
                                        $configStore.voiceSettings.engine
                                    ]}</span
                                >
                                <i
                                    class="chevron"
                                    class:open={openDropdown === "engine"}
                                ></i>
                            </button>
                            {#if openDropdown === "engine"}
                                <div class="options-container">
                                    {#each Object.entries(labelMaps["voiceSettings.engine"]) as [val, label]}
                                        <button
                                            class="option"
                                            class:selected={$configStore
                                                .voiceSettings.engine === val}
                                            on:click={() =>
                                                selectOption(
                                                    "voiceSettings.engine",
                                                    val,
                                                )}
                                        >
                                            <span class="dot"></span>
                                            {label}
                                        </button>
                                    {/each}
                                </div>
                            {/if}
                        </div>
                    </div>

                    {#if $configStore.voiceSettings.engine === "kokoro"}
                        <div class="field">
                            <label for="voice">Voz do Assistente (Neural)</label
                            >
                            <div class="custom-select" id="voice">
                                <button
                                    class="select-trigger"
                                    on:click={(e) =>
                                        toggleDropdown("voiceId", e)}
                                >
                                    <span
                                        >{labelMaps["voiceSettings.voiceId"][
                                            $configStore.voiceSettings.voiceId
                                        ]}</span
                                    >
                                    <i
                                        class="chevron"
                                        class:open={openDropdown === "voiceId"}
                                    ></i>
                                </button>
                                {#if openDropdown === "voiceId"}
                                    <div class="options-container">
                                        {#each Object.entries(labelMaps["voiceSettings.voiceId"]) as [val, label]}
                                            <button
                                                class="option"
                                                class:selected={$configStore
                                                    .voiceSettings.voiceId ===
                                                    val}
                                                on:click={() =>
                                                    selectOption(
                                                        "voiceSettings.voiceId",
                                                        val,
                                                    )}
                                            >
                                                <span class="dot"></span>
                                                {label}
                                            </button>
                                        {/each}
                                    </div>
                                {/if}
                            </div>
                        </div>
                    {:else if $configStore.voiceSettings.engine === "openai"}
                        <div class="field">
                            <label for="voice">Voz do Assistente (OpenAI)</label
                            >
                            <select
                                id="voice"
                                bind:value={$configStore.voiceSettings.voiceId}
                            >
                                <option value="alloy">Alloy (Padrão)</option>
                                <option value="echo">Echo</option>
                                <option value="fable">Fable</option>
                                <option value="onyx">Onyx</option>
                                <option value="nova">Nova</option>
                                <option value="shimmer">Shimmer</option>
                            </select>
                        </div>
                    {/if}

                    <div class="field toggle-item">
                        <label for="voiceEnabled" class="toggle-label">
                            <span class="label-text"
                                >Ativar Voz (Auto-Play)</span
                            >
                            <div class="switch">
                                <input
                                    id="voiceEnabled"
                                    type="checkbox"
                                    bind:checked={
                                        $configStore.voiceSettings.enabled
                                    }
                                />
                                <span class="slider"></span>
                            </div>
                        </label>
                    </div>

                    <div class="field toggle-item">
                        <label for="unlockModels" class="toggle-label">
                            <span class="label-text"
                                >Desativar Filtros (Unlock Models)</span
                            >
                            <div class="switch">
                                <input
                                    id="unlockModels"
                                    type="checkbox"
                                    bind:checked={
                                        $configStore.context.unlockModels
                                    }
                                />
                                <span class="slider"></span>
                            </div>
                        </label>
                    </div>
                </section>
            {:else}
                <section class="settings-content about-content">
                    <div class="about-hero">
                        <h3>Jarvis AI</h3>
                        <p>Assistente Técnico Avançado</p>
                    </div>

                    <div class="about-section">
                        <h4>🚀 O que eu faço?</h4>
                        <p>
                            Sou um assistente multimodal projetado para
                            automação e suporte técnico. Posso processar texto,
                            visão e voz, além de interagir diretamente com o seu
                            sistema operacional.
                        </p>
                    </div>

                    <div class="about-section">
                        <h4>🛠️ Recursos Principais</h4>
                        <ul>
                            <li>
                                <strong>Multi-Modelos:</strong> Suporte nativo a
                                Gemini 3.1, Groq, DeepSeek e OpenRouter.
                            </li>
                            <li>
                                <strong>Interação de Sistema:</strong> Leitura/Escrita
                                de arquivos, execução de comandos e automação de
                                e-mails.
                            </li>
                            <li>
                                <strong>Voz Neural:</strong> Motor de voz dual com
                                OpenAI Premium ou Edge TTS gratuito de alta qualidade.
                            </li>
                            <li>
                                <strong>Privacidade:</strong> Processamento focado
                                em segurança com armazenamento local de histórico
                                e chaves.
                            </li>
                        </ul>
                    </div>

                    <div class="about-section">
                        <h4>🔑 Como adicionar API Keys?</h4>
                        <ol>
                            <li>
                                Vá na aba <strong>Configurações</strong> acima.
                            </li>
                            <li>
                                Obtenha sua chave no site do provedor (ex: <a
                                    href="https://console.groq.com/keys"
                                    target="_blank">Groq</a
                                >
                                ou
                                <a
                                    href="https://aistudio.google.com/app/apikey"
                                    target="_blank">Gemini</a
                                >).
                            </li>
                            <li>
                                Cole a chave no campo correspondente e clique em <strong
                                    >Salvar Configurações</strong
                                >.
                            </li>
                        </ol>
                    </div>

                    <div class="about-footer">
                        <div class="signature">Erasmo Cardoso - Dev</div>
                        <div class="version">v2.1.0 - Era 2026</div>
                    </div>
                </section>
            {/if}

            {#if currentTab === "config"}
                <footer>
                    <button class="save-btn" on:click={save}
                        >Salvar Configurações</button
                    >
                </footer>
            {/if}
        </div>
    </div>
{/if}

<style>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.6);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 100;
        backdrop-filter: blur(4px);
    }

    .settings-modal {
        width: 95%;
        max-width: 550px;
        max-height: 90vh;
        border-radius: 24px;
        padding: 0;
        display: flex;
        flex-direction: column;
        overflow: hidden;
        border: 1px solid rgba(255, 255, 255, 0.1);
        box-shadow: 0 20px 50px rgba(0, 0, 0, 0.5);
    }

    header {
        padding: 24px 24px 16px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .tabs {
        display: flex;
        gap: 16px;
    }

    .tab-link {
        background: none;
        border: none;
        color: var(--text-muted);
        font-size: 16px;
        font-weight: 600;
        padding: 8px 0;
        cursor: pointer;
        position: relative;
        transition: color 0.2s;
    }

    .tab-link:hover {
        color: white;
    }

    .tab-link.active {
        color: var(--primary);
    }

    .tab-link.active::after {
        content: "";
        position: absolute;
        bottom: 0;
        left: 0;
        width: 100%;
        height: 2px;
        background: var(--primary);
        border-radius: 2px;
    }

    .settings-content {
        padding: 24px;
        display: flex;
        flex-direction: column;
        gap: 20px;
        overflow-y: auto;
        scrollbar-width: thin;
        scrollbar-color: var(--primary) transparent;
    }

    /* Custom Scrollbar for Webkit */
    .settings-content::-webkit-scrollbar {
        width: 6px;
    }
    .settings-content::-webkit-scrollbar-thumb {
        background: var(--primary);
        border-radius: 10px;
    }

    .field {
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .field label {
        font-size: 13px;
        color: var(--text-muted);
        margin-bottom: 2px;
    }

    .field input,
    .field select {
        width: 100%;
        background: rgba(
            0,
            0,
            0,
            0.4
        ); /* Fundo mais escuro para maior contraste */
        border: 1px solid rgba(255, 255, 255, 0.2);
        border-radius: 12px;
        padding: 12px 16px;
        color: white;
        font-size: 15px;
        transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
        appearance: none;
        background-repeat: no-repeat;
        background-position: right 14px center;
        background-size: 18px;
        cursor: pointer;
    }

    /* Estilos para o Dropdown Premium Customizado */
    .custom-select {
        position: relative;
        width: 100%;
    }

    .select-trigger {
        width: 100%;
        background: rgba(0, 0, 0, 0.4);
        border: 1px solid rgba(255, 255, 255, 0.2);
        border-radius: 12px;
        padding: 12px 16px;
        color: white;
        font-size: 15px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        cursor: pointer;
        transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .select-trigger:hover {
        background: rgba(255, 255, 255, 0.08);
        border-color: rgba(88, 166, 255, 0.5);
    }

    .select-trigger:focus {
        border-color: var(--primary);
        background: rgba(0, 0, 0, 0.6);
        outline: none;
        box-shadow: 0 0 0 4px rgba(88, 166, 255, 0.2);
    }

    .chevron {
        width: 18px;
        height: 18px;
        background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='%2358a6ff' stroke-width='2.5' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
        background-size: contain;
        background-repeat: no-repeat;
        transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    }

    .chevron.open {
        transform: rotate(180deg);
    }

    .options-container {
        position: absolute;
        top: calc(100% + 8px);
        left: 0;
        right: 0;
        background: #1a1a1a;
        border: 1px solid rgba(255, 255, 255, 0.15);
        border-radius: 12px;
        box-shadow: 0 10px 25px rgba(0, 0, 0, 0.5);
        z-index: 1000;
        overflow: hidden;
        animation: slideDown 0.2s cubic-bezier(0, 0, 0.2, 1);
        max-height: 250px;
        overflow-y: auto;
    }

    @keyframes slideDown {
        from {
            opacity: 0;
            transform: translateY(-10px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .option {
        width: 100%;
        padding: 12px 16px;
        background: transparent;
        color: rgba(255, 255, 255, 0.8);
        font-size: 14px;
        display: flex;
        align-items: center;
        gap: 12px;
        cursor: pointer;
        transition: all 0.2s;
        text-align: left;
    }

    .option:hover {
        background: rgba(88, 166, 255, 0.15);
        color: white;
    }

    .option.selected {
        background: rgba(88, 166, 255, 0.25);
        color: white;
        font-weight: 600;
    }

    .dot {
        width: 8px;
        height: 8px;
        border-radius: 50%;
        background: transparent;
        transition: all 0.2s;
        flex-shrink: 0;
    }

    .option.selected .dot {
        background: white; /* O ponto branco solicitado pelo Erasmo */
        box-shadow: 0 0 8px rgba(255, 255, 255, 0.8);
    }

    .field input,
    .field select {
        width: 100%;
        background: rgba(
            0,
            0,
            0,
            0.4
        ); /* Fundo mais escuro para maior contraste */
        border: 1px solid rgba(255, 255, 255, 0.2);
        border-radius: 12px;
        padding: 12px 16px;
        color: white;
        font-size: 15px;
        transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
        appearance: none;
        background-repeat: no-repeat;
        background-position: right 14px center;
        background-size: 18px;
        cursor: pointer;
    }

    .field input {
        cursor: text;
    }

    .field select {
        background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none' stroke='%2358a6ff' stroke-width='2.5' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
    }

    .field select option {
        background: #121212;
        color: white;
        padding: 12px;
    }

    .field input:hover,
    .field select:hover {
        background: rgba(255, 255, 255, 0.08);
        border-color: rgba(88, 166, 255, 0.5);
    }

    .field input:focus,
    .field select:focus {
        border-color: var(--primary);
        background: rgba(0, 0, 0, 0.6);
        outline: none;
        box-shadow: 0 0 0 4px rgba(88, 166, 255, 0.2);
        transform: translateY(-1px);
    }

    /* Estilos para Switches (Pill Toggles) */
    .toggle-item {
        background: rgba(255, 255, 255, 0.03);
        padding: 16px 20px;
        border-radius: 16px;
        transition: background 0.2s;
        border: 1px solid rgba(255, 255, 255, 0.05);
    }

    .toggle-item:hover {
        background: rgba(255, 255, 255, 0.06);
    }

    .toggle-label {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        cursor: pointer;
        margin: 0;
    }

    .label-text {
        font-size: 15px;
        font-weight: 500;
        color: rgba(255, 255, 255, 0.9);
    }

    .switch {
        position: relative;
        display: inline-block;
        width: 50px;
        height: 26px;
        flex-shrink: 0;
    }

    .switch input {
        opacity: 0;
        width: 0;
        height: 0;
        position: absolute;
    }

    .slider {
        position: absolute;
        cursor: pointer;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(255, 255, 255, 0.1);
        transition: 0.3s;
        border-radius: 34px;
        border: 1px solid rgba(255, 255, 255, 0.1);
    }

    .slider:before {
        position: absolute;
        content: "";
        height: 18px;
        width: 18px;
        left: 4px;
        bottom: 3px;
        background-color: white;
        transition: 0.3s;
        border-radius: 50%;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
    }

    input:checked + .slider {
        background-color: var(--primary);
        border-color: rgba(255, 255, 255, 0.2);
    }

    input:focus + .slider {
        box-shadow: 0 0 0 3px rgba(88, 166, 255, 0.3);
    }

    input:checked + .slider:before {
        transform: translateX(24px);
    }

    footer {
        padding: 16px 24px 24px;
        display: flex;
        justify-content: flex-end;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
    }

    .save-btn {
        background: var(--primary);
        color: white;
        padding: 14px 32px;
        font-weight: 700;
        border-radius: 12px;
        transition:
            transform 0.2s,
            background 0.2s;
    }

    .save-btn:hover {
        background: #4a95e6;
        transform: translateY(-2px);
    }

    .save-btn:active {
        transform: translateY(0);
    }

    /* About Section Styles */
    .about-content {
        gap: 24px;
    }

    .about-hero {
        text-align: center;
        padding: 10px 0;
    }

    .about-hero h3 {
        font-size: 24px;
        color: var(--primary);
        margin: 0;
        letter-spacing: 1px;
    }

    .about-hero p {
        color: var(--text-muted);
        font-size: 14px;
        margin-top: 4px;
    }

    .about-section h4 {
        font-size: 16px;
        color: white;
        margin-bottom: 10px;
        display: flex;
        align-items: center;
        gap: 8px;
    }

    .about-section p {
        font-size: 14px;
        line-height: 1.6;
        color: #cccccc;
    }

    .about-section ul,
    .about-section ol {
        padding-left: 20px;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .about-section li {
        font-size: 13px;
        color: #bbbbbb;
        line-height: 1.5;
    }

    .about-section strong {
        color: var(--primary);
    }

    .about-section a {
        color: var(--primary);
        text-decoration: none;
    }

    .about-section a:hover {
        text-decoration: underline;
    }

    .about-footer {
        margin-top: 10px;
        padding-top: 20px;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .signature {
        font-family: "Inter", sans-serif;
        font-weight: 700;
        color: white;
        font-style: italic;
        background: linear-gradient(90deg, #58a6ff, #bc85ff);
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
    }

    .version {
        font-size: 11px;
        color: rgba(255, 255, 255, 0.3);
    }
</style>
