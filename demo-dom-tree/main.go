package main

import (
    "strconv"
    "syscall/js"
)

var (
    dom = js.Global().Get("document")
    numElt = dom.Call("getElementById", "num")
    answerElt = dom.Call("getElementById", "answer")
    btnElt = js.Global().Get("btn")
)

// https://leetcode-cn.com/submissions/detail/41908622/
func fib(N int) int {
    if N < 2{
        return N
    }
    cache := make([]int, N+1)
    cache[0], cache[1] = 0, 1
    for i := 2; i <= N; i++ {
        cache[i] = cache[i-1] + cache[i-2]
    }
    return cache[N]
}

func fibFunc(this js.Value, args []js.Value) interface{} {
    v := numElt.Get("value")
    if num, err := strconv.Atoi(v.String()); err == nil {
        answerElt.Set("innerHTML", js.ValueOf(fib(num)))
    }
    return nil
}

func main() {
    done := make(chan int, 0)
    btnElt.Call("addEventListener", "click", js.FuncOf(fibFunc))
    <-done
}
