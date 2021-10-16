import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)
app.mount('#app')

import Go from './wasm_exec'

const go = new Go(); // Defined in wasm_exec.js

WebAssembly.instantiateStreaming(
  fetch("./main.wasm"), go.importObject
).then((result) => {
  go.run(result.instance)

  app.config.globalProperties.$go = {
    fibFunc: fibFunc,
  }
})
