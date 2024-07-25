export namespace types {
	
	export class Message {
	    id: string;
	    content: string;
	    userType: string;
	
	    static createFrom(source: any = {}) {
	        return new Message(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.content = source["content"];
	        this.userType = source["userType"];
	    }
	}

}

