/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-misused-promises */
import { boot } from 'quasar/wrappers'
import Go from '../assets/wasm_exec'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $wasm: Record<string, CallableFunction>;
  }
}

const go = new Go(); // Defined in wasm_exec.js

// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(async ({ app }) => {
  // something to do
  void await WebAssembly.instantiateStreaming(
    fetch('./main.wasm'), go.importObject
  ).then((result) => {
    const wasm = result.instance
    void go.run(wasm)
    app.config.globalProperties.$wasm = {
      fibFunc: window.fibFunc
    }
  })
})
