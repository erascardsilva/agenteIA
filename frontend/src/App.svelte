<!-- Erasmo Cardoso - Software Engineer | Electronics Specialist -->
<script>
  import Chat from "./lib/Chat.svelte";
  import Settings from "./lib/Settings.svelte";
  import Context from "./lib/Context.svelte";
  import { onMount } from "svelte";
  import { configStore } from "./lib/store";
  import { t } from "./lib/i18n";
  import LanguageBar from "./lib/LanguageBar.svelte";

  let currentTab = "chat";
  let showSettings = false;
  let settingsTab = "config";

  onMount(async () => {
    await configStore.load();
  });

  function toggleSettings(tab = "config") {
    settingsTab = tab;
    showSettings = !showSettings;
  }
</script>

<main>
  <aside class="sidebar glass">
    <div class="logo">
      <div class="logo-icon">&lt; IA &gt;</div>
      <h1>
        {$configStore.assistantName || $t("sidebar.default_assistant_name")}
      </h1>
    </div>

    <nav>
      <button
        class="nav-item {currentTab === 'chat' ? 'active' : ''}"
        on:click={() => (currentTab = "chat")}
      >
        <svg
          viewBox="0 0 24 24"
          width="20"
          height="20"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path
            d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"
          ></path>
        </svg>
        {$t("sidebar.chat")}
      </button>

      <button
        class="nav-item {currentTab === 'context' ? 'active' : ''}"
        on:click={() => (currentTab = "context")}
      >
        <svg
          viewBox="0 0 24 24"
          width="20"
          height="20"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path
            d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 18a8 8 0 1 1 8-8 8 8 0 0 1-8 8z"
          ></path>
          <path
            d="M12 6a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h3a1 1 0 0 0 0-2h-2V7a1 1 0 0 0-1-1z"
          ></path>
        </svg>
        {$t("sidebar.context")}
      </button>
    </nav>

    <div class="sidebar-footer">
      <button
        class="settings-trigger"
        on:click={() => toggleSettings("config")}
      >
        <svg
          viewBox="0 0 24 24"
          width="20"
          height="20"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <circle cx="12" cy="12" r="3"></circle>
          <path
            d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"
          ></path>
        </svg>
        {$t("sidebar.settings")}
      </button>

      <button class="settings-trigger" on:click={() => toggleSettings("about")}>
        <svg
          viewBox="0 0 24 24"
          width="20"
          height="20"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="12" y1="16" x2="12" y2="12"></line>
          <line x1="12" y1="8" x2="12.01" y2="8"></line>
        </svg>
        {$t("sidebar.about")}
      </button>

      <button
        class="support-trigger"
        on:click={() => window.go.main.App.OpenBrowser('https://www.paypal.com/ncp/payment/8V6WQCGN6HDCQ')}
        title="Apoie o Projeto"
      >
        <svg
          viewBox="0 0 24 24"
          width="20"
          height="20"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
        </svg>
        Apoie o Projeto
      </button>
    </div>
  </aside>

  <section class="content">
    <LanguageBar />
    {#if currentTab === "chat"}
      <Chat />
    {:else if currentTab === "context"}
      <Context />
    {/if}
  </section>

  <Settings bind:visible={showSettings} bind:currentTab={settingsTab} />
</main>

<style>
  main {
    display: flex;
    width: 100vw;
    height: 100vh;
    background: var(--bg-color);
  }

  .sidebar {
    width: 260px;
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 20px;
    border-right: 1px solid var(--border-color);
    z-index: 10;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 40px;
  }

  .logo-icon {
    background: var(--primary);
    color: white;
    padding: 6px 10px;
    border-radius: 8px;
    font-weight: 700;
    font-size: 14px;
  }

  .logo h1 {
    font-size: 18px;
    font-weight: 600;
    color: white;
  }

  nav {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .nav-item {
    width: 100%;
    padding: 12px 16px;
    justify-content: flex-start;
    background: transparent;
    color: var(--text-muted);
  }

  .nav-item:hover {
    background: rgba(255, 255, 255, 0.05);
    color: white;
  }

  .nav-item.active {
    background: rgba(88, 166, 255, 0.1);
    color: var(--primary);
  }

  .sidebar-footer {
    padding-top: 10px;
    border-top: 1px solid var(--border-color);
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .settings-trigger {
    width: 100%;
    padding: 12px 16px;
    justify-content: flex-start;
    background: transparent;
    color: var(--text-muted);
  }

  .settings-trigger:hover {
    background: rgba(255, 255, 255, 0.05);
    color: white;
  }

  .support-trigger {
    width: 100%;
    padding: 12px 16px;
    justify-content: flex-start;
    background: transparent;
    color: var(--text-muted);
    border: none;
    display: flex;
    align-items: center;
    gap: 12px;
    cursor: pointer;
    font-size: 14px;
    transition: all 0.2s;
  }

  .support-trigger:hover {
    background: rgba(0, 112, 186, 0.1);
    color: #58a6ff;
  }

  .support-trigger svg {
    color: #e13238; /* Cor de coração sutil */
  }

  .content {
    flex: 1;
    height: 100%;
    display: flex;
    flex-direction: column;
    background: radial-gradient(
      circle at top right,
      rgba(88, 166, 255, 0.05),
      transparent
    );
  }
</style>
