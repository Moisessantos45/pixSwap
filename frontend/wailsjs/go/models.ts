export namespace main {
	
	export class FileData {
	    id: number;
	    base64: string;
	    mimeType: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new FileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.base64 = source["base64"];
	        this.mimeType = source["mimeType"];
	        this.path = source["path"];
	    }
	}

}

