<template>
  <div class="hello">
    <h1>vue3 wasm demo</h1>
    <div>
      number: <input id="num" v-model="num" /> result:
      <input id="answer" v-model="answer" />
      <button id="btn" @click="onclick">Click</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      num: 1,
      answer: 1,
    };
  },
  methods: {
    onclick() {
      console.log('this---', this.$go);
      this.answer = this.$go.fibFunc(this.num)
      console.log('fib result is', this.answer);
    },
  },
};
</script>
