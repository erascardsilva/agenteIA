<!-- Erasmo Cardoso - Software Engineer |Electronics Technician -->
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
  let showFullVersion = false;
  let showSnapReminder = false;
  let isSnapMode = false;
  let settingsTab = "config";

  onMount(async () => {
    isSnapMode = await window.go.main.App.IsSnap();
    if (!isSnapMode) {
      showSnapReminder = true;
    }
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

      {#if isSnapMode}
        <button
          class="full-version-trigger"
          on:click={() => (showFullVersion = true)}
        >
          <svg
            viewBox="0 0 24 24"
            width="20"
            height="20"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14l-5-4.87 6.91-1.01L12 2z" />
          </svg>
          {$t("sidebar.full_version")}
        </button>
      {/if}

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

  {#if showFullVersion}
    <div class="modal-overlay full-version-overlay" on:click|self={() => (showFullVersion = false)}>
      <div class="full-version-modal glass fade-in">
        <header>
          <div class="modal-title-wrapper">
            <span class="modal-icon">🚀</span>
            <h2>{$t("context.cards.full_version.modal_title")}</h2>
          </div>
          <button class="close-btn" on:click={() => (showFullVersion = false)}>&times;</button>
        </header>

        <section class="modal-content">
          <p>{@html $t("context.cards.full_version.modal_desc")}</p>
          
          <div class="comparison-grid">
            <div class="comp-item demo">
              <span class="status-bad">✕</span>
              <strong>Demo (Snap Store)</strong>
              <p>Ambiente Isolado (Sandbox)</p>
            </div>
            <div class="comp-item full">
              <span class="status-good">✓</span>
              <strong>Full Version (GitHub)</strong>
              <p>Poder Total & Autônomo</p>
            </div>
          </div>
        </section>

        <footer>
          <button class="download-btn-premium" on:click={() => {
            window.go.main.App.OpenBrowser('https://github.com/erascardsilva/agenteIA/tree/master/build/bin');
            showFullVersion = false;
          }}>
            <svg viewBox="0 0 24 24" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4M7 10l5 5 5-5M12 15V3" />
            </svg>
            {$t("context.cards.full_version.download_btn")}
          </button>
        </footer>
      </div>
    </div>
  {/if}

  {#if showSnapReminder}
    <div class="modal-overlay" on:click|self={() => (showSnapReminder = false)}>
      <div class="full-version-modal glass fade-in">
        <header>
          <div class="modal-title-wrapper">
            <span class="modal-icon">💡</span>
            <h2>Agente IA: Full Version</h2>
          </div>
          <button class="close-btn" on:click={() => (showSnapReminder = false)}>&times;</button>
        </header>

        <section class="modal-content">
          <p style="text-align: center; font-size: 16px;">
            {@html $t("context.cards.full_version.keep_snap_reminder")}
          </p>
        </section>

        <footer>
          <button class="download-btn-premium" on:click={() => (showSnapReminder = false)}>
            OK, Entendi
          </button>
        </footer>
      </div>
    </div>
  {/if}
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

  .full-version-trigger {
    width: 100%;
    padding: 12px 16px;
    justify-content: flex-start;
    background: linear-gradient(90deg, rgba(88, 166, 255, 0.1), transparent);
    color: #58a6ff;
    font-weight: 600;
    margin-bottom: 4px;
    border-radius: 12px;
    border: 1px solid rgba(88, 166, 255, 0.1);
  }

  .full-version-trigger:hover {
    background: rgba(88, 166, 255, 0.15);
    border-color: rgba(88, 166, 255, 0.3);
    color: white;
  }

  /* Modal Styles */
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.7);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    backdrop-filter: blur(8px);
  }

  .full-version-modal {
    width: 90%;
    max-width: 500px;
    border-radius: 24px;
    overflow: hidden;
    border: 1px solid rgba(255, 255, 255, 0.1);
    box-shadow: 0 25px 50px rgba(0, 0, 0, 0.5);
  }

  .full-version-modal header {
    padding: 24px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  }

  .modal-title-wrapper {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .modal-icon {
    font-size: 24px;
  }

  .full-version-modal h2 {
    font-size: 20px;
    margin: 0;
    color: white;
  }

  .modal-content {
    padding: 24px;
    color: var(--text-muted);
    line-height: 1.6;
  }

  .comparison-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 16px;
    margin-top: 24px;
  }

  .comp-item {
    background: rgba(255, 255, 255, 0.03);
    padding: 16px;
    border-radius: 16px;
    text-align: center;
    border: 1px solid rgba(255, 255, 255, 0.05);
  }

  .status-bad { color: #f85149; font-size: 20px; display: block; margin-bottom: 8px; }
  .status-good { color: #3fb950; font-size: 20px; display: block; margin-bottom: 8px; }

  .comp-item strong { display: block; color: white; margin-bottom: 4px; }
  .comp-item p { font-size: 12px; margin: 0; }

  .full-version-modal footer {
    padding: 24px;
    display: flex;
    justify-content: center;
  }

  .download-btn-premium {
    width: 100%;
    padding: 16px;
    background: linear-gradient(135deg, #58a6ff, #1f6feb);
    color: white;
    border: none;
    border-radius: 12px;
    font-weight: 700;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    cursor: pointer;
    transition: all 0.2s;
  }

  .download-btn-premium:hover {
    transform: translateY(-2px);
    box-shadow: 0 10px 20px rgba(88, 166, 255, 0.3);
  }

  .close-btn {
    background: none;
    border: none;
    color: var(--text-muted);
    font-size: 28px;
    cursor: pointer;
  }
</style>
