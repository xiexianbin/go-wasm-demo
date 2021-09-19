package main

import (
    "syscall/js"
    "time"
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
    callback := args[len(args)-1]
    go func() {
        time.Sleep(3 * time.Second)
        v := fib(args[0].Int())
        callback.Invoke(v)
    }()

    js.Global().Get("answer").Set("innerHTML", "Sleep 3s...")
    return nil
}

func main() {
    done := make(chan int, 0)
    js.Global().Set("fibFunc", js.FuncOf(fibFunc))
    <-done
}
