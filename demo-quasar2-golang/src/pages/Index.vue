<template>
  <q-page class="row items-center justify-evenly">
    <div class="row">
      number: <input id="num" v-model="num" /> result:
      <input id="answer" v-model="answer" />

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
      num: 1,
      answer: 0,
    };
  },
  methods: {
    onclick() {
      // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
      this.answer = this.$wasm.fibFunc(this.num);
      console.log('fib result is', this.answer);
      console.log('update result is done');
    },
  },
});
</script>
