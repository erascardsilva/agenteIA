<!-- Erasmo Cardoso - Software Engineer | Electronics Specialist -->
<script>
    import {
        GetSystemVoices,
        GetAvailableModels,
    } from "../../wailsjs/go/main/App";
    import { configStore } from "./store";
    import { onMount } from "svelte";
    import { t } from "./i18n";

    export let visible = false;
    export let currentTab = "config"; // 'config' ou 'about'

    // Erasmo Cardoso - Software Engineer | Electronics Specialist
    // Lógica para dropdowns customizados
    let openDropdown = null; // ID do dropdown aberto atualmente

    const labelMaps = {
        "voiceSettings.engine": {
            openai: "OpenAI Premium (API Key)",
            native: "Nativo (Navegador)",
            linux_local: "Linux (spd-say / Local)",
        },
        "voiceSettings.voiceId": {},
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

    let systemVoices = [];
    let availableModels = [];
    let loadingModels = false;

    let currentFetchId = 0;
    async function fetchModels(providerOverride = null) {
        const provider = providerOverride || $configStore.preferredProvider;
        if (provider === "ollama") {
            availableModels = [];
            return;
        }

        const fetchId = ++currentFetchId;
        loadingModels = true;
        availableModels = []; // Limpa IMEDIATAMENTE a lista anterior

        try {
            console.log(`[Fetch ${fetchId}] Buscando modelos para:`, provider);
            const models = await GetAvailableModels(provider);

            // Proteção: Só atualiza se esta for a chamada mais recente
            if (fetchId !== currentFetchId) return;

            if (models && models.length > 0) {
                availableModels = models.sort();
            } else {
                availableModels = [];
            }
        } catch (e) {
            if (fetchId === currentFetchId) {
                console.error("Erro ao buscar modelos:", e);
                availableModels = [];
            }
        } finally {
            if (fetchId === currentFetchId) {
                loadingModels = false;
            }
        }
    }

    async function loadSystemVoices() {
        try {
            const voices = await GetSystemVoices();
            if (voices) systemVoices = voices;
        } catch (e) {
            console.error("Erro ao carregar vozes do sistema:", e);
        }
    }

    onMount(async () => {
        const handleOutsideClick = () => (openDropdown = null);
        window.addEventListener("click", handleOutsideClick);

        if ($configStore.voiceSettings.engine === "linux_local") {
            await loadSystemVoices();
        }

        if ($configStore.preferredProvider !== "ollama") {
            await fetchModels();
        }

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
                        >{$t("settings.config_tab")}</button
                    >
                    <button
                        class="tab-link {currentTab === 'about'
                            ? 'active'
                            : ''}"
                        on:click={() => (currentTab = "about")}
                        >{$t("settings.about_tab")}</button
                    >
                </div>
                <button class="close-btn" on:click={() => (visible = false)}
                    >&times;</button
                >
            </header>

            {#if currentTab === "config"}
                <section class="settings-content">
                    <div class="field">
                        <label for="assistantName"
                            >{$t("settings.assistant_name")}</label
                        >
                        <input
                            id="assistantName"
                            type="text"
                            bind:value={$configStore.assistantName}
                            placeholder={$t(
                                "settings.assistant_name_placeholder",
                            )}
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
                        <label for="humor">{$t("settings.mood")}</label>
                        <div class="custom-select" id="humor">
                            <button
                                class="select-trigger"
                                on:click={(e) => toggleDropdown("humor", e)}
                            >
                                <span
                                    >{$t(
                                        "settings.moods." +
                                            $configStore.context.humor,
                                    )}</span
                                >
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
                                            {$t("settings.moods." + hum)}
                                        </button>
                                    {/each}
                                </div>
                            {/if}
                        </div>
                    </div>

                    <div class="field">
                        <label for="provider"
                            >{$t("settings.preferred_provider")}</label
                        >
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
                                            on:click={async () => {
                                                await configStore.updateProvider(
                                                    val,
                                                );
                                                fetchModels(val);
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
                        <div class="field-header">
                            <label for="model">{$t("settings.model")}</label>
                            {#if $configStore.preferredProvider !== "ollama"}
                                <button
                                    class="refresh-btn mini"
                                    on:click={() => fetchModels()}
                                    disabled={loadingModels}
                                    title="Atualizar lista de modelos"
                                >
                                    <svg
                                        width="12"
                                        height="12"
                                        viewBox="0 0 24 24"
                                        fill="none"
                                        stroke="currentColor"
                                        stroke-width="2"
                                        class:spin={loadingModels}
                                    >
                                        <path
                                            d="M23 4v6h-6M1 20v-6h6M3.51 9a9 9 0 0114.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0020.49 15"
                                        />
                                    </svg>
                                </button>
                            {/if}
                        </div>

                        {#if $configStore.preferredProvider === "ollama"}
                            <input
                                id="model"
                                type="text"
                                bind:value={$configStore.preferredModel}
                                placeholder="ex: llama3, mistral, phi3..."
                            />
                        {:else if availableModels.length > 0}
                            <select
                                id="model"
                                bind:value={$configStore.preferredModel}
                                disabled={loadingModels}
                            >
                                {#if loadingModels}
                                    <option value="" disabled
                                        >Buscando modelos...</option
                                    >
                                {:else}
                                    <option value="" disabled
                                        >-- Selecione um modelo --</option
                                    >
                                    {#each availableModels as m}
                                        <option value={m}>{m}</option>
                                    {/each}
                                {/if}
                            </select>
                        {:else}
                            <div class="model-input-wrapper">
                                <input
                                    id="model"
                                    type="text"
                                    bind:value={$configStore.preferredModel}
                                    placeholder={loadingModels
                                        ? "Buscando modelos..."
                                        : $t("settings.model_placeholder")}
                                />
                                {#if loadingModels}
                                    <div class="spinner-small"></div>
                                {/if}
                            </div>
                        {/if}
                    </div>

                    <div class="field">
                        <label for="voiceEngine"
                            >{$t("settings.voice_engine")}</label
                        >
                        <div class="custom-select" id="voiceEngine">
                            <button
                                class="select-trigger"
                                on:click={(e) => toggleDropdown("engine", e)}
                            >
                                <span
                                    >{$t(
                                        "settings.vocal_engines." +
                                            $configStore.voiceSettings.engine,
                                    )}</span
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
                                            {$t(
                                                "settings.vocal_engines." + val,
                                            )}
                                        </button>
                                    {/each}
                                </div>
                            {/if}
                        </div>
                    </div>

                    {#if $configStore.voiceSettings.engine === "openai"}
                        <div class="field">
                            <label for="voice"
                                >{$t("settings.assistant_voice")}</label
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

                    {#if $configStore.voiceSettings.engine === "linux_local"}
                        <div class="field">
                            <label for="voiceLinux"
                                >{$t("settings.system_voice")}</label
                            >
                            <select
                                id="voiceLinux"
                                bind:value={$configStore.voiceSettings.voiceId}
                            >
                                <option value=""
                                    >{$t("settings.default_system")}</option
                                >
                                {#each systemVoices as voice}
                                    <option value={voice}>{voice}</option>
                                {/each}
                            </select>
                            <button
                                class="refresh-btn"
                                on:click|preventDefault={loadSystemVoices}
                                title="Atualizar Lista"
                            >
                                <svg
                                    viewBox="0 0 24 24"
                                    width="14"
                                    height="14"
                                    fill="currentColor"
                                >
                                    <path
                                        d="M17.65 6.35C16.2 4.9 14.21 4 12 4c-4.42 0-7.99 3.58-7.99 8s3.57 8 7.99 8c3.73 0 6.84-2.55 7.73-6h-2.08c-.82 2.33-3.04 4-5.65 4-3.31 0-6-2.69-6-6s2.69-6 6-6c1.66 0 3.14.69 4.22 1.78L13 11h7V4l-2.35 2.35z"
                                    ></path>
                                </svg>
                                {$t("settings.refresh_voices")}
                            </button>
                        </div>
                    {/if}

                    <div class="field toggle-item">
                        <label for="voiceEnabled" class="toggle-label">
                            <span class="label-text"
                                >{$t("settings.voice_enabled")}</span
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
                                >{$t("settings.unlock_models")}</span
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
                        <div class="logo-icon-large">&lt; IA &gt;</div>
                        <h3>Agente IA</h3>
                        <p>
                            {$t("about.description")}
                        </p>
                    </div>

                    <div class="about-section highlight">
                        <h4>{$t("about.what_does")}</h4>
                        <p>
                            {$t("about.what_does_text")}
                        </p>
                    </div>

                    <div class="about-grid">
                        <div class="about-card">
                            <div class="card-icon">🔓</div>
                            <h5>{$t("about.cards.unlock.title")}</h5>
                            <p>
                                {$t("about.cards.unlock.desc")}
                            </p>
                        </div>
                        <div class="about-card">
                            <div class="card-icon">🌐</div>
                            <h5>{$t("about.cards.multi_llm.title")}</h5>
                            <p>
                                {$t("about.cards.multi_llm.desc")}
                            </p>
                        </div>
                        <div class="about-card">
                            <div class="card-icon">🧠</div>
                            <h5>{$t("about.cards.context.title")}</h5>
                            <p>
                                {$t("about.cards.context.desc")}
                            </p>
                        </div>
                        <div class="about-card">
                            <div class="card-icon">💻</div>
                            <h5>{$t("about.cards.total_power.title")}</h5>
                            <p>
                                {$t("about.cards.total_power.desc")}
                            </p>
                        </div>
                        <div class="about-card">
                            <div class="card-icon">🎭</div>
                            <h5>{$t("about.cards.mood_adjust.title")}</h5>
                            <p>
                                {$t("about.cards.mood_adjust.desc")}
                            </p>
                        </div>
                        <div class="about-card">
                            <div class="card-icon">📂</div>
                            <h5>{$t("about.cards.kb.title")}</h5>
                            <p>
                                {$t("about.cards.kb.desc")}
                            </p>
                        </div>
                    </div>

                    <div class="about-footer">
                        <div class="signature">Erasmo Cardoso - Software Engineer | Electronics Specialist</div>
                        <div class="version">{$t("about.version")}</div>
                    </div>
                </section>
            {/if}

            {#if currentTab === "config"}
                <footer>
                    <button class="save-btn" on:click={save}
                        >{$t("settings.save_settings")}</button
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

    .field-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 8px;
    }

    .field-header label {
        margin-bottom: 0 !important;
    }

    .refresh-btn.mini {
        padding: 4px;
        background: transparent;
        border: none;
        color: var(--primary);
        opacity: 0.7;
        margin: 0;
    }

    .refresh-btn.mini:hover {
        opacity: 1;
        background: rgba(88, 166, 255, 0.1);
    }

    .spin {
        animation: spin 1s linear infinite;
    }

    @keyframes spin {
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(360deg);
        }
    }

    .model-input-wrapper {
        position: relative;
        display: flex;
        align-items: center;
    }

    .model-input-wrapper input {
        width: 100%;
    }

    .spinner-small {
        position: absolute;
        right: 12px;
        width: 16px;
        height: 16px;
        border: 2px solid rgba(255, 255, 255, 0.1);
        border-top-color: var(--primary);
        border-radius: 50%;
        animation: spin 0.8s linear infinite;
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

    .refresh-btn {
        background: rgba(255, 255, 255, 0.05);
        border: 1px solid rgba(255, 255, 255, 0.1);
        color: var(--text-muted);
        border-radius: 8px;
        padding: 6px 12px;
        font-size: 12px;
        display: flex;
        align-items: center;
        gap: 6px;
        cursor: pointer;
        transition: all 0.2s;
        margin-top: 4px;
        width: fit-content;
    }

    .refresh-btn:hover {
        background: rgba(255, 255, 255, 0.1);
        color: var(--primary);
        border-color: var(--primary);
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

    .about-section h4 {
        color: var(--primary);
        margin-bottom: 12px;
        font-size: 18px;
    }

    .about-content {
        gap: 32px;
    }

    .about-hero {
        text-align: center;
        padding: 20px 0;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 16px;
    }

    .logo-icon-large {
        background: var(--primary);
        color: white;
        padding: 12px 20px;
        border-radius: 16px;
        font-weight: 800;
        font-size: 24px;
        box-shadow: 0 10px 30px rgba(88, 166, 255, 0.3);
    }

    .about-hero h3 {
        font-size: 28px;
        margin-bottom: 4px;
        color: white;
    }

    .about-hero p {
        color: var(--text-muted);
        font-size: 16px;
    }

    .about-section.highlight {
        background: rgba(88, 166, 255, 0.1);
        padding: 24px;
        border-radius: 20px;
        border: 1px dashed rgba(88, 166, 255, 0.3);
        margin-bottom: 10px;
    }

    .about-section h4 {
        color: var(--primary);
        margin-bottom: 12px;
        font-size: 18px;
    }

    .about-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
        gap: 16px;
    }

    .about-card {
        background: rgba(255, 255, 255, 0.03);
        padding: 20px;
        border-radius: 16px;
        border: 1px solid rgba(255, 255, 255, 0.05);
        transition: all 0.3s ease;
    }

    .about-card:hover {
        background: rgba(255, 255, 255, 0.06);
        transform: translateY(-5px);
        border-color: rgba(88, 166, 255, 0.3);
    }

    .card-icon {
        font-size: 24px;
        margin-bottom: 12px;
    }

    .about-card h5 {
        color: white;
        font-size: 15px;
        margin-bottom: 8px;
    }

    .about-card p {
        font-size: 13px;
        color: var(--text-muted);
        line-height: 1.5;
    }

    .about-footer {
        margin-top: 20px;
        padding-top: 24px;
        border-top: 1px solid rgba(255, 255, 255, 0.05);
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .signature {
        font-family: "Courier New", Courier, monospace;
        color: var(--primary);
        font-weight: 600;
        font-size: 14px;
    }

    .version {
        font-size: 12px;
        color: rgba(255, 255, 255, 0.3);
    }
</style>
