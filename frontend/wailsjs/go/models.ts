export namespace config {
	
	export class ContextSettings {
	    humor: string;
	    unlockModels: boolean;
	    systemProcess: boolean;
	    userNome: string;
	    userProfissao: string;
	    userIdade: string;
	    funcoes: string;
	    regras: string;
	    autonomousMode: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ContextSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.humor = source["humor"];
	        this.unlockModels = source["unlockModels"];
	        this.systemProcess = source["systemProcess"];
	        this.userNome = source["userNome"];
	        this.userProfissao = source["userProfissao"];
	        this.userIdade = source["userIdade"];
	        this.funcoes = source["funcoes"];
	        this.regras = source["regras"];
	        this.autonomousMode = source["autonomousMode"];
	    }
	}
	export class VoiceSettings {
	    engine: string;
	    voiceId: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new VoiceSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.engine = source["engine"];
	        this.voiceId = source["voiceId"];
	        this.enabled = source["enabled"];
	    }
	}
	export class Config {
	    assistantName: string;
	    apiKeys: Record<string, string>;
	    preferredProvider: string;
	    preferredModel: string;
	    voiceSettings: VoiceSettings;
	    context: ContextSettings;
	    language: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assistantName = source["assistantName"];
	        this.apiKeys = source["apiKeys"];
	        this.preferredProvider = source["preferredProvider"];
	        this.preferredModel = source["preferredModel"];
	        this.voiceSettings = this.convertValues(source["voiceSettings"], VoiceSettings);
	        this.context = this.convertValues(source["context"], ContextSettings);
	        this.language = source["language"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

