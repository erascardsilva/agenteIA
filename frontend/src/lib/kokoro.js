// Erasmo Cardoso - Dev
export class KokoroManager {
    constructor() {
        this.worker = new Worker(new URL('./kokoro-worker.js', import.meta.url), { type: 'module' });
        this.onStatus = null;
        this.onError = null;

        this.worker.onmessage = (e) => {
            if (e.data.type === 'status' && this.onStatus) {
                this.onStatus(e.data.message);
            } else if (e.data.type === 'error' && this.onError) {
                this.onError(e.data.message);
            }
        };
    }

    async speak(text, voice) {
        return new Promise((resolve, reject) => {
            const handleMessage = (e) => {
                if (e.data.type === 'done') {
                    this.worker.removeEventListener('message', handleMessage);
                    resolve({ audio: e.data.audio, sampling_rate: e.data.sampling_rate });
                } else if (e.data.type === 'error') {
                    this.worker.removeEventListener('message', handleMessage);
                    reject(e.data.message);
                }
            };
            this.worker.addEventListener('message', handleMessage);
            this.worker.postMessage({ text, voice });
        });
    }
}

export const kokoro = new KokoroManager();
