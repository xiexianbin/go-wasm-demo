<template>
  <q-page class="row items-center justify-evenly">
    <div class="row">
      number: <input id="num" v-model="num" /> result:
      <input id="answer" value="" />

      <q-btn color="white" text-color="black" label="Click" @click="onclick" />
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'PageIndex',
  data() {
    return {
      num: 10,
      result: 0,
    };
  },
  methods: {
    onclick() {
      const fibFun = this.$wasm.exports.fib as CallableFunction;
      const updateFun = this.$wasm.exports.update as CallableFunction;
      console.log('fib result is', fibFun(this.num));
      updateFun();
      console.log('update result is done');
    },
  },
});
</script>
