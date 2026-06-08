import { WorkerEntrypoint } from "cloudflare:workers";
import app from "./app.wasm";
import "./wasm_exec.js";

globalThis.tryCatch = (o, fn, args) => {
    try {
        if (fn) {
            return { data: o[fn](...args) };
        }

        return { data: o(...args) };
    } catch (err) {
        if (!(err instanceof Error)) {
            if (err instanceof Object) {
                err = JSON.stringify(err);
            }
            err = new Error(err || "no error message");
        }
        return { error: err };
    }
};

let initiliazed = false;

let go = new Go();
let instance = new WebAssembly.Instance(app, go.importObject);

function init() {
    if (!initiliazed) {
        go.run(instance).finally(() => {
            initiliazed = false;
            instance = new WebAssembly.Instance(app, go.importObject);
        });
        initiliazed = true;
    }

    if (go.exited) {
        go = new Go();
        go.run(instance).finally(() => {
            instance = new WebAssembly.Instance(app, go.importObject);
        });
    }
}

export default class extends WorkerEntrypoint {
    constructor(ctx, env) {
        super(ctx, env);

        // Required to init tne `env` and `ctx` variables
        // and populate this class's prototype with RPC stubs
        globalThis.workerapp = this;
        init();
    }

    async fetch(request: Request): Response | Promise<Response> {
        let response = await cf.fetch(request, this.env, this.ctx) as Response;
        response.headers.set('x-datacenter', request.cf?.colo || 'Unknown')

        return response
    }
}

