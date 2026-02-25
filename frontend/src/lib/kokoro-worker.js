import { KokoroTTS } from 'kokoro-js';
import { env } from '@huggingface/transformers';

// Erasmo Cardoso - Dev
// HACK: A biblioteca kokoro-js (v1.2.1) tem uma lista fixa de vozes que bloqueia o Português.
// Este monkey-patch injeta as vozes PT-BR na validação interna sem mexer no node_modules.
const ptVoices = ['pf_dora', 'pm_alex', 'pm_santa'];
const originalValidate = KokoroTTS.prototype._validate_voice;
KokoroTTS.prototype._validate_voice = function (voice) {
    if (ptVoices.includes(voice)) {
        return voice.at(0); // Retorna o prefixo 'p' para o processamento
    }
    return originalValidate.call(this, voice);
};

// Erasmo Cardoso - Dev
// Worker para processamento neural de voz local via Kokoro-82M

// Otimização do ambiente Transformers
env.allowLocalModels = true;
env.allowRemoteModels = false;
env.localModelPath = '/models/';
env.useBrowserCache = true;
env.backends.onnx.wasm.proxy = false;
// Forçando SIMD e Multi-threading agressivo para CPU fallback
env.backends.onnx.wasm.simd = true;
env.backends.onnx.wasm.numThreads = Math.min(4, navigator.hardwareConcurrency || 4);

// Erasmo Cardoso - Dev
// Interceptador de rede ultra-robusto para garantir Offline Total
const originalFetch = self.fetch;
self.fetch = async (url, options) => {
    const urlString = url.toString();

    // 1. Vozes (.bin) - Captura qualquer requisição para a pasta /voices/
    if (urlString.includes('/voices/')) {
        const voiceFile = urlString.split('/').pop();
        const localPath = `/voices/${voiceFile}`;
        console.log(`[Kokoro-Fetch] Voz: ${voiceFile} -> ${localPath}`);
        return originalFetch(localPath, options);
    }

    // 2. Modelo e Configurações (HF ou caminhos locais do motor)
    if (urlString.includes('huggingface.co') || urlString.includes('/models/kokoro/')) {
        let filePath;

        if (urlString.includes('huggingface.co')) {
            const match = urlString.match(/resolve\/main\/(.+)$/);
            filePath = match ? match[1].split('?')[0] : urlString.split('/').pop().split('?')[0];
        } else {
            filePath = urlString.split('/models/kokoro/').pop();
        }

        const localPath = `/models/kokoro/${filePath}`;
        console.log(`[Kokoro-Fetch] Ativo: ${filePath} -> ${localPath}`);
        return originalFetch(localPath, options);
    }

    return originalFetch(url, options);
};

let tts = null;

const checkHardware = async () => {
    if ('gpu' in navigator) {
        try {
            // @ts-ignore
            const adapter = await navigator.gpu.requestAdapter();
            if (adapter) {
                const info = await adapter.requestAdapterInfo?.() || { architecture: 'unknown' };
                return `WebGPU (${info.architecture})`;
            }
        } catch (e) { }
    }
    return env.backends.onnx.wasm.simd ? "CPU (WASM + SIMD)" : "CPU (WASM)";
};

self.onmessage = async (e) => {
    const { text, voice } = e.data;

    try {
        if (!tts) {
            const hw = await checkHardware();
            self.postMessage({ type: 'status', message: `Iniciando Kokoro Local (${hw})...` });

            try {
                self.postMessage({ type: 'status', message: 'Carregando modelo 82M (local q8)...' });
                tts = await KokoroTTS.from_pretrained("kokoro", {
                    dtype: "q8" // Modelo quantizado de 83MB para maior velocidade
                });
                self.postMessage({ type: 'status', message: 'Modelo Kokoro carregado.' });
            } catch (initErr) {
                console.error('Erro ao inicializar Kokoro:', initErr);
                self.postMessage({ type: 'error', message: `Erro ao carregar modelo: ${initErr.message}` });
                return;
            }
        }

        self.postMessage({ type: 'status', message: 'Sintetizando áudio (processamento neural)...' });
        const startTime = performance.now();

        // Implementação de timeout de 90s para evitar travamento infinito
        const synthesisPromise = tts.generate(text, {
            voice: voice || 'af_heart',
        });

        const timeoutPromise = new Promise((_, reject) =>
            setTimeout(() => reject(new Error("Timeout: A síntese neural demorou mais de 120s. Tente enviar textos menores.")), 120000)
        );

        self.postMessage({ type: 'status', message: 'Computando tensores neurais...' });

        // Corrida entre a síntese e o timeout
        const output = await Promise.race([synthesisPromise, timeoutPromise]);

        const endTime = performance.now();
        const duration = ((endTime - startTime) / 1000).toFixed(1);

        self.postMessage({ type: 'status', message: `Concluído em ${duration}s.` });

        // Retornar o resultado para o frontend
        self.postMessage({
            type: 'done',
            audio: output.audio,
            sampling_rate: output.sampling_rate
        });

    } catch (err) {
        console.error('Erro no Worker Kokoro:', err);
        self.postMessage({ type: 'error', message: `Falha na geração: ${err.message}` });
    }
};
