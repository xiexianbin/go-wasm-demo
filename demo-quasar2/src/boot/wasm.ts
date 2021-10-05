/* eslint-disable @typescript-eslint/no-misused-promises */
import { boot } from 'quasar/wrappers'

import Go from 'assets/wasm_exec'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $wasm: WebAssembly.Instance;
  }
}

const go = new Go(); // Defined in wasm_exec.js

const wasmBrowserInstantiate = async (wasmModuleUrl:string, importObject:WebAssembly.Imports) => {
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

const wasmLoad = new Promise(async (reslove) => {
  const wasmModule = await wasmBrowserInstantiate('./main.wasm', go.importObject);
  const _wasm = wasmModule.instance;
  void go.run(_wasm);

  console.log('WASM Complated.')
  reslove(_wasm)
});

// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(async ({ app }) => {
  app.config.globalProperties.$wasm = await wasmLoad
})
