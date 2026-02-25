<!-- Erasmo Cardoso - Dev -->
<script>
  import Chat from "./lib/Chat.svelte";
  import Settings from "./lib/Settings.svelte";
  import Context from "./lib/Context.svelte";
  import { onMount } from "svelte";
  import { configStore } from "./lib/store";

  let currentTab = "chat";
  let showSettings = false;

  onMount(async () => {
    await configStore.load();
  });

  function toggleSettings() {
    showSettings = !showSettings;
  }
</script>

<main>
  <aside class="sidebar glass">
    <div class="logo">
      <div class="logo-icon">&lt; IA &gt;</div>
      <h1>{$configStore.assistantName || "Assistente"}</h1>
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
        Chat
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
        Contexto
      </button>
    </nav>

    <div class="sidebar-footer">
      <button class="settings-trigger" on:click={toggleSettings}>
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
        Configurações
      </button>
    </div>
  </aside>

  <section class="content">
    {#if currentTab === "chat"}
      <Chat />
    {:else if currentTab === "context"}
      <Context />
    {/if}
  </section>

  <Settings bind:visible={showSettings} />
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
    padding-top: 20px;
    border-top: 1px solid var(--border-color);
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
