<!-- Erasmo Cardoso - Software Engineer |Electronics Technician -->
<script>
    import { configStore } from "./store";
    import { locale } from "./i18n";

    async function setLanguage(lang) {
        if ($configStore.language === lang) return;
        await configStore.updateField("language", lang);
    }
</script>

<div class="lang-bar glass fade-in">
    <button
        class="lang-btn"
        class:active={$locale === "en"}
        on:click={() => setLanguage("en")}
    >
        <span class="flag">🇺🇸</span>
        <span class="label">English</span>
        {#if $locale === "en"}
            <div class="indicator"></div>
        {/if}
    </button>

    <div class="divider"></div>

    <button
        class="lang-btn"
        class:active={$locale === "pt-BR"}
        on:click={() => setLanguage("pt-BR")}
    >
        <span class="flag">🇧🇷</span>
        <span class="label">Português</span>
        {#if $locale === "pt-BR"}
            <div class="indicator"></div>
        {/if}
    </button>
</div>

<style>
    .lang-bar {
        display: flex;
        align-items: center;
        gap: 4px;
        padding: 4px;
        border-radius: 12px;
        background: rgba(255, 255, 255, 0.03);
        border: 1px solid rgba(255, 255, 255, 0.08);
        width: fit-content;
        margin: 10px 20px 0 auto; /* Alinhado à direita no topo */
    }

    .lang-btn {
        display: flex;
        align-items: center;
        gap: 8px;
        padding: 6px 12px;
        border-radius: 8px;
        background: transparent;
        border: none;
        color: var(--text-muted);
        font-size: 13px;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        position: relative;
    }

    .lang-btn:hover {
        background: rgba(255, 255, 255, 0.05);
        color: white;
    }

    .lang-btn.active {
        color: var(--primary);
        background: rgba(88, 166, 255, 0.1);
    }

    .flag {
        font-size: 14px;
    }

    .divider {
        width: 1px;
        height: 16px;
        background: rgba(255, 255, 255, 0.1);
        margin: 0 4px;
    }

    .indicator {
        position: absolute;
        bottom: -2px;
        left: 50%;
        transform: translateX(-50%);
        width: 12px;
        height: 2px;
        background: var(--primary);
        border-radius: 2px;
        box-shadow: 0 0 8px var(--primary);
        animation: slideIn 0.3s ease-out;
    }

    @keyframes slideIn {
        from {
            width: 0;
            opacity: 0;
        }
        to {
            width: 12px;
            opacity: 1;
        }
    }

    @keyframes fade-in {
        from {
            opacity: 0;
            transform: translateY(-5px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .fade-in {
        animation: fade-in 0.5s ease-out;
    }
</style>
