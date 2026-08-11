export namespace domain {
	
	export class Action {
	    ID: string;
	    Name: string;
	    Info: string;
	    Next: string[];
	
	    static createFrom(source: any = {}) {
	        return new Action(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Info = source["Info"];
	        this.Next = source["Next"];
	    }
	}
	export class BrowserConfig {
	    Width: number;
	    Height: number;
	
	    static createFrom(source: any = {}) {
	        return new BrowserConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Width = source["Width"];
	        this.Height = source["Height"];
	    }
	}
	export class GraphDiagnostic {
	    id: string;
	    nodeId?: string;
	    severity: string;
	    code: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new GraphDiagnostic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nodeId = source["nodeId"];
	        this.severity = source["severity"];
	        this.code = source["code"];
	        this.message = source["message"];
	    }
	}
	export class DiagnosticReport {
	    isValid: boolean;
	    errors: GraphDiagnostic[];
	    warnings: GraphDiagnostic[];
	
	    static createFrom(source: any = {}) {
	        return new DiagnosticReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isValid = source["isValid"];
	        this.errors = this.convertValues(source["errors"], GraphDiagnostic);
	        this.warnings = this.convertValues(source["warnings"], GraphDiagnostic);
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
	export class File {
	    Path: string;
	    Actions: Action[];
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Actions = this.convertValues(source["Actions"], Action);
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
	
	export class Project {
	    Path: string;
	    Vars: Record<string, string>;
	    Files: File[];
	    Browser: BrowserConfig;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Vars = source["Vars"];
	        this.Files = this.convertValues(source["Files"], File);
	        this.Browser = this.convertValues(source["Browser"], BrowserConfig);
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
	export class ProviderConfig {
	    Model: string;
	    APIKey: string;
	    BaseURL: string;
	
	    static createFrom(source: any = {}) {
	        return new ProviderConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Model = source["Model"];
	        this.APIKey = source["APIKey"];
	        this.BaseURL = source["BaseURL"];
	    }
	}
	export class Settings {
	    Provider: string;
	    Providers: Record<string, ProviderConfig>;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Provider = source["Provider"];
	        this.Providers = this.convertValues(source["Providers"], ProviderConfig, true);
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

export namespace testfile {
	
	export class Edge {
	    from: string;
	    to: string;
	    implicit?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Edge(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = source["from"];
	        this.to = source["to"];
	        this.implicit = source["implicit"];
	    }
	}
	export class Node {
	    id: string;
	    name: string;
	    info?: string;
	    source?: string;
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.info = source["info"];
	        this.source = source["source"];
	    }
	}
	export class Graph {
	    start: string;
	    nodes: Record<string, Node>;
	    edges: Edge[];
	
	    static createFrom(source: any = {}) {
	        return new Graph(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.start = source["start"];
	        this.nodes = this.convertValues(source["nodes"], Node, true);
	        this.edges = this.convertValues(source["edges"], Edge);
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
	
	export class View {
	    graph: Graph;
	    order?: string[];
	
	    static createFrom(source: any = {}) {
	        return new View(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.graph = this.convertValues(source["graph"], Graph);
	        this.order = source["order"];
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

