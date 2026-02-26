<!-- Erasmo Cardoso - Dev -->
<script>
    import { configStore } from "./store";
    import { t } from "./i18n";

    async function updateHumor(h) {
        $configStore.context.humor = h;
        await configStore.save($configStore);
    }

    async function toggleSetting() {
        await configStore.save($configStore);
    }
</script>

<div class="context-container fade-in">
    <header>
        <h2>{$t("context.header.title")}</h2>
        <p>
            {$t("context.header.subtitle")}
        </p>
    </header>

    <div class="context-grid">
        <div class="card glass">
            <h3>{$t("context.cards.mood.title")}</h3>
            <div class="humor-options">
                {#each ["Prestativo", "Sarcástico", "Profissional", "Amigável", "Direto", "Zombeteiro"] as h}
                    <button
                        class="humor-btn {$configStore.context.humor === h
                            ? 'active'
                            : ''}"
                        on:click={() => updateHumor(h)}
                    >
                        {$t("settings.moods." + h)}
                    </button>
                {/each}
            </div>
        </div>

        <div class="card glass">
            <h3>{$t("context.cards.security.title")}</h3>
            <div class="toggle-row">
                <div class="info">
                    <strong>{$t("context.cards.security.unlock_title")}</strong>
                    <span>{$t("context.cards.security.unlock_desc")}</span>
                </div>
                <input
                    type="checkbox"
                    bind:checked={$configStore.context.unlockModels}
                    on:change={toggleSetting}
                />
            </div>

            <div class="toggle-row">
                <div class="info">
                    <strong>{$t("context.cards.security.process_title")}</strong
                    >
                    <span>{$t("context.cards.security.process_desc")}</span>
                </div>
                <input
                    type="checkbox"
                    bind:checked={$configStore.context.systemProcess}
                    on:change={toggleSetting}
                />
            </div>
        </div>

        <div class="card glass agent-profile">
            <h3>{$t("context.cards.profile.title")}</h3>
            <div class="form-grid">
                <div class="input-group">
                    <label for="userName"
                        >{$t("context.cards.profile.name")}</label
                    >
                    <input
                        id="userName"
                        type="text"
                        bind:value={$configStore.context.userNome}
                        on:blur={toggleSetting}
                        placeholder={$t(
                            "context.cards.profile.name_placeholder",
                        )}
                    />
                </div>
                <div class="input-group">
                    <label for="userProfession"
                        >{$t("context.cards.profile.profession")}</label
                    >
                    <input
                        id="userProfession"
                        type="text"
                        bind:value={$configStore.context.userProfissao}
                        on:blur={toggleSetting}
                        placeholder={$t(
                            "context.cards.profile.profession_placeholder",
                        )}
                    />
                </div>
                <div class="input-group">
                    <label for="userAge"
                        >{$t("context.cards.profile.age")}</label
                    >
                    <input
                        id="userAge"
                        type="text"
                        bind:value={$configStore.context.userIdade}
                        on:blur={toggleSetting}
                        placeholder={$t(
                            "context.cards.profile.age_placeholder",
                        )}
                    />
                </div>
            </div>
            <div class="input-group">
                <label for="userFunctions"
                    >{$t("context.cards.profile.functions")}</label
                >
                <input
                    id="userFunctions"
                    type="text"
                    bind:value={$configStore.context.funcoes}
                    on:blur={toggleSetting}
                    placeholder={$t(
                        "context.cards.profile.functions_placeholder",
                    )}
                />
            </div>
            <div class="input-group">
                <label for="userRules"
                    >{$t("context.cards.profile.rules")}</label
                >
                <textarea
                    id="userRules"
                    bind:value={$configStore.context.regras}
                    on:blur={toggleSetting}
                    placeholder={$t("context.cards.profile.rules_placeholder")}
                ></textarea>
            </div>
        </div>

        <div class="card glass info-card">
            <h3>{$t("context.cards.tips.title")}</h3>
            <p>
                {$t("context.cards.tips.text")}
            </p>
        </div>
    </div>
</div>

<style>
    .context-container {
        padding: 40px;
        display: flex;
        flex-direction: column;
        gap: 32px;
        max-width: 900px;
        margin: 0 auto;
    }

    header h2 {
        font-size: 28px;
        margin-bottom: 8px;
    }
    header p {
        color: var(--text-muted);
    }

    .context-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 24px;
    }

    .card {
        padding: 24px;
        border-radius: 20px;
        display: flex;
        flex-direction: column;
        gap: 20px;
    }

    .card h3 {
        font-size: 18px;
        color: var(--primary);
    }

    .humor-options {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 10px;
    }

    .humor-btn {
        background: rgba(255, 255, 255, 0.05);
        color: var(--text-main);
        padding: 12px;
        font-size: 14px;
    }

    .humor-btn.active {
        background: var(--primary);
        color: white;
    }

    .toggle-row {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 20px;
    }

    .info {
        display: flex;
        flex-direction: column;
        gap: 4px;
    }
    .info strong {
        font-size: 15px;
    }
    .info span {
        font-size: 12px;
        color: var(--text-muted);
    }

    .agent-profile {
        grid-column: span 2;
        gap: 16px;
    }

    .form-grid {
        display: grid;
        grid-template-columns: 2fr 2fr 1fr;
        gap: 16px;
    }

    .input-group {
        display: flex;
        flex-direction: column;
        gap: 6px;
    }

    .input-group label {
        font-size: 13px;
        color: var(--text-muted);
    }

    .input-group input,
    .input-group textarea {
        background: rgba(255, 255, 255, 0.05);
        border: 1px solid var(--border-color);
        border-radius: 8px;
        padding: 10px;
        color: white;
        font-size: 14px;
    }

    .input-group textarea {
        min-height: 100px;
        resize: vertical;
    }

    .info-card {
        grid-column: span 2;
        background: rgba(88, 166, 255, 0.05);
        border: 1px dashed var(--primary);
    }

    .info-card p {
        font-size: 14px;
        line-height: 1.6;
    }
</style>
