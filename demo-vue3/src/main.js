import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)
app.mount('#app')

import Go from './wasm_exec'

const go = new Go(); // Defined in wasm_exec.js

const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
	let response = undefined;

	if (WebAssembly.instantiateStreaming) {
		// Fetch the module, and instantiate it as it is downloading
		response = await WebAssembly.instantiateStreaming(fetch(wasmModuleUrl), importObject);
	} else {
		const fetchAndInstantiateTask = async () => {
			const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
				response.arrayBuffer()
			);
			return WebAssembly.instantiate(wasmArrayBuffer, importObject);
		};
		response = await fetchAndInstantiateTask();
	}
	return response;
}

let wasm;
const wasmLoad = async () => {
  const wasmModule = await wasmBrowserInstantiate("./main.wasm", go.importObject);
  wasm = wasmModule.instance;
  go.run(wasm);

	app.config.globalProperties.$wasm = wasm
  console.log('WASM Complated.')
};
console.log('WASM Loading...')
console.log('main', wasm)
wasmLoad()
