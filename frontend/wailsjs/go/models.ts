export namespace main {
	
	export class Preset {
	    name: string;
	    before: string;
	    after: string;
	
	    static createFrom(source: any = {}) {
	        return new Preset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.before = source["before"];
	        this.after = source["after"];
	    }
	}
	export class Config {
	    supportedFormats: string[];
	    presets: Preset[];
	    locale: string;
	    theme: string;
	    compressTargetBytes: number;
	    closeBehavior: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.supportedFormats = source["supportedFormats"];
	        this.presets = this.convertValues(source["presets"], Preset);
	        this.locale = source["locale"];
	        this.theme = source["theme"];
	        this.compressTargetBytes = source["compressTargetBytes"];
	        this.closeBehavior = source["closeBehavior"];
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
	export class ImageFileInfo {
	    name: string;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new ImageFileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.size = source["size"];
	    }
	}

}

