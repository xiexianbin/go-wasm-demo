<template>
  <div class="hello">
    <h1>vue3 wasm demo</h1>
    <div>
      number: <input id="num" v-model="num" /> result:
      <input id="answer" value="" />
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
    };
  },
  methods: {
    onclick() {
      console.log('this---', this.$wasm);
      console.log('fib result is', this.$wasm.exports.fib(this.num));
      this.$wasm.exports.update();
      console.log('update result is done');
    },
  },
};
</script>
