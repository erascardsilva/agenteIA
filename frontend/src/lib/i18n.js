// Erasmo Cardoso - Software Engineer | Electronics Specialist
import { writable, derived } from 'svelte/store';
import { configStore } from './store';
import pt from './locales/pt.json';
import en from './locales/en.json';

const translations = {
    'en': en,
    'pt-BR': pt
};

export const locale = writable('en');

// Sincronizar com o configStore
configStore.subscribe(config => {
    if (config.language) {
        locale.set(config.language);
    }
});

export const t = derived(locale, ($locale) => (key, vars = {}) => {
    let text = key.split('.').reduce((obj, i) => obj ? obj[i] : null, translations[$locale]);

    if (!text) {
        // Fallback para inglês se não encontrar a chave no idioma atual
        text = key.split('.').reduce((obj, i) => obj ? obj[i] : null, translations['en']);
    }

    if (!text) return key;

    // Substituir variáveis se houver {varName}
    Object.keys(vars).forEach(k => {
        const regex = new RegExp(`{${k}}`, 'g');
        text = text.replace(regex, vars[k]);
    });

    return text;
});
