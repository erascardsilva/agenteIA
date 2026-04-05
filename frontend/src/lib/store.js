// Erasmo Cardoso - Software Engineer | Electronics Specialist
import { writable } from 'svelte/store';
import { GetConfig, SaveConfig } from '../../wailsjs/go/main/App';

function createConfigStore() {
    const { subscribe, set, update } = writable({
        assistantName: "Jarvis",
        apiKeys: { groq: "", gemini: "", openai: "", deepseek: "", openrouter: "" },
        preferredProvider: "groq",
        preferredModel: "",
        voiceSettings: { enabled: false, voiceId: "alloy", engine: "openai" },
        context: {
            humor: "Prestativo",
            unlockModels: false,
            systemProcess: false,
            userNome: "",
            userProfissao: "",
            userIdade: "",
            funcoes: "",
            regras: ""
        },
        language: "en"
    });

    return {
        subscribe,
        set,
        update,
        load: async () => {
            const cfg = await GetConfig();
            if (cfg) {
                update(current => {
                    return {
                        ...current,
                        ...cfg,
                        apiKeys: { ...current.apiKeys, ...(cfg.apiKeys || {}) },
                        context: { ...current.context, ...(cfg.context || {}) },
                        voiceSettings: { ...current.voiceSettings, ...(cfg.voiceSettings || {}) }
                    };
                });
            }
        },
        save: async (newConfig) => {
            set(newConfig);
            await SaveConfig(newConfig);
        },
        updateField: async (path, value) => {
            let updatedConfig;
            update(current => {
                const keys = path.split('.');
                let target = current;
                for (let i = 0; i < keys.length - 1; i++) {
                    target = target[keys[i]];
                }
                target[keys[keys.length - 1]] = value;
                updatedConfig = { ...current };
                return updatedConfig;
            });
            await SaveConfig(updatedConfig);
        },
        updateProvider: async (newProvider) => {
            let updatedConfig;
            update(current => {
                current.preferredProvider = newProvider;
                current.preferredModel = ""; // Limpa o modelo atômicamente
                updatedConfig = { ...current };
                return updatedConfig;
            });
            await SaveConfig(updatedConfig);
        }
    };
}

export const configStore = createConfigStore();
